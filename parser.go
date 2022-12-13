package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/exp/slog"
)

const (
	snapshotLen int32 = 1024
	prom              = false
	to                = 600 * time.Second
	msgQuery          = 'Q'
)

// capture listen tcp traffic with filter and parse SQL-query.
func capture(device string, port int) error {
	handle, err := pcap.OpenLive(device, snapshotLen, prom, to)
	if err != nil {
		return fmt.Errorf("open live: %w", err)
	}
	defer handle.Close()

	filter := fmt.Sprintf("tcp and port %d", port)

	if err := handle.SetBPFFilter(filter); err != nil {
		return fmt.Errorf("set filter: %w", err)
	}

	slog.Info("start capturing", slog.String("filter", filter), slog.String("device", device))

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	const lenMsgSize = 4

	for packet := range packetSource.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}
		tcp, _ := tcpLayer.(*layers.TCP)

		if tcp.DstPort == layers.TCPPort(port) {
			buf := bytes.NewBuffer(tcp.Payload)
			msgType, err := buf.ReadByte()
			if err != nil {
				continue
			}

			if msgType == msgQuery {
				size := make([]byte, lenMsgSize)
				_, err := buf.Read(size)
				if err != nil {
					slog.Error("failed to read size", err)
				}
				payload := make([]byte, int32(binary.BigEndian.Uint32(size))-lenMsgSize)
				if _, err = buf.Read(payload); err != nil {
					slog.Error("failed to read payload", err)
				}

				query := string(payload[:len(payload)-1])

				fmt.Println(query)
			}
		}
	}

	return nil
}

// deviceList provide network device list.
func deviceList() error {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return fmt.Errorf("find all devs: %w", err)
	}

	fmt.Println("network device list:")

	for i := range devices {
		fmt.Println(devices[i].Name)
	}

	return nil
}