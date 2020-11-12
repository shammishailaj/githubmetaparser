package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

func(u *Utils) PrintCloudflareIPV4IPs(appendDeny bool) {
	fmt.Printf("# Getting the Cloudflare IPv4 Lists at: %s\n", time.Now().String())
	url := "https://www.cloudflare.com/ips-v4" // https://docs.github.com/en/free-pro-team@latest/rest/reference/meta
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

	ips := strings.Split(resp.String(), "\n")


	fmt.Printf("# Now outputting IP data for nginx\n# Cloudflare IPs/CIDRs data for nginx\n")
	fmt.Printf("# IPv4 IPs\n# =================\n")
	for _, v := range ips {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")
	if appendDeny {
		fmt.Printf("\n\ndeny all;\n\n")
	}
}

func(u *Utils) PrintCloudflareIPV6IPs(appendDeny bool) {
	fmt.Printf("# Getting the Cloudflare IPv6 Lists at: %s\n", time.Now().String())
	url := "https://www.cloudflare.com/ips-v6" // https://docs.github.com/en/free-pro-team@latest/rest/reference/meta
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

	ips := strings.Split(resp.String(), "\n")


	fmt.Printf("# Now outputting IP data for nginx\n# Cloudflare IPs/CIDRs data for nginx\n")
	fmt.Printf("# IPv4 IPs\n# =================\n")
	for _, v := range ips {
		fmt.Printf("allow %s;\n", v)
	}
	fmt.Printf("\n\n")
	if appendDeny {
		fmt.Printf("\n\ndeny all;\n\n")
	}
}