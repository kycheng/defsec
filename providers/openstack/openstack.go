package openstack

import "github.com/aquasecurity/defsec/parsers/types"

type OpenStack struct {
	types.Metadata
	Compute Compute
}

type Compute struct {
	types.Metadata
	Instances []Instance
	Firewall  Firewall
}

type Firewall struct {
	types.Metadata
	AllowRules []Rule
	DenyRules  []Rule
}

type Rule struct {
	types.Metadata
	Source          types.StringValue
	Destination     types.StringValue
	SourcePort      types.StringValue
	DestinationPort types.StringValue
	Enabled         types.BoolValue
}

type Instance struct {
	types.Metadata
	AdminPassword types.StringValue
}