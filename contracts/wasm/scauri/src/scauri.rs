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
    let chargeWeight: u64 = 1000000000; // in miligramm
    let packageWeight: u64 = 2000;  // in miligramm
    let totalPackages: u64 = chargeWeight / packageWeight;
    let amountPerCharge: u64 = ctx.incoming().balance(&ScColor::IOTA);
    let amountPerPackage: u64 = amountPerCharge * packageWeight / chargeWeight;
    let rewardPerPackageProducer: u64 = amountPerCharge / totalPackages * ((100 - f.state.share_recycler().value()) / 100) as u64;
    let rewardPaerPackageRecycler: u64 = amountPerCharge / totalPackages * f.state.share_recycler().value() as u64 / 100;
    let packagesSorted: u64 = 0;
    let packagesWrongSorted: u64 = 0;
    let packagesAlreadyPaid = 0;
    let activationDate: u64 = ctx.timestamp() / NANO_TIME_DIVIDER;
    let expiryDate: u64 = f.params.expiry_date().value() / NANO_TIME_DIVIDER;
    let lastPayout = activationDate;

    let ppNew = ProductPass{
        id: id,
        did: did,
        name: name,
        issuer: issuer,
        version: version,
        purpose: purpose,
        charge_weight: chargeWeight, // in miligramm
        package_weight: packageWeight, //in miligramm
        total_packages: totalPackages,
        packages_sorted: packagesSorted,
        packages_wrong_sorted: packagesWrongSorted,
        packages_already_paid: packagesAlreadyPaid,
        amount_per_charge: amountPerCharge,
        reward_per_package_producer: rewardPerPackageProducer,
        reward_per_package_recycler: rewardPaerPackageRecycler,
        activation_date: activationDate,
        expiry_date: expiryDate,
        last_producer_payout: lastPayout

    };
    
    let requiredToken = &ppNew.charge_weight * f.state.price_per_mg().value();
    if &ppNew.amount_per_charge < &requiredToken {
        ctx.panic(&format!("Charge does not provide sufficient token. '{tokens}'are required", tokens=requiredToken.to_string()));
    }
    
    //let share_recycler = f.state.share_recycler().value() as i64;
    f.state.productpasses().get_product_pass(&ppNew.id).set_value(&ppNew);
    //f.state.payoffs().get_uint64(&ppNew.issuer).set_value((amount * share_recycler / 100) as u64); //noting the tokens the producer gets in case the packaging is recycled
    f.results.id().set_value(&ppNew.id);
    //f.events.ppcreated();
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
        composition_results_proxy.get_composition(i).set_value(&comp.get_composition(i).value());
    }
    
}

