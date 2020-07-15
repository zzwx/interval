package interval

import (
	"fmt"
	"strconv"
)

// WrapFloat64 normalizes the value that "wraps around" within the [min,max) range.
// WrapFloat64 always assumes that `min` is [inclusive, and `max` is exclusive).
func WrapFloat64(min, max, value float64) float64 {
	min, max = MinMaxFloat64(min, max)
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

// ClampFloat64 returns value capped to [min,max] range. Both ends of this range are inclusive.
func ClampFloat64(min, max, value float64) float64 {
	min, max = MinMaxFloat64(min, max)
	return MaxFloat64(min, MinFloat64(max, value))
}

// ValidateFloat64 tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func ValidateFloat64(min, max, value float64, minExclusive, maxExclusive bool) (float64, error) {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveFloat64(min, max, minExclusive, maxExclusive)
	if !TestFloat64(min, max, value, minExclusive, maxExclusive) {
		return 0, fmt.Errorf("%v is outside of range %v", value, ToStringFloat64(min, max, minExclusive, maxExclusive))
	}
	return value, nil
}

// TestFloat64 returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func TestFloat64(min, max, value float64, minExclusive, maxExclusive bool) bool {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveFloat64(min, max, minExclusive, maxExclusive)
	return !(value < min || value > max || (maxExclusive && (value == max)) || (minExclusive && (value == min)))
}

// ToStringFloat64 returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func ToStringFloat64(min, max float64, minExclusive, maxExclusive bool) string {
	min, max, minExclusive, maxExclusive = MinMaxExclusiveFloat64(min, max, minExclusive, maxExclusive)
	var minBracket = "["
	if minExclusive {
		minBracket = "("
	}
	var maxBracket = "]"
	if maxExclusive {
		maxBracket = ")"
	}
	return minBracket + strconv.FormatFloat(float64(min),'f', -1, 64) + "," + strconv.FormatFloat(float64(max),'f', -1, 64) + maxBracket
}

// MaxFloat64 returns the bigger of two numbers.
func MaxFloat64(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

// MinFloat64 returns the smaller of two numbers.
func MinFloat64(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

// MinMaxFloat64 swaps min and max to assure that min < max.
// It is automatically called for all the rest of the functions
// that do not expect minExclusive or maxExclusive.
func MinMaxFloat64(min, max float64) (float64, float64) {
	if min > max {
		return max, min
	}
	return min, max
}
// MinMaxExclusiveFloat64 swaps min and max as well as minExclusive, maxExclusive correspondingly
// to assure that min < max together with the interval endings. It is automatically called for all
// the rest of the functions that expect minExclusive or maxExclusive.
func MinMaxExclusiveFloat64(min, max float64, minExclusive, maxExclusive bool) (float64, float64, bool, bool) {
	if min > max {
		return max, min, maxExclusive, minExclusive
	}
	return min, max, minExclusive, maxExclusive
}

// Range represents a struct containing all the fields defining a range.
type RangeFloat64 struct {
	min          float64
	max          float64
	minExclusive bool
	maxExclusive bool
}

// NewRangeFloat64 makes a new Range and returns its pointer. RangeFloat64 can also be created with a RangeFloat64{...} literal or new(RangeFloat64).
func NewRangeFloat64(min float64, max float64, minExclusive bool, maxExclusive bool) *RangeFloat64 {
	return &RangeFloat64{min: min, max: max, minExclusive: minExclusive, maxExclusive: maxExclusive}
}

// Wrap does not obey minExclusive and maxExclusive and always assumes [min,max) range.
func (v RangeFloat64) Wrap(value float64) float64 {
	return WrapFloat64(v.min, v.max, value)
}

// Clamp does not obey minExclusive and maxExclusive and always assumes [min,max] range.
func (v RangeFloat64) Clamp(value float64) float64 {
	return ClampFloat64(v.min, v.max, value)
}

// Validate tests whether the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
// It returns the value if it is within the range, otherwise returns 0 and error.
func (v RangeFloat64) Validate(value float64) (float64, error) {
	return ValidateFloat64(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// Test returns true if the value is within [(min,max)] range (depending on minExclusive and maxExclusive).
func (v RangeFloat64) Test(value float64) bool {
	return TestFloat64(v.min, v.max, value, v.minExclusive, v.maxExclusive)
}

// ToString returns a string representation of the range using range notation
// (https://en.wikipedia.org/wiki/Interval_(mathematics)#Classification_of_intervals).
func (v RangeFloat64) ToString() string {
	return ToStringFloat64(v.min, v.max, v.minExclusive, v.maxExclusive)
}
