package convert

import "strconv"

// 对接口返回的结果进行类型转换

type ConvertStr string

func (s ConvertStr) String() string {
	return string(s)
}

func (s ConvertStr) Int() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s ConvertStr) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s ConvertStr) UInt32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s ConvertStr) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}