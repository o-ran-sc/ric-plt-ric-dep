..  ============LICENSE_START=======================================================
..
..  ================================================================================
..  Copyright (C) 2020 Hcl Technologies Limited.
..  ================================================================================
..  Licensed under the Apache License, Version 2.0 (the "License");
..  you may not use this file except in compliance with the License.
..  You may obtain a copy of the License at
..
..       http://www.apache.org/licenses/LICENSE-2.0
..
..  Unless required by applicable law or agreed to in writing, software
..  distributed under the License is distributed on an "AS IS" BASIS,
..  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
..  See the License for the specific language governing permissions and
..  limitations under the License.
..  ============LICENSE_END=========================================================

Important
^^^^^^^^^
The redis-cluster currently is NOT part of RIC platform & hence is completely optional.
This piece of document has been created as part of delivery item for below jira ticket 
https://jira.o-ran-sc.org/browse/RIC-109 
This ticket is about assessing the feasibility of redis-cluster (with data sharding) 
supporting desired pod anti-affinity for high availability as per the ticket.

Overview
^^^^^^^^
This document describes the environment/conditions used to test the feasibility of Redis 
cluster set-up as detailed in the above ticket. Redis Cluster is a distributed implementation 
of Redis with high performance goals. More details at https://redis.io/topics/cluster-spec

Environment Set-Up
^^^^^^^^^^^^^^^^^^
The set up was tested with kubernetes v1.19 cluster with 
   #. Pod topology spread constraint enabled
      Reference: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints
   #. CEPH as the Cluster Storage Solution.
   #. Three worker nodes in the kubernet cluster

Execution
^^^^^^^^^
Once environment is set-up,  a redis-cluster can be set up using the helm-chart (also provided with 
this commit). Once cluster is running, any master/slave of the redis instance pods can be deleted which
will be compensated automatically by new instances

At this stage the perl utility program (included with helm-chart) can be run. The helm chart installation
output generates the requirement commands to invoke.

This utility program identifies the missing anti-affinity(as per above ticket) of redis instances required
in a redis-cluster. When executed it communicates to redis nodes to switch roles (e.g. master/slave)
such that the end-state meets the desired anti-affinity. 

   
