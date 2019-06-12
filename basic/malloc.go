package main

import (
	"github.com/miekg/dns"
	"net"
	"time"
	"fmt"
)

func main() {
	message := dns.Msg{}

	startTime := time.Now()
	for i := 0; i <= 400000; i++ {
		DoUDP_new("", nil, &message)
	}
	fmt.Println("Do UDP with malloc memory each time, waste time:", time.Now().Sub(startTime))

}


func DoUDP(Net string, w dns.ResponseWriter, req *dns.Msg) {
	response := dns.Msg{}
	response.SetReply(req)
	rr_header := dns.RR_Header{
		Name: "test",
		Rrtype: dns.TypeA,
		Class: dns.ClassINET,
		Ttl: 10,
	}
	ans := &dns.A{Hdr:rr_header, A:net.ParseIP("127.0.0.1")}
	response.Answer = append(response.Answer, ans)

}

var
(
	response dns.Msg

	header = dns.RR_Header{
		Name: "test",
		Rrtype: dns.TypeA,
		Class: dns.ClassINET,
		Ttl: 10,
	}
)

func DoUDP_new(Net string, w dns.ResponseWriter, req *dns.Msg) {
	response.SetReply(req)
	ans := &dns.A{Hdr:header, A:net.ParseIP("127.0.0.1")}
	response.Answer = append(response.Answer, ans)
}