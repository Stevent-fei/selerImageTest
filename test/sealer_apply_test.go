package test

import (
	"blog/test/testhelper"
	. "github.com/onsi/ginkgo"

	"blog/test/suites/apply"
	"blog/test/testhelper/settings"
)

var _ = Describe("sealer apply", func() {
	Context("start apply", func() {
		rawClusterFilePath := apply.GetRawClusterFilePath()
		rawCluster := apply.LoadClusterFileFromDisk(rawClusterFilePath)
		rawCluster.Spec.Image = settings.TestImageName
		BeforeEach(func() {
			if rawCluster.Spec.Image != settings.TestImageName {
				//rawCluster imageName updated to customImageName
				rawCluster.Spec.Image = settings.TestImageName
				apply.MarshalClusterToFile(rawClusterFilePath, rawCluster)
			}
		})

		Context("check regular scenario that provider is bare metal, executes machine is not master0", func() {
			var tempFile string
			BeforeEach(func() {
				tempFile = testhelper.CreateTempFile()
			})

			AfterEach(func() {
				testhelper.RemoveTempFile(tempFile)
				testhelper.DeleteFileLocally(settings.GetClusterWorkClusterfile(rawCluster.Name))
			})
			It("init, scale up, scale down, clean up", func() {
				By("start to prepare infra")
				cluster := apply.LoadClusterFileFromDisk(rawClusterFilePath)
				cluster.Spec.Provider = settings.AliCloud
				usedCluster := apply.ChangeMasterOrderAndSave(cluster, tempFile)
				defer apply.CleanUpAliCloudInfra(usedCluster)
				sshClient := testhelper.NewSSHClientByCluster(usedCluster)
				testhelper.CheckFuncBeTrue(func() bool {
					err := sshClient.SSH.Copy(sshClient.RemoteHostIP, settings.DefaultSealerBin, settings.DefaultSealerBin)
					return err == nil
				}, settings.MaxWaiteTime)

				By("start to init cluster")
				apply.SendAndApplyCluster(sshClient, tempFile)
				apply.CheckNodeNumWithSSH(sshClient, 4)

				By("Wait for the cluster to be ready", func() {
					apply.WaitAllNodeRunningBySSH(sshClient.SSH, sshClient.RemoteHostIP)
				})

				By("start to delete cluster")
				err := sshClient.SSH.CmdAsync(sshClient.RemoteHostIP, apply.SealerDeleteCmd(tempFile))
				testhelper.CheckErr(err)
			})
		})
	})
})
