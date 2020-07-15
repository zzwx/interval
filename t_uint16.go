package interval

import (
	"fmt"
	"strconv"
)

// WrapUint16 normalizes the value that "wraps around" within the [min,max) range.
// WrapUint16 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapUint16(min, max, value uint16) uint16 {
	min, max = MinMaxUint16(min, max)
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

// ClampUint16 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampUint16(min, max, value uint16) uint16 {
	min, max = MinMaxUint16(min, max)
	return MaxUint16(min, MinUint16(max, value))
}

// ValidateUint16 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateUint16(min, max, value uint16, minExclusive, maxExclusive bool) (uint16, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint16(min, max, minExclusive, maxExclusive)
	if !TestUint16(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringUint16(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestUint16 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestUint16(min, max, value uint16, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint16(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringUint16 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringUint16(min, max uint16, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint16(min, max, minExclusive, maxExclusive)
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

// MaxUint16 returns the bigger of two numbers.
func MaxUint16(x, y uint16) uint16 {
	if x < y {
		return y
	}
	return x
}

// MinUint16 returns the smaller of two numbers.
func MinUint16(x, y uint16) uint16 {
	if x > y {
		return y
	}
	return x
}

// MinMaxUint16 swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxUint16(min, max uint16) (uint16, uint16) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveUint16 swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveUint16(min, max uint16, minExclusive, maxExclusive bool) (uint16, uint16, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeUint16 struct {
	min          uint16
	max          uint16
	minExclusive bool
	maxExclusive bool
}

// NewRangeUint16 makes a new Range and returns its pointer. RangeUint16 can also be created with a RangeUint16{...} literal or new(RangeUint16).
func NewRangeUint16(min uint16, max uint16, minExclusive bool, maxExclusive bool) *RangeUint16 {
	return &RangeUint16{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeUint16) Wrap(value uint16) uint16 {
	return WrapUint16(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeUint16) Clamp(value uint16) uint16 {
	return ClampUint16(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeUint16) Validate(value uint16) (uint16, error) {
	return ValidateUint16(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeUint16) Test(value uint16) bool {
	return TestUint16(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeUint16) ToString() string {
	return ToStringUint16(v.min, v.max, v.minExclusive, v.maxExclusive)
}
