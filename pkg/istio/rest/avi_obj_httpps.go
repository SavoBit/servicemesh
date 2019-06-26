/*
 * [2013] - [2018] Avi Networks Incorporated
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

package rest

import (
	"fmt"

	avimodels "github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/servicemesh/pkg/istio/nodes"
	"github.com/avinetworks/servicemesh/pkg/utils"

	"github.com/davecgh/go-spew/spew"
)

func AviHttpPSBuild(hps_meta *nodes.AviHttpPolicySetNode) *utils.RestOp {
	name := hps_meta.Name
	cksum := hps_meta.CloudConfigCksum
	cksumString := fmt.Sprint(cksum)
	tenant := fmt.Sprintf("/api/tenant/?name=%s", hps_meta.Tenant)
	cr := utils.OSHIFT_K8S_CLOUD_CONNECTOR

	http_req_pol := avimodels.HTTPRequestPolicy{}
	hps := avimodels.HTTPPolicySet{Name: &name, CloudConfigCksum: &cksumString,
		CreatedBy: &cr, TenantRef: &tenant, HTTPRequestPolicy: &http_req_pol}

	var idx int32
	idx = 0
	for _, hppmap := range hps_meta.HppMap {
		enable := true
		name := fmt.Sprintf("%s-%d", hps_meta.Name, idx)
		match_target := avimodels.MatchTarget{}
		if len(hppmap.Host) > 0 {
			match_crit := "HDR_EQUALS"
			host_hdr_match := avimodels.HostHdrMatch{MatchCriteria: &match_crit,
				Value: hppmap.Host}
			match_target.HostHdr = &host_hdr_match
		}
		if len(hppmap.Path) > 0 {
			match_crit := hppmap.MatchCriteria
			path_match := avimodels.PathMatch{MatchCriteria: &match_crit,
				MatchStr: hppmap.Path}
			match_target.Path = &path_match
		}
		if hppmap.Port != 0 {
			match_crit := "IS_IN"
			vsport_match := avimodels.PortMatch{MatchCriteria: &match_crit,
				Ports: []int64{int64(hppmap.Port)}}
			match_target.VsPort = &vsport_match
		}
		sw_action := avimodels.HttpswitchingAction{}
		if hppmap.Pool != "" {
			action := "HTTP_SWITCHING_SELECT_POOL"
			sw_action.Action = &action
			pool_ref := fmt.Sprintf("/api/pool/?name=%s", hppmap.Pool)
			sw_action.PoolRef = &pool_ref
		} else if hppmap.PoolGroup != "" {
			action := "HTTP_SWITCHING_SELECT_POOLGROUP"
			sw_action.Action = &action
			pg_ref := fmt.Sprintf("/api/poolgroup/?name=%s", hppmap.PoolGroup)
			sw_action.PoolGroupRef = &pg_ref
		}
		rule := avimodels.HTTPRequestRule{Enable: &enable, Index: &idx,
			Name: &name, Match: &match_target, SwitchingAction: &sw_action}
		http_req_pol.Rules = append(http_req_pol.Rules, &rule)
		idx = idx + 1
	}

	macro := utils.AviRestObjMacro{ModelName: "HTTPPolicySet", Data: hps}

	// TODO Version should be latest from configmap
	rest_op := utils.RestOp{Path: "/api/macro", Method: utils.RestPost, Obj: macro,
		Tenant: hps_meta.Tenant, Model: "HTTPPolicySet", Version: utils.CtrlVersion}

	utils.AviLog.Info.Print(spew.Sprintf("HTTPPolicySet Restop %v AviHttpPolicySetMeta %v\n",
		rest_op, *hps_meta))
	return &rest_op
}