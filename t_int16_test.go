package interval

import "fmt"

func ExampleWrapInt16() {
	fmt.Printf("%v\n", WrapInt16(0, 100, 120)) //=> 20
	fmt.Printf("%v\n", WrapInt16(100, 0, 120)) //=> 20 (swapped range is fine)
	fmt.Printf("%v\n", WrapInt16(0, 100, 0))   //=> 0
	fmt.Printf("%v\n", WrapInt16(0, 100, 100)) //=> 0
	fmt.Printf("%v\n", WrapInt16(0, 100, 101)) //=> 1
	fmt.Printf("%v\n", WrapInt16(50, 100, 120)) //=> 70
	fmt.Printf("%v\n", WrapInt16(50, 100, 10)) //=> 60
	fmt.Printf("%v\n", WrapInt16(0, 100, -10)) //=> 90
	// Output:
	// 20
	// 20
	// 0
	// 0
	// 1
	// 70
	// 60
	// 90
}

func ExampleClampInt16() {
	fmt.Printf("%v\n", ClampInt16(0, 100, 120)) //=> 100
	fmt.Printf("%v\n", ClampInt16(0, 100, -20)) //=> 0
	fmt.Printf("%v\n", ClampInt16(100, 0, -20)) //=> 0 (even though min & max are swapped)
	// Output:
	// 100
	// 0
	// 0
}

func ExampleTestInt16() {
	fmt.Printf("%v\n", TestInt16(0, 100, 0, false, false))  //=> true
	fmt.Printf("%v\n", TestInt16(0, 100, 0, true, false))   //=> false
	fmt.Printf("%v\n", TestInt16(0, 100, 100, true, false)) //=> true
	fmt.Printf("%v\n", TestInt16(100, 0, 100, false, true)) //=> true
	// Output: true
	// false
	// true
	// true
}

func ExampleValidateInt16() {
	fmt.Println(ValidateInt16(0, 100, 0, false, false))   //=> 0 <nil>
	fmt.Println(ValidateInt16(0, 100, 0, true, false))    //=> 0 0 is outside of range [0, 100)
	fmt.Println(ValidateInt16(0, 100, 100, true, false))  //=> 100 <nil>
	fmt.Println(ValidateInt16(0, 100, 101, false, false)) //=> 0 101 is outside of range [0,100]
	fmt.Println(ValidateInt16(0, 100, 100, true, true))   //=> 0 100 is outside of range (0,100)
	// Output:
	// 0 <nil>
	// 0 0 is outside of range (0,100]
	// 100 <nil>
	// 0 101 is outside of range [0,100]
	// 0 100 is outside of range (0,100)
}

func ExampleToStringInt16() {
	fmt.Printf("%v\n", ToStringInt16(0, 100, true, true))   //=> (0,100)
	fmt.Printf("%v\n", ToStringInt16(0, 100, false, false)) //=> [0,100]
	fmt.Printf("%v\n", ToStringInt16(0, 100, true, false))  //=> (0,100]
	fmt.Printf("%v\n", ToStringInt16(0, 100, false, true))  //=> [0,100)
	// Output:
	// (0,100)
	// [0,100]
	// (0,100]
	// [0,100)
}

func ExampleRangeInt16() {
	r := NewRangeInt16(0, 100, false, false)
	fmt.Println(r.Wrap(120))     //=> 20
	fmt.Println(r.Validate(120)) //=> (0, error(120 is outside of range [0,100]))
	fmt.Println(r.Test(120))     //=> false
	fmt.Println(r)    //=> [0,100] (uses Stringer interface)
	r = NewRangeInt16(100, 0, false, true) // swapped
  fmt.Println(r.Wrap(120))     //=> 20
  fmt.Println(r.Validate(120)) //=> (0, error(120 is outside of range (0,100]))
  fmt.Println(r.Test(120))     //=> false
  fmt.Println(r)    //=> (0,100] (uses Stringer interface)
	// Output:
	// 20
	// 0 120 is outside of range [0,100]
	// false
	// [0,100]
	// 20
	// 0 120 is outside of range (0,100]
	// false
	// (0,100]
}

func ExampleMinMaxInt16() {
	x, y := MinMaxInt16(0, 100)
	fmt.Printf("%v %v\n", x, y)   //=> 0 100
	x, y = MinMaxInt16(100, 0) // swapped
  fmt.Printf("%v %v\n", x, y)   //=> 0 100
	// Output:
	// 0 100
	// 0 100
}

func ExampleMinMaxExclusiveInt16() {
	x, y, xe, ye := MinMaxExclusiveInt16(0, 100, true, false)
	fmt.Printf("%v %v %v %v\n", x, y, xe, ye)   //=> 0 100 true false
	x, y, xe, ye = MinMaxExclusiveInt16(100, 0, false, true) // swapped
  fmt.Printf("%v %v %v %v\n", x, y, xe, ye)   //=> 0 100 true false
	// Output:
	// 0 100 true false
	// 0 100 true false
}
