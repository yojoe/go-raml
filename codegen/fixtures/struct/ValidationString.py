
from wtforms import TextField, FormField, IntegerField

class ValidationString():
    
    name = TextField(validators=[Length(min=8, max=40)])
    
