
using import "Admin.capnp".Admin;
using import "Animal.capnp".Animal;
@0xce31880b1fdb3001;

struct Cage {
  animal @0 :Animal;
  admin @1 :Admin;
}
