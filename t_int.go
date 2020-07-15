package interval

import (
	"fmt"
	"strconv"
)

// WrapInt normalizes the value that "wraps around" within the [min,max) range.
// WrapInt always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapInt(min, max, value int) int {
	min, max = MinMaxInt(min, max)
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

// ClampInt returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampInt(min, max, value int) int {
	min, max = MinMaxInt(min, max)
	return MaxInt(min, MinInt(max, value))
}

// ValidateInt tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateInt(min, max, value int, minExclusive, maxExclusive bool) (int, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt(min, max, minExclusive, maxExclusive)
	if !TestInt(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringInt(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestInt returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestInt(min, max, value int, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringInt returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringInt(min, max int, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt(min, max, minExclusive, maxExclusive)
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

// MaxInt returns the bigger of two numbers.
func MaxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// MinInt returns the smaller of two numbers.
func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// MinMaxInt swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxInt(min, max int) (int, int) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveInt swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveInt(min, max int, minExclusive, maxExclusive bool) (int, int, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeInt struct {
	min          int
	max          int
	minExclusive bool
	maxExclusive bool
}

// NewRangeInt makes a new Range and returns its pointer. RangeInt can also be created with a RangeInt{...} literal or new(RangeInt).
func NewRangeInt(min int, max int, minExclusive bool, maxExclusive bool) *RangeInt {
	return &RangeInt{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeInt) Wrap(value int) int {
	return WrapInt(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeInt) Clamp(value int) int {
	return ClampInt(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeInt) Validate(value int) (int, error) {
	return ValidateInt(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeInt) Test(value int) bool {
	return TestInt(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeInt) ToString() string {
	return ToStringInt(v.min, v.max, v.minExclusive, v.maxExclusive)
}
