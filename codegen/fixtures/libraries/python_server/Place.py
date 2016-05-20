
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange, required
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField, FieldList
from input_validators import multiple_of

from files.Directory import files.Directory

from datetime import datetime


class Place(Form):
    
    created = FormField(datetime)
    dir = FormField(files.Directory)
    name = TextField(validators=[DataRequired(message="")])
