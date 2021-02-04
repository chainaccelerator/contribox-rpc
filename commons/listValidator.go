package commons

// ValidateKeyEncryptedList ...
func ValidateKeyEncryptedList(keyEncryptedList []KeyEncrypted) bool {
	for _, keyEncrypted := range keyEncryptedList {
		if !ValidateKeyEncrypted(keyEncrypted) {
			return false
		}
	}
	return true
}

// ValidatePubKeyEncryptedList ...
func ValidatePubKeyEncryptedList(pubKeyEncryptedList []PubKeyEncrypted) bool {
	for _, pubKeyEncrypted := range pubKeyEncryptedList {
		if !ValidatePubKeyEncrypted(pubKeyEncrypted) {
			return false
		}
	}
	return true
}
