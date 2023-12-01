package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetConfigMap() []*corev1.ConfigMap {

	configMap1 := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-a1mediator-a1conf",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		Data: map[string]string{
			"local.rt": "newrt|start\n" +
				"# Warning! this is not a functioning table because the subscription manager and route manager are now involved in a1 flows\n" +
				"# the real routing table requires subscription ids as routing is now done over sub ids, but this isn't known until xapp deploy time, it's a dynamic process triggered by the xapp manager\n" +
				"# there is a single message type for all messages a1 sends out now, subid is the other necessary piece of info\n" +
				"# there are two message types a1 listens for; 20011 (instance response) and 20012 (query)\n" +
				"# xapps likely use rts to reply with 20012 so the routing entry isn't needed for that in most cases\n" +
				"mse|20010|SUBID|service-ricxapp-admctrl-rmr.ricxapp:4563\n" +
				"rte|20011|service-ricplt-a1mediator-rmr.ricplt:4562\n" +
				"rte|20012|service-ricplt-a1mediator-rmr.ricplt:4562\n" +
				"newrt|end\n" +
				"",
			"loglevel.txt": "log-level:",
		},
	}

	configMap2 := &corev1.ConfigMap{
		Data: map[string]string{
			"CONFIG_MAP_NAME":             "/opt/route/loglevel.txt",
			"INSTANCE_DELETE_NO_RESP_TTL": "5",
			"INSTANCE_DELETE_RESP_TTL":    "10",
			"PYTHONUNBUFFERED":            "1",
			"RMR_RTG_SVC":                 "4561",
			"RMR_SRC_ID":                  "service-ricplt-a1mediator-rmr.ricplt",
			"A1_RMR_RETRY_TIMES":          "20",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-a1mediator-env",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
	}

	configMap3 := &corev1.ConfigMap{
		Data: map[string]string{
			"ALARM_MGR_SERVICE_NAME": "service-ricplt-alarmmanager-rmr.ricplt",
			"ALARM_MGR_SERVICE_PORT": "4560",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ricplt",
			Name:      "configmap-ricplt-alarmmanager-appconfig",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
	}

	configMap4 := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alarm-appconfig",
			Namespace: "ricxapp",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: map[string]string{
			"ALARM_MGR_SERVICE_NAME": "service-ricplt-alarmmanager-rmr.ricplt",
			"ALARM_MGR_SERVICE_PORT": "4560",
		},
	}

	configMap5 := &corev1.ConfigMap{
		Data: map[string]string{
			"alarmmanagercfg": "{  \n" +
				"  \"local\": {\n" +
				"    \"host\": \":8080\"\n" +
				"  },\n" +
				"  \"logger\": {\n" +
				"    \"level\": 4\n" +
				"  },\n" +
				"  \"db\": {\n" +
				"    \"namespaces\": [\"sdl\", \"rnib\"]\n" +
				"  },\n" +
				"  \"rmr\": {\n" +
				"    \"protPort\": \"tcp:4560\",\n" +
				"    \"maxSize\": 1024,\n" +
				"    \"numWorkers\": 1\n" +
				"  },\n" +
				"  \"controls\": {\n" +
				"    \"promAlertManager\": {\n" +
				"      \"address\": \"cpro-alertmanager:80\",\n" +
				"      \"baseUrl\": \"api/v2\",\n" +
				"      \"schemes\": \"http\",\n" +
				"      \"alertInterval\": 30000\n" +
				"    },\n" +
				"    \"maxActiveAlarms\": 5000,\n" +
				"    \"maxAlarmHistory\": 20000,\n" +
				"    \"alarmInfoPvFile\": \"/mnt/pv-ricplt-alarmmanager/alarminfo.json\"\n" +
				"  }\n" +
				"}",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "configmap-ricplt-alarmmanager-alarmmanagercfg",
			Namespace: "ricplt",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
	}

	configMap6 := &corev1.ConfigMap{
		Data: map[string]string{
			"RMR_SEED_RT": "/cfg/uta_rtg.rt",
			"RMR_SRC_ID":  "service-ricplt-alarmmanager-rmr.ricplt",
			"RMR_RTG_SVC": "service-ricplt-rtmgr-rmr:4561",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "configmap-ricplt-alarmmanager-env",
			Namespace: "ricplt",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
	}

	configMap7 := &corev1.ConfigMap{
		Data: map[string]string{
			"appmgr.yaml": "\"local\":\n" +
				"  # Port on which the xapp-manager REST services are provided\n" +
				"  \"host\": \":8080\"\n" +
				"\"helm\":\n" +
				"  # Remote helm repo URL. UPDATE this as required.\n" +
				"  \"repo\": \"\\\"http://service-ricplt-xapp-onboarder-http:8080\\\"\"\n" +
				"\n" +
				"  # Repo name referred within the xapp-manager\n" +
				"  \"repo-name\": \"helm-repo\"\n" +
				"\n" +
				"  # Tiller service details in the cluster. UPDATE this as required.\n" +
				"  \"tiller-service\": service-tiller-ricxapp\n" +
				"  \"tiller-namespace\": ricinfra\n" +
				"  \"tiller-port\": \"44134\"\n" +
				"  # helm username and password files\n" +
				"  \"helm-username-file\": \"/opt/ric/secret/helm_repo_username\"\n" +
				"  \"helm-password-file\": \"/opt/ric/secret/helm_repo_password\"\n" +
				"  \"retry\": 1\n" +
				"\"xapp\":\n" +
				"  #Namespace to install xAPPs\n" +
				"  \"namespace\": \"ricxapp\"\n" +
				"  \"tarDir\": \"/tmp\"\n" +
				"  \"schema\": \"descriptors/schema.json\"\n" +
				"  \"config\": \"config/config-file.json\"\n" +
				"  \"tmpConfig\": \"/tmp/config-file.json\"\n" +
				"",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-appmgr-appconfig",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
	}

	configMap8 := &corev1.ConfigMap{
		Data: map[string]string{
			"appmgr-tiller-secret-copier.sh": "#!/bin/sh\n" +
				"if [ -x /svcacct-to-kubeconfig.sh ] ; then\n" +
				" /svcacct-to-kubeconfig.sh\n" +
				"fi\n" +
				"\n" +
				"if [ ! -z \"${HELM_TLS_CA_CERT}\" ]; then\n" +
				"  kubectl -n ${SECRET_NAMESPACE} get secret -o yaml ${SECRET_NAME} | \\\n" +
				"   grep 'ca.crt:' | \\\n" +
				"   awk '{print $2}' | \\\n" +
				"   base64 -d > ${HELM_TLS_CA_CERT}\n" +
				"fi\n" +
				"\n" +
				"if [ ! -z \"${HELM_TLS_CERT}\" ]; then\n" +
				"  kubectl -n ${SECRET_NAMESPACE} get secret -o yaml ${SECRET_NAME} | \\\n" +
				"   grep 'tls.crt:' | \\\n" +
				"   awk '{print $2}' | \\\n" +
				"   base64 -d > ${HELM_TLS_CERT}\n" +
				"fi\n" +
				"\n" +
				"if [ ! -z \"${HELM_TLS_KEY}\" ]; then\n" +
				"  kubectl -n ${SECRET_NAMESPACE} get secret -o yaml ${SECRET_NAME} | \\\n" +
				"   grep 'tls.key:' | \\\n" +
				"   awk '{print $2}' | \\\n" +
				"   base64 -d > ${HELM_TLS_KEY}\n" +
				"fi\n" +
				"",
			"svcacct-to-kubeconfig.sh": "#!/bin/sh\n" +
				"\n" +
				"# generate a kubconfig (at ${KUBECONFIG} file from the automatically-mounted\n" +
				"# service account token.\n" +
				"# ENVIRONMENT:\n" +
				"# SVCACCT_NAME: the name of the service account user.  default \"default\"\n" +
				"# CLUSTER_NAME: the name of the kubernetes cluster.  default \"kubernetes\"\n" +
				"# KUBECONFIG: where the generated file will be deposited.\n" +
				"SVCACCT_TOKEN=`cat /var/run/secrets/kubernetes.io/serviceaccount/token`\n" +
				"CLUSTER_CA=`base64 /var/run/secrets/kubernetes.io/serviceaccount/ca.crt|tr -d '\\n'`\n" +
				"\n" +
				"cat >${KUBECONFIG} <<__EOF__\n" +
				"ApiVersion: v1\n" +
				"kind: Config\n" +
				"users:\n" +
				"- name: ${SVCACCT_NAME:-default}\n" +
				"  user:\n" +
				"    token: ${SVCACCT_TOKEN}\n" +
				"clusters:\n" +
				"- cluster:\n" +
				"    certificate-authority-data: ${CLUSTER_CA}\n" +
				"    server: ${K8S_API_HOST:-https://kubernetes.default.svc.cluster.local/}\n" +
				"  name: ${CLUSTER_NAME:-kubernetes}\n" +
				"contexts:\n" +
				"- context:\n" +
				"    cluster: ${CLUSTER_NAME:-kubernetes}\n" +
				"    user: ${SVCACCT_NAME:-default}\n" +
				"  name: svcs-acct-context\n" +
				"current-context: svcs-acct-context\n" +
				"__EOF__\n" +
				"",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-appmgr-bin",
		},
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
	}

	configMap9 := &corev1.ConfigMap{
		Data: map[string]string{
			"RMR_RTG_SVC":       "4561",
			"HELM_TLS_CA_CERT":  "/opt/ric/secret/tiller-ca.cert",
			"HELM_TLS_CERT":     "/opt/ric/secret/helm-client.cert",
			"HELM_TLS_HOSTNAME": "service-tiller-ricxapp",
			"HELM_TLS_VERIFY":   "true",
			"NAME":              "xappmgr",
			"HELM_HOST":         "service-tiller-ricxapp.ricinfra:44134",
			"HELM_TLS_ENABLED":  "true",
			"HELM_TLS_KEY":      "/opt/ric/secret/helm-client.key",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "configmap-ricplt-appmgr-env",
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
	}
	
	return []*corev1.ConfigMap{configMap1, configMap2, configMap3, configMap4, configMap5, configMap6,configMap7, configMap8, configMap9}
}
