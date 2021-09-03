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
	health()
	system()

	//go func() {
	//	//time.Sleep(60*time.Second)
	//	health()
	//}()
	//
	//go func() {
	//	//time.Sleep(60*time.Second)
	//
	//}()
	//time.Sleep(15*time.Second)

}

func health() {

	mailTo := []string{
		"zhangyaqing59@126.com",
		//"30960425@qq.com",
	}

	//service : port
	service := map[string]int{
		"goods":     10050,
		"message":   10040,
		"cms":       10010,
		"container": 10030,
		"gateway":   10000,
		"user":      10020,
		"admin":     10001,
	}

	status := map[string]string{
		"goods":     "OK",
		"message":   "OK",
		"cms":       "OK",
		"container": "OK",
		"gateway":   "OK",
		"user":      "OK",
		"admin":     "OK",
	}

	var t string
	//测试服务是否正常，把不正常的服务记录
	for k, v := range service {
		//组装uri
		uri := "http://127.0.0.1:" + string(v) + "/test/check"
		// 测试uri
		t = time.Now().Format("2006-01-02 15:04:05")
		res := mod.HealthCheck(uri)
		if res == "TimeOut" {

			status[k] = "Failed"
		}
	}

	subject := "SERVICE INFO"
	body := "<html>\n\n<body>\n<h4>Service status: </h4>\n" +
		"<table border=\"1\">\n" +
		"<tr>\n  <td>Service</td>\n  <td>Port</td>\n  <td>Status</td>\n  <td>Time</td>\n  <td>备注</td>\n</tr>\n" +
		"<tr>\n  <td>goods</td>\n  <td>10050</td>\n  <td>" + status["goods"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"<tr>\n  <td>message</td>\n  <td>10040</td>\n  <td>" + status["message"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"<tr>\n  <td>cms</td>\n  <td>10010</td>\n  <td>" + status["cms"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"<tr>\n  <td>container</td>\n  <td>10030</td>\n  <td>" + status["container"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"<tr>\n  <td>gateway</td>\n  <td>10000</td>\n  <td>" + status["gateway"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"<tr>\n  <td>user</td>\n  <td>10020</td>\n  <td>" + status["user"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"<tr>\n  <td>admin</td>\n  <td>10001</td>\n  <td>" + status["admin"] + "</td>\n  <td>" + t + "</td>\n  <td>600</td>\n</tr>\n" +
		"</table>\n\n</body>\n</html>\n"

	err := email.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println("email send failed ,", err)
	}

}

func system() {
	hosts := mod.Monitor{}
	cpuinfo := hosts.Cpu.Cpuinfo()
	diskinfo := hosts.Disks.Diskinfo("/data")
	memroy := hosts.Memroy.Memoryinfo()

	mailTo := []string{
		"zhangyaqing59@126.com",
		//"30960425@qq.com",
	}
	//邮件主题为"Hello"
	info_subject := "SYSTEM INFO"
	// 邮件正文
	info_body := "<html>\n" +
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

	err := email.SendMail(mailTo, info_subject, info_body)
	if err != nil {
		fmt.Println("email send failed ,", err)
	}

}
