package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	device       string = "lo"
	snapshot_len int32  = 65536
	promiscuous  bool   = false
	err          error
	handle       *pcap.Handle
	// Will reuse these for each packet
	ethLayer layers.Ethernet
	ip4layer layers.IPv4
	ip6layer layers.IPv6
	tcpLayer layers.TCP
)

type Connection struct {
	TotalBytes  int64
	Opened      time.Time
	Closed      time.Time
	Identifier  string
	IsClosing   bool
	PacketCount int64
}
type DownloadCount struct {
	Count      int64
	Time       time.Time
	Identifier string
}
type CumilatedDownloadStats struct {
	Data map[int64][]*DownloadCount
	mux  sync.Mutex
}

var connectionList map[string]*Connection
var userDownloadStats *CumilatedDownloadStats
var channelStartTime time.Time

func processDump(dataChan <-chan *DownloadCount) {
	defer func() {
		log.Println("end!")
	}()
	log.Println("started go routine")
	var statsCount int64
	for {
		count := <-dataChan
		// log.Println("one")
		if strings.Contains(count.Identifier, "(redis)") {
			continue
		}
		unix := count.Time.Unix()
		userDownloadStats.mux.Lock()
		userDownloadStats.Data[unix] = append(userDownloadStats.Data[unix], count)
		userDownloadStats.mux.Unlock()
		if statsCount%100 == 1 && statsCount > 1 {
			log.Println("HEUEHUEHE!!")
			// do something every 100th element
			// go dumpToFile(unix - 2)
			statsCount = 0
			go dumpToFile()
		}
		statsCount++
		log.Println(statsCount)
	}
}
func dumpToFile() {
	intstr := strconv.FormatInt(time.Now().UnixNano(), 10)
	file, _ := os.OpenFile("./dumps/"+intstr, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	for i, v := range userDownloadStats.Data {
		for _, iv := range v {
			unixstring := strconv.FormatInt(iv.Time.Unix(), 10)
			file.WriteString(unixstring + "/" + iv.Identifier + "/" + strconv.FormatInt(iv.Count, 10) + "\n")
		}
		userDownloadStats.mux.Lock()
		delete(userDownloadStats.Data, i)
		userDownloadStats.mux.Unlock()
	}
}
func main() {
	connectionList = make(map[string]*Connection)
	userDownloadStats = &CumilatedDownloadStats{
		Data: make(map[int64][]*DownloadCount),
	}
	dataChan := make(chan *DownloadCount)
	go processDump(dataChan)
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeEthernet,
			&ethLayer,
			&ip4layer,
			&tcpLayer,
		)
		foundLayerTypes := []gopacket.LayerType{}

		_ = parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			fmt.Println("Trouble decoding layers: ", err)
		}

		identifier := ""
		identifier2 := ""
		activeIdentifier := ""
		var ack uint32
		var fin bool
		var syn bool
		var rst bool
		for _, layerType := range foundLayerTypes {

			if layerType == layers.LayerTypeIPv6 {
				identifier = identifier + ip6layer.SrcIP.String() + "/" + ip6layer.DstIP.String()
				identifier2 = identifier2 + ip6layer.DstIP.String() + "/" + ip6layer.SrcIP.String()
			}
			if layerType == layers.LayerTypeIPv4 {
				identifier = identifier + ip4layer.SrcIP.String() + "/" + ip4layer.DstIP.String()
				identifier2 = identifier2 + ip4layer.DstIP.String() + "/" + ip4layer.SrcIP.String()
			}
			if layerType == layers.LayerTypeTCP {
				ack = tcpLayer.Ack
				fin = tcpLayer.FIN
				rst = tcpLayer.RST
				syn = tcpLayer.SYN
				identifier = identifier + "/" + tcpLayer.SrcPort.String() + "/" + tcpLayer.DstPort.String()
				identifier2 = identifier2 + "/" + tcpLayer.DstPort.String() + "/" + tcpLayer.SrcPort.String()
			}

		}

		if connectionList[identifier] != nil {
			activeIdentifier = identifier
		} else if connectionList[identifier2] != nil {
			activeIdentifier = identifier2
		} else {
			connectionList[identifier] = &Connection{
				Opened: packet.Metadata().Timestamp,
			}
			activeIdentifier = identifier
		}

		if ack == 0 && connectionList[activeIdentifier].PacketCount == 1 {
			dataChan <- &DownloadCount{
				Identifier: activeIdentifier,
				Count:      connectionList[activeIdentifier].TotalBytes,
				Time:       connectionList[activeIdentifier].Opened,
			}
			delete(connectionList, activeIdentifier)
			connectionList[activeIdentifier] = &Connection{
				Opened: packet.Metadata().Timestamp,
			}
		}

		connectionList[activeIdentifier].TotalBytes = connectionList[activeIdentifier].TotalBytes + int64(packet.Metadata().Length)
		connectionList[activeIdentifier].PacketCount++
		if ack == 0 || rst || fin {
			if syn {
				continue
				// we don't want to close on syn
			}
			dataChan <- &DownloadCount{
				Identifier: activeIdentifier,
				Count:      connectionList[activeIdentifier].TotalBytes,
				Time:       connectionList[activeIdentifier].Opened,
			}
			delete(connectionList, activeIdentifier)
		}
	}
}
