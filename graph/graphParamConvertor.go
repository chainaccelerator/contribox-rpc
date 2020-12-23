package graph

import "bc_node_api/api3/commons"

// StoreParamConvert ...
func StoreParamConvert(params []interface{}) (string, commons.Trace, commons.StateReason) {
	_type := params[0].(string)

	traceParam := params[1].(map[string]interface{})

	hashParam := traceParam["hash"].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	proofMerkleTreeParam := traceParam["proofMerkleTree"].([]interface{})
	proofMerkleTree := commons.HashList(proofMerkleTreeParam)

	sigYesParam := traceParam["sigYes"].([]interface{})
	sigYes := commons.XPubList(sigYesParam)

	sigNoParam := traceParam["sigNo"].([]interface{})
	sigNo := commons.XPubList(sigNoParam)

	transactionHashParam := traceParam["transactionHash"].(map[string]interface{})
	transactionHash := commons.Hash{Hash: transactionHashParam["hash"].(string)}

	gitCommitHashParam := traceParam["gitCommitHash"].(map[string]interface{})
	gitCommitHash := commons.Hash{Hash: gitCommitHashParam["hash"].(string)}

	trace := commons.Trace{
		Hash:            hash,
		ProofMerkleTree: proofMerkleTree,
		SigYes:          sigYes,
		SigNo:           sigNo,
		TransactionHash: transactionHash,
		GitCommitHash:   gitCommitHash,
	}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, trace, state
}

// PeerAskParamConvert ...
func PeerAskParamConvert(params []interface{}) (string, commons.XPub, int, int, commons.StateReason) {
	_type := params[0].(string)

	xPubParam := params[1].(map[string]interface{})
	xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}

	depthMax := params[2].(int)
	depth := params[3].(int)

	stateParam := params[4].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, xPub, depthMax, depth, state
}
