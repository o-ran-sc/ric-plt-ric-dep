package controller

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)


func GetIngress() []*unstructured.Unstructured {

	ingress1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1beta1",
			"kind":       "Ingress",
			"metadata": map[string]interface{}{
				"name": "ingress-ricplt-a1mediator",
			},
			"spec": map[string]interface{}{
				"rules": []interface{}{
					map[string]interface{}{
						"http": map[string]interface{}{
							"paths": []interface{}{
								map[string]interface{}{
									"backend": map[string]interface{}{
										"serviceName": "service-ricplt-a1mediator-http",
										"servicePort": 10000,
									},
									"path": "/a1mediator",
								},
							},
						},
					},
				},
			},
		},
	}
	ingress2 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1beta1",
			"kind":       "Ingress",
			"metadata": map[string]interface{}{
				"name": "ingress-ricplt-appmgr",
			},
			"spec": map[string]interface{}{
				"rules": []interface{}{
					map[string]interface{}{
						"http": map[string]interface{}{
							"paths": []interface{}{
								map[string]interface{}{
									"backend": map[string]interface{}{
										"serviceName": "service-ricplt-appmgr-http",
										"servicePort": 8080,
									},
									"path": "/appmgr",
								},
							},
						},
					},
				},
			},
		},
	}

	ingress3 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "networking.k8s.io/v1beta1",
			"kind":       "Ingress",
			"metadata": map[string]interface{}{
				"name": "ingress-ricplt-e2mgr",
			},
			"spec": map[string]interface{}{
				"rules": []interface{}{
					map[string]interface{}{
						"http": map[string]interface{}{
							"paths": []interface{}{
								map[string]interface{}{
									"backend": map[string]interface{}{
										"serviceName": "service-ricplt-e2mgr-http",
										"servicePort": 3800,
									},
									"path": "/e2mgr",
								},
							},
						},
					},
				},
			},
		},
	}

	return []*unstructured.Unstructured{ingress1, ingress2,ingress3}
}