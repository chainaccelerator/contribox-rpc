package graph

import (
	"bc_node_api/api3/commons"
)

// Store ...
func Store(_type string, trace commons.Trace, state commons.StateReason) (bool, bool) {
	return true, false
}

// PeerAsk ...
func PeerAsk(_type string, xPub commons.XPub, depthMax int, depth int, state commons.StateReason) (commons.Trace, bool) {
	return commons.Trace{}, false
}

// HashAsk ...
// func HashAsk(_type string, trace commons.Trace, state commons.StateReason) ()
// Return type ?
