package loadbalancer

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lbCm "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/loadbalancer/v2"
)

// ************************************************* CreateOptsBuilder *************************************************

const (
	// For internet facing
	CreateOptsSchemeOptInternet CreateOptsSchemeOpt = "Internet"
	CreateOptsSchemeOptInternal CreateOptsSchemeOpt = "Internal"

	// For lb type
	CreateOptsTypeOptLayer4 CreateOptsTypeOpt = "Layer 4"
	CreateOptsTypeOptLayer7 CreateOptsTypeOpt = "Layer 7"
)

type (
	CreateOptsSchemeOpt string
	CreateOptsTypeOpt   string
)

type CreateOpts struct {
	Name               string              `json:"name"`
	PackageID          string              `json:"packageId"`
	Scheme             CreateOptsSchemeOpt `json:"scheme"`
	SubnetID           string              `json:"subnetId"`
	Type               CreateOptsTypeOpt   `json:"type"`
	CreateOptsListener *CreateOptsListener `json:"listener"`

	common.CommonOpts
}

type CreateOptsListener struct {
	AllowCidrs                  string    `json:"allowedCidrs"`
	CertificateAuthorities      *[]string `json:"certificateAuthorities,omitempty"`
	ClientCertificate           *string   `json:"clientCertificate,omitempty"`
	DefaultCertificateAuthority *string   `json:"defaultCertificateAuthority,omitempty"`
	Headers                     *[]string `json:"headers,omitempty"`
	ListenerName                string    `json:"listenerName"`
	ListenerProtocol            string    `json:"listenerProtocol"`
	ListenerProtocolPort        int       `json:"listenerProtocolPort"`
	TimeoutClient               int       `json:"timeoutClient"`
	TimeoutConnection           int       `json:"timeoutConnection"`
	TimeoutMember               int       `json:"timeoutMember"`
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}

// ************************************************** GetOptsBuilder ***************************************************

type GetOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ************************************************* DeleteOptsBuilder *************************************************

type DeleteOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ********************************************* GetBySubnetIDOptsBuilder **********************************************

type ListBySubnetIDOpts struct {
	SubnetID string
	common.CommonOpts
}

func (s *ListBySubnetIDOpts) GetSubnetID() string {
	return s.SubnetID
}

// ************************************************** ListOptsBuilder **************************************************

type ListOpts struct {
	common.CommonOpts
}

// ************************************************* UpdateOptsBuilder *************************************************

type UpdateOpts struct {
	PackageID string `json:"packageId"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

func (s *UpdateOpts) ToRequestBody() interface{} {
	return s
}
