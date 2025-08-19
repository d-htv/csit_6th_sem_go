package common

import "gorm.io/gorm"

type ApiServer struct{
	DB gorm.DB
}

var MyApiServer = ApiServer{}