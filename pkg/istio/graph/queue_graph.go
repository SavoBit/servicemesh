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

package graph

import (
	"fmt"
	"sync"
	"time"

	"github.com/avinetworks/servicemesh/pkg/utils"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
)

const (
	GraphLayer = "GraphLayer"
)

var queuewrapper sync.Once
var queueInstance *WorkQueueWrapper
var fixedQueues = [...]WorkerQueue{WorkerQueue{NumWorkers: utils.NumWorkers, WorkqueueName: GraphLayer, syncFunc: SyncFromGraphLayer}}

type WorkQueueWrapper struct {
	// This struct should manage a set of WorkerQueues for the various layers
	queueCollection map[string]*WorkerQueue
}

func (w *WorkQueueWrapper) GetQueueByName(queueName string) *WorkerQueue {
	workqueue, _ := w.queueCollection[queueName]
	return workqueue
}

func SharedWorkQueueWrappers() *WorkQueueWrapper {
	queuewrapper.Do(func() {
		queueInstance = &WorkQueueWrapper{}
		queueInstance.queueCollection = make(map[string]*WorkerQueue)
		for _, queue := range fixedQueues {
			workqueue := NewWorkQueue(queue.NumWorkers, queue.WorkqueueName, queue.syncFunc)
			queueInstance.queueCollection[queue.WorkqueueName] = workqueue
		}
	})
	return queueInstance
}

//Common utils like processing worker queue, that is common for all objects.
type WorkerQueue struct {
	NumWorkers    uint32
	Workqueue     []workqueue.RateLimitingInterface
	WorkqueueName string
	workerIdMutex sync.Mutex
	workerId      uint32
	syncFunc      func(string) error
}

func NewWorkQueue(num_workers uint32, workerQueueName string, syncFunc func(string) error) *WorkerQueue {
	queue := &WorkerQueue{}
	queue.Workqueue = make([]workqueue.RateLimitingInterface, num_workers)
	queue.workerId = (uint32(1) << num_workers) - 1
	queue.NumWorkers = num_workers
	queue.WorkqueueName = workerQueueName
	queue.syncFunc = syncFunc
	for i := uint32(0); i < num_workers; i++ {
		queue.Workqueue[i] = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), fmt.Sprintf("avi-%d", i))
	}
	return queue
}

func (c *WorkerQueue) Run(stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	utils.AviLog.Info.Printf("Starting workers to drain the %s layer queues", c.WorkqueueName)
	for i := uint32(0); i < c.NumWorkers; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	utils.AviLog.Info.Printf("Started the workers for: %s", c.WorkqueueName)

	return nil
}

func (c *WorkerQueue) StopWorkers(stopCh <-chan struct{}) {
	for i := uint32(0); i < c.NumWorkers; i++ {
		defer c.Workqueue[i].ShutDown()
	}
	utils.AviLog.Info.Printf("Shutting down the workers for %s", c.WorkqueueName)
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue. Pick a worker_id from worker_id mask
func (c *WorkerQueue) runWorker() {
	workerId := uint32(0xffffffff)
	c.workerIdMutex.Lock()
	for i := uint32(0); i < c.NumWorkers; i++ {
		if ((uint32(1) << i) & c.workerId) != 0 {
			workerId = i
			c.workerId = c.workerId & ^(uint32(1) << i)
			break
		}
	}
	c.workerIdMutex.Unlock()
	utils.AviLog.Info.Printf("Worker id %d", workerId)
	for c.processNextWorkItem(workerId) {
	}
	c.workerIdMutex.Lock()
	c.workerId = c.workerId | (uint32(1) << workerId)
	c.workerIdMutex.Unlock()
}

func (c *WorkerQueue) processNextWorkItem(worker_id uint32) bool {
	obj, shutdown := c.Workqueue[worker_id].Get()
	if shutdown {
		return false
	}
	var ok bool
	var ev string
	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.Workqueue[worker_id].Done(obj)
		if ev, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.Workqueue[worker_id].Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncToAvi, passing it the ev resource to be synced.
		err := c.syncFunc(ev)
		if err != nil {
			// TODO (sudswas): Do an add back logic via the retry layer here.
			utils.AviLog.Error.Printf("There was an error while syncing the key: %s", ev)
		}
		c.Workqueue[worker_id].Forget(obj)

		return nil
	}(obj)
	if err != nil {
		runtime.HandleError(err)
		return false
	}
	return true
}

func SyncFromGraphLayer(key string) error {
	// TBU
	return nil
}
