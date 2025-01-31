// Copyright 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package reference

import (
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	operatorv1alpha1 "github.com/gardener/gardener/pkg/apis/operator/v1alpha1"
	"github.com/gardener/gardener/pkg/controller/reference"
)

// AddToManager adds the garden-reference controller to the given manager.
func AddToManager(mgr manager.Manager, gardenNamespace string) error {
	return (&reference.Reconciler{
		ConcurrentSyncs:             pointer.Int(1),
		NewObjectFunc:               func() client.Object { return &operatorv1alpha1.Garden{} },
		NewObjectListFunc:           func() client.ObjectList { return &operatorv1alpha1.GardenList{} },
		GetNamespace:                func(client.Object) string { return gardenNamespace },
		GetReferencedSecretNames:    getReferencedSecretNames,
		GetReferencedConfigMapNames: getReferencedConfigMapNames,
		ReferenceChangedPredicate:   Predicate,
	}).AddToManager(mgr)
}

// Predicate is a predicate function for checking whether a reference changed in the Garden specification.
func Predicate(oldObj, newObj client.Object) bool {
	newGarden, ok := newObj.(*operatorv1alpha1.Garden)
	if !ok {
		return false
	}

	oldGarden, ok := oldObj.(*operatorv1alpha1.Garden)
	if !ok {
		return false
	}

	return kubeAPIServerAuditPolicyConfigMapChanged(oldGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer, newGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer) ||
		gardenerAPIServerAuditPolicyConfigMapChanged(oldGarden.Spec.VirtualCluster.Gardener.APIServer, newGarden.Spec.VirtualCluster.Gardener.APIServer) ||
		etcdBackupSecretChanged(oldGarden.Spec.VirtualCluster.ETCD, newGarden.Spec.VirtualCluster.ETCD) ||
		authenticationWebhookSecretChanged(oldGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer, newGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer) ||
		sniSecretChanged(oldGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer, newGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer) ||
		kubeAPIServerAuditWebhookSecretChanged(oldGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer, newGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer) ||
		gardenerAPIServerAuditWebhookSecretChanged(oldGarden.Spec.VirtualCluster.Gardener.APIServer, newGarden.Spec.VirtualCluster.Gardener.APIServer) ||
		kubeAPIServerAdmissionPluginSecretChanged(oldGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer, newGarden.Spec.VirtualCluster.Kubernetes.KubeAPIServer) ||
		gardenerAPIServerAdmissionPluginSecretChanged(oldGarden.Spec.VirtualCluster.Gardener.APIServer, newGarden.Spec.VirtualCluster.Gardener.APIServer)
}

func kubeAPIServerAuditPolicyConfigMapChanged(oldKubeAPIServer, newKubeAPIServer *operatorv1alpha1.KubeAPIServerConfig) bool {
	var oldConfigMap, newConfigMap string

	if oldKubeAPIServer != nil && oldKubeAPIServer.AuditConfig != nil && oldKubeAPIServer.AuditConfig.AuditPolicy != nil && oldKubeAPIServer.AuditConfig.AuditPolicy.ConfigMapRef != nil {
		oldConfigMap = oldKubeAPIServer.AuditConfig.AuditPolicy.ConfigMapRef.Name
	}
	if newKubeAPIServer != nil && newKubeAPIServer.AuditConfig != nil && newKubeAPIServer.AuditConfig.AuditPolicy != nil && newKubeAPIServer.AuditConfig.AuditPolicy.ConfigMapRef != nil {
		newConfigMap = newKubeAPIServer.AuditConfig.AuditPolicy.ConfigMapRef.Name
	}

	return oldConfigMap != newConfigMap
}

func gardenerAPIServerAuditPolicyConfigMapChanged(oldGardenerAPIServer, newGardenerAPIServer *operatorv1alpha1.GardenerAPIServerConfig) bool {
	var oldConfigMap, newConfigMap string

	if oldGardenerAPIServer != nil && oldGardenerAPIServer.AuditConfig != nil && oldGardenerAPIServer.AuditConfig.AuditPolicy != nil && oldGardenerAPIServer.AuditConfig.AuditPolicy.ConfigMapRef != nil {
		oldConfigMap = oldGardenerAPIServer.AuditConfig.AuditPolicy.ConfigMapRef.Name
	}
	if newGardenerAPIServer != nil && newGardenerAPIServer.AuditConfig != nil && newGardenerAPIServer.AuditConfig.AuditPolicy != nil && newGardenerAPIServer.AuditConfig.AuditPolicy.ConfigMapRef != nil {
		newConfigMap = newGardenerAPIServer.AuditConfig.AuditPolicy.ConfigMapRef.Name
	}

	return oldConfigMap != newConfigMap
}

