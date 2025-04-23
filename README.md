# musgen-go
musgen-go is a Golang code generator for the mus-go serializer.

## Capabilities
- Generates high-performance serialization code with optional unsafe 
  optimizations.
- Supports both in-memory and streaming data processing models.
- Can generate code for parameterized types and interfaces.
- Provides multi-package support.
- Enables cross-package code generation.
- Can be extended to support additional serialization formats beyond MUS.

# Contents
- [musgen-go](#musgen-go)
  - [Capabilities](#capabilities)
- [Contents](#contents)
- [Quick Example](#quick-example)
- [FileGenerator](#filegenerator)
  - [Configuration](#configuration)
    - [Required Options](#required-options)
    - [Streaming](#streaming)
    - [Unsafe Code](#unsafe-code)
    - [NotUnsafe Code](#notunsafe-code)
    - [Imports](#imports)
    - [Serializer Name](#serializer-name)
  - [Methods](#methods)
    - [AddDefinedType()](#adddefinedtype)
    - [AddStruct()](#addstruct)
    - [AddDTS()](#adddts)
    - [AddInterface()](#addinterface)
- [Multi-package support](#multi-package-support)
- [Cross-Package Code Generation](#cross-package-code-generation)
- [Serialization Options](#serialization-options)
  - [Numbers](#numbers)
  - [String](#string)
  - [Array](#array)
  - [Slice](#slice)
  - [Map](#map)
  - [time.Time](#timetime)
- [MUS Format](#mus-format)


# Quick Example
Here, we will generate a MUS serializer for the `Foo` type.

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

type Foo[T any] struct {
  s string
  t T
}
```

__gen/main.go__
```go
package main

import (
  "os"
  "reflect"

  "example.com/foo"

  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

func main() {
  g, err := musgen.NewFileGenerator(
    genops.WithPkgPath("example.com/foo"),
    // genops.WithPackage("bar"), // Can be used to specify the package name for
    // the generated file.
  )
  if err != nil {
    panic(err)
  }
  err = g.AddDefinedType(reflect.TypeFor[foo.MyInt]())
  if err != nil {
    panic(err)
  }
  err = g.AddStruct(reflect.TypeFor[foo.Foo[foo.MyInt]]())
  if err != nil {
    panic(err)
  }
  bs, err := g.Generate()
  if err != nil {
    // The error can be inspected for additional details.
    if engnErr, ok := err.(*musgen.TmplEngineError); ok {
      fmt.Println(string(engnErr.ByteSlice())) // Generated code that caused the error.
      fmt.Println(engnErr.Unwrap()) // Underlying error.
    }
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
$ go mod init example.com/foo
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
    foo = Foo[MyInt]{
      s: "hello world",
      t: MyInt(5),
    }
    size = FooMUS.Size(foo)
    bs   = make([]byte, size)
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

# FileGenerator
The `FileGenerator` is responsible for generating serialization code.

## Configuration
### Required Options
There are only one required configuration option:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g, err := musgen.NewFileGenerator(
  genops.WithPkgPath("pkgPath"),  // Sets the package path for the generated 
  // file. The path must match the standard Go package path format (e.g., 
  // github.com/user/project/pkg) and can be obtained using:
  //
  //   pkgPath := reflect.TypeFor[YourType]().PkgPath()
  //
)
```

### Streaming
To generate a streaming code:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(
  // ...
  genops.WithStream(),
)
```

In this case [mus-stream-go](https://github.com/mus-format/mus-stream-go) 
library will be used instead of mus-go.

### Unsafe Code
To generate an unsafe code:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(
  // ...
  genops.WithUnsafe(),
)
```

### NotUnsafe Code
In this mode, the unsafe package will be used for all types except string:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(
  // ...  
  genops.WithNotUnsafe(),
)
```

It produces the fastest serialization code without unsafe side effects.

### Imports
In some cases import statement of the generated file can miss one or more
packages. To fix this:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(
  // ...
  genops.WithImport("pkgPath"),
  genops.WithImportAlias("pkgPath", "alias"),
)
```

Also, `genops.WithImportAlias` helps prevent name conflicts when multiple 
packages are imported with the same alias.

### Serializer Name
Generated serializers follow the standard naming convention:
```
pkg.YouType[T,V] -> YouTypeMUS  // Serialization format is appended to the type name.
```

To override this behavior, use `genops.WithSerName()`:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(
  // ...
  genops.WithSerName(reflect.TypeFor[pkg.YourType](), "CustomSerName"),
)
```

## Methods
### AddDefinedType()
Supports types defined with the following source types:
- Number (`uint`, `int`, `float64`, ...)
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
  // Validator functions in general should have the following signature:
  //
  //   func(value Type) error
  //
  // where Type denotes the type the validator is applied to.
```

### AddStruct()
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
// want to specify options for some field, use structops.WithField() without
// any parameters.
err := g.AddStruct(reflect.TypeFor[MyStruct](),
  structops.WithField(), // No options for the first field.

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

### AddDTS()
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
err = g.AddDTS(t)  // The DTS definition will be generated for the specified 
// type.
```

### AddInterface()
Supports types defined with the `interface` source type, such as:
```go
type MyInterface interface { ... }
type MyAnotherInterface MyInterface
```

It can be used as follows:
```go
import ext "github.com/mus-format/ext-mus-go"

type MyInterface interface {...}
type Impl1 struct {...}
type Impl2 int

const (
  Impl1DTM com.DTM = iota + 1
  Impl2DTM
)

// ...

var (
  t1 = reflect.TypeFor[Impl1]()
  t2 = reflect.TypeFor[Impl2]()
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
  introps.WithImplType(t2),
  // introps.WithMarshaller(), // Enables serialization using the 
  // ext.MarshallerTypedMUS interface, that should be satisfied by all 
  // implementation types. Disabled by default.
)
```

# Multi-package support
By default, musgen-go expects a type’s serializer to reside in the same package 
as the type itself. For example, generating a serializer for the `Foo` type:
```go
package foo

type Foo struct{
  Bar bar.Bar
}
```

will result in:
```go
package foo

// ...

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
  return bar.BarMUS(v.Bar) // musgen-go assumes the Bar serializer is located
  // in the bar package and follows the default naming convention. However, this 
  // is not always the case.
}

// ...
```
To reference a `Bar` serializer defined in a different package or with a 
non-standard name, use the `genops.WithSerName` option:
```go
import (
  musgen "github.com/mus-format/musgen-go/mus"
  genops "github.com/mus-format/musgen-go/options/generate"
)

g := musgen.NewFileGenerator(
  // ...
  genops.WithSerName(reflect.TypeFor[bar.Bar](), "another.AwesomeBar"),
  // If only the name is non-standard:
  // genops.WithSerName(reflect.TypeFor[bar.Bar](), "AwesomeBar"),
)
```

The string `another.AwesomeBar` will be used as-is, with the serialization 
format appended:
```go
// ...

func (s fooMUS) Marshal(v Foo, bs []byte) (n int) {
  return another.AwesomeBarMUS(v.Bar)
}

// ...
```

# Cross-Package Code Generation
musgen-go allows to generate a serializer for a type from an external package, 
for example, `foo.BarMUS` for the `bar.Bar` type.

# Serialization Options
Different types support different serialization options. If an unsupported 
option is specified for a type, it will simply be ignored.

## Numbers
- `typeops.WithNumEncoding`
- `typeops.WithValidator`

## String
- `typeops.WithLenEncoding`
- `typeops.WithLenValidator`
- `typeops.WithValidator`

## Array
- `typeops.WithLenEncoding`
- `typeops.WithElem`
- `typeops.WithValidator`

## Slice
- `typeops.WithLenEncoding`
- `typeops.WithLenValidator`
- `typeops.WithElem`
- `typeops.WithValidator`

## Map
- `typeops.WithLenEncoding`
- `typeops.WithLenValidator`
- `typeops.WithKey`
- `typeops.WithElem`
- `typeops.WithValidator`

## time.Time
- `typeops.WithTimeUnit`
- `typeops.WithValidator`

# MUS Format
Defauls:
- Varint encoding is used for numbers.
- Varint without ZigZag encoding is used for the length of variable-length data 
  types, such as `string`, `array`, `slice`, or `map`.
- Varint without ZigZag encoding is used for DTM (Data Type Metadata).
- `raw.TimeUnix` is used for `time.Time`.
