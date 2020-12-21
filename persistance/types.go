package persistance

// Contribution ...
type Contribution struct {
	Id              int
	Hash            string
	Tx1Id           string
	Tx0IdAmount     int
	Tx0IdIssueAsset string
	Tx0IdSigA       string
}

// GroupName ...
type GroupName struct {
	Id   int
	Name string
	Type string
}

// Hash ...
type Hash struct {
	Id       int
	Hash     string
	HashType string
}

// KeyVal ...
type KeyVal struct {
	Id         int
	Key        string
	Val        map[string]interface{}
	KeyValType string
}

// Proof ...
type Proof struct {
	Id                int
	ProjectName       string
	LicenseSPDX       string
	LicenseSPDXChange string
	GroupRoleName     string
}

// PubKey ...
type PubKey struct {
	Id            int
	PubKey        string
	Base58Encoded bool
}

// Template ...
type Template struct {
	Id                 int
	Hash               string
	ProjectName        string
	LicenceSPDX        string
	GroupRoleName      string
	State              string
	UserRequirement    bool
	ProjectRequirement bool
	UserUser           string
	UserBackup         string
	UserLock           string
	UserWitness        string
	ProjectOld         string
	ProjectParent      string
	ProjectBoard       string
	ProjectMember      string
	ProjectCosigner    string
	ProjectWitness     string
}

// XPub ...
type XPub struct {
	Id       int
	XPub     string
	XPubType string
}
