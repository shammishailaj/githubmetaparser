package schemas

type GithubMetaSSHKeyFingerprints struct {
	MD5_RSA    string `json:"MD5_RSA"`
	MD5_DSA    string `json:"MD5_DSA"`
	SHA256_RSA string `json:"SHA256_RSA"`
	SHA256_DSA string `json:"SHA256_DSA"`
}

type GithubMeta struct {
	VerifiablePasswordAuthentication bool                         `json:"verifiable_password_authentication"`
	SSHKeyFingerprints               GithubMetaSSHKeyFingerprints `json:"ssh_key_fingerprints"`
	HooksIPs                         []string                     `json:"hooks"`
	WebIPs                           []string                     `json:"web"`
	ApiIPs                           []string                     `json:"api"`
	GitIPs                           []string                     `json:"git"`
	PagesIPs                         []string                     `json:"pages"`
	ImporterIPs                      []string                     `json:"importer"`
}
