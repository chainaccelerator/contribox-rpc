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
	if len(onBoarding.XPubList) != 0 {
		insert := InsertBoarding(onBoarding, state.Reason, dbConf)
		if !insert {
			return commons.Template{}
		}
	}
	if len(outBoarding.XPubList) != 0 {
		insert := InsertBoarding(outBoarding, state.Reason, dbConf)
		if !insert {
			return commons.Template{}
		}
	}

	templateDTO := commons.Template{
		Hash:          commons.Hash{Hash: "boardingTemplateGet Hash"},
		ProjectName:   commons.ProjectName{Name: "boardingTemplateGet project name"},
		GroupRoleName: commons.GroupRoleName{Name: "member"},
	}

	return templateDTO
}

// BoardingDb ...
func BoardingDb(_type string, resource commons.Template, state string, dbConf persistance.DbConf) commons.StateReason {
	template := persistance.Template{
		Hash:               resource.Hash.Hash,
		ProjectName:        resource.ProjectName.Name,
		LicenceSPDX:        resource.LicenceSPDX.SPDX,
		State:              resource.State.Reason,
		UserRequirement:    resource.UserRequirement,
		ProjectRequirement: resource.ProjectRequirement,
		UserUser:           resource.UserUser.Quorum,
		UserBackup:         resource.UserBackup.Quorum,
		UserLock:           resource.UserLock.Quorum,
		UserWitness:        resource.UserWitness.Quorum,
		ProjectOld:         resource.ProjectOld.Quorum,
		ProjectParent:      resource.ProjectParent.Quorum,
		ProjectBoard:       resource.ProjectBoard.Quorum,
		ProjectMember:      resource.ProjectMember.Quorum,
		ProjectCosigner:    resource.ProjectCosigner.Quorum,
		ProjectWitness:     resource.ProjectWitness.Quorum,
	}
	id := InsertTemplate(_type, template, state, dbConf)
	if id == 0 {
		return commons.StateReason{}
	}
	insertXPubs := InsertTemplateXPubs(
		id,
		resource.UserUserXPubList,
		resource.UserBackupXPubList,
		resource.UserLockXPubList,
		resource.UserWitnessXPubList,
		resource.ProjectOldXPubList,
		resource.ProjectParentXPubList,
		resource.ProjectBoardXPubList,
		resource.ProjectMemberXPubList,
		resource.ProjectCosignerXPubList,
		resource.ProjectWitnessXPubList,
		dbConf,
	)

	if !insertXPubs {
		return commons.StateReason{}
	}

	return commons.StateReason{Reason: "todo"}
}

// BoardingGetDb ...
func BoardingGetDb(_type string, xPubS commons.XPub, state commons.StateReason, dbConf persistance.DbConf) commons.Template {
	template := GetTemplateByTypeXPubAndState(_type, xPubS.XPub, state.Reason, dbConf)
	templateDTO := commons.BuildTemplateDTO(template, dbConf)

	return templateDTO
}
