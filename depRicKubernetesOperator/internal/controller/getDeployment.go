package controller

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func int32Ptr(val int) *int32 {
	var a int32
	a = int32(val)
	return &a
}

func GetDeployment() []*appsv1.Deployment {

	deployment1 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-a1mediator",
				"chart":    "a1mediator-3.0.0",
			},
			Name:      "deployment-ricplt-a1mediator",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-a1mediator",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"release": "release-name",
						"app":     "ricplt-a1mediator",
					},
				},
				Spec: corev1.PodSpec{
					HostPID:  false,
					Hostname: "a1mediator",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "a1conf",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-a1mediator-a1conf",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/a1-p/healthcheck",
										Port: intstr.IntOrString{
											StrVal: "http",
											Type:   intstr.Type(1),
										},
									},
								},
							},
							Name: "container-ricplt-a1mediator",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 10000,
									Name:          "http",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 4561,
									Name:          "rmrroute",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 4562,
									Name:          "rmrdata",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-a1mediator-env",
										},
									},
								},
								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-dbaas-appconfig",
										},
									},
								},
							},
							Image: "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-a1:2.5.0",
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/a1-p/healthcheck",
										Port: intstr.IntOrString{
											StrVal: "http",
											Type:   intstr.Type(1),
										},
									},
								},
							},
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/opt/route",
									Name:      "a1conf",
									ReadOnly:  false,
								},
							},
						},
					},
					HostIPC:     false,
					HostNetwork: false,
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment2 := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"release": "release-name",
					"app":     "ricplt-alarmmanager",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-alarmmanager",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "alarmmanager",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					ServiceAccountName: "svcacct-ricplt-alarmmanager",
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									Items: []corev1.KeyToPath{

										corev1.KeyToPath{
											Mode: int32Ptr(420),
											Path: "config-file.json",
											Key:  "alarmmanagercfg",
										},
									},
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-alarmmanager-alarmmanagercfg",
									},
								},
							},
						},
						corev1.Volume{
							Name: "am-persistent-storage",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "pvc-ricplt-alarmmanager",
									ReadOnly:  false,
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 4561,
									Name:          "rmrroute",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 4560,
									Name:          "rmrdata",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 8080,
									Name:          "http",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							StdinOnce: false,
							TTY:       false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/cfg",
									Name:      "config-volume",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									Name:      "am-persistent-storage",
									ReadOnly:  false,
									MountPath: "/mnt/pv-ricplt-alarmmanager",
								},
							},
							Image:           "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-alarmmanager:0.5.9",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "container-ricplt-alarmmanager",
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "PLT_NAMESPACE",
									Value: "ricplt",
								},
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-dbaas-appconfig",
										},
									},
								},
								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-alarmmanager-env",
										},
									},
								},
							},
							Stdin: false,
						},
					},
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"chart":    "alarmmanager-5.0.0",
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-alarmmanager",
			},
			Name: "deployment-ricplt-alarmmanager",
		},
	}

	return []*appsv1.Deployment{deployment1, deployment2}
}
