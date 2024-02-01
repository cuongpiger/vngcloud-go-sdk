package server

import (
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/common"
	lServerV2 "github.com/vngcloud/vngcloud-go-sdk/vngcloud/services/compute/v2"
)

type DataDiskEncryptionType string

const (
	DataDiskEncryptionAesXtsType DataDiskEncryptionType = "aes-xts-plain64_256"
)

// ************************************************** GetOptsBuilder ***************************************************

type GetOpts struct {
	common.CommonOpts
	lServerV2.ServerV2Common
}

// ************************************************* DeleteOptsBuilder *************************************************

type DeleteOpts struct {
	DeleteAllVolume bool `json:"deleteAllVolume"`

	common.CommonOpts
	lServerV2.ServerV2Common
}

func (s *DeleteOpts) ToRequestBody() interface{} {
	return s
}

// ************************************************* CreateOptsBuilder *************************************************

type CreateOpts struct {
	AttachFloating         bool                   `json:"attachFloating,omitempty"`
	BackupInstancePointId  string                 `json:"backupInstancePointId,omitempty"`
	DataDiskEncryptionType DataDiskEncryptionType `json:"dataDiskEncryptionType,omitempty"`
	DataDiskName           string                 `json:"dataDiskName,omitempty"`
	DataDiskSize           int                    `json:"dataDiskSize,omitempty"`
	DataDiskTypeId         string                 `json:"dataDiskTypeId,omitempty"`
	EnableBackup           bool                   `json:"enableBackup,omitempty"`
	EncryptionVolume       bool                   `json:"encryptionVolume"`
	ExpirePassword         bool                   `json:"expirePassword,omitempty"`
	FlavorId               string                 `json:"flavorId"`
	ImageId                string                 `json:"imageId"`
	Name                   string                 `json:"name"`
	NetworkId              string                 `json:"networkId"`
	OsLicence              bool                   `json:"osLicence,omitempty"`
	RestoreBackup          bool                   `json:"restoreBackup,omitempty"`
	RootDiskEncryptionType DataDiskEncryptionType `json:"rootDiskEncryptionType,omitempty"`
	RootDiskSize           int                    `json:"rootDiskSize"`
	RootDiskTypeId         string                 `json:"rootDiskTypeId"`
	SecurityGroup          []string               `json:"securityGroup,omitempty"`
	ServerGroupId          string                 `json:"serverGroupId,omitempty"`
	SshKeyId               string                 `json:"sshKeyId,omitempty"`
	SubnetId               string                 `json:"subnetId"`
	UserData               string                 `json:"userData,omitempty"`
	UserDataBase64Encoded  bool                   `json:"userDataBase64Encoded,omitempty"`
	UserName               string                 `json:"userName,omitempty"`
	UserPassword           string                 `json:"userPassword,omitempty"`

	common.CommonOpts
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}
