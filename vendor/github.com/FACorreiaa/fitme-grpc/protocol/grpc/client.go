package grpc

import (
	"context"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/FACorreiaa/fitme-grpc/protocol/grpc/middleware/grpclog"
	"github.com/FACorreiaa/fitme-grpc/protocol/grpc/middleware/grpcprometheus"
	"github.com/FACorreiaa/fitme-grpc/protocol/grpc/middleware/grpcspan"
)

func BootstrapClient(
	address string,
	log *zap.Logger,
	traceProvider trace.TracerProvider,
	prometheus *prometheus.Registry,
	opts ...grpc.DialOption,
) (*grpc.ClientConn, error) {
	// -- OpenTelemetry interceptor setup
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	//

	spanInterceptor, _ := grpcspan.Interceptors()

	// -- Zap logging interceptor setup
	logInterceptor, _ := grpclog.Interceptors(log)

	// trace to grafana
	ctx := context.Background()
	if err := grpcprometheus.SetupTracing(ctx); err != nil {
		log.Error("Failed to set up trace exporter", zap.Error(err))
		return nil, errors.Wrap(err, "failed to setup tracing")
	}

	// -- Prometheus exporter setup
	prometheusCollectors := grpcprometheus.NewPrometheusMetricsCollectors()
	if err := grpcprometheus.RegisterMetrics(prometheus, prometheusCollectors); err != nil {
		return nil, errors.Wrap(err, "failed to register grpc metrics")
	}

	clientMetrics := prometheusCollectors.Client

	connOptions := []grpc.DialOption{
		// We terminate TLS in the linkerd sidecar, so no need for TLS on the listen server
		grpc.WithTransportCredentials(insecure.NewCredentials()),

		// Add the unary interceptors
		grpc.WithChainUnaryInterceptor(
			spanInterceptor.Unary,
			logInterceptor.Unary,
			clientMetrics.UnaryClientInterceptor(),
		),

		// Add the stream interceptors
		grpc.WithChainStreamInterceptor(
			spanInterceptor.Stream,
			logInterceptor.Stream,
			clientMetrics.StreamClientInterceptor(),
		),
	}

	// Add any additional options
	connOptions = append(connOptions, opts...)

	return grpc.NewClient(address, connOptions...)
}
