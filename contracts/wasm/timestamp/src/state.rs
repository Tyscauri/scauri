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
pub struct ImmutabletimestampState {
	pub(crate) proxy: Proxy,
}

impl ImmutabletimestampState {
    pub fn timestamp(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(STATE_TIMESTAMP))
	}
}

#[derive(Clone)]
pub struct MutabletimestampState {
	pub(crate) proxy: Proxy,
}

impl MutabletimestampState {
    pub fn as_immutable(&self) -> ImmutabletimestampState {
		ImmutabletimestampState { proxy: self.proxy.root("") }
	}

    pub fn timestamp(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(STATE_TIMESTAMP))
	}
}
