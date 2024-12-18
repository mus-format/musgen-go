# musgen-go
musgen-go is a code generator for the [mus-go](https://github.com/mus-format/mus-go)
serializer. It can generate unsafe, streaming code, and currently supports 
only the MUS format.

# Contents
- [musgen-go](#musgen-go)
- [Contents](#contents)
  - [How to use](#how-to-use)
  - [Configuration](#configuration)
    - [Prefix](#prefix)
    - [Options](#options)
      - [Alias Options](#alias-options)
      - [Struct Options](#struct-options)
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
Here, we will generate `Marshal`, `Unmarshal`, `Size`, and `Skip` functions.

First, download and install Go (version 1.18 or later). Then, create a `foo` 
folder in your home directory with the following structure:
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
  s string
  b bool
  i IntAlias
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
`Marshal/Unmarshal/Size/Skip` functions for `IntAlias` and `Foo` types. 
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
      s: "hello world",
      b: true,
      i: IntAlias(5),
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

## Configuration
Configuration is done through the `FileGenerator.Add...With()` methods.

### Prefix
A prefix allows the generation of multiple `Marshal/Unmarshal/Size/Skip` 
functions for the same type. For example, `MarshalIntAliasMUS/...` and 
`MarshalAwesomeIntAliasMUS/...`, where `Awesome` is the prefix.
```go
// ...
prefix := "Awesome"
err = g.AddAliasWith(reflect.TypeFor[IntAlias](), prefix, nil)
if err != nil {
  panic(err)
}
// ...
```

### Options
Using the options, you can specify the encoding for the type, define a validator, 
ignore a structure field, etc.

#### Alias Options
Let's take an example:
```go
// ...
opts := basegen.NumOptions {
  Encoding: basegen.Raw, // IntAlias will be serialized with Raw encoding.
}
err = g.AddAliasWith(reflect.TypeFor[IntAlias](), "", opts)
if err != nil {
  panic(err)
}
// ...
```
Other available Options:
- `BoolOptions`
- `StringOptions`
- `SliceOptions`
- `ArrayOptions`
- `MapOptions`
- `PtrOptions`

The encoding of slice elements can be configured as follows:
```go
opts := basegen.SliceOptions {
  Elem: basegen.NumOpitions{ Encoding: basegen.Raw }, // Raw encoding will be
  // used for slice elements.
}
```
The same applies to arrays, maps, and pointers.

It should be noted that if an incorrect Options are used for a type, such as
`BoolOptions` for the `string` type, the worst that can happen is that some 
settings will not be applied.

#### Struct Options
The same options, but with the `FieldOptions` suffix, are available for struct 
fields: `BoolFieldOptions`, `NumFieldOptions`, etc. Additionally, 
`CustomTypeFieldOptions` can be used for fields with alias, struct, interface, 
or DTS types.
```go
// ...
opts := basegen.StructOptions{ // A slice whose elements must correspond to 
// struct fields.
  basegen.NumFieldOptions{ // For s field.
    NumOptions: basegen.NumOptions{
      Encoding: basegen.VarintPositive, // Sets a VarintPositive encoding fot this field.
    },
  },
  nil, // There is no Options for b field.
  basegen.CustomTypeFieldOptions{ // For i field.
    Prefix: "Awesome",
  },
}
err = g.AddStructWith(reflect.TypeFor[Foo](), "", opts)
// ...
```

##### Struct Prefix
Specifying a prefix for the entire structure means that it will be applied to 
all fields with alias, struct, interface, or DTS types.
```go
// ...
prefix := "Awesome"
err = g.AddStructWith(reflect.TypeFor[Foo](), prefix, opts)
// ...
```
In this case, for example, `MarshalAwesomeIntAlias()` will be used to marshal 
the `IntAlias` field. This common prefix can be ignored by the field:
```go
basegen.CustomTypeFieldOptions {
  Prefix: basegen.EmptyPrefix,
}
```
or can be overridden:
```go
basegen.CustomTypeFieldOptions {
  Prefix: "OwnPrefix",
}
```

##### Ignore a Field
The field also can be ignored:
```go
basegen.NumFieldOptions {
  Ignore: true,
}
```
All `FieldOptions` have an `Ignore` flag.

### Validation
Validation is performed during the unmarshalling process and requires one or 
more validators to be set. Each validator is just a function with the following 
signature `func (value Type) error`, where `Type` is a type of the value to 
which the validator is applied. To set a validator for an alias or struct 
field, use the `Validator` property. For example:
```go
func NotZero[T comparable](t T) (err error) { // Validator.
  if t == *new(T) {
    err = ErrZeroValue
  }
  return
}
// ...
opts := basegen.StructOptions{
  basegen.NumFieldOptions{
    NumOptions: basegen.NumOptions{
      Validator: "NotZero", // After unmarshalling the Foo.s field, its 
    // value will be checked by the NotZero validator. In general, we should 
    // write “packageName.ValidatorName” or just “ValidatorName” if the 
    // validator is from the same package.
    },
  },
  nil,
  nil,
}
err = g.AddStructWith(reflect.TypeFor[Foo], "", opts)
// ...
```
If the validator returns an error, it will be returned immediately by the
`Unmarshal` function, i.e. the rest of the struct will not be unmarshalled.

## Unsafe Code
To generate an unsafe code set the `Conf.Unsafe` flag:
```go
g, err := musgen.NewFileGenerator(basegen.Conf{
  Unsafe: true,
  // ...
})
```

## Streaming
musgen-go can also produce a streaming code:
```go
g, err := musgen.NewFileGenerator(basegen.Conf{
  Stream: true,
  // ...
})
```
In this case mus-stream-go library will be used instead of mus-go.

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
### Defaults
By default generated code uses:
- Varint encoding for numbers.
- VarintPositive encoding for length of variable length data types, such 
  as `string`, `slice` or `map`.
- VarintPositive encoding for DTM.
- No validation.

### Generate DTS
In addition to alias and struct a [DTS](https://github.com/mus-format/mus-dts-go) 
can be added to the `FileGenerator`. DTSs are useful when there is a need to 
deserialize data, but we don't know in advance what type it has. For example, it
could be `Foo` or `Bar`, we just don't know, but want to handle both of these 
cases.

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
way.

### Oneof Feature
Oneof feature is implemented using interfaces. Adding an interface to the 
`FileGenerator` requires `InterfaceOptions` with a non-empty `Oneof` property, 
which must contain one or more interface implementation types.
```go
// ...
opts := basegen.InterfaceOptions{
  Oneof: []reflect.Type{
    reflect.TypeFor[Copy](),
    reflect.TypeFor[Insert](),
  },
}
err = g.AddInterface(reflect.TypeFor[Instruction](), opts)
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