package commons

// XPubList ...
func XPubList(iList []interface{}) []XPub {
	var xPubList []XPub
	for _, iXPub := range iList {
		xPubList = append(xPubList, XPub{XPub: iXPub.(map[string]interface{})["xPub"].(string)})
	}
	return xPubList
}

// BuildProof ...
func BuildProof(proofParameter map[string]interface{}) Proof {
	projectNameParam := proofParameter["projectName"].(map[string]interface{})
	projectName := ProjectName{Name: projectNameParam["name"].(string)}

	licenceSPDXParam := proofParameter["licenseSPDX"].(map[string]interface{})
	licenceSPDX := Licence{SPDX: licenceSPDXParam["SPDX"].(string)}

	licenceSPDXChangeParam := proofParameter["licenseSPDXChange"].(map[string]interface{})
	licenceSPDXChange := Licence{SPDX: licenceSPDXChangeParam["SPDX"].(string)}

	groupRoleNameParam := proofParameter["groupRoleName"].(map[string]interface{})
	groupRoleName := GroupRoleName{Name: groupRoleNameParam["name"].(string)}

	descriptionPublicListParam := proofParameter["descriptionPublicList"].([]interface{})
	descriptionPublicList := keyValList(descriptionPublicListParam)

	identityListParam := proofParameter["identityList"].([]interface{})
	identityList := keyValList(identityListParam)

	eventListParam := proofParameter["eventList"].([]interface{})
	eventList := keyValList(eventListParam)

	environmentListParam := proofParameter["environmentList"].([]interface{})
	environmentList := keyValList(environmentListParam)

	qualityListParam := proofParameter["qualityList"].([]interface{})
	qualityList := keyValList(qualityListParam)

	contributeListParam := proofParameter["contributeList"].([]interface{})
	contributeList := keyValList(contributeListParam)

	originListParam := proofParameter["originList"].([]interface{})
	originList := HashList(originListParam)

	parentListParam := proofParameter["parentList"].([]interface{})
	parentList := HashList(parentListParam)

	previousListParam := proofParameter["previousList"].([]interface{})
	previousList := HashList(previousListParam)

	leftListParam := proofParameter["leftList"].([]interface{})
	leftList := HashList(leftListParam)

	ndaListParam := proofParameter["ndaList"].([]interface{})
	ndaList := keyValList(ndaListParam)

	confidentialDataListParam := proofParameter["confidentialDataList"].([]interface{})
	confidentialDataList := keyValList(confidentialDataListParam)

	metaDataListParam := proofParameter["metaDataList"].([]interface{})
	metaDataList := keyValList(metaDataListParam)

	officerListParam := proofParameter["officerList"].([]interface{})
	officerList := keyValList(officerListParam)

	editListParam := proofParameter["editList"].([]interface{})
	editList := keyValList(editListParam)

	certificateListParam := proofParameter["certificateList"].([]interface{})
	certificateList := keyValList(certificateListParam)

	exportControlListParam := proofParameter["exportControlList"].([]interface{})
	exportControlList := keyValList(exportControlListParam)

	keyValueListParam := proofParameter["keyValueList"].([]interface{})
	keyValueList := keyValList(keyValueListParam)

	forListParam := proofParameter["forList"].([]interface{})
	forList := XPubList(forListParam)

	toListParam := proofParameter["toList"].([]interface{})
	toList := XPubList(toListParam)

	tagListParam := proofParameter["tagList"].([]interface{})
	var tagList []string
	for _, tagParam := range tagListParam {
		tagList = append(tagList, tagParam.(string))
	}

	return Proof{
		ProjectName:           projectName,
		LicenseSPDX:           licenceSPDX,
		LicenseSPDXChange:     licenceSPDXChange,
		GroupRoleName:         groupRoleName,
		DescriptionPublicList: descriptionPublicList,
		IdentityList:          identityList,
		EventList:             eventList,
		EnvironmentList:       environmentList,
		QualityList:           qualityList,
		ContributeList:        contributeList,
		OriginList:            originList,
		ParentList:            parentList,
		PreviousList:          previousList,
		LeftList:              leftList,
		NdaList:               ndaList,
		ConfidentialDataList:  confidentialDataList,
		MetaDataList:          metaDataList,
		OfficerList:           officerList,
		EditList:              editList,
		CertificateList:       certificateList,
		ExportControlList:     exportControlList,
		KeyValueList:          keyValueList,
		ForList:               forList,
		ToList:                toList,
		TagList:               tagList,
	}
}

