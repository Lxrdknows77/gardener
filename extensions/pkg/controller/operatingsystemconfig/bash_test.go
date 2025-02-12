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

package operatingsystemconfig_test

import (
	"context"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"

	. "github.com/gardener/gardener/extensions/pkg/controller/operatingsystemconfig"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	. "github.com/gardener/gardener/pkg/utils/test/matchers"
)

var _ = Describe("Bash", func() {
	Describe("#FilesToDiskScript", func() {
		var (
			ctx        = context.Background()
			fakeClient client.Client
			namespace  = "namespace"
		)

		BeforeEach(func() {
			fakeClient = fakeclient.NewClientBuilder().Build()
		})

		It("should fail when a referenced secret cannot be read", func() {
			files := []extensionsv1alpha1.File{{
				Content: extensionsv1alpha1.FileContent{
					SecretRef: &extensionsv1alpha1.FileContentSecretRef{
						Name: "foo",
					},
				},
			}}

			script, err := FilesToDiskScript(ctx, fakeClient, namespace, files)
			Expect(err).To(BeNotFoundError())
			Expect(script).To(BeEmpty())
		})

		It("should generate the expected output", func() {
			var (
				folder1 = "/foo"
				file1   = folder1 + "/bar.txt"

				folder2 = "/bar"
				file2   = folder2 + "/baz"

				folder3 = "/baz"
				file3   = folder3 + "/foo"

				folder4 = "/foobar"
				file4   = folder4 + "/baz"
			)

			Expect(fakeClient.Create(ctx, &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: namespace,
				},
				Data: map[string][]byte{"bar": []byte("bar-content")},
			})).To(Succeed())

			files := []extensionsv1alpha1.File{
				{
					Path: file1,
					Content: extensionsv1alpha1.FileContent{
						SecretRef: &extensionsv1alpha1.FileContentSecretRef{
							Name:    "foo",
							DataKey: "bar",
						},
					},
				},
				{
					Path: file2,
					Content: extensionsv1alpha1.FileContent{
						Inline: &extensionsv1alpha1.FileContentInline{
							Encoding: "",
							Data:     "plain-text",
						},
					},
				},
				{
					Path: file3,
					Content: extensionsv1alpha1.FileContent{
						Inline: &extensionsv1alpha1.FileContentInline{
							Encoding: "b64",
							Data:     "YmFzZTY0",
						},
					},
				},
				{
					Path: file4,
					Content: extensionsv1alpha1.FileContent{
						Inline: &extensionsv1alpha1.FileContentInline{
							Encoding: "",
							Data:     "transmit-unencoded",
						},
						TransmitUnencoded: pointer.Bool(true),
					},
				},
			}

			By("Ensure the function generated the expected bash script")
			script, err := FilesToDiskScript(ctx, fakeClient, namespace, files)
			Expect(err).NotTo(HaveOccurred())
			Expect(script).To(Equal(`
mkdir -p "` + folder1 + `"

cat << EOF | base64 -d > "` + file1 + `"
YmFyLWNvbnRlbnQ=
EOF
mkdir -p "` + folder2 + `"

cat << EOF | base64 -d > "` + file2 + `"
cGxhaW4tdGV4dA==
EOF
mkdir -p "` + folder3 + `"

cat << EOF | base64 -d > "` + file3 + `"
YmFzZTY0
EOF
mkdir -p "` + folder4 + `"

cat << EOF > "` + file4 + `"
transmit-unencoded
EOF`))

			By("Ensure that the bash script can be executed and performs the desired operations")
			tempDir, err := os.MkdirTemp("", "tempdir")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(tempDir)

			script = strings.ReplaceAll(script, `"`+folder1, `"`+tempDir+folder1)
			script = strings.ReplaceAll(script, `"`+folder2, `"`+tempDir+folder2)
			script = strings.ReplaceAll(script, `"`+folder3, `"`+tempDir+folder3)
			script = strings.ReplaceAll(script, `"`+folder4, `"`+tempDir+folder4)

			runScriptAndCheckFiles(script,
				tempDir+file1,
				tempDir+file2,
				tempDir+file3,
				tempDir+file4,
			)
		})
	})

	Describe("#UnitsToDiskScript", func() {
		It("should generate the expected output", func() {
			var (
				unit1        = "unit1"
				unit1DropIn1 = "dropin1"
				unit1DropIn2 = "dropin2"

				unit2 = "unit2"

				units = []extensionsv1alpha1.Unit{
					{
						Name: unit1,
						DropIns: []extensionsv1alpha1.DropIn{
							{
								Name:    unit1DropIn1,
								Content: "dropdrop",
							},
							{
								Name:    unit1DropIn2,
								Content: "dropeldidrop",
							},
						},
					},
					{
						Name:    unit2,
						Content: pointer.String("content2"),
					},
				}
			)

			By("Ensure the function generated the expected bash script")
			script := UnitsToDiskScript(units)
			Expect(script).To(Equal(`
mkdir -p "/etc/systemd/system/` + unit1 + `.d"

cat << EOF | base64 -d > "/etc/systemd/system/` + unit1 + `.d/` + unit1DropIn1 + `"
ZHJvcGRyb3A=
EOF

cat << EOF | base64 -d > "/etc/systemd/system/` + unit1 + `.d/` + unit1DropIn2 + `"
ZHJvcGVsZGlkcm9w
EOF

cat << EOF | base64 -d > "/etc/systemd/system/` + unit2 + `"
Y29udGVudDI=
EOF`))

			By("Ensure that the bash script can be executed and performs the desired operations")
			tempDir, err := os.MkdirTemp("", "tempdir")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(tempDir)

			script = strings.ReplaceAll(script, "/etc/systemd/system/", tempDir+"/etc/systemd/system/")

			runScriptAndCheckFiles(script,
				tempDir+"/etc/systemd/system/"+unit2,
				tempDir+"/etc/systemd/system/"+unit1+".d/"+unit1DropIn1,
				tempDir+"/etc/systemd/system/"+unit1+".d/"+unit1DropIn2,
			)
		})
	})
})

func runScriptAndCheckFiles(script string, filePaths ...string) {
	ExpectWithOffset(1, exec.Command("bash", "-c", script).Run()).To(Succeed())

	for _, filePath := range filePaths {
		fileInfo, err := os.Stat(filePath)
		ExpectWithOffset(1, err).NotTo(HaveOccurred(), "file at path "+filePath)
		ExpectWithOffset(1, fileInfo.Mode().IsRegular()).To(BeTrue(), "file at path "+filePath)
		ExpectWithOffset(1, fileInfo.Mode().Perm()).To(Equal(fs.FileMode(0644)), "file at path "+filePath)
	}
}
