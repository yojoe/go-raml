package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateTime(t *testing.T) {
	Convey("datetime RFC3339", t, func() {
		Convey("not in struct", func() {
			dateStr := "2016-02-28T16:41:41.09Z"

			// create time
			tim, err := time.Parse(dateTimeFmt, dateStr)
			So(err, ShouldBeNil)

			to := DateTime(tim)

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
			jsonBytes := []byte(`{"name":"google","born":"2016-02-28T16:41:41.09Z"}`)
			var data = struct {
				Name string   `json:"name"`
				Born DateTime `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			So(err, ShouldBeNil)
			So(data.Born.String(), ShouldEqual, "2016-02-28T16:41:41.09Z")

			// marshal again
			b, err := json.Marshal(&data)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, string(jsonBytes))
		})
	})

}
