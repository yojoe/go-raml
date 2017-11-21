"""
Auto-generated class for WithDateTime
"""
import capnp
import os
from datetime import datetime

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class WithDateTime:
    """
    auto-generated. don't touch.
    """

    def __init__(self, birth: datetime, name: str) -> None:
        """
        :type birth: datetime
        :type name: str
        :rtype: WithDateTime
        """
        self.birth = birth  # type: datetime
        self.name = name  # type: str

    def to_capnp(self):
        template = capnp.load('%s/WithDateTime.capnp' % dir)
        return template.WithDateTime.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class WithDateTimeCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> WithDateTime:
        template = capnp.load('%s/WithDateTime.capnp' % dir)
        struct = template.WithDateTime.from_bytes(bin) if bin else template.WithDateTime.new_message()
        return WithDateTime(**struct.to_dict(verbose=True))
