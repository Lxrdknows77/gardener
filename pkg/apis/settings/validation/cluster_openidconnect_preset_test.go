// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package validation_test

import (
	"github.com/gardener/gardener/pkg/apis/settings"
	settings_validation "github.com/gardener/gardener/pkg/apis/settings/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

type clusterOpenIDConnectPresetProvider struct {
	new *settings.ClusterOpenIDConnectPreset
}

func (o *clusterOpenIDConnectPresetProvider) providerFunc() field.ErrorList {
	return settings_validation.ValidateClusterOpenIDConnectPreset(o.new)
}
func (o *clusterOpenIDConnectPresetProvider) preset() settings.Preset { return o.new }

type clusterOpenIDConnectPresetUpdateProvider struct {
	new *settings.ClusterOpenIDConnectPreset
	old *settings.ClusterOpenIDConnectPreset
}

func (o *clusterOpenIDConnectPresetUpdateProvider) providerFunc() field.ErrorList {
	return settings_validation.ValidateClusterOpenIDConnectPresetUpdate(o.new, o.old)
}
func (o *clusterOpenIDConnectPresetUpdateProvider) preset() settings.Preset { return o.new }

var _ = Describe("ClusterOpenIDConnectPreset", func() {

	var preset *settings.ClusterOpenIDConnectPreset

	BeforeEach(func() {
		preset = &settings.ClusterOpenIDConnectPreset{
			ObjectMeta: metav1.ObjectMeta{
				Name: "test",
			},
			OpenIDConnectPresetSpec: settings.OpenIDConnectPresetSpec{
				Weight: 1,
				Server: settings.KubeAPIServerOpenIDConnect{
					IssuerURL: "https://foo.bar",
					ClientID:  "client-baz",
				},
			},
		}

	})

	Describe("#ValidateClusterOpenIDConnectPreset", func() {
		provider := &clusterOpenIDConnectPresetProvider{}

		BeforeEach(func() {
			provider.new = preset.DeepCopy()
		})

		It("should forbid empty object", func() {

			provider.new.ObjectMeta.Name = ""
			provider.new.ObjectMeta.Namespace = ""
			provider.new.OpenIDConnectPresetSpec = settings.OpenIDConnectPresetSpec{}

			errorList := provider.providerFunc()

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("metadata.name"),
			})), PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeInvalid),
				"Field": Equal("weight"),
			})), PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("server.issuerURL"),
			})), PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":  Equal(field.ErrorTypeRequired),
				"Field": Equal("server.clientID"),
			})),
			))
		})

		validationAssertions(provider)
	})

	Describe("#ValidateClusterOpenIDConnectPresetUpdate", func() {
		provider := &clusterOpenIDConnectPresetUpdateProvider{}

		BeforeEach(func() {
			provider.old = preset.DeepCopy()
			provider.old.ObjectMeta.ResourceVersion = "2"

			provider.new = preset.DeepCopy()
			provider.new.ObjectMeta.ResourceVersion = "2"
		})

		It("should forbid update with mutation of objectmeta fields", func() {

			provider.new.ObjectMeta.Name = "changed-name"
			provider.new.ObjectMeta.ResourceVersion = ""

			errorList := provider.providerFunc()

			Expect(errorList).To(ConsistOf(PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":   Equal(field.ErrorTypeInvalid),
				"Field":  Equal("metadata.name"),
				"Detail": Equal("field is immutable"),
			})), PointTo(MatchFields(IgnoreExtras, Fields{
				"Type":   Equal(field.ErrorTypeInvalid),
				"Field":  Equal("metadata.resourceVersion"),
				"Detail": Equal("must be specified for an update"),
			})),
			))
		})

		validationAssertions(provider)
	})
})
