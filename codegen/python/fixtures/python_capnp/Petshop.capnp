
using import "Cat.capnp".Cat;
@0xdfb77c5498cdf75b;

struct Petshop {
  cats @0 :List(Cat);
  name @1 :Text;
}
