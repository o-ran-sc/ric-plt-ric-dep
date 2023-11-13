/*
==================================================================================

	Copyright (c) 2023 Samsung

	 Licensed under the Apache License, Version 2.0 (the "License");
	 you may not use this file except in compliance with the License.
	 You may obtain a copy of the License at

	     http://www.apache.org/licenses/LICENSE-2.0

	 Unless required by applicable law or agreed to in writing, software
	 distributed under the License is distributed on an "AS IS" BASIS,
	 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	 See the License for the specific language governing permissions and
	 limitations under the License.

	 This source code is part of the near-RT RIC (RAN Intelligent Controller)
	 platform project (RICP).

==================================================================================
*/
package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ricdeployv1 "ricdeploy/api/v1"
)

// RicPlatformReconciler reconciles a RicPlatform object
type RicPlatformReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ricdeploy.ricplt.com,resources=ricplatforms,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ricdeploy.ricplt.com,resources=ricplatforms/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ricdeploy.ricplt.com,resources=ricplatforms/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RicPlatform object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *RicPlatformReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	// TODO(user): your logic here
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RicPlatformReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ricdeployv1.RicPlatform{}).
		Complete(r)
}
