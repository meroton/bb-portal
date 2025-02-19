package buildqueuestateproxy

import (
	"context"
	"log"

	"github.com/buildbarn/bb-remote-execution/pkg/proto/buildqueuestate"
	"github.com/buildbarn/bb-storage/pkg/auth"
	"github.com/buildbarn/bb-storage/pkg/digest"
	"google.golang.org/protobuf/types/known/emptypb"
)

// BuildQueueStateServerImpl is a gRPC server that forwards requests to a BuildQueueStateClient.
type BuildQueueStateServerImpl struct {
	client     buildqueuestate.BuildQueueStateClient
	authorizer auth.Authorizer
}

// NewBuildQueueStateServerImpl creates a new BuildQueueStateServerImpl from a
// given client. It also takes an authorizer to filter out the queues that the
// user is not allowed to see.
func NewBuildQueueStateServerImpl(client buildqueuestate.BuildQueueStateClient, authorizer auth.Authorizer) *BuildQueueStateServerImpl {
	return &BuildQueueStateServerImpl{client: client, authorizer: authorizer}
}

// GetOperation proxies GetOperation requests to the client.
func (s *BuildQueueStateServerImpl) GetOperation(ctx context.Context, req *buildqueuestate.GetOperationRequest) (*buildqueuestate.GetOperationResponse, error) {
	return s.client.GetOperation(ctx, req)
}

// ListOperations proxies ListOperations requests to the client.
func (s *BuildQueueStateServerImpl) ListOperations(ctx context.Context, req *buildqueuestate.ListOperationsRequest) (*buildqueuestate.ListOperationsResponse, error) {
	return s.client.ListOperations(ctx, req)
}

// KillOperations proxies KillOperations requests to the client.
func (s *BuildQueueStateServerImpl) KillOperations(ctx context.Context, req *buildqueuestate.KillOperationsRequest) (*emptypb.Empty, error) {
	return s.client.KillOperations(ctx, req)
}

// ListPlatformQueues proxies ListPlatformQueues requests to the client.
func (s *BuildQueueStateServerImpl) ListPlatformQueues(ctx context.Context, req *emptypb.Empty) (*buildqueuestate.ListPlatformQueuesResponse, error) {
	response, err := s.client.ListPlatformQueues(ctx, req)
	if err != nil {
		return nil, err
	}
	response.PlatformQueues = filterPlatormQueues(ctx, response, s.authorizer)
	return response, err
}

// ListInvocationChildren proxies ListInvocationChildren requests to the client.
func (s *BuildQueueStateServerImpl) ListInvocationChildren(ctx context.Context, req *buildqueuestate.ListInvocationChildrenRequest) (*buildqueuestate.ListInvocationChildrenResponse, error) {
	return s.client.ListInvocationChildren(ctx, req)
}

// ListQueuedOperations proxies ListQueuedOperations requests to the client.
func (s *BuildQueueStateServerImpl) ListQueuedOperations(ctx context.Context, req *buildqueuestate.ListQueuedOperationsRequest) (*buildqueuestate.ListQueuedOperationsResponse, error) {
	return s.client.ListQueuedOperations(ctx, req)
}

// ListWorkers proxies ListWorkers requests to the client.
func (s *BuildQueueStateServerImpl) ListWorkers(ctx context.Context, req *buildqueuestate.ListWorkersRequest) (*buildqueuestate.ListWorkersResponse, error) {
	return s.client.ListWorkers(ctx, req)
}

// TerminateWorkers proxies TerminateWorkers requests to the client.
func (s *BuildQueueStateServerImpl) TerminateWorkers(ctx context.Context, req *buildqueuestate.TerminateWorkersRequest) (*emptypb.Empty, error) {
	return s.client.TerminateWorkers(ctx, req)
}

// ListDrains proxies ListDrains requests to the client.
func (s *BuildQueueStateServerImpl) ListDrains(ctx context.Context, req *buildqueuestate.ListDrainsRequest) (*buildqueuestate.ListDrainsResponse, error) {
	return s.client.ListDrains(ctx, req)
}

// AddDrain proxies AddDrain requests to the client.
func (s *BuildQueueStateServerImpl) AddDrain(ctx context.Context, req *buildqueuestate.AddOrRemoveDrainRequest) (*emptypb.Empty, error) {
	return s.client.AddDrain(ctx, req)
}

// RemoveDrain proxies RemoveDrain requests to the client.
func (s *BuildQueueStateServerImpl) RemoveDrain(ctx context.Context, req *buildqueuestate.AddOrRemoveDrainRequest) (*emptypb.Empty, error) {
	return s.client.RemoveDrain(ctx, req)
}

func filterPlatormQueues(ctx context.Context, response *buildqueuestate.ListPlatformQueuesResponse, authorizer auth.Authorizer) []*buildqueuestate.PlatformQueueState {
	queues := response.GetPlatformQueues()
	// Filter out the queues that the user is not allowed to see.
	allowedQueues := make([]*buildqueuestate.PlatformQueueState, 0, len(queues))
	for _, queue := range queues {
		instanceNamePrefix := queue.GetName().GetInstanceNamePrefix()
		instanceName, err := digest.NewInstanceName(instanceNamePrefix)
		if err != nil {
			log.Println("Error parsing instance name from scheduler: ", err)
			continue
		}
		if auth.AuthorizeSingleInstanceName(ctx, authorizer, instanceName) == nil {
			allowedQueues = append(allowedQueues, queue)
		}
	}
	return allowedQueues
}
