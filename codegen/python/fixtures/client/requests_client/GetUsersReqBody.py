"""
Auto-generated class for GetUsersReqBody
"""

from . import client_support


class GetUsersReqBody(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create():
        """
        :rtype: GetUsersReqBody
        """

        return GetUsersReqBody(
        )

    def __init__(self, json=None, **kwargs):
        if not json and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'GetUsersReqBody'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
