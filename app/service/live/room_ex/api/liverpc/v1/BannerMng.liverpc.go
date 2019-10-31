// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/BannerMng.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/BannerMng.proto
	v1/Banner.proto
	v1/RoomNews.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// ===================
// BannerMng Interface
// ===================

type BannerMng interface {
	// * 获取新后台配置的banner
	//
	GetNewBanner(context.Context, *BannerMngGetNewBannerReq) (*BannerMngGetNewBannerResp, error)
}

// =========================
// BannerMng Live Rpc Client
// =========================

type bannerMngRpcClient struct {
	client *liverpc.Client
}

// NewBannerMngRpcClient creates a Rpc client that implements the BannerMng interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewBannerMngRpcClient(client *liverpc.Client) BannerMng {
	return &bannerMngRpcClient{
		client: client,
	}
}

func (c *bannerMngRpcClient) GetNewBanner(ctx context.Context, in *BannerMngGetNewBannerReq) (*BannerMngGetNewBannerResp, error) {
	out := new(BannerMngGetNewBannerResp)
	err := doRpcRequest(ctx, c.client, 1, "BannerMng.getNewBanner", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRpcRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message) (err error) {
	err = client.Call(ctx, version, method, in, out)
	return
}
