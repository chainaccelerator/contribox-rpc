package boarding

import (
	"bc_node_api/api3/commons"
)

// BoardingTemplateGet ...
func BoardingTemplateGet(
	projectName commons.ProjectName,
	licenceSPDX commons.Licence,
	groupRoleName commons.GroupRoleName,
	onBoarding commons.Boarding,
	outBoarding commons.Boarding,
	hash commons.Hash,
	state commons.StateReason,
	dbConf commons.DbConf,
) (commons.Template, bool) {
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
func Boarding(_type string, resource commons.Template, state commons.StateReason, dbConf commons.DbConf) (commons.StateReason, bool) {
	return commons.StateReason{}, true
}

// BoardingGet ...
func BoardingGet(_type string, xPubS commons.XPub, state commons.StateReason, dbConf commons.DbConf) (commons.Template, bool) {
	return commons.Template{}, true
}

// BoardingBroadcast ...
func BoardingBroadcast(_type string, resourceList []commons.UTXO, hash commons.Hash, state commons.StateReason, dbConf commons.DbConf) (commons.StateReason, bool) {
	return commons.StateReason{}, true
}

// BoardingBroadcastGet ...
func BoardingBroadcastGet(_type string, hash commons.Hash, state commons.StateReason, dbConf commons.DbConf) (commons.UTXO, bool) {
	return commons.UTXO{}, true
}
