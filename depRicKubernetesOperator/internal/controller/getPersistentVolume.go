package controller

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

func GetPersistentVolume() []*unstructured.Unstructured {

	persistentVolume1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "PersistentVolume",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"type": "local",
				},
				"name":      "pv-ricplt-alarmmanager",
				"namespace": "ricplt",
			},
			"spec": map[string]interface{}{
				"accessModes": []interface{}{
					"ReadWriteOnce",
				},
				"capacity": map[string]interface{}{
					"storage": "100Mi",
				},
				"hostPath": map[string]interface{}{
					"path": "/mnt/pv-ricplt-alarmmanager",
				},
				"storageClassName": "local-storage",
			},
		},
	}

	return []*unstructured.Unstructured{persistentVolume1}
}
