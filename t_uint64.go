package interval

import (
	"fmt"
	"strconv"
)

// WrapUint64 normalizes the value that "wraps around" within the [min,max) range.
// WrapUint64 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapUint64(min, max, value uint64) uint64 {
	min, max = MinMaxUint64(min, max)
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

// ClampUint64 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampUint64(min, max, value uint64) uint64 {
	min, max = MinMaxUint64(min, max)
	return MaxUint64(min, MinUint64(max, value))
}

// ValidateUint64 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateUint64(min, max, value uint64, minExclusive, maxExclusive bool) (uint64, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint64(min, max, minExclusive, maxExclusive)
	if !TestUint64(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringUint64(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestUint64 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestUint64(min, max, value uint64, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint64(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringUint64 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringUint64(min, max uint64, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint64(min, max, minExclusive, maxExclusive)
	var minBracket = "["
	if minExclusive {
		minBracket = "("
	}
	var maxBracket = "]"
	if maxExclusive {
		maxBracket = ")"
	}
	return minBracket + strconv.FormatUint(uint64(min),10) + "," + strconv.FormatUint(uint64(max),10) + maxBracket
}

// MaxUint64 returns the bigger of two numbers.
func MaxUint64(x, y uint64) uint64 {
	if x < y {
		return y
	}
	return x
}

// MinUint64 returns the smaller of two numbers.
func MinUint64(x, y uint64) uint64 {
	if x > y {
		return y
	}
	return x
}

// MinMaxUint64 swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxUint64(min, max uint64) (uint64, uint64) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveUint64 swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveUint64(min, max uint64, minExclusive, maxExclusive bool) (uint64, uint64, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeUint64 struct {
	min          uint64
	max          uint64
	minExclusive bool
	maxExclusive bool
}

// NewRangeUint64 makes a new Range and returns its pointer. RangeUint64 can also be created with a RangeUint64{...} literal or new(RangeUint64).
func NewRangeUint64(min uint64, max uint64, minExclusive bool, maxExclusive bool) *RangeUint64 {
	return &RangeUint64{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeUint64) Wrap(value uint64) uint64 {
	return WrapUint64(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeUint64) Clamp(value uint64) uint64 {
	return ClampUint64(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeUint64) Validate(value uint64) (uint64, error) {
	return ValidateUint64(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeUint64) Test(value uint64) bool {
	return TestUint64(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeUint64) ToString() string {
	return ToStringUint64(v.min, v.max, v.minExclusive, v.maxExclusive)
}
