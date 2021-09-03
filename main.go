package main

import (
	"fmt"
	_ "github.com/shirou/gopsutil/host"
	"monitor/email"
	"monitor/mod"
	_ "net/smtp"
	"time"
)

func main() {

	go func() {
		for true {
			time.Sleep(300 * time.Second)

		}
	}()

	hosts := mod.Monitor{}

	cpuinfo := hosts.Cpu.Cpuinfo()
	diskinfo := hosts.Disks.Diskinfo("/data")
	memroy := hosts.Memroy.Memoryinfo()

	//定义收件人
	mailTo := []string{
		"zhangyaqing59@126.com",
	}
	//邮件主题为"Hello"
	subject := "Hello"
	// 邮件正文
	body := "<html>\n" +
		"<body>\n" +
		"<h4>CPU INFO：</h4>\n<table border=\"1\">\n<tr>\n  <td>Cores</td>\n  <td>" + cpuinfo.CoreNmub + "</td>\n</tr>\n" +
		"<tr>\n  <td>UseAge</td>\n  <td>" + cpuinfo.TotalPercent + "</td>\n</tr>\n</table>\n\n" +

		"<h4>MEM INFO：</h4>\n" +
		"<table border=\"1\">\n<tr>\n  <td>ToTal</td>\n  <td>" + memroy.Total + "</td>\n</tr>\n\n<tr>\n  " +
		"<td>Used</td>\n  <td>" + memroy.Used + "</td>\n</tr>\n\n" +
		"<tr>\n  <td>UseAge</td>\n  <td>" + memroy.UsedPercent + "</td>\n</tr>\n\n<tr>\n  " +
		"<td>Free</td>\n  <td>" + memroy.Free + "</td>\n</tr>\n\n</table>\n\n" +

		"<h4>DISK INFO：</h4>\n" +
		"<table border=\"1\">\n<tr>\n  <td>ToTal</td>\n  <td>" + diskinfo.Total + "</td>\n</tr>\n\n<tr>\n  " +
		"<td>Used</td>\n  <td>" + diskinfo.Used + "</td>\n</tr>\n\n<tr>\n  " +
		"<td>UseAge</td>\n  <td>" + diskinfo.UsedPercent + "</td>\n</tr>\n\n<tr>\n  " +
		"<td>Free</td>\n  <td>" + diskinfo.Free + "</td>\n</tr>\n\n</table>\n" +
		"</body>\n" +
		"</html>\n"

	err := email.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println("email send failed ,", err)
	}

}
