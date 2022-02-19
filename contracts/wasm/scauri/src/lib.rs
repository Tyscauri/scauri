// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use scauri::*;
use wasmlib::*;

use crate::consts::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;
use crate::structs::*;
use crate::typedefs::*;

mod consts;
mod contract;
mod params;
mod results;
mod state;
mod structs;
mod typedefs;

mod scauri;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_ADD_PP_TO_FRACTION,
    	FUNC_ADD_RECYCLER,
    	FUNC_ADD_SORTER,
    	FUNC_CREATE_FRACTION,
    	FUNC_CREATE_PP,
    	FUNC_CREATE_RECYCLATE,
    	FUNC_DELETE_PP,
    	FUNC_INIT,
    	FUNC_PAYOUT_PRODUCER,
    	FUNC_SET_OWNER,
    	VIEW_GET_AMOUNT_OF_REQUIRED_FUNDS,
    	VIEW_GET_FRACTION,
    	VIEW_GET_MATERIALS,
    	VIEW_GET_OWNER,
    	VIEW_GET_PP,
    	VIEW_GET_RECYCLATE,
    	VIEW_GET_TOKEN_PER_PACKAGE,
	],
    funcs: &[
    	func_add_pp_to_fraction_thunk,
    	func_add_recycler_thunk,
    	func_add_sorter_thunk,
    	func_create_fraction_thunk,
    	func_create_pp_thunk,
    	func_create_recyclate_thunk,
    	func_delete_pp_thunk,
    	func_init_thunk,
    	func_payout_producer_thunk,
    	func_set_owner_thunk,
	],
    views: &[
    	view_get_amount_of_required_funds_thunk,
    	view_get_fraction_thunk,
    	view_get_materials_thunk,
    	view_get_owner_thunk,
    	view_get_pp_thunk,
    	view_get_recyclate_thunk,
    	view_get_token_per_package_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	ScExports::call(index, &EXPORT_MAP);
}

#[no_mangle]
fn on_load() {
    ScExports::export(&EXPORT_MAP);
}

pub struct AddPPToFractionContext {
	params: ImmutableAddPPToFractionParams,
	results: MutableAddPPToFractionResults,
	state: MutablescauriState,
}

