package contribution

import (
	"bc_node_api/api3/commons"
)

// ContributionParamConvert ...
func ContributionParamConvert(params []interface{}) (string, commons.Contribution, commons.StateReason) {
	_type := params[0].(string)

	contributionParam := params[1].(map[string]interface{})
	contribution := buildContribution(contributionParam)

	stateParam := params[2].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, contribution, state
}

func buildContribution(contributionParam map[string]interface{}) commons.Contribution {
	hashParam := contributionParam["hash"].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	proofParam := contributionParam["proof"].(map[string]interface{})
	proof := commons.BuildProof(proofParam)

	templateParam := contributionParam["template"].(map[string]interface{})
	template := commons.BuildTemplate(templateParam)

	blindingKeyListParam := contributionParam["blindingKeyList"].([]interface{})
	var blindingKeyList []commons.BlindingKeyEncrypted
	for _, blindingKeyParam := range blindingKeyListParam {
		blindingKeyParamJSON := blindingKeyParam.(map[string]interface{})
		xPubParam := blindingKeyParamJSON["xPub"].(map[string]interface{})
		xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}
		blindingKeyEncrypted := commons.BlindingKeyEncrypted{
			XPub:   xPub,
			String: blindingKeyParamJSON["string"].(string),
		}
		blindingKeyList = append(blindingKeyList, blindingKeyEncrypted)
	}

	rangeListParam := contributionParam["rangeList"].([]interface{})
	var rangeList []commons.RangeEncrypted
	for _, rangeParam := range rangeListParam {
		rangeParamJSON := rangeParam.(map[string]interface{})
		xPubParam := rangeParamJSON["xPub"].(map[string]interface{})
		xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}
		rangeEncrypted := commons.RangeEncrypted{
			XPub:   xPub,
			String: rangeParamJSON["string"].(string),
		}
		rangeList = append(rangeList, rangeEncrypted)
	}

	onBoardingParam := contributionParam["onBoarding"].(map[string]interface{})
	onBoarding := commons.BuildBoarding(onBoardingParam)

	outBoardingParam := contributionParam["outBoarding"].(map[string]interface{})
	outBoarding := commons.BuildBoarding(outBoardingParam)

	tx1IdParam := contributionParam["Tx1Id"].(map[string]interface{})
	tx1Id := commons.TxId{Id: tx1IdParam["id"].(string)}

	tx0IdAmount := contributionParam["Tx0IdAmount"].(int)

	tx0IdIssueAssetParam := contributionParam["Tx0IdIssueAsset"].(map[string]interface{})
	issueAssetHashParam := tx0IdIssueAssetParam["hash"].(map[string]interface{})
	issueAssetHash := commons.Hash{Hash: issueAssetHashParam["hash"].(string)}
	tx0IdIssueAsset := commons.IssueAsset{IssueAsset: issueAssetHash}

	tx0IdSigAParam := contributionParam["Tx0IdSigA"].(map[string]interface{})
	tx0IdSigA := commons.SigData{Sig: tx0IdSigAParam["sig"].(string)}

	vout0PubKAParam := contributionParam["vout0PubKA"].(map[string]interface{})
	vout0PubKA := commons.PubKey{
		PubKey:        vout0PubKAParam["pubKey"].(string),
		Base58Encoded: vout0PubKAParam["base58Encoded"].(bool),
	}

	vout1PubKSParam := contributionParam["vout1PubKS"].(map[string]interface{})
	vout1PubKS := commons.PubKey{
		PubKey:        vout1PubKSParam["pubKey"].(string),
		Base58Encoded: vout1PubKSParam["base58Encoded"].(bool),
	}

	return commons.Contribution{
		Hash:            hash,
		Proof:           proof,
		Template:        template,
		BlindKeyList:    blindingKeyList,
		RangeList:       rangeList,
		OnBoarding:      onBoarding,
		OutBoarding:     outBoarding,
		Tx1Id:           tx1Id,
		Tx0IdAmount:     tx0IdAmount,
		Tx0IdIssueAsset: tx0IdIssueAsset,
		Tx0IdSigA:       tx0IdSigA,
		Vout0PubKA:      vout0PubKA,
		Vout1PubKS:      vout1PubKS,
	}
}

// ContributionGetParamConvert ...
func ContributionGetParamConvert(params []interface{}) (string, commons.XPub, commons.StateReason) {
	return commons.BuildGetWithXPubParams(params)
}

// ContributionConfirmParamConvert ...
func ContributionConfirmParamConvert(params []interface{}) (string, commons.Sig, commons.Hash, commons.XPub, string, commons.StateReason) {
	_type := params[0].(string)

	sigParam := params[1].(map[string]interface{})
	sig := commons.BuildSig(sigParam)

	hashParam := params[2].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	xPubParam := params[3].(map[string]interface{})
	xPub := commons.XPub{XPub: xPubParam["xPub"].(string)}

	resourceEncrypted := params[4].(string)

	stateParam := params[5].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, sig, hash, xPub, resourceEncrypted, state
}

// ContributionConfirmGetParamConvert ...
func ContributionConfirmGetParamConvert(params []interface{}) (string, commons.Hash, commons.StateReason) {
	return commons.BuildGetWithHashParams(params)
}

// ContributionBroadcastParamConvert ...
func ContributionBroadcastParamConvert(params []interface{}) (string, []commons.Contribution, commons.Hash, commons.StateReason) {
	_type := params[0].(string)

	resourceListParam := params[1].([]interface{})
	var resourceList []commons.Contribution
	for _, resourceParam := range resourceListParam {
		resourceList = append(resourceList, buildContribution(resourceParam.(map[string]interface{})))
	}

	hashParam := params[2].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	stateParam := params[3].(map[string]interface{})
	state := commons.StateReason{Reason: stateParam["reason"].(string)}

	return _type, resourceList, hash, state
}

// ContributionBroadcastGetParamConvert ...
func ContributionBroadcastGetParamConvert(params []interface{}) (string, commons.Hash, commons.StateReason) {
	return commons.BuildGetWithHashParams(params)
}
