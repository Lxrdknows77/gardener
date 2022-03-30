// Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
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

package controller

import (
	"time"

	testutils "github.com/gardener/gardener/landscaper/common/test-utils"
	landscaperutils "github.com/gardener/gardener/landscaper/common/utils"
	"github.com/gardener/gardener/landscaper/pkg/controlplane/apis/exports"
	"github.com/gardener/gardener/landscaper/pkg/controlplane/apis/imports"
	"github.com/gardener/gardener/pkg/logger"
	"github.com/gardener/gardener/pkg/utils/test"
	. "github.com/onsi/ginkgo/v2"
	"github.com/sirupsen/logrus"

	. "github.com/onsi/gomega"
)

var _ = Describe("#CheckForExpiringCertificates", func() {
	var (
		defaultValidity         = 30 * time.Second
		ca                      = testutils.GenerateCACertificateWithValidity("gardener.cloud:system:apiserver", defaultValidity)
		caCrt                   = string(ca.CertificatePEM)
		apiServerTLSServingCert = testutils.GenerateTLSServingCertificateWithValidity(&ca, defaultValidity)
		tlsServingCertString    = string(apiServerTLSServingCert.CertificatePEM)
		testOperation           operation
	)

	BeforeEach(func() {
		testOperation = operation{
			log: logrus.NewEntry(logger.NewNopLogger()),
			imports: &imports.Imports{
				GardenerAPIServer: imports.GardenerAPIServer{
					ComponentConfiguration: imports.APIServerComponentConfiguration{
						CA: &imports.CA{
							Crt: &caCrt,
						},
						TLS: &imports.TLSServer{
							Crt: &tlsServingCertString,
						},
					},
				},
				GardenerControllerManager: &imports.GardenerControllerManager{
					ComponentConfiguration: &imports.ControllerManagerComponentConfiguration{
						TLS: &imports.TLSServer{
							Crt: &tlsServingCertString,
						},
					},
				},
				GardenerAdmissionController: &imports.GardenerAdmissionController{
					Enabled: true,
					ComponentConfiguration: &imports.AdmissionControllerComponentConfiguration{
						CA: &imports.CA{
							Crt: &caCrt,
						},
						TLS: &imports.TLSServer{
							Crt: &tlsServingCertString,
						},
					},
				},
			},
			// initialized when creating operation
			exports: exports.Exports{
				GardenerAPIServerCA:                   exports.Certificate{},
				GardenerAPIServerTLSServing:           exports.Certificate{},
				GardenerAdmissionControllerCA:         &exports.Certificate{},
				GardenerAdmissionControllerTLSServing: &exports.Certificate{},
				GardenerControllerManagerTLSServing:   exports.Certificate{},
			},
		}
	})

	It("should not require certificate renewal - within validity threshold", func() {
		nowFunc := func() time.Time {
			return time.Now()
		}

		defer test.WithVar(&landscaperutils.NowFunc, nowFunc)()

		Expect(testOperation.CheckForExpiringCertificates(nil)).ToNot(HaveOccurred())
		Expect(testOperation.imports.GardenerAPIServer.ComponentConfiguration.CA.Crt).ToNot(BeNil())
		Expect(testOperation.imports.GardenerAPIServer.ComponentConfiguration.TLS.Crt).ToNot(BeNil())
		Expect(testOperation.exports.GardenerAPIServerCA.Rotated).To(Equal(false))
		Expect(testOperation.exports.GardenerAPIServerTLSServing.Rotated).To(Equal(false))

		Expect(testOperation.imports.GardenerAdmissionController.ComponentConfiguration.CA.Crt).ToNot(BeNil())
		Expect(testOperation.imports.GardenerAdmissionController.ComponentConfiguration.TLS.Crt).ToNot(BeNil())
		Expect(testOperation.exports.GardenerAdmissionControllerCA.Rotated).To(Equal(false))
		Expect(testOperation.exports.GardenerAdmissionControllerTLSServing.Rotated).To(Equal(false))

		Expect(testOperation.imports.GardenerControllerManager.ComponentConfiguration.TLS.Crt).ToNot(BeNil())
		Expect(testOperation.exports.GardenerControllerManagerTLSServing.Rotated).To(Equal(false))
	})

	It("should indicate rotation", func() {
		jumpToFuture := 25 * time.Second // within 80% validity threshold of 30 seconds expiration
		nowFunc := func() time.Time {
			return time.Now().Add(jumpToFuture)
		}

		defer test.WithVar(&landscaperutils.NowFunc, nowFunc)()

		Expect(testOperation.CheckForExpiringCertificates(nil)).ToNot(HaveOccurred())
		Expect(testOperation.imports.GardenerAPIServer.ComponentConfiguration.CA.Crt).To(BeNil())
		Expect(testOperation.imports.GardenerAPIServer.ComponentConfiguration.TLS.Crt).To(BeNil())
		Expect(testOperation.exports.GardenerAPIServerCA.Rotated).To(Equal(true))
		Expect(testOperation.exports.GardenerAPIServerTLSServing.Rotated).To(Equal(true))

		Expect(testOperation.imports.GardenerAdmissionController.ComponentConfiguration.CA.Crt).To(BeNil())
		Expect(testOperation.imports.GardenerAdmissionController.ComponentConfiguration.TLS.Crt).To(BeNil())
		Expect(testOperation.exports.GardenerAdmissionControllerCA.Rotated).To(Equal(true))
		Expect(testOperation.exports.GardenerAdmissionControllerTLSServing.Rotated).To(Equal(true))

		Expect(testOperation.imports.GardenerControllerManager.ComponentConfiguration.TLS.Crt).To(BeNil())
		Expect(testOperation.exports.GardenerControllerManagerTLSServing.Rotated).To(Equal(true))
	})
})