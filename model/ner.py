# -*- coding: utf-8 -*-
"""
Created on Sun Mar  8 00:02:22 2020

@author: Tanmay Thakur
"""

import spacy
import en_core_web_sm


nlp = en_core_web_sm.load()

# Test 
doc = nlp('Amogh Lele integrates USA arsenal into Amazon web archives, Russia protests vehemently due to Kremlin concerns.')
print([(X.text, X.label_) for X in doc.ents])