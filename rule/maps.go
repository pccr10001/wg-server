package rule

import (
	"net"
	"strings"
	"sync"
	"time"
)

var map1 sync.Map
var map2 sync.Map
var m *sync.Map

func InitMap() {
	m = &map1
}

func Lookup(key []byte) []net.IPNet {
	cidrs, ok := m.Load(key)
	if !ok {
		return nil
	}
	return cidrs.([]net.IPNet)
}
func Get(key []byte) string {
	cidrs, ok := m.Load(key)
	if !ok {
		return ""
	}
	cs := ""
	for _, cidr := range cidrs.([]net.IPNet) {
		cs += cidr.String()
		cs += ","
	}
	return cs[:len(cs)-1]
}
func Set(key string, cidrs string) {

	var cs []net.IPNet

	if strings.Index(cidrs, ",") >= 0 {
		for _, cidr := range strings.Split(cidrs, ",") {
			_, ipnet, err := net.ParseCIDR(cidr)
			if err == nil {
				cs = append(cs, *ipnet)
			}
		}
	} else {
		_, ipnet, err := net.ParseCIDR(cidrs)
		if err == nil {
			cs = append(cs, *ipnet)
		}
	}

	if m == &map1 {
		map2.Store(key, cs)
		m = &map2
	} else {
		map1.Store(key, cs)
		m = &map1
	}

	time.Sleep(500 * time.Millisecond)

	if m == &map2 {
		map1.Store(key, cs)
	} else {
		map2.Store(key, cs)
	}
}
