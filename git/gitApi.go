package git

import (
	"bc_node_api/api3/commons"
)

// Store ...
func Store(_type string, resourceEncrypted string, xPub commons.XPub, state commons.StateReason) (commons.Hash, bool) {
	if !ValidateStore(_type, xPub.XPub, state.Reason) {
		return commons.Hash{}, true
	}
	return commons.Hash{}, false
}

// PeerAsk ...
func PeerAsk(_type string, xPub commons.XPub, depthMax int, depth int, state commons.StateReason) (commons.IPv4, bool) {
	if !ValidatePeerAsk(_type, xPub.XPub, depthMax, depth, state.Reason) {
		return commons.IPv4{}, true
	}
	return commons.IPv4{}, false
}

// CommitHashAsk ...
func CommitHashAsk(_type string, hash commons.Hash, xPub commons.XPub, state commons.StateReason) (string, bool) {
	if !ValidateCommitHashAsk(_type, hash.Hash, xPub.XPub, state.Reason) {
		return "", true
	}
	return "", false
}
