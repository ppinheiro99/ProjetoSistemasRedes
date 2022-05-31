var express = require("express");
var socket = require('socket.io');
var cors = require('cors');
const { addUser, getUser, deleteUser, getUsers } = require('./users')

var app = express();
app.use(cors());
var index= 0;

// Initialize our websocket server on port 5000
// var server = app.listen(3000, () => {
//   console.log("started on port 3000");
// });
var server = app.listen(5000, () => {
  console.log("started on port 5000");
});


let userList = []
let socketID = []
let room = []

var io = socket(server);
io.on("connection", (socket) => {

  let userName = socket.handshake.query.userName
  let userID = socket.handshake.query.id

  socketID[userID] = socket.id // ArrayList de userId com os respectivos socketsId's
  let indexDelete = index
  console.log("user connect " + userName + "socket: " + socketID[userID])
  
  addUser(userName,userID)

  socket.broadcast.emit('user-list', [...userList])
  socket.emit('user-list', [...userList])

  console.log(userList)

  io.emit("user connect",userList)

  socket.on("message", (msg) =>{ 
    console.log("Id sender: " + userID);
    console.log("Id received: " + msg.receivedID);
    console.log("Data Received: " + msg.data);
    // Mensagem enviada para o utilizador que pretendemos enviar msg
    io.to(socketID[msg.receivedID]).emit("message",{
      message: msg.data,
      userName:userName,
      sender: userID,
      received:msg.receivedID,
      mine: false
    })
    // Mensagem enviada para ele próprio
    io.to(socketID[userID]).emit("message",{
      message: msg.data,
      userName:userName,
      sender: userID,
      received:msg.receivedID,
      mine: true
    })

  });

  socket.on('login', ({ name, room }, callback) => {
    const { user, error } = addUser(socket.id, name, room)
    if (error) return callback(error)
    socket.join(user.room)
    socket.in(room).emit('notification', { title: 'Someone\'s here', description: `${user.name} just entered the room` })
    io.in(room).emit('users', getUsers(room))
    callback()
})

socket.on('sendGroupMessage', message => {
    const user = getUser(socket.id)
    io.in(user.room).emit('message', { user: user.name, text: message });
})

// socket.on("disconnect", () => {
//     console.log("User disconnected");
//     const user = deleteUser(socket.id)
//     if (user) {
//         io.in(user.room).emit('notification', { title: 'Someone just left', description: `${user.name} just left the room` })
//         io.in(user.room).emit('users', getUsers(user.room))
//     }
// })
  socket.on("disconnect", () => {
     removeUser(userName,indexDelete,userID)
     socket.broadcast.emit('user-list', [...userList])
     socket.emit('user-list', [...userList])
  });
});

// function addUser(userName, id){
//   userList.push({
//     name: userName,
//     id: id
//   })
//   index++
// }

function removeUser(userName, indexDelete, userID){ // elimina o user quando faz logout ou faz reload
  console.log("user disconnect " + userName)
  userList.splice(indexDelete,1)
  socketID.splice(userID,1)
  index--
}


// const getUsers = (rooms) => room.filter(user => user.roomID === rooms)