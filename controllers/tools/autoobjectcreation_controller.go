/*


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

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	toolsaocv1 "github.com/ericogr/k8s-aoc/apis/tools/v1"
)

var (
	aocOnwerKey = ".metadata.controller"
	aocGV       = toolsaocv1.GroupVersion.String()
)

// AutoObjectCreationReconciler reconciles a AutoObjectCreation object
type AutoObjectCreationReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// SetupWithManager setup
func (r *AutoObjectCreationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// if err := mgr.GetFieldIndexer().IndexField(&corev1.Namespace{}, aocOnwerKey, func(rawObj runtime.Object) []string {
	// 	namespace := rawObj.(*corev1.Namespace)
	// 	owner := metav1.GetControllerOf(namespace)

	// 	if owner == nil {
	// 		return nil
	// 	}

	// 	if owner.APIVersion != aocGV || owner.Kind != "AutoObjectCreation" {
	// 		return nil
	// 	}

	// 	return []string{owner.Name}
	// }); err != nil {
	// 	return err
	// }

	return ctrl.NewControllerManagedBy(mgr).
		For(&toolsaocv1.AutoObjectCreation{}).
		Watches(&source.Kind{Type: &corev1.Namespace{}}, &handler.EnqueueRequestForObject{}).
		// Owns(&corev1.Namespace{}).
		Complete(r)
}

// +kubebuilder:rbac:groups=tools.aoc.github.com,resources=autoobjectcreations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tools.aoc.github.com,resources=autoobjectcreations/status,verbs=get;update;patch

// Reconcile k8s reconcile
func (r *AutoObjectCreationReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	fmt.Println("-----------------------")
	fmt.Println(req.NamespacedName.Name)
	fmt.Println("-----------------------")
	return ctrl.Result{}, nil

	// ctx := context.Background()
	// log := r.Log.WithValues("autoobjectcreation", req.NamespacedName)
	// var aoc toolsaocv1.AutoObjectCreation
	// err := r.Get(ctx, req.NamespacedName, &aoc)

	// if err != nil {
	// 	if errors.IsNotFound(err) {
	// 		// Object not found, return. Created objects are automatically garbage collected
	// 		return ctrl.Result{}, nil
	// 	}

	// 	// Error reading the object - requeue the request.
	// 	return ctrl.Result{}, err
	// }

	// namespaces, err := r.FindNamespacesByAnnotation(aoc.Spec.Trigger.Annotations)
	// processor := processor.Processor{Client: r.Client}
	// var listErrors string
	// for _, namespace := range namespaces {
	// 	reference := "[" + aoc.Spec.Template.Kind + "(" + aoc.Spec.Template.Name + ")] at " + namespace.ObjectMeta.Name + " namespace"
	// 	log.Info("Ready to process " + reference)

	// 	newObj, gvk, err := processor.ToObject(aoc.Spec.Template, namespace)

	// 	if err != nil {
	// 		strErr := "Error serializing " + reference
	// 		listErrors += strErr + "\n"
	// 		log.Error(err, strErr)
	// 		break
	// 	}

	// 	log.Info("Object encoded succefully " + reference)

	// 	findObj, err := processor.GetObject(
	// 		*gvk,
	// 		types.NamespacedName{
	// 			Namespace: namespace.Name,
	// 			Name:      aoc.Spec.Template.Name,
	// 		},
	// 	)

	// 	if err != nil && errors.IsNotFound(err) {
	// 		log.Info("Creating new object " + reference)
	// 		err := r.Client.Create(ctx, &newObj)

	// 		if err == nil {
	// 			log.Info("Create succefully " + reference)
	// 		} else {
	// 			strErr := "Error creating object " + reference
	// 			listErrors += strErr + "\n"
	// 			log.Error(err, strErr)
	// 		}
	// 	} else {
	// 		if err == nil {
	// 			findObj.Object["spec"] = newObj.Object["spec"]
	// 			err := r.Client.Update(ctx, &findObj)

	// 			if err == nil {
	// 				log.Info("Update succefully " + reference)
	// 			} else {
	// 				strErr := "Error updating object " + reference
	// 				listErrors += strErr + "\n"
	// 				log.Error(err, strErr)
	// 			}
	// 		} else {
	// 			strErr := "Error getting object " + reference
	// 			listErrors += strErr + "\n"
	// 			log.Error(err, strErr)
	// 		}
	// 	}
	// }

	// // //https://godoc.org/sigs.k8s.io/controller-runtime/pkg/predicate#GenerationChangedPredicate
	// if listErrors != "" {
	// 	aoc.Status.Status = listErrors
	// } else {
	// 	aoc.Status.Status = "OK"
	// }

	// if err := r.Status().Update(ctx, &aoc); err != nil {
	// 	log.Error(err, "Unable to update status")
	// 	return ctrl.Result{}, err
	// }

	// return ctrl.Result{}, nil
}

// FindNamespaces find all namespaces
func (r *AutoObjectCreationReconciler) FindNamespaces() (namespaces []corev1.Namespace, err error) {
	namespacesList := &corev1.NamespaceList{}
	err = r.Client.List(context.Background(), namespacesList)
	namespaces = namespacesList.Items

	return
}

// FindNamespacesByAnnotation find namespaces by annotation map
func (r *AutoObjectCreationReconciler) FindNamespacesByAnnotation(annotations map[string]string) ([]corev1.Namespace, error) {
	namespaces, err := r.FindNamespaces()
	if err != nil {
		return nil, err
	}

	var foundedNamespaces []corev1.Namespace

	for _, namespace := range namespaces {
		var found = true
		for annotation := range annotations {
			if _, found = namespace.Annotations[annotation]; found {
				break
			}
		}

		if found {
			foundedNamespaces = append(foundedNamespaces, namespace)
		}
	}

	return foundedNamespaces, nil
}
