"""
Auto-generated class for SingleInheritance
"""
from .EnumCity import EnumCity

from . import client_support


class SingleInheritance(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(cities, colours, name, single):
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :type single: bool
        :rtype: SingleInheritance
        """

        return SingleInheritance(
            cities=cities,
            colours=colours,
            name=name,
            single=single,
        )

    def __init__(self, json=None, **kwargs):
        if not json and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'SingleInheritance'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'cities'
        val = data.get(property_name)
        if val is not None:
            datatypes = [EnumCity]
            try:
                self.cities = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'colours'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.colours = client_support.list_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'name'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                self.name = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'single'
        val = data.get(property_name)
        if val is not None:
            datatypes = [bool]
            try:
                self.single = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
