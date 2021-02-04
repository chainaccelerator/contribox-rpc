package keys

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/hashes"
	"bc_node_api/api3/persistance"
)

// Service methods must return primitive types

// KeyShare ...
// Insert a share request
func KeyShare(pubKeyEncryptedList []commons.PubKeyEncrypted, pubKeySEncryptedList []commons.PubKeyEncrypted, keyEncryptedList []commons.KeyEncrypted, metadata commons.Metadata, dbConf persistance.DbConf) (KeyShareResponseBody, bool) {
	if !ValidateKeyShare(pubKeyEncryptedList, pubKeySEncryptedList, keyEncryptedList) || !commons.ValidateMetadata(metadata) {
		return KeyShareResponseBody{}, true
	}

	keyShared := KeyShareDb(pubKeyEncryptedList, pubKeySEncryptedList, keyEncryptedList, metadata, dbConf)
	if !keyShared {
		return KeyShareResponseBody{}, true
	}

	success := KeyShareResponseBody{Hash: metadata.Hash}
	return success, false
}

// KeyShareGet ...
// Check the existance of a share request
func KeyShareGet(pubKeyEncryptedList []commons.PubKeyEncrypted, metadata commons.Metadata, dbConf persistance.DbConf) (KeyShareGetResponseBody, bool) {
	if !ValidateKeyShareGet(pubKeyEncryptedList) || !commons.ValidateMetadata(metadata) {
		return KeyShareGetResponseBody{}, true
	}

	// TODO : process data
	success := KeyShareGetResponseBody{
		Hash: hashes.Hash{},
	}
	return success, false
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
