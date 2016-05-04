package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatetimeOnly(t *testing.T) {
	Convey("datetime-only", t, func() {
		Convey("not in struct", func() {
			dateStr := "2015-07-04T21:00:00"

			// create time
			tim, err := time.Parse("2006-01-02T15:04:05.99", dateStr)
			So(err, ShouldBeNil)

			to := DatetimeOnly(tim)

			// marshal
			b, err := json.Marshal(&to)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, `"`+dateStr+`"`)

			// unmarshal
			err = json.Unmarshal([]byte(`"`+dateStr+`"`), &to)
			So(err, ShouldBeNil)
			So(to.String(), ShouldEqual, dateStr)
		})

		Convey("in struct", func() {
			jsonBytes := []byte(`{"name":"google","born":"2015-07-04T21:00:00"}`)
			var data = struct {
				Name string       `json:"name"`
				Born DatetimeOnly `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			So(err, ShouldBeNil)
			So(data.Born.String(), ShouldEqual, "2015-07-04T21:00:00")

			// marshal again
			b, err := json.Marshal(&data)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
