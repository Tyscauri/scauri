// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

use crate::*;

#[derive(Clone)]
pub struct MapStringToImmutableStringArray {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableStringArray {
    pub fn get_string_array(&self, key: &str) -> ImmutableStringArray {
        ImmutableStringArray { proxy: self.proxy.key(&string_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapStringToImmutableStringMap {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableStringMap {
    pub fn get_string_map(&self, key: &str) -> ImmutableStringMap {
        ImmutableStringMap { proxy: self.proxy.key(&string_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct ImmutableTestWasmLibState {
	pub(crate) proxy: Proxy,
}

impl ImmutableTestWasmLibState {
    pub fn arrays(&self) -> MapStringToImmutableStringArray {
		MapStringToImmutableStringArray { proxy: self.proxy.root(STATE_ARRAYS) }
	}

    pub fn maps(&self) -> MapStringToImmutableStringMap {
		MapStringToImmutableStringMap { proxy: self.proxy.root(STATE_MAPS) }
	}

    pub fn random(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(STATE_RANDOM))
	}
}

#[derive(Clone)]
pub struct MapStringToMutableStringArray {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableStringArray {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_string_array(&self, key: &str) -> MutableStringArray {
        MutableStringArray { proxy: self.proxy.key(&string_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MapStringToMutableStringMap {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableStringMap {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_string_map(&self, key: &str) -> MutableStringMap {
        MutableStringMap { proxy: self.proxy.key(&string_to_bytes(key)) }
    }
}

#[derive(Clone)]
pub struct MutableTestWasmLibState {
	pub(crate) proxy: Proxy,
}

impl MutableTestWasmLibState {
    pub fn as_immutable(&self) -> ImmutableTestWasmLibState {
		ImmutableTestWasmLibState { proxy: self.proxy.root("") }
	}

    pub fn arrays(&self) -> MapStringToMutableStringArray {
		MapStringToMutableStringArray { proxy: self.proxy.root(STATE_ARRAYS) }
	}

    pub fn maps(&self) -> MapStringToMutableStringMap {
		MapStringToMutableStringMap { proxy: self.proxy.root(STATE_MAPS) }
	}

    pub fn random(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(STATE_RANDOM))
	}
}
