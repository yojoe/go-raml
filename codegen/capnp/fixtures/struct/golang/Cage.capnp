
using Go = import "/go.capnp";
using import "Admin.capnp".Admin;
using import "Animal.capnp".Animal;
@0xb2907c63dfc23a7d;

$Go.package("main");
$Go.import("main");
struct Cage {
  animal @0 :Animal;
  admin @1 :Admin;
}
