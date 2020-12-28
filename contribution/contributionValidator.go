package contribution

import (
	"bc_node_api/api3/commons"
)

// ValidateContribution ...
func ValidateContribution(_type string, contribution commons.FullContribution, state string) bool {
	return _type == "default" && validateContribution(contribution) && state == "todo"
}

func validateContribution(contribution commons.FullContribution) bool {
	return commons.ValidateKeyOrHash(contribution.Hash.Hash) &&
		commons.ValidateProof(contribution.Proof) &&
		validateBlindKeyList(contribution.BlindKeyList) &&
		// validateRangeList(contribution.RangeList) &&
		commons.ValidateBoarding(contribution.OnBoarding) &&
		commons.ValidateBoarding(contribution.OutBoarding) &&
		commons.ValidateTX(contribution.Tx1Id.Id) &&
		commons.ValidateKeyOrHash(contribution.Tx0IdIssueAsset.IssueAsset.Hash) &&
		commons.ValidateSigData(contribution.Tx0IdSigA.Sig) &&
		commons.ValidatePubKey(contribution.Vout0PubKA.PubKey) &&
		commons.ValidatePubKey(contribution.Vout1PubKS.PubKey)

}

func validateBlindKeyList(blindKeyList []commons.BlindingKeyEncrypted) bool {
	for _, blindKey := range blindKeyList {
		if !commons.ValidateXPub(blindKey.XPub.XPub) || len(blindKey.String) != 256 {
			return false
		}
	}
	return true
}

func validateRangeList(rangeList []commons.RangeEncrypted) bool {
	for _, _range := range rangeList {
		if !commons.ValidateXPub(_range.XPub.XPub) || len(_range.String) != 256 {
			return false
		}
	}
	return true
}

// ValidateContributionGet ...
func ValidateContributionGet(_type string, xPubS string, state string) bool {
	return _type == "default" && commons.ValidateXPub(xPubS) && state == "todo"
}

// ValidateContributionConfirm ...
func ValidateContributionConfirm(_type string, sig commons.Sig, hash string, xPub string, state string) bool {
	return _type == "default" && commons.ValidateSig(sig) && commons.ValidateKeyOrHash(hash) && commons.ValidateXPub(xPub) && state == "done"
}

// ValidateContributionConfirmGet ...
func ValidateContributionConfirmGet(_type string, hash string, state string) bool {
	return _type == "default" && commons.ValidateKeyOrHash(hash) && state == "done"
}

// ValidateContributionBroadcast ...
func ValidateContributionBroadcast(_type string, resourceList []commons.FullContribution, hash string, state string) bool {
	return _type == "default" && validateContributionList(resourceList) && commons.ValidateKeyOrHash(hash) && state == "done"
}

func validateContributionList(contributionList []commons.FullContribution) bool {
	for _, contribution := range contributionList {
		if !validateContribution(contribution) {
			return false
		}
	}
	return true
}

// ValidateContributionBroadcastGet ...
func ValidateContributionBroadcastGet(_type string, hash string, state string) bool {
	return _type == "default" && commons.ValidateKeyOrHash(hash) && state == "done"
}
