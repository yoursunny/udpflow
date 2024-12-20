package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/netip"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func main() {
	var localA, remoteA, localB, remoteB netip.AddrPort
	var sockA, sockB *net.UDPConn
	app := &cli.App{
		ArgsUsage: "localA:port remoteA:port localB:port remoteB:port",
		Before: func(ctx *cli.Context) (e error) {
			args := ctx.Args()
			addrPorts := []*netip.AddrPort{&localA, &remoteA, &localB, &remoteB}
			if args.Len() != len(addrPorts) {
				return fmt.Errorf("expect %d positional arguments", len(addrPorts))
			}
			for i, addrPort := range addrPorts {
				*addrPort, e = netip.ParseAddrPort(args.Get(i))
				if e != nil {
					return fmt.Errorf("ParseAddrPort(%d): %w", i, e)
				}
			}
			return nil
		},
		Action: func(ctx *cli.Context) (e error) {
			sockA, e = net.DialUDP("udp", net.UDPAddrFromAddrPort(localA), net.UDPAddrFromAddrPort(remoteA))
			if e != nil {
				return fmt.Errorf("DialUDP(%v,%v): %w", localA, remoteA, e)
			}
			defer sockA.Close()

			sockB, e = net.DialUDP("udp", net.UDPAddrFromAddrPort(localB), net.UDPAddrFromAddrPort(remoteB))
			if e != nil {
				return fmt.Errorf("DialUDP(%v,%v): %w", localB, remoteB, e)
			}
			defer sockB.Close()

			var g errgroup.Group
			g.Go(func() error {
				_, e := io.Copy(sockB, sockA)
				return e
			})
			g.Go(func() error {
				_, e := io.Copy(sockA, sockB)
				return e
			})
			return g.Wait()
		},
	}
	if e := app.Run(os.Args); e != nil {
		log.Fatal(e)
	}
}
