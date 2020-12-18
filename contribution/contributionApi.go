package contribution

import (
	"bc_node_api/api3/commons"
)

// Contribution ...
func Contribution(_type string, contribution commons.Contribution, state commons.StateReason, dbConf commons.DbConf) (commons.StateReason, bool) {
	return commons.StateReason{}, true
}

// ContributionGet ...
func ContributionGet(_type string, xPubS commons.XPub, state commons.StateReason, dbConf commons.DbConf) (commons.Contribution, bool) {
	return commons.Contribution{}, true
}

// ContributionConfirm ...
func ContributionConfirm(_type string, sig commons.Sig, hash commons.Hash, state commons.StateReason, dbConf commons.DbConf) (commons.Sig, bool) {
	return commons.Sig{}, true
}

// ContributionConfirmGet ...
func ContributionConfirmGet(_type string, hash commons.Hash, state commons.StateReason, dbConf commons.DbConf) (commons.Sig, bool) {
	return commons.Sig{}, true
}

// ContributionBroadcast ...
func ContributionBroadcast(_type string, resourceList []commons.Contribution, hash commons.Hash, state commons.StateReason, dbConf commons.DbConf) (commons.StateReason, bool) {
	return commons.StateReason{}, true
}

// ContributionBroadcastGet ...
func ContributionBroadcastGet(_type string, hash commons.Hash, state commons.StateReason, dbConf commons.DbConf) (commons.Contribution, bool) {
	return commons.Contribution{}, true
}
