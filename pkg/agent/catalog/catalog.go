package catalog

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager/disk"
	"github.com/spiffe/spire/pkg/agent/plugin/keymanager/memory"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor/aws"
	"github.com/spiffe/spire/pkg/agent/plugin/nodeattestor/jointoken"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/k8s"
	"github.com/spiffe/spire/pkg/agent/plugin/workloadattestor/unix"
	"github.com/spiffe/spire/proto/agent/keymanager"
	"github.com/spiffe/spire/proto/agent/nodeattestor"
	"github.com/spiffe/spire/proto/agent/workloadattestor"

	goplugin "github.com/hashicorp/go-plugin"
	common "github.com/spiffe/spire/pkg/common/catalog"
)

const (
	KeyManagerType       = "KeyManager"
	NodeAttestorType     = "NodeAttestor"
	WorkloadAttestorType = "WorkloadAttestor"
)

type Catalog interface {
	common.Catalog

	KeyManagers() []keymanager.KeyManager
	NodeAttestors() []nodeattestor.NodeAttestor
	WorkloadAttestors() []workloadattestor.WorkloadAttestor
}

var (
	supportedPlugins = map[string]goplugin.Plugin{
		KeyManagerType:       &keymanager.KeyManagerPlugin{},
		NodeAttestorType:     &nodeattestor.NodeAttestorPlugin{},
		WorkloadAttestorType: &workloadattestor.WorkloadAttestorPlugin{},
	}

	builtinPlugins = common.BuiltinPluginMap{
		KeyManagerType: {
			"disk":   disk.New(),
			"memory": memory.New(),
		},
		NodeAttestorType: {
			"aws_iid":    aws.NewIID(),
			"join_token": jointoken.New(),
		},
		WorkloadAttestorType: {
			"k8s":  k8s.New(),
			"unix": unix.New(),
		},
	}
)

type Config struct {
	PluginConfigs common.PluginConfigMap
	Log           logrus.FieldLogger
}

type catalog struct {
	com common.Catalog
	m   *sync.RWMutex
	log logrus.FieldLogger

	keyManagerPlugins       []keymanager.KeyManager
	nodeAttestorPlugins     []nodeattestor.NodeAttestor
	workloadAttestorPlugins []workloadattestor.WorkloadAttestor
}

func New(c *Config) Catalog {
	commonConfig := &common.Config{
		PluginConfigs:    c.PluginConfigs,
		SupportedPlugins: supportedPlugins,
		BuiltinPlugins:   builtinPlugins,
		Log:              c.Log,
	}

	return &catalog{
		log: c.Log,
		com: common.New(commonConfig),
		m:   new(sync.RWMutex),
	}
}

func (c *catalog) Run(ctx context.Context) error {
	c.m.Lock()
	defer c.m.Unlock()

	err := c.com.Run(ctx)
	if err != nil {
		return err
	}

	return c.categorize()
}

func (c *catalog) Stop() {
	c.m.Lock()
	defer c.m.Unlock()

	c.com.Stop()
	c.reset()

	return
}

func (c *catalog) Reload(ctx context.Context) error {
	c.m.Lock()
	defer c.m.Unlock()

	err := c.com.Reload(ctx)
	if err != nil {
		return err
	}

	return c.categorize()
}

func (c *catalog) Plugins() []*common.ManagedPlugin {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.com.Plugins()
}

func (c *catalog) Find(plugin common.Plugin) *common.ManagedPlugin {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.com.Find(plugin)
}

func (c *catalog) KeyManagers() []keymanager.KeyManager {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.keyManagerPlugins
}

func (c *catalog) NodeAttestors() []nodeattestor.NodeAttestor {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.nodeAttestorPlugins
}

func (c *catalog) WorkloadAttestors() []workloadattestor.WorkloadAttestor {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.workloadAttestorPlugins
}

// categorize iterates over all managed plugins and casts them into their
// respective client types. This method is called during Run and Reload
// to prevent the consumer from having to check for errors when fetching
// a client from the catalog
func (c *catalog) categorize() error {
	c.reset()

	errMsg := "Plugin %s does not adhere to %s interface"
	for _, p := range c.com.Plugins() {
		if !p.Config.Enabled {
			c.log.Debugf("%s plugin %s is disabled and will not be categorized", p.Config.PluginType, p.Config.PluginName)
			continue
		}

		switch p.Config.PluginType {
		case KeyManagerType:
			pl, ok := p.Plugin.(keymanager.KeyManager)
			if !ok {
				return fmt.Errorf(errMsg, p.Config.PluginName, KeyManagerType)
			}
			c.keyManagerPlugins = append(c.keyManagerPlugins, pl)
		case NodeAttestorType:
			pl, ok := p.Plugin.(nodeattestor.NodeAttestor)
			if !ok {
				return fmt.Errorf(errMsg, p.Config.PluginName, NodeAttestorType)
			}
			c.nodeAttestorPlugins = append(c.nodeAttestorPlugins, pl)
		case WorkloadAttestorType:
			pl, ok := p.Plugin.(workloadattestor.WorkloadAttestor)
			if !ok {
				return fmt.Errorf(errMsg, p.Config.PluginName, WorkloadAttestorType)
			}
			c.workloadAttestorPlugins = append(c.workloadAttestorPlugins, pl)
		default:
			return fmt.Errorf("Unsupported plugin type %s", p.Config.PluginType)
		}
	}

	// Guarantee we have at least one of each type
	pluginCount := map[string]int{}
	pluginCount[KeyManagerType] = len(c.keyManagerPlugins)
	pluginCount[NodeAttestorType] = len(c.nodeAttestorPlugins)
	pluginCount[WorkloadAttestorType] = len(c.workloadAttestorPlugins)
	for t, c := range pluginCount {
		if c < 1 {
			return fmt.Errorf("At least one plugin of type %s is required", t)
		}
	}

	return nil
}

func (c *catalog) reset() {
	c.keyManagerPlugins = nil
	c.nodeAttestorPlugins = nil
	c.workloadAttestorPlugins = nil
}
