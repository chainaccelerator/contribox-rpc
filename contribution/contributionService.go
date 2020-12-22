package contribution

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// ContributionDb ...
func ContributionDb(_type string, contribution commons.Contribution, state string, dbConf persistance.DbConf) commons.StateReason {
	return commons.StateReason{}
}

// ContributionGetDb ...
func ContributionGetDb(_type string, xPubS string, state string, dbConf persistance.DbConf) commons.Contribution {
	contribution := GetContributionByXPub(_type, xPubS, state, dbConf)

	proof := GetProofByContributionID(contribution.Id, dbConf)
	proofDTO := commons.BuildProofDTO(proof, dbConf)

	template := GetTemplateByContributionID(contribution.Id, dbConf)
	templateDTO := commons.BuildTemplateDTO(template, dbConf)

	contributionDTO := commons.Contribution{
		Hash:            commons.Hash{Hash: contribution.Hash},
		Proof:           proofDTO,
		Template:        templateDTO,
		Tx1Id:           commons.TxId{Id: contribution.Tx1Id},
		Tx0IdAmount:     contribution.Tx0IdAmount,
		Tx0IdIssueAsset: commons.IssueAsset{IssueAsset: commons.Hash{Hash: contribution.Tx0IdIssueAsset}},
		Tx0IdSigA:       commons.SigData{Sig: contribution.Tx0IdSigA},
	}

	pubKeys := GetPubKeysByContributionID(contribution.Id, dbConf)
	for _, pubKey := range pubKeys {
		pubKeyDTO := commons.PubKey{PubKey: pubKey.PubKey, Base58Encoded: pubKey.Base58Encoded}
		if pubKey.PubKeyType == "VOUT0PUBKA" {
			contributionDTO.Vout0PubKA = pubKeyDTO
		} else {
			contributionDTO.Vout1PubKS = pubKeyDTO
		}
	}

	return contributionDTO
}
