const express=require('express');
const expressLayouts=require('express-ejs-layouts');
const mongoose=require('mongoose');
const flash=require('connect-flash');
const session=require('express-session');
const app=express();
app.use(expressLayouts);
app.set("view engine","ejs");
app.use(express.urlencoded({extended: false}));
//Express Session
app.use(session({
    secret: 'secret',
    resave: false,
    saveUninitialized: true
}));
//Flash
app.use(flash());
//Global vars
app.use((req,res,next)=>{
    res.locals.success_msg=req.flash('success_msg');
    res.locals.error_msg=req.flash('error_msg');
    next();
})

app.use('/users',require("./routes/users"));
const PORT=process.env.PORT || 5000;

app.listen(PORT,console.log('Server started'));