package settings

import (
	"os"
	"time"
)

const (
	FileMode0755 = 0755
	FileMode0644 = 0644
)
const (
	BAREMETAL    = "BAREMETAL"
	AliCloud     = "ALI_CLOUD"
	DefaultImage = "docker.io/sealerio/kubernetes:v1-22-15-sealerio-2"
	TestK8s120   = "docker.io/sealerio/kubernetes:v1.20.4-sealerio-2"
	TestK8s118   = "docker.io/sealerio/kubernetes:v1.18.3-sealerio-2"
)

var (
	DefaultPollingInterval time.Duration
	MaxWaiteTime           time.Duration
	DefaultWaiteTime       time.Duration
	DefaultSealerBin       = ""
	DefaultTestEnvDir      = ""
	CustomImageName        = os.Getenv("IMAGE_NAME")
	LoadPath               = ""

	TestImageName      = "ack-agility-registry.cn-shanghai.cr.aliyuncs.com/ecp_builder/ackdistro:test"
	CustomCalicoEnv    = "Network=calico"
	CustomhybridnetEnv = "Network=hybridnet"
	CalicoEnv          = []string{"Network=calico"}
	HybridnetEnv       = []string{"Network=hybridnet"}
	aaa                = "aaa"
)
