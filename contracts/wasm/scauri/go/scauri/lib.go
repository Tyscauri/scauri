// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package scauri

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
    	FuncAddPPToFraction,
    	FuncAddRecycler,
    	FuncAddSorter,
    	FuncCreateFraction,
    	FuncCreatePP,
    	FuncCreateRecyclate,
    	FuncDeletePP,
    	FuncInit,
    	FuncPayoutProducer,
    	FuncSetOwner,
    	ViewGetAmountOfRequiredFunds,
    	ViewGetFraction,
    	ViewGetMaterials,
    	ViewGetOwner,
    	ViewGetPP,
    	ViewGetRecyclate,
    	ViewGetTokenPerPackage,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
    	funcAddPPToFractionThunk,
    	funcAddRecyclerThunk,
    	funcAddSorterThunk,
    	funcCreateFractionThunk,
    	funcCreatePPThunk,
    	funcCreateRecyclateThunk,
    	funcDeletePPThunk,
    	funcInitThunk,
    	funcPayoutProducerThunk,
    	funcSetOwnerThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
    	viewGetAmountOfRequiredFundsThunk,
    	viewGetFractionThunk,
    	viewGetMaterialsThunk,
    	viewGetOwnerThunk,
    	viewGetPPThunk,
    	viewGetRecyclateThunk,
    	viewGetTokenPerPackageThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type AddPPToFractionContext struct {
	Params  ImmutableAddPPToFractionParams
	Results MutableAddPPToFractionResults
	State   MutablescauriState
}

func funcAddPPToFractionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcAddPPToFraction")
	results := wasmlib.NewScDict()
	f := &AddPPToFractionContext{
		Params: ImmutableAddPPToFractionParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableAddPPToFractionResults{
			proxy: results.AsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.FracID().Exists(), "missing mandatory fracID")
	ctx.Require(f.Params.PpID().Exists(), "missing mandatory ppID")
	funcAddPPToFraction(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.funcAddPPToFraction ok")
}

type AddRecyclerContext struct {
	Params  ImmutableAddRecyclerParams
	State   MutablescauriState
}

func funcAddRecyclerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcAddRecycler")
	f := &AddRecyclerContext{
		Params: ImmutableAddRecyclerParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	access := f.State.Owner()
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	ctx.Require(f.Params.RecyclerID().Exists(), "missing mandatory recyclerID")
	funcAddRecycler(ctx, f)
	ctx.Log("scauri.funcAddRecycler ok")
}

type AddSorterContext struct {
	Params  ImmutableAddSorterParams
	State   MutablescauriState
}

func funcAddSorterThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcAddSorter")
	f := &AddSorterContext{
		Params: ImmutableAddSorterParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	access := f.State.Owner()
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	ctx.Require(f.Params.SorterID().Exists(), "missing mandatory sorterID")
	funcAddSorter(ctx, f)
	ctx.Log("scauri.funcAddSorter ok")
}

type CreateFractionContext struct {
	Params  ImmutableCreateFractionParams
	Results MutableCreateFractionResults
	State   MutablescauriState
}

func funcCreateFractionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcCreateFraction")
	results := wasmlib.NewScDict()
	f := &CreateFractionContext{
		Params: ImmutableCreateFractionParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableCreateFractionResults{
			proxy: results.AsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Purpose().Exists(), "missing mandatory purpose")
	funcCreateFraction(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.funcCreateFraction ok")
}

type CreatePPContext struct {
	Params  ImmutableCreatePPParams
	Results MutableCreatePPResults
	State   MutablescauriState
}

func funcCreatePPThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcCreatePP")
	results := wasmlib.NewScDict()
	f := &CreatePPContext{
		Params: ImmutableCreatePPParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableCreatePPResults{
			proxy: results.AsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.ExpiryDate().Exists(), "missing mandatory expiryDate")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	ctx.Require(f.Params.PackageWeight().Exists(), "missing mandatory packageWeight")
	ctx.Require(f.Params.PackagesNumber().Exists(), "missing mandatory packagesNumber")
	ctx.Require(f.Params.Purpose().Exists(), "missing mandatory purpose")
	ctx.Require(f.Params.RecyclateShare().Exists(), "missing mandatory recyclateShare")
	funcCreatePP(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.funcCreatePP ok")
}

type CreateRecyclateContext struct {
	Params  ImmutableCreateRecyclateParams
	Results MutableCreateRecyclateResults
	State   MutablescauriState
}

func funcCreateRecyclateThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcCreateRecyclate")
	results := wasmlib.NewScDict()
	f := &CreateRecyclateContext{
		Params: ImmutableCreateRecyclateParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableCreateRecyclateResults{
			proxy: results.AsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.FracID().Exists(), "missing mandatory fracID")
	funcCreateRecyclate(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.funcCreateRecyclate ok")
}

type DeletePPContext struct {
	Params  ImmutableDeletePPParams
	Results MutableDeletePPResults
	State   MutablescauriState
}

func funcDeletePPThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcDeletePP")
	results := wasmlib.NewScDict()
	f := &DeletePPContext{
		Params: ImmutableDeletePPParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableDeletePPResults{
			proxy: results.AsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.PpID().Exists(), "missing mandatory ppID")
	funcDeletePP(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.funcDeletePP ok")
}

type InitContext struct {
	Params  ImmutableInitParams
	State   MutablescauriState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcInit(ctx, f)
	ctx.Log("scauri.funcInit ok")
}

type PayoutProducerContext struct {
	Params  ImmutablePayoutProducerParams
	State   MutablescauriState
}

func funcPayoutProducerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcPayoutProducer")
	f := &PayoutProducerContext{
		Params: ImmutablePayoutProducerParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.FracID().Exists(), "missing mandatory fracID")
	funcPayoutProducer(ctx, f)
	ctx.Log("scauri.funcPayoutProducer ok")
}

type SetOwnerContext struct {
	Params  ImmutableSetOwnerParams
	State   MutablescauriState
}

func funcSetOwnerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("scauri.funcSetOwner")
	f := &SetOwnerContext{
		Params: ImmutableSetOwnerParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// current owner of this smart contract
	access := f.State.Owner()
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	ctx.Require(f.Params.Owner().Exists(), "missing mandatory owner")
	funcSetOwner(ctx, f)
	ctx.Log("scauri.funcSetOwner ok")
}

type GetAmountOfRequiredFundsContext struct {
	Params  ImmutableGetAmountOfRequiredFundsParams
	Results MutableGetAmountOfRequiredFundsResults
	State   ImmutablescauriState
}

func viewGetAmountOfRequiredFundsThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetAmountOfRequiredFunds")
	results := wasmlib.NewScDict()
	f := &GetAmountOfRequiredFundsContext{
		Params: ImmutableGetAmountOfRequiredFundsParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetAmountOfRequiredFundsResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.ChargeWeight().Exists(), "missing mandatory chargeWeight")
	viewGetAmountOfRequiredFunds(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetAmountOfRequiredFunds ok")
}

type GetFractionContext struct {
	Params  ImmutableGetFractionParams
	Results MutableGetFractionResults
	State   ImmutablescauriState
}

func viewGetFractionThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetFraction")
	results := wasmlib.NewScDict()
	f := &GetFractionContext{
		Params: ImmutableGetFractionParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetFractionResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.FracID().Exists(), "missing mandatory fracID")
	viewGetFraction(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetFraction ok")
}

type GetMaterialsContext struct {
	Params  ImmutableGetMaterialsParams
	Results MutableGetMaterialsResults
	State   ImmutablescauriState
}

func viewGetMaterialsThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetMaterials")
	results := wasmlib.NewScDict()
	f := &GetMaterialsContext{
		Params: ImmutableGetMaterialsParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetMaterialsResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Id().Exists(), "missing mandatory id")
	viewGetMaterials(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetMaterials ok")
}

type GetOwnerContext struct {
	Results MutableGetOwnerResults
	State   ImmutablescauriState
}

func viewGetOwnerThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetOwner")
	results := wasmlib.NewScDict()
	f := &GetOwnerContext{
		Results: MutableGetOwnerResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetOwner(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetOwner ok")
}

type GetPPContext struct {
	Params  ImmutableGetPPParams
	Results MutableGetPPResults
	State   ImmutablescauriState
}

func viewGetPPThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetPP")
	results := wasmlib.NewScDict()
	f := &GetPPContext{
		Params: ImmutableGetPPParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetPPResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Id().Exists(), "missing mandatory id")
	viewGetPP(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetPP ok")
}

type GetRecyclateContext struct {
	Params  ImmutableGetRecyclateParams
	Results MutableGetRecyclateResults
	State   ImmutablescauriState
}

func viewGetRecyclateThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetRecyclate")
	results := wasmlib.NewScDict()
	f := &GetRecyclateContext{
		Params: ImmutableGetRecyclateParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetRecyclateResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.RecyID().Exists(), "missing mandatory recyID")
	viewGetRecyclate(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetRecyclate ok")
}

type GetTokenPerPackageContext struct {
	Params  ImmutableGetTokenPerPackageParams
	Results MutableGetTokenPerPackageResults
	State   ImmutablescauriState
}

func viewGetTokenPerPackageThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("scauri.viewGetTokenPerPackage")
	results := wasmlib.NewScDict()
	f := &GetTokenPerPackageContext{
		Params: ImmutableGetTokenPerPackageParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetTokenPerPackageResults{
			proxy: results.AsProxy(),
		},
		State: ImmutablescauriState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetTokenPerPackage(ctx, f)
	ctx.Results(results)
	ctx.Log("scauri.viewGetTokenPerPackage ok")
}
