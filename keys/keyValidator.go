package keys

import "bc_node_api/api3/commons"

// ValidateKeyShare ...
func ValidateKeyShare(pubKeyEncryptedList []commons.PubKeyEncrypted, pubKeySEncryptedList []commons.PubKeyEncrypted, keyEncryptedList []commons.KeyEncrypted) bool {
	return commons.ValidatePubKeyEncryptedList(pubKeyEncryptedList) &&
		commons.ValidatePubKeyEncryptedList(pubKeySEncryptedList) &&
		commons.ValidateKeyEncryptedList(keyEncryptedList) &&
		len(pubKeySEncryptedList) == len(keyEncryptedList)
}

// ValidateKeyShareGet ...
func ValidateKeyShareGet(pubKeyEncryptedList []commons.PubKeyEncrypted) bool {
	return commons.ValidatePubKeyEncryptedList(pubKeyEncryptedList)
}

// ValidateKeyShareConfirm ...
func ValidateKeyShareConfirm(_type string, xPubS string, Hash string, state string) bool {
	return isCorrectType(_type) && commons.ValidateXPub(xPubS) && commons.ValidateKeyOrHash(Hash) && state == "done"
}

// ValidateKeyShareConfirmGet ...
func ValidateKeyShareConfirmGet(_type string, hash string, state string) bool {
	return isCorrectType(_type) && commons.ValidateKeyOrHash(hash) && state == "done"
}

func isCorrectType(_type string) bool {
	return _type == "share" || _type == "backup" || _type == "lock"
}
