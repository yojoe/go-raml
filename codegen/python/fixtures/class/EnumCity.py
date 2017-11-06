"""
Auto-generated class for EnumCity
"""
from .EnumEnumCityEnum_homeNum import EnumEnumCityEnum_homeNum
from .EnumEnumCityEnum_parks import EnumEnumCityEnum_parks

from . import client_support


class EnumCity(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type enum_homeNum: EnumEnumCityEnum_homeNum
        :type enum_parks: EnumEnumCityEnum_parks
        :type name: str
        :rtype: EnumCity
        """

        return EnumCity(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'EnumCity'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'enum_homeNum'
        val = data.get(property_name)
        if val is not None:
            datatypes = [EnumEnumCityEnum_homeNum]
            try:
                setattr(self, 'enum_homeNum', client_support.val_factory(val, datatypes))
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'enum_parks'
        val = data.get(property_name)
        if val is not None:
            datatypes = [EnumEnumCityEnum_parks]
            try:
                setattr(self, 'enum_parks', client_support.val_factory(val, datatypes))
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'name'
        val = data.get(property_name)
        if val is not None:
            datatypes = [str]
            try:
                setattr(self, 'name', client_support.val_factory(val, datatypes))
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
