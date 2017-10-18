package validators

type AddRoleValidator struct {
	RoleName string `json:"role" binding:"required,ADMIN|USER|LEADER"`
	UserID   int    `json:"userId" binding:"required"`
}
