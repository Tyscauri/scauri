// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"

	"contracts/wasm/scauri/go/scauri"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, scauri.ScName, scauri.OnLoad)
	require.NoError(t, ctx.ContractExists(scauri.ScName))
}
