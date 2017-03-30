package nim

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypeConversion(t *testing.T) {
	Convey("Test Type Conversion", t, func() {
		Convey("Type conversion", func() {
			So(toNimType("string"), ShouldEqual, "string")
			So(toNimType("number"), ShouldEqual, "float")
			So(toNimType("integer"), ShouldEqual, "int")
			So(toNimType("boolean"), ShouldEqual, "bool")
			So(toNimType("file"), ShouldEqual, "string")
			So(toNimType("date-only"), ShouldEqual, "Time")
			So(toNimType("time-only"), ShouldEqual, "Time")
			So(toNimType("datetime"), ShouldEqual, "Time")
			So(toNimType("Object"), ShouldEqual, "Object")
			So(toNimType("string[]"), ShouldEqual, "seq[string]")
			So(toNimType("string[][]"), ShouldEqual, "seq[seq[string]]")
			//So(toNimType("string | Person"), ShouldEqual, "interface{}")
			//So(toNimType("(string | Person)[]"), ShouldEqual, "[]interface{}")
		})
	})
}
