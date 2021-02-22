$(function () {
  if (!window.EventSource) {
    alert("No EventSource!");
  }

  let $chatlog = $("#chat-log");
  let $chatmsg = $("#chat-msg");

  let isBlank = function (string) {
    return string == null || string.trim() === "";
  };

  let username;
  while (isBlank(username)) {
    username = prompt("What is your name");
    if (!isBlank(username)) {
      $("#user-name").html("<b>" + username + "</b>");
    }
  }

  $("#input-form").on("submit", e => {
    e.preventDefault();
    $.post("/messages", {
      msg: $chatmsg.val(),
      name: username,
    });
    $chatmsg.val("");
    $chatmsg.focus();
  });

  var addMessage = data => {
    var text = "";
    if (!isBlank(data.name)) {
      text = "<strong> " + data.name + ": </strong>";
    }
    text += data.msg;
    $chatlog.prepend("<div><span>" + text + "</span></div>");
  };

  var es = new EventSource("/stream");
  es.onopen = e => $.post("users/", { name: username });
  es.onmessage = e => {
    let msg = JSON.parse(e.data);
    addMessage(msg);
  };

  window.onbeforeunload = () => {
    $.ajax({
      url: "/users?username=" + username,
      type: "DELETE",
    });
    es.close();
  };
});
