package admin

import (
	"github.com/astaxie/beego"
	"os"
	"runtime"
	"syscall"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

type MemStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

func MemStat() MemStatus {
	//自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	mem := MemStatus{}
	mem.Self = memStat.Alloc
	//系统占用,仅linux/mac下有效
	//system memory usage
	// sysInfo := new(syscall.Sysinfo_t)
	// err := syscall.Sysinfo(sysInfo)
	// if err == nil {
	// mem.All = sysInfo.Totalram * uint64(syscall.Getpagesize())
	// mem.Free = sysInfo.Freeram * uint64(syscall.Getpagesize())
	// mem.Used = mem.All - mem.Free
	// }
	return mem
}

type SysController struct {
	beego.Controller
}

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

func (this *SysController) Get() {

	hardwareDisk := DiskUsage("/")
	hardwareMem := MemStat()
	this.Data["osInfo"] = runtime.GOOS + runtime.GOARCH
	this.Data["goVer"] = runtime.Version()
	hostname, _ := os.Hostname()
	this.Data["hostname"] = hostname
	this.Data["goNumGoroutine"] = runtime.NumGoroutine()
	this.Data["hardwareNumCPU"] = runtime.NumCPU()
	this.Data["hardwareDiskAll"] = hardwareDisk.All / 1024 / 1024 / 1024
	this.Data["hardwareDiskUsed"] = hardwareDisk.Used / 1024 / 1024 / 1024
	this.Data["hardwareDiskFree"] = hardwareDisk.Free / 1024 / 1024 / 1024
	this.Data["hardwareMem"] = hardwareMem
	this.TplNames = "sys/sysinfo.html"
}
