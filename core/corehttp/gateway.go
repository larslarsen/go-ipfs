package corehttp

import (
	"fmt"
	"net"
	"net/http"

	core "github.com/ipfs/go-ipfs/core"
	config "github.com/ipfs/go-ipfs/repo/config"
	id "gx/ipfs/QmVCe3SNMjkcPgnpFhZs719dheq6xE7gJwjzV7aWcUM4Ms/go-libp2p/p2p/protocol/identify"
	bc "github.com/OpenBazaar/go-blockstackclient"
)

type GatewayConfig struct {
	Headers      map[string][]string
	Writable     bool
	PathPrefixes []string
	Resolver     *bc.BlockstackClient
	Authenticated bool
	Cookie        http.Cookie
	Username      string
	Password      string
}

func GatewayOption(resolver *bc.BlockstackClient, authenticated bool, authCookie http.Cookie, username, password string, paths ...string) ServeOption {
	return func(n *core.IpfsNode, _ net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		cfg, err := n.Repo.Config()
		if err != nil {
			return nil, err
		}

		gateway := newGatewayHandler(n, GatewayConfig{
			Headers:       cfg.Gateway.HTTPHeaders,
			Writable:      cfg.Gateway.Writable,
			PathPrefixes:  cfg.Gateway.PathPrefixes,
			Resolver:      resolver,
			Authenticated: authenticated,
			Cookie:        authCookie,
			Username:      username,
			Password:      password,
		})

		for _, p := range paths {
			mux.Handle(p+"/", gateway)
		}
		return mux, nil
	}
}

func VersionOption() ServeOption {
	return func(n *core.IpfsNode, _ net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Commit: %s\n", config.CurrentCommit)
			fmt.Fprintf(w, "Client Version: %s\n", id.ClientVersion)
			fmt.Fprintf(w, "Protocol Version: %s\n", id.LibP2PVersion)
		})
		return mux, nil
	}
}
