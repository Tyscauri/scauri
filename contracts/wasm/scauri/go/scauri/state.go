// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package scauri

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type MapHashToImmutableCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableCompositions) GetCompositions(key wasmtypes.ScHash) ImmutableCompositions {
	return ImmutableCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableFracCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableFracCompositions) GetFracCompositions(key wasmtypes.ScHash) ImmutableFracCompositions {
	return ImmutableFracCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableFraction struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableFraction) GetFraction(key wasmtypes.ScHash) ImmutableFraction {
	return ImmutableFraction{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapAgentIDToImmutableUint64 struct {
	proxy wasmtypes.Proxy
}

func (m MapAgentIDToImmutableUint64) GetUint64(key wasmtypes.ScAgentID) wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(m.proxy.Key(wasmtypes.AgentIDToBytes(key)))
}

type MapHashToImmutableProductPass struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableProductPass) GetProductPass(key wasmtypes.ScHash) ImmutableProductPass {
	return ImmutableProductPass{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableRecyCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableRecyCompositions) GetRecyCompositions(key wasmtypes.ScHash) ImmutableRecyCompositions {
	return ImmutableRecyCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToImmutableRecyclate struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToImmutableRecyclate) GetRecyclate(key wasmtypes.ScHash) ImmutableRecyclate {
	return ImmutableRecyclate{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type ArrayOfImmutableAgentID struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableAgentID) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableAgentID) GetAgentID(index uint32) wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(a.proxy.Index(index))
}

type ImmutablescauriState struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablescauriState) Compositions() MapHashToImmutableCompositions {
	return MapHashToImmutableCompositions{proxy: s.proxy.Root(StateCompositions)}
}

func (s ImmutablescauriState) DonationAddress() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(StateDonationAddress))
}

func (s ImmutablescauriState) FracCompositions() MapHashToImmutableFracCompositions {
	return MapHashToImmutableFracCompositions{proxy: s.proxy.Root(StateFracCompositions)}
}

func (s ImmutablescauriState) Fractions() MapHashToImmutableFraction {
	return MapHashToImmutableFraction{proxy: s.proxy.Root(StateFractions)}
}

func (s ImmutablescauriState) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(StateOwner))
}

func (s ImmutablescauriState) PricePerMg() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StatePricePerMg))
}

func (s ImmutablescauriState) ProducersBalances() MapAgentIDToImmutableUint64 {
	return MapAgentIDToImmutableUint64{proxy: s.proxy.Root(StateProducersBalances)}
}

func (s ImmutablescauriState) Productpasses() MapHashToImmutableProductPass {
	return MapHashToImmutableProductPass{proxy: s.proxy.Root(StateProductpasses)}
}

func (s ImmutablescauriState) RecyCompositions() MapHashToImmutableRecyCompositions {
	return MapHashToImmutableRecyCompositions{proxy: s.proxy.Root(StateRecyCompositions)}
}

func (s ImmutablescauriState) Recyclates() MapHashToImmutableRecyclate {
	return MapHashToImmutableRecyclate{proxy: s.proxy.Root(StateRecyclates)}
}

func (s ImmutablescauriState) Recyclers() ArrayOfImmutableAgentID {
	return ArrayOfImmutableAgentID{proxy: s.proxy.Root(StateRecyclers)}
}

func (s ImmutablescauriState) RecyclersBalances() MapAgentIDToImmutableUint64 {
	return MapAgentIDToImmutableUint64{proxy: s.proxy.Root(StateRecyclersBalances)}
}

func (s ImmutablescauriState) ShareRecycler() wasmtypes.ScImmutableUint8 {
	return wasmtypes.NewScImmutableUint8(s.proxy.Root(StateShareRecycler))
}

func (s ImmutablescauriState) Sorters() ArrayOfImmutableAgentID {
	return ArrayOfImmutableAgentID{proxy: s.proxy.Root(StateSorters)}
}

func (s ImmutablescauriState) TokenToDonate() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StateTokenToDonate))
}

type MapHashToMutableCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableCompositions) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableCompositions) GetCompositions(key wasmtypes.ScHash) MutableCompositions {
	return MutableCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableFracCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableFracCompositions) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableFracCompositions) GetFracCompositions(key wasmtypes.ScHash) MutableFracCompositions {
	return MutableFracCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableFraction struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableFraction) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableFraction) GetFraction(key wasmtypes.ScHash) MutableFraction {
	return MutableFraction{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapAgentIDToMutableUint64 struct {
	proxy wasmtypes.Proxy
}

func (m MapAgentIDToMutableUint64) Clear() {
	m.proxy.ClearMap()
}

func (m MapAgentIDToMutableUint64) GetUint64(key wasmtypes.ScAgentID) wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(m.proxy.Key(wasmtypes.AgentIDToBytes(key)))
}

type MapHashToMutableProductPass struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableProductPass) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableProductPass) GetProductPass(key wasmtypes.ScHash) MutableProductPass {
	return MutableProductPass{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableRecyCompositions struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableRecyCompositions) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableRecyCompositions) GetRecyCompositions(key wasmtypes.ScHash) MutableRecyCompositions {
	return MutableRecyCompositions{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type MapHashToMutableRecyclate struct {
	proxy wasmtypes.Proxy
}

func (m MapHashToMutableRecyclate) Clear() {
	m.proxy.ClearMap()
}

func (m MapHashToMutableRecyclate) GetRecyclate(key wasmtypes.ScHash) MutableRecyclate {
	return MutableRecyclate{proxy: m.proxy.Key(wasmtypes.HashToBytes(key))}
}

type ArrayOfMutableAgentID struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableAgentID) AppendAgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(a.proxy.Append())
}

func (a ArrayOfMutableAgentID) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableAgentID) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableAgentID) GetAgentID(index uint32) wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(a.proxy.Index(index))
}

type MutablescauriState struct {
	proxy wasmtypes.Proxy
}

func (s MutablescauriState) AsImmutable() ImmutablescauriState {
	return ImmutablescauriState(s)
}

func (s MutablescauriState) Compositions() MapHashToMutableCompositions {
	return MapHashToMutableCompositions{proxy: s.proxy.Root(StateCompositions)}
}

func (s MutablescauriState) DonationAddress() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(StateDonationAddress))
}

func (s MutablescauriState) FracCompositions() MapHashToMutableFracCompositions {
	return MapHashToMutableFracCompositions{proxy: s.proxy.Root(StateFracCompositions)}
}

func (s MutablescauriState) Fractions() MapHashToMutableFraction {
	return MapHashToMutableFraction{proxy: s.proxy.Root(StateFractions)}
}

func (s MutablescauriState) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(StateOwner))
}

func (s MutablescauriState) PricePerMg() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StatePricePerMg))
}

func (s MutablescauriState) ProducersBalances() MapAgentIDToMutableUint64 {
	return MapAgentIDToMutableUint64{proxy: s.proxy.Root(StateProducersBalances)}
}

func (s MutablescauriState) Productpasses() MapHashToMutableProductPass {
	return MapHashToMutableProductPass{proxy: s.proxy.Root(StateProductpasses)}
}

func (s MutablescauriState) RecyCompositions() MapHashToMutableRecyCompositions {
	return MapHashToMutableRecyCompositions{proxy: s.proxy.Root(StateRecyCompositions)}
}

func (s MutablescauriState) Recyclates() MapHashToMutableRecyclate {
	return MapHashToMutableRecyclate{proxy: s.proxy.Root(StateRecyclates)}
}

func (s MutablescauriState) Recyclers() ArrayOfMutableAgentID {
	return ArrayOfMutableAgentID{proxy: s.proxy.Root(StateRecyclers)}
}

func (s MutablescauriState) RecyclersBalances() MapAgentIDToMutableUint64 {
	return MapAgentIDToMutableUint64{proxy: s.proxy.Root(StateRecyclersBalances)}
}

func (s MutablescauriState) ShareRecycler() wasmtypes.ScMutableUint8 {
	return wasmtypes.NewScMutableUint8(s.proxy.Root(StateShareRecycler))
}

func (s MutablescauriState) Sorters() ArrayOfMutableAgentID {
	return ArrayOfMutableAgentID{proxy: s.proxy.Root(StateSorters)}
}

func (s MutablescauriState) TokenToDonate() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StateTokenToDonate))
}
