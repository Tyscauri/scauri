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


func TestSuccessfullRecyclingCircle(t * testing.T) {

	ctx:= setupTest(t)
	owner:= ctx.Creator()

	//create example compositions
	var composition1 *scauri.Composition
	composition1 = new(scauri.Composition) 
	composition1.Material = "PP"
	var ppMass uint64 = 8000 // mass of PP in mg
	composition1.Mass = ppMass
	
	var composition2 *scauri.Composition
	composition2 = new(scauri.Composition) 
	composition2.Material = "PE"
	var peMass uint64 = 2000
	composition2.Mass = peMass
	
	var composition3 *scauri.Composition
	composition3 = new(scauri.Composition)
	var hdpeMass uint64 = 2000
	composition3.Material = "HDPE"
	composition3.Mass = hdpeMass

	var compositionArray [3] *scauri.Composition

	compositionArray[0] = composition1
	compositionArray[1] = composition2
	compositionArray[2] = composition3
	
	//create single Productpass
	createPP:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("Food")
	createPP.Params.RecyclateShare().SetValue(0)
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Params.PackagesNumber().SetValue(100000)
	createPP.Params.PackageWeight().SetValue(12000)
	createPP.Params.Compositions().AppendComposition().SetValue(composition1)
	createPP.Params.Compositions().AppendComposition().SetValue(composition2)
	createPP.Params.Compositions().AppendComposition().SetValue(composition3)
	createPP.Func.TransferIotas(10).Call()
	var id1 = createPP.Results.Id().Value();
	fmt.Println("id1: " + fmt.Sprint(id1))

	getMat1 := scauri.ScFuncs.GetMaterials(ctx)
	getMat1.Params.Id().SetValue(id1)
	getMat1.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)



    // create multiple (numPackageTypes) test productpasses and the storage of their material compositions
	const numPackageTypes uint64 = 10
	var keys [numPackageTypes]wasmtypes.ScHash

	for k := 0; k < len(keys); k++ {

		createPPX:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
		createPPX.Params.Name().SetValue("tetrapack" + fmt.Sprint(k))
		createPPX.Params.Purpose().SetValue("Food")
		createPPX.Params.PackagesNumber().SetValue(100000)
		createPPX.Params.PackageWeight().SetValue(12000)
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
	

   //Test Materials set
	for k:= 0; k < len(keys); k++ {
		getMatX := scauri.ScFuncs.GetMaterials(ctx)
		getMatX.Params.Id().SetValue(keys[k])
		getMatX.Func.Call()
		require.NoError(t, ctx.Err)
		require.EqualValues(t, "PP", getMatX.Results.Compositions().GetComposition(0).Value().Material)
		require.EqualValues(t, ppMass, getMatX.Results.Compositions().GetComposition(0).Value().Mass)
		require.EqualValues(t, "PE", getMatX.Results.Compositions().GetComposition(1).Value().Material)
		require.EqualValues(t, peMass, getMatX.Results.Compositions().GetComposition(1).Value().Mass)
		require.EqualValues(t, "HDPE", getMatX.Results.Compositions().GetComposition(2).Value().Material)
		require.EqualValues(t, hdpeMass, getMatX.Results.Compositions().GetComposition(2).Value().Mass)
	}

	//fraction test
	var testSorter = ctx.NewSoloAgent()
	addSorter:= scauri.ScFuncs.AddSorter(ctx.Sign(owner))
	addSorter.Params.SorterID().SetValue(testSorter.ScAgentID())
	addSorter.Func.TransferIotas(1).Call()


	createFraction:= scauri.ScFuncs.CreateFraction(ctx.Sign(testSorter))
	createFraction.Params.Purpose().SetValue("Food")
	createFraction.Params.Name().SetValue("TestFraction")
	createFraction.Func.TransferIotas(1).Call()
	var fracID = createFraction.Results.FracID().Value()


    //add 15 packages of each type to a fraction
	var numPackages uint64 = 15

	for k:= 0; k < len(keys); k++ {

		for i:= 0; i < int(numPackages); i++ {
			addPPtoFractionX:= scauri.ScFuncs.AddPPToFraction(ctx)
			addPPtoFractionX.Params.PpID().SetValue(keys[k])
			addPPtoFractionX.Params.FracID().SetValue(fracID)
			addPPtoFractionX.Func.TransferIotas(1).Call()
		} 
	}

	
	var expectedFraction *scauri.Fraction
	expectedFraction = new(scauri.Fraction)
	expectedFraction.FracId = fracID
	expectedFraction.Did = "tbd"
	expectedFraction.Name = "TestFraction"
    expectedFraction.Purpose = "Food"
	expectedFraction.Issuer = testSorter.ScAgentID()
	expectedFraction.Amount = 1

	GetFraction:= scauri.ScFuncs.GetFraction(ctx)
	GetFraction.Params.FracID().SetValue(fracID)
	GetFraction.Func.Call()
	var fracResultsProxy = GetFraction.Results

	//Test fraction properties
	require.EqualValues(t, expectedFraction.FracId, fracResultsProxy.Fraction().Value().FracId)
	require.EqualValues(t, expectedFraction.Purpose, fracResultsProxy.Fraction().Value().Purpose)
	require.EqualValues(t, expectedFraction.Issuer, fracResultsProxy.Fraction().Value().Issuer)

    //Test resulting fraction compositions (for PP, PE and HDPE)
	var expectedFracCompositon1 *scauri.FracComposition
	expectedFracCompositon1 = new(scauri.FracComposition) 
	expectedFracCompositon1.Material = "PP"
	expectedFracCompositon1.Mass = numPackageTypes * numPackages * ppMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(0).Exists())
	require.EqualValues(t, expectedFracCompositon1.Material, fracResultsProxy.FracComposition().GetFracComposition(0).Value().Material)
	require.EqualValues(t, expectedFracCompositon1.Mass, fracResultsProxy.FracComposition().GetFracComposition(0).Value().Mass)
	require.NoError(t, ctx.Err)

	var expectedFracCompositon2 *scauri.FracComposition
	expectedFracCompositon2 = new(scauri.FracComposition) 
	expectedFracCompositon2.Material = "PE"
	expectedFracCompositon2.Mass = numPackageTypes * numPackages * peMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(1).Exists())
	require.EqualValues(t, expectedFracCompositon2.Material, fracResultsProxy.FracComposition().GetFracComposition(1).Value().Material)
	require.EqualValues(t, expectedFracCompositon2.Mass, fracResultsProxy.FracComposition().GetFracComposition(1).Value().Mass)
	require.NoError(t, ctx.Err)
	
	var expectedFracCompositon3 *scauri.FracComposition
	expectedFracCompositon3 = new(scauri.FracComposition) 
	expectedFracCompositon3.Material = "HDPE"
	expectedFracCompositon3.Mass = numPackageTypes * numPackages * hdpeMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(2).Exists())
	require.EqualValues(t, expectedFracCompositon3.Material, fracResultsProxy.FracComposition().GetFracComposition(2).Value().Material)
	require.EqualValues(t, expectedFracCompositon3.Mass, fracResultsProxy.FracComposition().GetFracComposition(2).Value().Mass)
	require.NoError(t, ctx.Err)

	var testRecycler = ctx.NewSoloAgent()
	addRecycler:= scauri.ScFuncs.AddRecycler(ctx.Sign(owner))
	addRecycler.Params.RecyclerID().SetValue(testRecycler.ScAgentID())
	addRecycler.Func.TransferIotas(1).Call()

	//test recyclate
	createRecyclate:= scauri.ScFuncs.CreateRecyclate(ctx.Sign(testRecycler))
	createRecyclate.Params.FracID().SetValue(fracID)
	createRecyclate.Params.Name().SetValue("TestRecyclate")
	createRecyclate.Func.TransferIotas(1).Call()
	var recyID = createRecyclate.Results.RecyID().Value()

	var expectedRecyclate *scauri.Recyclate
	expectedRecyclate = new(scauri.Recyclate)
	expectedRecyclate.RecyId = recyID
	expectedRecyclate.FracId = fracID
	expectedRecyclate.Did = "tbd"
	expectedRecyclate.Name = "TestRecyclate"
    expectedRecyclate.Purpose = "Food"
	expectedRecyclate.Issuer = testRecycler.ScAgentID()
	expectedRecyclate.Amount = 1

	GetRecyclate:= scauri.ScFuncs.GetRecyclate(ctx)
	GetRecyclate.Params.RecyID().SetValue(recyID)
	GetRecyclate.Func.Call()
	var recyResultsProxy = GetRecyclate.Results

	require.EqualValues(t, expectedRecyclate.RecyId, recyResultsProxy.Recyclate().Value().RecyId)
	require.EqualValues(t, expectedRecyclate.FracId, recyResultsProxy.Recyclate().Value().FracId)
	require.EqualValues(t, expectedRecyclate.Purpose, recyResultsProxy.Recyclate().Value().Purpose)
	require.EqualValues(t, expectedRecyclate.Issuer, recyResultsProxy.Recyclate().Value().Issuer)


	//Test resulting recycling compositions (for PP, PE and HDPE). Should be the same as for the fraction
	require.True(t, GetRecyclate.Results.RecyComposition().GetRecyComposition(0).Exists())
	require.EqualValues(t, expectedFracCompositon1.Material, recyResultsProxy.RecyComposition().GetRecyComposition(0).Value().Material)
	require.EqualValues(t, expectedFracCompositon1.Mass, recyResultsProxy.RecyComposition().GetRecyComposition(0).Value().Mass)
	require.NoError(t, ctx.Err)

	require.True(t, GetRecyclate.Results.RecyComposition().GetRecyComposition(1).Exists())
	require.EqualValues(t, expectedFracCompositon2.Material, recyResultsProxy.RecyComposition().GetRecyComposition(1).Value().Material)
	require.EqualValues(t, expectedFracCompositon2.Mass, recyResultsProxy.RecyComposition().GetRecyComposition(1).Value().Mass)
	require.NoError(t, ctx.Err)

	require.True(t, GetRecyclate.Results.RecyComposition().GetRecyComposition(2).Exists())
	require.EqualValues(t, expectedFracCompositon3.Material, recyResultsProxy.RecyComposition().GetRecyComposition(2).Value().Material)
	require.EqualValues(t, expectedFracCompositon3.Mass, recyResultsProxy.RecyComposition().GetRecyComposition(2).Value().Mass)
	require.NoError(t, ctx.Err)
	//test money distribution
	//test deletion
	//test access
}


