const express = require("express");
const jwt = require("jsonwebtoken")
const router = express.Router();
const fetch = require("node-fetch")

function authmiddleware(req, res, next) {
    token = req.cookies.jwt
    try {
        claims = jwt.verify(token, "smoketrees", {algorithm: "HS256"})

    } catch (err) {
        res.redirect("/login")
    }

    next()
}

router.get("/", (req, res) => {
    res.redirect("/login")
});

router.get("/login", (req, res) => {
    res.render('login');
})

router.get("/feed", authmiddleware, (req, res) => {
    token = req.cookies.jwt
    try {
        claims = jwt.verify(token, "smoketrees", {algorithm: "HS256"})

    } catch (err) {
        res.redirect("/login")
    }
    fetch("http://localhost:3000/forum/feed?username=" + claims.username).then(res1 => res1.json()).then(json => {
        console.log(json)
        res.render("feed", {json: json, user: claims.username, jwt: token})
    });
})

router.get("/register", (req, res) => {
    res.render('register')
})

router.get("/post/new", authmiddleware, (req, res) => {
    res.render("new")
})

module.exports = router;

