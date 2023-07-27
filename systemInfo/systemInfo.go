package systemInfo

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
)

type AllInfo struct {
	CpuName           string  `json:"cpuname"`
	StorageUnit       string  `json:"storageunit"`
	DiskAvailable     uint64  `json:"diskavailable"`
	DiskUsed          uint64  `json:"DiskUsed"`
	DiskAvailablePerc float64 `json:"DiskAvailablePerc"`
	Hostname          string  `json:"Hostname"`
}

func NewAllInfo() AllInfo {
	return AllInfo{
		"",
		"GB",
		0,
		0,
		0.0,
		"",
	}
}

func (a *AllInfo) Initialize(cpuInfo cpu.InfoStat, dUsageInfo disk.UsageStat, hostInfo host.InfoStat) {

	a.CpuName = cpuInfo.ModelName
	a.DiskAvailable = (uint64)(dUsageInfo.Free / 1000000000)
	a.DiskUsed = (uint64)(dUsageInfo.Used / 1000000000)
	a.DiskAvailablePerc = dUsageInfo.UsedPercent
	a.Hostname = hostInfo.Hostname
}

func CollectStats() AllInfo {

	allinfo := NewAllInfo()
	info, err := cpu.Info()
	if err != nil {
		fmt.Println("Error getting cpu info - ", err)
	}
	diskUsageInfo, err := disk.Usage("/")
	if err != nil {
		fmt.Println("Error getting disk info - ", err)
	}
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("Error getting host info - ", err)
	}

	allinfo.Initialize(info[0], *diskUsageInfo, *hostInfo)

	return allinfo

}
