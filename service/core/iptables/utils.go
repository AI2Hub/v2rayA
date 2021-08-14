package iptables

import (
	"github.com/v2rayA/v2rayA/common"
	"github.com/v2rayA/v2rayA/common/cmds"
	"github.com/v2rayA/v2rayA/global"
	"golang.org/x/net/nettest"
	"net"
	"strconv"
)

func IPNet2CIDR(ipnet *net.IPNet) string {
	ones, _ := ipnet.Mask.Size()
	return ipnet.IP.String() + "/" + strconv.Itoa(ones)
}

func GetLocalCIDR() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var cidrs []string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			cidrs = append(cidrs, IPNet2CIDR(ipnet))
		}
	}
	return cidrs, nil
}

func IsIPv6Supported() bool {
	if global.GetEnvironmentConfig().ForceIPV6On {
		return true
	}
	if common.IsInDocker() {
		return false
	}
	if !nettest.SupportsIPv6() {
		return false
	}
	return cmds.IsCommandValid("ip6tables")
}
