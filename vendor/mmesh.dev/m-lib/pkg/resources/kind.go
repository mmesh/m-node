package resources

import "mmesh.dev/m-api-go/grpc/resources/resource"

var ObjectKindMap = map[resource.Kind]string{
	resource.Kind_SERVICE_ITEM:   "ServiceItem",
	resource.Kind_LOCATION:       "Location",
	resource.Kind_FEDERATION:     "Federation",
	resource.Kind_CONTROLLER:     "Controller",
	resource.Kind_ACCOUNT:        "Account",
	resource.Kind_REALM:          "Realm",
	resource.Kind_USER:           "User",
	resource.Kind_SECURITY_GROUP: "SecurityGroup",
	resource.Kind_ROLE:           "Role",
	resource.Kind_ACL:            "ACL",
	resource.Kind_PROJECT:        "Project",
	resource.Kind_WORKFLOW:       "Workflow",
	resource.Kind_OPERATION:      "Operation",
	resource.Kind_TENANT:         "Tenant",
	resource.Kind_NETWORK:        "Network",
	resource.Kind_VRF:            "VRF",
	resource.Kind_NODE:           "Node",
	resource.Kind_ROUTING_TABLE:  "RoutingTable",
}

var ObjSetKindMap = map[resource.Kind]string{
	resource.Kind_SERVICE_ITEM:   "serviceItems",
	resource.Kind_LOCATION:       "locations",
	resource.Kind_FEDERATION:     "federations",
	resource.Kind_CONTROLLER:     "controllers",
	resource.Kind_ACCOUNT:        "accounts",
	resource.Kind_REALM:          "realms",
	resource.Kind_USER:           "users",
	resource.Kind_SECURITY_GROUP: "securityGroups",
	resource.Kind_ROLE:           "roles",
	resource.Kind_ACL:            "acls",
	resource.Kind_PROJECT:        "projects",
	resource.Kind_WORKFLOW:       "workflows",
	resource.Kind_OPERATION:      "operations",
	resource.Kind_TENANT:         "tenants",
	resource.Kind_NETWORK:        "networks",
	resource.Kind_VRF:            "vrfs",
	resource.Kind_NODE:           "nodes",
	resource.Kind_ROUTING_TABLE:  "rt",
}
