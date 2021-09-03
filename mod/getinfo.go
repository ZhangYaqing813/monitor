package mod

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strconv"
	"time"
)

const Gib = 1073741824

type cpuinfos struct {
	CoreNmub     string
	TotalPercent string
}

type memroyinfos struct {
	Total       string
	Used        string
	UsedPercent string
	Free        string
}

type diskinfos struct {
	Total       string
	Used        string
	UsedPercent string
	Free        string
}

type Monitor struct {
	Cpu    cpuinfos
	Memroy memroyinfos
	Disks  diskinfos
}

type Javaproc struct {
	Name string
	Pid  uint64
}

//获取CPU 的简要信息
func (C *cpuinfos) Cpuinfo() (cpuinfo *cpuinfos) {

	//physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	//fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	//perPercents, _ := cpu.Percent(3*time.Second, true)

	//folat64 转string
	cpus, _ := json.Marshal(totalPercent)
	cpuin := string(cpus)

	return &cpuinfos{
		CoreNmub:     string(logicalCnt),
		TotalPercent: cpuin,
	}

}

//获取磁盘的简要信息
func (D *diskinfos) Diskinfo(paths string) (disks *diskinfos) {

	info, _ := disk.Usage(paths)

	//folat64 转string
	usepre, _ := json.Marshal(info.UsedPercent)
	return &diskinfos{
		Total:       uinttostr(info.Total / Gib),
		Used:        uinttostr(info.Used / Gib),
		UsedPercent: string(usepre),
		Free:        uinttostr(info.Free / Gib),
	}
}

//获取内存的简要信息
func (M *memroyinfos) Memoryinfo() (memroy *memroyinfos) {
	info, _ := mem.VirtualMemory()
	usepre, _ := json.Marshal(info.UsedPercent)
	return &memroyinfos{
		Total:       uinttostr(info.Total / Gib),
		Used:        uinttostr(info.Used / Gib),
		UsedPercent: string(usepre),
		Free:        uinttostr(info.Free / Gib),
	}
}

//转换函数
func uinttostr(x uint64) (y string) {
	return strconv.FormatUint(x, 10)
}
