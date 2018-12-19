package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	synctime := flag.Int("sync-time", 30, "sync time defualt 30s")
	pullsourcedir := flag.String("pull-source-dir", "s3://dsdata01/rec_plt_svr/base_pull", "aws s3 pull dir")
	pulltargetdir := flag.String("pull-target-dir", "/data/base_pull", "aws s3 pull local dir")
	pushsourcedir := flag.String("push-source-dir", "/data/base_push", "aws s3 push local dir")
	pushtargetdir := flag.String("push-target-dir", "s3://dsdata01/rec_plt_svr/base_pull", "aws s3 push dir")

	flag.Parse()
	if *pushsourcedir != "/data/base_push" {
		if *pushtargetdir != "s3://dsdata01/rec_plt_svr/base_pull" {
			for {
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Println("当前时间：", timeStr, "同步间隔时间（s）：", *synctime)
				st := time.Duration(*synctime)
				time.Sleep(st * time.Second)
				runcmdpullandpush(*pushsourcedir, *pushtargetdir, *pullsourcedir, *pulltargetdir)
			}
		} else {
			log.Fatal("没有目的路径")
		}
	} else if *pullsourcedir != "s3://dsdata01/rec_plt_svr/base_pull" {
		if *pulltargetdir != "/data/base_pull" {
			for {
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Println("当前时间：", timeStr, "同步间隔时间（s）：", *synctime)
				st := time.Duration(*synctime)
				time.Sleep(st * time.Second)
				runcmdpull(*pullsourcedir, *pulltargetdir)
			}
		} else {
			log.Fatal("没有目的路径")
		}
	} else {
		log.Fatal("缺少源路径")
	}
}
func runcmdpullandpush(pushsourcedir string, pushtargetdir string, pullsourcedir string, pulltargetdir string) {
	runcmdpull(pullsourcedir, pulltargetdir)
	//time.Sleep(2 * time.Second)
	cmd := exec.Command("aws", "s3", "sync", pushsourcedir, pushtargetdir, "--delete")

	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pull源路径：", pullsourcedir, "pull目标路径：", pulltargetdir, "push源路径：", pushsourcedir, "push目标路径：", pushtargetdir)
}
func runcmdpull(pullsourcedir string, pulltargetdir string) {
	cmd := exec.Command("aws", "s3", "sync", pullsourcedir, pulltargetdir, "--delete")
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pull源路径：", pullsourcedir, "pull目标路径：", pulltargetdir)
}
