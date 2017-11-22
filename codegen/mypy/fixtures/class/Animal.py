"""
Auto-generated class for Animal
"""
import capnp
import os
from .EnumCity import EnumCity
from typing import List

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class Animal:
    """
    auto-generated. don't touch.
    """

    def __init__(self, cities: List[EnumCity], colours: List[str], name: str=None) -> None:
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :rtype: Animal
        """
        self.cities = cities  # type: List[EnumCity]
        self.colours = colours  # type: List[str]
        self.name = name  # type: str

    def to_capnp(self):
        """
        Load the class in capnp schema Animal.capnp
        :rtype bytes
        """
        template = capnp.load('%s/Animal.capnp' % dir)
        return template.Animal.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class AnimalCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> Animal:
        """
        Load the binary of Animal.capnp into class Animal
        :type bin: bytes. If none creates an empty capnp object.
        rtype: Animal
        """
        template = capnp.load('%s/Animal.capnp' % dir)
        struct = template.Animal.from_bytes(bin) if bin else template.Animal.new_message()
        return Animal(**struct.to_dict(verbose=True))
