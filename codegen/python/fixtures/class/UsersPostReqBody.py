"""
Auto-generated class for UsersPostReqBody
"""
from six import string_types

from . import client_support


class UsersPostReqBody(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type ID: string_types
        :type age: int
        :type grades: list[int]
        :type item: string_types
        :rtype: UsersPostReqBody
        """

        return UsersPostReqBody(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'UsersPostReqBody'
        data = json or kwargs

        # set attributes
        data_types = [string_types]
        self.ID = client_support.set_property('ID', data, data_types, False, [], False, True, class_name)
        data_types = [int]
        self.age = client_support.set_property('age', data, data_types, False, [], False, True, class_name)
        data_types = [int]
        self.grades = client_support.set_property('grades', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.item = client_support.set_property('item', data, data_types, False, [], False, True, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
