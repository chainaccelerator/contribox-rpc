package key

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// Service methods must return primitive types

// KeyShare ...
// Insert a share request
func KeyShare(_type string, keyShared commons.KeyShared, state commons.StateReason, dbConf persistance.DbConf) (string, bool) {
	if !ValidateKeyShare(_type, keyShared.XPubSList, keyShared.Key.Key, keyShared.Hash.Hash, state.Reason) {
		return "", true
	}
	var xPubStringList []string
	for _, xPubS := range keyShared.XPubSList {
		xPubStringList = append(xPubStringList, xPubS.XPub)
	}

	keyShare := KeyShareDb(_type, xPubStringList, keyShared.Key.Key, keyShared.Hash.Hash, state.Reason, dbConf)
	return keyShare, keyShare == ""
}

// KeyShareGet ...
// Check the existance of a share request
func KeyShareGet(_type string, xPubS commons.XPub, state commons.StateReason, dbConf persistance.DbConf) (commons.Key, bool) {
	if !ValidateKeyShareGet(_type, xPubS.XPub, state.Reason) {
		return commons.Key{Key: ""}, true
	}
	keyShareGet := KeyShareGetDb(_type, xPubS.XPub, state.Reason, dbConf)
	return commons.Key{Key: keyShareGet}, keyShareGet == ""
}

// KeyShareConfirm ...
// Approve a share request
func KeyShareConfirm(_type string, xPubS commons.XPub, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (string, bool) {
	if !ValidateKeyShareConfirm(_type, xPubS.XPub, hash.Hash, state.Reason) {
		return "", true
	}
	keyShareConfirm := KeyShareConfirmDb(_type, xPubS.XPub, hash.Hash, state.Reason, dbConf)
	return keyShareConfirm, keyShareConfirm == ""
}

// KeyShareConfirmGet ...
// Check the approval of a share request
func KeyShareConfirmGet(_type string, hash commons.Hash, state commons.StateReason, dbConf persistance.DbConf) (commons.XPub, bool) {
	if !ValidateKeyShareConfirmGet(_type, hash.Hash, state.Reason) {
		return commons.XPub{XPub: ""}, true
	}
	keyShareConfirmGet := KeyShareConfirmGetDb(_type, hash.Hash, state.Reason, dbConf)
	return commons.XPub{XPub: keyShareConfirmGet}, keyShareConfirmGet == ""
}
