// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package scauri

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type Composition struct {
	Mass     uint64 
	Material string 
}

func NewCompositionFromBytes(buf []byte) *Composition {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Composition{}
	data.Mass     = wasmtypes.Uint64Decode(dec)
	data.Material = wasmtypes.StringDecode(dec)
	dec.Close()
	return data
}

func (o *Composition) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.Mass)
		wasmtypes.StringEncode(enc, o.Material)
	return enc.Buf()
}

type ImmutableComposition struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableComposition) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableComposition) Value() *Composition {
	return NewCompositionFromBytes(o.proxy.Get())
}

type MutableComposition struct {
	proxy wasmtypes.Proxy
}

func (o MutableComposition) Delete() {
	o.proxy.Delete()
}

func (o MutableComposition) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableComposition) SetValue(value *Composition) {
	o.proxy.Set(value.Bytes())
}

func (o MutableComposition) Value() *Composition {
	return NewCompositionFromBytes(o.proxy.Get())
}

type FracComposition struct {
	Mass     uint64  //in mg
	Material string 
}

func NewFracCompositionFromBytes(buf []byte) *FracComposition {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &FracComposition{}
	data.Mass     = wasmtypes.Uint64Decode(dec)
	data.Material = wasmtypes.StringDecode(dec)
	dec.Close()
	return data
}

func (o *FracComposition) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.Mass)
		wasmtypes.StringEncode(enc, o.Material)
	return enc.Buf()
}

type ImmutableFracComposition struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableFracComposition) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableFracComposition) Value() *FracComposition {
	return NewFracCompositionFromBytes(o.proxy.Get())
}

type MutableFracComposition struct {
	proxy wasmtypes.Proxy
}

func (o MutableFracComposition) Delete() {
	o.proxy.Delete()
}

func (o MutableFracComposition) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableFracComposition) SetValue(value *FracComposition) {
	o.proxy.Set(value.Bytes())
}

func (o MutableFracComposition) Value() *FracComposition {
	return NewFracCompositionFromBytes(o.proxy.Get())
}

type Fraction struct {
	Amount  uint64 
	Did     string 
	FracId  wasmtypes.ScHash 
	Issuer  wasmtypes.ScAgentID 
	Name    string 
	Pure    bool 
	Purpose string 
}

func NewFractionFromBytes(buf []byte) *Fraction {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Fraction{}
	data.Amount  = wasmtypes.Uint64Decode(dec)
	data.Did     = wasmtypes.StringDecode(dec)
	data.FracId  = wasmtypes.HashDecode(dec)
	data.Issuer  = wasmtypes.AgentIDDecode(dec)
	data.Name    = wasmtypes.StringDecode(dec)
	data.Pure    = wasmtypes.BoolDecode(dec)
	data.Purpose = wasmtypes.StringDecode(dec)
	dec.Close()
	return data
}

func (o *Fraction) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.Amount)
		wasmtypes.StringEncode(enc, o.Did)
		wasmtypes.HashEncode(enc, o.FracId)
		wasmtypes.AgentIDEncode(enc, o.Issuer)
		wasmtypes.StringEncode(enc, o.Name)
		wasmtypes.BoolEncode(enc, o.Pure)
		wasmtypes.StringEncode(enc, o.Purpose)
	return enc.Buf()
}

type ImmutableFraction struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableFraction) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableFraction) Value() *Fraction {
	return NewFractionFromBytes(o.proxy.Get())
}

type MutableFraction struct {
	proxy wasmtypes.Proxy
}

func (o MutableFraction) Delete() {
	o.proxy.Delete()
}

func (o MutableFraction) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableFraction) SetValue(value *Fraction) {
	o.proxy.Set(value.Bytes())
}

func (o MutableFraction) Value() *Fraction {
	return NewFractionFromBytes(o.proxy.Get())
}

type ProductPass struct {
	ActivationDate           uint64 
	AmountPerCharge          uint64 
	ChargeWeight             uint64 
	Did                      string  //merged did:iota:id
	ExpiryDate               uint64 
	Id                       wasmtypes.ScHash 
	Issuer                   wasmtypes.ScAgentID  //packaging producer
	Name                     string 
	PackageWeight            uint64 
	PackagesAlreadyPaid      uint64 
	PackagesNumber           uint64 
	PackagesSorted           uint64 
	PackagesWrongSorted      uint64 
	Purpose                  string  //e.g. food, hygiene, others
	RewardPerPackageProducer uint64 
	RewardPerPackageRecycler uint64 
	Version                  uint8 
}

func NewProductPassFromBytes(buf []byte) *ProductPass {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &ProductPass{}
	data.ActivationDate           = wasmtypes.Uint64Decode(dec)
	data.AmountPerCharge          = wasmtypes.Uint64Decode(dec)
	data.ChargeWeight             = wasmtypes.Uint64Decode(dec)
	data.Did                      = wasmtypes.StringDecode(dec)
	data.ExpiryDate               = wasmtypes.Uint64Decode(dec)
	data.Id                       = wasmtypes.HashDecode(dec)
	data.Issuer                   = wasmtypes.AgentIDDecode(dec)
	data.Name                     = wasmtypes.StringDecode(dec)
	data.PackageWeight            = wasmtypes.Uint64Decode(dec)
	data.PackagesAlreadyPaid      = wasmtypes.Uint64Decode(dec)
	data.PackagesNumber           = wasmtypes.Uint64Decode(dec)
	data.PackagesSorted           = wasmtypes.Uint64Decode(dec)
	data.PackagesWrongSorted      = wasmtypes.Uint64Decode(dec)
	data.Purpose                  = wasmtypes.StringDecode(dec)
	data.RewardPerPackageProducer = wasmtypes.Uint64Decode(dec)
	data.RewardPerPackageRecycler = wasmtypes.Uint64Decode(dec)
	data.Version                  = wasmtypes.Uint8Decode(dec)
	dec.Close()
	return data
}

