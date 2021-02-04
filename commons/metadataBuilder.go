package commons

import "bc_node_api/api3/hashes"

// BuildMetadata ...
func BuildMetadata(metadataParam map[string]interface{}) Metadata {
	hashParam := metadataParam["hash"].(map[string]interface{})

	return Metadata{
		Hash: hashes.BuildHash(hashParam),
	}
}