fn func_add_pp_to_fraction_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcAddPPToFraction");
	let f = AddPPToFractionContext {
		params: ImmutableAddPPToFractionParams { proxy: params_proxy() },
		results: MutableAddPPToFractionResults { proxy: results_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	ctx.require(f.params.pp_id().exists(), "missing mandatory ppID");
	func_add_pp_to_fraction(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.funcAddPPToFraction ok");
}

pub struct AddRecyclerContext {
	params: ImmutableAddRecyclerParams,
	state: MutablescauriState,
}

fn func_add_recycler_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcAddRecycler");
	let f = AddRecyclerContext {
		params: ImmutableAddRecyclerParams { proxy: params_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.recycler_id().exists(), "missing mandatory recyclerID");
	func_add_recycler(ctx, &f);
	ctx.log("scauri.funcAddRecycler ok");
}

pub struct AddSorterContext {
	params: ImmutableAddSorterParams,
	state: MutablescauriState,
}

fn func_add_sorter_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcAddSorter");
	let f = AddSorterContext {
		params: ImmutableAddSorterParams { proxy: params_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.sorter_id().exists(), "missing mandatory sorterID");
	func_add_sorter(ctx, &f);
	ctx.log("scauri.funcAddSorter ok");
}

pub struct CreateFractionContext {
	params: ImmutableCreateFractionParams,
	results: MutableCreateFractionResults,
	state: MutablescauriState,
}

fn func_create_fraction_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcCreateFraction");
	let f = CreateFractionContext {
		params: ImmutableCreateFractionParams { proxy: params_proxy() },
		results: MutableCreateFractionResults { proxy: results_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.purpose().exists(), "missing mandatory purpose");
	func_create_fraction(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.funcCreateFraction ok");
}

pub struct CreatePPContext {
	params: ImmutableCreatePPParams,
	results: MutableCreatePPResults,
	state: MutablescauriState,
}

fn func_create_pp_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcCreatePP");
	let f = CreatePPContext {
		params: ImmutableCreatePPParams { proxy: params_proxy() },
		results: MutableCreatePPResults { proxy: results_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.expiry_date().exists(), "missing mandatory expiryDate");
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.package_weight().exists(), "missing mandatory packageWeight");
	ctx.require(f.params.packages_number().exists(), "missing mandatory packagesNumber");
	ctx.require(f.params.purpose().exists(), "missing mandatory purpose");
	ctx.require(f.params.recyclate_share().exists(), "missing mandatory recyclateShare");
	func_create_pp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.funcCreatePP ok");
}

pub struct CreateRecyclateContext {
	params: ImmutableCreateRecyclateParams,
	results: MutableCreateRecyclateResults,
	state: MutablescauriState,
}

fn func_create_recyclate_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcCreateRecyclate");
	let f = CreateRecyclateContext {
		params: ImmutableCreateRecyclateParams { proxy: params_proxy() },
		results: MutableCreateRecyclateResults { proxy: results_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	func_create_recyclate(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.funcCreateRecyclate ok");
}

pub struct DeletePPContext {
	params: ImmutableDeletePPParams,
	results: MutableDeletePPResults,
	state: MutablescauriState,
}

fn func_delete_pp_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcDeletePP");
	let f = DeletePPContext {
		params: ImmutableDeletePPParams { proxy: params_proxy() },
		results: MutableDeletePPResults { proxy: results_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.pp_id().exists(), "missing mandatory ppID");
	func_delete_pp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.funcDeletePP ok");
}

pub struct InitContext {
	params: ImmutableInitParams,
	state: MutablescauriState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcInit");
	let f = InitContext {
		params: ImmutableInitParams { proxy: params_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	func_init(ctx, &f);
	ctx.log("scauri.funcInit ok");
}

pub struct PayoutProducerContext {
	params: ImmutablePayoutProducerParams,
	state: MutablescauriState,
}

fn func_payout_producer_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcPayoutProducer");
	let f = PayoutProducerContext {
		params: ImmutablePayoutProducerParams { proxy: params_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	func_payout_producer(ctx, &f);
	ctx.log("scauri.funcPayoutProducer ok");
}

pub struct SetOwnerContext {
	params: ImmutableSetOwnerParams,
	state: MutablescauriState,
}

fn func_set_owner_thunk(ctx: &ScFuncContext) {
	ctx.log("scauri.funcSetOwner");
	let f = SetOwnerContext {
		params: ImmutableSetOwnerParams { proxy: params_proxy() },
		state: MutablescauriState { proxy: state_proxy() },
	};

	// current owner of this smart contract
	let access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	ctx.require(f.params.owner().exists(), "missing mandatory owner");
	func_set_owner(ctx, &f);
	ctx.log("scauri.funcSetOwner ok");
}

pub struct GetAmountOfRequiredFundsContext {
	params: ImmutableGetAmountOfRequiredFundsParams,
	results: MutableGetAmountOfRequiredFundsResults,
	state: ImmutablescauriState,
}

fn view_get_amount_of_required_funds_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetAmountOfRequiredFunds");
	let f = GetAmountOfRequiredFundsContext {
		params: ImmutableGetAmountOfRequiredFundsParams { proxy: params_proxy() },
		results: MutableGetAmountOfRequiredFundsResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.charge_weight().exists(), "missing mandatory chargeWeight");
	view_get_amount_of_required_funds(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetAmountOfRequiredFunds ok");
}

pub struct GetFractionContext {
	params: ImmutableGetFractionParams,
	results: MutableGetFractionResults,
	state: ImmutablescauriState,
}

fn view_get_fraction_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetFraction");
	let f = GetFractionContext {
		params: ImmutableGetFractionParams { proxy: params_proxy() },
		results: MutableGetFractionResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.frac_id().exists(), "missing mandatory fracID");
	view_get_fraction(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetFraction ok");
}

pub struct GetMaterialsContext {
	params: ImmutableGetMaterialsParams,
	results: MutableGetMaterialsResults,
	state: ImmutablescauriState,
}

fn view_get_materials_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetMaterials");
	let f = GetMaterialsContext {
		params: ImmutableGetMaterialsParams { proxy: params_proxy() },
		results: MutableGetMaterialsResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.id().exists(), "missing mandatory id");
	view_get_materials(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetMaterials ok");
}

pub struct GetOwnerContext {
	results: MutableGetOwnerResults,
	state: ImmutablescauriState,
}

fn view_get_owner_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetOwner");
	let f = GetOwnerContext {
		results: MutableGetOwnerResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	view_get_owner(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetOwner ok");
}

pub struct GetPPContext {
	params: ImmutableGetPPParams,
	results: MutableGetPPResults,
	state: ImmutablescauriState,
}

fn view_get_pp_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetPP");
	let f = GetPPContext {
		params: ImmutableGetPPParams { proxy: params_proxy() },
		results: MutableGetPPResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.id().exists(), "missing mandatory id");
	view_get_pp(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetPP ok");
}

pub struct GetRecyclateContext {
	params: ImmutableGetRecyclateParams,
	results: MutableGetRecyclateResults,
	state: ImmutablescauriState,
}

fn view_get_recyclate_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetRecyclate");
	let f = GetRecyclateContext {
		params: ImmutableGetRecyclateParams { proxy: params_proxy() },
		results: MutableGetRecyclateResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	ctx.require(f.params.recy_id().exists(), "missing mandatory recyID");
	view_get_recyclate(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetRecyclate ok");
}

pub struct GetTokenPerPackageContext {
	params: ImmutableGetTokenPerPackageParams,
	results: MutableGetTokenPerPackageResults,
	state: ImmutablescauriState,
}

fn view_get_token_per_package_thunk(ctx: &ScViewContext) {
	ctx.log("scauri.viewGetTokenPerPackage");
	let f = GetTokenPerPackageContext {
		params: ImmutableGetTokenPerPackageParams { proxy: params_proxy() },
		results: MutableGetTokenPerPackageResults { proxy: results_proxy() },
		state: ImmutablescauriState { proxy: state_proxy() },
	};
	view_get_token_per_package(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("scauri.viewGetTokenPerPackage ok");
}
