package interval

import (
	"fmt"
	"strconv"
)

// WrapInt32 normalizes the value that "wraps around" within the [min,max) range.
// WrapInt32 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapInt32(min, max, value int32) int32 {
	min, max = MinMaxInt32(min, max)
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

// ClampInt32 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampInt32(min, max, value int32) int32 {
	min, max = MinMaxInt32(min, max)
	return MaxInt32(min, MinInt32(max, value))
}

// ValidateInt32 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateInt32(min, max, value int32, minExclusive, maxExclusive bool) (int32, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt32(min, max, minExclusive, maxExclusive)
	if !TestInt32(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringInt32(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestInt32 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestInt32(min, max, value int32, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt32(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringInt32 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringInt32(min, max int32, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt32(min, max, minExclusive, maxExclusive)
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

// MaxInt32 returns the bigger of two numbers.
func MaxInt32(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

// MinInt32 returns the smaller of two numbers.
func MinInt32(x, y int32) int32 {
	if x > y {
		return y
	}
	return x
}

// MinMaxInt32 swaps x and y to assure that x <= y.
func MinMaxInt32(x, y int32) (int32, int32) {
	if x > y {
		return y, x
	}
	return x, y
}
// MinMaxExclusiveInt32 swaps x and y and minExclusive, maxExclusive to assure that x <= y together with the interval endings.
func MinMaxExclusiveInt32(x, y int32, minExclusive, maxExclusive bool) (int32, int32, bool, bool) {
	if x > y {
		return y, x, maxExclusive, minExclusive
	}
	return x, y, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeInt32 struct {
	min          int32
	max          int32
	minExclusive bool
	maxExclusive bool
}

// NewRangeInt32 makes a new Range and returns its pointer. RangeInt32 can also be created with a RangeInt32{...} literal or new(RangeInt32).
func NewRangeInt32(min int32, max int32, minExclusive bool, maxExclusive bool) *RangeInt32 {
	return &RangeInt32{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeInt32) Wrap(value int32) int32 {
	return WrapInt32(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeInt32) Clamp(value int32) int32 {
	return ClampInt32(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeInt32) Validate(value int32) (int32, error) {
	return ValidateInt32(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeInt32) Test(value int32) bool {
	return TestInt32(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeInt32) ToString() string {
	return ToStringInt32(v.min, v.max, v.minExclusive, v.maxExclusive)
}
