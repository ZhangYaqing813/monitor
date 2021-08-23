package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	_ "github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/gomail.v2"
	_ "net/smtp"
	"strconv"
	"time"
)

type cpuinfos struct {
	CoreNmub string
	TotalPercent string
}

type memroyinfos struct {
	Total string
	Used string
	UsedPercent string
	Free string
}

type diskinfos struct {
	Total string
	Used string
	UsedPercent string
	Free string
}



func main() {

	cpuinfomation := cpuinfo()
	diskinfomation := diskinfo()
	memoryinfomation := memoryinfo()
	//cpus:= strconv.FormatFloat(cpuinfomation,'G',5,64)
	cpus ,_:= json.Marshal(cpuinfomation)

	cpuin := string(cpus)
	//hostnames,_ := host.Info()
	////h:= string(hostnames.Hostname)
	total :=strconv.FormatUint(memoryinfomation.Total, 10)
	used :=strconv.FormatUint(memoryinfomation.Used, 10)
	free := strconv.FormatUint(memoryinfomation.Free, 10)

	dtotal := strconv.FormatUint(diskinfomation.Total, 10)
	dfree := strconv.FormatUint(diskinfomation.Free, 10)
	//定义收件人
	mailTo := []string {
		"zhangyaqing59@126.com",

	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body:="<html>\n" +
		"<title> 主机信息 </title>\n" +
		"<body>\n" +
		"<h1>CPU info </h1>\n" +
		"   <p>CPU_usepr:" + cpuin + "</p>\n"    +
		"<h1>MEMORY info</h1>\n" +
		"   <p>Total:" + total + "</p>\n    " +
		"   <p>Total:" + used + "</p>\n    " +
		"	<p>Free:" + free + "</p>\n    " +
		"<h1>DISK info</h1>\n\n    " +
		"   <p>Total:" + dtotal + "</p>\n    "  +
		"   <p>Free:" + dfree + "</p>\n " +
		"</body>\n\n" +
		"</html>"
	SendMail(mailTo, subject, body)
}
func cpuinfo() (otalPercent [] float64){
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	//perPercents, _ := cpu.Percent(3*time.Second, true)
	return totalPercent
}

func diskinfo() (info *disk.UsageStat ) {

	info, _ = disk.Usage("/")


	return info
}

func memoryinfo() (v *mem.VirtualMemoryStat){
	v, _ = mem.VirtualMemory()
	return v
	//fmt.Printf("Total: %v, Available: %v, UsedPercent:%f%%\n", v.Total, v.Available, v.UsedPercent)
}

func SendMail(mailTo []string,subject string, body string ) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": "notification@cmes.io",
		"pass": "nO06fca81oiT",
		"host": "smtp.zoho.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err
}