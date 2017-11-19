"""
Auto-generated class for Cage
"""
from .animal import animal

from . import client_support


class Cage(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type colours: str
        :type owner: animal
        :rtype: Cage
        """

        return Cage(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'Cage'
        data = json or kwargs

        # set attributes
        data_types = [str]
        self.colours = client_support.set_property('colours', data, data_types, False, [], False, True, class_name)
        data_types = [animal]
        self.owner = client_support.set_property('owner', data, data_types, False, [], False, True, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
