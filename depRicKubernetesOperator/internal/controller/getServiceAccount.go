package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetServiceAccount() []*corev1.ServiceAccount {

	serviceAccount1 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-alarmmanager",
			Namespace: "ricplt",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount2 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Name:      "svcacct-ricplt-appmgr",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount3 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name: "release-name-kong",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount4 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"helm.sh/hook":               "pre-upgrade",
				"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded",
			},
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name: "release-name-kong",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount5 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-tiller-ricxapp",
			Namespace: "ricinfra",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount6 := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tiller-secret-creator-xzhjjg",
			Namespace: "ricinfra",
		},
	}

	serviceAccount7 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name: "release-name-kong",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount8 := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"helm.sh/hook":               "pre-upgrade",
				"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded",
			},
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name: "release-name-kong",
		},
	}

	serviceAccount9 := &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "svcacct-ricplt-o1mediator",
			Namespace: "ricplt",
		},
	}

	serviceAccount10 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "release-name-prometheus-alertmanager",
			Namespace: "ricplt",
			Labels: map[string]string{
				"component": "alertmanager",
				"heritage":  "Helm",
				"release":   "release-name",
				"app":       "prometheus",
				"chart":     "prometheus-11.3.0",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	serviceAccount11 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":       "prometheus",
				"chart":     "prometheus-11.3.0",
				"component": "server",
				"heritage":  "Helm",
				"release":   "release-name",
			},
			Name:      "release-name-prometheus-server",
			Namespace: "ricplt",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
	}
 
	serviceAccount12 := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: "assigner-sa",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ServiceAccount",
		},
	}

	return []*corev1.ServiceAccount{serviceAccount1, serviceAccount2, serviceAccount3, serviceAccount4, serviceAccount5, serviceAccount6, serviceAccount7, serviceAccount8, serviceAccount9, serviceAccount10, serviceAccount11, serviceAccount12}
}