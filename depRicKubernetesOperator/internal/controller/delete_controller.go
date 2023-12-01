package controller

import (
	"context"
	"fmt"
)

func (r *RicPlatformReconciler) DeleteAll() {
	var err error
	namespaceProvided := "ricplt"

	for _, resource := range GetCustomResourceDefinition() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetCustomResourceDefinition()| Error --> |", err)
		}
	}

	for _, resource := range GetJob() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetJob()| Error --> |", err)
		}
	}

	for _, resource := range GetDeployment() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetDeployment()| Error --> |", err)
		}
	}

	for _, resource := range GetPersistentVolume() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetPersistentVolume()| Error --> |", err)
		}
	}

	for _, resource := range GetClusterRole() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetClusterRole()| Error --> |", err)
		}
	}

	for _, resource := range GetEndpoints() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetEndpoints()| Error --> |", err)
		}
	}

	for _, resource := range GetConfigMap() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetConfigMap()| Error --> |", err)
		}
	}

	for _, resource := range GetService() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetService()| Error --> |", err)
		}
	}

	for _, resource := range GetRoleBinding() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetRoleBinding()| Error --> |", err)
		}
	}

	for _, resource := range GetStatefulSet() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetStatefulSet()| Error --> |", err)
		}
	}

	for _, resource := range GetIngress() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetIngress()| Error --> |", err)
		}
	}

	for _, resource := range GetPersistentVolumeClaim() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetPersistentVolumeClaim()| Error --> |", err)
		}
	}

	for _, resource := range GetServiceAccount() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetServiceAccount()| Error --> |", err)
		}
	}

	for _, resource := range GetRole() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetRole()| Error --> |", err)
		}
	}

	for _, resource := range GetSecret() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetSecret()| Error --> |", err)
		}
	}

	for _, resource := range GetClusterRoleBinding() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetClusterRoleBinding()| Error --> |", err)
		}
	}

}