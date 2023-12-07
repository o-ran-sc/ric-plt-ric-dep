package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetSecret() []*corev1.Secret {

	secret1 := &corev1.Secret{
		Data: map[string][]uint8{
			"helm_repo_password": getDataForSecret("helm"),
			"helm_repo_username": getDataForSecret("helm"),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "secret-ricplt-appmgr",
		},
		Type: corev1.SecretType("Opaque"),
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
	}
return []*corev1.Secret{secret1}

}