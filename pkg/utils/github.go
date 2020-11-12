package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/shammishailaj/metaparser/pkg/schemas"
	"time"
)

func(u *Utils) PrintGithubIPs(appendDeny bool) {
	fmt.Printf("# Getting the Github Meta Data at: %s\n", time.Now().String())
	url := "https://api.github.com/meta" // https://docs.github.com/en/free-pro-team@latest/rest/reference/meta
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

	var responseData schemas.GithubMeta
	responseDataErr := json.Unmarshal([]byte(resp.String()), &responseData)
	if responseDataErr != nil {
		fmt.Printf("# Error Unmarshalling response from %s. %s\n", url, responseDataErr.Error())
		return
	}

	fmt.Printf("# Now outputting IP data for nginx\n# Github IPs/CIDRs data for nginx\n")
	fmt.Printf("# Hooks IPs\n# =================\n")
	for _, v := range responseData.HooksIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Web IPs\n# =================\n")
	for _, v := range responseData.WebIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# API IPs\n# =================\n")
	for _, v := range responseData.ApiIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Git IPs\n# =================\n")
	for _, v := range responseData.GitIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Pages IPs\n# =================\n")
	for _, v := range responseData.PagesIPs {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")

	fmt.Printf("# Importer IPs\n# =================\n")
	for _, v := range responseData.ImporterIPs {
		fmt.Printf("allow %s;\n", v)
	}

	if appendDeny {
		fmt.Printf("\n\ndeny all;\n\n")
	}
}
