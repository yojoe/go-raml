"""
Auto-generated class for NumberFormat
"""
import capnp
import os

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class NumberFormat(object):
    """
    auto-generated. don't touch.
    """

    def __init__(self, d: float, f: float, i16: int, i32: int, i64: int, i8: int, i: int, l: int, num: float) -> None:
        """
        :type d: float
        :type f: float
        :type i: int
        :type i16: int
        :type i32: int
        :type i64: int
        :type i8: int
        :type l: int
        :type num: float
        :rtype: NumberFormat
        """
        self.d = d  # type: float
        self.f = f  # type: float
        self.i = i  # type: int
        self.i16 = i16  # type: int
        self.i32 = i32  # type: int
        self.i64 = i64  # type: int
        self.i8 = i8  # type: int
        self.l = l  # type: int
        self.num = num  # type: float

    def to_capnp(self):
        template = capnp.load('%s/NumberFormat.capnp' % dir)
        return template.NumberFormat.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class NumberFormatCollection(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> NumberFormat:
        template = capnp.load('%s/NumberFormat.capnp' % dir)
        struct = template.NumberFormat.from_bytes(bin) if bin else template.NumberFormat.new_message()
        return NumberFormat(**struct.to_dict(verbose=True))
