module github.com/nekizz/final-project/backend/room

go 1.14

replace (
	github.com/nekizz/final-project/backend/booking => ../booking
	github.com/nekizz/final-project/backend/go-pbtype => ../go-pbtype
	github.com/nekizz/final-project/backend/used-service => ../used-service
)

require (
	github.com/Shopify/sarama v1.38.1 // indirect
	github.com/ThreeDotsLabs/watermill v1.2.0 // indirect
	github.com/ThreeDotsLabs/watermill-kafka/v2 v2.2.2
	github.com/go-gormigrate/gormigrate/v2 v2.0.2
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nekizz/final-project/backend/go-pbtype v0.0.0-20230301100314-746ac53c06f4
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	google.golang.org/grpc v1.53.0
	gorm.io/driver/mysql v1.4.7
	gorm.io/driver/postgres v1.4.5 // indirect
	gorm.io/gorm v1.24.5
)
