# BitField

The `BitField` package provides the primitives to work with bitfields in Go. This package is very simple and
can be used to implement more complex data structures such as bitsets, bloom filters, etc.

The `BitField` is implemented using `uint64` type, so it can hold up to 64 bits for one bitfield. To work with
bitfields larger than 64 bits, you can combine multiple `BitField` as an array to build bigger bitfields.   

## Installation

```bash
go get github.com/josestg/bitfield
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/josestg/bitfield"
)

func main() {
	var f bitfield.BitField
	f = f.SetBit(0).SetBit(1).SetBit(3).SetBit(4)
	fmt.Printf("%08b\n", f)   // 00011011
	fmt.Println(f.IsSet(0))   // true
	fmt.Println(f.IsSet(2))   // false
	fmt.Println(f.IsSet(3))   // true
	fmt.Println(f.Cardinal()) // 4

	g := f.DelBit(3)
	fmt.Println(g.IsSet(3)) // false
}
```