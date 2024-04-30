package netip

import (
	"net"
	"strings"
)

type NetType int32

const (
	IPV4     NetType = 0x1
	IPV6     NetType = 0x2
	LOOPBACK NetType = 0x4
)

func GetLocalIPList() (map[string]string, error) {
	interfaceList, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var byName *net.Interface
	var addrList []net.Addr
	var oneAddrs []string
	ipList := make(map[string]string, len(interfaceList))
	for i := range interfaceList {
		byName, err = net.InterfaceByName(interfaceList[i].Name)
		if err != nil {
			return nil, err
		}
		addrList, err = byName.Addrs()
		if err != nil {
			return nil, err
		}
		for ii := range addrList {
			oneAddrs = strings.SplitN(addrList[ii].String(), "/", 2)
			ipList[interfaceList[i].Name] = oneAddrs[0]
		}
	}
	return ipList, nil
}

func GetLocalIp(nt NetType) ([]string, error) {
	ipList := []string{}
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrList {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if nt&IPV4 != 0 {
				ipList = append(ipList, ipNet.IP.String())
			} else if nt&IPV6 != 0 {
				ipList = append(ipList, ipNet.IP.String())
			} else if nt&LOOPBACK != 0 {
				ipList = append(ipList, ipNet.IP.String())
			}
		}
	}
	return ipList, nil
}
