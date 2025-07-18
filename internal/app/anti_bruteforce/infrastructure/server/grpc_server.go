package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/gkarman/anti_bruteforce/api/pb"
	"github.com/gkarman/anti_bruteforce/internal/app/anti_bruteforce/infrastructure/app"
	"github.com/gkarman/anti_bruteforce/internal/config"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	cfg        config.GrpcServer
	app        app.AntiBruteForceApp
	grpcServer *grpc.Server
	pb.UnimplementedAntiBruteforceServiceServer
}

func NewGrpcServer(cfg config.GrpcServer, app app.AntiBruteForceApp) *GrpcServer {
	return &GrpcServer{
		cfg: cfg,
		app: app,
	}
}

func (s *GrpcServer) Start(_ context.Context) error {
	address := net.JoinHostPort(s.cfg.Host, s.cfg.Port)
	lsn, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)

	pb.RegisterAntiBruteforceServiceServer(s.grpcServer, s)

	log.Printf("gRPC server starting at " + address)
	if err := s.grpcServer.Serve(lsn); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		log.Printf("gRPC server failed: " + err.Error())
		return err
	}

	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	if s.grpcServer == nil {
		return nil
	}
	log.Println("Shutting down gRPS server...")

	// Graceful shutdown с таймаутом
	done := make(chan struct{})
	go func() {
		s.grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("gRPC server stopped gracefully")
		return nil
	case <-ctx.Done():
		log.Println("Timeout during gRPC shutdown, forcing stop")
		s.grpcServer.Stop()
		return ctx.Err()
	}
}

func (s *GrpcServer) IsCanLogin(ctx context.Context, req *pb.IsCanLoginRequest) (*pb.IsCanLoginResponse, error) {
	return &pb.IsCanLoginResponse{
		Ok: true,
	}, nil
}

func (s *GrpcServer) ClearBucket(ctx context.Context, req *pb.ClearBucketRequest) (*pb.ClearBucketResponse, error) {
	return &pb.ClearBucketResponse{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (s *GrpcServer) AddCIDRToBlackList(ctx context.Context, req *pb.AddCIDRToBlackListRequest) (*pb.AddCIDRToBlackListResponse, error) {
	return &pb.AddCIDRToBlackListResponse{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (s *GrpcServer) DeleteCIDRFromBlackList(ctx context.Context, req *pb.DeleteCIDRFromBlackListRequest) (*pb.DeleteCIDRFromBlackListResponse, error) {
	return &pb.DeleteCIDRFromBlackListResponse{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (s *GrpcServer) AddCIDRToWhiteList(ctx context.Context, req *pb.AddCIDRToWhiteListRequest) (*pb.AddCIDRToWhiteListResponse, error) {
	return &pb.AddCIDRToWhiteListResponse{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (s *GrpcServer) DeleteCIDRFromWhiteList(ctx context.Context, req *pb.DeleteCIDRFromWhiteListRequest) (*pb.DeleteCIDRFromWhiteListResponse, error) {
	return &pb.DeleteCIDRFromWhiteListResponse{
		Ok:      true,
		Message: "Success",
	}, nil
}
