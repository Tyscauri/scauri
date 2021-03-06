// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class CallOnChainCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncCallOnChain);
	params: sc.MutableCallOnChainParams = new sc.MutableCallOnChainParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableCallOnChainResults = new sc.ImmutableCallOnChainResults(wasmlib.ScView.nilProxy);
}

export class CallOnChainContext {
	params: sc.ImmutableCallOnChainParams = new sc.ImmutableCallOnChainParams(wasmlib.paramsProxy());
	results: sc.MutableCallOnChainResults = new sc.MutableCallOnChainResults(wasmlib.ScView.nilProxy);
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class CheckContextFromFullEPCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncCheckContextFromFullEP);
	params: sc.MutableCheckContextFromFullEPParams = new sc.MutableCheckContextFromFullEPParams(wasmlib.ScView.nilProxy);
}

export class CheckContextFromFullEPContext {
	params: sc.ImmutableCheckContextFromFullEPParams = new sc.ImmutableCheckContextFromFullEPParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class DoNothingCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncDoNothing);
}

export class DoNothingContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class GetMintedSupplyCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncGetMintedSupply);
	results: sc.ImmutableGetMintedSupplyResults = new sc.ImmutableGetMintedSupplyResults(wasmlib.ScView.nilProxy);
}

export class GetMintedSupplyContext {
	results: sc.MutableGetMintedSupplyResults = new sc.MutableGetMintedSupplyResults(wasmlib.ScView.nilProxy);
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class IncCounterCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncIncCounter);
}

export class IncCounterContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class InitCall {
	func: wasmlib.ScInitFunc = new wasmlib.ScInitFunc(sc.HScName, sc.HFuncInit);
	params: sc.MutableInitParams = new sc.MutableInitParams(wasmlib.ScView.nilProxy);
}

export class InitContext {
	params: sc.ImmutableInitParams = new sc.ImmutableInitParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class PassTypesFullCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncPassTypesFull);
	params: sc.MutablePassTypesFullParams = new sc.MutablePassTypesFullParams(wasmlib.ScView.nilProxy);
}

export class PassTypesFullContext {
	params: sc.ImmutablePassTypesFullParams = new sc.ImmutablePassTypesFullParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class RunRecursionCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncRunRecursion);
	params: sc.MutableRunRecursionParams = new sc.MutableRunRecursionParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableRunRecursionResults = new sc.ImmutableRunRecursionResults(wasmlib.ScView.nilProxy);
}

export class RunRecursionContext {
	params: sc.ImmutableRunRecursionParams = new sc.ImmutableRunRecursionParams(wasmlib.paramsProxy());
	results: sc.MutableRunRecursionResults = new sc.MutableRunRecursionResults(wasmlib.ScView.nilProxy);
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class SendToAddressCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncSendToAddress);
	params: sc.MutableSendToAddressParams = new sc.MutableSendToAddressParams(wasmlib.ScView.nilProxy);
}

export class SendToAddressContext {
	params: sc.ImmutableSendToAddressParams = new sc.ImmutableSendToAddressParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class SetIntCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncSetInt);
	params: sc.MutableSetIntParams = new sc.MutableSetIntParams(wasmlib.ScView.nilProxy);
}

export class SetIntContext {
	params: sc.ImmutableSetIntParams = new sc.ImmutableSetIntParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class SpawnCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncSpawn);
	params: sc.MutableSpawnParams = new sc.MutableSpawnParams(wasmlib.ScView.nilProxy);
}

export class SpawnContext {
	params: sc.ImmutableSpawnParams = new sc.ImmutableSpawnParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestBlockContext1Call {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestBlockContext1);
}

export class TestBlockContext1Context {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestBlockContext2Call {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestBlockContext2);
}

export class TestBlockContext2Context {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestCallPanicFullEPCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestCallPanicFullEP);
}

