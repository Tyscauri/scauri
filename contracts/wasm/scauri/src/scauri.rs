// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use wasmlib::*;

use crate::*;
use crate::structs::*;
use crate::typedefs::*;

const NANO_TIME_DIVIDER: u64 = 1_000_000_000;

pub fn func_init(ctx: &ScFuncContext, f: &InitContext) {
    if f.params.owner().exists() {
        f.state.owner().set_value(&f.params.owner().value());
        return;
    }
    f.state.owner().set_value(&ctx.contract_creator());
    f.state.share_recycler().set_value(75);
    f.state.price_per_mg().set_value(2); //assuming 1€ per MIOTA: 2 Iota per miligramm result in 2 cent per 10 gramms of weight
}

pub fn func_set_owner(ctx: &ScFuncContext, f: &SetOwnerContext) {
    f.state.owner().set_value(&f.params.owner().value());
}

pub fn view_get_owner(ctx: &ScViewContext, f: &GetOwnerContext) {
    f.results.owner().set_value(&f.state.owner().value());
}

pub fn func_create_pp(ctx: &ScFuncContext, f: &CreatePPContext) {
    let pps: MapHashToMutableProductPass = f.state.productpasses();
    
    let id = create_random_hash(ctx);

    //check if ID is already in Array;
    if pps.get_product_pass(&id).exists() {
           ctx.panic("id already exist");
    }
    let did: String = "did:iota:".to_owned() + &id.to_string();
    let name: String = f.params.name().value();
    let issuer: ScAgentID = ctx.caller();
    let version: u8 = 1;               //TEST IMPLEMENTATION - ADD as Parameter in create func like: f.params.version().value();
    let purpose: String = f.params.purpose().value();
    let packageWeight: u64 = f.params.package_weight().value();  // in miligramm
    let packagesNumber: u64 = f.params.packages_number().value();
    let chargeWeight: u64 = packageWeight * packagesNumber; // in miligramm
    let amountPerCharge: u64 = ctx.incoming().balance(&ScColor::IOTA);
    let amountPerPackage: u64 = amountPerCharge / packagesNumber;
    let rewardPerPackageProducer: u64 = amountPerPackage * (100 - f.state.share_recycler().value() as u64) / 100;
    let rewardPerPackageRecycler: u64 = amountPerPackage * f.state.share_recycler().value() as u64 / 100;
    let packagesSorted: u64 = 0;
    let packagesWrongSorted: u64 = 0;
    let packagesAlreadyPaid = 0;
    let activationDate: u64 = ctx.timestamp() / NANO_TIME_DIVIDER;
    let expiryDate: u64 = f.params.expiry_date().value() / NANO_TIME_DIVIDER;
    let lastPayout = activationDate;
    let recyclateShare = f.params.recyclate_share().value();


    let ppNew = ProductPass{
        id: id,
        did: did,
        name: name,
        issuer: issuer,
        version: version,
        purpose: purpose,
        charge_weight: chargeWeight, // in miligramm
        package_weight: packageWeight, //in miligramm
        packages_number: packagesNumber,
        packages_sorted: packagesSorted,
        packages_wrong_sorted: packagesWrongSorted,
        packages_already_paid: packagesAlreadyPaid,
        amount_per_charge: amountPerCharge,
        reward_per_package_producer: rewardPerPackageProducer,
        reward_per_package_recycler: rewardPerPackageRecycler,
        activation_date: activationDate,
        expiry_date: expiryDate,
        last_producer_payout: lastPayout,
        recyclate_share: recyclateShare

    };
    
    let mut requiredToken = &ppNew.charge_weight * f.state.price_per_mg().value();
    requiredToken = 10; //just for testing, delete later

    if &ppNew.amount_per_charge < &requiredToken {
        ctx.panic(&format!("Charge does not provide sufficient token. '{tokens}'are required", tokens=requiredToken.to_string()));
    }
    
    f.state.productpasses().get_product_pass(&ppNew.id).set_value(&ppNew);
    f.results.id().set_value(&ppNew.id);
    ctx.log(&format!("result set value id: {id}", id=f.results.id().value().to_string()));
    ctx.log(&format!("result id: {id}", id=&ppNew.id.to_string()));


    //set Materials
        
    let newComps= f.params.compositions();
    
    if newComps.length() == 0 {
        ctx.panic("List of Materials cannot have length 0");
    }
    
    let composition = f.state.compositions().get_compositions(&ppNew.id);
    
    composition.clear();
    
    let mut weightTester: u64 = 0;

    for i in 0 .. newComps.length() {

        let currentComp = newComps.get_composition(i).value();
        composition.append_composition().set_value(&currentComp);
        weightTester += currentComp.mass;
    }

    if weightTester != ppNew.package_weight {
        ctx.panic(&format!("Composition of package does not add up to the package weight. Package weight stated is '{weightStated}', but sums up to '{testedWeight}'. Are you missing some materials?", weightStated = ppNew.package_weight, testedWeight = weightTester))
    }

}

