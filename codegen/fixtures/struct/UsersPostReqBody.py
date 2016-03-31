
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField

class UsersPostReqBody(Form):
    
    ID = TextField(validators=[DataRequired(message=""), Length(min=4, max=8)])
    
    age = IntegerField(validators=[DataRequired(message=""), NumberRange(min=16, max=100)])
    
    item = TextField(validators=[DataRequired(message=""), Length(min=2), Regexp(regex="^[a-zA-Z]+$")])
    
