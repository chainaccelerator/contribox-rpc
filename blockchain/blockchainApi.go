package blockchain

import (
	"bc_node_api/api3/commons"
)

// Broadcast ...
func Broadcast(_type string, tx commons.Transaction, state commons.StateReason) (bool, bool) {
	if !ValidateBroadcast(_type, state.Reason) {
		return false, true
	}
	return true, false
}

// PeerAsk ...
func PeerAsk(_type string, txID commons.TxId, state commons.StateReason) (commons.Transaction, bool) {
	if !ValidatePeer(_type, txID.Id, state.Reason) {
		return commons.Transaction{}, true
	}
	return commons.Transaction{}, false
}

// PeerBlocValidation ...
func PeerBlocValidation(_type string, contributionTxID commons.TxId, state commons.StateReason) (bool, bool) {
	if !ValidatePeer(_type, contributionTxID.Id, state.Reason) {
		return false, true
	}
	return true, false
}

// PeerPegValidation ...
func PeerPegValidation(_type string, contributionTxID commons.TxId, state commons.StateReason) (bool, bool) {
	if !ValidatePeer(_type, contributionTxID.Id, state.Reason) {
		return false, true
	}
	return true, false
}
