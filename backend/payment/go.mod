module github.com/nekizz/final-project/backend/payment

go 1.14

replace github.com/nekizz/final-project/backend/go-pbtype => ../go-pbtype

require (
	github.com/Shopify/sarama v1.38.1
	github.com/ThreeDotsLabs/watermill v1.2.0
	github.com/ThreeDotsLabs/watermill-kafka/v2 v2.2.2
	github.com/go-gormigrate/gormigrate/v2 v2.0.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nekizz/final-project/backend/booking v0.0.0-20230330100303-6facc6898a88
	github.com/nekizz/final-project/backend/go-pbtype v0.0.0-20221220043446-54e621f710d6
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.13.0
	google.golang.org/grpc v1.50.0
	gorm.io/driver/mysql v1.4.1
	gorm.io/gorm v1.24.5
)
