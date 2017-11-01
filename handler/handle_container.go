package handler

import (
	"docker-doge/db"

	"docker-doge/handler/validators"

	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// AddNewDockerMathine ...
func AddNewDockerMathine(c *gin.Context) {
	d := db.GetDbInstance(c)
	var machine validators.DockerMathineValidator
	if err := c.ShouldBindWith(&machine, binding.JSON); err == nil {
		ma := db.DockerMachine{Domain: machine.Domain}
		if has := d.NewRecord(&ma); has {
			d.Create(&ma)
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		} else {
			c.JSON(405, gin.H{"message": "domain excited"})
		}
	} else {
		c.JSON(403, gin.H{"message": "valid error", "error": err.Error()})
	}
}

// GetAllDockerMathine ...
func GetAllDockerMathine(c *gin.Context) {
	d := db.GetDbInstance(c)
	machines := []db.DockerMachine{}
	domains := []string{}
	for _, machine := range machines {
		domains = append(domains, machine.Domain)
	}
	d.Find(&machines)
	c.JSON(http.StatusOK, gin.H{"domains": domains})
}

// GetAllContainersWithDomain ...
func GetAllContainersWithDomain(c *gin.Context) {
	domainKey := c.Param("dockerDomain")
	d := db.GetDbInstance(c)
	domain := db.DockerMachine{}
	containers := []db.DockerContainer{}
	d = d.First(&domain, "domain = ?", domainKey)
	if notFound := d.RecordNotFound(); notFound {
		c.JSON(404, gin.H{"message": "not found"})
	} else {
		d.Model(&domain).Related(&containers)
		c.JSON(200, gin.H{"containers": containers})
	}
}

// GetAllContainersForMachine ...
func GetAllContainersForMachine(c *gin.Context) {
	domainKey := c.Param("dockerDomain")
	d := db.GetDbInstance(c)
	domain := db.DockerMachine{}
	d = d.First(&domain, "domain = ?", domainKey)
	if notFound := d.RecordNotFound(); notFound {
		c.JSON(404, gin.H{"message": "not found"})
	} else {
		if dockerCli, err := docker.NewClient(domain.Domain); err == nil {
			containers, err := dockerCli.ListContainers(docker.ListContainersOptions{All: false})
			if err != nil {
				c.JSON(200, gin.H{"containers": containers})
			} else {
				c.JSON(401, gin.H{"error": err.Error()})
			}
		}

	}
}
