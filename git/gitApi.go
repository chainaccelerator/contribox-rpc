package git

import (
	"bc_node_api/api3/commons"
)

// Store ...
func Store(_type string, resourceEncrypted string, xPub commons.XPub, state commons.StateReason) (commons.Hash, bool) {
	return commons.Hash{}, false
}

// PeerAsk ...
func PeerAsk(_type string, xPub commons.XPub, depthMax int, depth int, state commons.StateReason) (commons.IPv4, bool) {
	return commons.IPv4{}, false
}

// CommitHashAsk ...
func CommitHashAsk(_type string, hash commons.Hash, xPub commons.XPub, state commons.StateReason) (string, bool) {
	return "", false
}
