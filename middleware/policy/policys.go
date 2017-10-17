package policy

import (
	"docker-doge/middleware"
	"fmt"
)

// CreatePolicys create policys slice
func CreatePolicys(groupName string) [][]string {
	containerDomin := fmt.Sprintf("%s/contarners*", groupName)
	policys := [][]string{
		// ------ contaners ------
		// admin
		[]string{groupName, middleware.ROLE_ADMIN, containerDomin, "GET"},
		[]string{groupName, middleware.ROLE_ADMIN, containerDomin, "POST"},
		[]string{groupName, middleware.ROLE_ADMIN, containerDomin, "PUT"},
		[]string{groupName, middleware.ROLE_ADMIN, containerDomin, "DELETE"},
		// leader
		[]string{groupName, middleware.ROLE_LEADER, containerDomin, "GET"},
		[]string{groupName, middleware.ROLE_LEADER, containerDomin, "POST"},
		[]string{groupName, middleware.ROLE_LEADER, containerDomin, "PUT"},
		[]string{groupName, middleware.ROLE_ADMIN, containerDomin, "GET"},
		// ------ auth -------
		[]string{groupName, middleware.ROLE_ADMIN, "/auth*", "GET"},
		[]string{groupName, middleware.ROLE_ADMIN, "/auth*", "POST"},
		[]string{groupName, middleware.ROLE_ADMIN, "/auth*", "PUT"},
		// user and leader
		[]string{groupName, middleware.ROLE_USER, "/auth*", "GET"},
		[]string{groupName, middleware.ROLE_LEADER, "/auth*", "GET"},
	}
	return policys
}
