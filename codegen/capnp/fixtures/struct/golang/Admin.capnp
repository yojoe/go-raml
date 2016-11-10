
using Go = import "/go.capnp";
using import "EnumAdminClearanceLevel.capnp".EnumAdminClearanceLevel;
@0xd055d40a513ed2f1;

$Go.package("main");
$Go.import("main");
struct Admin {
  clearanceLevel @0 :EnumAdminClearanceLevel;
}
