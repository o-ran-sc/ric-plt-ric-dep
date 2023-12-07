package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)


func GetEndpoints() []*unstructured.Unstructured {

	endpoints1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Endpoints",
			"metadata": map[string]interface{}{
				"name":      "aux-entry",
				"namespace": "ricplt",
			},
			"subsets": []interface{}{
				map[string]interface{}{
					"ports": []interface{}{
						map[string]interface{}{
							"name": "aux-entry-http-ingress-port",
							"port": 32080,
						},
						map[string]interface{}{
							"name": "aux-entry-https-ingress-port",
							"port": 32443,
						},
					},
					"addresses": []interface{}{
						map[string]interface{}{
							"ip": "10.0.0.1",
						},
					},
				},
			},
		},
	}

	endpoints2 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Endpoints",
			"metadata": map[string]interface{}{
				"name":      "aux-entry",
				"namespace": "ricxapp",
			},
			"subsets": []interface{}{
				map[string]interface{}{
					"addresses": []interface{}{
						map[string]interface{}{
							"ip": "10.0.0.1",
						},
					},
					"ports": []interface{}{
						map[string]interface{}{
							"name": "aux-entry-http-ingress-port",
							"port": 32080,
						},
						map[string]interface{}{
							"port": 32443,
							"name": "aux-entry-https-ingress-port",
						},
					},
				},
			},
		},
	}

	endpoints3 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Endpoints",
			"metadata": map[string]interface{}{
				"namespace": "ricplt",
				"name":      "aux-entry",
			},
			"subsets": []interface{}{
				map[string]interface{}{
					"addresses": []interface{}{
						map[string]interface{}{
							"ip": "10.0.0.1",
						},
					},
					"ports": []interface{}{
						map[string]interface{}{
							"name": "aux-entry-http-ingress-port",
							"port": 32080,
						},
						map[string]interface{}{
							"name": "aux-entry-https-ingress-port",
							"port": 32443,
						},
					},
				},
			},
		},
	}

	endpoints4 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "Endpoints",
			"metadata": map[string]interface{}{
				"name":      "aux-entry",
				"namespace": "ricxapp",
			},
			"subsets": []interface{}{
				map[string]interface{}{
					"addresses": []interface{}{
						map[string]interface{}{
							"ip": "10.0.0.1",
						},
					},
					"ports": []interface{}{
						map[string]interface{}{
							"name": "aux-entry-http-ingress-port",
							"port": 32080,
						},
						map[string]interface{}{
							"name": "aux-entry-https-ingress-port",
							"port": 32443,
						},
					},
				},
			},
		},
	}

	return []*unstructured.Unstructured{endpoints1, endpoints2, endpoints3, endpoints4}
}