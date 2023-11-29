
func GetRole() []*rbacv1.Role {

	role1 := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-alarmmanager-ricxapp-podreader",
			Namespace: "ricxapp",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"pods",
				},
				Verbs: []string{

					"get",
					"list",
					"watch",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
	}
	return []*rbacv1.Role{role1}
}