package blockchain

import "bc_node_api/api3/commons"

// ValidateBroadcast ...
func ValidateBroadcast(_type string, state string) bool {
	return _type == "test" && state == "stable"
}

// ValidatePeer ...
func ValidatePeer(_type string, txID string, state string) bool {
	return _type == "test" && commons.ValidateTX(txID) && state == "stable"
}
