"""
Auto-generated class for EnumCity
"""
import capnp
import os
from .EnumEnumCityEnumParks import EnumEnumCityEnumParks

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class EnumCity:
    """
    auto-generated. don't touch.
    """

    def __init__(self, enumParks: EnumEnumCityEnumParks, name: str) -> None:
        """
        :type enumParks: EnumEnumCityEnumParks
        :type name: str
        :rtype: EnumCity
        """
        self.enumParks = enumParks  # type: EnumEnumCityEnumParks
        self.name = name  # type: str

    def to_capnp(self):
        """
        Load the class in capnp schema EnumCity.capnp
        :rtype bytes
        """
        template = capnp.load('%s/EnumCity.capnp' % dir)
        return template.EnumCity.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class EnumCityCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> EnumCity:
        """
        Load the binary of EnumCity.capnp into class EnumCity
        :type bin: bytes. If none creates an empty capnp object.
        rtype: EnumCity
        """
        template = capnp.load('%s/EnumCity.capnp' % dir)
        struct = template.EnumCity.from_bytes(bin) if bin else template.EnumCity.new_message()
        return EnumCity(**struct.to_dict(verbose=True))
