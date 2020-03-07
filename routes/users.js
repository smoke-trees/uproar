const express=require('express');
const router=express.Router();
const bcrypt=require('bcryptjs');
const passport=require('passport');
//User Model
const User=require('../models/User');
router.get("/login",(req,res)=>res.render("login"));
router.get("/register",(req,res)=>res.render('register'));
//Register Handle
router.post("/register",(req,res)=>{
    const{name,email,password,password2}=req.body;
    let errors=[];
    if(!name || !email || !password || !password2){
        errors.push({msg: "Please fill in all the fields"});
    }
    if(password!==password2)
    {
        errors.push({msg: "Passwords don't match"});
    }
    if(password.length<6){
        errors.push({msg:"password should be atleast of 6 characters"});
    } 
    if(errors.length>0)
    {
        res.render('register',{
            errors,
            name,
            email,
            password,
            password2
        });
    }
    else
    {
        errors.push({msg: "Email is already registered"});
        User.findOne({email: email})
        .then(user=>{
            if(user){
                res.render('register',{
                    errors,
                    name,
                    email,
                    password,
                    password2
                });
            }
            else{
                const newUser=new  User({
                    name,
                    email,
                    password
                });
                //Hash password
                bcrypt.genSalt(10,(err,salt)=>bcrypt.hash(newUser.password,salt,(err,hash)=>
                {
                    if(err) throw err;
                    newUser.password=hash;
                    newUser.save()
                    .then(user=>{
                        req.flash('success_msg','You are now registered user');
                        console.log("You can now Log In");
                        res.redirect('/users/login');
                    })
                    .catch(err => console.log(err));
                }
                ));
            }
        });
    }
});
//Login Handle
router.post('/login',(req,res,next)=>{
    passport.authenticate('local',{
        successRedirect:'/dashboard',
        failureRedirect: '/users/login',
        failureFlash: true
    })(req,res,next);
});
//Logout handle
router.get("/logout",(req,res)=>{
    req.logOut();
    req.flash('success_msg','You are successfully logged out');
    res.redirect('/users/login');
})
module.exports=router;