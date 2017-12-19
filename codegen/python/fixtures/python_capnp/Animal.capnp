
using import "EnumCity.capnp".EnumCity;
@0xae962dc0f8d7e42d;

struct Animal {
  cities @0 :List(EnumCity);
  colours @1 :List(Text);
  name @2 :Text;
}
