package commons

import (
	"bc_node_api/api3/persistance"
)

// BuildProofDTO ...
func BuildProofDTO(proof persistance.Proof, dbConf persistance.DbConf) Proof {
	keyValList := persistance.GetKeyValListByProofID(proof.Id, dbConf)
	xPubList := persistance.GetXPubListByProofID(proof.Id, dbConf)
	hashList := persistance.GetHashListByProofID(proof.Id, dbConf)
	tagList := persistance.GetTagListByProofID(proof.Id, dbConf)

	proofDTO := Proof{
		ProjectName:       ProjectName{Name: proof.ProjectName},
		LicenseSPDX:       Licence{SPDX: proof.LicenceSPDX},
		LicenseSPDXChange: Licence{SPDX: proof.LicenceSPDXChange},
		GroupRoleName:     GroupRoleName{Name: proof.GroupRoleName},
	}

	for _, keyVal := range keyValList {
		keyValDTO := KeyVal{Key: keyVal.Key, Val: keyVal.Val}
		switch keyVal.KeyValType {
		case "DESCRIPTIONPUBLIC":
			proofDTO.DescriptionPublicList = append(proofDTO.DescriptionPublicList, keyValDTO)
		case "IDENTITY":
			proofDTO.IdentityList = append(proofDTO.IdentityList, keyValDTO)
		case "EVENT":
			proofDTO.EventList = append(proofDTO.EventList, keyValDTO)
		case "ENVIRONMENT":
			proofDTO.EnvironmentList = append(proofDTO.EnvironmentList, keyValDTO)
		case "QUALITY":
			proofDTO.QualityList = append(proofDTO.QualityList, keyValDTO)
		case "CONTRIBUTE":
			proofDTO.ContributeList = append(proofDTO.ContributeList, keyValDTO)
		case "ORIGIN":
			proofDTO.OriginList = append(proofDTO.OriginList, keyValDTO)
		case "NDA":
			proofDTO.NdaList = append(proofDTO.NdaList, keyValDTO)
		case "CONFIDENTIALDATA":
			proofDTO.ConfidentialDataList = append(proofDTO.ConfidentialDataList, keyValDTO)
		case "METADATA":
			proofDTO.MetaDataList = append(proofDTO.MetaDataList, keyValDTO)
		case "OFFICER":
			proofDTO.OfficerList = append(proofDTO.OfficerList, keyValDTO)
		case "EDIT":
			proofDTO.EditList = append(proofDTO.EditList, keyValDTO)
		case "CERTIFICATE":
			proofDTO.CertificateList = append(proofDTO.CertificateList, keyValDTO)
		case "EXPORTCONTROL":
			proofDTO.ExportControlList = append(proofDTO.ExportControlList, keyValDTO)
		case "KEYVALUE":
			proofDTO.KeyValueList = append(proofDTO.KeyValueList, keyValDTO)
		}
	}

	for _, hash := range hashList {
		hashDTO := Hash{Hash: hash.Hash}
		switch hash.HashType {
		case "PARENT":
			proofDTO.ParentList = append(proofDTO.ParentList, hashDTO)
		case "PREVIOUS":
			proofDTO.PreviousList = append(proofDTO.PreviousList, hashDTO)
		case "LEFT":
			proofDTO.LeftList = append(proofDTO.LeftList, hashDTO)
		}
	}

	for _, xPub := range xPubList {
		xPubDTO := XPub{XPub: xPub.XPub}
		switch xPub.XPubType {
		case "FOR":
			proofDTO.ForList = append(proofDTO.ForList, xPubDTO)
		case "TO":
			proofDTO.ToList = append(proofDTO.ToList, xPubDTO)
		}
	}

	for _, tag := range tagList {
		proofDTO.TagList = append(proofDTO.TagList, tag.Tag)
	}

	return proofDTO
}

// BuildTemplateDTO ...
func BuildTemplateDTO(template persistance.Template, dbConf persistance.DbConf) Template {
	proof := persistance.GetProofByTemplateID(template.Id, dbConf)
	proofDTO := BuildProofDTO(proof, dbConf)

	templateDTO := Template{
		Hash:               Hash{Hash: template.Hash},
		Proof:              proofDTO,
		ProjectName:        ProjectName{Name: template.ProjectName},
		LicenceSPDX:        Licence{SPDX: template.LicenceSPDX},
		GroupRoleName:      GroupRoleName{Name: template.GroupRoleName},
		UserRequirement:    template.UserRequirement,
		ProjectRequirement: template.ProjectRequirement,
		UserUser:           TemplateUser{Quorum: template.UserUser},
		UserBackup:         TemplateUser{Quorum: template.UserBackup},
		UserLock:           TemplateUser{Quorum: template.UserLock},
		UserWitness:        TemplateUser{Quorum: template.UserWitness},
		ProjectOld:         TemplateProject{Quorum: template.ProjectOld},
		ProjectParent:      TemplateProject{Quorum: template.ProjectParent},
		ProjectBoard:       TemplateProject{Quorum: template.ProjectBoard},
		ProjectMember:      TemplateProject{Quorum: template.ProjectMember},
		ProjectCosigner:    TemplateProject{Quorum: template.ProjectCosigner},
		ProjectWitness:     TemplateProject{Quorum: template.ProjectWitness},
	}

	xPubList := persistance.GetXPubListByTemplateID(template.Id, dbConf)

	for _, xPub := range xPubList {
		switch xPub.XPubType {
		case "USERUSER":
			templateDTO.UserUserXPubList = append(templateDTO.UserUserXPubList, XPub{XPub: xPub.XPub})
		case "USERBACKUP":
			templateDTO.UserBackupXPubList = append(templateDTO.UserBackupXPubList, XPub{XPub: xPub.XPub})
		case "USERLOCK":
			templateDTO.UserLockXPubList = append(templateDTO.UserLockXPubList, XPub{XPub: xPub.XPub})
		case "USERWITNESS":
			templateDTO.UserWitnessXPubList = append(templateDTO.UserWitnessXPubList, XPub{XPub: xPub.XPub})
		case "PROJECTOLD":
			templateDTO.ProjectOldXPubList = append(templateDTO.ProjectOldXPubList, XPub{XPub: xPub.XPub})
		case "PROJECTPARENT":
			templateDTO.ProjectParentXPubList = append(templateDTO.ProjectParentXPubList, XPub{XPub: xPub.XPub})
		case "PROJECTBOARD":
			templateDTO.ProjectBoardXPubList = append(templateDTO.ProjectBoardXPubList, XPub{XPub: xPub.XPub})
		case "PROJECTMEMBER":
			templateDTO.ProjectMemberXPubList = append(templateDTO.ProjectMemberXPubList, XPub{XPub: xPub.XPub})
		case "PROJECTCOSIGNER":
			templateDTO.ProjectCosignerXPubList = append(templateDTO.ProjectCosignerXPubList, XPub{XPub: xPub.XPub})
		case "PROJECTWITNESS":
			templateDTO.ProjectWitnessXPubList = append(templateDTO.ProjectWitnessXPubList, XPub{XPub: xPub.XPub})
		}
	}

	return templateDTO
}
