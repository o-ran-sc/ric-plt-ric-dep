################################################################################
#   Copyright (c) 2022 Nokia.                                                  #
#                                                                              #
#   Licensed under the Apache License, Version 2.0 (the "License");            #
#   you may not use this file except in compliance with the License.           #
#   You may obtain a copy of the License at                                    #
#                                                                              #
#       http://www.apache.org/licenses/LICENSE-2.0                             #
#                                                                              #
#   Unless required by applicable law or agreed to in writing, software        #
#   distributed under the License is distributed on an "AS IS" BASIS,          #
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.   #
#   See the License for the specific language governing permissions and        #
#   limitations under the License.                                             #
################################################################################
{{- define "dbaas.env.redis.port.list" -}}
  {{- $port := ( include "common.serviceport.dbaas.redis" . ) -}}
  {{- range $idx := until ($.Values.dbaas.clusterSize|int) -}}
    {{- printf "%d" ( add ($port|int) $idx ) -}}
    {{- if lt $idx  ( sub ($.Values.dbaas.clusterSize|int) 1 ) -}}
      {{- printf "," -}}
    {{- end -}}
  {{- end }}
{{- end -}}

{{- define "dbaas.env.sentinel.port.list" -}}
  {{- $port := ( include "common.serviceport.dbaas.sentinel" . ) -}}
  {{- range $idx := until ($.Values.dbaas.clusterSize|int) -}}
    {{- printf "%d" ( add ($port|int) $idx ) -}}
    {{- if lt $idx  ( sub ($.Values.dbaas.clusterSize|int) 1 ) -}}
      {{- printf "," -}}
    {{- end -}}
  {{- end }}
{{- end -}}

{{- define "dbaas.env.sentinel.master.name.list" -}}
  {{- $masterPrefix := ( $.Values.dbaas.redis.masterGroupName ) -}}
  {{- range $idx := until ($.Values.dbaas.clusterSize|int) -}}
    {{- printf "%s-cluster-%d" $masterPrefix $idx -}}
    {{- if lt $idx  ( sub ($.Values.dbaas.clusterSize|int) 1 ) -}}
      {{- printf "," -}}
    {{- end -}}
  {{- end }}
{{- end -}}
