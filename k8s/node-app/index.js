'use strict'

const express = require('express');
const app = express();

const port = process.env.PORT || 8080;
const param1 = process.env.PARAM1 || 'default-param';

app.get('/', (req, res) => res.send(`Hello from xendit! param1: ${param1}`));
app.listen(port, () => console.log(`Example app listening on port ${port}!`));
