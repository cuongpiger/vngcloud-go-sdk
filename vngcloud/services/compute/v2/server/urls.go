package server

import (
	"github.com/vngcloud/vngcloud-go-sdk/client"
)

func getServerURL(pSc *client.ServiceClient, pOpts IGetOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"servers",
		pOpts.GetServerID())
}

func deleteServerURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"servers",
		pOpts.GetServerID())
}

func createServerURL(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"servers")
}

func listURL(pSc *client.ServiceClient, pOpts IListOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"servers")
}

func updateSecGroupsServerURL(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) string {
	return pSc.ServiceURL(
		pOpts.GetProjectID(),
		"servers",
		pOpts.GetServerID(),
		"update-sec-group")
}
