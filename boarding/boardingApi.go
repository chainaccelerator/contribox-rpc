package boarding

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// BoardingTemplateGet ...
// Retrieve signature template, while creating boarding requests
// TODO : should the boarding requests be created further in the process ?
func BoardingTemplateGet(
	projectName commons.ProjectName,
	licenceSPDX commons.Licence,
	groupRoleName commons.GroupRoleName,
	onBoarding commons.Boarding,
	outBoarding commons.Boarding,
	hash commons.Hash,
	state commons.StateReason,
	dbConf persistance.DbConf,
) (commons.Template, bool) {
	if !ValidateBoardingTemplateGet(groupRoleName.Name, onBoarding, outBoarding, hash.Hash, state.Reason) {
		return commons.Template{}, true
	}
	template := BoardingTemplateGetDb(
		projectName,
		licenceSPDX,
		groupRoleName,
		onBoarding,
		outBoarding,
		hash,
		state,
		dbConf,
	)
	return template, template.GroupRoleName.Name == ""
}

// Boarding ...
// Save the signature template for a boarding request
func Boarding(_type string, resource commons.Template, state commons.StateReason, dbConf persistance.DbConf) (commons.StateReason, bool) {
	if !ValidateBoarding(_type, resource, state.Reason) {
		return commons.StateReason{}, true
	}
	stateResult := BoardingDb(_type, resource, state.Reason, dbConf)

	return stateResult, stateResult.Reason == ""
}

// BoardingGet ...
// Get a template associated with a boarding request
func BoardingGet(_type string, xPubS commons.XPub, state commons.StateReason, dbConf persistance.DbConf) (commons.Template, bool) {
	if !ValidateBoardingGet(_type, xPubS.XPub, state.Reason) {
		return commons.Template{}, true
	}
	template := BoardingGetDb(_type, xPubS, state, dbConf)
	return template, template.Hash.Hash == ""
}

// BoardingBroadcast ...
// Share the boarding request with other nodes
func BoardingBroadcast(_type string, resourceList []commons.UTXO, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.StateReason, bool) {
	if !ValidateBoardingBroadcast(_type, resourceList, hash.Hash, state.Reason) {
		return commons.StateReason{}, true
	}
	return commons.StateReason{Reason: "done"}, false
}

// BoardingBroadcastGet ...
func BoardingBroadcastGet(_type string, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.UTXO, bool) {
	if !ValidateBoardingBroadcastGet(_type, hash.Hash, state.Reason) {
		return commons.UTXO{}, true
	}
	return commons.UTXO{
		Hash:   commons.Hash{Hash: "mock utxo hash"},
		Tx0Id:  commons.TxId{Id: "mock utxo txid"},
		UTXO:   commons.UTXOData{UTXO: "mock utxodata"},
		Script: commons.Script{Script: "mock utxo script"},
	}, false
}
