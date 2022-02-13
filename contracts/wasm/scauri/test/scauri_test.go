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

	var start_time = time.Now().Unix()
	ctx:= setupTest(t)


	createPP:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("Food")
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Func.TransferIotas(10).Post()
	var id1 = createPP.Results.Id().Value();


	var keys [10]wasmtypes.ScHash

	for i := 0; i < len(keys); i++ {

		createPP:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
		createPP.Params.Name().SetValue("tetrapack" + fmt.Sprint(i))
		createPP.Params.Purpose().SetValue("Food")
		createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
		createPP.Func.TransferIotas(10).Post()
		id:= createPP.Results.Id().Value()
		keys[i] = id
		fmt.Println("id: " + fmt.Sprint(id));
    }

	var end_time = time.Now().Unix()
	fmt.Println("total duration: " + fmt.Sprint(end_time - start_time))

	require.NoError(t, ctx.Err)
	
	//var exampleComposition[3]scauri.Composition



	var composition *scauri.Composition
	composition = new(scauri.Composition) 
	composition.Material = "PP"
	composition.Proportion = 5

	setMaterial:= scauri.ScFuncs.SetMaterials(ctx.Sign(ctx.NewSoloAgent()))
	setMaterial.Params.Id().SetValue(id1)
	setMaterial.Params.Comp().AppendComposition().SetValue(composition)
	setMaterial.Func.Call()

	
	setMaterials:= scauri.ScFuncs.SetMaterials(ctx.Sign(ctx.NewSoloAgent()))
	for k:= 0; k < len(keys); k++ {
		setMaterials.Params.Id().SetValue(keys[k])

		for i:= 0; i < 3; i++ {
			setMaterials.Params.Comp().AppendComposition().SetValue(composition)
	
		}
	}
	setMaterials.Func.Call()

	for k:= 0; k < len(keys); k++ {
		getMat1 := scauri.ScFuncs.GetMaterials(ctx)
		getMat1.Params.Id().SetValue(keys[k])
		getMat1.Func.Call()
		require.NoError(t, ctx.Err)
		require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)
		require.EqualValues(t, 5, getMat1.Results.Compositions().GetComposition(0).Value().Proportion)
		//require.EqualValues(t, "PE", getMat1.Results.Compositions().GetComposition(1).Value().Material)
		//require.EqualValues(t, 40, getMat1.Results.Compositions().GetComposition(1).Value().Proportion)
	}


	/*
	addMaterial:=scauri.ScFuncs.AddMaterial(ctx.Sign(ctx.NewSoloAgent()))
	addMaterial.Params.Id().SetValue("123")
	addMaterial.Params.Mat().SetValue("PP")
	addMaterial.Params.Prop().SetValue(60)
	addMaterial.Func.TransferIotas(1).Post()
	
	addMaterial2:=scauri.ScFuncs.AddMaterial(ctx.Sign(ctx.NewSoloAgent()))
	addMaterial2.Params.Id().SetValue("123")
	addMaterial2.Params.Mat().SetValue("PE")
	addMaterial2.Params.Prop().SetValue(40)
	addMaterial2.Func.TransferIotas(1).Post()
			
	createPP2:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent()))
	createPP2.Params.Name().SetValue("wurscht")
	createPP2.Params.Id().SetValue("1234")
	createPP2.Func.TransferIotas(1).Post()
	
	//test getPP
	getPP := scauri.ScFuncs.GetPP(ctx)				//for views
	getPP.Params.Id().SetValue("123")                            // for views
	getPP.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "Chips", getPP.Results.Ppname().Value())
	
	getPP2 := scauri.ScFuncs.GetPP(ctx)			
	getPP2.Params.Id().SetValue("1234")
	getPP2.Func.Call()
	require.NoError(t, ctx.Err)
	require.NotEqualValues(t, "Chips", getPP.Results.Ppname().Value())
	require.NotEqualValues(t, "Chips", getPP.Results.Ppresult().Value().Name)
	
	getMat1 := scauri.ScFuncs.GetMaterials(ctx)
	getMat1.Params.Id().SetValue("123")
	getMat1.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)
	require.EqualValues(t, 60, getMat1.Results.Compositions().GetComposition(0).Value().Proportion)
	require.EqualValues(t, "PE", getMat1.Results.Compositions().GetComposition(1).Value().Material)
	require.EqualValues(t, 40, getMat1.Results.Compositions().GetComposition(1).Value().Proportion)
	*/
} 	
