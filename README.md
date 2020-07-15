# [github.com/zzwx/interval](https://pkg.go.dev/github.com/zzwx/interval)

Utility for normalizing a numeric range, with a wrapping function useful for polar coordinates.

It is a **golang clone** of a JavaScript project by [James Talmage](https://github.com/jamestalmage/normalize-range).

For dealing with the strict typing in Go, functions were simply auto-generated functions for all of the following types:

* `int`
* `int64`
* `int32`
* `int16`
* `int8`
* `uint`
* `uint64`
* `uint32`
* `uint16`
* `uint8`
* `float32`
* `float64`

## Motivation

I did not investigate whether generating all the variations is the best way or not to deal with all the numeric types. Making functions accept `interface{}` didn't sound like a robust approach. Certainly if only a subset of these auto-generated functions is used, the Go compiler is not going to include them all in the final build. This approach is inspired by [Rob Pike's article](https://blog.golang.org/generate) on code generation. In a meanwhile I found out that Generics may not be an overall remedy, especially where the code should be written differently depending on the actual numeric capacity. I ended up writing a universal code though.

## Installation

```
go get -u github.com/zzwx/interval
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/zzwx/interval"
)

func main() {
	fmt.Println(interval.WrapInt(0, 360, 400))  //=> 40
	fmt.Println(interval.WrapInt(0, 360, -90))  //=> 270
	fmt.Println(interval.ClampInt(0, 100, 500)) //=> 100
	fmt.Println(interval.ClampInt(0, 100, -20)) //=> 0

	r := interval.NewRangeFloat64(0, 100, false, false)
	fmt.Println(r.Wrap(120))     //=> 20
	fmt.Println(r.Validate(120)) //=> 0, error(120 is outside of range [0,100])
	fmt.Println(r.Test(120))     //=> false
	fmt.Println(r)               //=> [0,100] (uses Stringer interface)
}
```

[Go Playground](https://play.golang.org/p/c_cqte_YoAe)

## API

https://pkg.go.dev/github.com/zzwx/interval

## License

Original JavaScript author: [James Talmage](https://github.com/jamestalmage/normalize-range)

MIT © [Anton Veretennikov](https://github.com/zzwx)

