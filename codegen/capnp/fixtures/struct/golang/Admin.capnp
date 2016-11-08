
using Go = import "/go.capnp";
using import "EnumClearanceLevel.capnp".EnumClearanceLevel;
@0xdbcb1d47dbc9880d;

$Go.package("main");
$Go.import("main");
$Go.import("main");
struct Admin {
  clearanceLevel @0 :EnumClearanceLevel;
}
