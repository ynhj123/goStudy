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