pub fn view_get_pp(ctx: &ScViewContext, f: &GetPPContext) {
    let id = f.params.id().value();
    let pps: MapHashToImmutableProductPass = f.state.productpasses();
    
    if pps.get_product_pass(&id).exists() {
        let pp: ProductPass = pps.get_product_pass(&id).value();
        
        f.results.ppresult().set_value(&pp);
        f.results.ppname().set_value(&pp.name);
    }
    
    else {
    ctx.log("Product Charge ID not found");
    }
}

pub fn view_get_amount_of_required_funds(ctx: &ScViewContext, f: &GetAmountOfRequiredFundsContext) {
    let requiredFunds: u64 = f.params.charge_weight().value() * f.state.price_per_mg().value();
    f.results.token_required().set_value(requiredFunds);
}

pub fn view_get_token_per_package(ctx: &ScViewContext, f: &GetTokenPerPackageContext) {
    let pass = f.params.prod_pass().value();
    let pricePerMg = f.state.price_per_mg().value();
    let tokenPerPackage: u64 =  pass.package_weight * pricePerMg;
    
    f.results.token_per_package().set_value(tokenPerPackage);
}


pub fn view_get_materials(ctx: &ScViewContext, f: &GetMaterialsContext) {
    let id = f.params.id().value();
    
    let comp = f.state.compositions().get_compositions(&id);
    let composition_results_proxy = f.results.compositions();
    
    for i in 0..comp.length() {
        composition_results_proxy.append_composition().set_value(&comp.get_composition(i).value());
    }
    
}


pub fn func_create_fraction(ctx: &ScFuncContext, f: &CreateFractionContext) {
    
    let purpose: String = f.params.purpose().value();

    let fracID = create_random_hash(ctx);
    let did: String = "did:iota:".to_owned() + &"frac".to_owned() + &fracID.to_string();
    let name = f.params.name().value();
    let purpose: String = f.params.purpose().value();
    let issuer: ScAgentID = ctx.caller();    
    
    let newFrac = Fraction {
        frac_id: fracID,
        did: did,
        name: name,
        purpose: purpose,
        issuer: issuer,
        amount: 0
    };
    
    f.state.fractions().get_fraction(&newFrac.frac_id).set_value(&newFrac);
    f.results.frac_id().set_value(&newFrac.frac_id); 			
}

pub fn func_add_pp_to_fraction(ctx: &ScFuncContext, f: &AddPPToFractionContext) {
    
    let ppID = f.params.pp_id().value();
    let fracID = f.params.frac_id().value();
    let pp_proxy = f.state.productpasses().get_product_pass(&ppID);
    let pp: ProductPass = pp_proxy.value();
    let ppComps = f.state.compositions().get_compositions(&ppID);
    let fracComps = f.state.frac_compositions().get_frac_compositions(&fracID);
    let frac_proxy = f.state.fractions().get_fraction(&fracID);

    
    if fracComps.length() > 0{

        for i in 0..ppComps.length() {
    
            let mut foundMat: bool = false;
            
            for j in 0..fracComps.length() {
                if ppComps.get_composition(i).value().material == fracComps.get_frac_composition(j).value().material {
                    let newMass = fracComps.get_frac_composition(j).value().mass + ppComps.get_composition(i).value().mass;
                    let newShare = FracComposition {
                        material: ppComps.get_composition(i).value().material,
                        mass: newMass
                    };
                    fracComps.get_frac_composition(j).set_value(&newShare);
                    foundMat = true;
                }
            }    
            if !foundMat {
                
                let ppComp = ppComps.get_composition(i).value();
                let mat = ppComp.material;
                let mass: u64 = ppComp.mass;
                
                let newMat = FracComposition{
                    material: mat,
                    mass: mass
                };
                
                fracComps.append_frac_composition().set_value(&newMat);
            }
            
        }
    }

    else {
        for i in 0..ppComps.length() {

            let ppComp = ppComps.get_composition(i).value();
            let mat = ppComp.material;
            let mass: u64 = ppComp.mass;
            
            let newMat = FracComposition{
                material: mat,
                mass: mass
            };
            
            fracComps.append_frac_composition().set_value(&newMat);
            }
        
    }
    
    //organize money distribution
    //note that tracking packages which have been sorted to a fraction with the same purpose is basis for releasing funds
    if pp.purpose == frac_proxy.value().purpose {
        
        pp_proxy.value().packages_sorted = &pp.packages_sorted +1;
        frac_proxy.value().amount += &pp.reward_per_package_recycler;
    }

    else {          //currently sets the whole fraction unsuitable for the original application if one packaging added is not suitable for it
        if frac_proxy.value().purpose != "false" {
            frac_proxy.value().purpose = "false".to_string();
        }
        let donation_proxy = f.state.token_to_donate();
        donation_proxy.set_value(donation_proxy.value() + &pp.reward_per_package_producer + &pp.reward_per_package_recycler);
        pp_proxy.value().packages_wrong_sorted += 1;
    }    
}

