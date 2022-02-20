// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package scauri

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "scauri"
	ScDescription = "scauri description"
	HScName       = wasmtypes.ScHname(0x5559a7fd)
)

const (
	ParamChargeWeight   = "chargeWeight"
	ParamCompositions   = "compositions"
	ParamExpiryDate     = "expiryDate"
	ParamFracID         = "fracID"
	ParamId             = "id"
	ParamName           = "name"
	ParamOwner          = "owner"
	ParamPackageWeight  = "packageWeight"
	ParamPackagesNumber = "packagesNumber"
	ParamPpID           = "ppID"
	ParamProdPass       = "prodPass"
	ParamPurpose        = "purpose"
	ParamRecyID         = "recyID"
	ParamRecyclerID     = "recyclerID"
	ParamSorterID       = "sorterID"
)

const (
	ResultCompositions    = "compositions"
	ResultFracComposition = "fracComposition"
	ResultFracID          = "fracID"
	ResultFraction        = "fraction"
	ResultId              = "id"
	ResultOwner           = "owner"
	ResultPp              = "pp"
	ResultPpname          = "ppname"
	ResultPpresult        = "ppresult"
	ResultRecyComposition = "recyComposition"
	ResultRecyID          = "recyID"
	ResultRecyclate       = "recyclate"
	ResultTokenPerPackage = "tokenPerPackage"
	ResultTokenRequired   = "tokenRequired"
)

const (
	StateCompositions     = "compositions"
	StateDonationAddress  = "donationAddress"
	StateFracCompositions = "fracCompositions"
	StateFractions        = "fractions"
	StateOwner            = "owner"
	StatePricePerMg       = "pricePerMg"
	StateProductpasses    = "productpasses"
	StateRecyCompositions = "recyCompositions"
	StateRecyclates       = "recyclates"
	StateRecyclers        = "recyclers"
	StateShareRecycler    = "shareRecycler"
	StateSorters          = "sorters"
	StateTokenToDonate    = "tokenToDonate"
)

const (
	FuncAddPPToFraction          = "addPPToFraction"
	FuncAddRecycler              = "addRecycler"
	FuncAddSorter                = "addSorter"
	FuncCreateFraction           = "createFraction"
	FuncCreatePP                 = "createPP"
	FuncCreateRecyclate          = "createRecyclate"
	FuncDeletePP                 = "deletePP"
	FuncInit                     = "init"
	FuncPayoutProducer           = "payoutProducer"
	FuncSetOwner                 = "setOwner"
	ViewGetAmountOfRequiredFunds = "getAmountOfRequiredFunds"
	ViewGetFraction              = "getFraction"
	ViewGetMaterials             = "getMaterials"
	ViewGetOwner                 = "getOwner"
	ViewGetPP                    = "getPP"
	ViewGetRecyclate             = "getRecyclate"
	ViewGetTokenPerPackage       = "getTokenPerPackage"
)

const (
	HFuncAddPPToFraction          = wasmtypes.ScHname(0x50ac50a7)
	HFuncAddRecycler              = wasmtypes.ScHname(0x25ae8407)
	HFuncAddSorter                = wasmtypes.ScHname(0x735539b5)
	HFuncCreateFraction           = wasmtypes.ScHname(0x59842fc3)
	HFuncCreatePP                 = wasmtypes.ScHname(0x673fc3d7)
	HFuncCreateRecyclate          = wasmtypes.ScHname(0x5066d840)
	HFuncDeletePP                 = wasmtypes.ScHname(0x56dedc36)
	HFuncInit                     = wasmtypes.ScHname(0x1f44d644)
	HFuncPayoutProducer           = wasmtypes.ScHname(0x3a56494b)
	HFuncSetOwner                 = wasmtypes.ScHname(0x2a15fe7b)
	HViewGetAmountOfRequiredFunds = wasmtypes.ScHname(0xbcbaf2d6)
	HViewGetFraction              = wasmtypes.ScHname(0xe1c332fa)
	HViewGetMaterials             = wasmtypes.ScHname(0x1dca8ddb)
	HViewGetOwner                 = wasmtypes.ScHname(0x137107a6)
	HViewGetPP                    = wasmtypes.ScHname(0xbf711e08)
	HViewGetRecyclate             = wasmtypes.ScHname(0x4e87f53b)
	HViewGetTokenPerPackage       = wasmtypes.ScHname(0x522a14c0)
)
