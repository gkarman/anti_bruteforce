package valueobject

import (
	"fmt"
	"net"
)

type CIDR struct {
	value *net.IPNet
}

func NewCIDR(cidrStr string) (*CIDR, error) {
	_, ipNet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR format: %w", err)
	}
	return &CIDR{value: ipNet}, nil
}

func (c *CIDR) Contains(ip net.IP) bool {
	return c.value.Contains(ip)
}

func (c *CIDR) String() string {
	return c.value.String()
}
