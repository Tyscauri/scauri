package wasmtest

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/client/scclient"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/testutil"
	"github.com/iotaledger/wasp/packages/vm/vmconst"
	"testing"
	"time"
)

const (
	erc20_wasmPath    = "C:\\Users\\evaldas\\Documents\\proj\\Go\\src\\github.com\\iotaledger\\wasplib\\rust\\target\\wasm32-unknown-unknown\\release\\ERC20_ISCP.wasm"
	erc20_description = "ERC-20, a PoC smart contract"

	erc20_req_init_sc  = sctransaction.RequestCode(1 | sctransaction.RequestCodeProtected)
	erc20_req_transfer = sctransaction.RequestCode(2)
	erc20_req_approve  = sctransaction.RequestCode(3)

	erc20_var_supply = "s"
)

func TestDeploymentERC20(t *testing.T) {
	if !useWasm {
		t.Fatal("erc20 test is only for wasm SC code")
		return
	}
	wasps := setup(t, "TestDeploymentErc20")

	err := loadWasmIntoWasps(wasps, erc20_wasmPath, erc20_description)
	check(err, t)

	err = requestFunds(wasps, scOwnerAddr, "sc owner")
	check(err, t)

	scAddr, scColor, err := startSmartContract(wasps, "", erc20_description)
	checkSuccess(err, t, "smart contract has been created and activated")

	if !wasps.VerifyAddressBalances(scOwnerAddr, testutil.RequestFundsAmount-1, map[balance.Color]int64{
		balance.ColorIOTA: testutil.RequestFundsAmount - 1,
	}, "sc owner in the end") {
		t.Fail()
		return
	}

	if !wasps.VerifyAddressBalances(scAddr, 1, map[balance.Color]int64{
		*scColor: 1,
	}, "sc in the end") {
		t.Fail()
		return
	}

	if !wasps.VerifySCStateVariables2(scAddr, map[kv.Key]interface{}{
		vmconst.VarNameOwnerAddress: scOwnerAddr[:],
		vmconst.VarNameProgramHash:  programHash[:],
		vmconst.VarNameDescription:  erc20_description,
	}) {
		t.Fail()
	}
}

func TestInitERC20(t *testing.T) {
	if !useWasm {
		t.Fatal("erc20 test is only for wasm SC code")
		return
	}
	wasps := setup(t, "TestDeploymentErc20")

	err := loadWasmIntoWasps(wasps, erc20_wasmPath, erc20_description)
	check(err, t)

	err = requestFunds(wasps, scOwnerAddr, "sc owner")
	check(err, t)

	scAddr, scColor, err := startSmartContract(wasps, "", erc20_description)
	checkSuccess(err, t, "smart contract has been created and activated")

	client := scclient.New(
		wasps.NodeClient,
		wasps.WaspClient(0),
		scAddr,
		scOwner.SigScheme(),
		10*time.Second,
	)

	_, err = client.PostRequest(erc20_req_init_sc, nil, nil, map[string]interface{}{
		erc20_var_supply: 1000000,
	})
	checkSuccess(err, t, "success: posted initSC request")

	if !wasps.VerifyAddressBalances(scOwnerAddr, testutil.RequestFundsAmount-1, map[balance.Color]int64{
		balance.ColorIOTA: testutil.RequestFundsAmount - 1,
	}, "sc owner in the end") {
		t.Fail()
		return
	}

	if !wasps.VerifyAddressBalances(scAddr, 1, map[balance.Color]int64{
		*scColor: 1,
	}, "sc in the end") {
		t.Fail()
		return
	}

	if !wasps.VerifySCStateVariables2(scAddr, map[kv.Key]interface{}{
		vmconst.VarNameOwnerAddress: scOwnerAddr[:],
		vmconst.VarNameProgramHash:  programHash[:],
		vmconst.VarNameDescription:  erc20_description,
	}) {
		t.Fail()
	}
}