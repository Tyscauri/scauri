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
	owner:= ctx.Creator()

	    //create example compositions
	var composition1 *scauri.Composition
	composition1 = new(scauri.Composition) 
	composition1.Material = "PP"
	composition1.Mass = 80
	
	var composition2 *scauri.Composition
	composition2 = new(scauri.Composition) 
	composition2.Material = "PE"
	composition2.Mass = 10
	
	var composition3 *scauri.Composition
	composition3 = new(scauri.Composition) 
	composition3.Material = "HDPE"
	composition3.Mass = 10

	var compositionArray [3] *scauri.Composition

	compositionArray[0] = composition1
	compositionArray[1] = composition2
	compositionArray[2] = composition3
	
	createPP:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("Food")
	createPP.Params.RecyclateShare().SetValue(0)
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Params.PackagesNumber().SetValue(100000)
	createPP.Params.PackageWeight().SetValue(20000)
	createPP.Params.Compositions().AppendComposition().SetValue(composition1)
	createPP.Func.TransferIotas(10).Call()
	var id1 = createPP.Results.Id().Value();
	fmt.Println("id1: " + fmt.Sprint(id1))

	getMat1 := scauri.ScFuncs.GetMaterials(ctx)
	getMat1.Params.Id().SetValue(id1)
	getMat1.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)



    //test 10 productpasses and the storage of their material compositions
	var keys [10]wasmtypes.ScHash

	for k := 0; k < len(keys); k++ {

		createPPX:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
		createPPX.Params.Name().SetValue("tetrapack" + fmt.Sprint(k))
		createPPX.Params.Purpose().SetValue("Food")
		createPPX.Params.PackagesNumber().SetValue(100000)
		createPPX.Params.PackageWeight().SetValue(20000)
		createPPX.Params.RecyclateShare().SetValue(0)
		createPPX.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
		//add materials
		for i:= 0; i < len(compositionArray); i++ {
			createPPX.Params.Compositions().AppendComposition().SetValue(compositionArray[i])
			require.NoError(t, ctx.Err)

		}
		createPPX.Func.TransferIotas(10).Call()
		id:= createPPX.Results.Id().Value()
		keys[k] = id
		require.NoError(t, ctx.Err)

    }
	
/*	
	setMaterials:= scauri.ScFuncs.SetMaterials(ctx.Sign(ctx.NewSoloAgent()))
	for k:= 0; k < len(keys); k++ {

		setMaterials.Params.Id().SetValue(keys[k])

		for i:= 0; i < len(compositionArray); i++ {
			setMaterials.Params.Comp().AppendComposition().SetValue(compositionArray[i])
	
		}
		setMaterials.Func.TransferIotas(1).Call()
		require.NoError(t, ctx.Err)

	}
*/
   //Test Materials set
	for k:= 0; k < len(keys); k++ {
		getMatX := scauri.ScFuncs.GetMaterials(ctx)
		getMatX.Params.Id().SetValue(keys[k])
		getMatX.Func.Call()
		require.NoError(t, ctx.Err)
		require.EqualValues(t, "PP", getMatX.Results.Compositions().GetComposition(0).Value().Material)
		require.EqualValues(t, 80, getMatX.Results.Compositions().GetComposition(0).Value().Mass)
		require.EqualValues(t, "PE", getMatX.Results.Compositions().GetComposition(1).Value().Material)
		require.EqualValues(t, 10, getMatX.Results.Compositions().GetComposition(1).Value().Mass)
	}

	var end_time = time.Now().Unix()
	fmt.Println("total duration: " + fmt.Sprint(end_time - start_time))

	//fraction test
	var testSorter = ctx.NewSoloAgent()
	//var testSorterID:= testSorter.ScAgentID
	addSorter:= scauri.ScFuncs.AddSorter(ctx.Sign(owner))
	addSorter.Params.SorterID().SetValue(testSorter.ScAgentID())
	addSorter.Func.TransferIotas(1).Call()


	createFraction:= scauri.ScFuncs.CreateFraction(ctx.Sign(testSorter))
	createFraction.Params.Purpose().SetValue("Food")
	createFraction.Params.Name().SetValue("TestFraction")
	createFraction.Func.TransferIotas(1).Call()
	var fracID = createFraction.Results.FracID().Value()


    //add 15 packages of each type to a fraction
	for k:= 0; k < len(keys); k++ {

		for i:= 0; i < 15; i++ {
			addPPtoFractionX:= scauri.ScFuncs.AddPPToFraction(ctx)
			addPPtoFractionX.Params.PpID().SetValue(keys[k])
			addPPtoFractionX.Params.FracID().SetValue(fracID)
			addPPtoFractionX.Func.TransferIotas(1).Call()
		} 
	}

	/*
	var expectedFraction *scauri.Fraction
	expectedFraction.FracID = fracID
	expectedFraction.did = "tbd"
	expectedFraction.name = "TestFraction"
    expectedFraction.purpose = "Food"
	expectedFraction.issuer = testSorter.ScAgentID
	expectedFraction.amount = 1
*/


	var expectedFracCompositon1 *scauri.FracComposition
	expectedFracCompositon1 = new(scauri.FracComposition) 
	expectedFracCompositon1.Material = "PP"
	expectedFracCompositon1.Mass = 10 * 15 * 80

	GetFraction:= scauri.ScFuncs.GetFraction(ctx)
	GetFraction.Params.FracID().SetValue(fracID)
	GetFraction.Func.Call()
	fmt.Println("läuft")
	fmt.Println(GetFraction.Results.FracComposition().GetFracComposition(0).Value().Material)

	//require.True(t, GetFraction.Results.FracComposition().Exists())
	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(0).Exists())
	require.EqualValues(t, expectedFracCompositon1.Material, GetFraction.Results.FracComposition().GetFracComposition(0).Value().Material)
	require.EqualValues(t, expectedFracCompositon1.Mass, GetFraction.Results.FracComposition().GetFracComposition(0).Value().Mass)
	require.NoError(t, ctx.Err)


}
