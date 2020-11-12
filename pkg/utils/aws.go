package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shammishailaj/metaparser/pkg/schemas"
	"time"
)

// See: https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html
func(u *Utils) PrintAWSIPs(ipVersion int, region, service, networkBorderGroup string) {
	fmt.Printf("# Getting the AWS Meta Data at: %s\n", time.Now().String())
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json" // https://docs.github.com/en/free-pro-team@latest/rest/reference/meta
	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().EnableTrace().Get(url)

	if err != nil {
		fmt.Printf("# Error making request to %s. %s\n", url, err.Error())
		return
	}

	if resp.StatusCode() != 200 {
		fmt.Printf("# HTTP Status is %d (not 200). Won't continue\n", resp.StatusCode())
		return
	}

	var responseData schemas.AWSIPRanges
	responseDataErr := json.Unmarshal([]byte(resp.String()), &responseData)
	if responseDataErr != nil {
		fmt.Printf("# Error Unmarshalling response from %s. %s\n", url, responseDataErr.Error())
		return
	}

	if ipVersion != 4 && ipVersion != 6 {
		ipVersion = 4
	}

	fmt.Printf("# Now outputting IP data for nginx\n# Github IPs/CIDRs data for nginx\n")
	fmt.Printf("# Sync Token %s\n", responseData.SyncToken)
	fmt.Printf("# Create Date %s\n", responseData.CreateDate)
	if ipVersion == 4 {
		fmt.Printf("# IPV4 IPs\n# =================\n")
		for _, v := range responseData.Prefixes {
			if (region != "" && v.Region == region) || region == "all" {
				if (service != "" && v.Service == service) || service == "all" {
					if (networkBorderGroup != "" && v.NetworkBorderGroup == networkBorderGroup) || networkBorderGroup == "all" {
						fmt.Printf("allow %s;\n", v.IPPrefix)
					}
				}
			}
		}
		fmt.Printf("\n\n")
	}

	if ipVersion == 6 {
		fmt.Printf("# IPV6 IPs\n# =================\n")
		for _, v := range responseData.PrefixesV6 {
			if (region != "" && v.Region == region) || region == "all" {
				if (service != "" && v.Service == service) || service == "all" {
					if (networkBorderGroup != "" && v.NetworkBorderGroup == networkBorderGroup) || networkBorderGroup == "all" {
						fmt.Printf("allow %s;\n", v.IPPrefix)
					}
				}
			}
		}
		fmt.Printf("\n\n")
	}
}

