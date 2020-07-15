package interval

import (
	"fmt"
	"strconv"
)

// WrapFloat32 normalizes the value that "wraps around" within the [min,max) range.
// WrapFloat32 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapFloat32(min, max, value float32) float32 {
	min, max = MinMaxFloat32(min, max)
	var maxLessMin = max - min
	for value < min {
	  value += maxLessMin
	}
	for value >= max {
		value -= maxLessMin
	}
	return value
	//Original implementation is not working well for unsigned types:
	//return ((value-min)%maxLessMin+maxLessMin)%maxLessMin + min
}

// ClampFloat32 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampFloat32(min, max, value float32) float32 {
	min, max = MinMaxFloat32(min, max)
	return MaxFloat32(min, MinFloat32(max, value))
}

// ValidateFloat32 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateFloat32(min, max, value float32, minExclusive, maxExclusive bool) (float32, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveFloat32(min, max, minExclusive, maxExclusive)
	if !TestFloat32(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringFloat32(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestFloat32 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestFloat32(min, max, value float32, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveFloat32(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringFloat32 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringFloat32(min, max float32, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveFloat32(min, max, minExclusive, maxExclusive)
	var minBracket = "["
	if minExclusive {
		minBracket = "("
	}
	var maxBracket = "]"
	if maxExclusive {
		maxBracket = ")"
	}
	return minBracket + strconv.FormatFloat(float64(min),'f', -1, 32) + "," + strconv.FormatFloat(float64(max),'f', -1, 32) + maxBracket
}

// MaxFloat32 returns the bigger of two numbers.
func MaxFloat32(x, y float32) float32 {
	if x < y {
		return y
	}
	return x
}

// MinFloat32 returns the smaller of two numbers.
func MinFloat32(x, y float32) float32 {
	if x > y {
		return y
	}
	return x
}

// MinMaxFloat32 swaps x and y to assure that x <= y.
func MinMaxFloat32(x, y float32) (float32, float32) {
	if x > y {
		return y, x
	}
	return x, y
}
// MinMaxExclusiveFloat32 swaps x and y and minExclusive, maxExclusive to assure that x <= y together with the interval endings.
func MinMaxExclusiveFloat32(x, y float32, minExclusive, maxExclusive bool) (float32, float32, bool, bool) {
	if x > y {
		return y, x, maxExclusive, minExclusive
	}
	return x, y, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeFloat32 struct {
	min          float32
	max          float32
	minExclusive bool
	maxExclusive bool
}

// NewRangeFloat32 makes a new Range and returns its pointer. RangeFloat32 can also be created with a RangeFloat32{...} literal or new(RangeFloat32).
func NewRangeFloat32(min float32, max float32, minExclusive bool, maxExclusive bool) *RangeFloat32 {
	return &RangeFloat32{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeFloat32) Wrap(value float32) float32 {
	return WrapFloat32(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeFloat32) Clamp(value float32) float32 {
	return ClampFloat32(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeFloat32) Validate(value float32) (float32, error) {
	return ValidateFloat32(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeFloat32) Test(value float32) bool {
	return TestFloat32(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeFloat32) ToString() string {
	return ToStringFloat32(v.min, v.max, v.minExclusive, v.maxExclusive)
}