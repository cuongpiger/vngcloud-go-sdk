package network

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
	"github.com/vngcloud/vngcloud-go-sdk/vngcloud/objects"
)

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*objects.Network, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: response,
	})

	if err != nil {
		return nil, err
	}

	return response.ToNetworkObject(), nil
}
