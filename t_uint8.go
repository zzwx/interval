package interval

import (
	"fmt"
	"strconv"
)

// WrapUint8 normalizes the value that "wraps around" within the [min,max) range.
// WrapUint8 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapUint8(min, max, value uint8) uint8 {
	min, max = MinMaxUint8(min, max)
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

// ClampUint8 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampUint8(min, max, value uint8) uint8 {
	min, max = MinMaxUint8(min, max)
	return MaxUint8(min, MinUint8(max, value))
}

// ValidateUint8 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateUint8(min, max, value uint8, minExclusive, maxExclusive bool) (uint8, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint8(min, max, minExclusive, maxExclusive)
	if !TestUint8(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringUint8(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestUint8 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestUint8(min, max, value uint8, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint8(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringUint8 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringUint8(min, max uint8, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint8(min, max, minExclusive, maxExclusive)
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

// MaxUint8 returns the bigger of two numbers.
func MaxUint8(x, y uint8) uint8 {
	if x < y {
		return y
	}
	return x
}

// MinUint8 returns the smaller of two numbers.
func MinUint8(x, y uint8) uint8 {
	if x > y {
		return y
	}
	return x
}

// MinMaxUint8 swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxUint8(min, max uint8) (uint8, uint8) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveUint8 swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveUint8(min, max uint8, minExclusive, maxExclusive bool) (uint8, uint8, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeUint8 struct {
	min          uint8
	max          uint8
	minExclusive bool
	maxExclusive bool
}

// NewRangeUint8 makes a new Range and returns its pointer. RangeUint8 can also be created with a RangeUint8{...} literal or new(RangeUint8).
func NewRangeUint8(min uint8, max uint8, minExclusive bool, maxExclusive bool) *RangeUint8 {
	return &RangeUint8{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeUint8) Wrap(value uint8) uint8 {
	return WrapUint8(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeUint8) Clamp(value uint8) uint8 {
	return ClampUint8(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeUint8) Validate(value uint8) (uint8, error) {
	return ValidateUint8(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeUint8) Test(value uint8) bool {
	return TestUint8(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// String returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
// String implements Stringer interface.
func (v RangeUint8) String() string {
	return ToStringUint8(v.min, v.max, v.minExclusive, v.maxExclusive)
}
