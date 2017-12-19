
using import "Animal.capnp".Animal;
@0xdc061a9cc3b116dd;

struct Cage {
  colours @0 :Text;
  owner @1 :Animal;
}
