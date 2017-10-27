package db

import (
	"github.com/jinzhu/gorm"
)

// DockerMachine ...
type DockerMachine struct {
	gorm.Model
	Domain     string `gorm:"size:100"`
	Containers []DockerContainer
}

// DockerContainer ...
type DockerContainer struct {
	gorm.Model
	DomainMachineID uint
	CID             string `gorm:"size:255"`
	Names           string `gorm:"size:200"`
	Image           string `gorm:"size:200"`
	Command         string `gorm:"size:200"`
}
