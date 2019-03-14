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
	pullsourcedir := flag.String("pull-source-dir", "", "aws s3 pull dir")
	pulltargetdir := flag.String("pull-target-dir", "", "aws s3 pull local dir")
	pushsourcedir := flag.String("push-source-dir", "", "aws s3 push local dir")
	pushtargetdir := flag.String("push-target-dir", "", "aws s3 push dir")

	flag.Parse()
	if *pushsourcedir != "" && *pushtargetdir != "" && *pullsourcedir != "" && *pulltargetdir != "" {
		for {
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("当前时间：", timeStr, "同步间隔时间（s）：", *synctime)
			st := time.Duration(*synctime)
			time.Sleep(st * time.Second)
			runcmdpull(*pullsourcedir, *pulltargetdir)
			runcmdpush(*pushsourcedir, *pushtargetdir)
		}
	} else if *pullsourcedir != "" && *pulltargetdir != "" && *pushsourcedir == "" && *pushtargetdir == "" {
		for {
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("当前时间：", timeStr, "同步间隔时间（s）：", *synctime)
			st := time.Duration(*synctime)
			time.Sleep(st * time.Second)
			runcmdpull(*pullsourcedir, *pulltargetdir)
		}
	} else if *pullsourcedir == "" && *pulltargetdir == "" && *pushsourcedir != "" && *pushtargetdir != "" {
		for {
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("当前时间：", timeStr, "同步间隔时间（s）：", *synctime)
			st := time.Duration(*synctime)
			time.Sleep(st * time.Second)
			runcmdpush(*pullsourcedir, *pulltargetdir)
		}
	} else {
		log.Fatal("路径不能为空")
	}
}
func runcmdpush(pushsourcedir, pushtargetdir string) {
	//time.Sleep(2 * time.Second)
	cmd := exec.Command("aws", "s3", "sync", pushsourcedir, pushtargetdir, "--delete")

	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
	}
	fmt.Println("push源路径：", pushsourcedir, "push目标路径：", pushtargetdir)
}
func runcmdpull(pullsourcedir, pulltargetdir string) {
	cmd := exec.Command("aws", "s3", "sync", pullsourcedir, pulltargetdir, "--exact-timestamps", "--delete")
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
	}
	fmt.Println("pull源路径：", pullsourcedir, "pull目标路径：", pulltargetdir)
}
