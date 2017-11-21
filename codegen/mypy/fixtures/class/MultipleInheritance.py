"""
Auto-generated class for MultipleInheritance
"""
import capnp
import os
from .EnumCity import EnumCity
from typing import List

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class MultipleInheritance:
    """
    auto-generated. don't touch.
    """

    def __init__(self, cities: List[EnumCity], color: str, colours: List[str], kind: str, name: str=None) -> None:
        """
        :type cities: list[EnumCity]
        :type color: str
        :type colours: list[str]
        :type kind: str
        :type name: str
        :rtype: MultipleInheritance
        """
        self.cities = cities  # type: List[EnumCity]
        self.color = color  # type: str
        self.colours = colours  # type: List[str]
        self.kind = kind  # type: str
        self.name = name  # type: str

    def to_capnp(self):
        template = capnp.load('%s/MultipleInheritance.capnp' % dir)
        return template.MultipleInheritance.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class MultipleInheritanceCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> MultipleInheritance:
        template = capnp.load('%s/MultipleInheritance.capnp' % dir)
        struct = template.MultipleInheritance.from_bytes(bin) if bin else template.MultipleInheritance.new_message()
        return MultipleInheritance(**struct.to_dict(verbose=True))
