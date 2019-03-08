package descriptor

import (
	"context"
	"github.com/contiv/bgp-vpp/plugins/bgp/descriptor/adapter"
	"github.com/contiv/bgp-vpp/plugins/bgp/model"
	"github.com/ligato/cn-infra/db/keyval"
	"github.com/ligato/cn-infra/logging"
	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	api "github.com/osrg/gobgp/api"
	gobgp "github.com/osrg/gobgp/pkg/server"
)

const (
	globalDescriptorName = "global-conf"
)

//our descriptor
type GlobalDescriptor struct {
	log    logging.Logger
	broker keyval.ProtoBroker
	server *gobgp.BgpServer
	//scheduler kvs.KVScheduler
	//handlers GlobalConfigurator.GlobalConfAPI
}

// NewGlobalConfDescriptor creates a new instance of the descriptor.
func NewGlobalConfDescriptor(broker keyval.ProtoBroker, log logging.PluginLogger, server *gobgp.BgpServer) *GlobalDescriptor {
	// Set plugin descriptor init values
	return &GlobalDescriptor{
		log:    log.NewLogger("global-conf-descriptor"),
		broker: broker,
		server: server,
	}
}

// GetDescriptor returns descriptor suitable for registration (via adapter) with the KVScheduler.
func (d *GlobalDescriptor) GetDescriptor() *adapter.GlobalConfDescriptor {
	return &adapter.GlobalConfDescriptor{
		Name:          globalDescriptorName,
		NBKeyPrefix:   model.ModelBgpGlobal.KeyPrefix(),
		ValueTypeName: model.ModelBgpGlobal.ProtoName(),
		KeySelector:   model.ModelBgpGlobal.IsKeyValid,
		KeyLabel:      model.ModelBgpGlobal.StripKeyPrefix,
		Create:        d.Create,
		Delete:        d.Delete,
		//UpdateWithRecreate: d.UpdateWithRecreate,
		//Retrieve:             d.Retrieve,
		Dependencies:         d.Dependencies,
		RetrieveDependencies: []string{},
	}
}

// Create creates new value.
func (d *GlobalDescriptor) Create(key string, value *model.GlobalConf) (metadata interface{}, err error) {
	err = d.server.StartBgp(context.Background(), &api.StartBgpRequest{
		Global: &api.Global{
			As:         value.As,
			RouterId:   value.RouterId,
			ListenPort: value.ListenPort,
		},
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete removes an existing value.
func (d *GlobalDescriptor) Delete(key string, value *model.GlobalConf, metadata interface{}) error {
	err := d.server.StopBgp(context.Background(), &api.StopBgpRequest{})
	if err != nil {
		return err
	}

	return nil
}

// UpdateWithRecreate returns true if value update requires full re-creation.
/*func (d *GlobalDescriptor) UpdateWithRecreate(key string, old, new *model.GlobalConf, metadata interface{}) bool {
	return true
}*/

// Retrieve retrieves values from SB.
/*func (d *GlobalDescriptor) Retrieve(correlate []adapter.PluginKVWithMetadata) (retrieved []adapter.PluginKVWithMetadata, err error) {
	return retrieved, nil
}*/

// Dependencies lists dependencies of the given value.
func (d *GlobalDescriptor) Dependencies(key string, value *model.GlobalConf) (deps []kvs.Dependency) {
	return deps
}
