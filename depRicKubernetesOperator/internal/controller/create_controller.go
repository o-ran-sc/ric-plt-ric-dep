package controller

import (
	"context"
	"fmt"
)

func (r *RicPlatformReconciler) CreateAll() {
	var err error
	namespaceProvided := "ricplt"

	for _, resource := range GetConfigMap() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetConfigMap()| Error --> |", err)
		}
	}

	for _, resource := range getDeployment() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetDeployment()| Error --> |", err)
		}
	}

	for _, resource := range GetService() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetService()| Error --> |", err)
		}
	}
	for _, resource := range GetIngress() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetIngress()| Error --> |", err)
		}
	}

}
