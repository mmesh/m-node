package router

import (
	"net"
	"strings"

	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/xlog"
)

func (r *router) isValidRouteImport(route string) bool {
	for _, exportedRoute := range r.routes.exported {
		if exportedRoute == route {
			return false
		}
	}

	if strings.HasPrefix(route, "0.0.0.0") {
		return false
	}

	reply, err := isConnectedRoute(route)
	if err != nil {
		xlog.Errorf("Unable to check connected routes: %v", errors.Cause(err))
		return false
	}
	if reply {
		return false
	}

	for _, r := range r.routes.imported {
		if strings.ToLower(r) == "any" || r == route {
			return true
		}
	}

	return false
}

func isConnectedRoute(route string) (bool, error) {
	// check if route is directly connected
	connectedRoutes, err := getConnectedRoutes()
	if err != nil {
		return true, errors.Wrapf(err, "[%v] function getConnectedRoutes()", errors.Trace())
	}

	for _, ipNet := range connectedRoutes {
		xlog.Tracef("Checking connected route %s", ipNet.String())

		if ipNet.String() == route {
			return true, nil
		}
	}

	return false, nil
}

func getConnectedRoutes() ([]*net.IPNet, error) {
	connectedRoutes := make([]*net.IPNet, 0)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function net.InterfaceAddrs()", errors.Trace())
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok {
			connectedRoutes = append(connectedRoutes, ipNet)
		}
	}

	return connectedRoutes, nil
}
