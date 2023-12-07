package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func GetService() []*corev1.Service {

	service1 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-a1mediator",
				"chart":    "a1mediator-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-a1mediator-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
					Name: "http",
					Port: 10000,
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-a1mediator",
				"release": "release-name",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service2 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"release":  "release-name",
				"app":      "ricplt-a1mediator",
				"chart":    "a1mediator-3.0.0",
				"heritage": "Helm",
			},
			Name:      "service-ricplt-a1mediator-rmr",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-a1mediator",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "rmrroute",
					Port:     4561,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(1),
						StrVal: "rmrroute",
					},
				},
				corev1.ServicePort{
					Name:     "rmrdata",
					Port:     4562,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	return []*corev1.Service{service1, service2}
}
