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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ricdeployv1 "ricdeploy/api/v1"
)

// RicPlatformReconciler reconciles a RicPlatform object
type RicPlatformReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *RicPlatformReconciler) handle_deploy_using_generated_go_code(usage string) {
	// return
	if usage == "create" {
		r.CreateAll()
	} else {
		r.DeleteAll()
	}
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

	logger := log.FromContext(ctx)
	logger.Info("Reconcilling RIC")
	instance := &ricdeployv1.RicPlatform{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// object not found, could have been deleted after reconcile request, hence don't requeue
			return ctrl.Result{}, nil
		}
		// error reading the object, requeue the request
		return ctrl.Result{}, err
	}

	// name of our custom finalizer
	myFinalizerName := "batch.tutorial.kubebuilder.io/finalizer"
	// examine DeletionTimestamp to determine if object is under deletion
	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// Adding a Finaliser also adds the DeletionTimestamp while deleting
		if !controllerutil.ContainsFinalizer(instance, myFinalizerName) {
			logger.Info("--- Job is in Creation state")
			//r.get_replicas(logger, "intent-config", instance)
			r.handle_deploy_using_generated_go_code("create")
			logger.Info("--- Job has been Created")
			controllerutil.AddFinalizer(instance, myFinalizerName)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}
	} else {
		// The object is being deleted
		if controllerutil.ContainsFinalizer(instance, myFinalizerName) {
			logger.Info("--- Job is in Deletion state")
			r.handle_deploy_using_generated_go_code("delete")
			logger.Info("--- Job has been Delete")
			// remove our finalizer from the list and update it.
			controllerutil.RemoveFinalizer(instance, myFinalizerName)
			if err := r.Update(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RicPlatformReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ricdeployv1.RicPlatform{}).
		Complete(r)
}
