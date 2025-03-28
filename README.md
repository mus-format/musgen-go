# musgen-go
musgen-go is a code generator for the [mus-go](https://github.com/mus-format/mus-go) 
serializer. It can generate both unsafe and streaming code and currently 
supports only the MUS format.

## Quick Example
Here, we will generate a MUS serializer for the Foo type.

First, download and install Go (version 1.18 or later). Then, create a `foo` 
folder with the following structure:
```
foo/
 |‒‒‒gen/
 |    |‒‒‒main.go
 |‒‒‒foo.go
```

__foo.go__
```go
//go:generate go run gen/main.go
package foo

type MyInt int

type Foo struct {
  s string
  b bool
  i MyInt
}
```

__gen/main.go__
```go
package main

import (
  "os"
  "reflect"

  "foo"

  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

func main() {
  g := musgen.NewFileGenerator(genops.WithPackage("foo"))
  err := g.AddDefinedType(reflect.TypeFor[foo.MyInt]())
  if err != nil {
    panic(err)
  }
  err = g.AddStruct(reflect.TypeFor[foo.Foo]())
  if err != nil {
    panic(err)
  }
  bs, err := g.Generate()
  if err != nil {
    panic(err)
  }
  err = os.WriteFile("./mus-format.gen.go", bs, 0755)
  if err != nil {
    panic(err)
  }
}
```

Run from the command line:
```bash
$ cd ~/foo
$ go mod init foo
$ go mod tidy
$ go generate
$ go mod tidy
```

Now you can see `mus-format.gen.go` file in the `foo` folder with `MyIntMUS`
and `FooMUS` serializers. Let's write some tests. Create a `foo_test.go` file:
```
foo/
 |‒‒‒...
 |‒‒‒foo_test.go
```

__foo_test.go__
```go
package foo

import (
  "reflect"
  "testing"
)

func TestFooSerialization(t *testing.T) {
  var (
    foo = Foo{
      s: "hello world",
      b: true,
      i: MyInt(5),
    }
    size = FooMUS.Size(foo)
    bs = make([]byte, size)
  )
  FooMUS.Marshal(foo, bs)
  afoo, _, err := FooMUS.Unmarshal(bs)
  if err != nil {
    t.Fatal(err)
  }
  if !reflect.DeepEqual(foo, afoo) {
    t.Fatal("something went wrong")
  }
}
```

## FileGenerator
The `FileGenerator` is responsible for generating serialization code.

### Configuration
#### Streaming
To generate a streaming code:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(genops.WithPackage("package_name"),
  genops.WithStream())
```

In this case mus-stream-go library will be used instead of mus-go.

#### Unsafe Code
To generate an unsafe code:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(genops.WithPackage("package_name"),
  genops.WithUnsafe())
```

#### NotUnsafe Code
In this mode, the unsafe package will be used for all types except string:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(genops.WithPackage("package_name"),
  genops.WithNotUnsafe())
```

It produces the fastest serialization code without unsafe side effects.

#### Imports
In some cases import statement of the generated file can miss one or more
packages. To fix this:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(genops.WithPackage("package_name"),
  genops.WithImports([]string{"first import path", "second import path"}))
```

### Methods
#### AddDefinedType()
Supports types defined with the following source types:
- Number (`uint`, `uint64`, `uint32`, `uint16`, `uint8`, `int`, `int64`, 
  `int32`, `int16`, `int8`, `float64`, `float32`)
- String
- Array
- Slice
- Map
- Pointer

For example:
```go
type MyInt int
type MyStringSlice []string
type MyUintPtr *uint
// ...
```

It can be used as follows:
```go
import (
  "reflect"

  typeops "github.com/mus-format/musgen-go/options/type"
)

type MyInt int // Where int is the source type of MyInt.

// ...

err := g.AddDefinedType(reflect.TypeFor[MyInt]())
```

Or with serialization options, for example:
```go
err := g.AddDefinedType(reflect.TypeFor[MyInt](),
  typeops.WithNumEncoding(typeops.Raw), // The raw.Int serializer will be used
  // to serialize the source int type.
  typeops.WithValidator("ValidateMyInt")) // After unmarshalling, the MyInt
  // value will be validated using the ValidateMyInt function.
```

#### AddStruct()
Supports types defined with the `struct` source type, such as:
```go
type MyStruct struct { ... }
type MyAnotherStruct MyStruct
```

It can be used as follows:
```go
import (
  "reflect"

  genops "github.com/mus-format/musgen-go/options/generate"
  structops "github.com/mus-format/musgen-go/options/struct"
  typeops "github.com/mus-format/musgen-go/options/type"
)

type MyStruct struct {
  Str string
  Ignore int
  Slice []int
  // Interface MyInterface  // Interface fields are supported as well.
  // Any any                // But not the `any` type.
}

// ...

err := g.AddStruct(reflect.TypeFor[MyStruct]())
```

Or with serialization options, for example:
```go
// The number of options should be equal to the number of fields. If you don't
// want to specify options for some field, use structops.WithNil().
err := g.AddStruct(reflect.TypeFor[MyStruct](),
  structops.WithNil(), // No options for the first field.

  structops.WithField(typeops.WithIgnore()), // The second field will not be
  // serialized.

  structops.WithField( // Options for the third field.
    typeops.WithLenValidator("ValidateLength"), // The length of the slice
    // field will be validated using the ValidateLength function before the
    // rest of the slice is unmarshalled.
    typeops.WithElem( // Options for slice elements.
      typeops.WithNumEncoding(typeops.Raw), // The raw.Int serializer will be
      // used to serialize slice elements.
      typeops.WithValidator("ValidateSliceElem"), // Each slice element, after
      // unmarshalling, will be validated using the ValidateSliceElem function.
    ),
  ),
)
```

A special case for the `time.Time` source type:
```go
type MyTime time.Time
// ...
err = g.AddStruct(reflect.TypeFor[MyTime](),
  structops.WithSourceType(structops.Time, typeops.WithTimeUnit(typeops.Milli)), 
  // The raw.TimeUnixMilli serializer will be used  to serialize a time.Time 
  // value.
)
```

#### AddDTS()
Supports all types acceptable by the `AddDefinedType`, `AddStruct`, and 
`AddInterface` methods.

It can be used as follows:
```go
import (
  "reflect"
)

type MyInt int

// ...

t := reflect.TypeFor[MyInt]()
err := g.AddDefinedType(t)
// ...
err = g.AddDTS(t)  // Generator will generate a DTS definition for the specified 
// type.
```

#### AddInterface()
Supports types defined with the `interface` source type, such as:
```go
type MyInterface interface { ... }
type MyAnotherInterface MyInterface
```

It can be used as follows:
```go
type MyInterface interface {...}
type MyInterfaceImpl1 struct {...}
type MyInterfaceImpl2 int

// ...

var (
  t1 = reflect.TypeFor[MyInterfaceImpl1]()
  t2 = reflect.TypeFor[MyInterfaceImpl2]()
)

// ...

err := g.AddStruct(t1)
// ...
err = g.AddDTS(t1)
// ...
err = g.AddDefinedType(t2)
// ...
err = g.AddDTS(t2)
// ...
err = g.AddInterface(reflect.TypeFor[MyInterface](),
  introps.WithImplType(t1),
  introps.WithImplType(t2))
```

## Serialization Options
Different types support different serialization options. If an incorrect option 
is specified for a type, the worst that can happen is that it will be ignored.

### Numbers
- `typeops.WithNumEncoding`
- `typeops.WithValidator`

### String
- `typeops.WithLenEncoding`
- `typeops.WithLenValidator`
- `typeops.WithValidator`

### Array
- `typeops.WithLenEncoding`
- `typeops.WithElem`
- `typeops.WithValidator`

### Slice
- `typeops.WithLenEncoding`
- `typeops.WithLenValidator`
- `typeops.WithElem`
- `typeops.WithValidator`

### Map
- `typeops.WithLenEncoding`
- `typeops.WithLenValidator`
- `typeops.WithKey`
- `typeops.WithElem`
- `typeops.WithValidator`

### time.Time
- `typeops.WithTimeUnit`

## MUS Format
Defauls:
- Varint encoding is used for numbers.
- Varint without ZigZag encoding is used for the length of variable-length data 
  types, such as `string`, `array`, `slice`, or `map`.
- Varint without ZigZag encoding is used for DTM (Data Type Metadata).
- `raw.TimeUnix` is used for `time.Time`.
