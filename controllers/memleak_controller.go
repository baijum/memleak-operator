/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cachev1 "github.com/baijum/memleak-operator/api/v1"
	olmv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
)

// MemleakReconciler reconciles a Memleak object
type MemleakReconciler struct {
	client.Client
	dynClient dynamic.Interface
	Scheme    *runtime.Scheme
}

type customResourceDefinition struct {
	resource   *unstructured.Unstructured
	serviceGVR *schema.GroupVersionResource
	client     dynamic.Interface
	ns         string
}

//+kubebuilder:rbac:groups=cache.example.com,resources=memleaks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.example.com,resources=memleaks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.example.com,resources=memleaks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Memleak object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *MemleakReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	fmt.Println("BEFORE CSV CALLING")
	csvs, err := r.dynClient.Resource(olmv1alpha1.SchemeGroupVersion.WithResource("clusterserviceversions")).Namespace("default").List(context.Background(),
		metav1.ListOptions{})
	fmt.Printf("bbbbbbbbbb %v", csvs)
	_ = csvs
	_ = err
	fmt.Println("AFTER CSV CALLING")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MemleakReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.dynClient, _ = dynamic.NewForConfig(mgr.GetConfig())
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1.Memleak{}).
		Complete(r)
}
