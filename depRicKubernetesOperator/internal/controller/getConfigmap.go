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

	return []*corev1.ConfigMap{configMap1, configMap2}
}
