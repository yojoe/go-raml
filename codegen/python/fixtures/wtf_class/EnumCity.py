
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange, required
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField, FieldList
from input_validators import multiple_of


class EnumCity(Form):
    
    enum_homeNum = IntegerField(validators=[DataRequired(message="")])
    enum_parks = TextField(validators=[DataRequired(message="")])
    name = TextField(validators=[DataRequired(message="")])
