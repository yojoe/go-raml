
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField

class UsersPostReqBody(Form):
    
    age = IntegerField(validators=[])
    
    ID = TextField(validators=[Length(min=4, max=8)])
    
    item = TextField(validators=[Length(min=2)])
    
