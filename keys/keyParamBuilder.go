package keys

import "bc_node_api/api3/commons"

// KeyShareBuildParam ...
func KeyShareBuildParam(params []interface{}) ([]commons.PubKeyEncrypted, []commons.PubKeyEncrypted, []commons.KeyEncrypted) {
	pubKeyEncryptedListParam := params[0].([]interface{})
	pubKeySEncryptedListParam := params[1].([]interface{})
	keyEncryptedListParam := params[2].([]interface{})

	return commons.BuildPubKeyEncryptedList(pubKeyEncryptedListParam), commons.BuildPubKeyEncryptedList(pubKeySEncryptedListParam), commons.BuildKeyEncryptedList(keyEncryptedListParam)
}

// KeyShareGetBuildParam ...
func KeyShareGetBuildParam(params []interface{}) []commons.PubKeyEncrypted {
	pubKeyEncryptedListParam := params[0].([]interface{})
	return commons.BuildPubKeyEncryptedList(pubKeyEncryptedListParam)
}

// KeyShareConfirmBuildParam ...
func KeyShareConfirmBuildParam(params []interface{}) (string, commons.XPub, commons.Hash, commons.StateReason) {
	_type := params[0].(string)

	resourceParam := params[1].(map[string]interface{})
	resource := commons.XPub{XPub: resourceParam["xPub"].(string)}

	hashParam := params[2].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[3].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, resource, hash, state
}

// KeyShareConfirmGetBuildParam ...
func KeyShareConfirmGetBuildParam(params []interface{}) (string, commons.Hash, commons.StateReason) {
	return commons.BuildGetWithHashParams(params)
}
