
using import "Admin.capnp".Admin;
using import "Animal.capnp".Animal;
@0xce31880b1fdb3001;

struct Cage {
  admin @0 :Admin;
  animal @1 :Animal;
}
