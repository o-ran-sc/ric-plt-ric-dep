package controller

import (
	"context"
	"fmt"
)


func GetJob() []*unstructured.Unstructured {

	job1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "batch/v1",
			"kind":       "Job",
			"metadata": map[string]interface{}{
				"name":      "tiller-secret-generator",
				"namespace": "ricinfra",
			},
			"spec": map[string]interface{}{
				"template": map[string]interface{}{
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"env": []interface{}{
									map[string]interface{}{
										"name":  "ENTITIES",
										"value": "ricxapp-tiller-secret ricxapp-helm-secret",
									},
									map[string]interface{}{
										"name":  "TILLER_KEY_NAME",
										"value": "ricxapp-tiller-secret.key.pem",
									},
									map[string]interface{}{
										"name":  "TILLER_CERT_NAME",
										"value": "ricxapp-tiller-secret.cert.pem",
									},
									map[string]interface{}{
										"name":  "HELM_KEY_NAME",
										"value": "ricxapp-helm-secret.key.pem",
									},
									map[string]interface{}{
										"name":  "HELM_CERT_NAME",
										"value": "ricxapp-helm-secret.cert.pem",
									},
									map[string]interface{}{
										"name":  "TILLER_CN",
										"value": "service-tiller-ricxapp",
									},
									map[string]interface{}{
										"name":  "CLUSTER_SERVER",
										"value": "https://kubernetes.default.svc.cluster.local/",
									},
								},
								"image":           "nexus3.o-ran-sc.org:10002/o-ran-sc/it-dep-secret:0.0.2",
								"imagePullPolicy": "IfNotPresent",
								"name":            "tiller-secret-generator",
							},
						},
						"imagePullSecrets": []interface{}{
							map[string]interface{}{
								"name": "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
							},
						},
						"restartPolicy":      "Never",
						"serviceAccountName": "tiller-secret-creator-xzhjjg",
					},
				},
			},
		},
	}

	return []*unstructured.Unstructured{job1}
}