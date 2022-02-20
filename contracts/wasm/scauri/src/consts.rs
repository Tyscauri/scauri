// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME        : &str = "scauri";
pub const SC_DESCRIPTION : &str = "scauri description";
pub const HSC_NAME       : ScHname = ScHname(0x5559a7fd);

pub const PARAM_CHARGE_WEIGHT    : &str = "chargeWeight";
pub const PARAM_COMPOSITIONS     : &str = "compositions";
pub const PARAM_DONATION_ADDRESS : &str = "donationAddress";
pub const PARAM_EXPIRY_DATE      : &str = "expiryDate";
pub const PARAM_FRAC_ID          : &str = "fracID";
pub const PARAM_ID               : &str = "id";
pub const PARAM_NAME             : &str = "name";
pub const PARAM_OWNER            : &str = "owner";
pub const PARAM_PACKAGE_WEIGHT   : &str = "packageWeight";
pub const PARAM_PACKAGES_NUMBER  : &str = "packagesNumber";
pub const PARAM_PP_ID            : &str = "ppID";
pub const PARAM_PROD_PASS        : &str = "prodPass";
pub const PARAM_PURPOSE          : &str = "purpose";
pub const PARAM_RECY_ID          : &str = "recyID";
pub const PARAM_RECYCLER_ID      : &str = "recyclerID";
pub const PARAM_SORTER_ID        : &str = "sorterID";

pub const RESULT_COMPOSITIONS      : &str = "compositions";
pub const RESULT_FRAC_COMPOSITION  : &str = "fracComposition";
pub const RESULT_FRAC_ID           : &str = "fracID";
pub const RESULT_FRACTION          : &str = "fraction";
pub const RESULT_ID                : &str = "id";
pub const RESULT_OWNER             : &str = "owner";
pub const RESULT_PP                : &str = "pp";
pub const RESULT_PPNAME            : &str = "ppname";
pub const RESULT_PPRESULT          : &str = "ppresult";
pub const RESULT_RECY_COMPOSITION  : &str = "recyComposition";
pub const RESULT_RECY_ID           : &str = "recyID";
pub const RESULT_RECYCLATE         : &str = "recyclate";
pub const RESULT_TOKEN_PER_PACKAGE : &str = "tokenPerPackage";
pub const RESULT_TOKEN_REQUIRED    : &str = "tokenRequired";

pub const STATE_COMPOSITIONS      : &str = "compositions";
pub const STATE_DONATION_ADDRESS  : &str = "donationAddress";
pub const STATE_FRAC_COMPOSITIONS : &str = "fracCompositions";
pub const STATE_FRACTIONS         : &str = "fractions";
pub const STATE_OWNER             : &str = "owner";
pub const STATE_PRICE_PER_MG      : &str = "pricePerMg";
pub const STATE_PRODUCTPASSES     : &str = "productpasses";
pub const STATE_RECY_COMPOSITIONS : &str = "recyCompositions";
pub const STATE_RECYCLATES        : &str = "recyclates";
pub const STATE_RECYCLERS         : &str = "recyclers";
pub const STATE_SHARE_RECYCLER    : &str = "shareRecycler";
pub const STATE_SORTERS           : &str = "sorters";
pub const STATE_TOKEN_TO_DONATE   : &str = "tokenToDonate";

pub const FUNC_ADD_PP_TO_FRACTION           : &str = "addPPToFraction";
pub const FUNC_ADD_RECYCLER                 : &str = "addRecycler";
pub const FUNC_ADD_SORTER                   : &str = "addSorter";
pub const FUNC_CREATE_FRACTION              : &str = "createFraction";
pub const FUNC_CREATE_PP                    : &str = "createPP";
pub const FUNC_CREATE_RECYCLATE             : &str = "createRecyclate";
pub const FUNC_DELETE_PP                    : &str = "deletePP";
pub const FUNC_INIT                         : &str = "init";
pub const FUNC_PAYOUT_DONATION              : &str = "payoutDonation";
pub const FUNC_PAYOUT_PRODUCER              : &str = "payoutProducer";
pub const FUNC_SET_DONATION_ADDRESS         : &str = "setDonationAddress";
pub const FUNC_SET_OWNER                    : &str = "setOwner";
pub const VIEW_GET_AMOUNT_OF_REQUIRED_FUNDS : &str = "getAmountOfRequiredFunds";
pub const VIEW_GET_FRACTION                 : &str = "getFraction";
pub const VIEW_GET_MATERIALS                : &str = "getMaterials";
pub const VIEW_GET_OWNER                    : &str = "getOwner";
pub const VIEW_GET_PP                       : &str = "getPP";
pub const VIEW_GET_RECYCLATE                : &str = "getRecyclate";
pub const VIEW_GET_TOKEN_PER_PACKAGE        : &str = "getTokenPerPackage";

pub const HFUNC_ADD_PP_TO_FRACTION           : ScHname = ScHname(0x50ac50a7);
pub const HFUNC_ADD_RECYCLER                 : ScHname = ScHname(0x25ae8407);
pub const HFUNC_ADD_SORTER                   : ScHname = ScHname(0x735539b5);
pub const HFUNC_CREATE_FRACTION              : ScHname = ScHname(0x59842fc3);
pub const HFUNC_CREATE_PP                    : ScHname = ScHname(0x673fc3d7);
pub const HFUNC_CREATE_RECYCLATE             : ScHname = ScHname(0x5066d840);
pub const HFUNC_DELETE_PP                    : ScHname = ScHname(0x56dedc36);
pub const HFUNC_INIT                         : ScHname = ScHname(0x1f44d644);
pub const HFUNC_PAYOUT_DONATION              : ScHname = ScHname(0x4cf2009a);
pub const HFUNC_PAYOUT_PRODUCER              : ScHname = ScHname(0x3a56494b);
pub const HFUNC_SET_DONATION_ADDRESS         : ScHname = ScHname(0x5b375a7d);
pub const HFUNC_SET_OWNER                    : ScHname = ScHname(0x2a15fe7b);
pub const HVIEW_GET_AMOUNT_OF_REQUIRED_FUNDS : ScHname = ScHname(0xbcbaf2d6);
pub const HVIEW_GET_FRACTION                 : ScHname = ScHname(0xe1c332fa);
pub const HVIEW_GET_MATERIALS                : ScHname = ScHname(0x1dca8ddb);
pub const HVIEW_GET_OWNER                    : ScHname = ScHname(0x137107a6);
pub const HVIEW_GET_PP                       : ScHname = ScHname(0xbf711e08);
pub const HVIEW_GET_RECYCLATE                : ScHname = ScHname(0x4e87f53b);
pub const HVIEW_GET_TOKEN_PER_PACKAGE        : ScHname = ScHname(0x522a14c0);
