// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

#[derive(Clone)]
pub struct Composition {
    pub mass     : u64, 
    pub material : String, 
}

impl Composition {
    pub fn from_bytes(bytes: &[u8]) -> Composition {
        let mut dec = WasmDecoder::new(bytes);
        Composition {
            mass     : uint64_decode(&mut dec),
            material : string_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.mass);
		string_encode(&mut enc, &self.material);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableComposition {
    pub(crate) proxy: Proxy,
}

impl ImmutableComposition {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Composition {
        Composition::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableComposition {
    pub(crate) proxy: Proxy,
}

impl MutableComposition {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Composition) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Composition {
        Composition::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct FracComposition {
    pub mass     : u64,  //in mg
    pub material : String, 
}

impl FracComposition {
    pub fn from_bytes(bytes: &[u8]) -> FracComposition {
        let mut dec = WasmDecoder::new(bytes);
        FracComposition {
            mass     : uint64_decode(&mut dec),
            material : string_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.mass);
		string_encode(&mut enc, &self.material);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableFracComposition {
    pub(crate) proxy: Proxy,
}

impl ImmutableFracComposition {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> FracComposition {
        FracComposition::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableFracComposition {
    pub(crate) proxy: Proxy,
}

impl MutableFracComposition {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &FracComposition) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> FracComposition {
        FracComposition::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct Fraction {
    pub amount  : u64, 
    pub did     : String, 
    pub frac_id : ScHash, 
    pub issuer  : ScAgentID, 
    pub name    : String, 
    pub pure    : bool, 
    pub purpose : String, 
}

impl Fraction {
    pub fn from_bytes(bytes: &[u8]) -> Fraction {
        let mut dec = WasmDecoder::new(bytes);
        Fraction {
            amount  : uint64_decode(&mut dec),
            did     : string_decode(&mut dec),
            frac_id : hash_decode(&mut dec),
            issuer  : agent_id_decode(&mut dec),
            name    : string_decode(&mut dec),
            pure    : bool_decode(&mut dec),
            purpose : string_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.amount);
		string_encode(&mut enc, &self.did);
		hash_encode(&mut enc, &self.frac_id);
		agent_id_encode(&mut enc, &self.issuer);
		string_encode(&mut enc, &self.name);
		bool_encode(&mut enc, self.pure);
		string_encode(&mut enc, &self.purpose);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableFraction {
    pub(crate) proxy: Proxy,
}

impl ImmutableFraction {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Fraction {
        Fraction::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableFraction {
    pub(crate) proxy: Proxy,
}

impl MutableFraction {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Fraction) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Fraction {
        Fraction::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct ProductPass {
    pub activation_date             : u64, 
    pub amount_per_charge           : u64, 
    pub charge_weight               : u64, 
    pub did                         : String,  //merged did:iota:id
    pub expiry_date                 : u64, 
    pub id                          : ScHash, 
    pub issuer                      : ScAgentID,  //packaging producer
    pub name                        : String, 
    pub package_weight              : u64, 
    pub packages_number             : u64, 
    pub packages_sorted             : u64, 
    pub packages_wrong_sorted       : u64, 
    pub purpose                     : String,  //e.g. food, hygiene, others
    pub remaining_amount_per_charge : u64, 
    pub reward_per_package_producer : u64, 
    pub reward_per_package_recycler : u64, 
    pub version                     : u8, 
}

impl ProductPass {
    pub fn from_bytes(bytes: &[u8]) -> ProductPass {
        let mut dec = WasmDecoder::new(bytes);
        ProductPass {
            activation_date             : uint64_decode(&mut dec),
            amount_per_charge           : uint64_decode(&mut dec),
            charge_weight               : uint64_decode(&mut dec),
            did                         : string_decode(&mut dec),
            expiry_date                 : uint64_decode(&mut dec),
            id                          : hash_decode(&mut dec),
            issuer                      : agent_id_decode(&mut dec),
            name                        : string_decode(&mut dec),
            package_weight              : uint64_decode(&mut dec),
            packages_number             : uint64_decode(&mut dec),
            packages_sorted             : uint64_decode(&mut dec),
            packages_wrong_sorted       : uint64_decode(&mut dec),
            purpose                     : string_decode(&mut dec),
            remaining_amount_per_charge : uint64_decode(&mut dec),
            reward_per_package_producer : uint64_decode(&mut dec),
            reward_per_package_recycler : uint64_decode(&mut dec),
            version                     : uint8_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.activation_date);
		uint64_encode(&mut enc, self.amount_per_charge);
		uint64_encode(&mut enc, self.charge_weight);
		string_encode(&mut enc, &self.did);
		uint64_encode(&mut enc, self.expiry_date);
		hash_encode(&mut enc, &self.id);
		agent_id_encode(&mut enc, &self.issuer);
		string_encode(&mut enc, &self.name);
		uint64_encode(&mut enc, self.package_weight);
		uint64_encode(&mut enc, self.packages_number);
		uint64_encode(&mut enc, self.packages_sorted);
		uint64_encode(&mut enc, self.packages_wrong_sorted);
		string_encode(&mut enc, &self.purpose);
		uint64_encode(&mut enc, self.remaining_amount_per_charge);
		uint64_encode(&mut enc, self.reward_per_package_producer);
		uint64_encode(&mut enc, self.reward_per_package_recycler);
		uint8_encode(&mut enc, self.version);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableProductPass {
    pub(crate) proxy: Proxy,
}

impl ImmutableProductPass {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> ProductPass {
        ProductPass::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableProductPass {
    pub(crate) proxy: Proxy,
}

impl MutableProductPass {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &ProductPass) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> ProductPass {
        ProductPass::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct RecyComposition {
    pub mass     : u64,  //in mg
    pub material : String, 
}

impl RecyComposition {
    pub fn from_bytes(bytes: &[u8]) -> RecyComposition {
        let mut dec = WasmDecoder::new(bytes);
        RecyComposition {
            mass     : uint64_decode(&mut dec),
            material : string_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		uint64_encode(&mut enc, self.mass);
		string_encode(&mut enc, &self.material);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableRecyComposition {
    pub(crate) proxy: Proxy,
}

impl ImmutableRecyComposition {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> RecyComposition {
        RecyComposition::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableRecyComposition {
    pub(crate) proxy: Proxy,
}

impl MutableRecyComposition {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &RecyComposition) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> RecyComposition {
        RecyComposition::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct Recyclate {
    pub did     : String, 
    pub frac_id : ScHash, 
    pub issuer  : ScAgentID, 
    pub name    : String, 
    pub pure    : bool, 
    pub purpose : String, 
    pub recy_id : ScHash, 
}

impl Recyclate {
    pub fn from_bytes(bytes: &[u8]) -> Recyclate {
        let mut dec = WasmDecoder::new(bytes);
        Recyclate {
            did     : string_decode(&mut dec),
            frac_id : hash_decode(&mut dec),
            issuer  : agent_id_decode(&mut dec),
            name    : string_decode(&mut dec),
            pure    : bool_decode(&mut dec),
            purpose : string_decode(&mut dec),
            recy_id : hash_decode(&mut dec),
        }
    }

    pub fn to_bytes(&self) -> Vec<u8> {
        let mut enc = WasmEncoder::new();
		string_encode(&mut enc, &self.did);
		hash_encode(&mut enc, &self.frac_id);
		agent_id_encode(&mut enc, &self.issuer);
		string_encode(&mut enc, &self.name);
		bool_encode(&mut enc, self.pure);
		string_encode(&mut enc, &self.purpose);
		hash_encode(&mut enc, &self.recy_id);
        enc.buf()
    }
}

#[derive(Clone)]
pub struct ImmutableRecyclate {
    pub(crate) proxy: Proxy,
}

impl ImmutableRecyclate {
    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn value(&self) -> Recyclate {
        Recyclate::from_bytes(&self.proxy.get())
    }
}

#[derive(Clone)]
pub struct MutableRecyclate {
    pub(crate) proxy: Proxy,
}

impl MutableRecyclate {
    pub fn delete(&self) {
        self.proxy.delete();
    }

    pub fn exists(&self) -> bool {
        self.proxy.exists()
    }

    pub fn set_value(&self, value: &Recyclate) {
        self.proxy.set(&value.to_bytes());
    }

    pub fn value(&self) -> Recyclate {
        Recyclate::from_bytes(&self.proxy.get())
    }
}
