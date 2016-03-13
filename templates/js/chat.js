// $(function() {
            //     console.log("Loaded")
            //     var socket = null;
            //     var msgbox = $("#chatbox textarea");
            //     var messages = $("messages");
            //     $("#chatbox").submit(function() {
            //         if (!msgBox.val() return false);
            //         if (!socket) {
            //             alert("Error, there is no socket connection")
            //             return false
            //         };

            //         socket.send(msgBox.val());
            //         msgbox.val("");
            //         return false
            //     });

            //     if (!window["WebSocket"]) {
            //         alert("Error, Your browser does not support WebSockets")
            //     } else {
            //         socket = new WebSocket("ws://localhost:8080/room");
            //         socket.onclose = function() {
            //             alert("The socket has been closed")
            //         }
            //         socket.onmessage = function(e) {
            //             messages.append($("<li>").text(e.data))
            //         }
            //     }
            // });
