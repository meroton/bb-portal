package isccproxy

import (
	"context"

	"github.com/buildbarn/bb-storage/pkg/proto/iscc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// IsccServerImpl is a gRPC server that forwards requests to an InitialSizeClassCacheClient.
type IsccServerImpl struct {
	client iscc.InitialSizeClassCacheClient
}

// NewIsccServerImpl creates a new IsccServerImpl from a given client.
func NewIsccServerImpl(client iscc.InitialSizeClassCacheClient) *IsccServerImpl {
	return &IsccServerImpl{client: client}
}

// GetPreviousExecutionStats proxies GetPreviousExecutionStats requests to the client.
func (s *IsccServerImpl) GetPreviousExecutionStats(ctx context.Context, req *iscc.GetPreviousExecutionStatsRequest) (*iscc.PreviousExecutionStats, error) {
	return s.client.GetPreviousExecutionStats(ctx, req)
}

// UpdatePreviousExecutionStats proxies UpdatePreviousExecutionStats requests to the client.
func (s *IsccServerImpl) UpdatePreviousExecutionStats(ctx context.Context, req *iscc.UpdatePreviousExecutionStatsRequest) (*emptypb.Empty, error) {
	return s.client.UpdatePreviousExecutionStats(ctx, req)
}
