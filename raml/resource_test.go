package raml

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResourceTypeInheritance(t *testing.T) {
	Convey("resource type inheritance", t, func() {
		apiDef, err := ParseFile("./samples/resource_types.raml")
		So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		So(err, ShouldBeNil)

		Convey("checking users", func() {
			r := apiDef.Resources["/Users"]
			So(r.URI, ShouldEqual, "/Users")

			So(r.Get, ShouldNotBeNil)
			So(r.Get.Responses[200].Bodies.Type, ShouldEqual, "Users")
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})

	})
}
