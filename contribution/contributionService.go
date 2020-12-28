package contribution

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// ContributionDb ...
func ContributionDb(_type string, fullContribution commons.Contribution, state string, dbConf persistance.DbConf) commons.StateReason {
	contribution := persistance.Contribution{
		Hash:            fullContribution.Hash.Hash,
		XPub:            fullContribution.XPub.XPub,
		Tx1Id:           fullContribution.Tx1Id.Id,
		Tx0IdAmount:     fullContribution.Tx0IdAmount,
		Tx0IdIssueAsset: fullContribution.Tx0IdIssueAsset.IssueAsset.Hash,
		Tx0IdSigA:       fullContribution.Tx0IdSigA.Sig,
	}

	contributionID := InsertContribution(contribution, dbConf)
	if contributionID == 0 {
		return commons.StateReason{}
	}

	return commons.StateReason{}
}

// ContributionGetDb ...
func ContributionGetDb(_type string, xPubS string, state string, dbConf persistance.DbConf) commons.FullContribution {
	contribution := GetContributionByXPubAndState(_type, xPubS, state, dbConf)
	fullContributionDTO := buildContributionDTO(contribution, dbConf)
	return fullContributionDTO
}

// ContributionBroadcastGetDb ...
func ContributionBroadcastGetDb(_type string, hash string, state string, dbConf persistance.DbConf) commons.FullContribution {
	contribution := GetContributionByHashAndState(_type, hash, state, dbConf)
	fullContributionDTO := buildContributionDTO(contribution, dbConf)
	return fullContributionDTO
}

func buildContributionDTO(contribution persistance.Contribution, dbConf persistance.DbConf) commons.FullContribution {
	proof := GetProofByContributionID(contribution.Id, dbConf)
	proofDTO := commons.BuildProofDTO(proof, dbConf)

	contributionDTO := commons.Contribution{
		Hash:            commons.Hash{Hash: contribution.Hash},
		Proof:           proofDTO,
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

	template := GetTemplateByContributionID(contribution.Id, dbConf)
	templateDTO := commons.BuildTemplateDTO(template, dbConf)

	fullContributionDTO := commons.FullContribution{
		Contribution: contributionDTO,
		Template:     templateDTO,
	}

	return fullContributionDTO
}
