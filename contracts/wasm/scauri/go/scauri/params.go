// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package scauri

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableAddPPToFractionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddPPToFractionParams) FracID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamFracID))
}

func (s ImmutableAddPPToFractionParams) PpID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamPpID))
}

type MutableAddPPToFractionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddPPToFractionParams) FracID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamFracID))
}

func (s MutableAddPPToFractionParams) PpID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamPpID))
}

type ImmutableAddRecyclerParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddRecyclerParams) RecyclerID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamRecyclerID))
}

type MutableAddRecyclerParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddRecyclerParams) RecyclerID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamRecyclerID))
}

type ImmutableAddSorterParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableAddSorterParams) SorterID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamSorterID))
}

type MutableAddSorterParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableAddSorterParams) SorterID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamSorterID))
}

type ImmutableCreateFractionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableCreateFractionParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableCreateFractionParams) Purpose() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamPurpose))
}

type MutableCreateFractionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableCreateFractionParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableCreateFractionParams) Purpose() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamPurpose))
}

type ImmutableCreatePPParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableCreatePPParams) Compositions() ImmutableCompositions {
	return ImmutableCompositions{proxy: s.proxy.Root(ParamCompositions)}
}

func (s ImmutableCreatePPParams) ExpiryDate() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamExpiryDate))
}

func (s ImmutableCreatePPParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableCreatePPParams) PackageWeight() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamPackageWeight))
}

func (s ImmutableCreatePPParams) PackagesNumber() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamPackagesNumber))
}

func (s ImmutableCreatePPParams) Purpose() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamPurpose))
}

type MutableCreatePPParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableCreatePPParams) Compositions() MutableCompositions {
	return MutableCompositions{proxy: s.proxy.Root(ParamCompositions)}
}

func (s MutableCreatePPParams) ExpiryDate() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamExpiryDate))
}

func (s MutableCreatePPParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableCreatePPParams) PackageWeight() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamPackageWeight))
}

func (s MutableCreatePPParams) PackagesNumber() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamPackagesNumber))
}

func (s MutableCreatePPParams) Purpose() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamPurpose))
}

type ImmutableCreateRecyclateParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableCreateRecyclateParams) FracID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamFracID))
}

func (s ImmutableCreateRecyclateParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableCreateRecyclateParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableCreateRecyclateParams) FracID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamFracID))
}

func (s MutableCreateRecyclateParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableDeletePPParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableDeletePPParams) PpID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamPpID))
}

type MutableDeletePPParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableDeletePPParams) PpID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamPpID))
}

type ImmutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableInitParams) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOwner))
}

type MutableInitParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableInitParams) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOwner))
}

type ImmutablePayoutProducerParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutablePayoutProducerParams) PpID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamPpID))
}

type MutablePayoutProducerParams struct {
	proxy wasmtypes.Proxy
}

func (s MutablePayoutProducerParams) PpID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamPpID))
}

type ImmutableSetDonationAddressParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSetDonationAddressParams) DonationAddress() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamDonationAddress))
}

type MutableSetDonationAddressParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSetDonationAddressParams) DonationAddress() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamDonationAddress))
}

type ImmutableSetOwnerParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableSetOwnerParams) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOwner))
}

type MutableSetOwnerParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableSetOwnerParams) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOwner))
}

type ImmutableGetAmountOfRequiredFundsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetAmountOfRequiredFundsParams) ChargeWeight() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamChargeWeight))
}

type MutableGetAmountOfRequiredFundsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetAmountOfRequiredFundsParams) ChargeWeight() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamChargeWeight))
}

type ImmutableGetFractionParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetFractionParams) FracID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamFracID))
}

type MutableGetFractionParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetFractionParams) FracID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamFracID))
}

type ImmutableGetMaterialsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetMaterialsParams) Id() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamId))
}

type MutableGetMaterialsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetMaterialsParams) Id() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamId))
}

type ImmutableGetPPParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetPPParams) Id() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamId))
}

type MutableGetPPParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetPPParams) Id() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamId))
}

type ImmutableGetRecyclateParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetRecyclateParams) RecyID() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamRecyID))
}

type MutableGetRecyclateParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetRecyclateParams) RecyID() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamRecyID))
}

type ImmutableGetTokenPerPackageParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableGetTokenPerPackageParams) ProdPass() ImmutableProductPass {
	return ImmutableProductPass{proxy: s.proxy.Root(ParamProdPass)}
}

type MutableGetTokenPerPackageParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableGetTokenPerPackageParams) ProdPass() MutableProductPass {
	return MutableProductPass{proxy: s.proxy.Root(ParamProdPass)}
}
