package key

import "bc_node_api/api3/commons"

// KeyShareParamConvert ...
func KeyShareParamConvert(params ...interface{}) (string, commons.KeyShared, commons.StateReason) {
	_type := params[0].(string)

	keySharedParam := params[1].(map[string]interface{})
	xPubSListParam := keySharedParam["xPubSList"].([]interface{})
	xPubSList := commons.XPubList(xPubSListParam)

	keyShared := commons.KeyShared{
		XPubSList: xPubSList,
		Key:       commons.Key{Key: keySharedParam["key"].(string)},
		Hash:      commons.Hash{Hash: keySharedParam["hash"].(string)},
	}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, keyShared, state
}

// KeyShareGetParamConvert ...
func KeyShareGetParamConvert(params ...interface{}) (string, commons.XPub, commons.StateReason) {
	_type := params[0].(string)

	xPubSParam := params[1].(map[string]interface{})
	xPubS := commons.XPub{XPub: xPubSParam["xPub"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, xPubS, state
}

// KeyShareConfirmParamConvert ...
func KeyShareConfirmParamConvert(params ...interface{}) (string, commons.XPub, commons.Hash, commons.StateReason) {
	_type := params[0].(string)

	resourceParam := params[1].(map[string]interface{})
	resource := commons.XPub{XPub: resourceParam["xPub"].(string)}

	hashParam := params[2].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[3].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, resource, hash, state
}

// KeyShareConfirmGetParamConvert ...
func KeyShareConfirmGetParamConvert(params ...interface{}) (string, commons.Hash, commons.StateReason) {
	_type := params[0].(string)

	hashParam := params[1].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, hash, state
}
