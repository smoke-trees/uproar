const express = require('express');
const router = require('./routes/router')
const path = require('path');
const jwt = require('jsonwebtoken')
const cookieParser = require('cookie-parser')

const app = express();

function authmiddleware(req, res, next) {
    token = req.cookies.jwt
    try {
        claims = jwt.verify(token, "smoketrees", {algorithm: "HS256"})

    } catch (err) {
        res.redirect("/login")
    }

    next()
}

app.set("view engine", "ejs");
app.set("views", path.join(__dirname, "views"));
app.use(express.urlencoded({extended: false}));
app.use(cookieParser());
app.use("/public", express.static(__dirname + "/public"));

//Global vars

app.use('/', router);
const PORT = process.env.PORT || 5000;

app.listen(PORT, console.log('Server started'));