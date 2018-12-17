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
		st := time.Duration(*synctime)
		time.Sleep(st * time.Second)
		fmt.Println("时间：", *synctime, "拷贝目录：", *targetdir)
	}
}
func runcmd(targetdir string) {
	cmd := exec.Command("aws", "s3", "sync", targetdir, "/data")
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