func keyValList(iList []interface{}) []KeyVal {
	var keyValList []KeyVal
	for _, keyValParam := range iList {
		keyValParamJSON := keyValParam.(map[string]interface{})
		keyVal := KeyVal{
			Key: keyValParamJSON["key"].(string),
			Val: keyValParamJSON["val"].(string),
		}
		keyValList = append(keyValList, keyVal)
	}
	return keyValList
}

// HashList ...
func HashList(iList []interface{}) []Hash {
	var hashList []Hash
	for _, hashParam := range iList {
		hashList = append(hashList, Hash{Hash: hashParam.(map[string]interface{})["hash"].(string)})
	}
	return hashList
}

// BuildTemplate ...
func BuildTemplate(templateParam map[string]interface{}) Template {
	hashParameter := templateParam["hash"].(map[string]interface{})
	hash := Hash{Hash: hashParameter["hash"].(string)}

	proofParameter := templateParam["proof"].(map[string]interface{})
	proof := BuildProof(proofParameter)

	projectNameParam := templateParam["projectName"].(map[string]interface{})
	projectName := ProjectName{Name: projectNameParam["name"].(string)}

	licenceSPDXParam := templateParam["licenseSPDX"].(map[string]interface{})
	licenceSPDX := Licence{SPDX: licenceSPDXParam["SPDX"].(string)}

	groupRoleNameParam := templateParam["groupRoleName"].(map[string]interface{})
	groupRoleName := GroupRoleName{Name: groupRoleNameParam["name"].(string)}

	stateParam := templateParam["state"].(map[string]interface{})
	state := StateReason{Reason: stateParam["reason"].(string)}

	userRequirement := templateParam["userRequirement"].(bool)
	projectRequirement := templateParam["projectRequirement"].(bool)

	userUserParam := templateParam["userUser"].(map[string]interface{})
	userUser := TemplateUser{Quorum: userUserParam["quorum"].(string)}

	userBackupParam := templateParam["userBackup"].(map[string]interface{})
	userBackup := TemplateUser{Quorum: userBackupParam["quorum"].(string)}

	userLockParam := templateParam["userLock"].(map[string]interface{})
	userLock := TemplateUser{Quorum: userLockParam["quorum"].(string)}

	userWitnessParam := templateParam["userWitness"].(map[string]interface{})
	userWitness := TemplateUser{Quorum: userWitnessParam["quorum"].(string)}

	userUserXPubListParam := templateParam["userUserXPubList"].([]interface{})
	userUserXPubList := XPubList(userUserXPubListParam)

	userBackupXPubListParam := templateParam["userBackupXPubList"].([]interface{})
	userBackupXPubList := XPubList(userBackupXPubListParam)

	userLockXPubListParam := templateParam["userLockXPubList"].([]interface{})
	userLockXPubList := XPubList(userLockXPubListParam)

	userWitnessXPubListParam := templateParam["userWitnessXPubList"].([]interface{})
	userWitnessXPubList := XPubList(userWitnessXPubListParam)

	projectOldParam := templateParam["projectOld"].(map[string]interface{})
	projectOld := TemplateProject{Quorum: projectOldParam["quorum"].(string)}

	projectParentParam := templateParam["projectParent"].(map[string]interface{})
	projectParent := TemplateProject{Quorum: projectParentParam["quorum"].(string)}

	projectBoardParam := templateParam["projectBoard"].(map[string]interface{})
	projectBoard := TemplateProject{Quorum: projectBoardParam["quorum"].(string)}

	projectMemberParam := templateParam["projectMember"].(map[string]interface{})
	projectMember := TemplateProject{Quorum: projectMemberParam["quorum"].(string)}

	projectCosignerParam := templateParam["projectCosigner"].(map[string]interface{})
	projectCosigner := TemplateProject{Quorum: projectCosignerParam["quorum"].(string)}

	projectWitnessParam := templateParam["projectWitness"].(map[string]interface{})
	projectWitness := TemplateProject{Quorum: projectWitnessParam["quorum"].(string)}

	projectOldXPubListParam := templateParam["projectOldXPubList"].([]interface{})
	projectOldXPubList := XPubList(projectOldXPubListParam)

	projectParentXPubListParam := templateParam["projectParentXPubList"].([]interface{})
	projectParentXPubList := XPubList(projectParentXPubListParam)

	projectBoardXPubListParam := templateParam["projectBoardXPubList"].([]interface{})
	projectBoardXPubList := XPubList(projectBoardXPubListParam)

	projectMemberXPubListParam := templateParam["projectMemberXPubList"].([]interface{})
	projectMemberXPubList := XPubList(projectMemberXPubListParam)

	projectCosignerXPubListParam := templateParam["projectCosignerXPubList"].([]interface{})
	projectCosignerXPubList := XPubList(projectCosignerXPubListParam)

	projectWitnessXPubListParam := templateParam["projectWitnessXPubList"].([]interface{})
	projectWitnessXPubList := XPubList(projectWitnessXPubListParam)

	scriptTemplate := templateParam["scriptTemplate"].(string)

	return Template{
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

// BuildBoarding ...
func BuildBoarding(boardingParam map[string]interface{}) Boarding {
	boardingXPubListParam := boardingParam["xPubList"].([]interface{})
	boardingXPubList := XPubList(boardingXPubListParam)

	boardingIn := boardingParam["in"].(bool)

	boardingGroupRoleNameListParam := boardingParam["groupRoleNameList"].([]interface{})
	var boardingGroupRoleNameList []GroupRoleName
	for _, groupRoleNameParam := range boardingGroupRoleNameListParam {
		boardingGroupRoleNameList = append(boardingGroupRoleNameList, GroupRoleName{Name: groupRoleNameParam.(string)})
	}

	boardingGroupActionNameListParam := boardingParam["groupActionNameList"].([]interface{})
	var boardingGroupActionNameList []GroupActionName
	for _, groupActionNameParam := range boardingGroupActionNameListParam {
		boardingGroupActionNameList = append(boardingGroupActionNameList, GroupActionName{Name: groupActionNameParam.(string)})
	}

	return Boarding{
		XPubList:            boardingXPubList,
		In:                  boardingIn,
		GroupRoleNameList:   boardingGroupRoleNameList,
		GroupActionNameList: boardingGroupActionNameList,
	}
}

// BuildSig ...
func BuildSig(sigParam map[string]interface{}) Sig {
	sigDataParam := sigParam["sig"].(map[string]interface{})
	return Sig{
		Sig:   SigData{Sig: sigDataParam["sig"].(string)},
		XPub:  XPub{XPub: sigParam["xPub"].(string)},
		XPubS: XPub{XPub: sigParam["xPubS"].(string)},
	}
}

// BuildGetWithXPubParams ...
func BuildGetWithXPubParams(params []interface{}) (string, XPub, StateReason) {
	_type := params[0].(string)

	xPubParam := params[1].(map[string]interface{})
	xPub := XPub{XPub: xPubParam["xPub"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := StateReason{Reason: stateParam["reason"].(string)}

	return _type, xPub, state
}

// BuildGetWithHashParams ...
func BuildGetWithHashParams(params []interface{}) (string, Hash, StateReason) {
	_type := params[0].(string)

	hashParam := params[1].(map[string]interface{})
	hash := Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[2].(map[string]interface{})
	state := StateReason{Reason: stateParam["reason"].(string)}

	return _type, hash, state
}