pub fn func_create_recyclate(ctx: &ScFuncContext, f: &CreateRecyclateContext) {
    
    let fracID = f.params.frac_id().value();
    let fraction: Fraction = f.state.fractions().get_fraction(&fracID).value();

    let recyclateID = create_random_hash(ctx);    
    let did: String = "did:iota:".to_owned() + &"recy".to_owned() + &recyclateID.to_string();
    let name = f.params.name().value();
    let purpose = fraction.purpose;
    let issuer: ScAgentID = ctx.caller();    
    let amount: u64 = fraction.amount;
    
    let newRecy = Recyclate {
        recy_id: recyclateID,
        did: did,
        name: name,
        purpose: purpose,
        issuer: issuer,
        amount: amount,
        frac_id: fracID
    };
    
    let recyclates_proxy = f.state.recyclates();
    recyclates_proxy.get_recyclate(&newRecy.recy_id).set_value(&newRecy);
    
    let newRecyCompProxy = f.state.recy_compositions().get_recy_compositions(&newRecy.recy_id);
    let fracComp: ArrayOfMutableFracComposition = f.state.frac_compositions().get_frac_compositions(&newRecy.frac_id);
    
    for i in 0..fracComp.length() {
        
        let fComp = fracComp.get_frac_composition(i).value();
        let mat = fComp.material;
        let mass = fComp.mass;
        
        let recyComp = RecyComposition{
            material: mat,
            mass: mass
        };
        newRecyCompProxy.append_recy_composition().set_value(&recyComp);
    }
    
    //manage payouts
    if newRecy.amount > 0 {
        let mut address: ScAgentID;
        
        //only if the fraction is usable for the same purpose as the included packagings the money goes to the recycler, otherwise to a non profit address
        if newRecy.purpose != "false" {
            address = newRecy.issuer;
        }
        else {
            address = f.state.donation_address().value();
        }

        let payout = newRecy.amount;

        let transfers: ScTransfers = ScTransfers::iotas(payout);
        ctx.transfer_to_address(&address.address(), transfers);
        recyclates_proxy.get_recyclate(&newRecy.recy_id).value().amount = 0;
    }


    f.state.fractions().get_fraction(&newRecy.frac_id).delete();
    f.state.frac_compositions().get_frac_compositions(&newRecy.frac_id).clear();

    f.results.recy_id().set_value(&newRecy.recy_id);

}


fn create_random_hash(ctx: &ScFuncContext) -> ScHash {
    let random_value: u64 = ctx.random(u64::MAX - 1) + 1;
    let random_value_bytes: [u8; 8] = unsafe { std::mem::transmute(random_value.to_le()) };
    let random_hash: ScHash = ctx.utility().hash_sha3(&random_value_bytes);

    return random_hash;
}


