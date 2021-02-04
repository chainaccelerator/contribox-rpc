package blockchain

import (
	"bc_node_api/api3/commons"
)

// BroadcastBuildParam ...
func BroadcastBuildParam(params []interface{}) (string, commons.Transaction, commons.StateReason) {
	_type := params[0].(string)

	transactionParam := params[1].(map[string]interface{})
	transaction := commons.Transaction{Transaction: transactionParam["transaction"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, transaction, state
}

// PeerBuildParam ...
func PeerBuildParam(params []interface{}) (string, commons.TxId, commons.StateReason) {
	_type := params[0].(string)

	contributionTxIDParam := params[1].(map[string]interface{})
	contributionTxID := commons.TxId{Id: contributionTxIDParam["id"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, contributionTxID, state
}
