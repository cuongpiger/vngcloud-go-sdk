package snapshot

import lso "github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"

type (
	Snapshot struct {
		BackendID           string         `json:"backendId"`
		BackendSnapshotID   *string        `json:"backendSnapshotId"`
		BackendSnapshotName *string        `json:"backendSnapshotName"`
		CreatedAt           string         `json:"createdAt"`
		DeletedAt           *string        `json:"deletedAt"`
		ID                  string         `json:"id"`
		ParentID            *string        `json:"parentId"`
		ParentType          string         `json:"parentType"`
		ProjectID           string         `json:"projectId"`
		Size                int64          `json:"size"`
		SnapshotConfig      Config         `json:"snapshotConfig"`
		SnapshotTime        int64          `json:"snapshotTime"`
		SnapshotVolumeId    string         `json:"snapshotVolumeId"`
		Status              string         `json:"status"`
		UpdatedAt           string         `json:"updatedAt"`
		UserId              int64          `json:"userId"`
		VolumeId            string         `json:"volumeId"`
		VolumeSnapshot      VolumeSnapshot `json:"volumeSnapshot"`
		Name                string         `json:"name"`
		Description         string         `json:"description"`
	}

	Config struct {
		IsPermanently bool  `json:"isPermanently"`
		RetainedDays  int64 `json:"retainedDays"`
	}

	VolumeSnapshot struct {
		BackendPool      string  `json:"backendPool"`
		BackendPrefix    string  `json:"backendPrefix"`
		BackendStatus    string  `json:"backendStatus"`
		BackendUuid      string  `json:"backendUuid"`
		BootIndex        int     `json:"bootIndex"`
		Bootable         bool    `json:"bootable"`
		CreatedAt        string  `json:"createdAt"`
		Id               string  `json:"id"`
		MultiAttach      bool    `json:"multiAttach"`
		Name             string  `json:"name"`
		Product          string  `json:"product"`
		ProjectId        string  `json:"projectId"`
		Size             int64   `json:"size"`
		Status           string  `json:"status"`
		Type             string  `json:"type"`
		UpdatedAt        string  `json:"updatedAt"`
		Uuid             string  `json:"uuid"`
		VolumeId         int64   `json:"volumeId"`
		VolumeTypeId     string  `json:"volumeTypeId"`
		VolumeTypeZoneId string  `json:"volumeTypeZoneId"`
		EncryptionType   *string `json:"encryptionType"`
	}
)

/**
 * Define the response for the API List Snapshot of Block Storage volume
 */

type (
	ListVolumeSnapshotResponse struct {
		Items      []Snapshot `json:"items"`
		Page       int        `json:"page"`
		PageSize   int        `json:"pageSize"`
		TotalPages int        `json:"totalPages"`
		TotalItems int        `json:"totalItems"`
	}
)

func (s *Snapshot) toSnapshotObject() *lso.Snapshot {
	return &lso.Snapshot{
		ID:        s.ID,
		Name:      s.Name,
		CreatedAt: s.CreatedAt,
		VolumeID:  s.VolumeId,
		Status:    s.Status,
		Size:      s.Size,
	}
}

func (s *ListVolumeSnapshotResponse) ToSnapshotListObject() *lso.SnapshotList {
	if s == nil {
		return nil
	}

	sl := new(lso.SnapshotList)
	for _, item := range s.Items {
		sl.Items = append(sl.Items, *item.toSnapshotObject())
	}

	return sl
}
