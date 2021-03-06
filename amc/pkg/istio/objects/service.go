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

package objects

import "sync"

//This package gives relationship APIs to manage a kubernetes service object.

var svclisterinstance *SvcLister
var svconce sync.Once

func SharedSvcLister() *SvcLister {
	svconce.Do(func() {
		svcVsStore := NewObjectStore()
		svcDrStore := NewObjectStore()
		svclisterinstance = &SvcLister{}
		svclisterinstance.svcVsStore = svcVsStore
		svclisterinstance.svcDrStore = svcDrStore
	})
	return svclisterinstance
}

type SvcLister struct {
	svcVsStore *ObjectStore
	svcDrStore *ObjectStore
}

type SvcNSCache struct {
	namespace    string
	svcVsobjects *ObjectMapStore
	svcDrObjects *ObjectMapStore
}

func (v *SvcLister) Service(ns string) *SvcNSCache {
	namespacedsvcVsObjs := v.svcVsStore.GetNSStore(ns)
	namespacedsvcDrObjs := v.svcDrStore.GetNSStore(ns)
	//svcInstance := SharedSvcLister()
	return &SvcNSCache{namespace: ns, svcVsobjects: namespacedsvcVsObjs, svcDrObjects: namespacedsvcDrObjs}
}

func (v *SvcNSCache) GetSvcToVS(svcName string) (bool, []string) {
	// Need checks if it's found or not?
	found, vsNames := v.svcVsobjects.Get(svcName)
	if !found {
		return false, nil
	}
	return true, vsNames.([]string)
}

func (v *SvcNSCache) DeleteSvcToVSMapping(svcName string) bool {
	// Need checks if it's found or not?
	success := v.svcVsobjects.Delete(svcName)
	return success
}

func (v *SvcNSCache) UpdateSvcToVSMapping(svcName string, vsList []string) {
	v.svcVsobjects.AddOrUpdate(svcName, vsList)
}

func (v *SvcNSCache) UpdateSvcToDR(svcName string, drList []string) {
	v.svcDrObjects.AddOrUpdate(svcName, drList)
}

func (v *SvcNSCache) GetSvcToDR(svcName string) (bool, []string) {
	// Need checks if it's found or not?
	found, drNames := v.svcDrObjects.Get(svcName)
	if !found {
		return false, nil
	}
	return true, drNames.([]string)
}

func (v *SvcNSCache) DeleteSvcToDRMapping(svcName string) bool {
	// Need checks if it's found or not?
	success := v.svcDrObjects.Delete(svcName)
	return success
}
