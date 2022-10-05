/*
Copyright 2022 rewanthtammana.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var immutablekindlog = logf.Log.WithName("immutablekind-resource")

func (r *ImmutableKind) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-validate-rewanthtammana-com-v1-immutablekind,mutating=false,failurePolicy=fail,sideEffects=None,groups=validate.rewanthtammana.com,resources=immutablekinds,verbs=create;update,versions=v1,name=vimmutablekind.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &ImmutableKind{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ImmutableKind) ValidateCreate() error {
	immutablekindlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ImmutableKind) ValidateUpdate(old runtime.Object) error {
	immutablekindlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ImmutableKind) ValidateDelete() error {
	immutablekindlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
