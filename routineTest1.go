package main

import "fmt"
import "time"
import "math/rand"

func routine(name string, delay time.Duration){
	t0:= time.Now()
	fmt.Println(name, "开始于 ", t0)

	time.Sleep(delay)

	t1:= time.Now()
	fmt.Println(name, "结束于 ", t1)

	fmt.Println(name, "共执行 ", t1.Sub(t0))
}

func main1(){
	//生成随机种子
	rand.Seed(time.Now().Unix())

	var name string
	for i:=0; i<3;i++{
		//生成名称
		name = fmt.Sprintf("go_%d", i)

		go routine(name, time.Duration(rand.Intn(5)) * time.Second)
	}

	//主进程等待
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}