module github.com/nekizz/final-project/backend/user

go 1.14

replace (
	github.com/nekizz/final-project/backend/go-pbtype => ../go-pbtype
	github.com/nekizz/final-project/backend/hotel => ../hotel
	github.com/nekizz/final-project/backend/user => ../user
)

require (
	github.com/dinhtp/lets-go-company v0.0.0-20220726110114-1b63b6ca52be
	github.com/go-gormigrate/gormigrate/v2 v2.0.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/joho/godotenv v1.4.0
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nekizz/final-project/backend/go-pbtype v0.0.0-20221220043446-54e621f710d6
	github.com/nekizz/final-project/backend/hotel v0.0.0-20230106033736-a07e76bb4c2a
	github.com/redis/go-redis/v9 v9.0.4
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/valyala/fasthttp v1.47.0
	google.golang.org/grpc v1.48.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.24.5
)
