package boarding

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// BoardingTemplateGetDb ...
func BoardingTemplateGetDb(
	projectName commons.ProjectName,
	licenceSPDX commons.Licence,
	groupRoleName commons.GroupRoleName,
	onBoarding commons.Boarding,
	outBoarding commons.Boarding,
	hash commons.Hash,
	state commons.StateReason,
	dbConf persistance.DbConf,
) commons.Template {
	template := GetTemplate(projectName.Name, licenceSPDX.SPDX, groupRoleName.Name, state.Reason, dbConf)
	templateDTO := commons.BuildTemplateDTO(template, dbConf)

	// Update templateDTO data based on onBoarding and outBoarding parameters

	return templateDTO
}

// BoardingDB ...
func BoardingDB(_type string, resource commons.Template, state string, dbConf persistance.DbConf) commons.StateReason {
	InsertTemplate(_type, resource, state, dbConf)
	return commons.StateReason{}
}

// BoardingGetDb ...
func BoardingGetDb(_type string, xPubS commons.XPub, state commons.StateReason, dbConf persistance.DbConf) commons.Template {
	template := GetTemplateByTypeXPubAndState(_type, xPubS.XPub, state.Reason, dbConf)
	templateDTO := commons.BuildTemplateDTO(template, dbConf)

	return templateDTO
}