//resets the material composition of a package charge
pub fn func_set_materials(ctx: &ScFuncContext, f: &SetMaterialsContext) {
    
    let id = f.params.id().value();
    let newComp = f.params.comp();
    
    let composition = f.state.compositions().get_compositions(&id);
    
    composition.clear();
    
    for i in 0 .. newComp.length() {
        composition.append_composition().set_value(&newComp.get_composition(i).value());
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
    
    /* not needed probably
    let newFracComp = FracComposition {
        material: "init".to_string(),
        weight: 0
        };
    
    f.state.frac_compositions().get_frac_compositions(&newFrac.frac_id).get_frac_composition(0).set_value(&newFracComp);
    */
    f.results.frac_id().set_value(&newFrac.frac_id); 			
}

pub fn func_add_pp_to_fraction(ctx: &ScFuncContext, f: &AddPPToFractionContext) {
    
    let ppID = f.params.pp_id().value();
    let fracID = f.params.frac_id().value();
    let pp_proxy = f.state.productpasses().get_product_pass(&ppID);
    let pp: ProductPass = pp_proxy.value();
    let ppComp = f.state.compositions().get_compositions(&ppID);
    let fracComp = f.state.frac_compositions().get_frac_compositions(&fracID);
    let frac_proxy = f.state.fractions().get_fraction(&fracID);
    

    for i in 0..ppComp.length() {
    
        let mut foundMat: bool = false;
        
        for j in 0..fracComp.length() {
            if ppComp.get_composition(i).value().material == fracComp.get_frac_composition(j).value().material {
                let newWeight = fracComp.get_frac_composition(j).value().weight + ppComp.get_composition(i).value().proportion as u64 * pp.package_weight /1000;   // div by 1000 because proportion is in per mil
                
                let newShare = FracComposition {
                    material: ppComp.get_composition(i).value().material,
                    weight: newWeight
                };
                fracComp.get_frac_composition(j).set_value(&newShare);
                foundMat = true;
            }
        if !foundMat {
            let index = fracComp.length();
            
            let mat = ppComp.get_composition(i).value().material;
            let wei: u64 = ppComp.get_composition(i).value().proportion as u64 * pp.package_weight / 1000;
            
            let newMat = FracComposition{
                material: mat,
                weight: wei
            };
            
            fracComp.append_frac_composition().set_value(&newMat);
            }
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
        let wei = fComp.weight;
        
        let recyComp = RecyComposition{
            material: mat,
            weight: wei
        };
        newRecyCompProxy.get_recy_composition(i).set_value(&recyComp);
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

    f.results.recyclate_id().set_value(&newRecy.recy_id);

}


fn create_random_hash(ctx: &ScFuncContext) -> ScHash {
    let random_value: u64 = ctx.random(u64::MAX - 1) + 1;
    let random_value_bytes: [u8; 8] = unsafe { std::mem::transmute(random_value.to_le()) };
    let random_hash: ScHash = ctx.utility().hash_sha3(&random_value_bytes);

    return random_hash;
}


/*
pub fn func_payout(ctx: &ScFuncContext, f: &PayoutContext) {
    let keys: ArrayOfMutableAgentID = f.state.payoff_keys();
    let payoffs_proxy: MapAgentIDToMutableUint64 = f.state.payoffs();
    
    for i in 0..keys.length() {
        let address: ScAgentID = keys.get_agent_id(i).value();
        let payoff: i64 = payoffs_proxy.get_uint64(&address).value() as i64;
     
        if payoff > 1000000 {
        
            let transfers: ScTransfers = ScTransfers::iotas(payoff);

            ctx.transfer_to_address(&address.address(), transfers);
            payoffs_proxy.get_uint64(&address).set_value(0);
        }
    }
}
*/

//wird nicht mehr gebraucht
/*
pub fn func_payout_frac(ctx: &ScFuncContext, f: &PayoutFracContext) {
    
    let lastpayout = f.state.last_payout();
    let currentTime: u64 = ctx.timestamp() / NANO_TIME_DIVIDER;                   // timestamp in nano secs
    
    if lastpayout.value() + 86400 < currentTime{                                  //can be called only once per day (every 86400 seconds)
        let frac_id = f.params.frac_id().value();
        let keys: ArrayOfMutableAgentID = f.state.payoff_keys_frac().get_frac_payoff_keys(&frac_id);
        let payoffs_proxy: MapAgentIDToMutableUint64 = f.state.payoffs_frac().get_frac_payoffs(&frac_id);
    
        for i in 0..keys.length() {
            let address: ScAgentID = keys.get_agent_id(i).value();
            let payoff: u64 = payoffs_proxy.get_uint64(&address).value();
     
            if payoff > 0 {
        
                let transfers: ScTransfers = ScTransfers::iotas(payoff);

                ctx.transfer_to_address(&address.address(), transfers);
                payoffs_proxy.get_uint64(&address).set_value(0);
            }
        }
    
        keys.clear();
        payoffs_proxy.clear();
        lastpayout.set_value(currentTime);
    }
    else {
        ctx.panic("Functions can only be called once every 24 hours.");

    }

}*/

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

    //payout remaing packages
    if pp_value.expiry_date > ctx.timestamp() / NANO_TIME_DIVIDER || total_number_packages == pp_value.packages_sorted + pp_value.packages_wrong_sorted {
        if pp_value.packages_sorted > pp_value.packages_already_paid {
            let address: ScAgentID = pp_value.issuer;
            let payout = pp_value.reward_per_package_producer * (&pp_value.packages_sorted - &pp_value.packages_already_paid);
        
            pp_proxy.value().packages_already_paid = pp_value.packages_sorted;
        
            let transfers: ScTransfers = ScTransfers::iotas(payout);
            ctx.transfer_to_address(&address.address(), transfers);
        }

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
