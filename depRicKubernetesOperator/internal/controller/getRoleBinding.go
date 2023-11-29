package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)


func GetRoleBinding() []*rbacv1.RoleBinding {

	roleBinding1 := &rbacv1.RoleBinding{
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "svcacct-ricplt-alarmmanager",
				Namespace: "ricplt",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-alarmmanager-ricxapp-podreader",
			Namespace: "ricxapp",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "svcacct-ricplt-alarmmanager-ricxapp-podreader",
		},
	}
	
	return []*rbacv1.RoleBinding{roleBinding1, roleBinding2, roleBinding3, roleBinding4, roleBinding5, roleBinding6, roleBinding7, roleBinding8, roleBinding9}
}