package git

import (
	"bc_node_api/api3/commons"
)

// StoreBuildParam ...
func StoreBuildParam(params []interface{}) (string, string, commons.XPub, commons.StateReason) {
	_type := params[0].(string)

	resourceEncrypted := params[1].(string)

	xPubParam := params[2].(map[string]interface{})
	xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}

	stateParam := params[3].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, resourceEncrypted, xPub, state
}

// PeerAskBuildParam ...
func PeerAskBuildParam(params []interface{}) (string, commons.XPub, int, int, commons.StateReason) {
	_type := params[0].(string)

	xPubParam := params[1].(map[string]interface{})
	xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}

	depthMax := params[2].(int)
	depth := params[3].(int)

	stateParam := params[4].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, xPub, depthMax, depth, state
}

// CommitHashAskBuildParam ...
func CommitHashAskBuildParam(params []interface{}) (string, commons.Hash, commons.XPub, commons.StateReason) {
	_type := params[0].(string)

	hashParam := params[1].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	xPubParam := params[2].(map[string]interface{})
	xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}

	stateParam := params[3].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, hash, xPub, state
}
