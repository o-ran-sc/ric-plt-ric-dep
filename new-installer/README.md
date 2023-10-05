# Overview

**Note: This is WIP so your mileage may vary, but you are welcome to create an issue if there are problems you are facing using this documentation.**

This directory contains the source code for the new installer for the NearRT RIC Platform. The new installer leverages a single helm chart for installation. The components deployed during installations are determined by various flags present in the `values.yaml` file of the chart, which can be overridden using an override file.

The code is organized as follows

  |
  + helm
  | |
  | + charts
  |   |
  |   + Makefile
  |   + nearrtric/
  |     |
  |     + Makefile
  |     + a1mediator/
  |     + appmgr/
  |     + dbaas/
  |     + e2mgr/
  |     + e2term/
  |     + nearrt-ric-common/
	|		...
  |
  + helm-overrides/
    |
    + nearrtric/
      |
      + example-override-file.yaml

# Supported Versions

Following software versions are currently supported.

  - `minikube` with supported `k8s` version.
  - kubernetes 1.22 and above. A number of API versions were deprecated starting from this version. This should thus serve as a good
  - Helm Charts version 3

This setup is tested on Ubuntu 22.04 with following `minikube`, `k8s`, `kubectl` and `helm` versions.

```bash
$ uname -mvp
#26~22.04.1-Ubuntu SMP PREEMPT_DYNAMIC Thu Jul 13 16:27:29 UTC 2 x86_64 x86_64

$ kubectl version
WARNING: This version information is deprecated and will be replaced with the output from kubectl version --short.  Use --output=yaml|json to get the full version.
Client Version: version.Info{Major:"1", Minor:"24", GitVersion:"v1.24.1", GitCommit:"3ddd0f45aa91e2f30c70734b175631bec5b5825a", GitTreeState:"clean", BuildDate:"2022-05-24T12:26:19Z", GoVersion:"go1.18.2", Compiler:"gc", Platform:"linux/amd64"}
Kustomize Version: v4.5.4
Server Version: version.Info{Major:"1", Minor:"26", GitVersion:"v1.26.1", GitCommit:"8f94681cd294aa8cfd3407b8191f6c70214973a4", GitTreeState:"clean", BuildDate:"2023-01-18T15:51:25Z", GoVersion:"go1.19.5", Compiler:"gc", Platform:"linux/amd64"}
WARNING: version difference between client (1.24) and server (1.26) exceeds the supported minor version skew of +/-1

$ minikube version
minikube version: v1.29.0
commit: ddac20b4b34a9c8c857fc602203b6ba2679794d3

$ helm version
version.BuildInfo{Version:"v3.11.2", GitCommit:"912ebc1cd10d38d340f048efaf0abda047c3468e", GitTreeState:"clean", GoVersion:"go1.18.10"}

```

# Getting Started

## Setup and Pre-requisites

These instructions assume that you have `minikube` and `kubectl` installed.

## Starting `minikube`

First start `minikube` using the following command.

```bash
$ minikube start --driver=docker

```
We are using the 'docker' driver because this allows us to set up tunnels and expose services to the host using the Kubernetes Ingress Controller. The other drivers may also work but are not tested so far. This will start the 'minikube kubernetes cluster'. You can interact with this cluster using `kubectl`.

Also, note that you may have to set the environment variables like `HTTP_PROXY`, `HTTPS_PROXY` and `NO_PROXY` if you are behind an HTTP Proxy. Otherwise Kubernetes won't be able to pull the images.

```bash
$ kubectl get pods -A
```

## Creating Platform and Xapp Namespaces

After this step we need to create the namespaces for our deployment. Note: The namespaces need to be created every time we create a new cluster (not upon restarting the cluster).

```bash
$ kubectl create ns ricplt
$ kubectl create ns ricxapp
```

This sets up the cluster for subsequent installation of the NearRT RIC platform components.

## Installing and Running Chart Museum

One current limitation of this installation is - the XApp on-boarding functionality is not fully integrated into the `k8s` deployment as a result following steps need to be run manually for on-boarding and  creating installation packages of XApps. This will be eventually part of the `k8s` deployment (likely as a separate `helm` chart deployed in a separate namespace `k8s`).

Please follow the steps mentioned below to install `helm`, `chartmuseum` and `cm-push` plugin for `helm`.

```bash
# Following commands are run as non root user. The idea is to keep the privileges to minimum.

# Create a `temp/` directory in the home folder that will be used to download all the artifacts required during the bring up.
$ mkdir temp
$ cd temp
$ wget https://get.helm.sh/helm-v3.11.2-linux-amd64.tar.gz
$ wget https://get.helm.sh/chartmuseum-v0.13.1-linux-amd64.tar.gz
$ tar xvzpf helm-v3.11.2-linux-amd64.tar.gz
$ tar xvzpf chartmuseum-v0.13.1-linux-amd64.tar.gz

# It is assumed that ~/bin/ is in your $PATH
$ cp linux-amd64/chartmuseum linux-amd64/helm ~/bin
$ helm version
$ chartmuseum -version

# To Push to chart museum, we will need to install a Chart Museum push plugin.
$ helm plugin install https://github.com/chartmuseum/helm-push

# This should have the `cm-push` command available.
$ helm plugin list
NAME    VERSION DESCRIPTION
cm-push 0.10.3  Push chart package to ChartMuseum
```

