from app import app
from app.forms import TextForm
from flask import render_template,redirect,request
import spacy
import os, sys





@app.route('/')
def lol():
	return render_template('index.html')
@app.route('/form',methods=["GET","POST"])
def index():
	return render_template('ner_inp.html')
@app.route("/handledata",methods=["GET","POST"])
def calc():
	inp = request.form["inp"]
	nlp = spacy.load('en_core_web_sm')
		
	out = nlp(inp)
	ls = {}
	for ents in out.ents:
		ls[ents.text] = ents.label_

	return render_template('ner_inp.html',data=ls)



