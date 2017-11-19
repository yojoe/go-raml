"""
Auto-generated class for animal
"""
from .EnumCity import EnumCity

from . import client_support


class animal(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :rtype: animal
        """

        return animal(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'animal'
        data = json or kwargs

        # set attributes
        data_types = [EnumCity]
        self.cities = client_support.set_property('cities', data, data_types, False, [], True, True, class_name)
        data_types = [str]
        self.colours = client_support.set_property('colours', data, data_types, False, [], True, True, class_name)
        data_types = [str]
        self.name = client_support.set_property('name', data, data_types, False, [], False, False, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
