package jobcontroller

import (
	"github.com/go-logr/logr"
	"github.com/k8up-io/k8up/v2/operator/job"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// +kubebuilder:rbac:groups=batch,resources=jobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=batch,resources=jobs/status;jobs/finalizers,verbs=get;update;patch

// SetupWithManager configures the reconciler.
func (r *JobReconciler) SetupWithManager(mgr ctrl.Manager, _ logr.Logger) error {
	name := "job.k8up.io"
	pred, err := predicate.LabelSelectorPredicate(metav1.LabelSelector{MatchLabels: map[string]string{
		job.K8uplabel: "true",
	}})
	if err != nil {
		return err
	}
	r.Kube = mgr.GetClient()
	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&batchv1.Job{}, builder.WithPredicates(pred)).
		Complete(r)
}
