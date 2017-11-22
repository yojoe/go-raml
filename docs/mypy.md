#mypy

Table of Contents
=================

* [Install MyPy](#install-mypy)
* [Generation](#generation)


Other than source code, go-raml is able to produce python 3 classes that are typed using mypy and can load data from/to
capnp [plain schemas](./capnp.md#plain-schema).


## install mypy

```bash
 python3 -m pip install mypy
```

Other ways of installations can be found at http://mypy.readthedocs.io/en/latest/getting_started.html


**Install python library**

The generated classes use pycapnp to load data from/to capnp.

Installation procedures can be found at http://jparyani.github.io/pycapnp/install.html.





## generation
To generate the classes run the following command:
```
go-raml mypy --dir /result_dir --ramlfile api_file.raml
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
from typing import List

from . import client_support

dir = os.path.dirname(os.path.realpath(__file__))


class Animal:
    """
    auto-generated. don't touch.
    """

    def __init__(self, cities: List[EnumCity], colours: List[str], name: str=None) -> None:
        """
        :type cities: list[EnumCity]
        :type colours: list[str]
        :type name: str
        :rtype: Animal
        """
        self.cities = cities  # type: List[EnumCity]
        self.colours = colours  # type: List[str]
        self.name = name  # type: str

    def to_capnp(self):
        template = capnp.load('%s/Animal.capnp' % dir)
        return template.Animal.new_message(**self.as_dict()).to_bytes()

    def as_dict(self):
        return client_support.to_dict(self)


class AnimalCollection:
    """
    auto-generated. don't touch.
    """

    @staticmethod
    def new(bin=None) -> Animal:
        template = capnp.load('%s/Animal.capnp' % dir)
        struct = template.Animal.from_bytes(bin) if bin else template.Animal.new_message()
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