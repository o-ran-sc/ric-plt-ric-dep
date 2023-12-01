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
	service3 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-alarmmanager",
				"chart":    "alarmmanager-5.0.0",
			},
			Name:      "service-ricplt-alarmmanager-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     8080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 8080,
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-alarmmanager",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service4 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-alarmmanager",
				"chart":    "alarmmanager-5.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-alarmmanager-rmr",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
					Name:     "rmrdata",
					Port:     4560,
					Protocol: corev1.Protocol("TCP"),
				},
				corev1.ServicePort{
					Name:     "rmrroute",
					Port:     4561,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrroute",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-alarmmanager",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service5 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-appmgr",
				"chart":    "appmgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-appmgr-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     8080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(1),
						StrVal: "http",
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"release": "release-name",
				"app":     "ricplt-appmgr",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service6 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "service-ricplt-appmgr-rmr",
			Namespace: "ricplt",
			Labels: map[string]string{
				"chart":    "appmgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-appmgr",
			},
		},
		Spec: corev1.ServiceSpec{
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
					Port:     4560,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-appmgr",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	return []*corev1.Service{service1, service2, service3, service4,service5,service6 }
}
