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
	dbConf commons.DbConf,
) commons.Template {

	template := GetTemplate(projectName.Name, licenceSPDX.SPDX, groupRoleName.Name, state.Reason, dbConf)
	proof := GetProofByID(template.Id, dbConf)
	xPubList := GetXPubListByTemplateID(template.Id, dbConf)
	templateDTO := buildTemplateDTO(template, proof, xPubList)

	return templateDTO
}

// BoardingGetDb ...
func BoardingGetDb(_type string, xPubS commons.XPub, state commons.StateReason, dbConf commons.DbConf) commons.Template {
	template := GetTemplateByTypeXPubAndState(_type, xPubS.XPub, state.Reason, dbConf)
	proof := GetProofByID(template.Id, dbConf)
	xPubList := GetXPubListByTemplateID(template.Id, dbConf)
	templateDTO := buildTemplateDTO(template, proof, xPubList)

	return templateDTO
}

func buildTemplateDTO(template persistance.Template, proof persistance.Proof, xPubList []persistance.XPub) commons.Template {
	proofDTO := commons.Proof{
		ProjectName:       commons.ProjectName{Name: proof.ProjectName},
		LicenseSPDX:       commons.Licence{SPDX: proof.LicenseSPDX},
		LicenseSPDXChange: commons.Licence{SPDX: proof.LicenseSPDXChange},
		GroupRoleName:     commons.GroupRoleName{Name: proof.GroupRoleName},
	}

	templateDTO := commons.Template{
		Hash:               commons.Hash{Hash: template.Hash},
		Proof:              proofDTO,
		ProjectName:        commons.ProjectName{Name: template.ProjectName},
		LicenceSPDX:        commons.Licence{SPDX: template.LicenceSPDX},
		GroupRoleName:      commons.GroupRoleName{Name: template.GroupRoleName},
		UserRequirement:    template.UserRequirement,
		ProjectRequirement: template.ProjectRequirement,
		UserUser:           commons.TemplateUser{Quorum: template.UserUser},
		UserBackup:         commons.TemplateUser{Quorum: template.UserBackup},
		UserLock:           commons.TemplateUser{Quorum: template.UserLock},
		UserWitness:        commons.TemplateUser{Quorum: template.UserWitness},
		ProjectOld:         commons.TemplateProject{Quorum: template.ProjectOld},
		ProjectParent:      commons.TemplateProject{Quorum: template.ProjectParent},
		ProjectBoard:       commons.TemplateProject{Quorum: template.ProjectBoard},
		ProjectMember:      commons.TemplateProject{Quorum: template.ProjectMember},
		ProjectCosigner:    commons.TemplateProject{Quorum: template.ProjectCosigner},
		ProjectWitness:     commons.TemplateProject{Quorum: template.ProjectWitness},
	}

	for _, xPub := range xPubList {
		switch xPub.XPubType {
		case "USERUSER":
			templateDTO.UserUserXPubList = append(templateDTO.UserUserXPubList, commons.XPub{XPub: xPub.XPub})
		case "USERBACKUP":
			templateDTO.UserBackupXPubList = append(templateDTO.UserBackupXPubList, commons.XPub{XPub: xPub.XPub})
		case "USERLOCK":
			templateDTO.UserLockXPubList = append(templateDTO.UserLockXPubList, commons.XPub{XPub: xPub.XPub})
		case "USERWITNESS":
			templateDTO.UserWitnessXPubList = append(templateDTO.UserWitnessXPubList, commons.XPub{XPub: xPub.XPub})
		case "PROJECTOLD":
			templateDTO.ProjectOldXPubList = append(templateDTO.ProjectOldXPubList, commons.XPub{XPub: xPub.XPub})
		case "PROJECTPARENT":
			templateDTO.ProjectParentXPubList = append(templateDTO.ProjectParentXPubList, commons.XPub{XPub: xPub.XPub})
		case "PROJECTBOARD":
			templateDTO.ProjectBoardXPubList = append(templateDTO.ProjectBoardXPubList, commons.XPub{XPub: xPub.XPub})
		case "PROJECTMEMBER":
			templateDTO.ProjectMemberXPubList = append(templateDTO.ProjectMemberXPubList, commons.XPub{XPub: xPub.XPub})
		case "PROJECTCOSIGNER":
			templateDTO.ProjectCosignerXPubList = append(templateDTO.ProjectCosignerXPubList, commons.XPub{XPub: xPub.XPub})
		case "PROJECTWITNESS":
			templateDTO.ProjectWitnessXPubList = append(templateDTO.ProjectWitnessXPubList, commons.XPub{XPub: xPub.XPub})
		}
	}

	return templateDTO
}
