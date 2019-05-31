const app = require('express')()
const server = require('http').Server(app)
const io = require('socket.io')(server)

server.listen(3000)

app.get('/', (req, res) => {
  res.sendFile(__dirname + '/index.html')
});

io.on('connection', socket => {
  socket.on('JOIN_ROOM', lessonNumber => {
    socket.join(lessonNumber)
  })

  socket.on('CREATE_POLL', ({ lesson_number, ...poll }) => {
    socket.broadcast.to(lesson_number).emit('POLL_STARTED', poll)
  })

  socket.on('POLL_ANSWERED', ({ lesson_number, ...answer }) => {
    socket.broadcast.to(lesson_number).emit('POLL_ANSWERED_NOTIFICATION', answer)
  })
})
