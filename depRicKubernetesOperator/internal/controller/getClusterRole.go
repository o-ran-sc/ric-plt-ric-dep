package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)	

func GetClusterRole() []*rbacv1.ClusterRole {

	clusterRole1 := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: "svcacct-ricplt-appmgr-ricxapp-access",
		},
		Rules: []rbacv1.PolicyRule{

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
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"secrets",
				},
				Verbs: []string{

					"get",
					"list",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
	}

	clusterRole2 := &rbacv1.ClusterRole{
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"configmaps",
					"endpoints",
					"services",
				},
				Verbs: []string{

					"get",
					"list",
					"create",
					"update",
					"delete",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "svcacct-ricplt-appmgr-ricxapp-getappconfig",
		},
	}

	clusterRole3 := &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
			},
			Name: "release-name-kong",
		},
		Rules: []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				Resources: []string{

					"endpoints",
					"nodes",
					"pods",
					"secrets",
				},
				Verbs: []string{

					"list",
					"watch",
				},
				APIGroups: []string{

					"",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"nodes",
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

					"services",
				},
				Verbs: []string{

					"get",
					"list",
					"watch",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"extensions",
					"networking.k8s.io",
				},
				Resources: []string{

					"ingresses",
				},
				Verbs: []string{

					"get",
					"list",
					"watch",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"events",
				},
				Verbs: []string{

					"create",
					"patch",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"extensions",
					"networking.k8s.io",
				},
				Resources: []string{

					"ingresses/status",
				},
				Verbs: []string{

					"update",
				},
			},
			rbacv1.PolicyRule{
				Verbs: []string{

					"get",
					"list",
					"watch",
				},
				APIGroups: []string{

					"configuration.konghq.com",
				},
				Resources: []string{

					"kongplugins",
					"kongcredentials",
					"kongconsumers",
					"kongingresses",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRole",
		},
	}

	clusterRole4 := &rbacv1.ClusterRole{
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
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"endpoints",
					"nodes",
					"pods",
					"secrets",
				},
				Verbs: []string{

					"list",
					"watch",
				},
			},
			rbacv1.PolicyRule{
				Resources: []string{

					"nodes",
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
					"list",
					"watch",
				},
				APIGroups: []string{

					"",
				},
				Resources: []string{

					"services",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"extensions",
					"networking.k8s.io",
				},
				Resources: []string{

					"ingresses",
				},
				Verbs: []string{

					"get",
					"list",
					"watch",
				},
			},
			rbacv1.PolicyRule{
				Resources: []string{

					"events",
				},
				Verbs: []string{

					"create",
					"patch",
				},
				APIGroups: []string{

					"",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"extensions",
					"networking.k8s.io",
				},
				Resources: []string{

					"ingresses/status",
				},
				Verbs: []string{

					"update",
				},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{

					"configuration.konghq.com",
				},
				Resources: []string{

					"kongplugins",
					"kongcredentials",
					"kongconsumers",
					"kongingresses",
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
			Kind:       "ClusterRole",
		},
	}

	return []*rbacv1.ClusterRole{clusterRole1, clusterRole2, clusterRole3, clusterRole4}
}