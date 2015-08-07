package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinData(t *testing.T) {
	Convey("It should access assets in bindata", t, func() {
		common, err := Asset(CommonScriptPath)
		So(err, ShouldBeNil)

		result, err := runScript(common, AnalyticsScriptPrefix+"poi")
		So(err, ShouldBeNil)

		So(result.Type, ShouldEqual, "boolean")
		So(result.Name, ShouldEqual, "postgresql installed")
		So(result.Value, ShouldEqual, true)

		So(result.Error, ShouldEqual, "")
	})
}
