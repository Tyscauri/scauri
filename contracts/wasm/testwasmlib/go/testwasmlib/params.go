// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayAppendParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableArrayAppendParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableArrayAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayAppendParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableArrayAppendParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableArrayClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayClearParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableArrayClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayClearParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArraySetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArraySetParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableArraySetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableArraySetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArraySetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArraySetParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableArraySetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableMapClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapClearParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapClearParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapSetParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableMapSetParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableMapSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableMapSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapSetParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableMapSetParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableMapSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type MapStringToImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableBytes) GetBytes(key string) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableParamTypesParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableParamTypesParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutableParamTypesParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutableParamTypesParams) Bool() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamBool))
}

func (s ImmutableParamTypesParams) Bytes() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamBytes))
}

func (s ImmutableParamTypesParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutableParamTypesParams) Color() wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(s.proxy.Root(ParamColor))
}

func (s ImmutableParamTypesParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

func (s ImmutableParamTypesParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

func (s ImmutableParamTypesParams) Int16() wasmtypes.ScImmutableInt16 {
	return wasmtypes.NewScImmutableInt16(s.proxy.Root(ParamInt16))
}

func (s ImmutableParamTypesParams) Int32() wasmtypes.ScImmutableInt32 {
	return wasmtypes.NewScImmutableInt32(s.proxy.Root(ParamInt32))
}

func (s ImmutableParamTypesParams) Int64() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64))
}

func (s ImmutableParamTypesParams) Int8() wasmtypes.ScImmutableInt8 {
	return wasmtypes.NewScImmutableInt8(s.proxy.Root(ParamInt8))
}

func (s ImmutableParamTypesParams) Param() MapStringToImmutableBytes {
	//nolint:gosimple
	return MapStringToImmutableBytes{proxy: s.proxy}
}

func (s ImmutableParamTypesParams) RequestID() wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(s.proxy.Root(ParamRequestID))
}

func (s ImmutableParamTypesParams) String() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamString))
}

func (s ImmutableParamTypesParams) Uint16() wasmtypes.ScImmutableUint16 {
	return wasmtypes.NewScImmutableUint16(s.proxy.Root(ParamUint16))
}

func (s ImmutableParamTypesParams) Uint32() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamUint32))
}

func (s ImmutableParamTypesParams) Uint64() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamUint64))
}

func (s ImmutableParamTypesParams) Uint8() wasmtypes.ScImmutableUint8 {
	return wasmtypes.NewScImmutableUint8(s.proxy.Root(ParamUint8))
}

type MapStringToMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableBytes) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableBytes) GetBytes(key string) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableParamTypesParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableParamTypesParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutableParamTypesParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutableParamTypesParams) Bool() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamBool))
}

func (s MutableParamTypesParams) Bytes() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamBytes))
}

func (s MutableParamTypesParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutableParamTypesParams) Color() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(s.proxy.Root(ParamColor))
}

func (s MutableParamTypesParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

func (s MutableParamTypesParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}

func (s MutableParamTypesParams) Int16() wasmtypes.ScMutableInt16 {
	return wasmtypes.NewScMutableInt16(s.proxy.Root(ParamInt16))
}

func (s MutableParamTypesParams) Int32() wasmtypes.ScMutableInt32 {
	return wasmtypes.NewScMutableInt32(s.proxy.Root(ParamInt32))
}

func (s MutableParamTypesParams) Int64() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64))
}

func (s MutableParamTypesParams) Int8() wasmtypes.ScMutableInt8 {
	return wasmtypes.NewScMutableInt8(s.proxy.Root(ParamInt8))
}

func (s MutableParamTypesParams) Param() MapStringToMutableBytes {
	//nolint:gosimple
	return MapStringToMutableBytes{proxy: s.proxy}
}

func (s MutableParamTypesParams) RequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(s.proxy.Root(ParamRequestID))
}

func (s MutableParamTypesParams) String() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamString))
}

func (s MutableParamTypesParams) Uint16() wasmtypes.ScMutableUint16 {
	return wasmtypes.NewScMutableUint16(s.proxy.Root(ParamUint16))
}

func (s MutableParamTypesParams) Uint32() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamUint32))
}

func (s MutableParamTypesParams) Uint64() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamUint64))
}

func (s MutableParamTypesParams) Uint8() wasmtypes.ScMutableUint8 {
	return wasmtypes.NewScMutableUint8(s.proxy.Root(ParamUint8))
}

type ImmutableTriggerEventParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTriggerEventParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutableTriggerEventParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableTriggerEventParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTriggerEventParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutableTriggerEventParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableArrayLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayLengthParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableArrayLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayLengthParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayValueParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableArrayValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayValueParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableBlockRecordParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBlockRecordParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

func (s ImmutableBlockRecordParams) RecordIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamRecordIndex))
}

type MutableBlockRecordParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBlockRecordParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

func (s MutableBlockRecordParams) RecordIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamRecordIndex))
}

type ImmutableBlockRecordsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBlockRecordsParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableBlockRecordsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBlockRecordsParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapValueParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableMapValueParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapValueParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableMapValueParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}
