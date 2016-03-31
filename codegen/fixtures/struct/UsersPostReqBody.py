
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange, required
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField, FieldList
from input_validators import multiple_of

class UsersPostReqBody(Form):
    
    ID = TextField(validators=[DataRequired(message=""), Length(min=4, max=8)])
    
    
    age = IntegerField(validators=[DataRequired(message=""), NumberRange(min=16, max=100), multiple_of(mult=4)])
    
    
    
    grades = FieldList(IntegerField('grades', [required()]), min_entries=2,max_entries=5)
    
    item = TextField(validators=[DataRequired(message=""), Length(min=2), Regexp(regex="^[a-zA-Z]+$")])
    
    