Run the following command to run Chart museum.

```bash
# This step is fairly simple, simply run the chartmusem command with appropriate options
# We are using local storage backend and port 6873 (MUSE) and a directory called `helm/chartmuseum/`
# The setup is kept simple.
$ chartmuseum --debug --port 6873 --storage local --storage-local-rootdir $HOME/helm/chartsmuseum/
```

## Setting up Helm Repositories

### Setting up `local` Helm Repo

We will be building the `helm` charts locally and those charts will be 'push'ed to the local repository. Later during installations charts will be downloaded from this `local` repository. The `chartmuseum` server described above will serve this 'local' repository.

```bash
$ helm repo add local http://localhost:6873/
$ helm repo list
```
### Setting up Third Party repositories

This installation leverages the chart available from the third-party repositories, instead of maintaining those charts ourselves. This allows us to update third party dependencies easily and we don't need to maintain our own 'charts' just for this alone.

For the installation of the Near RT RIC, we are making use of `influxdb2` (for storing KPI data by `kpimon-go` application). The charts for `influxdb2` are available from the Influxdata repository. Hence this repository needs to be added inside `helm`.

```bash
$ helm repo add influxdata https://helm.influxdata.com
$ helm repo list
```


## Running Near RT RIC

### Preparing the Charts

We will first build the charts based on the data from this repository. We use `make` for building the charts. To build the charts, following instructions can be followed.

```bash
# Make sure that chart museum is running already as discussed above
#
# First we will prepare chart(s) for the installation. This is done using Makefile
# `make` tool needs to be installed for this.
$ cd helm/charts
$ make nearrtric

##### Output ignored

$ helm search repo local/nearrtric
NAME           	CHART VERSION	APP VERSION	DESCRIPTION
local/nearrtric	0.1.0        	           	Umbrella Helm Chart for NearRT RAN Intelligent ...
```
This means the `local/nearrtric` chart is available and can be installed using `helm install`.


### Installing the Charts

Currently, the Near RT RIC contains mandatory components and some optional components. Following are mandatory components -

1. `appmgr`
2. `dbaas`
3. `e2mgr`
4. `e2term`
5. `submgr`
6. `rtmgr`

and following are optional components -

1. `a1mediator`
2. `alarmmanager`
3. `jaegeradapter`
4. `o1mediator`
5. `vespamgr`
6. `xapp-onboarder`

and following are third party (optional) components -

1. `influxdb2`

The default `values.yaml` file disables all optional components. The optional components can be enabled using the override file inside `helm-overrides/nearrtric/` directory. The provided file enables `influxdb2` and `a1mediator` optional components. In addition, additional configurations like using docker image etc can be changed for all the components. Please refer to the provided example file for the supported configurations.

Run the following command to install the chart.

```bash
$ helm install nearrtric -n ricplt local/nearrtric -f example_recipe_latest_stable.yaml
```
Note: You can refer to `helm-overrides/nearrtric/minimal-nearrt-ric.yaml` file and follow the conventions used in this override file.

This should install the main `nearrtric` chart and it's dependencies. The installation can be verified using the following `kubectl` commands.

