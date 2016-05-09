package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeOnly(t *testing.T) {
	Convey("time-only", t, func() {
		Convey("not in struct", func() {
			dateStr := "10:09:08"

			// create time
			tim, err := time.Parse("15:04:05", dateStr)
			So(err, ShouldBeNil)

			to := TimeOnly(tim)

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
			jsonBytes := []byte(`{"name":"google","return":"10:09:08"}`)
			var data = struct {
				Name   string   `json:"name"`
				Return TimeOnly `json:"return"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			So(err, ShouldBeNil)
			So(data.Return.String(), ShouldEqual, "10:09:08")

			// marshal again
			b, err := json.Marshal(&data)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
