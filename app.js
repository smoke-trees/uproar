const express=require('express');
const expressLayouts=require('express-ejs-layouts');
const mongoose=require('mongoose');
const flash=require('connect-flash');
const session=require('express-session');
const passport=require('passport');
const app=express();
//passport Config
require('./config/passport')(passport);
//DB Config
const db=require('./config/keys').MongoURI;
app.use(expressLayouts);
mongoose.connect(db,{useNewUrlParser: true,useUnifiedTopology: true})
    .then(()=>console.log("MongoDB connceted"))
    .catch(err=>console.log(err));
app.set("view engine","ejs");
app.use(express.urlencoded({extended: false}));
//Express Session
app.use(session({
    secret: 'secret',
    resave: false,
    saveUninitialized: true
}));
app.use(passport.initialize());
app.use(passport.session());
//Flash
app.use(flash());
//Global vars
app.use((req,res,next)=>{
    res.locals.success_msg=req.flash('success_msg');
    res.locals.error_msg=req.flash('error_msg');
    next();
})
app.use("/",require("./routes/index"));
app.use('/users',require("./routes/users"));
const PORT=process.env.PORT || 5000;

app.listen(PORT,console.log('Server started'));