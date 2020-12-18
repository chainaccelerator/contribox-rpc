package persistance

// Proof ...
type Proof struct {
	Id                int
	ProjectName       string
	LicenseSPDX       string
	LicenseSPDXChange string
	GroupRoleName     string
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

// GroupName ...
type GroupName struct {
	Id int
}