func TestUnsuccessfullRecyclingCircle(t * testing.T) {



	ctx:= setupTest(t)
	owner:= ctx.Creator()

	//create example compositions
	var composition1 *scauri.Composition
	composition1 = new(scauri.Composition) 
	composition1.Material = "PP"
	var ppMass uint64 = 8000 // mass of PP in mg
	composition1.Mass = ppMass
	
	var composition2 *scauri.Composition
	composition2 = new(scauri.Composition) 
	composition2.Material = "PE"
	var peMass uint64 = 2000
	composition2.Mass = peMass
	
	var composition3 *scauri.Composition
	composition3 = new(scauri.Composition)
	var hdpeMass uint64 = 2000
	composition3.Material = "HDPE"
	composition3.Mass = hdpeMass

	var compositionArray [3] *scauri.Composition

	compositionArray[0] = composition1
	compositionArray[1] = composition2
	compositionArray[2] = composition3
	

    // create multiple (numPackageTypes) test productpasses and the storage of their material compositions
	const numPackageTypes uint64 = 10
	var keys [numPackageTypes]wasmtypes.ScHash

	for k := 0; k < len(keys); k++ {

		createPPX:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
		createPPX.Params.Name().SetValue("tetrapack" + fmt.Sprint(k))
		createPPX.Params.Purpose().SetValue("Food")
		createPPX.Params.PackagesNumber().SetValue(100000)
		createPPX.Params.PackageWeight().SetValue(12000)
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
	

   //Test Materials set
	for k:= 0; k < len(keys); k++ {
		getMatX := scauri.ScFuncs.GetMaterials(ctx)
		getMatX.Params.Id().SetValue(keys[k])
		getMatX.Func.Call()
		require.NoError(t, ctx.Err)
		require.EqualValues(t, "PP", getMatX.Results.Compositions().GetComposition(0).Value().Material)
		require.EqualValues(t, ppMass, getMatX.Results.Compositions().GetComposition(0).Value().Mass)
		require.EqualValues(t, "PE", getMatX.Results.Compositions().GetComposition(1).Value().Material)
		require.EqualValues(t, peMass, getMatX.Results.Compositions().GetComposition(1).Value().Mass)
		require.EqualValues(t, "HDPE", getMatX.Results.Compositions().GetComposition(2).Value().Material)
		require.EqualValues(t, hdpeMass, getMatX.Results.Compositions().GetComposition(2).Value().Mass)
	}

	//fraction test
	var testSorter = ctx.NewSoloAgent()
	addSorter:= scauri.ScFuncs.AddSorter(ctx.Sign(owner))
	addSorter.Params.SorterID().SetValue(testSorter.ScAgentID())
	addSorter.Func.TransferIotas(1).Call()


	createFraction:= scauri.ScFuncs.CreateFraction(ctx.Sign(testSorter))
	createFraction.Params.Purpose().SetValue("Food")
	createFraction.Params.Name().SetValue("TestFraction")
	createFraction.Func.TransferIotas(1).Call()
	var fracID = createFraction.Results.FracID().Value()


    //add 15 packages of each type to a fraction
	var numPackages uint64 = 15

	for k:= 0; k < len(keys); k++ {

		for i:= 0; i < int(numPackages); i++ {
			addPPtoFractionX:= scauri.ScFuncs.AddPPToFraction(ctx)
			addPPtoFractionX.Params.PpID().SetValue(keys[k])
			addPPtoFractionX.Params.FracID().SetValue(fracID)
			addPPtoFractionX.Func.TransferIotas(1).Call()
		} 
	}

	
	//create impure Productpass
	createPP:= scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) 	//for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("NuclearWaste")
	createPP.Params.RecyclateShare().SetValue(0)
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Params.PackagesNumber().SetValue(100000)
	createPP.Params.PackageWeight().SetValue(12000)
	createPP.Params.Compositions().AppendComposition().SetValue(composition1)
	createPP.Params.Compositions().AppendComposition().SetValue(composition2)
	createPP.Params.Compositions().AppendComposition().SetValue(composition3)
	createPP.Func.TransferIotas(10).Call()
	var id1 = createPP.Results.Id().Value();
	fmt.Println("id1: " + fmt.Sprint(id1))

	addImpurePP:= scauri.ScFuncs.AddPPToFraction(ctx)
	addImpurePP.Params.PpID().SetValue(id1)
	addImpurePP.Params.FracID().SetValue(fracID)
	addImpurePP.Func.TransferIotas(1).Call()

	
	var expectedFraction *scauri.Fraction
	expectedFraction = new(scauri.Fraction)
	expectedFraction.FracId = fracID
	expectedFraction.Did = "tbd"
	expectedFraction.Name = "TestFraction"
    expectedFraction.Purpose = "false"					//because non food article was added
	expectedFraction.Issuer = testSorter.ScAgentID()
	expectedFraction.Amount = 1

	GetFraction:= scauri.ScFuncs.GetFraction(ctx)
	GetFraction.Params.FracID().SetValue(fracID)
	GetFraction.Func.Call()
	var fracResultsProxy = GetFraction.Results

	//Test fraction properties
	require.EqualValues(t, expectedFraction.FracId, fracResultsProxy.Fraction().Value().FracId)
	require.EqualValues(t, expectedFraction.Purpose, fracResultsProxy.Fraction().Value().Purpose)
	require.EqualValues(t, expectedFraction.Issuer, fracResultsProxy.Fraction().Value().Issuer)

    //Test resulting fraction compositions (for PP, PE and HDPE)
	var expectedFracCompositon1 *scauri.FracComposition
	expectedFracCompositon1 = new(scauri.FracComposition) 
	expectedFracCompositon1.Material = "PP"
	expectedFracCompositon1.Mass = numPackageTypes * numPackages * ppMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(0).Exists())
	require.EqualValues(t, expectedFracCompositon1.Material, fracResultsProxy.FracComposition().GetFracComposition(0).Value().Material)
	require.EqualValues(t, expectedFracCompositon1.Mass, fracResultsProxy.FracComposition().GetFracComposition(0).Value().Mass)
	require.NoError(t, ctx.Err)

	var expectedFracCompositon2 *scauri.FracComposition
	expectedFracCompositon2 = new(scauri.FracComposition) 
	expectedFracCompositon2.Material = "PE"
	expectedFracCompositon2.Mass = numPackageTypes * numPackages * peMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(1).Exists())
	require.EqualValues(t, expectedFracCompositon2.Material, fracResultsProxy.FracComposition().GetFracComposition(1).Value().Material)
	require.EqualValues(t, expectedFracCompositon2.Mass, fracResultsProxy.FracComposition().GetFracComposition(1).Value().Mass)
	require.NoError(t, ctx.Err)
	
	var expectedFracCompositon3 *scauri.FracComposition
	expectedFracCompositon3 = new(scauri.FracComposition) 
	expectedFracCompositon3.Material = "HDPE"
	expectedFracCompositon3.Mass = numPackageTypes * numPackages * hdpeMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(2).Exists())
	require.EqualValues(t, expectedFracCompositon3.Material, fracResultsProxy.FracComposition().GetFracComposition(2).Value().Material)
	require.EqualValues(t, expectedFracCompositon3.Mass, fracResultsProxy.FracComposition().GetFracComposition(2).Value().Mass)
	require.NoError(t, ctx.Err)

	var testRecycler = ctx.NewSoloAgent()
	addRecycler:= scauri.ScFuncs.AddRecycler(ctx.Sign(owner))
	addRecycler.Params.RecyclerID().SetValue(testRecycler.ScAgentID())
	addRecycler.Func.TransferIotas(1).Call()

	//test recyclate
	createRecyclate:= scauri.ScFuncs.CreateRecyclate(ctx.Sign(testRecycler))
	createRecyclate.Params.FracID().SetValue(fracID)
	createRecyclate.Params.Name().SetValue("TestRecyclate")
	createRecyclate.Func.TransferIotas(1).Call()
	var recyID = createRecyclate.Results.RecyID().Value()

	var expectedRecyclate *scauri.Recyclate
	expectedRecyclate = new(scauri.Recyclate)
	expectedRecyclate.RecyId = recyID
	expectedRecyclate.FracId = fracID
	expectedRecyclate.Did = "tbd"
	expectedRecyclate.Name = "TestRecyclate"
    expectedRecyclate.Purpose = "Food"
	expectedRecyclate.Issuer = testRecycler.ScAgentID()
	expectedRecyclate.Amount = 1

	GetRecyclate:= scauri.ScFuncs.GetRecyclate(ctx)
	GetRecyclate.Params.RecyID().SetValue(recyID)
	GetRecyclate.Func.Call()
	var recyResultsProxy = GetRecyclate.Results

	require.EqualValues(t, expectedRecyclate.RecyId, recyResultsProxy.Recyclate().Value().RecyId)
	require.EqualValues(t, expectedRecyclate.FracId, recyResultsProxy.Recyclate().Value().FracId)
	require.EqualValues(t, expectedRecyclate.Purpose, recyResultsProxy.Recyclate().Value().Purpose)
	require.EqualValues(t, expectedRecyclate.Issuer, recyResultsProxy.Recyclate().Value().Issuer)


	//Test resulting recycling compositions (for PP, PE and HDPE). Should be the same as for the fraction
	require.True(t, GetRecyclate.Results.RecyComposition().GetRecyComposition(0).Exists())
	require.EqualValues(t, expectedFracCompositon1.Material, recyResultsProxy.RecyComposition().GetRecyComposition(0).Value().Material)
	require.EqualValues(t, expectedFracCompositon1.Mass, recyResultsProxy.RecyComposition().GetRecyComposition(0).Value().Mass)
	require.NoError(t, ctx.Err)

	require.True(t, GetRecyclate.Results.RecyComposition().GetRecyComposition(1).Exists())
	require.EqualValues(t, expectedFracCompositon2.Material, recyResultsProxy.RecyComposition().GetRecyComposition(1).Value().Material)
	require.EqualValues(t, expectedFracCompositon2.Mass, recyResultsProxy.RecyComposition().GetRecyComposition(1).Value().Mass)
	require.NoError(t, ctx.Err)

	require.True(t, GetRecyclate.Results.RecyComposition().GetRecyComposition(2).Exists())
	require.EqualValues(t, expectedFracCompositon3.Material, recyResultsProxy.RecyComposition().GetRecyComposition(2).Value().Material)
	require.EqualValues(t, expectedFracCompositon3.Mass, recyResultsProxy.RecyComposition().GetRecyComposition(2).Value().Mass)
	require.NoError(t, ctx.Err)
	//test money distribution
	//test deletion
	//test access
}