func etcdBackupSecretChanged(oldETCD, newETCD *operatorv1alpha1.ETCD) bool {
	var oldSecret, newSecret string

	if oldETCD != nil && oldETCD.Main != nil && oldETCD.Main.Backup != nil {
		oldSecret = oldETCD.Main.Backup.SecretRef.Name
	}

	if newETCD != nil && newETCD.Main != nil && newETCD.Main.Backup != nil {
		newSecret = newETCD.Main.Backup.SecretRef.Name
	}

	return oldSecret != newSecret
}

func authenticationWebhookSecretChanged(oldKubeAPIServer, newKubeAPIServer *operatorv1alpha1.KubeAPIServerConfig) bool {
	var oldSecret, newSecret string

	if oldKubeAPIServer != nil && oldKubeAPIServer.Authentication != nil && oldKubeAPIServer.Authentication.Webhook != nil {
		oldSecret = oldKubeAPIServer.Authentication.Webhook.KubeconfigSecretName
	}

	if newKubeAPIServer != nil && newKubeAPIServer.Authentication != nil && newKubeAPIServer.Authentication.Webhook != nil {
		newSecret = newKubeAPIServer.Authentication.Webhook.KubeconfigSecretName
	}

	return oldSecret != newSecret
}

func sniSecretChanged(oldKubeAPIServer, newKubeAPIServer *operatorv1alpha1.KubeAPIServerConfig) bool {
	var oldSecret, newSecret string

	if oldKubeAPIServer != nil && oldKubeAPIServer.SNI != nil {
		oldSecret = oldKubeAPIServer.SNI.SecretName
	}

	if newKubeAPIServer != nil && newKubeAPIServer.SNI != nil {
		newSecret = newKubeAPIServer.SNI.SecretName
	}

	return oldSecret != newSecret
}

func kubeAPIServerAuditWebhookSecretChanged(oldKubeAPIServer, newKubeAPIServer *operatorv1alpha1.KubeAPIServerConfig) bool {
	var oldSecret, newSecret string

	if oldKubeAPIServer != nil && oldKubeAPIServer.AuditWebhook != nil {
		oldSecret = oldKubeAPIServer.AuditWebhook.KubeconfigSecretName
	}
	if newKubeAPIServer != nil && newKubeAPIServer.AuditWebhook != nil {
		newSecret = newKubeAPIServer.AuditWebhook.KubeconfigSecretName
	}

	return oldSecret != newSecret
}

func gardenerAPIServerAuditWebhookSecretChanged(oldGardenerAPIServer, newGardenerAPIServer *operatorv1alpha1.GardenerAPIServerConfig) bool {
	var oldSecret, newSecret string

	if oldGardenerAPIServer != nil && oldGardenerAPIServer.AuditWebhook != nil {
		oldSecret = oldGardenerAPIServer.AuditWebhook.KubeconfigSecretName
	}
	if newGardenerAPIServer != nil && newGardenerAPIServer.AuditWebhook != nil {
		newSecret = newGardenerAPIServer.AuditWebhook.KubeconfigSecretName
	}

	return oldSecret != newSecret
}

func kubeAPIServerAdmissionPluginSecretChanged(oldKubeAPIServer, newKubeAPIServer *operatorv1alpha1.KubeAPIServerConfig) bool {
	oldSecrets, newSecrets := sets.Set[string]{}, sets.Set[string]{}

	if oldKubeAPIServer != nil {
		for _, plugin := range oldKubeAPIServer.AdmissionPlugins {
			if plugin.KubeconfigSecretName != nil {
				oldSecrets.Insert(*plugin.KubeconfigSecretName)
			}
		}
	}
	if newKubeAPIServer != nil {
		for _, plugin := range newKubeAPIServer.AdmissionPlugins {
			if plugin.KubeconfigSecretName != nil {
				newSecrets.Insert(*plugin.KubeconfigSecretName)
			}
		}
	}

	return !oldSecrets.Equal(newSecrets)
}

