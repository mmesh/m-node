package rib

import (
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

func (r *ribData) GetNodeAppSvcs() []*topology.AppSvc {
	r.RLock()
	defer r.RUnlock()

	appSvcs := make([]*topology.AppSvc, 0)

	for _, as := range r.appSvcs {
		appSvcs = append(appSvcs, as)
	}

	return appSvcs
}

func (r *ribData) AddNodeAppSvc(as *topology.AppSvc) {
	r.Lock()
	defer r.Unlock()

	r.appSvcs[nodeAppSvcID(as.AppSvcID)] = as
}

func (r *ribData) RemoveNodeAppSvc(appSvcID string) {
	r.Lock()
	defer r.Unlock()

	delete(r.appSvcs, nodeAppSvcID(appSvcID))
}
