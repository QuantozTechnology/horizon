package stellarbase

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHashing(t *testing.T) {
	Convey("Given a message to hash", t, func() {
		message := []byte("hello world")
		expected := [32]byte{
			0xB9, 0x4D, 0x27, 0xB9,
			0x93, 0x4D, 0x3E, 0x08,
			0xA5, 0x2E, 0x52, 0xD7,
			0xDA, 0x7D, 0xAB, 0xFA,
			0xC4, 0x84, 0xEF, 0xE3,
			0x7A, 0x53, 0x80, 0xEE,
			0x90, 0x88, 0xF7, 0xAC,
			0xE2, 0xEF, 0xCD, 0xE9,
		}

		Convey("When hashing the message", func() {
			actual := Hash(message)

			Convey("The hash should be correct", func() {
				So(actual, ShouldResemble, expected)
			})
		})

	})
}
