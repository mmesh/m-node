package router

import (
	"mmesh.dev/m-api-go/grpc/network/routing"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
	"mmesh.dev/m-node/internal/app/node/mnet/router/rib"
)

func (r *router) eventProcessor(closeCh <-chan struct{}) {
	for {
		select {
		case nh := <-r.RIB().RelayConnQueue():
			if err := r.rconnect(nh); err != nil {
				xlog.Errorf("Unable to connect to relay: %v", errors.Cause(err))
			}
		case nh := <-r.RIB().RouterConnQueue():
			if err := r.rconnect(nh); err != nil {
				xlog.Errorf("Unable to connect to router: %v", errors.Cause(err))
			}
		case evt := <-r.RIB().RouteEventQueue():
			switch evt.Type {
			case rib.RouteEventTypeADD:
				if !r.localForwarding || r.networkInterface == nil {
					continue
				}

				if evt.RouteType == routing.RouteType_STATIC {
					if !r.isValidRouteImport(evt.Addr) {
						continue
					}
				}

				if err := r.networkInterface.routeAdd(cidrIPDst(evt.Addr)); err != nil {
					xlog.Errorf("Unable to add route: %v", errors.Cause(err))
				}
			case rib.RouteEventTypeDELETE:
				if !r.localForwarding || r.networkInterface == nil {
					continue
				}

				if err := r.networkInterface.routeDel(cidrIPDst(evt.Addr)); err != nil {
					xlog.Errorf("Unable to remove route: %v", errors.Cause(err))
				}
			}
		case <-closeCh:
			return
		}
	}
}
