package git

import "bc_node_api/api3/commons"

// ValidateStore ...
func ValidateStore(_type string, xPub string, state string) bool {
	return _type == "test" && commons.ValidateXPub(xPub) && state == "stable"
}

// ValidatePeerAsk ...
func ValidatePeerAsk(_type string, xPub string, depthMax int, depth int, state string) bool {
	return commons.ValidatePeerAsk(_type, xPub, depthMax, depth, state)
}

// ValidateCommitHashAsk ...
func ValidateCommitHashAsk(_type string, hash string, xPub string, state string) bool {
	return _type == "test" && commons.ValidateKeyOrHash(hash) && commons.ValidateXPub(xPub) && state == "stable"
}
