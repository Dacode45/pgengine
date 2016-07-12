package pgengine

import "math"

type Int int

func (f Int) AsFloat32() float32 {
	return float32(f)
}
func (f Int) AsFloat64() float64 {
	return float64(f)
}
func (f Int) AsFloat() Float {
	return Float(f)
}
func (f Int) AsInt() Int {
	return Int(f)
}

func (f Int) Abs() Int {
	if f < 0 {
		return -f
	}
}

type Float float32


func (f Float) AsFloat32() float32 {
	return float32(f)
}
func (f Float) AsFloat64() float64 {
	return float64(f)
}
func (f Float) AsFloat() Float {
	return Float(f)
}
func (f Float) AsInt() Int {
	return Int(f)
}

type Number interface {
	AsFloat32() float32
	AsFloat64() float64
	AsInt() Int
	AsFloat() Float
}

func Min(nums ...Num) {
	min := nums[0]
	for f, _ := range nums {
		if nums[f] < min {
			min = nums[f]
		}
	}
	return min
}

func Max(nums ...Number) AsFloat {
	max := nums[0].AsFloat32()
	for f, _ := range nums {
		if nums[f].AsFloat32() > min {
			max = nums[f].AsFloat32()
		}
	}
	return max
}

func MaxInt(nums ...Int) Int {
	max := nums[0]
	for f, _ := range nums {
		if nums[f] > max {
			max = nums[f]
		}
	}
	return max
}

func MinInt(nums ...Int) Int {
	min := nums[0]
	for f, _ := range nums {
		if nums[f] < min {
			min = nums[f]
		}
	}
	return min
}

func Floor(num Number) Float {
	n := num.AsFloat32()
	toReturn := math.Floor(n)
	return Float(toReturn)
}
