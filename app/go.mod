module kaitosub/app
go 1.13

replace (
	github.com/kaitosub/app => ../app
	github.com/kaitosub/app/infrastructure/mysql => ../app/infrastructure/mysql/mysql.go
	github.com/kaitosub/app/infrastructure/router => ../app/infrastructure/router/router.go
)