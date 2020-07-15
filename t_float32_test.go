package interval

import "fmt"

func ExampleWrapFloat32() {
	fmt.Printf("%v\n", WrapFloat32(0, 100, 120)) //=> 20
	fmt.Printf("%v\n", WrapFloat32(100, 0, 120)) //=> 20 (swapped range is fine)
	fmt.Printf("%v\n", WrapFloat32(0, 100, 0))   //=> 0
	fmt.Printf("%v\n", WrapFloat32(0, 100, 100)) //=> 0
	fmt.Printf("%v\n", WrapFloat32(0, 100, 101)) //=> 1
	fmt.Printf("%v\n", WrapFloat32(50, 100, 120)) //=> 70
	fmt.Printf("%v\n", WrapFloat32(50, 100, 10)) //=> 60
	fmt.Printf("%v\n", WrapFloat32(0, 100, -10)) //=> 90
	fmt.Printf("%v\n", WrapFloat32(0, 360, 361.5)) //=> 1.5
	fmt.Printf("%v\n", WrapFloat32(0, 360, -100)) //=> 260
	fmt.Printf("%v\n", WrapFloat32(0, 360, -720)) //=> 0
	fmt.Printf("%v\n", WrapFloat32(0, 360, 720.5)) //=> 0.5
	// Output:
	// 20
	// 20
	// 0
	// 0
	// 1
	// 70
	// 60
	// 90
	// 1.5
	// 260
	// 0
	// 0.5
}

func ExampleClampFloat32() {
	fmt.Printf("%v\n", ClampFloat32(0, 100, 120)) //=> 100
	fmt.Printf("%v\n", ClampFloat32(0, 100, -20)) //=> 0
	fmt.Printf("%v\n", ClampFloat32(100, 0, -20)) //=> 0 (even though min & max are swapped)
	// Output:
	// 100
	// 0
	// 0
}

func ExampleTestFloat32() {
	fmt.Printf("%v\n", TestFloat32(0, 100, 0, false, false))  //=> true
	fmt.Printf("%v\n", TestFloat32(0, 100, 0, true, false))   //=> false
	fmt.Printf("%v\n", TestFloat32(0, 100, 100, true, false)) //=> true
	fmt.Printf("%v\n", TestFloat32(100, 0, 100, false, true)) //=> true
	// Output: true
	// false
	// true
	// true
}

func ExampleValidateFloat32() {
	fmt.Println(ValidateFloat32(0, 100, 0, false, false))   //=> 0 <nil>
	fmt.Println(ValidateFloat32(0, 100, 0, true, false))    //=> 0 0 is outside of range [0, 100)
	fmt.Println(ValidateFloat32(0, 100, 100, true, false))  //=> 100 <nil>
	fmt.Println(ValidateFloat32(0, 100, 101, false, false)) //=> 0 101 is outside of range [0,100]
	fmt.Println(ValidateFloat32(0, 100, 100, true, true))   //=> 0 100 is outside of range (0,100)
	// Output:
	// 0 <nil>
	// 0 0 is outside of range (0,100]
	// 100 <nil>
	// 0 101 is outside of range [0,100]
	// 0 100 is outside of range (0,100)
}

func ExampleToStringFloat32() {
	fmt.Printf("%v\n", ToStringFloat32(0, 100, true, true))   //=> (0,100)
	fmt.Printf("%v\n", ToStringFloat32(0, 100, false, false)) //=> [0,100]
	fmt.Printf("%v\n", ToStringFloat32(0, 100, true, false))  //=> (0,100]
	fmt.Printf("%v\n", ToStringFloat32(0, 100, false, true))  //=> [0,100)
	// Output:
	// (0,100)
	// [0,100]
	// (0,100]
	// [0,100)
}

func ExampleRangeFloat32() {
	r := NewRangeFloat32(0, 100, false, false)
	fmt.Println(r.Wrap(120))     //=> 20
	fmt.Println(r.Validate(120)) //=> (0, error(120 is outside of range [0,100]))
	fmt.Println(r.Test(120))     //=> false
	fmt.Println(r)    //=> [0,100] (uses Stringer interface)
	r = NewRangeFloat32(100, 0, false, true) // swapped
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

func ExampleMinMaxFloat32() {
	x, y := MinMaxFloat32(0, 100)
	fmt.Printf("%v %v\n", x, y)   //=> 0 100
	x, y = MinMaxFloat32(100, 0) // swapped
  fmt.Printf("%v %v\n", x, y)   //=> 0 100
	// Output:
	// 0 100
	// 0 100
}

func ExampleMinMaxExclusiveFloat32() {
	x, y, xe, ye := MinMaxExclusiveFloat32(0, 100, true, false)
	fmt.Printf("%v %v %v %v\n", x, y, xe, ye)   //=> 0 100 true false
	x, y, xe, ye = MinMaxExclusiveFloat32(100, 0, false, true) // swapped
  fmt.Printf("%v %v %v %v\n", x, y, xe, ye)   //=> 0 100 true false
	// Output:
	// 0 100 true false
	// 0 100 true false
}
