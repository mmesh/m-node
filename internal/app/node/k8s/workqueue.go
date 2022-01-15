package k8s

import (
	"time"

	// "k8s.io/klog"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	util_runtime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"mmesh.dev/m-lib/pkg/xlog"
)

type eventType string

const (
	eventAdd    eventType = "ADD"
	eventUpdate eventType = "UPDATE"
	eventDelete eventType = "DELETE"
)

type controller struct {
	informer cache.Controller
	indexer  cache.Indexer
	queue    workqueue.RateLimitingInterface
}

type objEvent struct {
	key string
	obj interface{}
	old interface{}
	evt eventType
}

func (c *controller) processNextItem() bool {
	// Wait until there is a new item in the working queue
	objEvt, quit := c.queue.Get()
	if quit {
		return false
	}
	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
	// This allows safe parallel processing because two pods with the same key are never processed in
	// parallel.
	defer c.queue.Done(objEvt)

	// Invoke the method containing the business logic
	// err := c.syncToStdout(key.(string))

	var err error
	oevt := objEvt.(*objEvent)
	switch o := oevt.obj.(type) {
	case *v1.Service:
		err = c.manageSvcEvent(o, oevt.evt)
	case *v1.Pod:
		// err = c.managePodEvent(o, oevt.evt)
	}

	// Handle the error if something went wrong during the execution of the business logic
	c.handleErr(err, objEvt)

	return true
}

/*
// syncToStdout is the business logic of the controller. In this controller it simply prints
// information about the pod to stdout. In case an error happened, it has to simply return the error.
// The retry logic should not be part of the business logic.
func (c *controller) syncToStdout(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		// Below we will warm up our cache with a Service, so that we will see a delete for one pod
		fmt.Printf("Pod %s does not exist anymore\n", key)
	} else {
		// Note that you also have to check the uid if you have a local controlled resource, which
		// is dependent on the actual instance, to detect that a Service was recreated with the same name
		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Service).GetName())
	}

	return nil
}
*/

// handleErr checks if an error happened and makes sure we will retry later.
func (c *controller) handleErr(err error, objEvt interface{}) {
	if err == nil {
		// Forget about the #AddRateLimited history of the objEvt on every successful synchronization.
		// This ensures that future processing of updates for this objEvt is not delayed because of
		// an outdated error history.
		c.queue.Forget(objEvt)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(objEvt) < 5 {
		xlog.Warnf("Retrying to sync object %s: %v", objEvt.(*objEvent).key, err)

		// Re-enqueue the objEvt rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the objEvt will be processed later again.
		c.queue.AddRateLimited(objEvt)
		return
	}

	c.queue.Forget(objEvt)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	// runtime.HandleError(err)
	// runtime.Hanror(err)
	xlog.Errorf("Dropping object %s out of the queue: %v", objEvt.(*objEvent).key, err)
}

func (c *controller) run(name string, threadiness int, stopCh chan struct{}) {
	defer util_runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	xlog.Infof("Starting kubernetes %s controller", name)

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		// runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		// runtime.HandleError(fmt.Errorf("Timed out waiting for cach sync"))
		xlog.Errorf("[kubernetes %s controller] Timed out waiting for caches to sync", name)
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	xlog.Infof("Stopping kubernetes %s controller", name)
}

func (c *controller) runWorker() {
	for c.processNextItem() {
	}
}

func KubernetesController(quitCh chan struct{}) error {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// svc watcher
	svcListWatcher := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		string(v1.ResourceServices),
		v1.NamespaceAll,
		fields.Everything(),
	)
	// svc controller
	svcController := newController(svcListWatcher, &v1.Service{})

	// pod watcher
	podListWatcher := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		string(v1.ResourcePods),
		v1.NamespaceAll,
		fields.Everything(),
	)
	// pod controller
	podController := newController(podListWatcher, &v1.Pod{})

	// Now let's start the svcController
	svcStop := make(chan struct{})
	defer close(svcStop)
	go svcController.run("svc", 1, svcStop)

	// Now let's start the podController
	podStop := make(chan struct{})
	defer close(podStop)
	go podController.run("pod", 1, podStop)

	<-quitCh

	return nil
}

func newController(lw cache.ListerWatcher, objType runtime.Object) *controller {
	// create the workqueue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
	// whenever the cache is updated, the svc key is added to the workqueue.
	// Note that when we finally process the item from the workqueue, we might see a newer version
	// of the Service than the version which was responsible for triggering the update.
	indexer, informer := cache.NewIndexerInformer(lw, objType, 0, cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(&objEvent{
					key: key,
					obj: obj,
					evt: eventAdd,
				})
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			if err == nil {
				queue.Add(&objEvent{
					key: key,
					obj: new,
					old: old,
					evt: eventUpdate,
				})
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
			// key function.
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(&objEvent{
					key: key,
					obj: obj,
					evt: eventDelete,
				})
			}
		},
	}, cache.Indexers{})

	// We can now warm up the cache for initial synchronization.
	// Let's suppose that we knew about a pod "mypod" on our last run, therefore add it to the cache.
	// If this pod is not there anymore, the controller will be notified about the removal after the
	// cache has synchronized.
	indexer.Add(&v1.Service{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      "mysvc",
			Namespace: v1.NamespaceDefault,
		},
	})

	return &controller{
		informer: informer,
		indexer:  indexer,
		queue:    queue,
	}
}