```bash
$ kubectl get pods -n ricplt
NAME                                              READY   STATUS    RESTARTS        AGE
deployment-ricplt-a1mediator-84fc865778-x846h     1/1     Running   1 (4m13s ago)   5m16s
deployment-ricplt-appmgr-57cc4d665b-lb8dg         1/1     Running   0               5m16s
deployment-ricplt-e2mgr-9748f9585-mg2zl           1/1     Running   3 (4m47s ago)   5m16s
deployment-ricplt-e2term-alpha-5ffb57bf9f-slmrz   1/1     Running   0               5m16s
deployment-ricplt-rtmgr-57f7c7797f-mpkg9          1/1     Running   0               5m16s
deployment-ricplt-submgr-74f67bf444-qh5rn         1/1     Running   0               5m16s
ricplt-influxdb-0                                 1/1     Running   0               5m16s
statefulset-ricplt-dbaas-server-0                 1/1     Running   0               5m16s

$ kubectl get svc -n ricplt
NAME                                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)              AGE
ricplt-influxdb                          ClusterIP   10.105.150.100   <none>        8086/TCP             6m
service-ricplt-a1mediator-http           ClusterIP   10.101.218.159   <none>        10000/TCP            6m1s
service-ricplt-a1mediator-rmr            ClusterIP   10.103.242.185   <none>        4561/TCP,4562/TCP    6m1s
service-ricplt-appmgr-http               ClusterIP   10.111.10.61     <none>        8080/TCP             6m1s
service-ricplt-appmgr-rmr                ClusterIP   10.105.106.102   <none>        4561/TCP,4560/TCP    6m1s
service-ricplt-dbaas-tcp                 ClusterIP   None             <none>        6379/TCP             6m1s
service-ricplt-e2mgr-http                ClusterIP   10.97.133.251    <none>        3800/TCP             6m1s
service-ricplt-e2mgr-rmr                 ClusterIP   10.108.221.142   <none>        4561/TCP,3801/TCP    6m1s
service-ricplt-e2term-prometheus-alpha   ClusterIP   10.109.194.105   <none>        8088/TCP             6m1s
service-ricplt-e2term-rmr-alpha          ClusterIP   10.109.22.225    <none>        4561/TCP,38000/TCP   6m
service-ricplt-e2term-sctp-alpha         NodePort    10.96.249.236    <none>        36422:32222/SCTP     6m
service-ricplt-rtmgr-http                ClusterIP   10.96.22.95      <none>        3800/TCP             6m
service-ricplt-rtmgr-rmr                 ClusterIP   10.99.242.83     <none>        4561/TCP,4560/TCP    6m
service-ricplt-submgr-http               ClusterIP   10.105.226.85    <none>        8088/TCP             6m
service-ricplt-submgr-rmr                ClusterIP   10.96.246.248    <none>        4560/TCP,4561/TCP    6m
```

### Deploying the Simulator

You can test the installed Near RT RIC by installing the simulator from the [sim-e2-interface](https://gerrit.o-ran-sc.org/r/admin/repos/sim/e2-interface,general) repository. If you are installing on Ubuntu 22.04, The image needs to be built using the Ubuntu 22.04 Builder Image. Change the appropriate lines in the `Dockerfile` in the directory `e2sim/e2sm_examples/kpm_e2sm`. The Ubuntu 22.04 Builder image can be used as follows

```bash

 # the builder has: git, wget, cmake, gcc/g++, make, python2/3. v7 dropped nng support
 #
-ARG CONTAINER_PULL_REGISTRY=nexus3.o-ran-sc.org:10001
-FROM ${CONTAINER_PULL_REGISTRY}/o-ran-sc/bldr-ubuntu18-c-go:1.9.0 as buildenv
+ARG CONTAINER_PULL_REGISTRY=nexus3.o-ran-sc.org:10004
+# FROM nexus3.o-ran-sc.org:10004/o-ran-sc/bldr-ubunut22-c-go:0.1.0 as buildenv
+FROM ${CONTAINER_PULL_REGISTRY}/o-ran-sc/bldr-ubuntu22-c-go:0.1.0 as buildenv

```

It is also required to update the IP Address of the `service-ricplt-e2term-sctp-alpha` service above in the `Dockerfile`. (Agreed this is a bit ugly and will be fixed soon!).

```bash
-CMD kpm_sim 10.110.102.29 36422
+CMD kpm_sim 10.96.249.236 36422
```
The Docker Image that needs to be built should be using the 'docker environment' of the `minikube`. This environment can be set by running the following command.

```bash
```

The simulator then can be started as follows

```bash

# Setup the 'docker env'
$ eval $(minikube -p minikube docker-env)

# Build docker image
$ docker build -t e2sim:latest .

# Make sure the `helm/deployments/deployment.yaml` file points to `e2sim:latest` images

$ helm install e2sim -n ricplt helm/

```

### Deploying XApps

You can deploy the XApps. The approach we follow is preparing the chart images for the XApps first and then simply using `helm install`. The Chart images can be prepared by using the Xapp Onboarder. (Note: A future version of deployment would make this available as a single `helm` chart installation, currently this manual step needs to be followed.)

```bash
$ git clone https://gerrit.o-ran-sc.org/r/ric-plt/appmgr
$ cd appmgr/xapp_orchestrater/dev/xapp_onboarder

# Create a Virtual Environment to install the `xapp_onboarder`
$ python3 -m venv venv3
$ . venv/bin/activate
$ pip install -r requirements.txt

# Note: The Chart Repo URL points to the `chartmusem` server running started.
# Also note the `shcema` typo :-)
$ CHART_REPO_URL=http://localhost:6873 dms_cli onboard --config-file-path <path-to-app-config> --shcema_file_path <path-to-schema-json>
$ CHART_REPO_URL=http://localhost:6873 dms_cli download_helm_chart <chart-name> <version>
# This step downloads the file `chart-name-chart-version.tgz`

# Finally install the Chart in `ricxapp` namespace
$ helm install <release-name> -n ricxapp <chart-name-chart-version.tgz>
```

You can then check the logs of individual pods using `kubectl logs ...` command to check if the pods, Xapps are running as expected.
