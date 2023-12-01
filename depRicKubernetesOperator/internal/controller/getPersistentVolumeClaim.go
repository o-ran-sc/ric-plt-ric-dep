package controller

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPersistentVolumeClaim() []*corev1.PersistentVolumeClaim {

	persistentVolumeClaim1 := &corev1.PersistentVolumeClaim{
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{

				corev1.PersistentVolumeAccessMode("ReadWriteOnce"),
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					"storage": resource.MustParse("100Mi"),
				},
			},
			StorageClassName: stringPtr("local-storage"),
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "PersistentVolumeClaim",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pvc-ricplt-alarmmanager",
			Namespace: "ricplt",
		},
	}
	return []*corev1.PersistentVolumeClaim{persistentVolumeClaim1}
}
