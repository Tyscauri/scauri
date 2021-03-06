// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ArrayOfImmutableString struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableString) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableString) GetString(index uint32) wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(a.proxy.Index(index))
}

type ImmutableStringArray = ArrayOfImmutableString

type ArrayOfMutableString struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableString) AppendString() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(a.proxy.Append())
}

func (a ArrayOfMutableString) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableString) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableString) GetString(index uint32) wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(a.proxy.Index(index))
}

type MutableStringArray = ArrayOfMutableString

type MapStringToImmutableString struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableString) GetString(key string) wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableStringMap = MapStringToImmutableString

type MapStringToMutableString struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableString) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableString) GetString(key string) wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableStringMap = MapStringToMutableString
