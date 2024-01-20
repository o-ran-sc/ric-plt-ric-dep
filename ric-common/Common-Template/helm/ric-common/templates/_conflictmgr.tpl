################################################################################
#   Copyright (c) 2023 Capgemini                                               #
#                                                                              #
################################################################################
{{- define "common.name.conflictmgr" -}}
  {{- printf "conflictmgr" -}}
{{- end -}}

{{- define "common.fullname.conflictmgr" -}}
  {{- $name := ( include "common.name.conflictmgr" . ) -}}
  {{- $namespace := ( include "common.namespace.platform" . ) -}}
  {{- printf "%s-%s" $namespace $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.configmapname.conflictmgr" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "configmap-%s" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.deploymentname.conflictmgr" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "deployment-%s" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.containername.conflictmgr" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "container-%s" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.serviceaccountname.conflictmgr" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "svcacct-%s" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.ingressname.conflictmgr" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "ingress-%s" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.kongpath.ric.conflictmgr" -}}/conflictmgr{{- end -}}

{{- define "common.servicename.conflictmgr.rmr" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "service-%s-rmr" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.serviceport.conflictmgr.rmr.data" -}}4560{{- end -}}
{{- define "common.serviceport.conflictmgr.rmr.route" -}}4561{{- end -}}
{{- define "common.serviceport.conflictmgr.http" -}}10000{{- end -}}


{{- define "common.servicename.conflictmgr.grpc" -}}
  {{- $name := ( include "common.fullname.conflictmgr" . ) -}}
  {{- printf "service-%s-grpc" $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "common.serviceport.conflictmgr.grpc.data" -}}50051{{- end -}}
