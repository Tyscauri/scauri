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
pub struct ImmutableAddPPToFractionParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableAddPPToFractionParams {
    pub fn frac_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}

    pub fn pp_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_PP_ID))
	}
}

#[derive(Clone)]
pub struct MutableAddPPToFractionParams {
	pub(crate) proxy: Proxy,
}

impl MutableAddPPToFractionParams {
    pub fn frac_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}

    pub fn pp_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_PP_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableAddRecyclerParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableAddRecyclerParams {
    pub fn recycler_id(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(PARAM_RECYCLER_ID))
	}
}

#[derive(Clone)]
pub struct MutableAddRecyclerParams {
	pub(crate) proxy: Proxy,
}

impl MutableAddRecyclerParams {
    pub fn recycler_id(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(PARAM_RECYCLER_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableAddSorterParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableAddSorterParams {
    pub fn sorter_id(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(PARAM_SORTER_ID))
	}
}

#[derive(Clone)]
pub struct MutableAddSorterParams {
	pub(crate) proxy: Proxy,
}

impl MutableAddSorterParams {
    pub fn sorter_id(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(PARAM_SORTER_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableCreateFractionParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableCreateFractionParams {
    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn purpose(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_PURPOSE))
	}
}

#[derive(Clone)]
pub struct MutableCreateFractionParams {
	pub(crate) proxy: Proxy,
}

impl MutableCreateFractionParams {
    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn purpose(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_PURPOSE))
	}
}

#[derive(Clone)]
pub struct ImmutableCreatePPParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableCreatePPParams {
    pub fn compositions(&self) -> ImmutableCompositions {
		ImmutableCompositions { proxy: self.proxy.root(PARAM_COMPOSITIONS) }
	}

    pub fn expiry_date(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_EXPIRY_DATE))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn package_weight(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_PACKAGE_WEIGHT))
	}

    pub fn packages_number(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_PACKAGES_NUMBER))
	}

    pub fn purpose(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_PURPOSE))
	}

    pub fn recyclate_share(&self) -> ScImmutableUint8 {
		ScImmutableUint8::new(self.proxy.root(PARAM_RECYCLATE_SHARE))
	}
}

#[derive(Clone)]
pub struct MutableCreatePPParams {
	pub(crate) proxy: Proxy,
}

impl MutableCreatePPParams {
    pub fn compositions(&self) -> MutableCompositions {
		MutableCompositions { proxy: self.proxy.root(PARAM_COMPOSITIONS) }
	}

    pub fn expiry_date(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_EXPIRY_DATE))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}

    pub fn package_weight(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_PACKAGE_WEIGHT))
	}

    pub fn packages_number(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_PACKAGES_NUMBER))
	}

    pub fn purpose(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_PURPOSE))
	}

    pub fn recyclate_share(&self) -> ScMutableUint8 {
		ScMutableUint8::new(self.proxy.root(PARAM_RECYCLATE_SHARE))
	}
}

#[derive(Clone)]
pub struct ImmutableCreateRecyclateParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableCreateRecyclateParams {
    pub fn frac_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}

    pub fn name(&self) -> ScImmutableString {
		ScImmutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct MutableCreateRecyclateParams {
	pub(crate) proxy: Proxy,
}

impl MutableCreateRecyclateParams {
    pub fn frac_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}

    pub fn name(&self) -> ScMutableString {
		ScMutableString::new(self.proxy.root(PARAM_NAME))
	}
}

#[derive(Clone)]
pub struct ImmutableDeletePPParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableDeletePPParams {
    pub fn pp_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_PP_ID))
	}
}

#[derive(Clone)]
pub struct MutableDeletePPParams {
	pub(crate) proxy: Proxy,
}

impl MutableDeletePPParams {
    pub fn pp_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_PP_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableInitParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableInitParams {
    pub fn owner(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(PARAM_OWNER))
	}
}

#[derive(Clone)]
pub struct MutableInitParams {
	pub(crate) proxy: Proxy,
}

impl MutableInitParams {
    pub fn owner(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(PARAM_OWNER))
	}
}

#[derive(Clone)]
pub struct ImmutablePayoutProducerParams {
	pub(crate) proxy: Proxy,
}

impl ImmutablePayoutProducerParams {
    pub fn frac_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}
}

#[derive(Clone)]
pub struct MutablePayoutProducerParams {
	pub(crate) proxy: Proxy,
}

impl MutablePayoutProducerParams {
    pub fn frac_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableSetOwnerParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableSetOwnerParams {
    pub fn owner(&self) -> ScImmutableAgentID {
		ScImmutableAgentID::new(self.proxy.root(PARAM_OWNER))
	}
}

#[derive(Clone)]
pub struct MutableSetOwnerParams {
	pub(crate) proxy: Proxy,
}

impl MutableSetOwnerParams {
    pub fn owner(&self) -> ScMutableAgentID {
		ScMutableAgentID::new(self.proxy.root(PARAM_OWNER))
	}
}

#[derive(Clone)]
pub struct ImmutableGetAmountOfRequiredFundsParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetAmountOfRequiredFundsParams {
    pub fn charge_weight(&self) -> ScImmutableUint64 {
		ScImmutableUint64::new(self.proxy.root(PARAM_CHARGE_WEIGHT))
	}
}

#[derive(Clone)]
pub struct MutableGetAmountOfRequiredFundsParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetAmountOfRequiredFundsParams {
    pub fn charge_weight(&self) -> ScMutableUint64 {
		ScMutableUint64::new(self.proxy.root(PARAM_CHARGE_WEIGHT))
	}
}

#[derive(Clone)]
pub struct ImmutableGetFractionParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetFractionParams {
    pub fn frac_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}
}

#[derive(Clone)]
pub struct MutableGetFractionParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetFractionParams {
    pub fn frac_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_FRAC_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableGetMaterialsParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetMaterialsParams {
    pub fn id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_ID))
	}
}

#[derive(Clone)]
pub struct MutableGetMaterialsParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetMaterialsParams {
    pub fn id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableGetPPParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetPPParams {
    pub fn id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_ID))
	}
}

#[derive(Clone)]
pub struct MutableGetPPParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetPPParams {
    pub fn id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableGetRecyclateParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetRecyclateParams {
    pub fn recy_id(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(PARAM_RECY_ID))
	}
}

#[derive(Clone)]
pub struct MutableGetRecyclateParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetRecyclateParams {
    pub fn recy_id(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(PARAM_RECY_ID))
	}
}

#[derive(Clone)]
pub struct ImmutableGetTokenPerPackageParams {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetTokenPerPackageParams {
    pub fn prod_pass(&self) -> ImmutableProductPass {
		ImmutableProductPass { proxy: self.proxy.root(PARAM_PROD_PASS) }
	}
}

#[derive(Clone)]
pub struct MutableGetTokenPerPackageParams {
	pub(crate) proxy: Proxy,
}

impl MutableGetTokenPerPackageParams {
    pub fn prod_pass(&self) -> MutableProductPass {
		MutableProductPass { proxy: self.proxy.root(PARAM_PROD_PASS) }
	}
}
