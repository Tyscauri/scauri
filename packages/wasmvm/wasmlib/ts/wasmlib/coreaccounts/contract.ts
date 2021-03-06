// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class DepositCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncDeposit);
	params: sc.MutableDepositParams = new sc.MutableDepositParams(wasmlib.ScView.nilProxy);
}

export class HarvestCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncHarvest);
	params: sc.MutableHarvestParams = new sc.MutableHarvestParams(wasmlib.ScView.nilProxy);
}

export class WithdrawCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncWithdraw);
}

export class AccountsCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewAccounts);
	results: sc.ImmutableAccountsResults = new sc.ImmutableAccountsResults(wasmlib.ScView.nilProxy);
}

export class BalanceCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewBalance);
	params: sc.MutableBalanceParams = new sc.MutableBalanceParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableBalanceResults = new sc.ImmutableBalanceResults(wasmlib.ScView.nilProxy);
}

export class GetAccountNonceCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetAccountNonce);
	params: sc.MutableGetAccountNonceParams = new sc.MutableGetAccountNonceParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetAccountNonceResults = new sc.ImmutableGetAccountNonceResults(wasmlib.ScView.nilProxy);
}

export class TotalAssetsCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewTotalAssets);
	results: sc.ImmutableTotalAssetsResults = new sc.ImmutableTotalAssetsResults(wasmlib.ScView.nilProxy);
}

export class ScFuncs {
	static deposit(_ctx: wasmlib.ScFuncCallContext): DepositCall {
		const f = new DepositCall();
		f.params = new sc.MutableDepositParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static harvest(_ctx: wasmlib.ScFuncCallContext): HarvestCall {
		const f = new HarvestCall();
		f.params = new sc.MutableHarvestParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static withdraw(_ctx: wasmlib.ScFuncCallContext): WithdrawCall {
		return new WithdrawCall();
	}

	static accounts(_ctx: wasmlib.ScViewCallContext): AccountsCall {
		const f = new AccountsCall();
		f.results = new sc.ImmutableAccountsResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static balance(_ctx: wasmlib.ScViewCallContext): BalanceCall {
		const f = new BalanceCall();
		f.params = new sc.MutableBalanceParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableBalanceResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getAccountNonce(_ctx: wasmlib.ScViewCallContext): GetAccountNonceCall {
		const f = new GetAccountNonceCall();
		f.params = new sc.MutableGetAccountNonceParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetAccountNonceResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static totalAssets(_ctx: wasmlib.ScViewCallContext): TotalAssetsCall {
		const f = new TotalAssetsCall();
		f.results = new sc.ImmutableTotalAssetsResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}
}
