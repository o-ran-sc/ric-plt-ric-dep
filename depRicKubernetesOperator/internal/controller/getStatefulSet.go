package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



func GetStatefulSet() []*appsv1.StatefulSet {

	statefulSet1 := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":      "ricplt-dbaas",
				"chart":    "dbaas-2.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "statefulset-ricplt-dbaas-server",
		},
		Spec: appsv1.StatefulSetSpec{
			PodManagementPolicy: appsv1.PodManagementPolicyType("OrderedReady"),
			Replicas:            int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-dbaas",
					"release": "release-name",
				},
			},
			ServiceName: "service-ricplt-dbaas-tcp",
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-dbaas",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					ShareProcessNamespace:         boolPtr(true),
					TerminationGracePeriodSeconds: int64Ptr(5),
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-dbaas-config",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Command: []string{

								"redis-server",
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-dbaas-appconfig",
										},
									},
								},
							},
							LivenessProbe: &corev1.Probe{
								PeriodSeconds: 5,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{

											"/bin/sh",
											"-c",
											"timeout 10 redis-cli -p 6379 ping",
										},
									},
								},
								InitialDelaySeconds: 15,
							},
							Stdin: false,
							TTY:   false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/data",
									Name:      "config",
									ReadOnly:  false,
								},
							},
							Args: []string{

								"/data/redis.conf",
							},
							Image: "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-dbaas:0.6.1",
							Name:  "container-ricplt-dbaas-redis",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 6379,
									Name:          "redis",
								},
							},
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 15,
								PeriodSeconds:       5,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{

											"/bin/sh",
											"-c",
											"timeout 10 redis-cli -p 6379 ping",
										},
									},
								},
							},
							StdinOnce: false,
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
				},
			},
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.StatefulSetUpdateStrategyType("RollingUpdate"),
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "StatefulSet",
		},
	}

	statefulSet2 := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "redis-cluster",
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: int32Ptr(9),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/instance": "release-name",
					"app.kubernetes.io/name":     "redis-cluster",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app.kubernetes.io/instance": "release-name",
						"app.kubernetes.io/name":     "redis-cluster",
					},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "conf",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									DefaultMode: int32Ptr(493),
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "redis-cluster-cm",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
							Command: []string{

								"/conf/update-node.sh",
								"redis-server",
								"/conf/redis.conf",
							},
							Name: "redis",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 6379,
									Name:          "client",
								},
								corev1.ContainerPort{
									Name:          "gossip",
									ContainerPort: 16379,
								},
							},
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									Name:      "conf",
									ReadOnly:  false,
									MountPath: "/conf",
								},
								corev1.VolumeMount{
									MountPath: "/data",
									Name:      "data",
									ReadOnly:  false,
								},
							},
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
							},
							Image:           "redis:5.0.1-alpine",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					TopologySpreadConstraints: []corev1.TopologySpreadConstraint{

						corev1.TopologySpreadConstraint{
							TopologyKey:       "kubernetes.io/hostname",
							WhenUnsatisfiable: corev1.UnsatisfiableConstraintAction("DoNotSchedule"),
							LabelSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"app.kubernetes.io/instance": "release-name",
									"app.kubernetes.io/name":     "redis-cluster",
								},
							},
							MaxSkew: 1,
						},
					},
				},
			},
			VolumeClaimTemplates: []corev1.PersistentVolumeClaim{

				corev1.PersistentVolumeClaim{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app.kubernetes.io/name":     "redis-cluster",
							"app.kubernetes.io/instance": "release-name",
						},
						Name: "data",
					},
					Spec: corev1.PersistentVolumeClaimSpec{
						AccessModes: []corev1.PersistentVolumeAccessMode{

							corev1.PersistentVolumeAccessMode("ReadWriteOnce"),
						},
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								"storage": resource.MustParse("1Gi"),
							},
						},
					},
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "StatefulSet",
		},
	}

	return []*appsv1.StatefulSet{statefulSet1, statefulSet2}
}