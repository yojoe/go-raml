"""
Auto-generated class for SingleInheritance
"""
from .EnumCity import EnumCity
from six import string_types

from . import client_support


class SingleInheritance(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type cities: list[EnumCity]
        :type colours: list[string_types]
        :type name: string_types
        :type single: bool
        :rtype: SingleInheritance
        """

        return SingleInheritance(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'SingleInheritance'
        data = json or kwargs

        # set attributes
        data_types = [EnumCity]
        self.cities = client_support.set_property('cities', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.colours = client_support.set_property('colours', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.name = client_support.set_property('name', data, data_types, False, [], False, True, class_name)
        data_types = [bool]
        self.single = client_support.set_property('single', data, data_types, False, [], False, True, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
