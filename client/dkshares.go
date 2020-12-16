// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package client

// This API is used to maintain the distributed key shares.
// The Golang API in this file tries to follow the REST conventions.

import (
	"net/http"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
)

// CreateDKSharesRoute is relative to the AdminRoutePrefix.
func DKSharesPostRoute() string {
	return "dks"
}

// GetDKSharesRoute is relative to the AdminRoutePrefix.
func DKSharesGetRoute(sharedAddress string) string {
	return "dks/" + sharedAddress
}

// DKSharesPostRequest is a POST request for creating new DKShare.
type DKSharesPostRequest struct {
	PeerNetIDs  []string `json:"peerNetIDs"`  // NetIDs of the nodes sharing the key.
	PeerPubKeys [][]byte `json:"peerPubKeys"` // Optional, base58 encoded public keys of the peers generating the DKS.
	Threshold   uint16   `json:"threshold"`   // Should be =< len(PeerPubKeys)
	TimeoutMS   uint16   `json:"timeoutMS"`   // Timeout in milliseconds.
}

// DKSharesInfo stands for the DKShare representation, returned by the GET and POST methods.
type DKSharesInfo struct {
	Address      string   `json:"address"`      // New generated shared address.
	SharedPubKey []byte   `json:"sharedPubKey"` // Shared public key.
	PubKeyShares [][]byte `json:"pubKeyShares"` // Public key shares for all the peers.
	Threshold    uint16   `json:"threshold"`    //
	PeerIndex    *uint16  `json:"peerIndex"`    // Index of the node returning the share, if it is a member of the sharing group.
}

// DKSharesPost creates new DKShare and returns its state.
func (c *WaspClient) DKSharesPost(request *DKSharesPostRequest) (*DKSharesInfo, error) {
	var response DKSharesInfo
	err := c.do(http.MethodPost, AdminRoutePrefix+"/"+DKSharesPostRoute(), request, &response)
	return &response, err
}

// DKSharesGet retrieves representation of an existing DKShare.
func (c *WaspClient) DKSharesGet(sharedAddress *address.Address) (*DKSharesInfo, error) {
	var sharedAddressStr = sharedAddress.String()
	var response DKSharesInfo
	err := c.do(http.MethodGet, AdminRoutePrefix+"/"+DKSharesGetRoute(sharedAddressStr), nil, &response)
	return &response, err
}