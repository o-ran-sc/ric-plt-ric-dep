package controller

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	roleBinding2 := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name:      "release-name-kong",
			Namespace: "ricplt",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "release-name-kong",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "release-name-kong",
				Namespace: "ricplt",
			},
		},
	}

	roleBinding3 := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-tiller-ricxapp-ricxapp-tiller-base",
			Namespace: "ricxapp",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "ricxapp-tiller-base",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "svcacct-tiller-ricxapp",
				Namespace: "ricinfra",
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
	}

	roleBinding4 := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricinfra",
			Name:      "svcacct-tiller-ricxapp-ricxapp-tiller-operation",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "ricxapp-tiller-operation",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "svcacct-tiller-ricxapp",
				Namespace: "ricinfra",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
	}

	roleBinding5 := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricxapp",
			Name:      "svcacct-tiller-ricxapp-ricxapp-tiller-deployer",
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "Role",
			Name:     "ricxapp-tiller-deployer",
			APIGroup: "rbac.authorization.k8s.io",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "svcacct-tiller-ricxapp",
				Namespace: "ricinfra",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
	}

	roleBinding6 := &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricinfra",
			Name:      "tiller-secret-creator-xzhjjg-secret-create",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "tiller-secret-creator-xzhjjg-secret-create",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "tiller-secret-creator-xzhjjg",
				Namespace: "ricinfra",
			},
		},
	}

	roleBinding7 := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name:      "release-name-kong",
			Namespace: "ricplt",
		},
		RoleRef: rbacv1.RoleRef{
			Name:     "release-name-kong",
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
		},
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Namespace: "ricplt",
				Kind:      "ServiceAccount",
				Name:      "release-name-kong",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
	}

	roleBinding8 := &rbacv1.RoleBinding{
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind:      "ServiceAccount",
				Name:      "svcacct-ricplt-o1mediator",
				Namespace: "ricplt",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-o1mediator-ricxapp-podreader",
			Namespace: "ricxapp",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "svcacct-ricplt-o1mediator-ricxapp-podreader",
		},
	}

	roleBinding9 := &rbacv1.RoleBinding{
		Subjects: []rbacv1.Subject{

			rbacv1.Subject{
				Kind: "ServiceAccount",
				Name: "assigner-sa",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "assigner-rb",
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "assigner-role",
		},
	}

	return []*rbacv1.RoleBinding{roleBinding1, roleBinding2, roleBinding3, roleBinding4, roleBinding5, roleBinding6, roleBinding7, roleBinding8, roleBinding9}
}
