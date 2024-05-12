package main

import "fmt"

type TCPStreamInfo struct {
	Protocol string
}

type UDPStreamInfo struct {
	Protocol string
}

func appendStreams[T any](streamsMap map[string]T) []T {
	var streams []T
	for _, stats := range streamsMap {
		streams = append(streams, stats)
	}
	return streams
}

func main() {
	tcpStreamsMap := make(map[string]*TCPStreamInfo)
	udpStreamsMap := make(map[string]*UDPStreamInfo)

	tcpStreams := appendStreams(tcpStreamsMap)
	udpStreams := appendStreams(udpStreamsMap)
	// 不用泛型的话，就会像这样写两遍
	// var tcpStreams []*TCPStreamInfo
	// for _, stats := range tcpStreamsMap {
	// 	tcpStreams = append(tcpStreams, stats)
	// }

	// var udpStreams []*UDPStreamInfo
	// for _, stats := range udpStreamsMap {
	// 	udpStreams = append(udpStreams, stats)
	// }

	fmt.Println(tcpStreams, udpStreams)
}