func (o *ProductPass) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.ActivationDate)
		wasmtypes.Uint64Encode(enc, o.AmountPerCharge)
		wasmtypes.Uint64Encode(enc, o.ChargeWeight)
		wasmtypes.StringEncode(enc, o.Did)
		wasmtypes.Uint64Encode(enc, o.ExpiryDate)
		wasmtypes.HashEncode(enc, o.Id)
		wasmtypes.AgentIDEncode(enc, o.Issuer)
		wasmtypes.StringEncode(enc, o.Name)
		wasmtypes.Uint64Encode(enc, o.PackageWeight)
		wasmtypes.Uint64Encode(enc, o.PackagesAlreadyPaid)
		wasmtypes.Uint64Encode(enc, o.PackagesNumber)
		wasmtypes.Uint64Encode(enc, o.PackagesSorted)
		wasmtypes.Uint64Encode(enc, o.PackagesWrongSorted)
		wasmtypes.StringEncode(enc, o.Purpose)
		wasmtypes.Uint64Encode(enc, o.RewardPerPackageProducer)
		wasmtypes.Uint64Encode(enc, o.RewardPerPackageRecycler)
		wasmtypes.Uint8Encode(enc, o.Version)
	return enc.Buf()
}

type ImmutableProductPass struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableProductPass) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableProductPass) Value() *ProductPass {
	return NewProductPassFromBytes(o.proxy.Get())
}

type MutableProductPass struct {
	proxy wasmtypes.Proxy
}

func (o MutableProductPass) Delete() {
	o.proxy.Delete()
}

func (o MutableProductPass) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableProductPass) SetValue(value *ProductPass) {
	o.proxy.Set(value.Bytes())
}

func (o MutableProductPass) Value() *ProductPass {
	return NewProductPassFromBytes(o.proxy.Get())
}

type RecyComposition struct {
	Mass     uint64  //in mg
	Material string 
}

func NewRecyCompositionFromBytes(buf []byte) *RecyComposition {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &RecyComposition{}
	data.Mass     = wasmtypes.Uint64Decode(dec)
	data.Material = wasmtypes.StringDecode(dec)
	dec.Close()
	return data
}

func (o *RecyComposition) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.Uint64Encode(enc, o.Mass)
		wasmtypes.StringEncode(enc, o.Material)
	return enc.Buf()
}

type ImmutableRecyComposition struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableRecyComposition) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableRecyComposition) Value() *RecyComposition {
	return NewRecyCompositionFromBytes(o.proxy.Get())
}

type MutableRecyComposition struct {
	proxy wasmtypes.Proxy
}

func (o MutableRecyComposition) Delete() {
	o.proxy.Delete()
}

func (o MutableRecyComposition) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableRecyComposition) SetValue(value *RecyComposition) {
	o.proxy.Set(value.Bytes())
}

func (o MutableRecyComposition) Value() *RecyComposition {
	return NewRecyCompositionFromBytes(o.proxy.Get())
}

type Recyclate struct {
	Did     string 
	FracId  wasmtypes.ScHash 
	Issuer  wasmtypes.ScAgentID 
	Name    string 
	Pure    bool 
	Purpose string 
	RecyId  wasmtypes.ScHash 
}

func NewRecyclateFromBytes(buf []byte) *Recyclate {
	dec := wasmtypes.NewWasmDecoder(buf)
	data := &Recyclate{}
	data.Did     = wasmtypes.StringDecode(dec)
	data.FracId  = wasmtypes.HashDecode(dec)
	data.Issuer  = wasmtypes.AgentIDDecode(dec)
	data.Name    = wasmtypes.StringDecode(dec)
	data.Pure    = wasmtypes.BoolDecode(dec)
	data.Purpose = wasmtypes.StringDecode(dec)
	data.RecyId  = wasmtypes.HashDecode(dec)
	dec.Close()
	return data
}

func (o *Recyclate) Bytes() []byte {
	enc := wasmtypes.NewWasmEncoder()
		wasmtypes.StringEncode(enc, o.Did)
		wasmtypes.HashEncode(enc, o.FracId)
		wasmtypes.AgentIDEncode(enc, o.Issuer)
		wasmtypes.StringEncode(enc, o.Name)
		wasmtypes.BoolEncode(enc, o.Pure)
		wasmtypes.StringEncode(enc, o.Purpose)
		wasmtypes.HashEncode(enc, o.RecyId)
	return enc.Buf()
}

type ImmutableRecyclate struct {
	proxy wasmtypes.Proxy
}

func (o ImmutableRecyclate) Exists() bool {
	return o.proxy.Exists()
}

func (o ImmutableRecyclate) Value() *Recyclate {
	return NewRecyclateFromBytes(o.proxy.Get())
}

type MutableRecyclate struct {
	proxy wasmtypes.Proxy
}

func (o MutableRecyclate) Delete() {
	o.proxy.Delete()
}

func (o MutableRecyclate) Exists() bool {
	return o.proxy.Exists()
}

func (o MutableRecyclate) SetValue(value *Recyclate) {
	o.proxy.Set(value.Bytes())
}

func (o MutableRecyclate) Value() *Recyclate {
	return NewRecyclateFromBytes(o.proxy.Get())
}
