
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange, required
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField, FieldList
from input_validators import multiple_of



class PersonInclude(Form):
    
    age = IntegerField(validators=[DataRequired(message=""), NumberRange(min=0)])
    firstName = TextField(validators=[DataRequired(message="")])
    lastName = TextField(validators=[DataRequired(message="")])
