/*
Copyright 2022 rewanthtammana.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	validatev1 "github.com/rewanthtammana/crd-immutable-validation-webhook/crd-validation/api/v1"
)

// ImmutableKindReconciler reconciles a ImmutableKind object
type ImmutableKindReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=validate.rewanthtammana.com,resources=immutablekinds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=validate.rewanthtammana.com,resources=immutablekinds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=validate.rewanthtammana.com,resources=immutablekinds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ImmutableKind object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *ImmutableKindReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ImmutableKindReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&validatev1.ImmutableKind{}).
		Complete(r)
}
