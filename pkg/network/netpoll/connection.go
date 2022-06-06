/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package netpoll

import (
	"errors"
	"io"

	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/netpoll"
)

type Conn struct {
	network.Conn
}

func (c *Conn) Peek(n int) (b []byte, err error) {
	b, err = c.Conn.Peek(n)
	err = normalizeErr(err)
	return
}

func (c *Conn) Skip(n int) error {
	return c.Conn.Skip(n)
}

func (c *Conn) Release() error {
	return c.Conn.Release()
}

func (c *Conn) Len() int {
	return c.Conn.Len()
}

func (c *Conn) ReadByte() (b byte, err error) {
	b, err = c.Conn.ReadByte()
	err = normalizeErr(err)
	return
}

func (c *Conn) ReadBinary(n int) (b []byte, err error) {
	b, err = c.Conn.ReadBinary(n)
	err = normalizeErr(err)
	return
}

func (c *Conn) Malloc(n int) (buf []byte, err error) {
	return c.Conn.Malloc(n)
}

func (c *Conn) WriteBinary(b []byte) (n int, err error) {
	return c.Conn.WriteBinary(b)
}

func (c *Conn) Flush() error {
	return c.Conn.Flush()
}

func normalizeErr(err error) error {
	if errors.Is(err, netpoll.ErrEOF) {
		return io.EOF
	}

	return err
}

func newConn(c netpoll.Connection) network.Conn {
	return &Conn{Conn: c.(network.Conn)}
}