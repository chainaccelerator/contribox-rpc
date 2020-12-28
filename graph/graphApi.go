package graph

import (
	"bc_node_api/api3/commons"
)

// Store ...
func Store(_type string, trace commons.Trace, state commons.StateReason) (bool, bool) {
	if !ValidateStore(_type, trace, state.Reason) {
		return false, true
	}
	return true, false
}

// PeerAsk ...
func PeerAsk(_type string, xPub commons.XPub, depthMax int, depth int, state commons.StateReason) (commons.Trace, bool) {
	if !ValidatePeerAsk(_type, xPub.XPub, depthMax, depth, state.Reason) {
		return commons.Trace{}, true
	}
	return commons.Trace{}, false
}

// HashAsk ...
func HashAsk(_type string, trace commons.Trace, state commons.StateReason) (bool, bool) {
	if !ValidateHashAsk(_type, trace, state.Reason) {
		return false, true
	}
	return true, false
}
