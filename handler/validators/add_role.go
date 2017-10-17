package validators

type AddRoleValidator struct {
	UserID uint   `json:"userId" binding:"required"`
	Role   string `json:"role binding:"required"`
}
