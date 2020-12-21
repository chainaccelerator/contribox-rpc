package key

import "bc_node_api/api3/commons"

// As much as possible, pass only string in validation methods

// ValidateKeyShare ...
func ValidateKeyShare(_type string, xPubSList []commons.XPub, Key string, Hash string, state string) bool {
	return isCorrectType(_type) && commons.ValidateXPubList(xPubSList) && commons.ValidateKeyOrHash(Key) && commons.ValidateKeyOrHash(Hash) && state == "todo"
}

// ValidateKeyShareGet ...
func ValidateKeyShareGet(_type string, xPubS string, state string) bool {
	return isCorrectType(_type) && commons.ValidateXPub(xPubS) && state == "todo"
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
