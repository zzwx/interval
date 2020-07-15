package interval

import (
	"fmt"
	"strconv"
)

// WrapInt16 normalizes the value that "wraps around" within the [min,max) range.
// WrapInt16 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapInt16(min, max, value int16) int16 {
	min, max = MinMaxInt16(min, max)
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

// ClampInt16 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampInt16(min, max, value int16) int16 {
	min, max = MinMaxInt16(min, max)
	return MaxInt16(min, MinInt16(max, value))
}

// ValidateInt16 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateInt16(min, max, value int16, minExclusive, maxExclusive bool) (int16, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt16(min, max, minExclusive, maxExclusive)
	if !TestInt16(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringInt16(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestInt16 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestInt16(min, max, value int16, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt16(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringInt16 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringInt16(min, max int16, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt16(min, max, minExclusive, maxExclusive)
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

// MaxInt16 returns the bigger of two numbers.
func MaxInt16(x, y int16) int16 {
	if x < y {
		return y
	}
	return x
}

// MinInt16 returns the smaller of two numbers.
func MinInt16(x, y int16) int16 {
	if x > y {
		return y
	}
	return x
}

// MinMaxInt16 swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxInt16(min, max int16) (int16, int16) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveInt16 swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveInt16(min, max int16, minExclusive, maxExclusive bool) (int16, int16, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeInt16 struct {
	min          int16
	max          int16
	minExclusive bool
	maxExclusive bool
}

// NewRangeInt16 makes a new Range and returns its pointer. RangeInt16 can also be created with a RangeInt16{...} literal or new(RangeInt16).
func NewRangeInt16(min int16, max int16, minExclusive bool, maxExclusive bool) *RangeInt16 {
	return &RangeInt16{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeInt16) Wrap(value int16) int16 {
	return WrapInt16(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeInt16) Clamp(value int16) int16 {
	return ClampInt16(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeInt16) Validate(value int16) (int16, error) {
	return ValidateInt16(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeInt16) Test(value int16) bool {
	return TestInt16(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeInt16) ToString() string {
	return ToStringInt16(v.min, v.max, v.minExclusive, v.maxExclusive)
}
