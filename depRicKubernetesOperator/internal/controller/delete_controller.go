package controller

import (
	"context"
	"fmt"
)

func (r *RicPlatformReconciler) DeleteAll() {
	var err error
	namespaceProvided := "ricxapp"

	for _, resource := range GetConfigMap() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetConfigMap()| Error --> |", err)
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

	for _, resource := range GetService() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetService()| Error --> |", err)
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

}
