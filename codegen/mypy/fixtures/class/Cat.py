"""
Auto-generated class for Cat
"""
import capnp
import os

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class Cat(object):
    """
    auto-generated. don't touch.
    """

    def __init__(self, kind: str) -> None:
        """
        :type kind: str
        :rtype: Cat
        """
        self.kind = kind  # type: str

    def to_capnp(self):
        template = capnp.load('%s/Cat.capnp' % dir)
        return template.Cat.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class CatCollection(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> Cat:
        template = capnp.load('%s/Cat.capnp' % dir)
        struct = template.Cat.from_bytes(bin) if bin else template.Cat.new_message()
        return Cat(**struct.to_dict(verbose=True))
