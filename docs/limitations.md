# Limitations

## Unsupported RAML 1.0 features

* Modularization
    * [References to inner elements of external files](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md#references-to-inner-elements)
    * [Libraries](http://docs.raml.org/specs/1.0/#libraries) not supported by Nim generator
    * [Overlays and extensions](http://docs.raml.org/specs/1.0/#overlays-and-extensions)
    * [type declaration inside the `type` facet](https://github.com/raml-org/raml-spec/wiki/RAML-1.0-RC1-vs-RC2) in `Value of the `type` facet.` section
