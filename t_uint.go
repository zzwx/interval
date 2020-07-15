package interval

import (
	"fmt"
	"strconv"
)

// WrapUint normalizes the value that "wraps around" within the [min,max) range.
// WrapUint always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapUint(min, max, value uint) uint {
	min, max = MinMaxUint(min, max)
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

// ClampUint returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampUint(min, max, value uint) uint {
	min, max = MinMaxUint(min, max)
	return MaxUint(min, MinUint(max, value))
}

// ValidateUint tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateUint(min, max, value uint, minExclusive, maxExclusive bool) (uint, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint(min, max, minExclusive, maxExclusive)
	if !TestUint(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringUint(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestUint returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestUint(min, max, value uint, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringUint returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringUint(min, max uint, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveUint(min, max, minExclusive, maxExclusive)
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

// MaxUint returns the bigger of two numbers.
func MaxUint(x, y uint) uint {
	if x < y {
		return y
	}
	return x
}

// MinUint returns the smaller of two numbers.
func MinUint(x, y uint) uint {
	if x > y {
		return y
	}
	return x
}

// MinMaxUint swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxUint(min, max uint) (uint, uint) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveUint swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveUint(min, max uint, minExclusive, maxExclusive bool) (uint, uint, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeUint struct {
	min          uint
	max          uint
	minExclusive bool
	maxExclusive bool
}

// NewRangeUint makes a new Range and returns its pointer. RangeUint can also be created with a RangeUint{...} literal or new(RangeUint).
func NewRangeUint(min uint, max uint, minExclusive bool, maxExclusive bool) *RangeUint {
	return &RangeUint{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeUint) Wrap(value uint) uint {
	return WrapUint(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeUint) Clamp(value uint) uint {
	return ClampUint(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeUint) Validate(value uint) (uint, error) {
	return ValidateUint(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeUint) Test(value uint) bool {
	return TestUint(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeUint) ToString() string {
	return ToStringUint(v.min, v.max, v.minExclusive, v.maxExclusive)
}
