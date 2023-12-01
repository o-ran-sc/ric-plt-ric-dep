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

	service7 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-dbaas",
				"chart":    "dbaas-2.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "service-ricplt-dbaas-tcp",
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "server",
					Port:     6379,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "redis",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-dbaas",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service8 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-e2mgr",
				"chart":    "e2mgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-e2mgr-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
					Name:     "http",
					Port:     3800,
					Protocol: corev1.Protocol("TCP"),
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-e2mgr",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service9 := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-e2mgr",
				"chart":    "e2mgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-e2mgr-rmr",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":     "ricplt-e2mgr",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrroute",
						Type:   intstr.Type(1),
					},
					Name: "rmrroute",
					Port: 4561,
				},
				corev1.ServicePort{
					Name:     "rmrdata",
					Port:     3801,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
		},
	}

	service10 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"prometheus.io/path":   "/metrics",
				"prometheus.io/port":   "8088",
				"prometheus.io/scrape": "true",
			},
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-e2term-alpha",
				"chart":    "e2term-3.0.0",
			},
			Name:      "service-ricplt-e2term-prometheus-alpha",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "prmts-alpha",
					Port:     8088,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "prmts-alpha",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"release": "release-name",
				"app":     "ricplt-e2term-alpha",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service11 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"chart":    "e2term-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-e2term-alpha",
			},
			Name:      "service-ricplt-e2term-rmr-alpha",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "rmrroute-alpha",
					Port:     4561,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrroute-alpha",
						Type:   intstr.Type(1),
					},
				},
				corev1.ServicePort{
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata-alpha",
						Type:   intstr.Type(1),
					},
					Name:     "rmrdata-alpha",
					Port:     38000,
					Protocol: corev1.Protocol("TCP"),
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-e2term-alpha",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service12 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":      "ricplt-e2term-alpha",
				"chart":    "e2term-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "service-ricplt-e2term-sctp-alpha",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "sctp-alpha",
					NodePort: 32222,
					Port:     36422,
					Protocol: corev1.Protocol("SCTP"),
					TargetPort: intstr.IntOrString{
						IntVal: 36422,
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-e2term-alpha",
				"release": "release-name",
			},
			Type: corev1.ServiceType("NodePort"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service13 := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "aux-entry",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "aux-entry-http-ingress-port",
					Port:     80,
					Protocol: corev1.Protocol("TCP"),
				},
				corev1.ServicePort{
					Name:     "aux-entry-https-ingress-port",
					Port:     443,
					Protocol: corev1.Protocol("TCP"),
				},
			},
			PublishNotReadyAddresses: false,
		},
	}

	service14 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "aux-entry",
			Namespace: "ricxapp",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "aux-entry-http-ingress-port",
					Port:     80,
					Protocol: corev1.Protocol("TCP"),
				},
				corev1.ServicePort{
					Protocol: corev1.Protocol("TCP"),
					Name:     "aux-entry-https-ingress-port",
					Port:     443,
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service15 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "aux-entry",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "aux-entry-http-ingress-port",
					Port:     80,
					Protocol: corev1.Protocol("TCP"),
				},
				corev1.ServicePort{
					Name:     "aux-entry-https-ingress-port",
					Port:     443,
					Protocol: corev1.Protocol("TCP"),
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service16 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "aux-entry",
			Namespace: "ricxapp",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Port:     80,
					Protocol: corev1.Protocol("TCP"),
					Name:     "aux-entry-http-ingress-port",
				},
				corev1.ServicePort{
					Name:     "aux-entry-https-ingress-port",
					Port:     443,
					Protocol: corev1.Protocol("TCP"),
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service17 := &corev1.Service{
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app.kubernetes.io/instance":  "release-name",
				"app.kubernetes.io/name":      "kong",
				"app.kubernetes.io/component": "app",
			},
			Type: corev1.ServiceType("NodePort"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "kong-proxy",
					NodePort: 32080,
					Port:     32080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 32080,
					},
				},
				corev1.ServicePort{
					Port:     32443,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 32443,
					},
					Name:     "kong-proxy-tls",
					NodePort: 32443,
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "release-name-kong-proxy",
			Labels: map[string]string{
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
		},
	}

	service18 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":  "helm",
				"name": "tiller",
			},
			Name:      "service-tiller-ricxapp",
			Namespace: "ricinfra",
		},
		Spec: corev1.ServiceSpec{
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":  "helm",
				"name": "tiller",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name: "tiller",
					Port: 44134,
					TargetPort: intstr.IntOrString{
						StrVal: "tiller",
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

	service19 := &corev1.Service{
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Port:     5775,
					Protocol: corev1.Protocol("UDP"),
					TargetPort: intstr.IntOrString{
						IntVal: 5775,
					},
					Name: "zipkincompact",
				},
				corev1.ServicePort{
					Name:     "jaegercompact",
					Port:     6831,
					Protocol: corev1.Protocol("UDP"),
					TargetPort: intstr.IntOrString{
						IntVal: 6831,
					},
				},
				corev1.ServicePort{
					Name:     "jaegerbinary",
					Port:     6832,
					Protocol: corev1.Protocol("UDP"),
					TargetPort: intstr.IntOrString{
						IntVal: 6832,
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"release": "release-name",
				"app":     "ricplt-jaegeradapter",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-jaegeradapter",
				"chart":    "jaegeradapter-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-jaegeradapter-agent",
			Namespace: "ricplt",
		},
	}

	service20 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-jaegeradapter",
				"chart":    "jaegeradapter-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-jaegeradapter-collector",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "jaegerhttpt",
					Port:     14267,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 14267,
					},
				},
				corev1.ServicePort{
					Port:     14268,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 14268,
					},
					Name: "jaegerhttp",
				},
				corev1.ServicePort{
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 9411,
					},
					Name: "zipkinhttp",
					Port: 9411,
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-jaegeradapter",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service21 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-jaegeradapter",
				"chart":    "jaegeradapter-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-jaegeradapter-query",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":     "ricplt-jaegeradapter",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "httpquery",
					Port:     16686,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 16686,
					},
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service22 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
				"app.kubernetes.io/instance":   "release-name",
			},
			Name: "release-name-kong-proxy",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType("NodePort"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					NodePort: 32080,
					Port:     32080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 32080,
					},
					Name: "kong-proxy",
				},
				corev1.ServicePort{
					Port:     32443,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 32443,
					},
					Name:     "kong-proxy-tls",
					NodePort: 32443,
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app.kubernetes.io/component": "app",
				"app.kubernetes.io/instance":  "release-name",
				"app.kubernetes.io/name":      "kong",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service23 := &corev1.Service{
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					TargetPort: intstr.IntOrString{
						IntVal: 9001,
					},
					Name:     "http-supervise",
					Port:     9001,
					Protocol: corev1.Protocol("TCP"),
				},
				corev1.ServicePort{
					Port:     8080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 8080,
					},
					Name: "http-mediation",
				},
				corev1.ServicePort{
					Port:     3000,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 3000,
					},
					Name: "http-event",
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-o1mediator",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":      "ricplt-o1mediator",
				"chart":    "o1mediator-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "service-ricplt-o1mediator-http",
		},
	}

	service24 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-o1mediator",
				"chart":    "o1mediator-3.0.0",
			},
			Name: "service-ricplt-o1mediator-tcp-netconf",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "tcp-netconf",
					NodePort: 30830,
					Port:     830,
					Protocol: corev1.Protocol("TCP"),
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-o1mediator",
				"release": "release-name",
			},
			Type: corev1.ServiceType("NodePort"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service25 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":       "prometheus",
				"chart":     "prometheus-11.3.0",
				"component": "alertmanager",
				"heritage":  "Helm",
				"release":   "release-name",
			},
			Name:      "release-name-prometheus-alertmanager",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     80,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 9093,
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":       "prometheus",
				"component": "alertmanager",
				"release":   "release-name",
			},
			SessionAffinity: corev1.ServiceAffinity("None"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service26 := &corev1.Service{
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
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     80,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						IntVal: 9090,
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"release":   "release-name",
				"app":       "prometheus",
				"component": "server",
			},
			SessionAffinity: corev1.ServiceAffinity("None"),
			Type:            corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service27 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-cluster-svc",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app.kubernetes.io/instance": "release-name",
				"app.kubernetes.io/name":     "redis-cluster",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name: "client",
					Port: 6379,
					TargetPort: intstr.IntOrString{
						IntVal: 6379,
					},
				},
				corev1.ServicePort{
					Name: "gossip",
					Port: 16379,
					TargetPort: intstr.IntOrString{
						IntVal: 16379,
					},
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service28 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-rsm",
				"chart":    "rsm-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-rsm-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-rsm",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Port:     4800,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
					Name: "http",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service29 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":      "ricplt-rsm",
				"chart":    "rsm-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "service-ricplt-rsm-rmr",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "rmrroute",
					Port:     4561,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrroute",
						Type:   intstr.Type(1),
					},
				},
				corev1.ServicePort{
					Name:     "rmrdata",
					Port:     4801,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"release": "release-name",
				"app":     "ricplt-rsm",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service30 := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-rtmgr",
				"chart":    "rtmgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-rtmgr-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-rtmgr",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     3800,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
				},
			},
		},
	}

	service31 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":      "ricplt-rtmgr",
				"chart":    "rtmgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "service-ricplt-rtmgr-rmr",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":     "ricplt-rtmgr",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(1),
						StrVal: "rmrroute",
					},
					Name: "rmrroute",
					Port: 4561,
				},
				corev1.ServicePort{
					Port:     4560,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "rmrdata",
						Type:   intstr.Type(1),
					},
					Name: "rmrdata",
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service32 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-submgr",
				"chart":    "submgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-submgr-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     3800,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-submgr",
				"release": "release-name",
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	service33 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-submgr",
				"chart":    "submgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-submgr-rmr",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Port:     4560,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(1),
						StrVal: "rmrdata",
					},
					Name: "rmrdata",
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
				"app":     "ricplt-submgr",
				"release": "release-name",
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service34 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-vespamgr",
				"chart":    "vespamgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-vespamgr-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":     "ricplt-vespamgr",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Name:     "http",
					Port:     8080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "http",
						Type:   intstr.Type(1),
					},
				},
				corev1.ServicePort{
					Name:     "alert",
					Port:     9095,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "alert",
						Type:   intstr.Type(1),
					},
				},
			},
			PublishNotReadyAddresses: false,
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
	}

	service35 := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-xapp-onboarder",
				"chart":    "xapp-onboarder-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "service-ricplt-xapp-onboarder-http",
			Namespace: "ricplt",
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{

				corev1.ServicePort{
					Port:     8888,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "server",
						Type:   intstr.Type(1),
					},
					Name: "server",
				},
				corev1.ServicePort{
					Port:     8080,
					Protocol: corev1.Protocol("TCP"),
					TargetPort: intstr.IntOrString{
						StrVal: "chartmuseum",
						Type:   intstr.Type(1),
					},
					Name: "chartmuseum",
				},
			},
			PublishNotReadyAddresses: false,
			Selector: map[string]string{
				"app":     "ricplt-xapp-onboarder",
				"release": "release-name",
			},
			Type: corev1.ServiceType("ClusterIP"),
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
	}

	return []*corev1.Service{service1, service2, service3, service4, service5, service6, service7, service8, service9, service10, service11, service12, service13, service14, service15, service16, service17, service18, service19, service20, service21, service22, service23, service24, service25, service26, service27, service28, service29, service30, service31, service32, service33, service34, service35}
}
