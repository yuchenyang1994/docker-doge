package validators

type SingupVlidator struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	UserGroupID uint   `json:"userGroupId" binding:"required"`
}
