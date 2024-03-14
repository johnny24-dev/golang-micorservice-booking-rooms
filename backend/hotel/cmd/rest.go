package cmd

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	pba "github.com/nekizz/final-project/backend/go-pbtype/address"
	pbc "github.com/nekizz/final-project/backend/go-pbtype/comment"
	pb "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	pbs "github.com/nekizz/final-project/backend/go-pbtype/service"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "hotel service serve rest command",
	Run:   runRestCommand,
}

func init() {
	serveCmd.AddCommand(restCmd)

	restCmd.Flags().StringP("backend", "", "backend", "grpc backend address")
	restCmd.Flags().StringP("address", "", "address", "rest gateway address")

	_ = viper.BindPFlag("address", restCmd.Flags().Lookup("address"))
	_ = viper.BindPFlag("backend", restCmd.Flags().Lookup("backend"))
}

func runRestCommand(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	address := viper.GetString("address")
	backend := viper.GetString("backend")

	mux := http.NewServeMux()
	GatewayMux := runtime.NewServeMux()

	go func() {
		opts := []grpc.DialOption{grpc.WithInsecure()}
		initializeGatewayService(ctx, GatewayMux, backend, opts)
	}()

	a := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(mux)

	mux.Handle("/", GatewayMux)
	srv := &http.Server{
		Addr:         address,
		Handler:      a,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if nil != err {
			panic(err)
		}
	}()

	logrus.WithFields(logrus.Fields{
		"service": "hotel-service",
		"type":    "rest",
	}).Info("hotel service server started")

	<-c
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logrus.WithFields(logrus.Fields{
		"service": "hotel-service",
		"type":    "rest",
	}).Info("hotel service gracefully shutdowns")
}

//func cors(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
//
//		if r.Method == "OPTIONS" {
//			return
//		}
//		h.ServeHTTP(w, r)
//	})
//}

func initializeGatewayService(ctx context.Context, gw *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	hotelGwErr := pb.RegisterHotelServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != hotelGwErr {
		fmt.Println(hotelGwErr)
		os.Exit(1)
	}

	addressGwErr := pba.RegisterAddressServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != addressGwErr {
		fmt.Println(addressGwErr)
		os.Exit(1)
	}

	commentGwErr := pbc.RegisterCommentServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != commentGwErr {
		fmt.Println(commentGwErr)
		os.Exit(1)
	}

	serviceGwErr := pbs.RegisterServiceServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != serviceGwErr {
		fmt.Println(serviceGwErr)
		os.Exit(1)
	}
}
