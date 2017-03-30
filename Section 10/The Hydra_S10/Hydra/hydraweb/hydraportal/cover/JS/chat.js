$(function() {
  var socket = null;
  var msgBox = $("#chatbox textarea");
  var label = $("#chatbox label");
  var messages = $("#messages");
  $("#chatbox").submit(function() {
    if (!msgBox.val()) return false;
    if (!socket) {
      alert("Error: There is no socket connection.");
      return false;
    }
    socket.send(label.text() + " " + msgBox.val() + "\n");
    msgBox.val("");
    return false;
  });
  if (!window["WebSocket"]) {
    alert("Error: Your browser does not support web sockets.")
  } else {
    if (window.location.protocol == "https:"){
      socket = new WebSocket("wss://localhost:8062/ChatRoom/");
    } else {
      socket = new WebSocket("ws://localhost:8061/ChatRoom/");
    }
    socket.onclose = function() {
      alert("Connection has been closed.");
    }
    socket.onmessage = function(e) {
      messages.append(
        $("<li>").append(
          e.data
        ));
    }
  }
});
