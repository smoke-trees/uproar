# -*- coding: utf-8 -*-
"""
Created on Sun Mar  8 00:03:34 2020

@author: Tanmay Thakur
"""

import pickle
import load as l

from tensorflow.keras.preprocessing.sequence import pad_sequences


tokenizer = pickle.load(open("tokenizer_instance.pickle","rb"))

model, _ = l.init()

def make_sentence(sentence):
    return [sentence]


def preprocessing(sentence, max_features = 20000, maxlen = 50, tokenizer = tokenizer):
    list_tokenized_train = tokenizer.texts_to_sequences(sentence)
    
    X_t = pad_sequences(list_tokenized_train, maxlen = maxlen)
    return X_t

def prediction(x):
    x = preprocessing(make_sentence(x))
    list_classes = ["Toxic", "Severely Toxic", "Obscene", "Threat", "Insult", "Identity Hate"]     
    x = str(dict(zip(list_classes, 100*model.predict([x,]).flatten())))
    return x