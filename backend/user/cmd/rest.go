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

	pba "github.com/nekizz/final-project/backend/go-pbtype/account"
	pbc "github.com/nekizz/final-project/backend/go-pbtype/customer"
	pb "github.com/nekizz/final-project/backend/go-pbtype/user"
	pbw "github.com/nekizz/final-project/backend/go-pbtype/wishlist"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "user service serve rest command",
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

	corsMux := cors.New(cors.Options{
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
		Handler:      corsMux,
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
		"service": "user-service",
		"type":    "rest",
	}).Info("user service server started")

	<-c
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logrus.WithFields(logrus.Fields{
		"service": "user-service",
		"type":    "rest",
	}).Info("user service gracefully shutdowns")
}

func initializeGatewayService(ctx context.Context, gw *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	userGwErr := pb.RegisterUserServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != userGwErr {
		fmt.Println(userGwErr)
		os.Exit(1)
	}

	customerGwErr := pbc.RegisterCustomerServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != customerGwErr {
		fmt.Println(customerGwErr)
		os.Exit(1)
	}

	accountGwErr := pba.RegisterAccountServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != accountGwErr {
		fmt.Println(accountGwErr)
		os.Exit(1)
	}

	wishlistGwErr := pbw.RegisterWishlistServiceHandlerFromEndpoint(ctx, gw, endpoint, opts)
	if nil != wishlistGwErr {
		fmt.Println(wishlistGwErr)
		os.Exit(1)
	}
}
