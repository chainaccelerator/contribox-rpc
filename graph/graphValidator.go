package graph

import "bc_node_api/api3/commons"

// ValidateStore ...
func ValidateStore(_type string, trace commons.Trace, state string) bool {
	return _type == "test" && validateTrace(trace) && state == "stable"
}

// ValidatePeerAsk ...
func ValidatePeerAsk(_type string, xPub string, depthMax int, depth int, state string) bool {
	return commons.ValidatePeerAsk(_type, xPub, depthMax, depth, state)
}

// ValidateHashAsk ...
func ValidateHashAsk(_type string, trace commons.Trace, state string) bool {
	return _type == "test" && validateTrace(trace) && state == "stable"
}

func validateTrace(trace commons.Trace) bool {
	return commons.ValidateKeyOrHash(trace.Hash.Hash) &&
		commons.ValidateHashList(trace.ProofMerkleTree) &&
		commons.ValidateXPubList(trace.SigYes) &&
		commons.ValidateXPubList(trace.SigNo) &&
		commons.ValidateKeyOrHash(trace.GitCommitHash.Hash)
}
