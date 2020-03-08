# -*- coding: utf-8 -*-
"""
Created on Sun Mar  8 00:49:35 2020

@author: Tanmay Thakur
"""

from flask import Flask, request
from twilio.twiml.messaging_response import MessagingResponse


app = Flask(__name__)


@app.route('/sms', methods=['POST'])
def sms():
    number = request.form['From']
    message_body = request.form['Body']

    resp = MessagingResponse()
    resp.message('Hello {}, you said: {}'.format(number, message_body))
    return str(resp)

if __name__ == '__main__':
    app.run()