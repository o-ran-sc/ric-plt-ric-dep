package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	return []*rbacv1.ClusterRole{clusterRole1, clusterRole2}

}