# Limitations

## Unsupported RAML 1.0 features

* Types
    * [Map](http://docs.raml.org/specs/1.0/#raml-10-spec-map-types) no support for `patternProperties`
    * No support for the `[]` way of defining maps, you need to use `"[]"` since `[]` is not accepted by the yaml parser.
      This will change in the final raml spec so no plans to fix this: https://github.com/raml-org/raml-spec/issues/286


* Modularization
    * [References to inner elements of external files](http://docs.raml.org/specs/1.0/#references-to-inner-elements-of-external-files)
    * [Libraries](http://docs.raml.org/specs/1.0/#libraries)
    * [Overlays and extensions](http://docs.raml.org/specs/1.0/#overlays-and-extensions)
