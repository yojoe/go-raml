
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField

class ValidationString(Form):
    
    name = TextField(validators=[DataRequired(message=""), Length(min=8, max=40)])
    
