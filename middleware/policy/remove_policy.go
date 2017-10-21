package policy

import (
	"docker-doge/middleware"
)

// RemovePolicyForUserGroups remove policy
func RemovePolicyForUserGroups(groupName string) {
	e := middleware.CreateAuthz()
	policys := CreatePolicys(groupName)
	for _, policy := range policys {
		e.RemovePolicy(policy)
	}
}
