package shortcode

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNewShortCode(t *testing.T) {
	convey.Convey("default value check", t, func() {
		sl := NewShortCodeBuilder()
		fmt.Println("")
		fmt.Println(sl.Do(1))
		fmt.Println(sl.Do(101))
		fmt.Println(sl.Do(1001))
		fmt.Println(sl.Do(10001))
		fmt.Println(sl.Do(100001))
		fmt.Println(sl.Do(1000001))
		fmt.Println(sl.Do(10000001))
		convey.So(true, convey.ShouldBeTrue)
	})

	convey.Convey("set value check", t, func() {
		sl := NewShortCodeBuilder(WithCodeLen(3))
		fmt.Println("")
		fmt.Println(sl.Do(1))
		fmt.Println(sl.Do(101))
		fmt.Println(sl.Do(1001))
		fmt.Println(sl.Do(10001))
		fmt.Println(sl.Do(100001))
		fmt.Println(sl.Do(1000001))
		fmt.Println(sl.Do(10000001))
		convey.So(true, convey.ShouldBeTrue)
	})

	convey.Convey("conflict check two day", t, func() {
		length := 6
		sl := NewShortCodeBuilder(WithCodeLen(length))
		codeMap := make(map[string]uint64)
		var i uint64 = 1629557050
		flag := true
		for ; i < 1629729850; i++ {
			s, _ := sl.Do(i)
			id, ex := codeMap[s]
			if ex {
				fmt.Printf("confilict %d, %d\n", id, i)
				flag = false
			}
			if len(s) != length {
				fmt.Printf("confilict %s, %d\n", s, i)
			}
			codeMap[s] = i
		}
		convey.So(flag, convey.ShouldBeTrue)
	})

	convey.Convey("conflict check uid:100000000-150000000", t, func() {
		length := 6
		sl := NewShortCodeBuilder(WithCodeLen(length))
		codeMap := make(map[string]uint64)
		var i uint64 = 100000000
		flag := true
		for ; i < 150000000; i++ {
			s, _ := sl.Do(i)
			id, ex := codeMap[s]
			if ex {
				fmt.Printf("confilict %d, %d\n", id, i)
				flag = false
			}
			if len(s) != length {
				fmt.Printf("confilict %s, %d\n", s, i)
			}
			codeMap[s] = i
		}
		convey.So(flag, convey.ShouldBeTrue)
	})

	convey.Convey("conflict check uid:150000000-200000000", t, func() {
		length := 6
		sl := NewShortCodeBuilder(WithCodeLen(length))
		codeMap := make(map[string]uint64)
		var i uint64 = 150000000
		flag := true
		for ; i < 200000000; i++ {
			s, _ := sl.Do(i)
			id, ex := codeMap[s]
			if ex {
				fmt.Printf("confilict %d, %d\n", id, i)
				flag = false
			}
			if len(s) != length {
				fmt.Printf("confilict %s, %d\n", s, i)
			}
			codeMap[s] = i
		}
		convey.So(flag, convey.ShouldBeTrue)
	})
}