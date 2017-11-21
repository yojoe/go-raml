"""
Auto-generated class for SingleInheritance
"""
import capnp
import os
from .EnumCity import EnumCity
from typing import List

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class SingleInheritance(object):
    """
    auto-generated. don't touch.
    """

    def __init__(self, cities: List[EnumCity], colours: List[str], name: str, single: bool) -> None:
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :type single: bool
        :rtype: SingleInheritance
        """
        self.cities = cities  # type: List[EnumCity]
        self.colours = colours  # type: List[str]
        self.name = name  # type: str
        self.single = single  # type: bool

    def to_capnp(self):
        template = capnp.load('%s/SingleInheritance.capnp' % dir)
        return template.SingleInheritance.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class SingleInheritanceCollection(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> SingleInheritance:
        template = capnp.load('%s/SingleInheritance.capnp' % dir)
        struct = template.SingleInheritance.from_bytes(bin) if bin else template.SingleInheritance.new_message()
        return SingleInheritance(**struct.to_dict(verbose=True))
