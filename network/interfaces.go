package network

import (
    "github.com/wsvn53/goutils/multivalues"
    "net"
    "golang.org/x/net/route"
    "reflect"
    "fmt"
    "syscall"
)

type Interface struct {
    Name string
    Index int
    Gateway net.IP
    Addrs []string
}

func (i Interface) String() string {
    return fmt.Sprintf(`{
    Name:  %s,
    Index: %d,
    Gateway: %s,
    Addrs:   %v,
}`, i.Name, i.Index, i.Gateway.String(), i.Addrs)
}

func GetInterfaces() []*Interface {
    outputInterfaces := make([]*Interface, 0)
    interfaceMap := make(map[int]*Interface)

    // get interface basic infomations
    interfaceList, _ := net.Interfaces()
    for _, iface := range interfaceList {
        niface := &Interface{
            Name: iface.Name,
            Index: iface.Index,
            Addrs: stringifyAddrs(multivalues.First(iface.Addrs()).([]net.Addr)),
        }
        outputInterfaces = append(outputInterfaces, niface)
        interfaceMap[iface.Index] = niface
    }

    // fetach route gateway from RIB
    ribBytes, _ := route.FetchRIB(0, route.RIBTypeRoute, 0)
    ribMessages, _ := route.ParseRIB(route.RIBTypeRoute, ribBytes)
    for _, ribMsg := range ribMessages {
        routeMsg := ribMsg.(*route.RouteMessage)
        if len(routeMsg.Addrs) < 2 {
            continue
        }

        // detect flags contains G = gateway
        if routeMsg.Flags & syscall.RTF_GATEWAY == 0 {
            continue
        }

        niface := interfaceMap[routeMsg.Index]
        if reflect.TypeOf(routeMsg.Addrs[1]).String() != "*route.Inet4Addr" {
            continue
        }
        gwAddr := routeMsg.Addrs[1].(*route.Inet4Addr)
        gwIpAddr := net.IPv4(gwAddr.IP[0], gwAddr.IP[1], gwAddr.IP[2], gwAddr.IP[3])
        niface.Gateway = gwIpAddr
    }

    return outputInterfaces
}

func stringifyAddrs(addrs interface{}) []string {
    outputAddrs := make([]string, 0)
    addrsValue := reflect.ValueOf(addrs)
    for i, l := 0, addrsValue.Len(); i<l; i++ {
        outputAddrs = append(outputAddrs, fmt.Sprint(addrsValue.Index(i)))
    }
    return outputAddrs
}

