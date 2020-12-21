package commons

import (
	"bc_node_api/api3/persistance"
)

// BuildProofDTO ...
func BuildProofDTO(proof persistance.Proof, keyValList []persistance.KeyVal, hashList []persistance.Hash, xPubList []persistance.XPub) Proof {
	proofDTO := Proof{
		ProjectName:       ProjectName{Name: proof.ProjectName},
		LicenseSPDX:       Licence{SPDX: proof.LicenseSPDX},
		LicenseSPDXChange: Licence{SPDX: proof.LicenseSPDXChange},
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

	return proofDTO
}
