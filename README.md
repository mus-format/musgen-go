# musgen-go
musgen-go is a Golang code generator that can produce code for various binary
serialization formats. It uses serialization primitives from the 
[mus-go](https://github.com/mus-format/mus-go) library, and for now supports 
only the [MUS](https://medium.com/p/21d7be309e8d) format.

For all supported formats, musgen-go can generate unsafe and streaming code. In 
addition, it has flexible customization options.

# Contents
- [musgen-go](#musgen-go)
- [Contents](#contents)
  - [How to use](#how-to-use)
  - [Custom Serialization](#custom-serialization)
    - [Prefix](#prefix)
    - [Metadata](#metadata)
      - [Alias Metadata](#alias-metadata)
      - [Struct Metadata](#struct-metadata)
        - [Struct Prefix](#struct-prefix)
        - [Ignore a Field](#ignore-a-field)
    - [Validation](#validation)
  - [Unsafe Code](#unsafe-code)
  - [Streaming](#streaming)
  - [Imports](#imports)
  - [MUS Format](#mus-format)
    - [Defaults](#defaults)
    - [Generate DTS](#generate-dts)
    - [Oneof Feature](#oneof-feature)

## How to use
Here we will generate the `Marshal/Unmarshal/Size/Skip` functions of the MUS 
format.

First, you should download and install Go, version 1.17 or later.

Create in your home directory a `foo` folder with the following structure:
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

type IntAlias int

type Foo struct {
  fld0 string
  fld1 bool
  fld2 IntAlias
}
```

__gen/main.go__
```go
package main

import (
  "os"
  "reflect"

  "foo"

  "github.com/mus-format/musgen-go/basegen"
  musgen "github.com/mus-format/musgen-go/mus"
)

func main() {
  g, err := musgen.NewFileGenerator(basegen.Conf{Package: "foo"})
  if err != nil {
    panic(err)
  }
  err = g.AddAlias(reflect.TypeFor[foo.IntAlias]())
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

Now you can see `mus-format.gen.go` file in the `foo` folder with the
`Marshal/Unmarshal/Size/Skip` MUS functions for `IntAlias` and `Foo` types. 
Let's write some tests. Create a `foo_test.go` file:
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
      fld0: "hello world",
      fld1: true,
      fld2: IntAlias(5),
    }
    bs = make([]byte, SizeFooMUS(foo))
  )
  MarshalFooMUS(foo, bs)
  afoo, _, err := UnmarshalFooMUS(bs)
  if err != nil {
    t.Fatal(err)
  }
  if !reflect.DeepEqual(foo, afoo) {
    t.Fatal("something went wrong")
  }
}
```

## Custom Serialization
musgen-go provides flexible options for customizing the serialization process.
It is done by `FileGenerator.Add...With()` methods.

### Prefix
Prefix allows to have several `Marshal/Unmarshal/Size/Skip` functions for one 
type. For example, at the same time we can have both `MarshalIntAliasMUS()` and 
`MarshalAwesomeIntAliasMUS()`, where `Awesome` is the prefix.
```go
// ...
prefix := "Awesome"
err = g.AddAliasWith(reflect.TypeFor[IntAlias](), prefix, nil)
if err != nil {
  panic(err)
}
// ...
```

### Metadata
Metadata also allows to customize the serialization of individual data types.

#### Alias Metadata
Let's look at an example:
```go
// ...
meta := basegen.NumMetadata{
  Encoding: basegen.Raw, // The IntAlias will be serialized using Raw encoding.
}
err = g.AddAliasWith(reflect.TypeFor[IntAlias](), "", meta)
if err != nil {
  panic(err)
}
// ...
```
There are other metadata types, such as:
- `BoolMetadata`
- `StringMetadata` - Allows to set length encoding and length validator to
  protect against too long strings.
- `SliceMetadata` - In addition to length encoding and length validator, allows
  to customize the serialization of elements.
- `MapMetadata` - In addition to length encoding and length validator, allows to
  customize the serialization of keys and values.
- `PtrMetadata`
  
, each has own customization options. It should be noted that if an incorrect 
metadata is set for a type (for example, `BoolMetadata` for a `string` type), 
the worst that can happen is that some settings will not be applied.

#### Struct Metadata
For struct fields there are `BoolFieldMetadata`, `NumFieldMetadata`, 
`CustomTypeFieldMetadata` (can be used for alias, struct, interface or DTS 
types), etc., all ending in `FieldMetadata`. Let's look at an example:
```go
// ...
meta := basegen.StructMetadata{ // basegen.StructMetadata is a slice whose 
// elements must correspond to struct fields.
  basegen.NumFieldMetadata{ // Corresponds to Foo.fld0.
    NumMetadata: basegen.NumMetadata{
      Encoding: basegen.VarintPositive, // Sets a VarintPositive encoding fot this field.
    },
  },
  nil, // Corresponds to Foo.fld1. There is no metadata for this field.
  basegen.CustomTypeFieldMetadata{ // Corresponds to Foo.fld2.
    Prefix: "Awesome",
  },
}
err = g.AddStructWith(reflect.TypeFor[Foo](), "", meta)
// ...
```

##### Struct Prefix
Specifying a prefix for the entire struct means that it will be applied to 
all fields with custom types (such as alias, struct, interface, or DTS). 
```go
// ...
prefix := "Awesome"
err = g.AddStructWith(reflect.TypeFor[Struct](), prefix, meta)
// ...
```
In this case, for example, `MarshalAwesomeIntAlias()` will be used to marshal
`fld2` field. This common prefix can be ignored by the field:
```go
basegen.CustomTypeFieldMetadata {
  Prefix: basegen.EmptyPrefix,
}
```
or can be overridden:
```go
basegen.CustomTypeFieldMetadata {
  Prefix: "OwnPrefix",
}
```

##### Ignore a Field
The field also can be ignored:
```go
basegen.NumFieldMetadata {
  Ignore: true,
}
```
All `FieldMetadata` types have an `Ignore` flag.

### Validation
Validation is performed during the unmarshalling process and requires one or 
more validators to be set. Each validator is just a function with the following 
signature `func (value Type) error`, where `Type` is a type of the value to 
which the validator is applied. To set a validator for an alias or struct 
field, use the `Validator` metadata property. For example:
```go
func NotZero[T comparable](t T) (err error) { // Validator.
  if t == *new(T) {
    err = ErrZeroValue
  }
  return
}
// ...
meta := basegen.StructMetadata{
  basegen.NumFieldMetadata{
    NumMetadata: basegen.NumMetadata{
      Validator: "NotZero", // After unmarshalling the Foo.fld0 field, its 
    // value will be checked by the NotZero validator. In general, we should 
    // write “packageName.ValidatorName” or just “ValidatorName” if the 
    // validator is from the same package.
    },
  },
  nil,
  nil,
}
err = g.AddStructWith(reflect.TypeFor[Foo], "", meta)
// ...
```
If the validator returns an error, it will be returned immediately by the
`Unmarshal` function, i.e. the rest of the struct will not be unmarshalled.

## Unsafe Code
To generate an unsafe code just set the `Conf.Unsafe` flag:
```go
g, err := musgen.NewFileGenerator(basegen.Conf{
  Unsafe: true,
  // ...
})
```

## Streaming
mesgen-go can also produce a streaming code:
```go
g, err := musgen.NewFileGenerator(basegen.Conf{
  Stream: true,
  // ...
})
```
In this case mus-stream-go library will be used.

## Imports
In some cases import statement of the generated file can miss one or more
packages. To fix this use `Conf.Imports`:
```go
g, err := musgen.NewFileGenerator(basegen.Conf{
  Imports: []string{
    "first import path",
    "second import path",
  },
  // ...
})
```

## MUS Format
To generate the MUS format code, use the `github.com/mus-format/musgen-go/mus` 
package.

### Defaults
By default generated code:
- Uses Varint encoding for numbers.
- The length of variable length data types (such as `string`, `slice` or `map`) 
  is encoded using Varint Postitive.
- DTMs also encoded using Varint Positive.
- There is no validation.

### Generate DTS
In addition to alias and struct, we can add DTS to the `FileGenerator`. DTSs are
useful when we need to deserialize data, but we don't know in advance what type 
it has. For example, it could be `Foo` or `Bar`, we just don't know..., but want
to handle both of these cases.

To add DTS generation, we need to define a DTM:
```go
const (
  IntAliasDTM = 1
)
```
and 
```go
// ...
err = g.AddAliasDTS(reflect.TypeFor[IntAlias]()) // Marshal/Unmarshal/Size/Skip
// functions and IntAliasDTS will be generated for the IntAlias type.
// ...
```
There is also `FileGenerator.AddStructDTS()` method that behaves in a similar 
way. More information about DTS can be found [here](https://github.com/mus-format/mus-dts-go).

### Oneof Feature
Oneof feature is implemented using interfaces. Adding an interface to the 
`FileGenerator` requires `InterfaceMetadata` with a non-empty `OneOf` property, 
which must contain one or more interface implementation types.
```go
// ...
meta := basegen.InterfaceMetadata{
  OneOf: []reflect.Type{
    reflect.TypeFor[Copy](),
    reflect.TypeFor[Insert](),
  },
}
err = g.AddInterface(reflect.TypeFor[Instruction](), meta)
// ...
```
, where `Instruction` is an interface implemented by `Copy` and `Insert`. Also, 
the latter must have DTMs:
```go
const (
  CopyDTM = 1
  InsertDTM  = 2
)
```
and DTSs:
```go
err = g.AddStructDTS(reflect.TypeFor[Copy]())
// ...
err = g.AddStructDTS(reflect.TypeFor[Insert]())
// ...
```