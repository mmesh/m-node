package resources

import "mmesh.dev/m-api-go/grpc/resources/resource"

var ObjectKindMap = map[resource.Kind]string{
	// providerAPI
	resource.Kind_WEBHOOK:    "Webhook",
	resource.Kind_LOCATION:   "Location",
	resource.Kind_FEDERATION: "Federation",
	resource.Kind_CONTROLLER: "Controller",
	resource.Kind_ACCOUNT:    "Account",

	// coreAPI
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

	// billingAPI
	resource.Kind_CUSTOMER:         "Customer", // stripe customer
	resource.Kind_SUBSCRIPTION:     "Subscription",
	resource.Kind_BILLED_ITEM:      "BilledItem",
	resource.Kind_INVOICE:          "Invoice",
	resource.Kind_PAYMENT_INTENT:   "PaymentIntent",   // stripe payment_intent
	resource.Kind_CHECKOUT_SESSION: "CheckoutSession", // stripe checkout session

	// servicesAPI
	resource.Kind_SERVICE:      "Service",
	resource.Kind_PRICING_PLAN: "PricingPlan",
	resource.Kind_PROVIDER:     "Provider",
	resource.Kind_PRODUCT:      "Product",
	resource.Kind_PRICE:        "Price",

	resource.Kind_ITSM_INCIDENT: "Incident",
	resource.Kind_ITSM_REQUEST:  "Request",

	resource.Kind_CLOUD_COMPUTE_INSTANCE:   "ComputeInstance",
	resource.Kind_CLOUD_APP:                "CloudApp",
	resource.Kind_CLOUD_KUBERNETES_CLUSTER: "KubernetesCluster",
	resource.Kind_MANAGED_SERVICE_CONTRACT: "ServiceContract",
}

var ObjSetKindMap = map[resource.Kind]string{
	// providerAPI
	resource.Kind_WEBHOOK:    "webhooks",
	resource.Kind_LOCATION:   "locations",
	resource.Kind_FEDERATION: "federations",
	resource.Kind_CONTROLLER: "controllers",
	resource.Kind_ACCOUNT:    "accounts",

	// coreAPI
	resource.Kind_USER:           "iam:users",
	resource.Kind_SECURITY_GROUP: "iam:securityGroups",
	resource.Kind_ROLE:           "iam:roles",
	resource.Kind_ACL:            "iam:acls",
	resource.Kind_PROJECT:        "projects",
	resource.Kind_WORKFLOW:       "workflows",
	resource.Kind_OPERATION:      "operations",
	resource.Kind_TENANT:         "tenants",
	resource.Kind_NETWORK:        "networks",
	resource.Kind_VRF:            "vrfs",
	resource.Kind_NODE:           "nodes",
	resource.Kind_ROUTING_TABLE:  "rt",

	// billingAPI
	resource.Kind_CUSTOMER:         "customers", // stripe customer
	resource.Kind_SUBSCRIPTION:     "billing:subscriptions",
	resource.Kind_BILLED_ITEM:      "billing:items",
	resource.Kind_INVOICE:          "billing:invoices",
	resource.Kind_PAYMENT_INTENT:   "paymentIntents",   // stripe payment_intent
	resource.Kind_CHECKOUT_SESSION: "checkoutSessions", // stripe checkout sessions

	// servicesAPI
	resource.Kind_SERVICE:      "services",
	resource.Kind_PRICING_PLAN: "pricingPlans",
	resource.Kind_PROVIDER:     "providers",
	resource.Kind_PRODUCT:      "products",
	resource.Kind_PRICE:        "prices",

	resource.Kind_ITSM_INCIDENT: "itsm:incidents",
	resource.Kind_ITSM_REQUEST:  "itsm:requests",

	resource.Kind_CLOUD_COMPUTE_INSTANCE:   "services:cloud:computeInstances",
	resource.Kind_CLOUD_APP:                "services:cloud:apps",
	resource.Kind_CLOUD_KUBERNETES_CLUSTER: "services:cloud:k8sClusters",
	resource.Kind_MANAGED_SERVICE_CONTRACT: "services:managed:serviceContracts",
}
