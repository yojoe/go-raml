
from flask_wtf import Form
from wtforms.validators import DataRequired, Length, Regexp, NumberRange, required
from wtforms import TextField, FormField, IntegerField, FloatField, FileField, BooleanField, DateField, FieldList
from input_validators import multiple_of

from EnumCity import EnumCity


class animal(Form):
    
    cities = FieldList(FormField(EnumCity))
    colours = FieldList(TextField('colours', [required()]), DataRequired(message=""))
    name = TextField(validators=[])
