"""
Auto-generated class for NumberFormat
"""

from . import client_support


class NumberFormat(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(d, f, i, i16, i32, i64, i8, l, num):
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

        return NumberFormat(
            d=d,
            f=f,
            i=i,
            i16=i16,
            i32=i32,
            i64=i64,
            i8=i8,
            l=l,
            num=num,
        )

    def __init__(self, json=None, **kwargs):
        if not json and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'NumberFormat'
        create_error = '{cls}: unable to create {prop} from value: {val}: {err}'
        required_error = '{cls}: missing required property {prop}'

        data = json or kwargs

        property_name = 'd'
        val = data.get(property_name)
        if val is not None:
            datatypes = [float]
            try:
                self.d = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'f'
        val = data.get(property_name)
        if val is not None:
            datatypes = [float]
            try:
                self.f = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'i'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.i = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'i16'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.i16 = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'i32'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.i32 = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'i64'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.i64 = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'i8'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.i8 = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'l'
        val = data.get(property_name)
        if val is not None:
            datatypes = [int]
            try:
                self.l = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

        property_name = 'num'
        val = data.get(property_name)
        if val is not None:
            datatypes = [float]
            try:
                self.num = client_support.val_factory(val, datatypes)
            except ValueError as err:
                raise ValueError(create_error.format(cls=class_name, prop=property_name, val=val, err=err))
        else:
            raise ValueError(required_error.format(cls=class_name, prop=property_name))

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)
