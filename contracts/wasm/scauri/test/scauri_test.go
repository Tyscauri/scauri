// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/iotaledger/wasp/contracts/wasm/scauri/go/scauri"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/stretchr/testify/require"
)

//some test parameters
var iotasAddedToCharge uint64 = 1200
var packagesPerCharge uint64 = 100
var numPackages uint64 = 15       //FRACTION: Add 15 packages of each type to a fraction
const numPackageTypes uint64 = 10 //creates 10 package charges

func setupTest(t *testing.T) *wasmsolo.SoloContext {
	return wasmsolo.NewSoloContext(t, scauri.ScName, scauri.OnLoad)
}

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, scauri.ScName, scauri.OnLoad)
	require.NoError(t, ctx.ContractExists(scauri.ScName))
}

func TestSuccessfullRecyclingCircle(t *testing.T) {

	ctx := setupTest(t)
	owner := ctx.Creator()

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

	var compositionArray [3]*scauri.Composition

	compositionArray[0] = composition1
	compositionArray[1] = composition2
	compositionArray[2] = composition3

	//PP: Create single Productpass
	createPP := scauri.ScFuncs.CreatePP(ctx.Sign(ctx.NewSoloAgent())) //for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("Food")
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Params.PackagesNumber().SetValue(packagesPerCharge)
	createPP.Params.PackageWeight().SetValue(12000)
	createPP.Params.Compositions().AppendComposition().SetValue(composition1)
	createPP.Params.Compositions().AppendComposition().SetValue(composition2)
	createPP.Params.Compositions().AppendComposition().SetValue(composition3)
	createPP.Func.TransferIotas(iotasAddedToCharge).Call()
	var id1 = createPP.Results.Id().Value()
	fmt.Println("id1: " + fmt.Sprint(id1))

	getMat1 := scauri.ScFuncs.GetMaterials(ctx)
	getMat1.Params.Id().SetValue(id1)
	getMat1.Func.Call()
	require.NoError(t, ctx.Err)
	require.EqualValues(t, "PP", getMat1.Results.Compositions().GetComposition(0).Value().Material)

	var keys [numPackageTypes]wasmtypes.ScHash

	var testProducer = ctx.NewSoloAgent()
	var preBalanceProducer uint64 = testProducer.Balance() - uint64(len(keys))*iotasAddedToCharge // subtraction for the iota transfers afterwards, one for creation and second for getting payoff

	// PP: Create multiple (numPackageTypes) test productpasses and the storage of their material compositions
	for k := 0; k < len(keys); k++ {

		createPPX := scauri.ScFuncs.CreatePP(ctx.Sign(testProducer)) //for funcs
		createPPX.Params.Name().SetValue("tetrapack" + fmt.Sprint(k))
		createPPX.Params.Purpose().SetValue("Food")
		createPPX.Params.PackagesNumber().SetValue(packagesPerCharge)
		createPPX.Params.PackageWeight().SetValue(12000)
		createPPX.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
		//add materials
		for i := 0; i < len(compositionArray); i++ {
			createPPX.Params.Compositions().AppendComposition().SetValue(compositionArray[i])
			require.NoError(t, ctx.Err)

		}
		createPPX.Func.TransferIotas(iotasAddedToCharge).Call()
		id := createPPX.Results.Id().Value()
		keys[k] = id
		require.NoError(t, ctx.Err)

	}

	//PP: Test Materials set
	for k := 0; k < len(keys); k++ {
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

	//FRACTION: Testing
	var testSorter = ctx.NewSoloAgent()
	addSorter := scauri.ScFuncs.AddSorter(ctx.Sign(owner))
	addSorter.Params.SorterID().SetValue(testSorter.ScAgentID())
	addSorter.Func.TransferIotas(1).Call()

	createFraction := scauri.ScFuncs.CreateFraction(ctx.Sign(testSorter))
	createFraction.Params.Purpose().SetValue("Food")
	createFraction.Params.Name().SetValue("TestFraction")
	createFraction.Func.TransferIotas(1).Call()
	var fracID = createFraction.Results.FracID().Value()

	//FRACTION: Add 15 packages of each type to a fraction
	for k := 0; k < len(keys); k++ {

		for i := 0; i < int(numPackages); i++ {
			addPPtoFractionX := scauri.ScFuncs.AddPPToFraction(ctx)
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
	expectedFraction.Pure = true
	expectedFraction.Issuer = testSorter.ScAgentID()
	expectedFraction.Amount = (iotasAddedToCharge / packagesPerCharge) * 75 / 100 * numPackageTypes * numPackages

	GetFraction := scauri.ScFuncs.GetFraction(ctx)
	GetFraction.Params.FracID().SetValue(fracID)
	GetFraction.Func.Call()
	var fracResultsProxy = GetFraction.Results

	//FRACTION: Test fraction properties
	require.EqualValues(t, expectedFraction.FracId, fracResultsProxy.Fraction().Value().FracId)
	require.EqualValues(t, expectedFraction.Purpose, fracResultsProxy.Fraction().Value().Purpose)
	require.EqualValues(t, expectedFraction.Pure, fracResultsProxy.Fraction().Value().Pure)
	require.EqualValues(t, expectedFraction.Issuer, fracResultsProxy.Fraction().Value().Issuer)
	require.EqualValues(t, expectedFraction.Amount, fracResultsProxy.Fraction().Value().Amount)

	//FRACTION: Test resulting fraction compositions (for PP, PE and HDPE)
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

	//FRACTION: Test payoffs producer
	var expectedPayoffProducer uint64 = (iotasAddedToCharge/packagesPerCharge)*25/100*numPackageTypes*numPackages - 1

	payoutProducer := scauri.ScFuncs.PayoutProducer(ctx.Sign(testProducer))
	payoutProducer.Params.ProdID().SetValue(testProducer.ScAgentID())
	payoutProducer.Func.TransferIotas(1).Call()

	require.EqualValues(t, expectedPayoffProducer, testProducer.Balance()-preBalanceProducer)

	var testRecycler = ctx.NewSoloAgent()
	addRecycler := scauri.ScFuncs.AddRecycler(ctx.Sign(owner))
	addRecycler.Params.RecyclerID().SetValue(testRecycler.ScAgentID())
	addRecycler.Func.TransferIotas(1).Call()

	//RECYCLATE: Test recyclate
	var preBalanceRecycler uint64 = testRecycler.Balance() - 1 //-1 for the func call afterwards

	createRecyclate := scauri.ScFuncs.CreateRecyclate(ctx.Sign(testRecycler))
	createRecyclate.Params.FracID().SetValue(fracID)
	createRecyclate.Params.Name().SetValue("TestRecyclate")
	createRecyclate.Func.TransferIotas(1).Call()
	var recyID = createRecyclate.Results.RecyID().Value()

	//RECYCLATE: Test recyclate properties
	var expectedRecyclate *scauri.Recyclate
	expectedRecyclate = new(scauri.Recyclate)
	expectedRecyclate.RecyId = recyID
	expectedRecyclate.FracId = fracID
	expectedRecyclate.Did = "tbd"
	expectedRecyclate.Name = "TestRecyclate"
	expectedRecyclate.Purpose = "Food"
	expectedRecyclate.Pure = true
	expectedRecyclate.Issuer = testRecycler.ScAgentID()

	GetRecyclate := scauri.ScFuncs.GetRecyclate(ctx)
	GetRecyclate.Params.RecyID().SetValue(recyID)
	GetRecyclate.Func.Call()
	var recyResultsProxy = GetRecyclate.Results

	require.EqualValues(t, expectedRecyclate.RecyId, recyResultsProxy.Recyclate().Value().RecyId)
	require.EqualValues(t, expectedRecyclate.FracId, recyResultsProxy.Recyclate().Value().FracId)
	require.EqualValues(t, expectedRecyclate.Purpose, recyResultsProxy.Recyclate().Value().Purpose)
	require.EqualValues(t, expectedRecyclate.Issuer, recyResultsProxy.Recyclate().Value().Issuer)

	//RECYCLATE: Test resulting recycling compositions (for PP, PE and HDPE). Should be the same as for the fraction
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

	//RECYCLATE: Test payoff recyclate
	var expectedPayoffRecycler uint64 = expectedFraction.Amount
	require.EqualValues(t, expectedPayoffRecycler, testRecycler.Balance()-preBalanceRecycler)

	//test access
}

func TestUnsuccessfullRecyclingCircle(t *testing.T) {

	ctx := setupTest(t)
	owner := ctx.Creator()

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

	var compositionArray [3]*scauri.Composition

	compositionArray[0] = composition1
	compositionArray[1] = composition2
	compositionArray[2] = composition3

	//PP: create multiple (numPackageTypes) test productpasses and the storage of their material compositions
	var keys [numPackageTypes]wasmtypes.ScHash
	var testProducer = ctx.NewSoloAgent()
	var testDonationReceiver = ctx.NewSoloAgent()

	setDonationAddress := scauri.ScFuncs.SetDonationAddress(ctx.Sign(owner))
	setDonationAddress.Params.DonationAddress().SetValue(testDonationReceiver.ScAgentID())
	setDonationAddress.Func.TransferIotas(1).Call()

	var donPreBalance = testDonationReceiver.Balance()

	for k := 0; k < len(keys); k++ {

		createPPX := scauri.ScFuncs.CreatePP(ctx.Sign(testProducer))
		createPPX.Params.Name().SetValue("tetrapack" + fmt.Sprint(k))
		createPPX.Params.Purpose().SetValue("Food")
		createPPX.Params.PackagesNumber().SetValue(packagesPerCharge)
		createPPX.Params.PackageWeight().SetValue(12000)
		createPPX.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
		//add materials
		for i := 0; i < len(compositionArray); i++ {
			createPPX.Params.Compositions().AppendComposition().SetValue(compositionArray[i])
			require.NoError(t, ctx.Err)

		}
		createPPX.Func.TransferIotas(iotasAddedToCharge).Call()
		id := createPPX.Results.Id().Value()
		keys[k] = id
		require.NoError(t, ctx.Err)

	}

	//PP: Test materials
	for k := 0; k < len(keys); k++ {
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

	//FRACTION: Fraction test
	var testSorter = ctx.NewSoloAgent()
	addSorter := scauri.ScFuncs.AddSorter(ctx.Sign(owner))
	addSorter.Params.SorterID().SetValue(testSorter.ScAgentID())
	addSorter.Func.TransferIotas(1).Call()

	createFraction := scauri.ScFuncs.CreateFraction(ctx.Sign(testSorter))
	createFraction.Params.Purpose().SetValue("Food")
	createFraction.Params.Name().SetValue("TestFraction")
	createFraction.Func.TransferIotas(1).Call()
	var fracID = createFraction.Results.FracID().Value()

	//FRACTION: Add 'numPackages' packages of each type to a fraction
	for k := 0; k < len(keys); k++ {

		for i := 0; i < int(numPackages); i++ {
			addPPtoFractionX := scauri.ScFuncs.AddPPToFraction(ctx)
			addPPtoFractionX.Params.PpID().SetValue(keys[k])
			addPPtoFractionX.Params.FracID().SetValue(fracID)
			addPPtoFractionX.Func.TransferIotas(1).Call()
		}
	}

	//create impure Productpass
	createPP := scauri.ScFuncs.CreatePP(ctx.Sign(testProducer))
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("NuclearWaste")
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Params.PackagesNumber().SetValue(packagesPerCharge)
	createPP.Params.PackageWeight().SetValue(12000)
	createPP.Params.Compositions().AppendComposition().SetValue(composition1)
	createPP.Params.Compositions().AppendComposition().SetValue(composition2)
	createPP.Params.Compositions().AppendComposition().SetValue(composition3)
	createPP.Func.TransferIotas(iotasAddedToCharge).Call()
	var id1 = createPP.Results.Id().Value()
	fmt.Println("id1: " + fmt.Sprint(id1))

	var preBalanceProducer = testProducer.Balance()

	addImpurePP := scauri.ScFuncs.AddPPToFraction(ctx.Sign(testSorter))
	addImpurePP.Params.PpID().SetValue(id1)
	addImpurePP.Params.FracID().SetValue(fracID)
	addImpurePP.Func.TransferIotas(1).Call()

	//FRACTION: Test fraction properties
	var expectedFraction *scauri.Fraction
	expectedFraction = new(scauri.Fraction)
	expectedFraction.FracId = fracID
	expectedFraction.Did = "tbd"
	expectedFraction.Name = "TestFraction"
	expectedFraction.Purpose = "Food"
	expectedFraction.Pure = false //because non food article was added
	expectedFraction.Issuer = testSorter.ScAgentID()
	expectedFraction.Amount = (iotasAddedToCharge / packagesPerCharge) * 75 / 100 * numPackageTypes * numPackages

	GetFraction := scauri.ScFuncs.GetFraction(ctx)
	GetFraction.Params.FracID().SetValue(fracID)
	GetFraction.Func.Call()
	var fracResultsProxy = GetFraction.Results

	require.EqualValues(t, expectedFraction.FracId, fracResultsProxy.Fraction().Value().FracId)
	require.EqualValues(t, expectedFraction.Purpose, fracResultsProxy.Fraction().Value().Purpose)
	require.EqualValues(t, expectedFraction.Pure, fracResultsProxy.Fraction().Value().Pure)
	require.EqualValues(t, expectedFraction.Issuer, fracResultsProxy.Fraction().Value().Issuer)

	//FRACTION: Test resulting fraction compositions (for PP, PE and HDPE)
	var expectedFracCompositon1 *scauri.FracComposition
	expectedFracCompositon1 = new(scauri.FracComposition)
	expectedFracCompositon1.Material = "PP"
	expectedFracCompositon1.Mass = numPackageTypes*numPackages*ppMass + ppMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(0).Exists())
	require.EqualValues(t, expectedFracCompositon1.Material, fracResultsProxy.FracComposition().GetFracComposition(0).Value().Material)
	require.EqualValues(t, expectedFracCompositon1.Mass, fracResultsProxy.FracComposition().GetFracComposition(0).Value().Mass)
	require.NoError(t, ctx.Err)

	var expectedFracCompositon2 *scauri.FracComposition
	expectedFracCompositon2 = new(scauri.FracComposition)
	expectedFracCompositon2.Material = "PE"
	expectedFracCompositon2.Mass = numPackageTypes*numPackages*peMass + peMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(1).Exists())
	require.EqualValues(t, expectedFracCompositon2.Material, fracResultsProxy.FracComposition().GetFracComposition(1).Value().Material)
	require.EqualValues(t, expectedFracCompositon2.Mass, fracResultsProxy.FracComposition().GetFracComposition(1).Value().Mass)
	require.NoError(t, ctx.Err)

	var expectedFracCompositon3 *scauri.FracComposition
	expectedFracCompositon3 = new(scauri.FracComposition)
	expectedFracCompositon3.Material = "HDPE"
	expectedFracCompositon3.Mass = numPackageTypes*numPackages*hdpeMass + hdpeMass

	require.True(t, GetFraction.Results.FracComposition().GetFracComposition(2).Exists())
	require.EqualValues(t, expectedFracCompositon3.Material, fracResultsProxy.FracComposition().GetFracComposition(2).Value().Material)
	require.EqualValues(t, expectedFracCompositon3.Mass, fracResultsProxy.FracComposition().GetFracComposition(2).Value().Mass)
	require.NoError(t, ctx.Err)

	//FRACTION: Test correct payoff for producer
	var expectedPayoffProducer uint64 = (iotasAddedToCharge/packagesPerCharge)*25/100*numPackageTypes*numPackages - 1

	payoutProducer := scauri.ScFuncs.PayoutProducer(ctx.Sign(testProducer))
	payoutProducer.Params.ProdID().SetValue(testProducer.ScAgentID())
	payoutProducer.Func.TransferIotas(1).Call()
	require.EqualValues(t, expectedPayoffProducer, testProducer.Balance()-preBalanceProducer)

	//RECYCLATE
	var testRecycler = ctx.NewSoloAgent()
	addRecycler := scauri.ScFuncs.AddRecycler(ctx.Sign(owner))
	addRecycler.Params.RecyclerID().SetValue(testRecycler.ScAgentID())
	addRecycler.Func.TransferIotas(1).Call()

	var preBalanceRecycler uint64 = testRecycler.Balance() - 1 //-1 for the func call afterwards

	//RECYLCATE: Test recyclate properties
	createRecyclate := scauri.ScFuncs.CreateRecyclate(ctx.Sign(testRecycler))
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
	expectedRecyclate.Pure = false
	expectedRecyclate.Issuer = testRecycler.ScAgentID()

	GetRecyclate := scauri.ScFuncs.GetRecyclate(ctx)
	GetRecyclate.Params.RecyID().SetValue(recyID)
	GetRecyclate.Func.Call()
	var recyResultsProxy = GetRecyclate.Results

	require.EqualValues(t, expectedRecyclate.RecyId, recyResultsProxy.Recyclate().Value().RecyId)
	require.EqualValues(t, expectedRecyclate.FracId, recyResultsProxy.Recyclate().Value().FracId)
	require.EqualValues(t, expectedRecyclate.Purpose, recyResultsProxy.Recyclate().Value().Purpose)
	require.EqualValues(t, expectedRecyclate.Pure, recyResultsProxy.Recyclate().Value().Pure)
	require.EqualValues(t, expectedRecyclate.Issuer, recyResultsProxy.Recyclate().Value().Issuer)
	require.NoError(t, ctx.Err)

	//RECYCLATE: Test resulting recycling compositions (for PP, PE and HDPE). Should be the same as for the fraction
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

	//RECYCLATE: Test payoff recyclate
	//Sorter/Recycler shall obtain no reward
	var expectedPayoffRecycler uint64 = 0 // because fraction was polluted
	require.EqualValues(t, expectedPayoffRecycler, testRecycler.Balance()-preBalanceRecycler)

	//Money shall be transferred to donation address
	var expectedPayoffDonation uint64 = expectedFraction.Amount
	require.EqualValues(t, expectedPayoffDonation, testDonationReceiver.Balance()-donPreBalance)

	//test token to donate & implement get token to donate
	payoutDonation := scauri.ScFuncs.PayoutDonation(ctx.Sign(testDonationReceiver))
	payoutDonation.Func.TransferIotas(1).Call()

	expectedPayoffDonation += (iotasAddedToCharge / packagesPerCharge) - 1
	require.EqualValues(t, expectedPayoffDonation, testDonationReceiver.Balance()-donPreBalance)

	//test access
}

func TestResourceOverdrain(t *testing.T) {

	ctx := setupTest(t)
	owner := ctx.Creator()

	var composition *scauri.Composition
	composition = new(scauri.Composition)
	composition.Material = "PP"
	var ppMass uint64 = 8000 // mass of PP in mg
	composition.Mass = ppMass

	var testProducer = ctx.NewSoloAgent()

	createPP := scauri.ScFuncs.CreatePP(ctx.Sign(testProducer)) //for funcs
	createPP.Params.Name().SetValue("Chips")
	createPP.Params.Purpose().SetValue("Food")
	createPP.Params.ExpiryDate().SetValue(uint64(time.Now().Unix()))
	createPP.Params.PackagesNumber().SetValue(numPackages)
	createPP.Params.PackageWeight().SetValue(8000)
	createPP.Params.Compositions().AppendComposition().SetValue(composition)
	createPP.Func.TransferIotas(iotasAddedToCharge).Call()
	var id = createPP.Results.Id().Value()
	fmt.Println("id: " + fmt.Sprint(id))

	var testSorter = ctx.NewSoloAgent()
	addSorter := scauri.ScFuncs.AddSorter(ctx.Sign(owner))
	addSorter.Params.SorterID().SetValue(testSorter.ScAgentID())
	addSorter.Func.TransferIotas(1).Call()

	createFraction := scauri.ScFuncs.CreateFraction(ctx.Sign(testSorter))
	createFraction.Params.Purpose().SetValue("Cosmetics")
	createFraction.Params.Name().SetValue("TestFraction")
	createFraction.Func.TransferIotas(1).Call()
	var fracID = createFraction.Results.FracID().Value()
	require.NoError(t, ctx.Err)

	//sort all existing packages

	for i := 0; i < int(numPackages); i++ {
		addPPtoFractionX := scauri.ScFuncs.AddPPToFraction(ctx.Sign(testSorter))
		addPPtoFractionX.Params.PpID().SetValue(id)
		addPPtoFractionX.Params.FracID().SetValue(fracID)
		addPPtoFractionX.Func.TransferIotas(1).Call()
		require.NoError(t, ctx.Err)
	}

	//try to sort 2 packages more than exist
	for i := 0; i < 2; i++ {
		addPPtoFractionX := scauri.ScFuncs.AddPPToFraction(ctx.Sign(testSorter))
		addPPtoFractionX.Params.PpID().SetValue(id)
		addPPtoFractionX.Params.FracID().SetValue(fracID)
		addPPtoFractionX.Func.TransferIotas(1).Call()
		require.Error(t, ctx.Err)
	}

}

func TestUnauthorizedAccess(t *testing.T) {
	//somebody else than the owner tries to manipulate the address for donations

	ctx := setupTest(t)
	var owner = ctx.Creator()
	var plantTreesNGO = ctx.NewSoloAgent()
	var unauthorized = ctx.NewSoloAgent()

	setDonationAddressValid := scauri.ScFuncs.SetDonationAddress(ctx.Sign(owner))
	setDonationAddressValid.Params.DonationAddress().SetValue(plantTreesNGO.ScAgentID())
	setDonationAddressValid.Func.TransferIotas(1).Call()

	setDonationAddressInvalid := scauri.ScFuncs.SetDonationAddress(ctx.Sign(unauthorized))
	setDonationAddressInvalid.Params.DonationAddress().SetValue(unauthorized.ScAgentID())
	setDonationAddressInvalid.Func.TransferIotas(1).Call()

	getDonationAddress := scauri.ScFuncs.GetDonationAddress(ctx)
	getDonationAddress.Func.Call()
	require.EqualValues(t, getDonationAddress.Results.DonationAddress().Value(), plantTreesNGO.ScAgentID())
	require.NoError(t, ctx.Err)

}
