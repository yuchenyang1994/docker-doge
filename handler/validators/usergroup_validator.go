package validators

// UserGroupVlidator Vlidator for UserGroup
type UserGroupVlidator struct {
	GroupName string `json:"groupName" binding:"required"`
}

// ChangeGroupNameVlidator ...
type ChangeGroupNameVlidator struct {
	GroupName    string `json:"groupName" binding:"required"`
	NewGroupName string `json:"newGroupName" binding:"required"`
}
