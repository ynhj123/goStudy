package heap

func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}

func (self *Object) Shorts() []int16 {
	return self.data.([]int16)
}

func (self *Object) Ints() []int32 {
	return self.data.([]int32)
}

func (self *Object) Longs() []int64 {
	return self.data.([]int64)
}
func (self *Object) Chars() []uint16 {
	return self.data.([]uint16)
}

func (self *Object) Floats() []float32 {
	return self.data.([]float32)
}

func (self *Object) Doubles() []float64 {
	return self.data.([]float64)
}

func (self *Object) Refs() []*Object {
	return self.data.([]*Object)
}

func (self *Object) ArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.Bytes()))
	case []int16:
		return int32(len(self.Shorts()))
	case []int32:
		return int32(len(self.Ints()))
	case []int64:
		return int32(len(self.Longs()))
	case []uint16:
		return int32(len(self.Chars()))
	case []float32:
		return int32(len(self.Floats()))
	case []float64:
		return int32(len(self.Doubles()))
	case []*Object:
		return int32(len(self.Refs()))
	default:
		panic("No array!")
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dst.data.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dst.data.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos+length]
		_dst := dst.data.([]int64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dst.data.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dst.data.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}
