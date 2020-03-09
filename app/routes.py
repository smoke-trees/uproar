from app import app
from flask import render_template,redirect,request
import model.load_toxic as l
from model.predict_toxic import make_sentence,preprocessing,prediction
import pickle
from tensorflow.keras.preprocessing.sequence import pad_sequences
from tensorflow.keras.models import load_model   
import tensorflow as tf
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
	place_model = spacy.load("/home/rosguy/uproar/model/model")
	out_1 = nlp(inp)
	out_2 = place_model(inp)
	ls = {}
	#for ents in out_1.ents:
	#	ls[ents.text] = ents.label_
	for ents in out_2.ents:
		ls[ents.text] = ents.label_
	return render_template('ner_inp.html',data=ls)


@app.route("/toxicform",methods=["GET","POST"])
def display():
	return render_template('toxicity.html')

@app.route("/toxic",methods=["GET","POST"])
def toxic():
	model = load_model('/home/rosguy/uproar/model/toxic_model.h5')
	inp = request.form["sentence"]
	
	sentence = preprocessing(make_sentence(inp))
	list_classes = ["Toxic", "Severely Toxic", "Obscene", "Threat", "Insult", "Identity Hate"]     
	pred = str(dict(zip(list_classes, 100*model.predict(sentence).flatten())))
	
	#with graph.as_default():
	#	with tf.Session() as sess:
	#		sess.run(tf.global_variables_initializer())
			
	#		pred = prediction(inp)
		
	return render_template('toxicity.html',data=pred)

