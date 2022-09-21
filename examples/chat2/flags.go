package main

import (
	"fmt"

	wcli "github.com/status-im/go-waku/waku/cliutils"
	"github.com/status-im/go-waku/waku/v2/protocol"

	"github.com/urfave/cli/v2"
)

type FleetValue struct {
	Value   *Fleet
	Default Fleet
}

func (v *FleetValue) Set(value string) error {
	if value == string(fleetProd) || value == string(fleetTest) || value == string(fleetNone) {
		*v.Value = Fleet(value)
		return nil
	}
	return fmt.Errorf("%s is not a valid option. need %+v", value, []Fleet{fleetProd, fleetTest, fleetNone})
}

func (v *FleetValue) String() string {
	if v.Value == nil {
		return string(v.Default)
	}
	return string(*v.Value)
}

func getFlags() []cli.Flag {
	// Defaults
	options.Fleet = fleetProd

	return []cli.Flag{
		&cli.IntFlag{
			Name:        "tcp-port",
			Aliases:     []string{"port", "p"},
			Value:       60000,
			Usage:       "Libp2p TCP listening port (0 for random)",
			Destination: &options.Port,
		},
		&cli.StringFlag{
			Name:        "address",
			Aliases:     []string{"host", "listen-address"},
			Value:       "0.0.0.0",
			Usage:       "Listening address",
			Destination: &options.Address,
		},
		&cli.GenericFlag{
			Name:  "nodekey",
			Usage: "P2P node private key as hex. (default random)",
			Value: &wcli.PrivateKeyValue{
				Value: &options.NodeKey,
			},
		},
		&cli.StringFlag{
			Name:        "content-topic",
			Usage:       "content topic to use for the chat",
			Value:       protocol.NewContentTopic("toy-chat", 2, "luzhou", "proto").String(),
			Destination: &options.ContentTopic,
		},
		&cli.GenericFlag{
			Name:  "fleet",
			Usage: "Select the fleet to connect to",
			Value: &FleetValue{
				Default: fleetProd,
				Value:   &options.Fleet,
			},
		},
		&cli.GenericFlag{
			Name:  "staticnode",
			Usage: "Multiaddr of peer to directly connect with. Option may be repeated",
			Value: &wcli.MultiaddrSlice{
				Values: &options.StaticNodes,
			},
		},
		&cli.StringFlag{
			Name:        "nickname",
			Usage:       "nickname to use in chat.",
			Destination: &options.Nickname,
			Value:       "Anonymous",
		},

		&cli.BoolFlag{
			Name:        "relay",
			Value:       true,
			Usage:       "Enable relay protocol",
			Destination: &options.Relay.Enable,
		},
		&cli.BoolFlag{
			Name:        "payloadV1",
			Value:       false,
			Usage:       "use Waku v1 payload encoding/encryption",
			Destination: &options.UsePayloadV1,
		},
		&cli.StringSliceFlag{
			Name:        "topics",
			Usage:       "List of topics to listen",
			Destination: &options.Relay.Topics,
		},
		&cli.BoolFlag{
			Name:        "store",
			Usage:       "Enable relay protocol",
			Value:       true,
			Destination: &options.Store.Enable,
		},
		&cli.GenericFlag{
			Name:  "storenode",
			Usage: "Multiaddr of a peer that supports store protocol.",
			Value: &wcli.MultiaddrValue{
				Value: &options.Store.Node,
			},
		},
		&cli.BoolFlag{
			Name:        "filter",
			Usage:       "Enable filter protocol",
			Destination: &options.Filter.Enable,
		},
		&cli.GenericFlag{
			Name:  "filternode",
			Usage: "Multiaddr of a peer that supports filter protocol.",
			Value: &wcli.MultiaddrValue{
				Value: &options.Filter.Node,
			},
		},
		&cli.BoolFlag{
			Name:        "lightpush",
			Usage:       "Enable lightpush protocol",
			Destination: &options.LightPush.Enable,
		},
		&cli.GenericFlag{
			Name:  "lightpushnode",
			Usage: "Multiaddr of a peer that supports lightpush protocol.",
			Value: &wcli.MultiaddrValue{
				Value: &options.LightPush.Node,
			},
		},
		&cli.BoolFlag{
			Name:        "discv5-discovery",
			Usage:       "Enable discovering nodes via Node Discovery v5",
			Destination: &options.DiscV5.Enable,
		},
		&cli.StringSliceFlag{
			Name:        "discv5-bootstrap-node",
			Usage:       "Text-encoded ENR for bootstrap node. Used when connecting to the network. Option may be repeated",
			Destination: &options.DiscV5.Nodes,
		},
		&cli.IntFlag{
			Name:        "discv5-udp-port",
			Value:       9000,
			Usage:       "Listening UDP port for Node Discovery v5.",
			Destination: &options.DiscV5.Port,
		},
		&cli.BoolFlag{
			Name:        "discv5-enr-auto-update",
			Usage:       "Discovery can automatically update its ENR with the IP address as seen by other nodes it communicates with.",
			Destination: &options.DiscV5.AutoUpdate,
		},
		&cli.BoolFlag{
			Name:        "dns-discovery",
			Usage:       "Enable DNS discovery",
			Destination: &options.DNSDiscovery.Enable,
		},
		&cli.StringFlag{
			Name:        "dns-discovery-url",
			Usage:       "URL for DNS node list in format 'enrtree://<key>@<fqdn>'",
			Destination: &options.DNSDiscovery.URL,
		},
		&cli.StringFlag{
			Name:        "dns-discovery-name-server",
			Aliases:     []string{"dns-discovery-nameserver"},
			Usage:       "DNS nameserver IP to query (empty to use system's default)",
			Destination: &options.DNSDiscovery.Nameserver,
		},
		&cli.BoolFlag{
			Name:        "rln-relay",
			Value:       false,
			Usage:       "Enable spam protection through rln-relay",
			Destination: &options.RLNRelay.Enable,
		},
		&cli.IntFlag{
			Name:        "rln-relay-membership-index",
			Value:       0,
			Usage:       "(experimental) the index of node in the rln-relay group: a value between 0-99 inclusive",
			Destination: &options.RLNRelay.MembershipIndex,
		},
		&cli.StringFlag{
			Name:        "rln-relay-pubsub-topic",
			Value:       "/waku/2/default-waku/proto",
			Usage:       "the pubsub topic for which rln-relay gets enabled",
			Destination: &options.RLNRelay.PubsubTopic,
		},
		&cli.StringFlag{
			Name:        "rln-relay-content-topic",
			Value:       "/toy-chat/2/luzhou/proto",
			Usage:       "the content topic for which rln-relay gets enabled",
			Destination: &options.RLNRelay.ContentTopic,
		},
		&cli.BoolFlag{
			Name:        "rln-relay-dynamic",
			Usage:       "Enable waku-rln-relay with on-chain dynamic group management",
			Destination: &options.RLNRelay.Dynamic,
		},
		&cli.StringFlag{
			Name:        "rln-relay-id",
			Usage:       "Rln relay identity secret key as a Hex string",
			Destination: &options.RLNRelay.IDKey,
		},
		&cli.StringFlag{
			Name:        "rln-relay-id-commitment",
			Usage:       "Rln relay identity commitment key as a Hex string",
			Destination: &options.RLNRelay.IDCommitment,
		},
		&cli.PathFlag{
			Name:        "rln-relay-cred-path",
			Usage:       "The path for persisting rln-relay credential",
			Value:       "",
			Destination: &options.RLNRelay.CredentialsPath,
		},
		// TODO: this is a good candidate option for subcommands
		// TODO: consider accepting a private key file and passwd
		&cli.GenericFlag{
			Name:  "eth-account-privatekey",
			Usage: "Ethereum Goerli testnet account private key used for registering in member contract",
			Value: &wcli.PrivateKeyValue{
				Value: &options.RLNRelay.ETHPrivateKey,
			},
		},
		&cli.StringFlag{
			Name:        "eth-client-address",
			Usage:       "Ethereum testnet client address",
			Value:       "ws://localhost:8545",
			Destination: &options.RLNRelay.ETHClientAddress,
		},
		&cli.GenericFlag{
			Name:  "eth-mem-contract-address",
			Usage: "Address of membership contract on an Ethereum testnet",
			Value: &wcli.AddressValue{
				Value: &options.RLNRelay.MembershipContractAddress,
			},
		},
	}
}