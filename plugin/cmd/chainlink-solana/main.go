package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/pelletier/go-toml/v2"
	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink-relay/pkg/logger"
	"github.com/smartcontractkit/chainlink-relay/pkg/loop"
	"github.com/smartcontractkit/chainlink-relay/pkg/types"
	pkgsol "github.com/smartcontractkit/chainlink-solana/pkg/solana"

	"github.com/smartcontractkit/chainlink/core/chains/solana"
)

func main() {
	lggr, err := logger.New()
	cp := &chainPlugin{lggr: lggr}
	//TODO graceful shutdown via signal.Notify()? then cp.Close()
	if err != nil {
		fmt.Println("Failed to create logger:", err)
		os.Exit(1)
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: loop.HandshakeConfig(),
		Plugins: map[string]plugin.Plugin{
			loop.NameChain: loop.NewGRPCChainPlugin(cp, lggr),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

type chainPlugin struct {
	lggr    logger.Logger
	closers []io.Closer //TODO need to sync
}

func (c *chainPlugin) NewRelayer(ctx context.Context, config string, keystore loop.Keystore) (types.Relayer, error) {
	d := toml.NewDecoder(strings.NewReader(config))
	d.DisallowUnknownFields()
	var cfg solana.SolanaConfig
	if err := d.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config toml: %w", err)
	}

	//TODO ensure nothing else is running? or just not this chain ID?

	chainSet, err := solana.NewChainSetImmut(solana.ChainSetOpts{
		Logger:   c.lggr,   //TODO expand Logger interface
		KeyStore: keystore, //TODO simplify Keystore interface
	}, solana.SolanaConfigs{&cfg})
	if err != nil {
		return nil, fmt.Errorf("failed to create chain: %w", err)
	}
	if err := chainSet.Start(ctx); err != nil {
		return nil, fmt.Errorf("failed to start chain: %w", err)
	}
	r := pkgsol.NewRelayer(c.lggr, chainSet)
	c.closers = append(c.closers, chainSet, r)
	return r, nil
}

func (c *chainPlugin) Close() (err error) {
	for _, cl := range c.closers {
		if e := cl.Close(); e != nil {
			err = multierr.Append(err, e)
		}
	}
	return
}
