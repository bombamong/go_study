package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("No Users", string(data))

	body := `{"first_name": "mimi",
			 "last_name" : "mi", 
			 "email" : "something@something.com"}`
	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	body = `{"first_name": "momo",
			 "last_name" : "mi", 
			 "email" : "something@something.com"}`
	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(2, len(users))

	// assert.Equal("Get UserInfo by /users/{id}", string(data))
}
func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(string(data), "No User ID:89")

	resp, err = http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(string(data), "No User ID:56")
}
func TestCreateUserInfo(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	body := `{"first_name": "bombu",
			 "last_name" : "mi", 
			 "email" : "something@something.com"}`
	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)

}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// how to delete
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	// create user
	body := `{"first_name": "bombu",
	"last_name" : "mi", 
	"email" : "something@something.com"}`
	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// erase user
	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User ID:1")
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	//update user when user does not exist
	body := `{
			"id": 1,
			"first_name": "KIMCHI",
			"last_name" : "MANDU", 
			"email" : "KIMCHI@MANDU.YUM"
			}`
	req, _ := http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(body))
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID:1")

	//create user
	body = `{
			"first_name": "bombu",
			"last_name" : "mi", 
			"email" : "something@something.com"
			}`
	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	json.NewDecoder(resp.Body).Decode(user)

	// update user when user exists
	body = fmt.Sprintf(`{"id" : %d, "first_name": "KIMCHI"}`, user.ID)
	req, _ = http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(body))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	//compare if user data is updated
	updatedUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updatedUser)
	assert.NoError(err)
	assert.Equal(updatedUser.ID, user.ID)
	assert.NotEqual(updatedUser.FirstName, user.FirstName)
	assert.Equal(updatedUser.LastName, user.LastName)
	assert.Equal(updatedUser.Email, user.Email)

}
