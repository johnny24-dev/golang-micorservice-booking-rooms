module github.com/nekizz/final-project/backend/hotel

go 1.14

replace (
	github.com/nekizz/final-project/backend/booking => ../booking
	github.com/nekizz/final-project/backend/go-pbtype => ../go-pbtype/
	github.com/nekizz/final-project/backend/user => ../user
)

require (
	github.com/go-gormigrate/gormigrate/v2 v2.0.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nekizz/final-project/backend/go-pbtype v0.0.0-20221220043446-54e621f710d6
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	golang.org/x/crypto v0.7.0 // indirect
	google.golang.org/grpc v1.48.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.24.5
)
