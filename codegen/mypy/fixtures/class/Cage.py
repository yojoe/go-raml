"""
Auto-generated class for Cage
"""
import capnp
import os
from .Animal import Animal

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class Cage:
    """
    auto-generated. don't touch.
    """

    def __init__(self, colours: str, owner: Animal) -> None:
        """
        :type colours: str
        :type owner: Animal
        :rtype: Cage
        """
        self.colours = colours  # type: str
        self.owner = owner  # type: Animal

    def to_capnp(self):
        template = capnp.load('%s/Cage.capnp' % dir)
        return template.Cage.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class CageCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> Cage:
        template = capnp.load('%s/Cage.capnp' % dir)
        struct = template.Cage.from_bytes(bin) if bin else template.Cage.new_message()
        return Cage(**struct.to_dict(verbose=True))
