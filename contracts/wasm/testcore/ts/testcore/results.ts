// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableCallOnChainResults extends wasmtypes.ScProxy {
	intValue(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ResultIntValue));
	}
}

export class MutableCallOnChainResults extends wasmtypes.ScProxy {
	intValue(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ResultIntValue));
	}
}

export class ImmutableGetMintedSupplyResults extends wasmtypes.ScProxy {
	mintedColor(): wasmtypes.ScImmutableColor {
		return new wasmtypes.ScImmutableColor(this.proxy.root(sc.ResultMintedColor));
	}

	mintedSupply(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ResultMintedSupply));
	}
}

export class MutableGetMintedSupplyResults extends wasmtypes.ScProxy {
	mintedColor(): wasmtypes.ScMutableColor {
		return new wasmtypes.ScMutableColor(this.proxy.root(sc.ResultMintedColor));
	}

	mintedSupply(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ResultMintedSupply));
	}
}

export class ImmutableRunRecursionResults extends wasmtypes.ScProxy {
	intValue(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ResultIntValue));
	}
}

export class MutableRunRecursionResults extends wasmtypes.ScProxy {
	intValue(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ResultIntValue));
	}
}

export class ImmutableTestChainOwnerIDFullResults extends wasmtypes.ScProxy {
	chainOwnerID(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ResultChainOwnerID));
	}
}

export class MutableTestChainOwnerIDFullResults extends wasmtypes.ScProxy {
	chainOwnerID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ResultChainOwnerID));
	}
}

export class ImmutableFibonacciResults extends wasmtypes.ScProxy {
	intValue(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ResultIntValue));
	}
}

export class MutableFibonacciResults extends wasmtypes.ScProxy {
	intValue(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ResultIntValue));
	}
}

export class ImmutableGetCounterResults extends wasmtypes.ScProxy {
	counter(): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.root(sc.ResultCounter));
	}
}

export class MutableGetCounterResults extends wasmtypes.ScProxy {
	counter(): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.root(sc.ResultCounter));
	}
}

export class MapStringToImmutableInt64 extends wasmtypes.ScProxy {

	getInt64(key: string): wasmtypes.ScImmutableInt64 {
		return new wasmtypes.ScImmutableInt64(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class ImmutableGetIntResults extends wasmtypes.ScProxy {
	values(): sc.MapStringToImmutableInt64 {
		return new sc.MapStringToImmutableInt64(this.proxy);
	}
}

export class MapStringToMutableInt64 extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getInt64(key: string): wasmtypes.ScMutableInt64 {
		return new wasmtypes.ScMutableInt64(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MutableGetIntResults extends wasmtypes.ScProxy {
	values(): sc.MapStringToMutableInt64 {
		return new sc.MapStringToMutableInt64(this.proxy);
	}
}

export class MapStringToImmutableString extends wasmtypes.ScProxy {

	getString(key: string): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class ImmutableGetStringValueResults extends wasmtypes.ScProxy {
	vars(): sc.MapStringToImmutableString {
		return new sc.MapStringToImmutableString(this.proxy);
	}
}

export class MapStringToMutableString extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getString(key: string): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.key(wasmtypes.stringToBytes(key)));
	}
}

export class MutableGetStringValueResults extends wasmtypes.ScProxy {
	vars(): sc.MapStringToMutableString {
		return new sc.MapStringToMutableString(this.proxy);
	}
}

export class ImmutableTestChainOwnerIDViewResults extends wasmtypes.ScProxy {
	chainOwnerID(): wasmtypes.ScImmutableAgentID {
		return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ResultChainOwnerID));
	}
}

export class MutableTestChainOwnerIDViewResults extends wasmtypes.ScProxy {
	chainOwnerID(): wasmtypes.ScMutableAgentID {
		return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ResultChainOwnerID));
	}
}

export class ImmutableTestSandboxCallResults extends wasmtypes.ScProxy {
	sandboxCall(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ResultSandboxCall));
	}
}

export class MutableTestSandboxCallResults extends wasmtypes.ScProxy {
	sandboxCall(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ResultSandboxCall));
	}
}
