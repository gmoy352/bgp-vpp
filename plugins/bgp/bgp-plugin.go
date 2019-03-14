package bgp

//go:generate protoc --proto_path=model --proto_path=$GOPATH/src --gogo_out=model model/bgp.proto
//go:generate descriptor-adapter --descriptor-name GlobalConf --value-type *model.GlobalConf --import "model" --output-dir "descriptor"
//go:generate descriptor-adapter --descriptor-name PeerConf --value-type *model.PeerConf --import "model" --output-dir "descriptor"
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/contiv/bgp-vpp/plugins/bgp/descriptor"
	"github.com/contiv/vpp/plugins/ipv4net"
	"github.com/contiv/vpp/plugins/netctl/remote"
	"github.com/contiv/vpp/plugins/nodesync/vppnode"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/ligato/cn-infra/datasync"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/logging"

	//"strconv"

	//"github.com/ligato/cn-infra/db/keyval/etcd"
	"github.com/ligato/cn-infra/infra"
	"github.com/ligato/cn-infra/rpc/rest"
	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	bgp_api "github.com/osrg/gobgp/api"
	gobgp "github.com/osrg/gobgp/pkg/server"
	"io/ioutil"
	"log"
	"strings"
)

type BgpPlugin struct {
	Deps
	watchCloser chan string
}

//Deps is only for external dependencies
type Deps struct {
	//plugins - initialized in options.go NewPlugin()
	infra.PluginDeps
	Rest        *rest.Plugin
	BGPServer   *gobgp.BgpServer
	KVScheduler kvs.KVScheduler
	KVStore     keyval.KvProtoPlugin
}

const nodePrefix = "/vnf-agent/contiv-ksr/allocatedIds/"
const getIpamDataCmd = "contiv/v1/ipam"

func (p *BgpPlugin) String() string {
	return "Starting BgpPlugin Application"
}
func (p *BgpPlugin) Init() error {
	if p.BGPServer == nil {
		p.BGPServer = gobgp.NewBgpServer()
		go p.BGPServer.Serve()
	}

	// register descriptor for bgp global config
	gd := descriptor.NewGlobalConfDescriptor(p.Log, p.BGPServer)
	p.KVScheduler.RegisterKVDescriptor(gd)

	// register descriptor for bgp peer config

	pd := descriptor.NewPeerConfDescriptor(p.Log, p.BGPServer)
	p.KVScheduler.RegisterKVDescriptor(pd)

	p.watchCloser = make(chan string)
	watcher := p.Deps.KVStore.NewWatcher(nodePrefix)
	err := watcher.Watch(p.onChange, p.watchCloser, "")
	if err != nil {
		return err

	}
	return nil
}
func (p *BgpPlugin) Close() error {
	log.Println("Closing BgpPlugin Application.")
	return nil
}

//tutorial says keyval.protowatchresp but i couldnt find it
//in the  cn-infra/db/keyval/proto_watcher_api.go in our vendor folder there isnt a ProtoWatchResp
// but on the github for cn infra, there is one
func (p *BgpPlugin) onChange(resp datasync.ProtoWatchResp) {
	//key := resp.GetKey()

	//Getting
	value := new(vppnode.VppNode)
	changeType := resp.GetChangeType()
	if changeType == datasync.Delete {
		if prevValExist, err := resp.GetPrevValue(value); err != nil {
			logging.Errorf("get value error: %v", err)
		} else if !prevValExist {
			logging.Errorf("Delete and Previous Value Does not Exist")
		}
	} else {
		if err := resp.GetValue(value); err != nil {
			logging.Errorf("get value error: %v", err)
		}
	}
	ip := value.IpAddresses[0]
	ipParts := strings.Split(ip, "/")
	ip = ipParts[0]
	//prefixlen, err4 := strconv.Atoi(ipParts[1])

	client, err := remote.CreateHTTPClient("")
	if err != nil {
		logging.Error("create http client error: %v", err)
	}

	//Getting Ipam Info
	b, err2 := getNodeInfo(client, ip, getIpamDataCmd)
	if err2 != nil {
		logging.Errorf("getnodeinfo error: %v", err2)
	}

	ipam := ipv4net.IPAMData{}
	err3 := json.Unmarshal(b, &ipam)
	if err3 != nil {
		logging.Errorf("unmarshal error: %v", err3)
	}

	//Setting Route info
	nlri, _ := ptypes.MarshalAny(&bgp_api.IPAddressPrefix{
		Prefix:    ipam.PodSubnetThisNode,
		PrefixLen: 24,
	})

	a1, _ := ptypes.MarshalAny(&bgp_api.OriginAttribute{
		Origin: 0,
	})
	a2, _ := ptypes.MarshalAny(&bgp_api.NextHopAttribute{
		NextHop: ip,
	})
	attrs := []*any.Any{a1, a2}
	if changeType == datasync.Put {
		_, err6 := p.Deps.BGPServer.AddPath(context.Background(), &bgp_api.AddPathRequest{
			Path: &bgp_api.Path{
				Family: &bgp_api.Family{Afi: bgp_api.Family_AFI_IP, Safi: bgp_api.Family_SAFI_UNICAST},
				Nlri:   nlri,
				Pattrs: attrs,
			},
		})
		if err6 != nil {
			logging.Errorf("AddPath: %v", err6)
		}
	} else if changeType == datasync.Delete {
		err7 := p.Deps.BGPServer.DeletePath(context.Background(), &bgp_api.DeletePathRequest{
			Path: &bgp_api.Path{
				Family: &bgp_api.Family{Afi: bgp_api.Family_AFI_IP, Safi: bgp_api.Family_SAFI_UNICAST},
				Nlri:   nlri,
				Pattrs: attrs,
			},
		})
		if err7 != nil {
			logging.Errorf("AddPath: %v", err7)
		}
	} else {
		logging.Errorf("GetChangeType: %v", changeType)
	}

}

// getNodeInfo will make an http request for the given command and return an indented slice of bytes.
//copied from vpp/plugins/netctl/cmdimpl/nodes.go
func getNodeInfo(client *remote.HTTPClient, base string, cmd string) ([]byte, error) {
	res, err := client.Get(base, cmd)
	if err != nil {
		err := fmt.Errorf("getNodeInfo: url: %s Get Error: %s", cmd, err.Error())
		fmt.Printf("http get error: %s ", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		err := fmt.Errorf("getNodeInfo: url: %s HTTP res.Status: %s", cmd, res.Status)
		fmt.Printf("http get error: %s ", err.Error())
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
