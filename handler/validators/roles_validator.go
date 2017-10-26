package validators

type AddRoleValidator struct {
	RoleName string `json:"role" binding:"required"`
	UserID   int    `json:"userId" binding:"required"`
}
