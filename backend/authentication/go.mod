module github.com/nekizz/final-project/backend/authentication

go 1.14

replace github.com/nekizz/final-project/backend/go-pbtype => ./../go-pbtype

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-gormigrate/gormigrate/v2 v2.0.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/joho/godotenv v1.5.1
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nekizz/final-project/backend/go-pbtype v0.0.0-20221012040806-d8841d677437
	github.com/redis/go-redis/v9 v9.0.5
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	google.golang.org/grpc v1.48.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.8
)
