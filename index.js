var express = require('express');
var app = express();
var path = require('path');
var server = require('http').createServer(app);
var io = require('socket.io')(server);
var port = process.env.PORT || 8080;

server.listen(port, () => {
  console.log('Server listening at port %d', port);
});

// Routing
app.use(express.static(path.join(__dirname, 'public')));

// Chatroom
var numUsers = 0;

io.on('connection', (socket) => {
  var addedUser = false;

  // 发送消息
  socket.on('new message', (data) => {
    socket.broadcast.emit('new message', {
      username: socket.username,
      message: data
    });
  });

  // 用户进入聊天室
  socket.on('add user', (username) => {
    if (addedUser) return;
    socket.username = username;
    ++numUsers;
    addedUser = true;
    socket.emit('login', {
      numUsers: numUsers
    });
    socket.broadcast.emit('user joined', {
      username: socket.username,
      numUsers: numUsers
    });
  });

  // 用户正在输入
  socket.on('typing', () => {
    socket.broadcast.emit('typing', {
      username: socket.username
    });
  });

  // 用户停止输入
  socket.on('stop typing', () => {
    socket.broadcast.emit('stop typing', {
      username: socket.username
    });
  });

  // 用户离开房间
  socket.on('disconnect', () => {
    if (addedUser) {
      --numUsers;
      socket.broadcast.emit('user left', {
        username: socket.username,
        numUsers: numUsers
      });
    }
  });
});
