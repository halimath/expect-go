# expect-go

![CI Status][ci-img-url] 
[![Go Report Card][go-report-card-img-url]][go-report-card-url] 
[![Package Doc][package-doc-img-url]][package-doc-url] 
[![Releases][release-img-url]][release-url]


A library for writing test expectations using golang 1.18 generics to provide a fluent, readable and type-safe 
expectations.

# Installation

This module uses golang modules and can be installed with

```shell
go get github.com/halimath/expect-go@main
```

# Usage in tests

The following example demonstates the basic use:

```go
expect.That(t, got).
	Is(expect.DeepEqual(MyStruc{
    	Foo: "bar",
    	Spam: "eggs",
	}))
```

To start a new chain of expectations, use the `That` function providing a `testing.T` or `testing.B` and the
value to run expections on. Then, use one of the chaining methods `Is`, `Has`, `And` or `Matches` to add a 
matcher to the chain. `expect-go` provides a set of predefined matchers (see below) but you can also define
your own matchers.

If you want to stop the test's execution on the first failing expectation, provide the `StopImmediate` clause
to `That`: 

```go
expect.That(t, got, expect.StopImmediately{}).
	Is(expect.DeepEqual(MyStruc{
    	Foo: "bar",
    	Spam: "eggs",
	}))

```

## Predefined matchers

The following table shows the predefined matchers.

Matcher | Type constraints | Description
-- | -- | --
`Nil` | `any` | Expects a pointer to be `nil`
`NotNil` | `any` | Expects a pointer to be non `nil`
`Equal` | `comparable` | Compares given and wanted for equality using the go `==` operator.
`DeepEqual` | `any` | Compares given and wanted for deep equality using reflection.
`NoError` | `error` | Expects the given error value to be `nil`.
`Error` | `error` | Expects that the given error to be a non-`nil` error that is of the given target error by using `errors.Is` 
`HTTPStatus` | `httptest.ResponseRecorder` | Expects that the response recorded a given status code
`HTTPHeader` | `*http.Request | httptest.ResponseRecorder` | Expects that the HTTP entity conatins a given header with value

## Defining you own matcher

Defining you own matcher is very simple: Implement a type that implements the `Matcher` interface which
contains a single method: `Match`. The method receives a `Context` and the actual value. Perform the matching
steps and call `Fail` of `Failf` from the `Context` to fail the test with a given message. As most matchers
can be implemented by a closure function, `expect-go` provides the `MatcherFunc` convenience type.

The following example shows how to implement a matcher for asserting that a given number is even. The example
uses generics to handle all kinds of integral numbers.

```go
type Mod interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Even[M Mod]() expect.Matcher[M] {
	return expect.MatcherFunc[M](func(ctx expect.Context, got M) {
		if got%2 != 0 {
			ctx.Failf("expected %v to be even", got)
		}
	})
}

func TestCustomMatcher(t *testing.T) {
	var i int = 22
	expect.That(t, i).Is(Even[int]())
}
```

This example creates a type constraint interface assembling a union of all the number types a modulo operation
is useful for. It then defines a generic factory function `Even` to create a custom matcher for a given
integral type implemented as a closure using the `MatcherFunc` type. 

Note that we need to specify the generic type argument when using the matcher. This is due to the fact, that 
`Even` is not accepting any kind of argument. Hopefully, a later version of the go compiler will be able to 
interfer the type argument based on the context it is used in. 

We can rewrite this matcher to be a little bit more versatile, we get the following:

```go
func DivisableBy[M Mod](d M) expect.Matcher[M] {
	return expect.MatcherFunc[M](func(ctx expect.Context, got M) {
		if got%d != 0 {
			ctx.Failf("expected %d to be divisable by %d", got, d)
		}
	})
}

func TestCustomMatcher2(t *testing.T) {
	var i int = 22
	expect.That(t, i).Is(DivisableBy(2))
}
```

As you can see here, there is no need to specify any generic arguments.

# License

Copyright 2022 Alexander Metzner.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licensassertthat-go
assertthat-go
assertthat-go
assertthat-go
assertthat-go
assertthat-go
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[ci-img-url]: https://github.com/halimath/expect-go/workflows/CI/badge.svg
[go-report-card-img-url]: https://goreportcard.com/badge/github.com/halimath/expect-go
[go-report-card-url]: https://goreportcard.com/report/github.com/halimath/expect-go
[package-doc-img-url]: https://img.shields.io/badge/GoDoc-Reference-blue.svg
[package-doc-url]: https://pkg.go.dev/github.com/halimath/expect-go
[release-img-url]: https://img.shields.io/github/v/release/halimath/expect-go.svg
[release-url]: https://github.com/halimath/expect-go/releases