package namesgenerator

import (
	"math/rand"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNameFormat(t *testing.T) {
	Convey("Test format", t, func() {
		name := GetRandomName(0)
		So(strings.Contains(name, " "), ShouldBeTrue)
		So(strings.ContainsAny(name, "0123456789"), ShouldBeFalse)
	})
}

func TestNameRetries(t *testing.T) {
	Convey("Test retries", t, func() {
		name := GetRandomName(1)
		So(strings.Contains(name, " "), ShouldBeTrue)
		So(strings.ContainsAny(name, "0123456789"), ShouldBeTrue)
	})
}

func TestWozniak(t *testing.T) {
	Convey("Boring Wozniak", t, func() {
		rand.Seed(-96681)
		name := GetRandomName(0)
		So(name, ShouldNotEqual, "Boring Wozniak")
	})
}
