package interval

// Code generated by generate/generate.go. DO NOT EDIT

import "fmt"

func ExampleWrapInt32() {
	fmt.Printf("%v\n", WrapInt32(0, 100, 120))  //=> 20
	fmt.Printf("%v\n", WrapInt32(100, 0, 120))  //=> 20 (swapped range is fine)
	fmt.Printf("%v\n", WrapInt32(0, 100, 0))    //=> 0
	fmt.Printf("%v\n", WrapInt32(0, 100, 100))  //=> 0
	fmt.Printf("%v\n", WrapInt32(0, 100, 101))  //=> 1
	fmt.Printf("%v\n", WrapInt32(50, 100, 120)) //=> 70
	fmt.Printf("%v\n", WrapInt32(50, 100, 10))  //=> 60
	fmt.Printf("%v\n", WrapInt32(0, 100, -10))  //=> 90
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

func ExampleClampInt32() {
	fmt.Printf("%v\n", ClampInt32(0, 100, 120)) //=> 100
	fmt.Printf("%v\n", ClampInt32(0, 100, -20)) //=> 0
	fmt.Printf("%v\n", ClampInt32(100, 0, -20)) //=> 0 (even though min & max are swapped)
	// Output:
	// 100
	// 0
	// 0
}

func ExampleTestInt32() {
	fmt.Printf("%v\n", TestInt32(0, 100, 0, false, false))  //=> true
	fmt.Printf("%v\n", TestInt32(0, 100, 0, true, false))   //=> false
	fmt.Printf("%v\n", TestInt32(0, 100, 100, true, false)) //=> true
	fmt.Printf("%v\n", TestInt32(100, 0, 100, false, true)) //=> true
	// Output: true
	// false
	// true
	// true
}

func ExampleValidateInt32() {
	fmt.Println(ValidateInt32(0, 100, 0, false, false))   //=> 0 <nil>
	fmt.Println(ValidateInt32(0, 100, 0, true, false))    //=> 0 0 is outside of range [0, 100)
	fmt.Println(ValidateInt32(0, 100, 100, true, false))  //=> 100 <nil>
	fmt.Println(ValidateInt32(0, 100, 101, false, false)) //=> 0 101 is outside of range [0,100]
	fmt.Println(ValidateInt32(0, 100, 100, true, true))   //=> 0 100 is outside of range (0,100)
	// Output:
	// 0 <nil>
	// 0 0 is outside of range (0,100]
	// 100 <nil>
	// 0 101 is outside of range [0,100]
	// 0 100 is outside of range (0,100)
}

func ExampleToStringInt32() {
	fmt.Printf("%v\n", ToStringInt32(0, 100, true, true))   //=> (0,100)
	fmt.Printf("%v\n", ToStringInt32(0, 100, false, false)) //=> [0,100]
	fmt.Printf("%v\n", ToStringInt32(0, 100, true, false))  //=> (0,100]
	fmt.Printf("%v\n", ToStringInt32(0, 100, false, true))  //=> [0,100)
	// Output:
	// (0,100)
	// [0,100]
	// (0,100]
	// [0,100)
}

func ExampleRangeInt32() {
	r := NewRangeInt32(10, 100, false, false)
	fmt.Println(r.Wrap(120))                //=> 30
	fmt.Println(r.Validate(120))            //=> (0, error(120 is outside of range [0,100]))
	fmt.Println(r.Test(120))                //=> false
	fmt.Println(r)                          //=> [0,100] (uses Stringer interface)
	r = NewRangeInt32(100, 10, false, true) // swapped
	fmt.Println(r.Wrap(120))                //=> 30
	fmt.Println(r.Validate(120))            //=> (0, error(120 is outside of range (0,100]))
	fmt.Println(r.Test(120))                //=> false
	fmt.Println(r.Clamp(120))               //=> 100
	fmt.Println(r.Clamp(0))                 //=> 10
	fmt.Println(r)                          //=> (0,100] (uses Stringer interface)
	// Output:
	// 30
	// 0 120 is outside of range [10,100]
	// false
	// [10,100]
	// 30
	// 0 120 is outside of range (10,100]
	// false
	// 100
	// 10
	// (10,100]
}

func ExampleMinMaxInt32() {
	x, y := MinMaxInt32(0, 100)
	fmt.Printf("%v %v\n", x, y) //=> 0 100
	x, y = MinMaxInt32(100, 0)  // swapped
	fmt.Printf("%v %v\n", x, y) //=> 0 100
	// Output:
	// 0 100
	// 0 100
}

func ExampleMinMaxExclusiveInt32() {
	x, y, xe, ye := MinMaxExclusiveInt32(0, 100, true, false)
	fmt.Printf("%v %v %v %v\n", x, y, xe, ye)                //=> 0 100 true false
	x, y, xe, ye = MinMaxExclusiveInt32(100, 0, false, true) // swapped
	fmt.Printf("%v %v %v %v\n", x, y, xe, ye)                //=> 0 100 true false
	// Output:
	// 0 100 true false
	// 0 100 true false
}
