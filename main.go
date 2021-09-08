package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	_ "github.com/shirou/gopsutil/host"
	"monitor/email"
	"monitor/mod"
	_ "net/smtp"
	"time"
)

func main() {

	systemCrontab := cron.New()
	systmeTask := func() {
		system()
	}
	// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	_, err := systemCrontab.AddFunc("* * * * *", systmeTask)
	if err != nil {
		fmt.Println(err)
	}
	// 启动定时器
	systemCrontab.Start()

	healthCrontab := cron.New()

	healthTask := func() {
		health()
	}

	_, err = healthCrontab.AddFunc("* * * * *", healthTask)
	if err != nil {
		fmt.Println(err)
	}
	healthCrontab.Start()

	select {}

}

func health() {

	mailTo := []string{
		"zhangyaqing59@126.com",
		//"30960425@qq.com",
	}

	//service : port
	service := map[string]string{
		"goods":     "http://127.0.0.1:10050",
		"message":   "http://127.0.0.1:10040",
		"cms":       "http://127.0.0.1:10010",
		"container": "http://127.0.0.1:10030",
		"gateway":   "http://127.0.0.1:10000",
		"user":      "http://127.0.0.1:10020",
		"admin":     "http://127.0.0.1:10001",
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
		uri := string(v) + "/test/check"
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

	for _, v := range status {
		if v == "Failed" {
			err := email.SendMail(mailTo, subject, body)
			if err != nil {
				fmt.Println("email send failed ,", err)
			}
		}
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
