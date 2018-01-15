#python-capnp

Table of Contents
=================

* [Install pycapnp](#install-pycapnp)
* [Generation](#generation)


Other than source code, go-raml is able to produce python 3 classes that can load data from/to
capnp [plain schemas](./capnp.md#plain-schema).


## install pycapnp

**Install python library**

The generated classes use pycapnp to load data from/to capnp.

Installation procedures can be found at http://jparyani.github.io/pycapnp/install.html.


## generation
To generate the classes run the following command:
```
go-raml python-capnp --dir /result_dir --ramlfile api_file.raml
```

This will generate python classes and capnp schemas that will be used by the python classes.

For example if api_file.raml includes the following:

```raml
#%RAML 1.0
title: Struct API Test
mediaType: application/json
types:
  EnumCity:
    description: |
      first line
      second line
      third line
    properties:
      name: string
      enumParks:
        type: string
        enum: [ parkA, parkB ]
  Animal:
    description: |
      Animal represent animal object.
      It contains field that construct animal
      such as : name, colours, and cities.
    type: object
    properties:
      name?: string
      colours:
        type: array
        items: string
      cities:
        type: EnumCity[]
        minItems: 1
        maxItems: 10
```
This will be the generated python class for Animal:

```python
"""
Auto-generated class for Animal
"""
import capnp
import os
from .EnumCity import EnumCity
from six import string_types

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class Animal(object):
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def create(**kwargs):
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :rtype: Animal
        """

        return Animal(**kwargs)

    def __init__(self, json=None, **kwargs):
        if json is None and not kwargs:
            raise ValueError('No data or kwargs present')

        class_name = 'Animal'
        data = json or kwargs

        # set attributes
        data_types = [EnumCity]
        self.cities = client_support.set_property('cities', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.colours = client_support.set_property('colours', data, data_types, False, [], True, True, class_name)
        data_types = [string_types]
        self.name = client_support.set_property('name', data, data_types, False, [], False, False, class_name)

    def __str__(self):
        return self.as_json(indent=4)

    def as_json(self, indent=0):
        return client_support.to_json(self, indent=indent)

    def as_dict(self):
        return client_support.to_dict(self)

    def to_capnp(self):
        """
        Load the class in capnp schema Animal.capnp
        :rtype bytes
        """
        template = capnp.load('%s/Animal.capnp' % dir)
        return template.Animal.new_message(**self.as_dict()).to_bytes()


class AnimalCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(binary=None):
        """
        Load the binary of Animal.capnp into class Animal
        :type binary: bytes. If none creates an empty capnp object.
        rtype: Animal
        """
        template = capnp.load('%s/Animal.capnp' % dir)
        struct = template.Animal.from_bytes(binary) if bin else template.Animal.new_message()
        return Animal(**struct.to_dict(verbose=True))

```

where `Animal.to_capnp` loads the class into the capnp schema and returns the binary values and `AnimalCollection.new` loads a capnp binary into the python class.

This will be the genrated capnp schema:
```capnp

using import "EnumCity.capnp".EnumCity;
@0xae962dc0f8d7e42d;

struct Animal {
  cities @0 :List(EnumCity);
  colours @1 :List(Text);
  name @2 :Text;
}

```

## Notes on type alias

No capnp will be generated for type alias:
- alias of builtin type: capnp doesn't support it
- alias of non builtin type: use capnp of the aliased type.

In the python class of builtin type alias,  there will be no functions to load from/to python/capnp like we have when we generate a gevent client/server.