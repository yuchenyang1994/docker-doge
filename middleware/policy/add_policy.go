package policy

import (
	"docker-doge/middleware"
)

// AddPolicyForUserGroups add policy
func AddPolicyForUserGroups(groupName string) {
	e := middleware.GetAuthzInstance()
	// 容器相关权限
	policys := CreatePolicys(groupName)
	for _, policy := range policys {
		e.AddPolicy(policy)

	}
}
