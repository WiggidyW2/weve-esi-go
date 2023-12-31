package client

import (
	"context"

	"github.com/WiggidyW/weve-esi/client/crude_client"
	"github.com/WiggidyW/weve-esi/client/crude_client/response"
	"github.com/WiggidyW/weve-esi/client/db"
)

const NULL_AUTH = ""

type Client struct {
	Inner *crude_client.CrudeClient
	Db    db.Db
	// proto.UnimplementedWeveEsiServer
}

func NewClient(run_local bool) *Client {
	return &Client{
		Inner: crude_client.NewCrudeClient(run_local),
		Db:    db.NewDb(),
	}
}

func (c *Client) crudeRequestNoCache(
	ctx context.Context,
	url string,
	method string,
	auth string,
) (*response.EsiResponse, error) {
	return c.Inner.RequestNoCache(ctx, url, method, auth)
}

func (c *Client) crudeRequest(
	ctx context.Context,
	url string,
	method string,
	auth string,
) (*response.EsiResponse, error) {
	return c.Inner.Request(ctx, url, method, auth)
}

func (c *Client) crudeRequestNoArray(
	ctx context.Context,
	url string,
	method string,
	auth string,
) (*response.EsiResponse, error) {
	return c.Inner.RequestNoArray(ctx, url, method, auth)
}

func (c *Client) crudeRequestHead(
	ctx context.Context,
	url string,
	auth string,
) (*response.EsiHeadResponse, error) {
	return c.Inner.RequestHead(ctx, url, auth)
}

func (c *Client) crudeRequestAuth(
	ctx context.Context,
	token string,
) (string, error) {
	return c.Inner.RequestAuth(ctx, token)
}

func (c *Client) dbGetRegionId(
	ctx context.Context,
	system_id uint64,
) (int, error) {
	return c.Db.GetRegionId(ctx, system_id)
}

func (c *Client) dbGetStationSystemId(
	ctx context.Context,
	station_id uint64,
) (uint32, error) {
	return c.Db.GetStationSystemId(ctx, station_id)
}

func (c *Client) dbGetSystemRegionId(
	ctx context.Context,
	system_id uint32,
) (uint32, error) {
	return c.Db.GetSystemRegionId(ctx, system_id)
}
