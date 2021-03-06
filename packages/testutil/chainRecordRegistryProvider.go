// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package testutil

import (
	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/iotaledger/wasp/packages/registry"
)

// Mock implementation of a ChainRecordRegistryProvider for testing purposes

type ChainRecordRegistryProvider struct {
	DB map[[ledgerstate.AddressLength]byte]*registry.ChainRecord
}

func NewChainRecordRegistryProvider() *ChainRecordRegistryProvider {
	return &ChainRecordRegistryProvider{
		DB: map[[ledgerstate.AddressLength]byte]*registry.ChainRecord{},
	}
}

func (p *ChainRecordRegistryProvider) SaveChainRecord(chainRecord *registry.ChainRecord) error {
	p.DB[chainRecord.ChainID.Array()] = chainRecord
	return nil
}

func (p *ChainRecordRegistryProvider) LoadChainRecord(chainID *iscp.ChainID) (*registry.ChainRecord, error) {
	ret := p.DB[chainID.Array()]
	return ret, nil
}
