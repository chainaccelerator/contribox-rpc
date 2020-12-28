package persistance

// Contribution ...
type Contribution struct {
	Id              int
	Hash            string
	XPub            string
	Tx1Id           string
	Tx0IdAmount     int
	Tx0IdIssueAsset string
	Tx0IdSigA       string
	State           string
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
	Val        string
	KeyValType string
}

// Proof ...
type Proof struct {
	Id                int
	XPub              string
	ProjectName       string
	LicenceSPDX       string
	LicenceSPDXChange string
	GroupRoleName     string
}

// PubKey ...
type PubKey struct {
	Id            int
	PubKey        string
	Base58Encoded bool
	PubKeyType    string
}

// Sig ...
type Sig struct {
	Id    int
	Sig   string
	XPub  string
	XPubS string
}

// Tag ...
type Tag struct {
	Id  int
	Tag string
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
