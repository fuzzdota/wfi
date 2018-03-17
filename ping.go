package wfi

import (
	"net"
	"time"
)

// Ping it
func Ping(host string, duration time.Duration) error {
	// TODO: Support all types, not just tcp
	// "tcp",
	// "tcp4" (IPv4-only),
	// "tcp6" (IPv6-only),
	// "udp",
	// "udp4" (IPv4-only),
	// "udp6" (IPv6-only),
	// "ip",
	// "ip4" (IPv4-only),
	// "ip6" (IPv6-only),
	// "unix",
	// "unixgram"
	// "unixpacket"
	conn, err := net.DialTimeout("tcp", host, duration)
	if err == nil {
		return conn.Close()
	}
	return err
}
