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
	"errors"
	"fmt"

	avimodels "github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/servicemesh/amc/pkg/istio/nodes"
	"github.com/avinetworks/servicemesh/utils"

	"github.com/davecgh/go-spew/spew"
)

func AviHttpPSBuild(hps_meta *nodes.AviHttpPolicySetNode, cache_obj *utils.AviHTTPCache) *utils.RestOp {
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
		var j int32
		j = idx
		rule := avimodels.HTTPRequestRule{Enable: &enable, Index: &j,
			Name: &name, Match: &match_target, SwitchingAction: &sw_action}
		http_req_pol.Rules = append(http_req_pol.Rules, &rule)
		idx = idx + 1
	}

	for _, hppmap := range hps_meta.RedirectPorts {
		enable := true
		name := fmt.Sprintf("%s-%d", hps_meta.Name, idx)
		match_target := avimodels.MatchTarget{}
		if len(hppmap.Hosts) > 0 {
			match_crit := "HDR_EQUALS"
			host_hdr_match := avimodels.HostHdrMatch{MatchCriteria: &match_crit,
				Value: hppmap.Hosts}
			match_target.HostHdr = &host_hdr_match
			port_match_crit := "IS_IN"
			match_target.VsPort = &avimodels.PortMatch{MatchCriteria: &port_match_crit, Ports: []int64{int64(hppmap.VsPort)}}
		}
		redirect_action := avimodels.HTTPRedirectAction{}
		protocol := "HTTP"
		redirect_action.StatusCode = &hppmap.StatusCode
		redirect_action.Protocol = &protocol
		redirect_action.Port = &hppmap.RedirectPort
		var j int32
		j = idx
		rule := avimodels.HTTPRequestRule{Enable: &enable, Index: &j,
			Name: &name, Match: &match_target, RedirectAction: &redirect_action}
		http_req_pol.Rules = append(http_req_pol.Rules, &rule)
		idx = idx + 1
	}

	macro := utils.AviRestObjMacro{ModelName: "HTTPPolicySet", Data: hps}
	var path string
	var rest_op utils.RestOp
	if cache_obj != nil {
		path = "/api/httppolicyset/" + cache_obj.Uuid
		rest_op = utils.RestOp{Path: path, Method: utils.RestPut, Obj: hps,
			Tenant: hps_meta.Tenant, Model: "HTTPPolicySet", Version: utils.CtrlVersion}

	} else {
		path = "/api/macro"
		rest_op = utils.RestOp{Path: path, Method: utils.RestPost, Obj: macro,
			Tenant: hps_meta.Tenant, Model: "HTTPPolicySet", Version: utils.CtrlVersion}
	}

	utils.AviLog.Info.Print(spew.Sprintf("HTTPPolicySet Restop %v AviHttpPolicySetMeta %v\n",
		rest_op, *hps_meta))
	return &rest_op
}

func AviHttpPolicyDel(uuid string, tenant string) *utils.RestOp {
	path := "/api/httppolicyset/" + uuid
	rest_op := utils.RestOp{Path: path, Method: "DELETE",
		Tenant: tenant, Model: "HTTPPolicySet", Version: utils.CtrlVersion}
	utils.AviLog.Info.Print(spew.Sprintf("HTTP Policy Set DELETE Restop %v \n",
		utils.Stringify(rest_op)))
	return &rest_op
}

func AviHTTPPolicyCacheAdd(cache *utils.AviObjCache, rest_op *utils.RestOp, vsKey utils.NamespaceName) error {
	if (rest_op.Err != nil) || (rest_op.Response == nil) {
		utils.AviLog.Warning.Printf("rest_op has err or no reponse for HTTP Policy Set Objects")
		return errors.New("Errored rest_op")
	}

	resp_elems, ok := RestRespArrToObjByType(rest_op, "httppolicyset")
	if ok != nil || resp_elems == nil {
		utils.AviLog.Warning.Printf("Unable to find HTTP Policy Set obj in resp %v", rest_op.Response)
		return errors.New("HTTP Policy Set object not found")
	}

	for _, resp := range resp_elems {
		name, ok := resp["name"].(string)
		if !ok {
			utils.AviLog.Warning.Printf("Name not present in response %v", resp)
			continue
		}

		uuid, ok := resp["uuid"].(string)
		if !ok {
			utils.AviLog.Warning.Printf("Uuid not present in response %v", resp)
			continue
		}

		cksum := resp["cloud_config_cksum"].(string)

		http_cache_obj := utils.AviHTTPCache{Name: name, Tenant: rest_op.Tenant,
			Uuid:             uuid,
			CloudConfigCksum: cksum}

		k := utils.NamespaceName{Namespace: rest_op.Tenant, Name: name}
		cache.HTTPCache.AviCacheAdd(k, &http_cache_obj)
		vs_cache, ok := cache.VsCache.AviCacheGet(vsKey)
		if ok {
			vs_cache_obj, found := vs_cache.(*utils.AviVsCache)
			if found {
				if vs_cache_obj.HTTPKeyCollection == nil {
					vs_cache_obj.HTTPKeyCollection = []utils.NamespaceName{k}
				} else {
					if !utils.HasElem(vs_cache_obj.HTTPKeyCollection, k) {
						vs_cache_obj.HTTPKeyCollection = append(vs_cache_obj.HTTPKeyCollection, k)
					}
				}
				utils.AviLog.Info.Printf("Modified the VS cache for https object. The cache now is :%v", utils.Stringify(vs_cache_obj))
			}

		} else {
			vs_cache_obj := utils.AviVsCache{Name: vsKey.Name, Tenant: vsKey.Namespace,
				HTTPKeyCollection: []utils.NamespaceName{k}}
			cache.VsCache.AviCacheAdd(vsKey, &vs_cache_obj)
			utils.AviLog.Info.Print(spew.Sprintf("Added VS cache key during http policy update %v val %v\n", vsKey,
				vs_cache_obj))
		}
		utils.AviLog.Info.Print(spew.Sprintf("Added Http Policy Set cache k %v val %v\n", k,
			http_cache_obj))
	}

	return nil
}

func AviHTTPCacheDel(cache *utils.AviObjCache, rest_op *utils.RestOp, vsKey utils.NamespaceName) error {
	key := utils.NamespaceName{Namespace: rest_op.Tenant, Name: rest_op.ObjName}
	cache.HTTPCache.AviCacheDelete(key)
	vs_cache, ok := cache.VsCache.AviCacheGet(vsKey)
	if ok {
		vs_cache_obj, found := vs_cache.(*utils.AviVsCache)
		if found {
			vs_cache_obj.HTTPKeyCollection = Remove(vs_cache_obj.HTTPKeyCollection, key)
		}
	}

	return nil
}
