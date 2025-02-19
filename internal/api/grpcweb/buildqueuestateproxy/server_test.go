package buildqueuestateproxy

import (
	"context"
	"io"
	"log"
	"testing"

	"github.com/buildbarn/bb-remote-execution/pkg/proto/buildqueuestate"
	"github.com/buildbarn/bb-storage/pkg/auth"
	auth_pb "github.com/buildbarn/bb-storage/pkg/proto/auth"
	"github.com/jmespath/go-jmespath"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestFilterPlatformQueues(t *testing.T) {
	a := auth.NewJMESPathExpressionAuthorizer(
		jmespath.MustCompile("contains(authenticationMetadata.private.permittedInstanceNames, instanceName) || instanceName == ''"),
	)

	ctx := auth.NewContextWithAuthenticationMetadata(context.Background(), auth.MustNewAuthenticationMetadataFromProto(&auth_pb.AuthenticationMetadata{
		Private: structpb.NewStructValue(&structpb.Struct{
			Fields: map[string]*structpb.Value{
				"permittedInstanceNames": structpb.NewListValue(&structpb.ListValue{
					Values: []*structpb.Value{
						structpb.NewStringValue("allowed"),
					},
				}),
			},
		}),
	}))

	t.Run("NoPlatformQueues", func(t *testing.T) {
		platformQueues := buildqueuestate.ListPlatformQueuesResponse{
			PlatformQueues: []*buildqueuestate.PlatformQueueState{},
		}
		allowedQueues := filterPlatormQueues(ctx, &platformQueues, a)
		if len(allowedQueues) != 0 {
			t.Errorf("Expected no platform queues, got %d", len(allowedQueues))
		}
	})

	t.Run("FilterQueues", func(t *testing.T) {
		platformQueues := buildqueuestate.ListPlatformQueuesResponse{
			PlatformQueues: []*buildqueuestate.PlatformQueueState{
				{
					Name: &buildqueuestate.PlatformQueueName{
						InstanceNamePrefix: "allowed",
					},
				},
				{
					Name: &buildqueuestate.PlatformQueueName{
						InstanceNamePrefix: "forbidden",
					},
				},
			},
		}
		allowedQueues := filterPlatormQueues(ctx, &platformQueues, a)
		if len(allowedQueues) != 1 {
			t.Errorf("Expected one platform queue, got %d", len(allowedQueues))
		}
		expected := platformQueues.PlatformQueues[0]
		if allowedQueues[0] != expected {
			t.Errorf("Expected platform queue %+v, got %+v", expected, allowedQueues[0])
		}
	})

	t.Run("AllowEmptyInstanceNames", func(t *testing.T) {
		platformQueues := buildqueuestate.ListPlatformQueuesResponse{
			PlatformQueues: []*buildqueuestate.PlatformQueueState{
				{
					Name: &buildqueuestate.PlatformQueueName{
						InstanceNamePrefix: "",
					},
				},
				{
					Name: &buildqueuestate.PlatformQueueName{
						InstanceNamePrefix: "forbidden",
					},
				},
			},
		}
		allowedQueues := filterPlatormQueues(ctx, &platformQueues, a)
		if len(allowedQueues) != 1 {
			t.Errorf("Expected one platform queue, got %d", len(allowedQueues))
		}
		expected := platformQueues.PlatformQueues[0]
		if allowedQueues[0] != expected {
			t.Errorf("Expected platform queue %+v, got %+v", expected, allowedQueues[0])
		}
	})

	t.Run("InvalidPlatformQueue", func(t *testing.T) {
		log.SetOutput(io.Discard)
		platformQueues := buildqueuestate.ListPlatformQueuesResponse{
			PlatformQueues: []*buildqueuestate.PlatformQueueState{
				{
					Name: &buildqueuestate.PlatformQueueName{
						InstanceNamePrefix: "asdff//////DF////",
					},
				},
			},
		}
		allowedQueues := filterPlatormQueues(ctx, &platformQueues, a)
		if len(allowedQueues) != 0 {
			t.Errorf("Expected no platform queues, got %d", len(allowedQueues))
		}
	})
}
