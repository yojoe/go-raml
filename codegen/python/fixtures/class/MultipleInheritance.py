"""
Auto-generated class for MultipleInheritance
"""
from .EnumCity import EnumCity

from . import client_support


class MultipleInheritance(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type cities: list[EnumCity]
        :type color: str
        :type colours: list[str]
        :type kind: str
        :type name: str
        :rtype: MultipleInheritance
        """

        return MultipleInheritance(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'MultipleInheritance'
        data = json or kwargs

        # set attributes
        data_types = [EnumCity]
        self.cities = client_support.set_property('cities', data, data_types, False, [], True, True, class_name)
        data_types = [str]
        self.color = client_support.set_property('color', data, data_types, False, [], False, True, class_name)
        data_types = [str]
        self.colours = client_support.set_property('colours', data, data_types, False, [], True, True, class_name)
        data_types = [str]
        self.kind = client_support.set_property('kind', data, data_types, False, [], False, True, class_name)
        data_types = [str]
        self.name = client_support.set_property('name', data, data_types, False, [], False, False, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
