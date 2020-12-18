package authentication

import "bc_node_api/api3/commons"

// MetaData ...
type MetaData struct {
	Time       commons.TimeStamp
	PeerList   []string
	PubKey     commons.PubKey
	Sig        commons.Sig
	Hash       commons.Hash
	PowDiff    int
	PowPreffix string
	PowNonce   int
	PowHash    commons.Hash
	Id         string
	State      string
}
