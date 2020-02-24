.. This work is licensed under a Creative Commons Attribution 4.0 International License.
.. SPDX-License-Identifier: CC-BY-4.0
.. ===============LICENSE_START=======================================================
.. Copyright (C) 2019-2020 AT&T Intellectual Property
.. ===================================================================================
.. This documentation file is distributed under the Creative Commons Attribution
.. 4.0 International License (the "License"); you may not use this file except in
.. compliance with the License.  You may obtain a copy of the License at
..
.. http://creativecommons.org/licenses/by/4.0
..
.. This file is distributed on an "AS IS" BASIS,
.. WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
.. See the License for the specific language governing permissions and
.. limitations under the License.
.. ===============LICENSE_END=========================================================

===================
Installation Guides
===================

This document describes how to install the RIC components deployed by scripts and Helm charts
under the ric-plt/dep repository, including the dependencies and required system resources.

.. contents::
   :depth: 3
   :local:

Version history
===============

+--------------------+--------------------+--------------------+--------------------+
| **Date**           | **Ver.**           | **Author**         | **Comment**        |
|                    |                    |                    |                    |
+--------------------+--------------------+--------------------+--------------------+
| 2020-02-20         | 0.1.0              |Abdulwahid W        |                    |
|                    |                    |                    |                    |
+--------------------+--------------------+--------------------+--------------------+


Overview
========

This section explains the installation of Near Realtime RAN Intelligent Controller Platform only.

Prerequisites
=============

Both RIC and AUX clusters need to fulfill the following prerequisites.

- Kubernetes v.1.16.0 or above
- helm v2.12.3 or above
- Read-write access to directory /mnt

The following two sections show two example methods to create an environment for installing RIC.

VirtualBox VMs as Installation Hosts
------------------------------------

The deployment of Near Realtime RIC can be done on a wide range of hosts, including
bare metal servers, OpenStack VMs, and VirtualBox VMs.  This section provides detailed instructions
for setting up Oracle VirtualBox VMs to be used as installation hosts.

.. include:: ./installation-virtualbox.rst


One-Node Kubernetes Cluster
---------------------------

This section describes how to set up a one-node Kubernetes cluster onto a VM installation host.

.. include:: ./installation-k8s1node.rst


Installing Near Realtime RIC in RIC Cluster
===========================================

After the Kubernetes cluster is installed, the next step is to install the (Near Realtime) RIC Platform.

.. include:: ./installation-ric.rst
