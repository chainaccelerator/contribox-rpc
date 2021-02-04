package commons

// BuildKeyEncryptedList ...
func BuildKeyEncryptedList(keyEncryptedListParam []interface{}) []KeyEncrypted {
	var keyEncryptedList []KeyEncrypted
	for _, keyEncryptedParam := range keyEncryptedListParam {
		keyEncryptedList = append(keyEncryptedList, BuildKeyEncrypted(keyEncryptedParam.(map[string]interface{})))
	}
	return keyEncryptedList
}

// BuildPubKeyEncryptedList ...
func BuildPubKeyEncryptedList(pubKeyEncryptedListParam []interface{}) []PubKeyEncrypted {
	var pubKeyEncryptedList []PubKeyEncrypted
	for _, pubKeyEncryptedParam := range pubKeyEncryptedListParam {
		pubKeyEncryptedList = append(pubKeyEncryptedList, BuildPubKeyEncrypted(pubKeyEncryptedParam.(map[string]interface{})))
	}
	return pubKeyEncryptedList
}
