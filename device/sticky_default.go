//go:build !linux

package device

import (
	"github.com/pccr10001/wireguard-go/conn"
	"github.com/pccr10001/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
