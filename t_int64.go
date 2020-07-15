package interval

import (
	"fmt"
	"strconv"
)

// WrapInt64 normalizes the value that "wraps around" within the [min,max) range.
// WrapInt64 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapInt64(min, max, value int64) int64 {
	min, max = MinMaxInt64(min, max)
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

// ClampInt64 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampInt64(min, max, value int64) int64 {
	min, max = MinMaxInt64(min, max)
	return MaxInt64(min, MinInt64(max, value))
}

// ValidateInt64 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateInt64(min, max, value int64, minExclusive, maxExclusive bool) (int64, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt64(min, max, minExclusive, maxExclusive)
	if !TestInt64(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringInt64(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestInt64 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestInt64(min, max, value int64, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt64(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringInt64 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringInt64(min, max int64, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveInt64(min, max, minExclusive, maxExclusive)
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

// MaxInt64 returns the bigger of two numbers.
func MaxInt64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// MinInt64 returns the smaller of two numbers.
func MinInt64(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

// MinMaxInt64 swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxInt64(min, max int64) (int64, int64) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveInt64 swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveInt64(min, max int64, minExclusive, maxExclusive bool) (int64, int64, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeInt64 struct {
	min          int64
	max          int64
	minExclusive bool
	maxExclusive bool
}

// NewRangeInt64 makes a new Range and returns its pointer. RangeInt64 can also be created with a RangeInt64{...} literal or new(RangeInt64).
func NewRangeInt64(min int64, max int64, minExclusive bool, maxExclusive bool) *RangeInt64 {
	return &RangeInt64{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeInt64) Wrap(value int64) int64 {
	return WrapInt64(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeInt64) Clamp(value int64) int64 {
	return ClampInt64(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeInt64) Validate(value int64) (int64, error) {
	return ValidateInt64(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeInt64) Test(value int64) bool {
	return TestInt64(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// String returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
// String implements Stringer interface.
func (v RangeInt64) String() string {
	return ToStringInt64(v.min, v.max, v.minExclusive, v.maxExclusive)
}
