package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	synctime := flag.Int("sync-time", 10, "sync time defualt 10s")
	targetdir := flag.String("target-dir", "s3://dsdata01/rec_plt_svr", "aws s3 dir defualt /dsdata01/rec_plt_svr")

	flag.Parse()

	for {
		timeStr := time.Now().Format("2018-12-17 11:11:11")
		st := time.Duration(*synctime)
		time.Sleep(st * time.Second)
		fmt.Println("当前实际：", timeStr, "同步间隔时间：", st, "拷贝目录：", *targetdir)
		runcmd(*targetdir)
	}
}
func runcmd(targetdir string) {
	cmd := exec.Command("aws", "s3", "sync", targetdir, "/data")
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
