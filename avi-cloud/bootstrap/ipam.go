/*
* [2013] - [2019] Avi Networks Incorporated
* All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package bootstrap

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	avimodels "github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/servicemesh/utils"
)

func IPAMRestOps(ipamfilename string, avi_rest_client_pool *utils.AviRestClientPool) {
	network := HydrateNetwork(ipamfilename)
	cloud_name := os.Getenv("CLOUD_NAME")
	if cloud_name != "" {
		cloud_ref := "/api/cloud/?name=" + cloud_name
		utils.AviLog.Info.Printf("Setting cloud reference to the network profile :%s", cloud_name)
		network.CloudRef = &cloud_ref
	}
	network_name := os.Getenv("NETWORK_NAME")
	// Override the network name if it's provided
	if network_name != "" {
		utils.AviLog.Info.Printf("Overriding the network name with :%s", network_name)
		network.Name = &network_name
	}
	var rest_ops []*utils.RestOp

	path := "/api/network/"
	rest_op := utils.RestOp{Path: path, Method: "POST", Obj: network,
		Tenant: "admin", Model: "Network", Version: utils.CtrlVersion}
	rest_ops = append(rest_ops, &rest_op)
	err := avi_rest_client_pool.AviRestOperate(avi_rest_client_pool.AviClient[0], rest_ops)
	if err != nil {
		utils.AviLog.Warning.Printf("Couldn't create the network IPAM name:%s due to the following error :%s ", *network.Name, err.Error())

	} else {
		utils.AviLog.Info.Printf("Successfully created the network IPAM :%s", *network.Name)
	}
}

func HydrateNetwork(ipamfilename string) avimodels.Network {
	file, err := os.Open(ipamfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	network := avimodels.Network{}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&network); err != nil {
		utils.AviLog.Warning.Printf("parsing config file %s", err.Error())
		os.Exit(1)
	}
	// Let's set some values based on environment variables.
	cidr := os.Getenv("CIDR")
	var ipaddr string
	var mask int32
	addressType := "V4"
	subnet := avimodels.Subnet{}
	if cidr != "" {
		// CIDR should be of the format ip/<mask>
		splitcidr := strings.Split(cidr, "/")
		if len(splitcidr) != 2 {
			// wrong cidr provided. Let's exit
			utils.AviLog.Info.Printf("Wrong CIDR provided. Format is: ipaddress/mask")
			os.Exit(1)
		} else {
			ipaddr = splitcidr[0]
			maskint, _ := strconv.Atoi(splitcidr[1])
			mask = int32(maskint)
			// Process ip address
			ipaddrobj := avimodels.IPAddr{Addr: &ipaddr, Type: &addressType}
			ipprefixobj := avimodels.IPAddrPrefix{IPAddr: &ipaddrobj, Mask: &mask}
			subnet.Prefix = &ipprefixobj
		}
	}
	rangeStart := os.Getenv("STATIC_RANGE_START")
	rangeEnd := os.Getenv("STATIC_RANGE_END")
	if rangeStart != "" && rangeEnd != "" {
		// We should process the static range.
		startIppAddr := avimodels.IPAddr{Addr: &rangeStart, Type: &addressType}
		endIpAddr := avimodels.IPAddr{Addr: &rangeEnd, Type: &addressType}
		staticRange := avimodels.IPAddrRange{Begin: &startIppAddr, End: &endIpAddr}
		subnet.StaticRanges = append(subnet.StaticRanges, &staticRange)
	} else {
		// Just log
		utils.AviLog.Info.Printf("Static range not provided")
	}
	network.ConfiguredSubnets = append(network.ConfiguredSubnets, &subnet)
	return network
}

func IPAMProviderProfileRestOps(ipamprofilefilename string, avi_rest_client_pool *utils.AviRestClientPool) {
	file, err := os.Open(ipamprofilefilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dnsProfile := avimodels.IPAMDNSProviderProfile{}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&dnsProfile); err != nil {
		utils.AviLog.Warning.Printf("parsing config file %s", err.Error())
		os.Exit(1)
	}

	var rest_ops []*utils.RestOp

	network_name := os.Getenv("NETWORK_NAME")
	// Override the network name if it's provided
	if network_name != "" {
		utils.AviLog.Info.Printf("Overriding the network name with :%s", network_name)
		networkRef := "/api/network/?name=" + network_name
		dnsInternalProfile := avimodels.IPAMDNSInternalProfile{}
		dnsInternalProfile.UsableNetworkRefs = append(dnsInternalProfile.UsableNetworkRefs, networkRef)
		dnsProfile.InternalProfile = &dnsInternalProfile
	}
	path := "/api/ipamdnsproviderprofile/"
	rest_op := utils.RestOp{Path: path, Method: "POST", Obj: dnsProfile,
		Tenant: "admin", Model: "IPAMDNSProviderProfile", Version: utils.CtrlVersion}
	rest_ops = append(rest_ops, &rest_op)
	err = avi_rest_client_pool.AviRestOperate(avi_rest_client_pool.AviClient[0], rest_ops)
	if err != nil {
		utils.AviLog.Warning.Printf("Couldn't create the DNS Provider IPAM due to the following error :%s ", err.Error())
	} else {
		utils.AviLog.Info.Printf("Successfully created the DNS network IPAM")
	}

}

func IPAMDNSProfileRestOps(ipamprofilefilename string, avi_rest_client_pool *utils.AviRestClientPool) {
	file, err := os.Open(ipamprofilefilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dnsProfile := avimodels.IPAMDNSProviderProfile{}
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&dnsProfile); err != nil {
		utils.AviLog.Warning.Printf("parsing config file %s", err.Error())
		os.Exit(1)
	}

	var rest_ops []*utils.RestOp
	dnsSubDomain := os.Getenv("DNS_SUBDOMAIN")
	utils.AviLog.Info.Printf("Sending DNS Profile info %s for creation", utils.Stringify(dnsProfile))
	if dnsSubDomain == "" {
		utils.AviLog.Info.Printf("DNS subdomain not provided, will use default value.")
	} else {
		dnsSD := avimodels.DNSServiceDomain{DomainName: &dnsSubDomain}
		dnsInternalProfile := avimodels.IPAMDNSInternalProfile{}
		dnsInternalProfile.DNSServiceDomain = append(dnsInternalProfile.DNSServiceDomain, &dnsSD)
		dnsProfile.InternalProfile = &dnsInternalProfile
	}
	utils.AviLog.Info.Printf("Sending DNS Profile info %s for creation", utils.Stringify(dnsProfile))
	path := "/api/ipamdnsproviderprofile/"
	rest_op := utils.RestOp{Path: path, Method: "POST", Obj: dnsProfile,
		Tenant: "admin", Model: "IPAMDNSProviderProfile", Version: utils.CtrlVersion}
	rest_ops = append(rest_ops, &rest_op)
	err = avi_rest_client_pool.AviRestOperate(avi_rest_client_pool.AviClient[0], rest_ops)
	if err != nil {
		utils.AviLog.Warning.Printf("Couldn't create the DNS Profile due to the following error :%s ", err.Error())
	} else {
		utils.AviLog.Info.Printf("Successfully created the DNS Profile")
	}

}
