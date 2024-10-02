/*
Copyright 2024.

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

package controller

import (
	"context"
	crdv1 "github.com/arnab-baishnab/learning-kubebuilder/api/v1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	_ "strings"
)

const (
	ourKind = "MyKind"
)

// MyKindReconciler reconciles a MyKind object
type MyKindReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
	ctx    context.Context
	myKind *crdv1.MyKind
}

//+kubebuilder:rbac:groups=mygroup.mydomain.com,resources=mykinds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=mygroup.mydomain.com,resources=mykinds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=mygroup.mydomain.com,resources=mykinds/finalizers,verbs=update

func (r *MyKindReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	r.Log = ctrl.Log.WithValues("MyKind", req.NamespacedName)
	r.ctx = ctx

	// TODO(user): your logic here

	klog.Info("Reconciling MyKind ->>>>>>>>>>>>>>")

	var mykind crdv1.MyKind

	if err := r.Get(ctx, req.NamespacedName, &mykind); err != nil {
		r.Log.Info("mykind not found")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.myKind = &mykind

	klog.Info("name-namespace ", mykind.Name, mykind.Namespace)

	if err := r.CheckDeployment(); err != nil {
		return ctrl.Result{}, err
	}
	//if err := r.CheckService(); err != nil {
	//	return ctrl.Result{}, err
	//}
	return ctrl.Result{}, nil
}
