package commons

// ValidateGroupRoleName ...
func ValidateGroupRoleName(groupRoleName string) bool {
	return groupRoleName == "old" ||
		groupRoleName == "parent" ||
		groupRoleName == "board" ||
		groupRoleName == "member" ||
		groupRoleName == "cosigner" ||
		groupRoleName == "witness"
}

// ValidateGroupRoleNameList ...
func ValidateGroupRoleNameList(groupRoleNameList []GroupRoleName) bool {
	for _, groupRoleName := range groupRoleNameList {
		if !ValidateGroupRoleName(groupRoleName.Name) {
			return false
		}
	}
	return true
}

// ValidateKeyOrHash ...
func ValidateKeyOrHash(keyOrHash string) bool {
	return len(keyOrHash) == 64
}

// ValidateXPub ...
func ValidateXPub(xPub string) bool {
	return len(xPub) <= 112
}

// ValidateXPubList ...
func ValidateXPubList(xPubList []XPub) bool {
	for _, xPub := range xPubList {
		if !ValidateXPub(xPub.XPub) {
			return false
		}
	}
	return true
}

// ValidateTemplate ...
func ValidateTemplate(template Template) bool {
	return ValidateKeyOrHash(template.Hash.Hash) &&
		ValidateProof(template.Proof) &&
		ValidateGroupRoleName(template.GroupRoleName.Name) &&
		validateTemplateUser(template.UserUser.Quorum) &&
		validateTemplateUser(template.UserBackup.Quorum) &&
		validateTemplateUser(template.UserLock.Quorum) &&
		validateTemplateUser(template.UserWitness.Quorum)
}

// ValidateProof ...
func ValidateProof(proof Proof) bool {
	return ValidateGroupRoleName(proof.GroupRoleName.Name) &&
		validateHashList(proof.ParentList) &&
		validateHashList(proof.PreviousList) &&
		validateHashList(proof.LeftList)
}

func validateHashList(hashList []Hash) bool {
	for _, hash := range hashList {
		if !ValidateKeyOrHash(hash.Hash) {
			return false
		}
	}
	return true
}

func validateTemplateUser(quorum string) bool {
	return quorum == "x" ||
		quorum == "any||timeout" ||
		quorum == "timeout||any" ||
		quorum == "any"
}

// ValidateUTXOList ...
func ValidateUTXOList(utxoList []UTXO) bool {
	for _, utxo := range utxoList {
		if !validateUTXO(utxo) {
			return false
		}
	}
	return true
}

func validateUTXO(utxo UTXO) bool {
	return ValidateKeyOrHash(utxo.Hash.Hash) &&
		ValidateTX(utxo.Tx0Id.Id) &&
		ValidateTX(utxo.UTXO.UTXO) &&
		validateScript(utxo.Script.Script)
}

// ValidateTX ...
func ValidateTX(tx string) bool {
	return len(tx) == 32
}

func validateScript(script string) bool {
	return len(script) <= 1650
}

// ValidateBoarding ...
func ValidateBoarding(boarding Boarding) bool {
	return ValidateXPubList(boarding.XPubList) && ValidateGroupRoleNameList(boarding.GroupRoleNameList)
}

// ValidateSigData ...
func ValidateSigData(sig string) bool {
	return len(sig) >= 65 && len(sig) <= 71
}

// ValidatePubKey ...
func ValidatePubKey(pubKey string) bool {
	return len(pubKey) == 65
}

// ValidateSig ...
func ValidateSig(sig Sig) bool {
	return ValidateSigData(sig.Sig.Sig) && ValidateXPub(sig.XPub.XPub) && ValidateXPub(sig.XPubS.XPub)
}
