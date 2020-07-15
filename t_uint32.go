package interval

import (
	"fmt"
	"strconv"
)

// WrapUint32 normalizes the value that "wraps around" within the [min,max) range.
// WrapUint32 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapUint32(min, max, value uint32) uint32 {
	min, max = MinMaxUint32(min, max)
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

// ClampUint32 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampUint32(min, max, value uint32) uint32 {
	min, max = MinMaxUint32(min, max)
	return MaxUint32(min, MinUint32(max, value))
}

// ValidateUint32 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateUint32(min, max, value uint32, minExclusive, maxExclusive bool) (uint32, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint32(min, max, minExclusive, maxExclusive)
	if !TestUint32(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringUint32(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestUint32 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestUint32(min, max, value uint32, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint32(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringUint32 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringUint32(min, max uint32, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint32(min, max, minExclusive, maxExclusive)
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

// MaxUint32 returns the bigger of two numbers.
func MaxUint32(x, y uint32) uint32 {
	if x < y {
		return y
	}
	return x
}

// MinUint32 returns the smaller of two numbers.
func MinUint32(x, y uint32) uint32 {
	if x > y {
		return y
	}
	return x
}

// MinMaxUint32 swaps x and y to assure that x <= y.
func MinMaxUint32(x, y uint32) (uint32, uint32) {
	if x > y {
		return y, x
	}
	return x, y
}
// MinMaxExclusiveUint32 swaps x and y and minExclusive, maxExclusive to assure that x <= y together with the interval endings.
func MinMaxExclusiveUint32(x, y uint32, minExclusive, maxExclusive bool) (uint32, uint32, bool, bool) {
	if x > y {
		return y, x, maxExclusive, minExclusive
	}
	return x, y, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeUint32 struct {
	min          uint32
	max          uint32
	minExclusive bool
	maxExclusive bool
}

// NewRangeUint32 makes a new Range and returns its pointer. RangeUint32 can also be created with a RangeUint32{...} literal or new(RangeUint32).
func NewRangeUint32(min uint32, max uint32, minExclusive bool, maxExclusive bool) *RangeUint32 {
	return &RangeUint32{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeUint32) Wrap(value uint32) uint32 {
	return WrapUint32(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeUint32) Clamp(value uint32) uint32 {
	return ClampUint32(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeUint32) Validate(value uint32) (uint32, error) {
	return ValidateUint32(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeUint32) Test(value uint32) bool {
	return TestUint32(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeUint32) ToString() string {
	return ToStringUint32(v.min, v.max, v.minExclusive, v.maxExclusive)
}
