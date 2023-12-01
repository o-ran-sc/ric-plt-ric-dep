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


	return []*appsv1.Deployment{deployment1, deployment2,deployment3,deployment4,deployment5}
}
