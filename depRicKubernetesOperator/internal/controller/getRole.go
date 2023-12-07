package controller

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

	role2 := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
				"app.kubernetes.io/instance":   "release-name",
			},
			Name: "release-name-kong",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				Resources: []string{

					"configmaps",
					"pods",
					"secrets",
					"namespaces",
				},
				Verbs: []string{

					"get",
				},
				APIGroups: []string{

					"",
				},
			},
			rbacv1.PolicyRule{
				Verbs: []string{

					"get",
					"update",
				},
				APIGroups: []string{

					"",
				},
				ResourceNames: []string{

					"kong-ingress-controller-leader-kong-kong",
				},
				Resources: []string{

					"configmaps",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"configmaps",
				},
				Verbs: []string{

					"create",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"endpoints",
				},
				Verbs: []string{

					"get",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
	}

	role3 := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ricxapp-tiller-base",
			Namespace: "ricxapp",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				ResourceNames: []string{

					"ricxapp-tiller-secret",
				},
				Resources: []string{

					"secrets",
				},
				Verbs: []string{

					"get",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"pods/portforward",
				},
				Verbs: []string{

					"create",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"namespaces",
				},
				Verbs: []string{

					"get",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"pods",
					"configmaps",
					"deployments",
					"services",
				},
				Verbs: []string{

					"get",
					"list",
					"create",
					"delete",
				},
			},
		},
	}

	role4 := &rbacv1.Role{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ricxapp-tiller-operation",
			Namespace: "ricinfra",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"configmaps",
				},
				Verbs: []string{

					"get",
					"list",
					"create",
					"delete",
					"update",
				},
			},
		},
	}

	role5 := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ricxapp-tiller-deployer",
			Namespace: "ricxapp",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"pods",
					"configmaps",
					"services",
				},
				Verbs: []string{

					"get",
					"list",
					"create",
					"delete",
				},
			},
			rbacv1.PolicyRule{
				Verbs: []string{

					"get",
					"list",
					"create",
					"delete",
				},
				APIGroups: []string{

					"extensions",
					"apps",
				},
				Resources: []string{

					"deployments",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
	}

	role6 := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tiller-secret-creator-xzhjjg-secret-create",
			Namespace: "ricinfra",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"secrets",
				},
				Verbs: []string{

					"create",
					"get",
					"patch",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
	}

	role7 := &rbacv1.Role{
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"configmaps",
					"pods",
					"secrets",
					"namespaces",
				},
				Verbs: []string{

					"get",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				ResourceNames: []string{

					"kong-ingress-controller-leader-kong-kong",
				},
				Resources: []string{

					"configmaps",
				},
				Verbs: []string{

					"get",
					"update",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"configmaps",
				},
				Verbs: []string{

					"create",
				},
			},
			rbacv1.PolicyRule{
				Resources: []string{

					"endpoints",
				},
				Verbs: []string{

					"get",
				},
				APIGroups: []string{

					"",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "Role",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-kong",
			Labels: map[string]string{
				"helm.sh/chart":                "kong-0.36.6",
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
			},
		},
	}

	role8 := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-o1mediator-ricxapp-podreader",
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

	role9 := &rbacv1.Role{
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
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"pods/exec",
				},
				Verbs: []string{

					"create",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Role",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "assigner-role",
		},
	}

	return []*rbacv1.Role{role1, role2, role3, role4, role5, role6, role7, role8, role9}
}