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
	for _, resource := range GetPersistentVolume() {
		if resource.GetNamespace() == "" {
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetPersistentVolume()| Error --> |", err)
		}
	}
	for _, resource := range GetRoleBinding() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetRoleBinding()| Error --> |", err)
		}
	}
	for _, resource := range GetPersistentVolumeClaim() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetPersistentVolumeClaim()| Error --> |", err)
		}
	}

	for _, resource := range GetServiceAccount() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetServiceAccount()| Error --> |", err)
		}
	}

	for _, resource := range GetRole() {
		if resource.ObjectMeta.Namespace == "" {
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetRole()| Error --> |", err)
		}
	}

}
