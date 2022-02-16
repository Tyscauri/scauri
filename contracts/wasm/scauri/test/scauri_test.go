// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"testing"
	"time"
	"fmt"

	"github.com/iotaledger/wasp/contracts/wasm/scauri/go/scauri"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
	"github.com/stretchr/testify/require"
	
)

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContext(t, scauri.ScName, scauri.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, scauri.ScName, scauri.OnLoad)
	require.NoError(t, ctx.ContractExists(scauri.ScName))
}


func TestCreateAndGetPP(t * testing.T) {


	//create single Productpass
	var start_time = time.Now().Unix()
	ctx:= setupTest(t)


	createPP:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("Food")
	createPP.Params.RecyclateShare().SetValue(0)
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Func.TransferIotas(10).Call()
	var id1 = createPP.Results.Id().Value();
	fmt.Println("id1: " + fmt.Sprint(id1))


    //create example compositions
	var composition1 *scauri.Composition
	composition1 = new(scauri.Composition) 
	composition1.Material = "PP"
	composition1.Proportion = 80

	var composition2 *scauri.Composition
	composition2 = new(scauri.Composition) 
	composition2.Material = "PE"
	composition2.Proportion = 10

	var composition3 *scauri.Composition
	composition3 = new(scauri.Composition) 
	composition3.Material = "HDPE"
	composition3.Proportion = 10

	var compositionArray [3] *scauri.Composition

	compositionArray[0] = composition1
	compositionArray[1] = composition2
	compositionArray[2] = composition3

	//Test single material setting and getting
	setMaterial:= scauri.ScFuncs.SetMaterials(ctx.Sign(ctx.NewSoloAgent()))
	setMaterial.Params.Id().SetValue(id1)
	setMaterial.Params.Comp().AppendComposition().SetValue(composition1)
	setMaterial.Func.TransferIotas(1).Call()

	getMat1 := scauri.ScFuncs.GetMaterials(ctx)
	getMat1.Params.Id().SetValue(id1)
	getMat1.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)

	/*
	createPP2:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
	createPP2.Params.Name().SetValue("Fish")
	createPP2.Params.Purpose().SetValue("Food")
	createPP2.Params.RecyclateShare().SetValue(0)
	createPP2.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP2.Func.TransferIotas(10).Call()
	var id2 = createPP.Results.Id().Value();
	fmt.Println("id2: " + fmt.Sprint(id2))
*/

    //test 10 productpasses and the storage of their material compositions
	var keys [10]wasmtypes.ScHash

	for i := 0; i < len(keys); i++ {

		createPPX:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
		createPPX.Params.Name().SetValue("tetrapack" + fmt.Sprint(i))
		createPPX.Params.Purpose().SetValue("Food")
		createPPX.Params.RecyclateShare().SetValue(0)
		createPPX.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
		createPPX.Func.TransferIotas(10).Call()
		id:= createPPX.Results.Id().Value()
		keys[i] = id
		require.NoError(t, ctx.Err)

    }
	
	
	setMaterials:= scauri.ScFuncs.SetMaterials(ctx.Sign(ctx.NewSoloAgent()))
	for k:= 0; k < len(keys); k++ {

		setMaterials.Params.Id().SetValue(keys[k])

		for i:= 0; i < len(compositionArray); i++ {
			setMaterials.Params.Comp().AppendComposition().SetValue(compositionArray[i])
	
		}
		setMaterials.Func.TransferIotas(1).Call()
		require.NoError(t, ctx.Err)

	}


	for k:= 0; k < len(keys); k++ {
		getMatX := scauri.ScFuncs.GetMaterials(ctx)
		getMatX.Params.Id().SetValue(keys[k])
		getMatX.Func.Call()
		require.NoError(t, ctx.Err)
		require.EqualValues(t, "PP", getMatX.Results.Compositions().GetComposition(0).Value().Material)
		require.EqualValues(t, 80, getMatX.Results.Compositions().GetComposition(0).Value().Proportion)
		require.EqualValues(t, "PE", getMatX.Results.Compositions().GetComposition(1).Value().Material)
		require.EqualValues(t, 10, getMatX.Results.Compositions().GetComposition(1).Value().Proportion)
	}

	var end_time = time.Now().Unix()
	fmt.Println("total duration: " + fmt.Sprint(end_time - start_time))

} 	
