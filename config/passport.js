const LocalStrategy=require('passport-local').Strategy;
const mongoose=require('mongoose');
const bcrypt=require('bcryptjs');
const User=require('../models/User');
module.exports=function(passport){
    passport.use(
        new LocalStrategy({usernameField: 'email'},(email,password,done)=>{
            //Match user
            User.findOne({email: email})
            .then(user=>{
                if(!user){
                    console.log("That email isn't registered");
                    return done(null,false,{message: 'That email is not registered'});
                }
                //Match Password
                bcrypt.compare(password,user.password,(err,isMatch)=>{
                    if(err) throw err;
                    if(isMatch){
                        return(done(null,user));
                    }
                    else{
                        return done(null,{message:"Password is in-correct"});
                    }
                });
            })
            .catch(err=>console.log(err));
        })
    );
    passport.serializeUser(function(user,done){
        done(null,user.id);
    });
    passport.deserializeUser((id,done)=>{
        User.findById(id,function(err,user){
            done(err,user);
        });
    });
}