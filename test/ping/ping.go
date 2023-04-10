package ping

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	ICMP_ECHO_REQUEST = 8
	DEFAULT_TIMEOUT   = 100000
)

type ICMPMessage struct {
	Type       uint8
	Code       uint8
	Checksum   uint16
	Identifier uint16
	Sequence   uint16
	Data       []byte
}

func checksum(data []byte) uint16 {
	var sum uint32

	for i := 0; i < len(data)-1; i += 2 {
		sum += uint32(data[i])<<8 | uint32(data[i+1])
	}

	if len(data)%2 != 0 {
		sum += uint32(data[len(data)-1]) << 8
	}

	for sum > 0xffff {
		sum = (sum >> 16) + (sum & 0xffff)
	}

	return ^uint16(sum)
}

func sendPingRequest(conn *net.IPConn, icmpMessage *ICMPMessage) error {
	buf := make([]byte, 8+len(icmpMessage.Data))
	binary.BigEndian.PutUint16(buf[0:2], uint16(icmpMessage.Type)<<8|uint16(icmpMessage.Code))
	binary.BigEndian.PutUint16(buf[2:4], icmpMessage.Checksum)
	binary.BigEndian.PutUint16(buf[4:6], icmpMessage.Identifier)
	binary.BigEndian.PutUint16(buf[6:8], icmpMessage.Sequence)
	copy(buf[8:], icmpMessage.Data)

	icmpMessage.Checksum = checksum(buf)
	binary.BigEndian.PutUint16(buf[2:4], icmpMessage.Checksum)

	_, err := conn.Write(buf)
	return err
}

func receivePingReply(conn *net.IPConn, sendTime time.Time) (*ICMPMessage, time.Duration, error) {
	buf := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(DEFAULT_TIMEOUT))
	bytesRead, err := conn.Read(buf)
	if err != nil {
		return nil, 0, err
	}

	receiveTime := time.Now()
	rtt := receiveTime.Sub(sendTime)

	icmpMessage := &ICMPMessage{
		Type:       buf[0],
		Code:       buf[1],
		Checksum:   binary.BigEndian.Uint16(buf[2:4]),
		Identifier: binary.BigEndian.Uint16(buf[4:6]),
		Sequence:   binary.BigEndian.Uint16(buf[6:8]),
		Data:       buf[8:bytesRead],
	}

	return icmpMessage, rtt, nil
}

func Ping(hostname string, timeout time.Duration, count int) error {
	addr, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		return err
	}

	conn, err := net.DialIP("ip4:icmp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	icmpMessage := &ICMPMessage{
		Type:       ICMP_ECHO_REQUEST,
		Code:       0,
		Identifier: uint16(os.Getpid() & 0xffff),
		Sequence:   0,
		Data:       []byte("PING"),
	}

	packetsSent := 0
	packetsReceived := 0

	for i := 0; i < count; i++ {
		icmpMessage.Sequence = uint16(i)
		err = sendPingRequest(conn, icmpMessage)
		if err != nil {
			return err
		}
		packetsSent++

		reply, rtt, err := receivePingReply(conn, time.Now())
		if err == nil {
			fmt.Printf("Received reply from %s: icmp_seq=%d time=%v\n", addr.String(), reply.Sequence, rtt)
			packetsReceived++
		} else {
			fmt.Printf("Request timeout for icmp_seq %d\n", icmpMessage.Sequence)
		}

		time.Sleep(timeout)
	}

	fmt.Printf("\n--- %s ping statistics ---\n", hostname)
	fmt.Printf("%d packets transmitted, %d packets received\n", packetsSent, packetsReceived)

	return nil
}
