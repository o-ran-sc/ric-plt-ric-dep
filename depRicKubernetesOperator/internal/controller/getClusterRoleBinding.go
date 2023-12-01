package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetClusterRoleBinding() []*rbacv1.ClusterRoleBinding {
	clusterRoleBinding1 := &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-appmgr-ricxapp-access",
			Namespace: "ricplt",
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "ClusterRole",
			Name:     "svcacct-ricplt-appmgr-ricxapp-access",
			APIGroup: "rbac.authorization.k8s.io",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Namespace: "ricplt",
				Kind:      "ServiceAccount",
				Name:      "svcacct-ricplt-appmgr",
			},
		},
	}

	clusterRoleBinding2 := &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-appmgr-ricxapp-getappconfig",
			Namespace: "ricxapp",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "svcacct-ricplt-appmgr-ricxapp-getappconfig",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Namespace: "ricplt",
				Kind:      "ServiceAccount",
				Name:      "svcacct-ricplt-appmgr",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "ClusterRoleBinding",
		},
	}
	return []*rbacv1.ClusterRoleBinding{clusterRoleBinding1, clusterRoleBinding2}
}
