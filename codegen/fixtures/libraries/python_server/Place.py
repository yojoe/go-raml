
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange, required
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField, FieldList
from input_validators import multiple_of

from libraries.files.Directory import Directory
from datetime import datetime


class Place(Form):
    
    created = FormField(datetime)
    dir = FormField(Directory)
    name = TextField(validators=[DataRequired(message="")])