func gardenerAPIServerAdmissionPluginSecretChanged(oldGardenerAPIServer, newGardenerAPIServer *operatorv1alpha1.GardenerAPIServerConfig) bool {
	oldSecrets, newSecrets := sets.Set[string]{}, sets.Set[string]{}

	if oldGardenerAPIServer != nil {
		for _, plugin := range oldGardenerAPIServer.AdmissionPlugins {
			if plugin.KubeconfigSecretName != nil {
				oldSecrets.Insert(*plugin.KubeconfigSecretName)
			}
		}
	}
	if newGardenerAPIServer != nil {
		for _, plugin := range newGardenerAPIServer.AdmissionPlugins {
			if plugin.KubeconfigSecretName != nil {
				newSecrets.Insert(*plugin.KubeconfigSecretName)
			}
		}
	}

	return !oldSecrets.Equal(newSecrets)
}

func getReferencedSecretNames(obj client.Object) []string {
	garden, ok := obj.(*operatorv1alpha1.Garden)
	if !ok {
		return nil
	}

	var (
		virtualCluster = garden.Spec.VirtualCluster
		out            []string
	)

	if virtualCluster.ETCD != nil && virtualCluster.ETCD.Main != nil && virtualCluster.ETCD.Main.Backup != nil {
		out = append(out, virtualCluster.ETCD.Main.Backup.SecretRef.Name)
	}

	if virtualCluster.Kubernetes.KubeAPIServer != nil {
		for _, plugin := range virtualCluster.Kubernetes.KubeAPIServer.AdmissionPlugins {
			if plugin.KubeconfigSecretName != nil {
				out = append(out, *plugin.KubeconfigSecretName)
			}
		}
	}

	if virtualCluster.Gardener.APIServer != nil {
		for _, plugin := range virtualCluster.Gardener.APIServer.AdmissionPlugins {
			if plugin.KubeconfigSecretName != nil {
				out = append(out, *plugin.KubeconfigSecretName)
			}
		}
	}

	if virtualCluster.Kubernetes.KubeAPIServer != nil && virtualCluster.Kubernetes.KubeAPIServer.Authentication != nil && virtualCluster.Kubernetes.KubeAPIServer.Authentication.Webhook != nil {
		out = append(out, virtualCluster.Kubernetes.KubeAPIServer.Authentication.Webhook.KubeconfigSecretName)
	}

	if virtualCluster.Kubernetes.KubeAPIServer != nil && virtualCluster.Kubernetes.KubeAPIServer.SNI != nil {
		out = append(out, virtualCluster.Kubernetes.KubeAPIServer.SNI.SecretName)
	}

	if virtualCluster.Kubernetes.KubeAPIServer != nil && virtualCluster.Kubernetes.KubeAPIServer.AuditWebhook != nil {
		out = append(out, virtualCluster.Kubernetes.KubeAPIServer.AuditWebhook.KubeconfigSecretName)
	}

	if virtualCluster.Gardener.APIServer != nil && virtualCluster.Gardener.APIServer.AuditWebhook != nil {
		out = append(out, virtualCluster.Gardener.APIServer.AuditWebhook.KubeconfigSecretName)
	}

	return out
}

func getReferencedConfigMapNames(obj client.Object) []string {
	garden, ok := obj.(*operatorv1alpha1.Garden)
	if !ok {
		return nil
	}

	var (
		virtualCluster = garden.Spec.VirtualCluster
		out            []string
	)

	if virtualCluster.Kubernetes.KubeAPIServer != nil && virtualCluster.Kubernetes.KubeAPIServer.AuditConfig != nil && virtualCluster.Kubernetes.KubeAPIServer.AuditConfig.AuditPolicy != nil && virtualCluster.Kubernetes.KubeAPIServer.AuditConfig.AuditPolicy.ConfigMapRef != nil {
		out = append(out, virtualCluster.Kubernetes.KubeAPIServer.AuditConfig.AuditPolicy.ConfigMapRef.Name)
	}

	if virtualCluster.Gardener.APIServer != nil && virtualCluster.Gardener.APIServer.AuditConfig != nil && virtualCluster.Gardener.APIServer.AuditConfig.AuditPolicy != nil && virtualCluster.Gardener.APIServer.AuditConfig.AuditPolicy.ConfigMapRef != nil {
		out = append(out, virtualCluster.Gardener.APIServer.AuditConfig.AuditPolicy.ConfigMapRef.Name)
	}

	return out
}
