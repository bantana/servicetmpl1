// Package cacheclient is the wrapper around the third party gRPC Cache Micro-service.
// It encapsulates the logic to call outside service, to make it transparent to the business logic layer.

package cacheclient

import (
	"context"
	"github.com/jfeng45/servicetmpl1/app/logger"
	"github.com/jfeng45/servicetmpl1/applicationservice/cacheclient/generatedclient"
	"google.golang.org/grpc"
)

// CacheDataGrpc represents the gRPC connection handler
type CacheDataGrpc struct {
	Conn *grpc.ClientConn
}

// getCacheClient creates a gRPC client
func getCacheClient(conn *grpc.ClientConn) generatedclient.CacheServiceClient {
	return generatedclient.NewCacheServiceClient(conn)
}

// Get makes a call to Get function on Cache service
func (cdg CacheDataGrpc) Get(key string) ([]byte, error) {
	cacheClient := getCacheClient(cdg.Conn)
	resp, err := cacheClient.Get(context.Background(), &generatedclient.GetReq{Key: key})
	if err != nil {
		return nil, err
	} else {
		return resp.Value, err
	}
}

// Store makes a call to Store function on Cache service
func (cdg CacheDataGrpc) Store(key string, value []byte) error {
	cacheClient := getCacheClient(cdg.Conn)
	ctx := context.Background()
	_, err := cacheClient.Store(ctx, &generatedclient.StoreReq{Key: key, Value: value})

	if err != nil {
		return err
	} else {
		logger.Log.Debug("store called")
	}
	return nil
}
