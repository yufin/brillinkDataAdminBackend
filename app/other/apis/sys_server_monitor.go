package apis

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	_ "go-admin/common/response/antd"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	//Version           string
	expectDiskFsTypes = []string{
		"apfs", "ext4", "ext3", "ext2", "f2fs", "reiserfs", "jfs", "btrfs",
		"fuseblk", "zfs", "simfs", "ntfs", "fat32", "exfat", "xfs", "fuse.rclone",
	}
	//excludeNetInterfaces = []string{
	//	"lo", "tun", "docker", "veth", "br-", "vmbr", "vnet", "kube",
	//}
	getMacDiskNo = regexp.MustCompile(`\/dev\/disk(\d)s.*`)
)

var (
	// netInSpeed, netOutSpeed, netInTransfer, netOutTransfer, lastUpdateNetStats uint64
	cachedBootTime time.Time
)

type ServerMonitor struct {
	api.Api
}

// 获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

// ServerInfo 获取系统信息
// @Summary 系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {object} antd.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/server-monitor [get]
// @Security Bearer
func (e ServerMonitor) ServerInfo(c *gin.Context) {
	e.Context = c

	sysInfo, err := host.Info()
	cpuType := "Physical"
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = sysInfo.OS
	osDic["arch"] = sysInfo.KernelArch
	if sysInfo.VirtualizationSystem != "" {
		cpuType = "Vrtual"
	}
	osDic["cpuType"] = cpuType
	cpuModelCount := make(map[string]int)
	ci, _ := cpu.Info()
	for i := 0; i < len(ci); i++ {
		cpuModelCount[ci[i].ModelName]++
	}
	var cpus []string
	for model, count := range cpuModelCount {
		cpus = append(cpus, fmt.Sprintf("%s %d %s Core", model, count, cpuType))
	}
	mv, _ := mem.VirtualMemory()
	diskTotal, diskUsed := getDiskTotalAndUsed()

	var swapMemTotal uint64
	if runtime.GOOS == "windows" {
		ms, _ := mem.SwapMemory()
		swapMemTotal = ms.Total
	} else {
		swapMemTotal = mv.SwapTotal
	}

	if cachedBootTime.IsZero() {
		cachedBootTime = time.Unix(int64(sysInfo.BootTime), 0)
	}
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = sysInfo.KernelVersion
	osDic["numGoroutine"] = runtime.NumGoroutine()
	osDic["ip"] = pkg.GetLocaHonst()
	osDic["projectDir"] = pkg.GetCurrentPath()
	osDic["hostName"] = sysInfo.Hostname
	osDic["swapMemTotal"] = swapMemTotal
	osDic["bootTime"] = sysInfo.BootTime
	osDic["time"] = time.Now().Format("2006-01-02 15:04:05")

	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = diskTotal / GB
	diskDic["free"] = diskUsed / GB

	mem, _ := mem.VirtualMemory()
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = memTotalMB
	memDic["used"] = memUsedMB
	memDic["free"] = memFreeMB
	memDic["usage"] = memUsedPercent

	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuInfo"], _ = cpu.Info()
	percent, _ := cpu.Percent(0, false)
	cpuDic["Percent"] = pkg.Round(percent[0], 2)
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	//服务器磁盘信息
	disklist := make([]disk.UsageStat, 0)
	//所有分区
	diskInfo, err := disk.Partitions(true)
	if err == nil {
		for _, p := range diskInfo {
			diskDetail, err := disk.Usage(p.Mountpoint)
			if err == nil {
				diskDetail.UsedPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", diskDetail.UsedPercent), 64)
				diskDetail.Total = diskDetail.Total / 1024 / 1024
				diskDetail.Used = diskDetail.Used / 1024 / 1024
				diskDetail.Free = diskDetail.Free / 1024 / 1024
				disklist = append(disklist, *diskDetail)
			}
		}
	}

	e.Custom(gin.H{
		"code":     200,
		"os":       osDic,
		"mem":      memDic,
		"cpu":      cpuDic,
		"disk":     diskDic,
		"diskList": disklist,
	})
}

func getDiskTotalAndUsed() (total uint64, used uint64) {
	diskList, _ := disk.Partitions(false)
	devices := make(map[string]string)
	countedDiskForMac := make(map[string]struct{})
	for _, d := range diskList {
		fsType := strings.ToLower(d.Fstype)
		// 不统计 K8s 的虚拟挂载点：https://github.com/shirou/gopsutil/issues/1007
		if devices[d.Device] == "" && isListContainsStr(expectDiskFsTypes, fsType) && !strings.Contains(d.Mountpoint, "/var/lib/kubelet") {
			devices[d.Device] = d.Mountpoint
		}
	}
	for device, mountPath := range devices {
		diskUsageOf, _ := disk.Usage(mountPath)
		// 这里是针对 Mac 机器的处理，https://github.com/giampaolo/psutil/issues/1509
		matches := getMacDiskNo.FindStringSubmatch(device)
		if len(matches) == 2 {
			if _, has := countedDiskForMac[matches[1]]; !has {
				countedDiskForMac[matches[1]] = struct{}{}
				total += diskUsageOf.Total
			}
		} else {
			total += diskUsageOf.Total
		}
		used += diskUsageOf.Used
	}

	// Fallback 到这个方法,仅统计根路径,适用于OpenVZ之类的.
	if runtime.GOOS == "linux" {
		if total == 0 && used == 0 {
			cmd := exec.Command("df")
			out, err := cmd.CombinedOutput()
			if err == nil {
				s := strings.Split(string(out), "\n")
				for _, c := range s {
					info := strings.Fields(c)
					if len(info) == 6 {
						if info[5] == "/" {
							total, _ = strconv.ParseUint(info[1], 0, 64)
							used, _ = strconv.ParseUint(info[2], 0, 64)
							total = total * 1024
							used = used * 1024
						}
					}
				}
			}
		}
	}
	return
}

func isListContainsStr(list []string, str string) bool {
	for i := 0; i < len(list); i++ {
		if strings.Contains(str, list[i]) {
			return true
		}
	}
	return false
}
