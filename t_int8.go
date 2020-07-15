package interval

import (
	"fmt"
	"strconv"
)

// WrapInt8 normalizes the value that "wraps around" within the [min,max) range.
// WrapInt8 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapInt8(min, max, value int8) int8 {
	min, max = MinMaxInt8(min, max)
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

// ClampInt8 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampInt8(min, max, value int8) int8 {
	min, max = MinMaxInt8(min, max)
	return MaxInt8(min, MinInt8(max, value))
}

// ValidateInt8 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateInt8(min, max, value int8, minExclusive, maxExclusive bool) (int8, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt8(min, max, minExclusive, maxExclusive)
	if !TestInt8(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringInt8(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestInt8 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestInt8(min, max, value int8, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt8(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringInt8 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringInt8(min, max int8, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt8(min, max, minExclusive, maxExclusive)
	var minBracket = "["
	if minExclusive {
		minBracket = "("
	}
	var maxBracket = "]"
	if maxExclusive {
		maxBracket = ")"
	}
	return minBracket + strconv.FormatInt(int64(min),10) + "," + strconv.FormatInt(int64(max),10) + maxBracket
}

// MaxInt8 returns the bigger of two numbers.
func MaxInt8(x, y int8) int8 {
	if x < y {
		return y
	}
	return x
}

// MinInt8 returns the smaller of two numbers.
func MinInt8(x, y int8) int8 {
	if x > y {
		return y
	}
	return x
}

// MinMaxInt8 swaps x and y to assure that x <= y.
func MinMaxInt8(x, y int8) (int8, int8) {
	if x > y {
		return y, x
	}
	return x, y
}
// MinMaxExclusiveInt8 swaps x and y and minExclusive, maxExclusive to assure that x <= y together with the interval endings.
func MinMaxExclusiveInt8(x, y int8, minExclusive, maxExclusive bool) (int8, int8, bool, bool) {
	if x > y {
		return y, x, maxExclusive, minExclusive
	}
	return x, y, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeInt8 struct {
	min          int8
	max          int8
	minExclusive bool
	maxExclusive bool
}

// NewRangeInt8 makes a new Range and returns its pointer. RangeInt8 can also be created with a RangeInt8{...} literal or new(RangeInt8).
func NewRangeInt8(min int8, max int8, minExclusive bool, maxExclusive bool) *RangeInt8 {
	return &RangeInt8{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeInt8) Wrap(value int8) int8 {
	return WrapInt8(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeInt8) Clamp(value int8) int8 {
	return ClampInt8(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeInt8) Validate(value int8) (int8, error) {
	return ValidateInt8(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeInt8) Test(value int8) bool {
	return TestInt8(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeInt8) ToString() string {
	return ToStringInt8(v.min, v.max, v.minExclusive, v.maxExclusive)
}
