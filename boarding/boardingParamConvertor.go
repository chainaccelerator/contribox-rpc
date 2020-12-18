package boarding

import "bc_node_api/api3/commons"

// BoardingTemplateGetParamConvert ...
func BoardingTemplateGetParamConvert(params []interface{}) (commons.ProjectName, commons.Licence, commons.GroupRoleName, commons.Boarding, commons.Boarding, commons.Hash, commons.StateReason) {
	projectNameParam := params[0].(map[string]interface{})
	projectName := commons.ProjectName{Name: projectNameParam["name"].(string)}

	licenceSPDXParam := params[1].(map[string]interface{})
	licenceSPDX := commons.Licence{SPDX: licenceSPDXParam["licenceSPDX"].(string)}

	groupRoleNameParam := params[2].(map[string]interface{})
	groupRoleName := commons.GroupRoleName{Name: groupRoleNameParam["groupRoleName"].(string)}

	onBoardingParam := params[3].(map[string]interface{})
	onBoarding := commons.BuildBoarding(onBoardingParam)

	outBoardingParam := params[4].(map[string]interface{})
	outBoarding := commons.BuildBoarding(outBoardingParam)

	hashParam := params[5].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[6].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return projectName, licenceSPDX, groupRoleName, onBoarding, outBoarding, hash, state
}

// BoardingParamConvert ...
func BoardingParamConvert(params []interface{}) (string, commons.Template, commons.StateReason) {
	_type := params[0].(string)

	templateParam := params[1].(map[string]interface{})
	template := commons.BuildTemplate(templateParam)

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, template, state
}

// BoardingGetParamConvert ...
func BoardingGetParamConvert(params []interface{}) (string, commons.XPub, commons.StateReason) {
	return commons.BuildGetWithXPubParams(params)
}

// BoardingBroadcastParamConvert ...
func BoardingBroadcastParamConvert(params []interface{}) (string, []commons.UTXO, commons.Hash, commons.StateReason) {
	_type := params[0].(string)

	resourceListParam := params[1].([]interface{})
	var resourceList []commons.UTXO
	for _, resourceParam := range resourceListParam {
		utxoParam := resourceParam.(map[string]interface{})

		hashParam := utxoParam["hash"].(map[string]interface{})
		hash := commons.Hash{Hash: hashParam["hash"].(string)}

		txIdParam := utxoParam["Tx0Id"].(map[string]interface{})
		txId := commons.TxId{Id: txIdParam["id"].(string)}

		UTXODataParam := utxoParam["UTXO"].(map[string]interface{})
		UTXOData := commons.UTXOData{UTXO: UTXODataParam["UTXO"].(string)}

		scriptParam := utxoParam["script"].(map[string]interface{})
		script := commons.Script{Script: scriptParam["script"].(string)}

		utxo := commons.UTXO{
			Hash:   hash,
			Tx0Id:  txId,
			UTXO:   UTXOData,
			Script: script,
		}

		resourceList = append(resourceList, utxo)
	}

	hashParam := params[2].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[3].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, resourceList, hash, state
}

// BoardingBroadcastGetParamConvert ...
func BoardingBroadcastGetParamConvert(params []interface{}) (string, commons.Hash, commons.StateReason) {
	return commons.BuildGetWithHashParams(params)
}
