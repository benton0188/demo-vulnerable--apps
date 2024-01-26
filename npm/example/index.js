// A simple web server. Note this demo does not actually use any vulnerable dependencies.

const express = require('express')
const app = express()
const port = 8080

app.get('/', (req, res) => {
    res.send('Hello, World!')
})

app.listen(port, () => {
    console.log(`Server started on port ${port}`)
})