export class TestCallPanicFullEPContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestCallPanicViewEPFromFullCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestCallPanicViewEPFromFull);
}

export class TestCallPanicViewEPFromFullContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestChainOwnerIDFullCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestChainOwnerIDFull);
	results: sc.ImmutableTestChainOwnerIDFullResults = new sc.ImmutableTestChainOwnerIDFullResults(wasmlib.ScView.nilProxy);
}

export class TestChainOwnerIDFullContext {
	results: sc.MutableTestChainOwnerIDFullResults = new sc.MutableTestChainOwnerIDFullResults(wasmlib.ScView.nilProxy);
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestEventLogDeployCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestEventLogDeploy);
}

export class TestEventLogDeployContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestEventLogEventDataCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestEventLogEventData);
}

export class TestEventLogEventDataContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestEventLogGenericDataCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestEventLogGenericData);
	params: sc.MutableTestEventLogGenericDataParams = new sc.MutableTestEventLogGenericDataParams(wasmlib.ScView.nilProxy);
}

export class TestEventLogGenericDataContext {
	params: sc.ImmutableTestEventLogGenericDataParams = new sc.ImmutableTestEventLogGenericDataParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestPanicFullEPCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncTestPanicFullEP);
}

export class TestPanicFullEPContext {
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class WithdrawToChainCall {
	func: wasmlib.ScFunc = new wasmlib.ScFunc(sc.HScName, sc.HFuncWithdrawToChain);
	params: sc.MutableWithdrawToChainParams = new sc.MutableWithdrawToChainParams(wasmlib.ScView.nilProxy);
}

export class WithdrawToChainContext {
	params: sc.ImmutableWithdrawToChainParams = new sc.ImmutableWithdrawToChainParams(wasmlib.paramsProxy());
	state: sc.MutableTestCoreState = new sc.MutableTestCoreState(wasmlib.ScState.proxy());
}

export class CheckContextFromViewEPCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewCheckContextFromViewEP);
	params: sc.MutableCheckContextFromViewEPParams = new sc.MutableCheckContextFromViewEPParams(wasmlib.ScView.nilProxy);
}

export class CheckContextFromViewEPContext {
	params: sc.ImmutableCheckContextFromViewEPParams = new sc.ImmutableCheckContextFromViewEPParams(wasmlib.paramsProxy());
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class FibonacciCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewFibonacci);
	params: sc.MutableFibonacciParams = new sc.MutableFibonacciParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableFibonacciResults = new sc.ImmutableFibonacciResults(wasmlib.ScView.nilProxy);
}

export class FibonacciContext {
	params: sc.ImmutableFibonacciParams = new sc.ImmutableFibonacciParams(wasmlib.paramsProxy());
	results: sc.MutableFibonacciResults = new sc.MutableFibonacciResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class GetCounterCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetCounter);
	results: sc.ImmutableGetCounterResults = new sc.ImmutableGetCounterResults(wasmlib.ScView.nilProxy);
}

export class GetCounterContext {
	results: sc.MutableGetCounterResults = new sc.MutableGetCounterResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class GetIntCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetInt);
	params: sc.MutableGetIntParams = new sc.MutableGetIntParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetIntResults = new sc.ImmutableGetIntResults(wasmlib.ScView.nilProxy);
}

export class GetIntContext {
	params: sc.ImmutableGetIntParams = new sc.ImmutableGetIntParams(wasmlib.paramsProxy());
	results: sc.MutableGetIntResults = new sc.MutableGetIntResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class GetStringValueCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewGetStringValue);
	params: sc.MutableGetStringValueParams = new sc.MutableGetStringValueParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetStringValueResults = new sc.ImmutableGetStringValueResults(wasmlib.ScView.nilProxy);
}

export class GetStringValueContext {
	params: sc.ImmutableGetStringValueParams = new sc.ImmutableGetStringValueParams(wasmlib.paramsProxy());
	results: sc.MutableGetStringValueResults = new sc.MutableGetStringValueResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class JustViewCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewJustView);
}

