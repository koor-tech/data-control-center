const express = require('express')
const dotenv = require('dotenv').config()
const history = require('connect-history-api-fallback')

const app = express()
const port = process.env.PORT || 3000

app.use(express.static('dist/spa'))
app.use(history())

app.get('/ping', (req, res) => {
  res.send('pong')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})
