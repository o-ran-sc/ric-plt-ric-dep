#!/bin/bash
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

echo "Installing servecm (Chart Manager) and common templates to helm3"

helm plugin install https://github.com/jdolitsky/helm-servecm
eval $(helm env |grep HELM_REPOSITORY_CACHE) 
echo ${HELM_REPOSITORY_CACHE}

# servecm can download chartmuseum binary, but the hardcoded location does not work anymore
# so, we download it now before starting servecm, so that it's available when servecm
# tries to start the chartmuseum binary
curl -LO https://get.helm.sh/chartmuseum-v0.15.0-linux-386.tar.gz
tar xzvf chartmuseum-v0.15.0-linux-386.tar.gz
chmod +x ./linux-386/chartmuseum
cp ./linux-386/chartmuseum /usr/local/bin

nohup helm servecm --port=8879 --context-path=/charts --storage local --storage-local-rootdir $HELM_REPOSITORY_CACHE/local/ <<EOF &
yes
EOF

CURL_CMD="curl --silent --output /dev/null  http://127.0.0.1:8879/charts"
`${CURL_CMD}`
READY=$?
while [ ${READY} != 0 ]; do
        echo "servecm not yet running. sleeping for 2 seconds"
        sleep 2
        `${CURL_CMD}`
        READY=$?
done
echo "servcm up and running"

eval $(helm env |grep HELM_REPOSITORY_CACHE)
echo ${HELM_REPOSITORY_CACHE}
mkdir -p "${HELM_REPOSITORY_CACHE}/local/"

export COMMON_CHART_VERSION=$(cat ../ric-common/Common-Template/helm/ric-common/Chart.yaml | grep version | awk '{print $2}')
helm package -d /tmp ../ric-common/Common-Template/helm/ric-common
cp /tmp/ric-common-${COMMON_CHART_VERSION}.tgz "${HELM_REPOSITORY_CACHE}/local/"
helm repo remove local
helm repo add local http://127.0.0.1:8879/charts

echo "checking that ric-common templates were added"
helm search repo local/ric-common



