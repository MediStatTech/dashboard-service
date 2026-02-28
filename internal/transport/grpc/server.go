package grpc

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	pb "github.com/MediStatTech/dashboard-client/pb/go/services/v1"
	"github.com/MediStatTech/dashboard-service/internal/app"
	s_options "github.com/MediStatTech/dashboard-service/internal/app/options"
	grpc_auth "github.com/MediStatTech/dashboard-service/internal/transport/grpc/auth"
	grpc_diseas "github.com/MediStatTech/dashboard-service/internal/transport/grpc/diseas"
	"github.com/MediStatTech/dashboard-service/internal/transport/grpc/middleware"
	grpc_patient "github.com/MediStatTech/dashboard-service/internal/transport/grpc/patient"
	grpc_staff "github.com/MediStatTech/dashboard-service/internal/transport/grpc/staff"
	"github.com/MediStatTech/dashboard-service/pkg"
)

type Server struct {
	addr         string
	server       *grpc.Server
	healthServer *health.Server
}

func New(p *pkg.Facade, appInstance *app.Facade) (*Server, error) {
	grpcTLS := grpc.ServerOption(grpc.EmptyServerOption{})
	if p.Config.TLSCertFilePath != "" && p.Config.TLSKeyFilePath != "" {
		p.Logger.Info("TLS enabled", map[string]any{
			"cert_file": p.Config.TLSCertFilePath,
			"key_file":  p.Config.TLSKeyFilePath,
		})
		creds, err := credentials.NewServerTLSFromFile(p.Config.TLSCertFilePath, p.Config.TLSKeyFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create new server tls from file: %w", err)
		}
		grpcTLS = grpc.Creds(creds)
	}

	server := grpc.NewServer(
		grpcTLS,
		grpc.MaxRecvMsgSize(1024*1024*50*50),
		grpc.MaxSendMsgSize(1024*1024*50),
		grpc.ConnectionTimeout(24*time.Hour),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     24 * time.Hour,
			MaxConnectionAge:      24 * time.Hour,
			MaxConnectionAgeGrace: 24 * time.Hour,
			Time:                  4 * time.Hour,
			Timeout:               4 * time.Hour,
		}),
		grpc.UnaryInterceptor(middleware.AuthInterceptor(appInstance.JwtService)),
	)

	opts := &s_options.Options{
		App: appInstance,
		PKG: p,
	}

	// Register gRPC services
	authHandler := grpc_auth.New(opts)
	pb.RegisterAuthServiceServer(server, authHandler)

	staffHandler := grpc_staff.New(opts)
	pb.RegisterStaffServiceServer(server, staffHandler)

	patientHandler := grpc_patient.New(opts)
	pb.RegisterPatientServiceServer(server, patientHandler)

	diseasHandler := grpc_diseas.New(opts)
	pb.RegisterDiseasServiceServer(server, diseasHandler)

	// Health checks
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	healthServer.SetServingStatus("services.v1.AuthService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("services.v1.StaffService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("services.v1.PatientService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("services.v1.DiseasService", grpc_health_v1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	reflection.Register(server)

	addr := ":50051"

	return &Server{
		addr:         addr,
		server:       server,
		healthServer: healthServer,
	}, nil
}

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to create a listener on %s: %w", s.addr, err)
	}

	if err = s.server.Serve(lis); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	select {
	case <-ctx.Done():
		s.server.Stop()
		return fmt.Errorf("forced shutdown due to context: %w", ctx.Err())
	default:
		done := make(chan struct{})
		go func() {
			s.server.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			return nil
		case <-ctx.Done():
			s.server.Stop()
			return fmt.Errorf("shutdown timeout: %w", ctx.Err())
		}
	}
}

func (s *Server) Address() string {
	return s.addr
}
