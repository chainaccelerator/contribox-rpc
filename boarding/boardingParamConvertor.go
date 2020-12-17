package boarding

import "bc_node_api/api3/commons"

// BoardingTemplateGetParamConvert ...
func BoardingTemplateGetParamConvert(params ...interface{}) (commons.ProjectName, commons.Licence, commons.GroupRoleName, commons.Boarding, commons.Boarding, commons.Hash, commons.StateReason) {
	projectNameParam := params[0].(map[string]interface{})
	projectName := commons.ProjectName{Name: projectNameParam["name"].(string)}

	licenceSPDXParam := params[1].(map[string]interface{})
	licenceSPDX := commons.Licence{SPDX: licenceSPDXParam["licenceSPDX"].(string)}

	groupRoleNameParam := params[2].(map[string]interface{})
	groupRoleName := commons.GroupRoleName{Name: groupRoleNameParam["groupRoleName"].(string)}

	onBoardingParam := params[3].(map[string]interface{})
	onBoarding := boarding(onBoardingParam)

	outBoardingParam := params[4].(map[string]interface{})
	outBoarding := boarding(outBoardingParam)

	hashParam := params[5].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[6].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return projectName, licenceSPDX, groupRoleName, onBoarding, outBoarding, hash, state
}

func boarding(boardingParam map[string]interface{}) commons.Boarding {
	boardingXPubListParam := boardingParam["xPubList"].([]interface{})
	boardingXPubList := commons.XPubList(boardingXPubListParam)

	boardingIn := boardingParam["in"].(bool)

	boardingGroupRoleNameListParam := boardingParam["groupRoleNameList"].([]interface{})
	var boardingGroupRoleNameList []commons.GroupRoleName
	for _, groupRoleNameParam := range boardingGroupRoleNameListParam {
		boardingGroupRoleNameList = append(boardingGroupRoleNameList, commons.GroupRoleName{Name: groupRoleNameParam.(string)})
	}

	boardingGroupActionNameListParam := boardingParam["groupActionNameList"].([]interface{})
	var boardingGroupActionNameList []commons.GroupActionName
	for _, groupActionNameParam := range boardingGroupActionNameListParam {
		boardingGroupActionNameList = append(boardingGroupActionNameList, commons.GroupActionName{Name: groupActionNameParam.(string)})
	}

	return commons.Boarding{
		XPubList:            boardingXPubList,
		In:                  boardingIn,
		GroupRoleNameList:   boardingGroupRoleNameList,
		GroupActionNameList: boardingGroupActionNameList,
	}
}

// BoardingParamConvert ...
func BoardingParamConvert(params ...interface{}) (string, commons.Template, commons.StateReason) {
	_type := params[0].(string)

	templateParam := params[1].(map[string]interface{})
	template := template(templateParam)

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, template, state
}

