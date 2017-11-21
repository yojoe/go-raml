"""
Auto-generated class for PlainObject
"""
import capnp
import os

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class PlainObject(object):
    """
    auto-generated. don't touch.
    """

    def __init__(self, obj: dict) -> None:
        """
        :type obj: dict
        :rtype: PlainObject
        """
        self.obj = obj  # type: dict

    def to_capnp(self):
        template = capnp.load('%s/PlainObject.capnp' % dir)
        return template.PlainObject.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class PlainObjectCollection(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> PlainObject:
        template = capnp.load('%s/PlainObject.capnp' % dir)
        struct = template.PlainObject.from_bytes(bin) if bin else template.PlainObject.new_message()
        return PlainObject(**struct.to_dict(verbose=True))
