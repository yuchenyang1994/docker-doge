package db

import (
	"github.com/jinzhu/gorm"
)

// DockerMachine ...
type DockerMachine struct {
	gorm.Model
	Domain     string `gorm:"size:100" json:"domain"`
	Containers []DockerContainer
}

// DockerContainer ...
type DockerContainer struct {
	gorm.Model
	DomainMachineID uint
	CID             string `gorm:"size:255" json:"containerId"`
	Name            string `gorm:"size:200" json:"name"`
	Image           string `gorm:"size:200" json:"image"`
	Command         string `gorm:"size:200" json:"command"`
}
