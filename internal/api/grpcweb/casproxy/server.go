package casproxy

import (
	"context"
	"io"

	"google.golang.org/genproto/googleapis/bytestream"
)

// CasServerImpl is an gRPC server that forwards requests to a ByteStreamClient.
type CasServerImpl struct {
	client bytestream.ByteStreamClient
}

// NewCasServerImpl creates a new CasServerImpl from a given client.
func NewCasServerImpl(client bytestream.ByteStreamClient) *CasServerImpl {
	return &CasServerImpl{client: client}
}

// Read proxies Read requests to the client.
func (s *CasServerImpl) Read(req *bytestream.ReadRequest, stream bytestream.ByteStream_ReadServer) error {
	clientStream, err := s.client.Read(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		resp, err := clientStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

// Write proxies Write requests to the client.
func (s *CasServerImpl) Write(stream bytestream.ByteStream_WriteServer) error {
	clientStream, err := s.client.Write(context.Background())
	if err != nil {
		return err
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp, err := clientStream.CloseAndRecv()
			if err != nil {
				return err
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			return err
		}
		if err := clientStream.Send(req); err != nil {
			return err
		}
	}
}

// QueryWriteStatus proxies QueryWriteStatus requests to the client.
func (s *CasServerImpl) QueryWriteStatus(ctx context.Context, req *bytestream.QueryWriteStatusRequest) (*bytestream.QueryWriteStatusResponse, error) {
	return s.client.QueryWriteStatus(ctx, req)
}
