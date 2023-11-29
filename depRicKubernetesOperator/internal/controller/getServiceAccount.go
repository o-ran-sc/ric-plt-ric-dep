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
	return []*corev1.ServiceAccount{serviceAccount1}
}