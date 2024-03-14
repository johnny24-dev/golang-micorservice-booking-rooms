package cmd

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nekizz/final-project/backend/authentication/authentication"
	pb "github.com/nekizz/final-project/backend/go-pbtype/authentication"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
)

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "authentication service serve grpc command",
	Run:   runGrpcCommand,
}

func init() {
	serveCmd.AddCommand(grpcCmd)

	grpcCmd.Flags().StringP("backend", "", "grpc-address", "gRPC address")
	grpcCmd.Flags().StringP("mysqlDsn", "", "mysql-dsn", "mysql connection string")

	_ = viper.BindPFlag("backend", grpcCmd.Flags().Lookup("backend"))
	_ = viper.BindPFlag("mysqlDsn", grpcCmd.Flags().Lookup("mysqlDsn"))
}

func runGrpcCommand(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// init DB Connection
	mysqlChan := make(chan *gorm.DB, 1)
	go initializeDbConnection("mysqlDsn", c, mysqlChan)
	orm := <-mysqlChan

	//init config.env
	err := godotenv.Load("/app/config.env")
	if err != nil {
		log.Fatal("Some error occured. Err: %s", err)
	}

	// services
	grpcServer := grpc.NewServer()
	grpcServer = initializeServices(orm, grpcServer)

	// init GRPC backend
	grpcAddr := viper.GetString("backend")
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(err)
	}

	// Serve GRPC
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	logrus.WithFields(logrus.Fields{
		"service": "authentication",
		"type":    "grpc",
	}).Info("authentication service server started")

	<-c
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logrus.WithFields(logrus.Fields{
		"service": "authentication-service",
		"type":    "grpc",
	}).Info("authentication service gracefully shutdowns")

}

func initializeDbConnection(mysqlDsnField string, c chan os.Signal, mysqlChan chan *gorm.DB) {
	mysqlDsn := viper.GetString(mysqlDsnField)
	orm, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if nil != err {
		fmt.Println(err)
		c <- syscall.SIGTERM
	}

	sqlDB, err := orm.DB()
	if nil != err {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	fmt.Println(fmt.Sprintf("MySQL connection established"))

	mysqlChan <- orm
}

func initializeServices(orm *gorm.DB, grpcServer *grpc.Server) *grpc.Server {
	authService := authentication.NewService(orm)
	pb.RegisterAuthenticationServiceServer(grpcServer, authService)

	return grpcServer
}
