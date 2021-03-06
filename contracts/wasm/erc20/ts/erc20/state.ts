// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class MapAgentIDToImmutableAllowancesForAgent extends wasmtypes.ScProxy {

	getAllowancesForAgent(key: wasmtypes.ScAgentID): sc.ImmutableAllowancesForAgent {
		return new sc.ImmutableAllowancesForAgent(this.proxy.key(wasmtypes.agentIDToBytes(key)));
	}
}

export class ImmutableErc20State extends wasmtypes.ScProxy {
	allAllowances(): sc.MapAgentIDToImmutableAllowancesForAgent {
		return new sc.MapAgentIDToImmutableAllowancesForAgent(this.proxy.root(sc.StateAllAllowances));
	}

	balances(): sc.MapAgentIDToImmutableUint64 {
		return new sc.MapAgentIDToImmutableUint64(this.proxy.root(sc.StateBalances));
	}

	supply(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.StateSupply));
	}
}

export class MapAgentIDToMutableAllowancesForAgent extends wasmtypes.ScProxy {

	clear(): void {
		this.proxy.clearMap();
	}

	getAllowancesForAgent(key: wasmtypes.ScAgentID): sc.MutableAllowancesForAgent {
		return new sc.MutableAllowancesForAgent(this.proxy.key(wasmtypes.agentIDToBytes(key)));
	}
}

export class MutableErc20State extends wasmtypes.ScProxy {
	asImmutable(): sc.ImmutableErc20State {
		return new sc.ImmutableErc20State(this.proxy);
	}

	allAllowances(): sc.MapAgentIDToMutableAllowancesForAgent {
		return new sc.MapAgentIDToMutableAllowancesForAgent(this.proxy.root(sc.StateAllAllowances));
	}

	balances(): sc.MapAgentIDToMutableUint64 {
		return new sc.MapAgentIDToMutableUint64(this.proxy.root(sc.StateBalances));
	}

	supply(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.StateSupply));
	}
}
