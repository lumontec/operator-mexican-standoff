/*
Copyright 2021.

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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	hellogroupv1 "example.com/hello/api/v1"
)

// HelloReconciler reconciles a Hello object
type HelloReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=hellogroup.example.com,resources=hellos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hellogroup.example.com,resources=hellos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hellogroup.example.com,resources=hellos/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Hello object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *HelloReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("start reconcile")

	var hello hellogroupv1.Hello

	// Gets the cronJob object from the cluster by key ---

	if err := r.Get(ctx, req.NamespacedName, &hello); err != nil {
		log.Error(err, "unable to fetch CronJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("retrieved resource by key", req.NamespacedName.String(), hello)

	// Construct our pod for Hello object
	constructPod := func(hello *hellogroupv1.Hello) (*corev1.Pod, error) {

		// create the pod obj
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				Name:        hello.Name,
				Namespace:   hello.Namespace,
			},
			Spec: corev1.PodSpec{
				RestartPolicy: corev1.RestartPolicyOnFailure,
				Containers: []corev1.Container{{
					Image:   "busybox",
					Name:    hello.Namespace,
					Command: []string{"echo", "Hello " + hello.Spec.Foo},
				}},
			},
		}

		return pod, nil
	}

	pod, err := constructPod(&hello)
	if err != nil {
		log.Error(err, "unable to construct pod from hello")
		// don't bother requeuing until we get a change to the spec
		return ctrl.Result{}, nil
	}

	if err := r.Create(ctx, pod); err != nil {
		log.Error(err, "unable to create Pod for Hello", "pod", pod)
		return ctrl.Result{}, err
	}

	log.V(1).Info("created Pod for Hello run", "pod", pod)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HelloReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hellogroupv1.Hello{}).
		Complete(r)
}
