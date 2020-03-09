# -*- coding: utf-8 -*-
"""
Created on Mon Mar  9 07:13:00 2020

@author: Tanmay Thakur
"""

import spacy
import logging

from telegram.ext import Updater, MessageHandler, Filters


logging.basicConfig(format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
                     level=logging.INFO)


nlp = spacy.load('en_core_web_sm')

def some_func(bot, update):
    pass
    if not update.effective_message.text:
        update.effective_message.reply_text(text = "Cannot handle given format, getting aware now")
    else:
        msg = update.effective_message.text
        text = nlp(msg)
        text = str([(X.text, X.label_) for X in text.ents])
        update.effective_message.reply_text(text = text)
        
def main():
    updater = Updater('983042438:AAFoRWHUelWKHpjs3fYPuPM9p7r9SVY4MFM')
    dp = updater.dispatcher
    dp.add_handler(MessageHandler(Filters.all, some_func))
    updater.start_polling()
    updater.idle()
    
if __name__ == '__main__':
    main()