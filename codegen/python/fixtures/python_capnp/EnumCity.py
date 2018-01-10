"""
Auto-generated class for EnumCity
"""
import capnp
import os
from .EnumEnumCityEnumParks import EnumEnumCityEnumParks
from six import string_types

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class EnumCity(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type enumParks: EnumEnumCityEnumParks
        :type name: string_types
        :rtype: EnumCity
        """

        return EnumCity(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'EnumCity'
        data = json or kwargs

        # set attributes
        data_types = [EnumEnumCityEnumParks]
        self.enumParks = client_support.set_property('enumParks', data, data_types, False, [], False, True, class_name)
        data_types = [string_types]
        self.name = client_support.set_property('name', data, data_types, False, [], False, True, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)

    def to_capnp(self):
        """
        Load the class in capnp schema EnumCity.capnp
        :rtype bytes
        """
        template = capnp.load('%s/EnumCity.capnp' % dir)
        return template.EnumCity.new_message(**self.as_dict()).to_bytes()


class EnumCityCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(binary=None):
        """
        Load the binary of EnumCity.capnp into class EnumCity
        :type binary: bytes. If none creates an empty capnp object.
        rtype: EnumCity
        """
        template = capnp.load('%s/EnumCity.capnp' % dir)
        struct = template.EnumCity.from_bytes(binary) if binary else template.EnumCity.new_message()
        return EnumCity(**struct.to_dict(verbose=True))
