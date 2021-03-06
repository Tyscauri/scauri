// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";
import * as sc from "./index";

export class ImmutableDonateParams extends wasmtypes.ScProxy {
	feedback(): wasmtypes.ScImmutableString {
		return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamFeedback));
	}
}

export class MutableDonateParams extends wasmtypes.ScProxy {
	feedback(): wasmtypes.ScMutableString {
		return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamFeedback));
	}
}

export class ImmutableWithdrawParams extends wasmtypes.ScProxy {
	amount(): wasmtypes.ScImmutableUint64 {
		return new wasmtypes.ScImmutableUint64(this.proxy.root(sc.ParamAmount));
	}
}

export class MutableWithdrawParams extends wasmtypes.ScProxy {
	amount(): wasmtypes.ScMutableUint64 {
		return new wasmtypes.ScMutableUint64(this.proxy.root(sc.ParamAmount));
	}
}

export class ImmutableDonationParams extends wasmtypes.ScProxy {
	nr(): wasmtypes.ScImmutableUint32 {
		return new wasmtypes.ScImmutableUint32(this.proxy.root(sc.ParamNr));
	}
}

export class MutableDonationParams extends wasmtypes.ScProxy {
	nr(): wasmtypes.ScMutableUint32 {
		return new wasmtypes.ScMutableUint32(this.proxy.root(sc.ParamNr));
	}
}
