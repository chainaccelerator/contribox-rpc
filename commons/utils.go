package commons

// XPubList ...
func XPubList(iList []interface{}) []XPub {
	var xPubList []XPub
	for _, iXPub := range iList {
		xPubList = append(xPubList, XPub{XPub: iXPub.(map[string]interface{})["xPub"].(string)})
	}
	return xPubList
}
