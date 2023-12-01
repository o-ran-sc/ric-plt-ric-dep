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

	deployment3 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":      "ricplt-appmgr",
				"chart":    "appmgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name: "deployment-ricplt-appmgr",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-appmgr",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{

						corev1.Container{
							Stdin: false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/opt/ric/config/appmgr.yaml",
									Name:      "config-volume",
									ReadOnly:  false,
									SubPath:   "appmgr.yaml",
								},
								corev1.VolumeMount{
									MountPath: "/opt/ric/secret",
									Name:      "helm-secret-volume",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									ReadOnly:  false,
									SubPath:   "helm_repo_username",
									MountPath: "/opt/ric/secret/helm_repo_username",
									Name:      "secret-volume",
								},
								corev1.VolumeMount{
									SubPath:   "helm_repo_password",
									MountPath: "/opt/ric/secret/helm_repo_password",
									Name:      "secret-volume",
									ReadOnly:  false,
								},
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-appmgr-env",
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
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "container-ricplt-appmgr",
							TTY:             false,
							Image:           "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-appmgr:0.2.0",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8080,
									Name:          "http",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 4561,
									Name:          "rmrroute",
								},
								corev1.ContainerPort{
									ContainerPort: 4560,
									Name:          "rmrdata",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							StdinOnce: false,
						},
					},
					HostIPC:            false,
					HostPID:            false,
					Hostname:           "appmgr",
					RestartPolicy:      corev1.RestartPolicy("Always"),
					ServiceAccountName: "svcacct-ricplt-appmgr",
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-appmgr-appconfig",
									},
								},
							},
						},
						corev1.Volume{
							Name: "secret-volume",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: "secret-ricplt-appmgr",
								},
							},
						},
						corev1.Volume{
							Name: "helm-secret-volume",
						},
						corev1.Volume{
							Name: "appmgr-bin-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									DefaultMode: int32Ptr(493),
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-appmgr-bin",
									},
								},
							},
						},
					},
					HostNetwork: false,
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					InitContainers: []corev1.Container{

						corev1.Container{
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "SVCACCT_NAME",
									Value: "svcacct-ricplt-appmgr",
								},
								corev1.EnvVar{
									Name:  "CLUSTER_NAME",
									Value: "kubernetes",
								},
								corev1.EnvVar{
									Name:  "KUBECONFIG",
									Value: "/tmp/kubeconfig",
								},
								corev1.EnvVar{
									Name:  "K8S_API_HOST",
									Value: "https://kubernetes.default.svc.cluster.local/",
								},
								corev1.EnvVar{
									Name:  "SECRET_NAMESPACE",
									Value: "ricinfra",
								},
								corev1.EnvVar{
									Name:  "SECRET_NAME",
									Value: "ricxapp-helm-secret",
								},
							},
							Image:     "nexus3.o-ran-sc.org:10002/o-ran-sc/it-dep-init:0.0.1",
							Name:      "container-ricplt-appmgr-copy-tiller-secret",
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									Name:      "helm-secret-volume",
									ReadOnly:  false,
									MountPath: "/opt/ric/secret",
								},
								corev1.VolumeMount{
									MountPath: "/svcacct-to-kubeconfig.sh",
									Name:      "appmgr-bin-volume",
									ReadOnly:  false,
									SubPath:   "svcacct-to-kubeconfig.sh",
								},
								corev1.VolumeMount{
									ReadOnly:  false,
									SubPath:   "appmgr-tiller-secret-copier.sh",
									MountPath: "/appmgr-tiller-secret-copier.sh",
									Name:      "appmgr-bin-volume",
								},
							},
							Command: []string{

								"/appmgr-tiller-secret-copier.sh",
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-appmgr-env",
										},
									},
								},
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
						},
					},
				},
			},
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"release": "release-name",
					"app":     "ricplt-appmgr",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment4 := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-e2mgr",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-e2mgr",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "e2mgr",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "local-router-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-e2mgr-router-configmap",
									},
								},
							},
						},
						corev1.Volume{
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-e2mgr-configuration-configmap",
									},
								},
							},
							Name: "local-configuration-file",
						},
						corev1.Volume{
							Name: "e2mgr-loglevel-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									Items: []corev1.KeyToPath{

										corev1.KeyToPath{
											Key:  "logcfg",
											Mode: int32Ptr(420),
											Path: "log-level.yaml",
										},
									},
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-e2mgr-loglevel-configmap",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 3800,
									Name:          "http",
								},
								corev1.ContainerPort{
									ContainerPort: 4561,
									Name:          "rmrroute",
								},
								corev1.ContainerPort{
									ContainerPort: 3801,
									Name:          "rmrdata",
								},
							},
							SecurityContext: &corev1.SecurityContext{
								Privileged: boolPtr(false),
							},
							StdinOnce: false,
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: 3,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "v1/health",
										Port: intstr.IntOrString{
											IntVal: 3800,
										},
									},
								},
							},
							Name:            "container-ricplt-e2mgr",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 3,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "v1/health",
										Port: intstr.IntOrString{
											IntVal: 3800,
										},
									},
								},
							},
							Stdin: true,
							TTY:   true,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/opt/E2Manager/router.txt",
									Name:      "local-router-file",
									ReadOnly:  false,
									SubPath:   "router.txt",
								},
								corev1.VolumeMount{
									MountPath: "/etc/config",
									Name:      "e2mgr-loglevel-volume",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									MountPath: "/opt/E2Manager/resources/configuration.yaml",
									Name:      "local-configuration-file",
									ReadOnly:  false,
									SubPath:   "configuration.yaml",
								},
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-e2mgr-env",
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
							Image: "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-e2mgr:3.0.1",
						},
					},
					HostIPC: false,
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-e2mgr",
				"chart":    "e2mgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "deployment-ricplt-e2mgr",
			Namespace: "ricplt",
		},
	}

	deployment5 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"release":  "release-name",
				"app":      "ricplt-e2term-alpha",
				"chart":    "e2term-3.0.0",
				"heritage": "Helm",
			},
			Name:      "deployment-ricplt-e2term-alpha",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"release": "release-name",
					"app":     "ricplt-e2term-alpha",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-e2term-alpha",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{

						corev1.Container{
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 120,
								PeriodSeconds:       60,
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{

											"/bin/sh",
											"-c",
											"ip=`hostname -i`;export RMR_SRC_ID=$ip;/opt/e2/rmr_probe -h $ip:38000",
										},
									},
								},
							},
							SecurityContext: &corev1.SecurityContext{
								Privileged: boolPtr(false),
							},
							Stdin: true,
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "SYSTEM_NAME",
									Value: "SEP",
								},
								corev1.EnvVar{
									Name:  "CONFIG_MAP_NAME",
									Value: "/etc/config/log-level",
								},
								corev1.EnvVar{
									Name: "HOST_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "spec.nodeName",
										},
									},
								},
								corev1.EnvVar{
									Name:  "SERVICE_NAME",
									Value: "RIC_E2_TERM",
								},
								corev1.EnvVar{
									Name:  "CONTAINER_NAME",
									Value: "container-ricplt-e2term",
								},
								corev1.EnvVar{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-e2term-env-alpha",
										},
									},
								},
							},
							Image: "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-e2:3.0.1",
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									Exec: &corev1.ExecAction{
										Command: []string{

											"/bin/sh",
											"-c",
											"ip=`hostname -i`;export RMR_SRC_ID=$ip;/opt/e2/rmr_probe -h $ip:38000",
										},
									},
								},
								InitialDelaySeconds: 10,
								PeriodSeconds:       10,
							},
							Name:            "container-ricplt-e2term",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 4561,
									Name:          "rmrroute-alpha",
								},
								corev1.ContainerPort{
									ContainerPort: 38000,
									Name:          "rmrdata-alpha",
								},
								corev1.ContainerPort{
									ContainerPort: 36422,
									Name:          "sctp-alpha",
									Protocol:      corev1.Protocol("SCTP"),
								},
								corev1.ContainerPort{
									ContainerPort: 8088,
									Name:          "prmts-alpha",
								},
							},
							StdinOnce: false,
							TTY:       true,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/opt/e2/router.txt",
									Name:      "local-router-file",
									ReadOnly:  false,
									SubPath:   "router.txt",
								},
								corev1.VolumeMount{
									MountPath: "/tmp/rmr_verbose",
									Name:      "local-router-file",
									ReadOnly:  false,
									SubPath:   "rmr_verbose",
								},
								corev1.VolumeMount{
									MountPath: "/etc/config",
									Name:      "local-loglevel-file",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									ReadOnly:  false,
									MountPath: "/data/outgoing/",
									Name:      "vol-shared",
								},
							},
						},
					},
					DNSPolicy:   corev1.DNSPolicy("ClusterFirstWithHostNet"),
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "e2term-alpha",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "local-router-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-e2term-router-configmap",
									},
								},
							},
						},
						corev1.Volume{
							Name: "local-loglevel-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-e2term-loglevel-configmap",
									},
								},
							},
						},
						corev1.Volume{
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: "pvc-ricplt-e2term-alpha",
									ReadOnly:  false,
								},
							},
							Name: "vol-shared",
						},
					},
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}


	deployment6 := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/component":  "app",
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name: "release-name-kong",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "app",
					"app.kubernetes.io/instance":  "release-name",
					"app.kubernetes.io/name":      "kong",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app.kubernetes.io/managed-by": "Helm",
						"app.kubernetes.io/name":       "kong",
						"app.kubernetes.io/version":    "1.4",
						"helm.sh/chart":                "kong-0.36.6",
						"app.kubernetes.io/component":  "app",
						"app.kubernetes.io/instance":   "release-name",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "release-name-kong",
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "release-name-kong-prefix-dir",
						},
						corev1.Volume{
							Name: "release-name-kong-tmp",
						},
						corev1.Volume{
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "release-name-kong-default-custom-server-blocks",
									},
								},
							},
							Name: "custom-nginx-template-volume",
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							TTY: false,
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.name",
										},
									},
								},
								corev1.EnvVar{
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.namespace",
										},
									},
									Name: "POD_NAMESPACE",
								},
							},
							Image: "kong/kubernetes-ingress-controller:0.7.0",
							LivenessProbe: &corev1.Probe{
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{
											IntVal: 10254,
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
								SuccessThreshold: 1,
								TimeoutSeconds:   5,
							},
							ReadinessProbe: &corev1.Probe{
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Port: intstr.IntOrString{
											IntVal: 10254,
										},
										Scheme: corev1.URIScheme("HTTP"),
										Path:   "/healthz",
									},
								},
								SuccessThreshold: 1,
								TimeoutSeconds:   5,
							},
							Stdin:     false,
							StdinOnce: false,
							Args: []string{

								"/kong-ingress-controller",
								"--publish-service=ricplt/release-name-kong-proxy",
								"--ingress-class=kong",
								"--election-id=kong-ingress-controller-leader-kong",
								"--kong-url=https://localhost:8444",
								"--admin-tls-skip-verify",
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "ingress-controller",
						},
						corev1.Container{
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Lifecycle: &corev1.Lifecycle{
								PreStop: &corev1.LifecycleHandler{
									Exec: &corev1.ExecAction{
										Command: []string{

											"/bin/sh",
											"-c",
											"kong quit",
										},
									},
								},
							},
							Name:      "proxy",
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/kong_prefix/",
									Name:      "release-name-kong-prefix-dir",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									ReadOnly:  false,
									MountPath: "/tmp",
									Name:      "release-name-kong-tmp",
								},
								corev1.VolumeMount{
									MountPath: "/kong",
									Name:      "custom-nginx-template-volume",
									ReadOnly:  false,
								},
							},
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "KONG_LUA_PACKAGE_PATH",
									Value: "/opt/?.lua;;",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_LISTEN",
									Value: "0.0.0.0:8444 ssl",
								},
								corev1.EnvVar{
									Name:  "KONG_PROXY_LISTEN",
									Value: "0.0.0.0:32080,0.0.0.0:32443 ssl",
								},
								corev1.EnvVar{
									Name:  "KONG_NGINX_DAEMON",
									Value: "off",
								},
								corev1.EnvVar{
									Name:  "KONG_NGINX_HTTP_INCLUDE",
									Value: "/kong/servers.conf",
								},
								corev1.EnvVar{
									Name:  "KONG_PLUGINS",
									Value: "bundled",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_ERROR_LOG",
									Value: "/dev/stderr",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_GUI_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_GUI_ERROR_LOG",
									Value: "/dev/stderr",
								},
								corev1.EnvVar{
									Name:  "KONG_DATABASE",
									Value: "off",
								},
								corev1.EnvVar{
									Name:  "KONG_NGINX_WORKER_PROCESSES",
									Value: "1",
								},
								corev1.EnvVar{
									Name:  "KONG_PORTAL_API_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_PORTAL_API_ERROR_LOG",
									Value: "/dev/stderr",
								},
								corev1.EnvVar{
									Name:  "KONG_PREFIX",
									Value: "/kong_prefix/",
								},
								corev1.EnvVar{
									Name:  "KONG_PROXY_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_PROXY_ERROR_LOG",
									Value: "/dev/stderr",
								},
							},
							Image: "kong:1.4",
							LivenessProbe: &corev1.Probe{
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/status",
										Port: intstr.IntOrString{
											StrVal: "metrics",
											Type:   intstr.Type(1),
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
								SuccessThreshold: 1,
								TimeoutSeconds:   5,
							},
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8444,
									Name:          "admin",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 32080,
									Name:          "proxy",
								},
								corev1.ContainerPort{
									ContainerPort: 32443,
									Name:          "proxy-tls",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 9542,
									Name:          "metrics",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							ReadinessProbe: &corev1.Probe{
								TimeoutSeconds:      5,
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/status",
										Port: intstr.IntOrString{
											StrVal: "metrics",
											Type:   intstr.Type(1),
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
								SuccessThreshold: 1,
							},
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: int64Ptr(1000),
					},
				},
			},
		},
	}

	deployment7 := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":  "helm",
				"name": "tiller",
			},
			Name:      "deployment-tiller-ricxapp",
			Namespace: "ricinfra",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":  "helm",
					"name": "tiller",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":  "helm",
						"name": "tiller",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "svcacct-tiller-ricxapp",
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "tiller-certs",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: "ricxapp-tiller-secret",
								},
							},
						},
					},
					AutomountServiceAccountToken: boolPtr(true),
					Containers: []corev1.Container{

						corev1.Container{
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "TILLER_NAMESPACE",
									Value: "ricinfra",
								},
								corev1.EnvVar{
									Name:  "TILLER_HISTORY_MAX",
									Value: "0",
								},
								corev1.EnvVar{
									Name:  "TILLER_TLS_VERIFY",
									Value: "1",
								},
								corev1.EnvVar{
									Name:  "TILLER_TLS_ENABLE",
									Value: "1",
								},
								corev1.EnvVar{
									Name:  "TILLER_TLS_CERTS",
									Value: "/etc/certs",
								},
							},
							Image: "ghcr.io/helm/tiller:v2.16.12",
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: 1,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/liveness",
										Port: intstr.IntOrString{
											IntVal: 44135,
										},
									},
								},
								TimeoutSeconds: 1,
							},
							Name: "tiller",
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 1,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/readiness",
										Port: intstr.IntOrString{
											IntVal: 44135,
										},
									},
								},
								TimeoutSeconds: 1,
							},
							Stdin: false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									ReadOnly:  true,
									MountPath: "/etc/certs",
									Name:      "tiller-certs",
								},
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 44134,
									Name:          "tiller",
								},
								corev1.ContainerPort{
									Name:          "http",
									ContainerPort: 44135,
								},
							},
							StdinOnce: false,
							TTY:       false,
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-ghcr-io",
						},
					},
				},
			},
		},
	}

	deployment8 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-jaegeradapter",
				"chart":    "jaegeradapter-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "deployment-ricplt-jaegeradapter",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-jaegeradapter",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-docker-io",
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-jaegeradapter",
										},
									},
								},
							},
							Image: "docker.io/jaegertracing/all-in-one:1.12",
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/",
										Port: intstr.IntOrString{
											IntVal: 16686,
										},
									},
								},
							},
							TTY:             false,
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "container-ricplt-jaegeradapter",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 5775,
									Name:          "zipkincompact",
									Protocol:      corev1.Protocol("UDP"),
								},
								corev1.ContainerPort{
									Name:          "jaegercompact",
									Protocol:      corev1.Protocol("UDP"),
									ContainerPort: 6831,
								},
								corev1.ContainerPort{
									ContainerPort: 6832,
									Name:          "jaegerbinary",
									Protocol:      corev1.Protocol("UDP"),
								},
								corev1.ContainerPort{
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 16686,
									Name:          "httpquery",
								},
								corev1.ContainerPort{
									ContainerPort: 5778,
									Name:          "httpconfig",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 9411,
									Name:          "zipkinhttp",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 14268,
									Name:          "jaegerhttp",
								},
								corev1.ContainerPort{
									ContainerPort: 14267,
									Name:          "jaegerhttpt",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 5,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/",
										Port: intstr.IntOrString{
											IntVal: 16686,
										},
									},
								},
							},
							Stdin:     false,
							StdinOnce: false,
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "jaegeradapter",
				},
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-jaegeradapter",
						"release": "release-name",
					},
				},
			},
			Paused: false,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment9 := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/component":  "app",
				"app.kubernetes.io/instance":   "release-name",
				"app.kubernetes.io/managed-by": "Helm",
				"app.kubernetes.io/name":       "kong",
				"app.kubernetes.io/version":    "1.4",
				"helm.sh/chart":                "kong-0.36.6",
			},
			Name: "release-name-kong",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "app",
					"app.kubernetes.io/instance":  "release-name",
					"app.kubernetes.io/name":      "kong",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app.kubernetes.io/version":    "1.4",
						"helm.sh/chart":                "kong-0.36.6",
						"app.kubernetes.io/component":  "app",
						"app.kubernetes.io/instance":   "release-name",
						"app.kubernetes.io/managed-by": "Helm",
						"app.kubernetes.io/name":       "kong",
					},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "release-name-kong-prefix-dir",
						},
						corev1.Volume{
							Name: "release-name-kong-tmp",
						},
						corev1.Volume{
							Name: "custom-nginx-template-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "release-name-kong-default-custom-server-blocks",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							TTY: false,
							Args: []string{

								"/kong-ingress-controller",
								"--publish-service=ricplt/release-name-kong-proxy",
								"--ingress-class=kong",
								"--election-id=kong-ingress-controller-leader-kong",
								"--kong-url=https://localhost:8444",
								"--admin-tls-skip-verify",
							},
							Image:           "kong/kubernetes-ingress-controller:0.7.0",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "ingress-controller",
							ReadinessProbe: &corev1.Probe{
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{
											IntVal: 10254,
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
								SuccessThreshold: 1,
								TimeoutSeconds:   5,
							},
							StdinOnce: false,
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath:  "metadata.name",
											APIVersion: "v1",
										},
									},
								},
								corev1.EnvVar{
									Name: "POD_NAMESPACE",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.namespace",
										},
									},
								},
							},
							LivenessProbe: &corev1.Probe{
								SuccessThreshold:    1,
								TimeoutSeconds:      5,
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/healthz",
										Port: intstr.IntOrString{
											IntVal: 10254,
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
							},
							Stdin: false,
						},
						corev1.Container{
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Lifecycle: &corev1.Lifecycle{
								PreStop: &corev1.LifecycleHandler{
									Exec: &corev1.ExecAction{
										Command: []string{

											"/bin/sh",
											"-c",
											"kong quit",
										},
									},
								},
							},
							Stdin:     false,
							StdinOnce: false,
							Image:     "kong:1.4",
							LivenessProbe: &corev1.Probe{
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/status",
										Port: intstr.IntOrString{
											StrVal: "metrics",
											Type:   intstr.Type(1),
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
								SuccessThreshold: 1,
								TimeoutSeconds:   5,
							},
							Name: "proxy",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8444,
									Name:          "admin",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 32080,
									Name:          "proxy",
								},
								corev1.ContainerPort{
									ContainerPort: 32443,
									Name:          "proxy-tls",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 9542,
									Name:          "metrics",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							ReadinessProbe: &corev1.Probe{
								SuccessThreshold:    1,
								TimeoutSeconds:      5,
								FailureThreshold:    3,
								InitialDelaySeconds: 5,
								PeriodSeconds:       10,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/status",
										Port: intstr.IntOrString{
											StrVal: "metrics",
											Type:   intstr.Type(1),
										},
										Scheme: corev1.URIScheme("HTTP"),
									},
								},
							},
							TTY: false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/kong_prefix/",
									Name:      "release-name-kong-prefix-dir",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									MountPath: "/tmp",
									Name:      "release-name-kong-tmp",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									MountPath: "/kong",
									Name:      "custom-nginx-template-volume",
									ReadOnly:  false,
								},
							},
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "KONG_LUA_PACKAGE_PATH",
									Value: "/opt/?.lua;;",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_LISTEN",
									Value: "0.0.0.0:8444 ssl",
								},
								corev1.EnvVar{
									Name:  "KONG_PROXY_LISTEN",
									Value: "0.0.0.0:32080,0.0.0.0:32443 ssl",
								},
								corev1.EnvVar{
									Name:  "KONG_NGINX_DAEMON",
									Value: "off",
								},
								corev1.EnvVar{
									Value: "/kong/servers.conf",
									Name:  "KONG_NGINX_HTTP_INCLUDE",
								},
								corev1.EnvVar{
									Name:  "KONG_PLUGINS",
									Value: "bundled",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_ERROR_LOG",
									Value: "/dev/stderr",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_GUI_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_ADMIN_GUI_ERROR_LOG",
									Value: "/dev/stderr",
								},
								corev1.EnvVar{
									Name:  "KONG_DATABASE",
									Value: "off",
								},
								corev1.EnvVar{
									Name:  "KONG_NGINX_WORKER_PROCESSES",
									Value: "1",
								},
								corev1.EnvVar{
									Name:  "KONG_PORTAL_API_ACCESS_LOG",
									Value: "/dev/stdout",
								},
								corev1.EnvVar{
									Name:  "KONG_PORTAL_API_ERROR_LOG",
									Value: "/dev/stderr",
								},
								corev1.EnvVar{
									Name:  "KONG_PREFIX",
									Value: "/kong_prefix/",
								},
								corev1.EnvVar{
									Value: "/dev/stdout",
									Name:  "KONG_PROXY_ACCESS_LOG",
								},
								corev1.EnvVar{
									Name:  "KONG_PROXY_ERROR_LOG",
									Value: "/dev/stderr",
								},
							},
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: int64Ptr(1000),
					},
					ServiceAccountName: "release-name-kong",
				},
			},
		},
	}

	deployment10 := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"chart":    "o1mediator-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-o1mediator",
			},
			Name:      "deployment-ricplt-o1mediator",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-o1mediator",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-o1mediator",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "appconfig-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-o1mediator-appconfig-configmap",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 9001,
									Name:          "http-supervise",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 8080,
									Name:          "http-mediation",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									ContainerPort: 3000,
									Name:          "http-event",
									Protocol:      corev1.Protocol("TCP"),
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
											Name: "configmap-ricplt-o1mediator-env",
										},
									},
								},
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Stdin:           false,
							StdinOnce:       false,
							TTY:             false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									Name:      "appconfig-file",
									ReadOnly:  false,
									MountPath: "/etc/o1agent",
								},
							},
							Image: "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-o1:0.3.1",
							Name:  "container-ricplt-o1mediator",
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "o1mediator",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					ServiceAccountName: "svcacct-ricplt-o1mediator",
				},
			},
		},
	}

	deployment11 := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":       "prometheus",
					"component": "alertmanager",
					"release":   "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "release-name-prometheus-alertmanager",
									},
								},
							},
						},
						corev1.Volume{
							Name: "storage-volume",
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Args: []string{

								"--config.file=/etc/config/alertmanager.yml",
								"--storage.path=/data",
								"--cluster.advertise-address=$(POD_IP):6783",
								"--web.external-url=http://localhost:9093",
							},
							Image: "prom/alertmanager:v0.20.0",
							Name:  "prometheus-alertmanager",
							Stdin: false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/etc/config",
									Name:      "config-volume",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									MountPath: "/data",
									Name:      "storage-volume",
									ReadOnly:  false,
								},
							},
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "status.podIP",
										},
									},
								},
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 9093,
								},
							},
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 30,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/-/ready",
										Port: intstr.IntOrString{
											IntVal: 9093,
										},
									},
								},
								TimeoutSeconds: 30,
							},
							StdinOnce: false,
							TTY:       false,
						},
						corev1.Container{
							Name:      "prometheus-alertmanager-configmap-reload",
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									Name:      "config-volume",
									ReadOnly:  true,
									MountPath: "/etc/config",
								},
							},
							Args: []string{

								"--volume-dir=/etc/config",
								"--webhook-url=http://127.0.0.1:9093/-/reload",
							},
							Image:           "jimmidyson/configmap-reload:v0.3.0",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					SecurityContext: &corev1.PodSecurityContext{
						FSGroup:      int64Ptr(65534),
						RunAsGroup:   int64Ptr(65534),
						RunAsNonRoot: boolPtr(true),
						RunAsUser:    int64Ptr(65534),
					},
					ServiceAccountName: "release-name-prometheus-alertmanager",
				},
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":       "prometheus",
						"chart":     "prometheus-11.3.0",
						"component": "alertmanager",
						"heritage":  "Helm",
						"release":   "release-name",
					},
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
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
	}

	deployment12 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "release-name-prometheus-server",
			Namespace: "ricplt",
			Labels: map[string]string{
				"app":       "prometheus",
				"chart":     "prometheus-11.3.0",
				"component": "server",
				"heritage":  "Helm",
				"release":   "release-name",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":       "prometheus",
					"component": "server",
					"release":   "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"release":   "release-name",
						"app":       "prometheus",
						"chart":     "prometheus-11.3.0",
						"component": "server",
						"heritage":  "Helm",
					},
				},
				Spec: corev1.PodSpec{
					HostNetwork: false,
					HostPID:     false,
					SecurityContext: &corev1.PodSecurityContext{
						FSGroup:      int64Ptr(65534),
						RunAsGroup:   int64Ptr(65534),
						RunAsNonRoot: boolPtr(true),
						RunAsUser:    int64Ptr(65534),
					},
					ServiceAccountName:            "release-name-prometheus-server",
					TerminationGracePeriodSeconds: int64Ptr(300),
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "release-name-prometheus-server",
									},
								},
							},
						},
						corev1.Volume{
							Name: "storage-volume",
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							StdinOnce: false,
							TTY:       false,
							Image:     "prom/prometheus:v2.18.1",
							LivenessProbe: &corev1.Probe{
								SuccessThreshold:    1,
								TimeoutSeconds:      30,
								FailureThreshold:    3,
								InitialDelaySeconds: 30,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/-/healthy",
										Port: intstr.IntOrString{
											IntVal: 9090,
										},
									},
								},
							},
							Stdin: false,
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 9090,
								},
							},
							ReadinessProbe: &corev1.Probe{
								FailureThreshold:    3,
								InitialDelaySeconds: 30,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/-/ready",
										Port: intstr.IntOrString{
											IntVal: 9090,
										},
									},
								},
								SuccessThreshold: 1,
								TimeoutSeconds:   30,
							},
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/etc/config",
									Name:      "config-volume",
									ReadOnly:  false,
								},
								corev1.VolumeMount{
									MountPath: "/data",
									Name:      "storage-volume",
									ReadOnly:  false,
								},
							},
							Args: []string{

								"--storage.tsdb.retention.time=15d",
								"--config.file=/etc/config/prometheus.yml",
								"--storage.tsdb.path=/data",
								"--web.console.libraries=/etc/prometheus/console_libraries",
								"--web.console.templates=/etc/prometheus/consoles",
								"--web.enable-lifecycle",
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "prometheus-server",
						},
					},
					HostIPC: false,
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment13 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"chart":   "redis-cluster-0.1.0",
				"release": "release-name",
			},
			Name:      "assigner-dep",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "assigner",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{

						corev1.Container{
							Image:     "bitnami/kubectl:1.18",
							Name:      "kubectl",
							StdinOnce: false,
							TTY:       false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/conf",
									Name:      "conf",
									ReadOnly:  false,
								},
							},
							Args: []string{

								"-c",
								"sleep 3000",
							},
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "POD_LABEL",
									Value: "app.kubernetes.io/instance=release-name",
								},
							},
							Stdin: false,
							Command: []string{

								"/bin/sh",
							},
						},
					},
					HostIPC:            false,
					HostNetwork:        false,
					HostPID:            false,
					ServiceAccountName: "assigner-sa",
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "conf",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									DefaultMode: int32Ptr(493),
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "assigner-cm",
									},
								},
							},
						},
					},
				},
			},
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "assigner",
					"release": "release-name",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment14 := &appsv1.Deployment{
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-rsm",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-rsm",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Hostname: "rsm",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					InitContainers: []corev1.Container{

						corev1.Container{
							Command: []string{

								"sh",
								"-c",
								"redis-cli -h service-ricplt-dbaas-tcp -p 6379 set \"{rsm},CFG:GENERAL:v1.0.0\"  \"{\\\"enableResourceStatus\\\":true,\\\"partialSuccessAllowed\\\":true,\\\"prbPeriodic\\\":true,\\\"tnlLoadIndPeriodic\\\":true,\\\"wwLoadIndPeriodic\\\":true,\\\"absStatusPeriodic\\\":true,\\\"rsrpMeasurementPeriodic\\\":true,\\\"csiPeriodic\\\":true,\\\"periodicityMs\\\":1,\\\"periodicityRsrpMeasurementMs\\\":3,\\\"periodicityCsiMs\\\":4}\" nx",
							},
							Image:           "docker.io/redis:latest",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "redis-init",
							Stdin:           false,
							StdinOnce:       false,
							TTY:             false,
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "local-router-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-rsm-router-configmap",
									},
								},
							},
						},
						corev1.Volume{
							Name: "local-configuration-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-rsm",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Name: "container-ricplt-rsm",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 4800,
									Name:          "http",
								},
								corev1.ContainerPort{
									ContainerPort: 4561,
									Name:          "rmrroute",
								},
								corev1.ContainerPort{
									ContainerPort: 4801,
									Name:          "rmrdata",
								},
							},
							Stdin: true,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath: "/opt/RSM/router.txt",
									Name:      "local-router-file",
									ReadOnly:  false,
									SubPath:   "router.txt",
								},
								corev1.VolumeMount{
									MountPath: "/opt/RSM/resources/configuration.yaml",
									Name:      "local-configuration-file",
									ReadOnly:  false,
									SubPath:   "configuration.yaml",
								},
							},
							Image:           "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-resource-status-manager:3.0.1",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							StdinOnce:       false,
							TTY:             true,
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-rsm-env",
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
							SecurityContext: &corev1.SecurityContext{
								Privileged: boolPtr(false),
							},
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-rsm",
				"chart":    "rsm-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "deployment-ricplt-rsm",
			Namespace: "ricplt",
		},
	}

	deployment15 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app":      "ricplt-rtmgr",
				"chart":    "rtmgr-3.0.0",
				"heritage": "Helm",
				"release":  "release-name",
			},
			Name:      "deployment-ricplt-rtmgr",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Paused: false,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-rtmgr",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"release": "release-name",
						"app":     "ricplt-rtmgr",
					},
				},
				Spec: corev1.PodSpec{
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "rtmgrcfg",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "configmap-ricplt-rtmgr-rtmgrcfg",
									},
									Items: []corev1.KeyToPath{

										corev1.KeyToPath{
											Key:  "rtmgrcfg",
											Mode: int32Ptr(420),
											Path: "rtmgr-config.yaml",
										},
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 3800,
									Name:          "http",
								},
								corev1.ContainerPort{
									ContainerPort: 4561,
									Name:          "rmrroute",
								},
								corev1.ContainerPort{
									ContainerPort: 4560,
									Name:          "rmrdata",
								},
							},
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 5,
								PeriodSeconds:       15,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "ric/v1/health/ready",
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
									},
								},
							},
							Stdin: false,
							TTY:   false,
							Command: []string{

								"/run_rtmgr.sh",
							},
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-rtmgr-env",
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
							LivenessProbe: &corev1.Probe{
								PeriodSeconds: 15,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
										Path: "ric/v1/health/alive",
									},
								},
								InitialDelaySeconds: 5,
							},
							Name: "container-ricplt-rtmgr",
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									ReadOnly:  true,
									MountPath: "/cfg",
									Name:      "rtmgrcfg",
								},
							},
							Image:           "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-rtmgr:0.3.8",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							StdinOnce:       false,
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "rtmgr",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment16 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-submgr",
				"chart":    "submgr-3.0.0",
			},
			Name:      "deployment-ricplt-submgr",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-submgr",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Hostname: "submgr",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					Volumes: []corev1.Volume{

						corev1.Volume{
							Name: "config-volume",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									Items: []corev1.KeyToPath{

										corev1.KeyToPath{
											Key:  "submgrcfg",
											Mode: int32Ptr(420),
											Path: "submgr-config.yaml",
										},
										corev1.KeyToPath{
											Key:  "submgrutartg",
											Mode: int32Ptr(420),
											Path: "submgr-uta-rtg.rt",
										},
									},
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "submgrcfg",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							TTY: false,
							VolumeMounts: []corev1.VolumeMount{

								corev1.VolumeMount{
									ReadOnly:  false,
									MountPath: "/cfg",
									Name:      "config-volume",
								},
							},
							Args: []string{

								"-f",
								"/cfg/submgr-config.yaml",
							},
							Command: []string{

								"/submgr",
							},
							Image: "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-submgr:0.10.7",
							Name:  "container-ricplt-submgr",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 3800,
									Name:          "http",
									Protocol:      corev1.Protocol("TCP"),
								},
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
							},
							Stdin: false,
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-submgr-env",
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
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							LivenessProbe: &corev1.Probe{
								PeriodSeconds: 15,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "ric/v1/health/alive",
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
									},
								},
								InitialDelaySeconds: 5,
							},
							ReadinessProbe: &corev1.Probe{
								InitialDelaySeconds: 5,
								PeriodSeconds:       15,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "ric/v1/health/ready",
										Port: intstr.IntOrString{
											IntVal: 8080,
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
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-submgr",
					"release": "release-name",
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
	}

	deployment17 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"heritage": "Helm",
				"release":  "release-name",
				"app":      "ricplt-vespamgr",
				"chart":    "vespamgr-3.0.0",
			},
			Name:      "deployment-ricplt-vespamgr",
			Namespace: "ricplt",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-vespamgr",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-vespamgr",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					Hostname: "vespamgr",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10002-o-ran-sc",
						},
					},
					Containers: []corev1.Container{

						corev1.Container{
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-vespamgr",
										},
									},
								},
								corev1.EnvFromSource{
									SecretRef: &corev1.SecretEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "vespa-secrets",
										},
									},
								},
							},
							LivenessProbe: &corev1.Probe{
								InitialDelaySeconds: 30,
								PeriodSeconds:       60,
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/supervision",
										Port: intstr.IntOrString{
											IntVal: 8080,
										},
									},
								},
								TimeoutSeconds: 20,
							},
							Name:      "container-ricplt-vespamgr",
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
							Env: []corev1.EnvVar{

								corev1.EnvVar{
									Name:  "VESMGR_APPMGRDOMAN",
									Value: "service-ricplt-appmgr-http",
								},
							},
							Image:           "nexus3.o-ran-sc.org:10002/o-ran-sc/ric-plt-vespamgr:0.4.0",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8080,
									Name:          "http",
									Protocol:      corev1.Protocol("TCP"),
								},
								corev1.ContainerPort{
									Name:          "alert",
									Protocol:      corev1.Protocol("TCP"),
									ContainerPort: 9095,
								},
							},
						},
					},
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
				},
			},
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
	}

	deployment18 := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Labels: map[string]string{
				"release":  "release-name",
				"app":      "ricplt-xapp-onboarder",
				"chart":    "xapp-onboarder-3.0.0",
				"heritage": "Helm",
			},
			Name: "deployment-ricplt-xapp-onboarder",
		},
		Spec: appsv1.DeploymentSpec{
			Paused:   false,
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app":     "ricplt-xapp-onboarder",
					"release": "release-name",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":     "ricplt-xapp-onboarder",
						"release": "release-name",
					},
				},
				Spec: corev1.PodSpec{
					HostIPC:     false,
					HostNetwork: false,
					HostPID:     false,
					Hostname:    "xapp-onboarder",
					ImagePullSecrets: []corev1.LocalObjectReference{

						corev1.LocalObjectReference{
							Name: "secret-nexus3-o-ran-sc-org-10004-o-ran-sc",
						},
					},
					RestartPolicy: corev1.RestartPolicy("Always"),
					Containers: []corev1.Container{

						corev1.Container{
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-xapp-onboarder-chartmuseum-env",
										},
									},
								},
							},
							Image:           "docker.io/chartmuseum/chartmuseum:v0.8.2",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "container-ricplt-xapp-onboarder-chartmuseum",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8080,
									Name:          "chartmuseum",
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							Stdin:     false,
							StdinOnce: false,
							TTY:       false,
						},
						corev1.Container{
							StdinOnce: false,
							TTY:       false,
							EnvFrom: []corev1.EnvFromSource{

								corev1.EnvFromSource{
									ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "configmap-ricplt-xapp-onboarder-env",
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
							Image:           "nexus3.o-ran-sc.org:10004/o-ran-sc/xapp-onboarder:1.0.0",
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
							Name:            "container-ricplt-xapp-onboarder",
							Ports: []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort: 8888,
									Name:          "server",
									Protocol:      corev1.Protocol("TCP"),
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
	}

	return []*appsv1.Deployment{deployment1, deployment2, deployment3, deployment4, deployment5, deployment6, deployment7, deployment8, deployment9, deployment10, deployment11, deployment12, deployment13, deployment14, deployment15, deployment16, deployment17, deployment18}
}
