package contribution

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// Contribution ...
func Contribution(_type string, contribution commons.FullContribution, state commons.StateReason, dbConf persistance.DbConf) (commons.StateReason, bool) {
	if !ValidateContribution(_type, contribution, state.Reason) {
		return commons.StateReason{}, true
	}
	return commons.StateReason{Reason: "todo"}, false
}

// ContributionGet ...
func ContributionGet(_type string, xPubS commons.XPub, state commons.StateReason, dbConf persistance.DbConf) (commons.FullContribution, bool) {
	if !ValidateContributionGet(_type, xPubS.XPub, state.Reason) {
		return commons.FullContribution{}, true
	}
	return ContributionGetDb(_type, xPubS.XPub, state.Reason, dbConf), false
}

// ContributionConfirm ...
func ContributionConfirm(_type string, sig commons.Sig, hash commons.Hash, xPub commons.XPub, resourceEncrypted string, state commons.StateReason, dbConf persistance.DbConf) (commons.Sig, bool) {
	if !ValidateContributionConfirm(_type, sig, hash.Hash, xPub.XPub, state.Reason) {
		return commons.Sig{}, true
	}
	return sig, false
}

// ContributionConfirmGet ...
func ContributionConfirmGet(_type string, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.Sig, bool) {
	if !ValidateContributionConfirmGet(_type, hash.Hash, state.Reason) {
		return commons.Sig{}, true
	}
	return commons.Sig{Sig: commons.SigData{Sig: "mock signature"}, XPub: commons.XPub{XPub: "mock xPub"}, XPubS: commons.XPub{XPub: "mock xPub signer"}}, false
}

// ContributionBroadcast ...
func ContributionBroadcast(_type string, resourceList []commons.FullContribution, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.StateReason, bool) {
	if !ValidateContributionBroadcast(_type, resourceList, hash.Hash, state.Reason) {
		return commons.StateReason{}, true
	}
	return commons.StateReason{Reason: "todo"}, false
}

// ContributionBroadcastGet ...
func ContributionBroadcastGet(_type string, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.FullContribution, bool) {
	if !ValidateContributionBroadcastGet(_type, hash.Hash, state.Reason) {
		return commons.FullContribution{}, true
	}
	return ContributionBroadcastGetDb(_type, hash.Hash, state.Reason, dbConf), false
}
