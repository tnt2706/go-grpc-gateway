package server

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"

	"grpc-gateway/internal/service"
	"grpc-gateway/pkg/pb"

	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in the allowed header, don't send the header
	return s, false
}

func setupRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})

	e.GET("/custom", func(c echo.Context) error {
		return c.String(http.StatusOK, "Custom route")
	})

	e.GET("/console", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/console/")
	})

	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "")
	})

}

func StartGrpcServer() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, service.NewCalculatorService())
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	ec := echo.New()

	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux(
		// convert header in response(going from gateway) from metadata received.
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			header := request.Header.Get("Authorization")
			// send all the headers received from the client
			md := metadata.Pairs("auth", header)
			return md
		}),
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			//creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}),
		// Custom response default camelCase
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(1024*1024*20), grpc.MaxCallSendMsgSize(1024*1024*20))}

	// setting up a dail up for gRPC service by specifying endpoint/target url
	err = pb.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, "0.0.0.0:8080", opts)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	setupRoutes(ec)

	ec.Any("/api/*", func(c echo.Context) error {
		mux.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	})

	if err := ec.Start(":8090"); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}

func main() {
	StartGrpcServer()
}
