package raml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResourceTypeInheritance(t *testing.T) {
	Convey("resource type & traits inheritance", t, func() {
		apiDef, err := ParseFile("./samples/resource_types.raml")
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
			So(r.Post.Bodies.ApplicationJson.Type, ShouldEqual, "User")
			So(r.Post.Responses[200].Bodies.Type, ShouldEqual, "User")
		})

		Convey("checking queues - optional method", func() {
			r := apiDef.Resources["/queues"]
			So(r, ShouldNotBeNil)

			So(r.Get, ShouldNotBeNil)
			So(r.Get.Description, ShouldEqual, "Get all queues")

			So(r.Post, ShouldBeNil)
		})

		Convey("checking corps - header - resourcePath - request body", func() {
			r := apiDef.Resources["/corps"]
			So(r, ShouldNotBeNil)

			So(r.Post, ShouldNotBeNil)

			props := r.Post.Bodies.ApplicationJson.Properties
			So(ToProperty("name", props["name"]).Type, ShouldEqual, "string")
			So(ToProperty("age", props["age"]).Type, ShouldEqual, "int")
			So(r.Post.Headers["X-Chargeback"].Required, ShouldBeTrue)

			mem := r.Nested["/{id}"]
			So(mem, ShouldNotBeNil)
			So(mem.Get.Description, ShouldEqual, "get /corps/{id}")
		})

		Convey("books - query parameters", func() {
			r := apiDef.Resources["/books"]
			So(r, ShouldNotBeNil)

			qps := r.Get.QueryParameters
			So(qps["title"].Description, ShouldEqual, "Return books that have their title matching the given value")
			So(qps["digest_all_fields"].Description, ShouldEqual,
				"If no values match the value given for title, use digest_all_fields instead")
		})

		Convey("traits test", func() {
			r := apiDef.Resources["/books"]
			So(r, ShouldNotBeNil)

			qps := r.Get.QueryParameters
			So(qps["numPages"].Description, ShouldEqual, "The number of pages to return, not to exceed 10")
		})
	})
}
