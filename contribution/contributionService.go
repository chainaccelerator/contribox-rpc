package contribution

import (
	"bc_node_api/api3/commons"
	"bc_node_api/api3/persistance"
)

// ContributionDb ...
func ContributionDb(_type string, contribution commons.Contribution, state string, dbConf commons.DbConf) commons.StateReason {
	return commons.StateReason{}
}

// ContributionGetDb ...
func ContributionGetDb(_type string, xPubS string, state string) commons.Contribution {
	contribution := GetContributionByID(_type, xPubS, state)
	return commons.Contribution{Hash: commons.Hash{Hash: contribution.Hash}}
}

func buildContributionDTO(
	contribution persistance.Contribution,
	proof persistance.Proof, template persistance.Template,
	vout0pubKeyA persistance.PubKey,
	vout1pubKeyS persistance.PubKey,
) commons.Contribution {
	return commons.Contribution{}
}