pub fn func_payout_producer(ctx: &ScFuncContext, f: &PayoutProducerContext) {


    let ppID = f.params.frac_id().value();
    let pp_proxy = f.state.productpasses().get_product_pass(&ppID);
    let pp_value = pp_proxy.value();

    let lastpayout = pp_value.last_producer_payout;
    let currentTime: u64 = ctx.timestamp() / NANO_TIME_DIVIDER;    

    //can be called only once per day (every 86400 seconds)
    if lastpayout + 86400 < currentTime && pp_value.packages_sorted > pp_value.packages_already_paid {

        let address: ScAgentID = pp_value.issuer;
        let payout = pp_value.reward_per_package_producer * (&pp_value.packages_sorted - &pp_value.packages_already_paid);
    
        pp_proxy.value().packages_already_paid = pp_value.packages_sorted;
    
        let transfers: ScTransfers = ScTransfers::iotas(payout);
        ctx.transfer_to_address(&address.address(), transfers);
        pp_proxy.value().last_producer_payout = currentTime;
    }
    else {
        ctx.panic("Function can only be called once every 24 hours.");
    }

}

pub fn func_delete_pp(ctx: &ScFuncContext, f: &DeletePPContext) {

    let ppID = f.params.pp_id().value();
    let pp_proxy: MutableProductPass = f.state.productpasses().get_product_pass(&ppID);
    let pp_value = pp_proxy.value();
    let total_number_packages = pp_value.charge_weight / pp_value.package_weight;
    
    if pp_value.expiry_date > ctx.timestamp() / NANO_TIME_DIVIDER || total_number_packages == pp_value.packages_sorted + pp_value.packages_wrong_sorted {
    
        //payout remaing packages
            if pp_value.packages_sorted > pp_value.packages_already_paid {
                let address: ScAgentID = pp_value.issuer;
                let payout = pp_value.reward_per_package_producer * (&pp_value.packages_sorted - &pp_value.packages_already_paid);
            
                pp_proxy.value().packages_already_paid = pp_value.packages_sorted;
            
                let transfers: ScTransfers = ScTransfers::iotas(payout);
                ctx.transfer_to_address(&address.address(), transfers);
            }
    
        let updated_pp_value = pp_proxy.value();
    
        let leftover_token = updated_pp_value.amount_per_charge - 
            (updated_pp_value.packages_sorted + updated_pp_value.packages_wrong_sorted) * (updated_pp_value.reward_per_package_producer + updated_pp_value.reward_per_package_recycler);
    
        let token_to_donate_proxy = f.state.token_to_donate();
        token_to_donate_proxy.set_value(token_to_donate_proxy.value() + leftover_token);
        
        //remove pp
        f.state.productpasses().get_product_pass(&ppID).delete();
        f.state.compositions().get_compositions(&ppID).clear();      //probably not functional right now
    }
    
    else {
        ctx.panic("Product pass cannot be deleted now.")
    }
}

pub fn func_add_recycler(ctx: &ScFuncContext, f: &AddRecyclerContext) {
    f.state.recyclers().append_agent_id().set_value(&f.params.recycler_id().value());
}

pub fn func_add_sorter(ctx: &ScFuncContext, f: &AddSorterContext) {
    f.state.sorters().append_agent_id().set_value(&f.params.sorter_id().value());

}

pub fn view_get_fraction(ctx: &ScViewContext, f: &GetFractionContext) {

    let id = f.params.frac_id().value();
    let fractions: MapHashToImmutableFraction = f.state.fractions();
    
    if fractions.get_fraction(&id).exists() {

        let frac: Fraction = fractions.get_fraction(&id).value();        
        f.results.fraction().set_value(&frac);
        let frac_composition = f.state.frac_compositions().get_frac_compositions(&id);
        let fcomposition_results_proxy = f.results.frac_composition();

        for i in 0..frac_composition.length() {
            let comp =frac_composition.get_frac_composition(i).value();
            fcomposition_results_proxy.append_frac_composition().set_value(&comp);

        }
    }
    
    else {
    ctx.panic("Fraction ID not found");
    }

}

pub fn view_get_recyclate(ctx: &ScViewContext, f: &GetRecyclateContext) {
    let id = f.params.recy_id().value();
    let recyclates: MapHashToImmutableRecyclate = f.state.recyclates();
    
    if recyclates.get_recyclate(&id).exists() {
        let recy: Recyclate = recyclates.get_recyclate(&id).value();        
        f.results.recyclate().set_value(&recy);

        let recy_composition = f.state.recy_compositions().get_recy_compositions(&id);
        let rcomposition_results_proxy = f.results.recy_composition();

        for i in 0..recy_composition.length() {
            rcomposition_results_proxy.append_recy_composition().set_value(&recy_composition.get_recy_composition(i).value());
        }
    }
    
    else {
    ctx.panic("Recyclate ID not found");
    }
}
