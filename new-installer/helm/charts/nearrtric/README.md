# Introduction

This directory contains the charts required to deploy the Near RT RIC platform. Near RT RIC platform can be deployed using a single chart called `nearrtric`. All the platform components are dependencies of this chart. The goal of this project is it should be possible to easily deploy the platform during development and/or testing as the current procedure is a bit involved. The goal is to avoid manual steps as much as possible and keep things simple to a few commands and/or charts.

Also, the goal of this is to make sure that the platform can be deployed with latest `k8s` and it's many desktop flavors like `minikube`, `microk8s` etc. Initial focus is on making sure that things work on `minikube`. The Goal is also to make sure that the charts can be deployed on latest version of Ubuntu, but in general it should be possible to work with any of the latest distribution.

It is planned that this will be actively maintained.

# Getting Started

Each of the individual component is a helm chart. Individual components can be enabled during deployment through the values like following. This allows us to select a subset of components that are of interest. For example, to simply test an App SDK and an App that consumes messages from E2 node(s). We might want to skip the entire Policy/VES set of components.


## Chart Museum

In the current deployment we are also deploying chart museum along with the platform. While this may be useful for preparing charts for the XApps etc. the way it is deployed is quite confusing and it's role is not quite clear during deployment. This approach keeps the chart museum totally separate. This has been found very useful.

## Third Party components

Often we are using third party components for certain functionality. For example we are using `influxdb2` for saving KPI-Mon data in the time series database. We will be using the official charts of the third party dependencies. Also, the latest versions of those charts will be used (as appropriate).

## Minikube

This section describes steps for getting started with `minikube`. While the steps below should work with different `minikube` drivers, for exposing a set of services from the platform through ingress controllers, the current discussion only discusses the `minikube` with `docker` driver. It should be possible to work with other drivers, but at the moment is not actively supported, so your mileage may vary.

## Local Development

Often the developers would want to test their code in the environment in which it would be deployed. It is possible to build all the docker images required to use in the platform locally, following instructions can be followed. This takes example of the `e2` repository, similar approach can be followed for other platform components and/or Xapp development.
