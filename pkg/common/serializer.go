package common

import (
	"gorm.io/gorm"
)

type Serializer interface {
	Create() *gorm.DB
	Validate() error
}
