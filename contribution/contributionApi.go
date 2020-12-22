package contribution

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// Contribution ...
func Contribution(_type string, contribution commons.Contribution, state commons.StateReason, dbConf persistance.DbConf) (commons.StateReason, bool) {
	if !ValidateContribution(_type, contribution, state.Reason) {
		return commons.StateReason{}, true
	}
	return commons.StateReason{}, false
}

// ContributionGet ...
func ContributionGet(_type string, xPubS commons.XPub, state commons.StateReason, dbConf persistance.DbConf) (commons.Contribution, bool) {
	if !ValidateContributionGet(_type, xPubS.XPub, state.Reason) {
		return commons.Contribution{}, true
	}
	return commons.Contribution{}, false
}

// ContributionConfirm ...
func ContributionConfirm(_type string, sig commons.Sig, hash commons.Hash, xPub commons.XPub, resourceEncrypted string, state commons.StateReason, dbConf persistance.DbConf) (commons.Sig, bool) {
	if !ValidateContributionConfirm(_type, sig, hash.Hash, xPub.XPub, state.Reason) {
		return commons.Sig{}, true
	}
	return commons.Sig{}, false
}

// ContributionConfirmGet ...
func ContributionConfirmGet(_type string, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.Sig, bool) {
	if !ValidateContributionConfirmGet(_type, hash.Hash, state.Reason) {
		return commons.Sig{}, true
	}
	return commons.Sig{}, false
}

// ContributionBroadcast ...
func ContributionBroadcast(_type string, resourceList []commons.Contribution, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.StateReason, bool) {
	if !ValidateContributionBroadcast(_type, resourceList, hash.Hash, state.Reason) {
		return commons.StateReason{}, true
	}
	return commons.StateReason{}, false
}

// ContributionBroadcastGet ...
func ContributionBroadcastGet(_type string, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.Contribution, bool) {
	if !ValidateContributionBroadcastGet(_type, hash.Hash, state.Reason) {
		return commons.Contribution{}, true
	}
	return commons.Contribution{}, false
}
