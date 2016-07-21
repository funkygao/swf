package server

import (
	"net"

	"golang.org/x/net/context"
)

type key int

const remoteAddrKey key = 0

func NewContext(ctx context.Context, conn net.Conn) context.Context {
	return context.WithValue(ctx, remoteAddrKey, conn)
}

func FromContext(ctx context.Context) (conn net.Conn, ok bool) {
	conn, ok = ctx.Value(remoteAddrKey).(net.Conn)
	return
}
