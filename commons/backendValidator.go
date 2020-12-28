package commons

// ValidatePeerAsk ...
func ValidatePeerAsk(_type string, xPub string, depthMax int, depth int, state string) bool {
	return _type == "test" && ValidateXPub(xPub) && validateDepth(depthMax) && validateDepth(depth) && state == "stable"
}

func validateDepth(depth int) bool {
	return depth >= 1 && depth <= 3
}
