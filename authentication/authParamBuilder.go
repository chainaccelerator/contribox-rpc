package authentication

import "bc_node_api/api3/commons"

func authPeerListGetBuildParam(params []interface{}) (MetaData, string, []interface{}) {
	metaDataParam := params[0].(map[string]interface{})

	timeStampParam := metaDataParam["time"].(map[string]interface{})
	timeStamp := commons.TimeStamp{Time: timeStampParam["time"].(int), Delay: timeStampParam["delay"].(int)}

	peerListParam := metaDataParam["peerList"].([]interface{})
	var peerList []string
	for _, peerParam := range peerListParam {
		peerList = append(peerList, peerParam.(string))
	}

	pubKeyParam := metaDataParam["pubKey"].(map[string]interface{})
	pubKey := commons.PubKey{PubKey: pubKeyParam["pubKey"].(string), Base58Encoded: pubKeyParam["base58Encoded"].(bool)}

	sigParam := metaDataParam["sig"].(map[string]interface{})
	sig := commons.BuildSig(sigParam)

	hashParam := metaDataParam["hash"].(map[string]interface{})
	hash := commons.Hash{Hash: hashParam["hash"].(string)}

	powDiff := metaDataParam["PowDiff"].(int)
	powPrefix := metaDataParam["PowPrefix"].(string)
	powNonce := metaDataParam["PowNonce"].(int)
	powHashParam := metaDataParam["PowHash"].(map[string]interface{})
	powHash := commons.Hash{Hash: powHashParam["hash"].(string)}

	id := metaDataParam["id"].(string)
	state := metaDataParam["state"].(string)

	metaData := MetaData{
		Time:       timeStamp,
		PeerList:   peerList,
		PubKey:     pubKey,
		Sig:        sig,
		Hash:       hash,
		PowDiff:    powDiff,
		PowPreffix: powPrefix,
		PowNonce:   powNonce,
		PowHash:    powHash,
		Id:         id,
		State:      state,
	}

	method := params[1].(string)
	_params := params[2].([]interface{})

	return metaData, method, _params
}