export class JustViewContext {
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class PassTypesViewCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewPassTypesView);
	params: sc.MutablePassTypesViewParams = new sc.MutablePassTypesViewParams(wasmlib.ScView.nilProxy);
}

export class PassTypesViewContext {
	params: sc.ImmutablePassTypesViewParams = new sc.ImmutablePassTypesViewParams(wasmlib.paramsProxy());
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestCallPanicViewEPFromViewCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewTestCallPanicViewEPFromView);
}

export class TestCallPanicViewEPFromViewContext {
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestChainOwnerIDViewCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewTestChainOwnerIDView);
	results: sc.ImmutableTestChainOwnerIDViewResults = new sc.ImmutableTestChainOwnerIDViewResults(wasmlib.ScView.nilProxy);
}

export class TestChainOwnerIDViewContext {
	results: sc.MutableTestChainOwnerIDViewResults = new sc.MutableTestChainOwnerIDViewResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestPanicViewEPCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewTestPanicViewEP);
}

export class TestPanicViewEPContext {
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class TestSandboxCallCall {
	func: wasmlib.ScView = new wasmlib.ScView(sc.HScName, sc.HViewTestSandboxCall);
	results: sc.ImmutableTestSandboxCallResults = new sc.ImmutableTestSandboxCallResults(wasmlib.ScView.nilProxy);
}

export class TestSandboxCallContext {
	results: sc.MutableTestSandboxCallResults = new sc.MutableTestSandboxCallResults(wasmlib.ScView.nilProxy);
	state: sc.ImmutableTestCoreState = new sc.ImmutableTestCoreState(wasmlib.ScState.proxy());
}

export class ScFuncs {
	static callOnChain(_ctx: wasmlib.ScFuncCallContext): CallOnChainCall {
		const f = new CallOnChainCall();
		f.params = new sc.MutableCallOnChainParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableCallOnChainResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static checkContextFromFullEP(_ctx: wasmlib.ScFuncCallContext): CheckContextFromFullEPCall {
		const f = new CheckContextFromFullEPCall();
		f.params = new sc.MutableCheckContextFromFullEPParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static doNothing(_ctx: wasmlib.ScFuncCallContext): DoNothingCall {
		return new DoNothingCall();
	}

	static getMintedSupply(_ctx: wasmlib.ScFuncCallContext): GetMintedSupplyCall {
		const f = new GetMintedSupplyCall();
		f.results = new sc.ImmutableGetMintedSupplyResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static incCounter(_ctx: wasmlib.ScFuncCallContext): IncCounterCall {
		return new IncCounterCall();
	}

	static init(_ctx: wasmlib.ScFuncCallContext): InitCall {
		const f = new InitCall();
		f.params = new sc.MutableInitParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static passTypesFull(_ctx: wasmlib.ScFuncCallContext): PassTypesFullCall {
		const f = new PassTypesFullCall();
		f.params = new sc.MutablePassTypesFullParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static runRecursion(_ctx: wasmlib.ScFuncCallContext): RunRecursionCall {
		const f = new RunRecursionCall();
		f.params = new sc.MutableRunRecursionParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableRunRecursionResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static sendToAddress(_ctx: wasmlib.ScFuncCallContext): SendToAddressCall {
		const f = new SendToAddressCall();
		f.params = new sc.MutableSendToAddressParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static setInt(_ctx: wasmlib.ScFuncCallContext): SetIntCall {
		const f = new SetIntCall();
		f.params = new sc.MutableSetIntParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static spawn(_ctx: wasmlib.ScFuncCallContext): SpawnCall {
		const f = new SpawnCall();
		f.params = new sc.MutableSpawnParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static testBlockContext1(_ctx: wasmlib.ScFuncCallContext): TestBlockContext1Call {
		return new TestBlockContext1Call();
	}

	static testBlockContext2(_ctx: wasmlib.ScFuncCallContext): TestBlockContext2Call {
		return new TestBlockContext2Call();
	}

	static testCallPanicFullEP(_ctx: wasmlib.ScFuncCallContext): TestCallPanicFullEPCall {
		return new TestCallPanicFullEPCall();
	}

	static testCallPanicViewEPFromFull(_ctx: wasmlib.ScFuncCallContext): TestCallPanicViewEPFromFullCall {
		return new TestCallPanicViewEPFromFullCall();
	}

	static testChainOwnerIDFull(_ctx: wasmlib.ScFuncCallContext): TestChainOwnerIDFullCall {
		const f = new TestChainOwnerIDFullCall();
		f.results = new sc.ImmutableTestChainOwnerIDFullResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static testEventLogDeploy(_ctx: wasmlib.ScFuncCallContext): TestEventLogDeployCall {
		return new TestEventLogDeployCall();
	}

	static testEventLogEventData(_ctx: wasmlib.ScFuncCallContext): TestEventLogEventDataCall {
		return new TestEventLogEventDataCall();
	}

	static testEventLogGenericData(_ctx: wasmlib.ScFuncCallContext): TestEventLogGenericDataCall {
		const f = new TestEventLogGenericDataCall();
		f.params = new sc.MutableTestEventLogGenericDataParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static testPanicFullEP(_ctx: wasmlib.ScFuncCallContext): TestPanicFullEPCall {
		return new TestPanicFullEPCall();
	}

	static withdrawToChain(_ctx: wasmlib.ScFuncCallContext): WithdrawToChainCall {
		const f = new WithdrawToChainCall();
		f.params = new sc.MutableWithdrawToChainParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static checkContextFromViewEP(_ctx: wasmlib.ScViewCallContext): CheckContextFromViewEPCall {
		const f = new CheckContextFromViewEPCall();
		f.params = new sc.MutableCheckContextFromViewEPParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static fibonacci(_ctx: wasmlib.ScViewCallContext): FibonacciCall {
		const f = new FibonacciCall();
		f.params = new sc.MutableFibonacciParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableFibonacciResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getCounter(_ctx: wasmlib.ScViewCallContext): GetCounterCall {
		const f = new GetCounterCall();
		f.results = new sc.ImmutableGetCounterResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getInt(_ctx: wasmlib.ScViewCallContext): GetIntCall {
		const f = new GetIntCall();
		f.params = new sc.MutableGetIntParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetIntResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getStringValue(_ctx: wasmlib.ScViewCallContext): GetStringValueCall {
		const f = new GetStringValueCall();
		f.params = new sc.MutableGetStringValueParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetStringValueResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static justView(_ctx: wasmlib.ScViewCallContext): JustViewCall {
		return new JustViewCall();
	}

	static passTypesView(_ctx: wasmlib.ScViewCallContext): PassTypesViewCall {
		const f = new PassTypesViewCall();
		f.params = new sc.MutablePassTypesViewParams(wasmlib.newCallParamsProxy(f.func));
		return f;
	}

	static testCallPanicViewEPFromView(_ctx: wasmlib.ScViewCallContext): TestCallPanicViewEPFromViewCall {
		return new TestCallPanicViewEPFromViewCall();
	}

	static testChainOwnerIDView(_ctx: wasmlib.ScViewCallContext): TestChainOwnerIDViewCall {
		const f = new TestChainOwnerIDViewCall();
		f.results = new sc.ImmutableTestChainOwnerIDViewResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static testPanicViewEP(_ctx: wasmlib.ScViewCallContext): TestPanicViewEPCall {
		return new TestPanicViewEPCall();
	}

	static testSandboxCall(_ctx: wasmlib.ScViewCallContext): TestSandboxCallCall {
		const f = new TestSandboxCallCall();
		f.results = new sc.ImmutableTestSandboxCallResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}
}
