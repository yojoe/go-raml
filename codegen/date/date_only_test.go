package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateOnly(t *testing.T) {
	Convey("date-only", t, func() {
		Convey("not in struct", func() {
			dateStr := "2016-05-04"

			// create time
			tim, err := time.Parse("2006-01-02", dateStr)
			So(err, ShouldBeNil)

			do := DateOnly(tim)

			// marshal
			b, err := json.Marshal(&do)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, `"`+dateStr+`"`)

			// unmarshal
			err = json.Unmarshal([]byte(`"`+dateStr+`"`), &do)
			So(err, ShouldBeNil)
			So(do.String(), ShouldEqual, dateStr)
		})

		Convey("in struct", func() {
			jsonBytes := []byte(`{"name":"google","born":"2016-05-04"}`)
			var data = struct {
				Name string   `json:"name"`
				Born DateOnly `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			So(err, ShouldBeNil)

			// marshal again
			b, err := json.Marshal(&data)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
