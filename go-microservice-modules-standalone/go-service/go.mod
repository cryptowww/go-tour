module go_service_1

go 1.16

require (
	dbmodel v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
	github.com/mattn/go-colorable v0.1.12
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)

replace dbmodel => ../dbmodel
