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
			So(r.Description, ShouldEqual, "The collection of Users")

			So(r.Get, ShouldNotBeNil)
			So(r.Get.Description, ShouldEqual, "Get all Users, optionally filtered")
			So(r.Get.Responses[200].Bodies.Type, ShouldEqual, "Users")

			So(r.Post, ShouldNotBeNil)
			So(r.Post.Description, ShouldEqual, "Create a new User")
			So(r.Post.Responses[200].Bodies.Type, ShouldEqual, "User")
		})

		Convey("checking queues - optional method", func() {
			r := apiDef.Resources["/queues"]
			So(r, ShouldNotBeNil)

			So(r.Get, ShouldNotBeNil)
			So(r.Get.Description, ShouldEqual, "Get all queues")

			So(r.Post, ShouldBeNil)
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})

	})
}
