package schemas

type AWSIPRangesPrefixes struct {
	IPPrefix           string `json:"ip_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

type AWSIPRangesPrefixesV6 struct {
	IPPrefix           string `json:"ipv6_prefix"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	NetworkBorderGroup string `json:"network_border_group"`
}

type AWSIPRanges struct {
	SyncToken string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes []AWSIPRangesPrefixes `json:"prefixes"`
	PrefixesV6 []AWSIPRangesPrefixesV6 `json:"ipv6_prefixes"`
}
