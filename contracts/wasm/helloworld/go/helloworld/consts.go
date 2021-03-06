// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package helloworld

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "helloworld"
	ScDescription = "The ubiquitous hello world demo"
	HScName       = wasmtypes.ScHname(0x0683223c)
)

const (
	ResultHelloWorld = "helloWorld"
)

const (
	FuncHelloWorld    = "helloWorld"
	ViewGetHelloWorld = "getHelloWorld"
)

const (
	HFuncHelloWorld    = wasmtypes.ScHname(0x9d042e65)
	HViewGetHelloWorld = wasmtypes.ScHname(0x210439ce)
)
