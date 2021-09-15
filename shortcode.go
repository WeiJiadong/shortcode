package shortcode

import (
	"fmt"
	"strings"
)

var (
	// charSet hash词典
	charSet = []byte{
		'8', 'U', 'C', 'p', '9', 'G', 'B', 'A', 'q', 'g', 'n', 'Z', 'Y', '6', 'P', 'J', 'K',
		'c', 'L', '4', '5', 'e', 'a', 'x', 'd', 'W', '7', 'i', 'V', 'f', 'N', 'E', 'u', 'h',
		'v', 'H', 'I', 'y', 't', 'j', 'r', 'M', 'T', 'F', 'w', 'k', 'Q', 'm', 'X', 'b', 'R',
		'3', 's', 'D', '2', 'S', 'z', '0', 'O', 'o', '1', 'l',
	}
	// charSetLen hash词典长度
	charSetLen = uint64(len(charSet))
)

const (
	// defaultShortCodeLen 默认短码长度
	defaultShortCodeLen = 6

	slat  = 1610612741
	prime = 3 //prime和charSetLen互质, 可保证 (cid * prime) % L 在 [0,charSetLen)上均匀分布
)

// shortCodeBuilder 短码生成器
type shortCodeBuilder struct {
	len int // 生成码的长度
}

// shortCodeOpt shortCodeBuilder helper
type shortCodeOpt func(code *shortCodeBuilder)

// WithCodeLen 设置生成的码长
func WithCodeLen(len int) shortCodeOpt {
	return func(code *shortCodeBuilder) {
		if len < 0 {
			panic(fmt.Errorf("WithCodeLen error, len:%d", len))
		}
		code.len = len
	}
}

// NewShortCodeBuilder shortCodeBuilder 构造函数
func NewShortCodeBuilder(opt ...shortCodeOpt) *shortCodeBuilder {
	sl := &shortCodeBuilder{len: defaultShortCodeLen}
	for i := range opt {
		opt[i](sl)
	}

	return sl
}

//  Do 执行转换
func (c *shortCodeBuilder) Do(cid uint64) (string, error) {
	id := cid*prime + slat
	arr := make([]uint64, c.len+1)
	arr[0] = id
	for i := 1; i < c.len+1; i++ {
		arr[i] = arr[i-1] / charSetLen
		arr[i-1] = (arr[i-1] + uint64(i-1)*arr[0]) % charSetLen
	}

	sb := strings.Builder{}
	for i := 0; i < c.len; i++ {
		if et := sb.WriteByte(charSet[int(arr[i])]); et != nil {
			fmt.Print("sb.WriteByte error, et:", et)
			return "", et
		}
	}

	return sb.String(), nil
}