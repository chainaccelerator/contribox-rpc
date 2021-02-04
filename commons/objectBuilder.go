package commons

import "bc_node_api/api3/hashes"

// BuildKey ...
func BuildKey(keyParam map[string]interface{}) Key {
	hashParam := keyParam["hash"].(map[string]interface{})
	return Key{
		Data:  keyParam["data"].(string),
		Hash:  hashes.BuildHash(hashParam),
		State: keyParam["state"].(string),
	}
}

// BuildKeyEncrypted ...
func BuildKeyEncrypted(keyEncryptedParam map[string]interface{}) KeyEncrypted {
	keyParam := keyEncryptedParam["key"].(map[string]interface{})
	hashParam := keyEncryptedParam["hash"].(map[string]interface{})
	return KeyEncrypted{
		Key:       BuildKey(keyParam),
		Encrypted: keyEncryptedParam["encrypted"].(string),
		Hash:      hashes.BuildHash(hashParam),
		State:     keyEncryptedParam["state"].(string),
	}
}

// BuildPubKeyEncrypted ...
func BuildPubKeyEncrypted(pubKeyEncryptedParam map[string]interface{}) PubKeyEncrypted {
	hashParam := pubKeyEncryptedParam["hash"].(map[string]interface{})
	return PubKeyEncrypted{
		Encrypted: pubKeyEncryptedParam["encrypted"].(string),
		Hash:      hashes.BuildHash(hashParam),
		State:     pubKeyEncryptedParam["state"].(string),
	}
}
