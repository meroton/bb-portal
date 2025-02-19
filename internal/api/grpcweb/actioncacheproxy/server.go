package actioncacheproxy

import (
	"context"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
)

// ActionCacheServerImpl is a gRPC server that forwards requests to an ActionCacheClient.
type ActionCacheServerImpl struct {
	client remoteexecution.ActionCacheClient
}

// NewAcctionCacheServerImpl creates a new ActionCacheServerImpl from a given client.
func NewAcctionCacheServerImpl(client remoteexecution.ActionCacheClient) *ActionCacheServerImpl {
	return &ActionCacheServerImpl{client: client}
}

// GetActionResult proxies GetActionResult requests to the client.
func (s *ActionCacheServerImpl) GetActionResult(ctx context.Context, req *remoteexecution.GetActionResultRequest) (*remoteexecution.ActionResult, error) {
	return s.client.GetActionResult(ctx, req)
}

// UpdateActionResult proxies UpdateActionResult requests to the client.
func (s *ActionCacheServerImpl) UpdateActionResult(ctx context.Context, req *remoteexecution.UpdateActionResultRequest) (*remoteexecution.ActionResult, error) {
	return s.client.UpdateActionResult(ctx, req)
}
