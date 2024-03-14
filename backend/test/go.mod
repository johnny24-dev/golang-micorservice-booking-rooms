module test

go 1.18

replace (
	github.com/nekizz/final-project/backend/hotel => ../hotel
	github.com/nekizz/final-project/backend/user => ../user
	github.com/nekizz/final-project/backend/wishlist => ../wishlist
)

require (
	github.com/go-sql-driver/mysql v1.7.0
	github.com/gofiber/fiber/v2 v2.41.0
	github.com/nekizz/final-project/backend/user v0.0.0-20230225103444-d1676a06ef2b
	github.com/valyala/fasthttp v1.47.0
	github.com/xuri/excelize/v2 v2.7.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.5
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.16.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/nekizz/final-project/backend/hotel v0.0.0-20230106033736-a07e76bb4c2a // indirect
	github.com/richardlehane/mscfb v1.0.4 // indirect
	github.com/richardlehane/msoleps v1.0.3 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xuri/efp v0.0.0-20220603152613-6918739fd470 // indirect
	github.com/xuri/nfp v0.0.0-20220409054826-5e722a1d9e22 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
