# -*- coding: utf-8 -*-
"""
Created on Sun Mar  8 00:05:37 2020

@author: Tanmay Thakur
"""

import tensorflow as tf
from tensorflow.keras.models import load_model   


def init():
    model = load_model('toxic_model.h5')
    graph = tf.compat.v1.get_default_graph()
    
    return model,graph