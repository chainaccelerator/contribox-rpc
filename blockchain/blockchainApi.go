package blockchain

import (
	"bc_node_api/api3/commons"
)

// Broadcast ...
// func Broadcast(_type string, )
// Where is type transaction

// PeerAsk ...

// PeerBlocValidation ...
func PeerBlocValidation(_type string, contributionTxId commons.TxId, state commons.StateReason) (bool, bool) {
	return true, false
}

// PeerPegValidation ...
func PeerPegValidation(_type string, contributionTxId commons.TxId, state commons.StateReason) (bool, bool) {
	return true, false
}
