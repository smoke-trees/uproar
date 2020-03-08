from flask_wtf import FlaskForm
from wtforms import StringField,SubmitField

class TextForm(FlaskForm):
	text = StringField('inp')
	submit = SubmitField('submit')


