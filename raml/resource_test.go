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

		Convey("books - query parameters", func() {
			r := apiDef.Resources["/books"]
			So(r, ShouldNotBeNil)

			qps := r.Get.QueryParameters
			So(qps["title"].Description, ShouldEqual, "Return books that have their title matching the given value")
			So(qps["digest_all_fields"].Description, ShouldEqual,
				"If no values match the value given for title, use digest_all_fields instead")
		})

		Reset(func() {
			os.RemoveAll(targetDir)
		})

	})
}
