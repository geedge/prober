package main

import (
	"github.com/miekg/dns"
	"github.com/panjf2000/ants/v2"
	"github.com/thanhpk/randstr"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	defer ants.Release()

	var wg sync.WaitGroup

	p, _ := ants.NewPool(10000, ants.WithPreAlloc(true))

	c := new(dns.Client)
	args := os.Args
	qname := args[1]
	runcount, _ := strconv.Atoi(args[2])
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
		}
	}
	wg.Wait()
	print("完成！！")
}
