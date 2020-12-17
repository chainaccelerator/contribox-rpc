package key

import "bc_node_api/api3/commons"

// As much as possible, pass only string in validation methods

// ValidateKeyShare ...
func ValidateKeyShare(_type string, xPubSList []commons.XPub, Key string, Hash string, state string) bool {
	return isCorrectType(_type) && isCorrectXPubSList(xPubSList) && isCorrectKeyOrHash(Key) && isCorrectKeyOrHash(Hash) && state == "todo"
}

// ValidateKeyShareGet ...
func ValidateKeyShareGet(_type string, xPubS string, state string) bool {
	return isCorrectType(_type) && isCorrectXPubS(xPubS) && state == "todo"
}

// ValidateKeyShareConfirm ...
func ValidateKeyShareConfirm(_type string, xPubS string, Hash string, state string) bool {
	return isCorrectType(_type) && isCorrectXPubS(xPubS) && isCorrectKeyOrHash(Hash) && state == "done"
}

// ValidateKeyShareConfirmGet ...
func ValidateKeyShareConfirmGet(_type string, hash string, state string) bool {
	return isCorrectType(_type) && isCorrectKeyOrHash(hash) && state == "done"
}

func isCorrectType(_type string) bool {
	return _type == "share" || _type == "backup" || _type == "lock"
}

func isCorrectXPubS(xPubS string) bool {
	return len(xPubS) <= 112
}

func isCorrectXPubSList(xPubSList []commons.XPub) bool {
	for _, xPubS := range xPubSList {
		if !isCorrectXPubS(xPubS.XPub) {
			return false
		}
	}
	return true
}

func isCorrectKeyOrHash(keyOrHash string) bool {
	return len(keyOrHash) == 64
}