func template(templateParam map[string]interface{}) commons.Template {
	hashParameter := templateParam["hash"].(map[string]interface{})
	hash := commons.Hash{Hash: hashParameter["hash"].(string)}

	proofParameter := templateParam["proof"].(map[string]interface{})
	proof := proof(proofParameter)

	projectNameParam := templateParam["projectName"].(map[string]interface{})
	projectName := commons.ProjectName{Name: projectNameParam["name"].(string)}

	licenceSPDXParam := templateParam["licenseSPDX"].(map[string]interface{})
	licenceSPDX := commons.Licence{SPDX: licenceSPDXParam["SPDX"].(string)}

	groupRoleNameParam := templateParam["groupRoleName"].(map[string]interface{})
	groupRoleName := commons.GroupRoleName{Name: groupRoleNameParam["name"].(string)}

	stateParam := templateParam["state"].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	userRequirement := templateParam["userRequirement"].(bool)
	projectRequirement := templateParam["projectRequirement"].(bool)

	userUserParam := templateParam["userUser"].(map[string]interface{})
	userUser := commons.TemplateUser{Quorum: userUserParam["quorum"].(string)}

	userBackupParam := templateParam["userBackup"].(map[string]interface{})
	userBackup := commons.TemplateUser{Quorum: userBackupParam["quorum"].(string)}

	userLockParam := templateParam["userLock"].(map[string]interface{})
	userLock := commons.TemplateUser{Quorum: userLockParam["quorum"].(string)}

	userWitnessParam := templateParam["userWitness"].(map[string]interface{})
	userWitness := commons.TemplateUser{Quorum: userWitnessParam["quorum"].(string)}

	userUserXPubListParam := templateParam["userUserXPubList"].([]interface{})
	userUserXPubList := commons.XPubList(userUserXPubListParam)

	userBackupXPubListParam := templateParam["userBackupXPubList"].([]interface{})
	userBackupXPubList := commons.XPubList(userBackupXPubListParam)

	userLockXPubListParam := templateParam["userLockXPubList"].([]interface{})
	userLockXPubList := commons.XPubList(userLockXPubListParam)

	userWitnessXPubListParam := templateParam["userWitnessXPubList"].([]interface{})
	userWitnessXPubList := commons.XPubList(userWitnessXPubListParam)

	projectOldParam := templateParam["projectOld"].(map[string]interface{})
	projectOld := commons.TemplateProject{Quorum: projectOldParam["quorum"].(string)}

	projectParentParam := templateParam["projectParent"].(map[string]interface{})
	projectParent := commons.TemplateProject{Quorum: projectParentParam["quorum"].(string)}

	projectBoardParam := templateParam["projectBoard"].(map[string]interface{})
	projectBoard := commons.TemplateProject{Quorum: projectBoardParam["quorum"].(string)}

	projectMemberParam := templateParam["projectMember"].(map[string]interface{})
	projectMember := commons.TemplateProject{Quorum: projectMemberParam["quorum"].(string)}

	projectCosignerParam := templateParam["projectCosigner"].(map[string]interface{})
	projectCosigner := commons.TemplateProject{Quorum: projectCosignerParam["quorum"].(string)}

	projectWitnessParam := templateParam["projectWitness"].(map[string]interface{})
	projectWitness := commons.TemplateProject{Quorum: projectWitnessParam["quorum"].(string)}

	projectOldXPubListParam := templateParam["projectOldXPubList"].([]interface{})
	projectOldXPubList := commons.XPubList(projectOldXPubListParam)

	projectParentXPubListParam := templateParam["projectParentXPubList"].([]interface{})
	projectParentXPubList := commons.XPubList(projectParentXPubListParam)

	projectBoardXPubListParam := templateParam["projectBoardXPubList"].([]interface{})
	projectBoardXPubList := commons.XPubList(projectBoardXPubListParam)

	projectMemberXPubListParam := templateParam["projectMemberXPubList"].([]interface{})
	projectMemberXPubList := commons.XPubList(projectMemberXPubListParam)

	projectCosignerXPubListParam := templateParam["projectCosignerXPubList"].([]interface{})
	projectCosignerXPubList := commons.XPubList(projectCosignerXPubListParam)

	projectWitnessXPubListParam := templateParam["projectWitnessXPubList"].([]interface{})
	projectWitnessXPubList := commons.XPubList(projectWitnessXPubListParam)

	scriptTemplate := templateParam["scriptTemplate"].(string)

	return commons.Template{
		Hash:                    hash,
		Proof:                   proof,
		ProjectName:             projectName,
		LicenceSPDX:             licenceSPDX,
		GroupRoleName:           groupRoleName,
		State:                   state,
		UserRequirement:         userRequirement,
		ProjectRequirement:      projectRequirement,
		UserUser:                userUser,
		UserBackup:              userBackup,
		UserLock:                userLock,
		UserWitness:             userWitness,
		UserUserXPubList:        userUserXPubList,
		UserBackupXPubList:      userBackupXPubList,
		UserLockXPubList:        userLockXPubList,
		UserWitnessXPubList:     userWitnessXPubList,
		ProjectOld:              projectOld,
		ProjectParent:           projectParent,
		ProjectBoard:            projectBoard,
		ProjectMember:           projectMember,
		ProjectCosigner:         projectCosigner,
		ProjectWitness:          projectWitness,
		ProjectOldXPubList:      projectOldXPubList,
		ProjectParentXPubList:   projectParentXPubList,
		ProjectBoardXPubList:    projectBoardXPubList,
		ProjectMemberXPubList:   projectMemberXPubList,
		ProjectCosignerXPubList: projectCosignerXPubList,
		ProjectWitnessXPubList:  projectWitnessXPubList,
		ScriptTemplate:          scriptTemplate,
	}
}

func proof(proofParameter map[string]interface{}) commons.Proof {
	return commons.Proof{}
}

// BoardingGetParamConvert ...
func BoardingGetParamConvert(params ...interface{}) (string, commons.XPub, commons.StateReason) {
	_type := params[0].(string)

	xPubSParam := params[1].(map[string]interface{})
	xPubS := commons.XPub{XPub: xPubSParam["xPub"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, xPubS, state
}

// BoardingBroadcastParamConvert ...
func BoardingBroadcastParamConvert(params ...interface{}) (string, []commons.UTXO, commons.Hash, commons.StateReason) {
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
func BoardingBroadcastGetParamConvert(params ...interface{}) (string, commons.Hash, commons.StateReason) {
	_type := params[0].(string)

	hashParam := params[1].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, hash, state
}
