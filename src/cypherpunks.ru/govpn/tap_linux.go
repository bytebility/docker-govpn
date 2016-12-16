// +build linux

/*
GoVPN -- simple secure free software virtual private network daemon
Copyright (C) 2014-2016 Sergey Matveev <stargrave@stargrave.org>
*/

package govpn

import (
	"io"

	"github.com/bigeagle/water"
)

func newTAPer(ifaceName string) (io.ReadWriter, error) {
	return water.NewTAP(ifaceName)
}
