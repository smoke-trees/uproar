const express=require('express');
const router=express.Router();
const bcrypt=require('bcryptjs');
//User Model

router.get("/login",(req,res)=>res.render("login"));
router.get("/register",(req,res)=>res.render('register'));
//Register Handle
router.post("/register",(req,res)=>{
    console.log(req.body)
    
});
//Login Handle
router.post('/login',(req,res,next)=>{
    
});
//Logout handle
router.get("/logout",(req,res)=>{
    
})
module.exports=router;