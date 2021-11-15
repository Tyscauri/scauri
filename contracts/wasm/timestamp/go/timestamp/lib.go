// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package timestamp

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncNow, funcNowThunk)
	exports.AddView(ViewGetTimestamp, viewGetTimestampThunk)

	for i, key := range keyMap {
		idxMap[i] = key.KeyID()
	}
}

type NowContext struct {
	State MutabletimestampState
}

func funcNowThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("timestamp.funcNow")
	f := &NowContext{
		State: MutabletimestampState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	funcNow(ctx, f)
	ctx.Log("timestamp.funcNow ok")
}

type GetTimestampContext struct {
	Results MutableGetTimestampResults
	State   ImmutabletimestampState
}

func viewGetTimestampThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("timestamp.viewGetTimestamp")
	f := &GetTimestampContext{
		Results: MutableGetTimestampResults{
			id: wasmlib.OBJ_ID_RESULTS,
		},
		State: ImmutabletimestampState{
			id: wasmlib.OBJ_ID_STATE,
		},
	}
	viewGetTimestamp(ctx, f)
	ctx.Log("timestamp.viewGetTimestamp ok")
}