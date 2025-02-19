package fsacproxy

import (
	"context"

	"github.com/buildbarn/bb-storage/pkg/proto/fsac"
	"google.golang.org/protobuf/types/known/emptypb"
)

// FsacServerImpl is a gRPC server that forwards requests to an FileSystemAccessCacheClient.
type FsacServerImpl struct {
	client fsac.FileSystemAccessCacheClient
}

// NewFsacServerImpl creates a new FsacServerImpl from a given client.
func NewFsacServerImpl(client fsac.FileSystemAccessCacheClient) *FsacServerImpl {
	return &FsacServerImpl{client: client}
}

// GetFileSystemAccessProfile proxies GetFileSystemAccessProfile requests to the client.
func (s *FsacServerImpl) GetFileSystemAccessProfile(ctx context.Context, req *fsac.GetFileSystemAccessProfileRequest) (*fsac.FileSystemAccessProfile, error) {
	return s.client.GetFileSystemAccessProfile(ctx, req)
}

// UpdateFileSystemAccessProfile proxies UpdateFileSystemAccessProfile requests to the client.
func (s *FsacServerImpl) UpdateFileSystemAccessProfile(ctx context.Context, req *fsac.UpdateFileSystemAccessProfileRequest) (*emptypb.Empty, error) {
	return s.client.UpdateFileSystemAccessProfile(ctx, req)
}
