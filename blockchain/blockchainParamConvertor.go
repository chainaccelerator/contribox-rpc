package blockchain

import (
	"bc_node_api/api3/commons"
)

// PeerValidationParamConvert ...
func PeerValidationParamConvert(params []interface{}) (string, commons.TxId, commons.StateReason) {
	_type := params[0].(string)

	contributionTxIDParam := params[1].(map[string]interface{})
	contributionTxID := commons.TxId{Id: contributionTxIDParam["id"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, contributionTxID, state
}
