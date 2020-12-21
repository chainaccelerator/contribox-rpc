package boarding

import (
	"bc_node_api/api3/commons"
)

// ValidateBoardingTemplateGet ...
func ValidateBoardingTemplateGet(groupRoleName string, onBoarding commons.Boarding, outBoarding commons.Boarding, hash string, state string) bool {
	return commons.ValidateGroupRoleName(groupRoleName) &&
		commons.ValidateBoarding(onBoarding) &&
		commons.ValidateBoarding(outBoarding) &&
		commons.ValidateKeyOrHash(hash) &&
		state == "done"
}

// ValidateBoarding ...
func ValidateBoarding(_type string, resource commons.Template, state string) bool {
	return _type == "default" && commons.ValidateTemplate(resource) && state == "todo"
}

// ValidateBoardingGet ...
func ValidateBoardingGet(_type string, xPubS string, state string) bool {
	return _type == "default" && commons.ValidateXPub(xPubS) && state == "todo"
}

// ValidateBoardingBroadcast ...
func ValidateBoardingBroadcast(_type string, resourceList []commons.UTXO, hash string, state string) bool {
	return _type == "default" && commons.ValidateUTXOList(resourceList) && commons.ValidateKeyOrHash(hash) && state == "done"
}

// ValidateBoardingBroadcastGet ...
func ValidateBoardingBroadcastGet(_type string, hash string, state string) bool {
	return _type == "default" && commons.ValidateKeyOrHash(hash) && state == "done"
}
