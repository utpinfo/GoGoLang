package main

import (
	"flag"

	"github.com/sjlleo/netflix-verify/printer"
	"github.com/sjlleo/netflix-verify/verify"
)

var custom = flag.String("custom", "", "自定义测试NF影片ID\n绝命毒师的ID是70143836")
var address = flag.String("address", "", "本机网卡的IP")
var proxy = flag.String("proxy", "", "代理地址")

func main() {

	flag.Parse()

	r := verify.NewVerify(verify.Config{
		LocalAddr: *address,
		Custom:    *custom,
		Proxy:     *proxy,
	})

	printer.Print(*r)
}
