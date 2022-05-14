package common

import (
	"gorm.io/gorm"
)

type Serializer interface {
	Create() *gorm.DB
	Retrive() *gorm.DB
	Validate() error
}
