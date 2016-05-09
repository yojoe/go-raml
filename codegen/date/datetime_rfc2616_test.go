package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateTimeRFC2616(t *testing.T) {

	Convey("datetime RF2616", t, func() {
		Convey("not in struct", func() {
			dateStr := "Sun, 28 Feb 2016 16:41:41 GMT"

			// create time
			tim, err := time.Parse(dateTimeRFC2616Fmt, dateStr)
			So(err, ShouldBeNil)

			to := DateTimeRFC2616(tim)

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
			jsonBytes := []byte(`{"name":"google","born":"Sun, 28 Feb 2016 16:41:41 GMT"}`)
			var data = struct {
				Name string          `json:"name"`
				Born DateTimeRFC2616 `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			So(err, ShouldBeNil)
			So(data.Born.String(), ShouldEqual, "Sun, 28 Feb 2016 16:41:41 GMT")

			// marshal again
			b, err := json.Marshal(&data)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, string(jsonBytes))
		})
	})

}
