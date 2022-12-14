package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"golang.org/x/exp/slog"
)

const (
	snapshotLen int32 = 1024
	prom              = false
	timeout           = 1 * time.Millisecond
	msgQuery          = 'Q'
)

// capture listen tcp traffic with filter and parse SQL-query.
func capture(device string, port int, highlight bool) error {
	handle, err := pcap.OpenLive(device, snapshotLen, prom, timeout)
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
	packetSource.DecodeOptions.Lazy = true
	packetSource.DecodeOptions.NoCopy = true

	for packet := range packetSource.Packets() {
		if packet == nil {
			continue
		}

		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		if tcpLayer == nil {
			continue
		}

		tcp, ok := tcpLayer.(*layers.TCP)
		if !ok {
			slog.Warn("could`t cast packet layer")
			continue
		}

		if tcp.DstPort == layers.TCPPort(port) {
			buf := bytes.NewBuffer(tcp.Payload)

			msgType, err := buf.ReadByte()
			if err != nil {
				continue
			}

			if msgType == msgQuery {
				fmt.Println("- - - >")

				query := extractQuery(buf)

				if highlight {
					highlightSQL(query)
				} else {
					fmt.Println(query)
				}

				fmt.Println("- - - >")
			}
		}
	}

	return nil
}

// extractQuery extract SQL-query from pocket payload.
func extractQuery(buf io.Reader) string {
	const lenMsgSize = 4

	size := make([]byte, lenMsgSize)

	if _, err := buf.Read(size); err != nil {
		slog.Error("failed timeout read size", err)
	}

	payload := make([]byte, int32(binary.BigEndian.Uint32(size))-lenMsgSize)

	if _, err := buf.Read(payload); err != nil {
		slog.Error("failed timeout read payload", err)
	}

	return string(payload[:len(payload)-1])
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
