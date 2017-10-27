package validators

// DockerMathineValidator ...
type DockerMathineValidator struct {
	Domain string `json:"domain" binding:"required"`
}
