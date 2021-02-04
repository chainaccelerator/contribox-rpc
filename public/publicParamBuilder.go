package public

// PublicPeerListGetBuildParam ...
func PublicPeerListGetBuildParam(params []interface{}) []string {
	tagListParam := params[0].([]interface{})
	var tagList []string
	for _, tagParam := range tagListParam {
		tagList = append(tagList, tagParam.(string))
	}

	return tagList
}
