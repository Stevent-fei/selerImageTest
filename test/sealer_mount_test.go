// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"fmt"

	"blog/test/testhelper/settings"
	"github.com/sealerio/sealer/test/suites/build"
	"github.com/sealerio/sealer/test/suites/image"
	"github.com/sealerio/sealer/test/testhelper"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("sealer mount and umount", func() {
	Context("test mount and unmount a single image", func() {
		BeforeEach(func() {
			image.DoImageOps("pull", settings.TestImageName)
			testhelper.CheckBeTrue(build.CheckIsImageExist(settings.TestImageName))
		})

		It("start to mount and umount", func() {

			By("start to mount a image")
			mountCmd := fmt.Sprintf("%s alpha mount %s -d", settings.DefaultSealerBin, settings.TestImageName)
			testhelper.RunCmdAndCheckResult(mountCmd, 0)

			By("start to umount a image")
			umountCmd := fmt.Sprintf("%s alpha umount %s -d", settings.DefaultSealerBin, settings.TestImageName)
			testhelper.RunCmdAndCheckResult(umountCmd, 0)

		})
	})

	Context("test mount and unmount multiple images", func() {
		BeforeEach(func() {
			image.DoImageOps("pull", settings.TestImageName)
			image.DoImageOps("pull", settings.TestK8s120)
			image.DoImageOps("pull", settings.TestK8s118)
			testhelper.CheckBeTrue(build.CheckIsImageExist(settings.TestImageName))
			testhelper.CheckBeTrue(build.CheckIsImageExist(settings.TestK8s120))
			testhelper.CheckBeTrue(build.CheckIsImageExist(settings.TestK8s118))
		})
		It("start to mount and umount", func() {

			By("start to mount multiple images")
			mountCmd := fmt.Sprintf("%s alpha mount %s %s %s -d", settings.DefaultSealerBin, settings.TestImageName, settings.TestK8s120, settings.TestK8s118)
			testhelper.RunCmdAndCheckResult(mountCmd, 0)

			By("start to umount multiple images")
			umountCmd := fmt.Sprintf("%s alpha umount %s %s %s -d", settings.DefaultSealerBin, settings.TestImageName, settings.TestK8s120, settings.TestK8s118)
			testhelper.RunCmdAndCheckResult(umountCmd, 0)
		})
	})
})
