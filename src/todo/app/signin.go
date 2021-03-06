package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserID struct {
	ID            string `json:"ID"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

// 구글 oAuth 설정
var googleOauthConfig = oauth2.Config{
	// oAuth 요청 후 응답 받을 주소
	RedirectURL: "http://localhost:3000/auth/google/callback",
	// 요청자 아이디
	ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
	// 요청자 비밀번호
	ClientSecret: os.Getenv("GOOGLE_SECRET"),
	// 요청하고자 하는 것들, 여기선 유저 이메일 정보
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
	// 토큰 받는 주소, 제공자마다 다르고, 이번엔 구글 정보 요청 중이니 구글 endpoint 사용
	Endpoint: google.Endpoint,
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	// Cookie 에 특수 문자 추가
	expiration := time.Now().Add(1 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, cookie)
	// Cookie에 포함된 동일한 문자열 스테이트로 반환
	return state
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	// CSRF 방지하기 위해 state 생성 - generateStateOauthCookie 참조
	state := generateStateOauthCookie(w)
	// state 외에 config 내용이 포함된 URL 반환
	url := googleOauthConfig.AuthCodeURL(state)
	// 위에서 조립한 URL로 잠시 이동
	// 구글 기본 redirect auth url = https://accounts.google.com/o/oauth2/auth
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

const oAuthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleUserInfo(code string) ([]byte, error) {
	// Exchange 함수로 구글에서 받은 코드를 토큰으로 변경
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to Exchange %s", err.Error())
	}
	// 토큰을 포함해서 제공자 API로 유저 정보 요청
	resp, err := http.Get(oAuthGoogleURLAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get userInfo %s", err.Error())
	}
	// 유저 정보 []byte 로 반환
	return ioutil.ReadAll(resp.Body)
}

func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	// request 쿠키에 포함 시켰던 state 확인
	oauthstate, _ := r.Cookie("oauthstate")
	if r.FormValue("state") != oauthstate.Value {
		errMsg := fmt.Sprintf("invalid google oauth state cookie: %s state: %s", oauthstate.Value, r.FormValue("state"))
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	// state 가 같으면 유저 정보 요청
	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Store Id info into session cookie
	var userInfo GoogleUserID
	err = json.Unmarshal(data, &userInfo)

	session, _ := store.Get(r, "session")
	// Set some session values.
	session.Values["id"] = userInfo.ID
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//
	// 유저 정보 출력
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
