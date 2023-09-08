package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn    io.ReadWriteCloser
	buf     *bufio.Writer
	decoder *gob.Decoder
	encoder *gob.Encoder
}

var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn:    conn,
		buf:     buf,
		decoder: gob.NewDecoder(conn),
		encoder: gob.NewEncoder(buf),
	}
}

func (g *GobCodec) ReadHeader(h *Header) error {
	return g.decoder.Decode(h)
}

func (g *GobCodec) ReadBody(body interface{}) error {
	return g.decoder.Decode(body)
}

func (g *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()
	if err = g.encoder.Encode(h); err != nil {
		log.Println("rpc: gob encode header error:", err)
		return
	}
	if err = g.encoder.Encode(body); err != nil {
		log.Println("rpc: gob encode body error:", err)
		return
	}
	return
}

func (g *GobCodec) Close() error {
	return g.conn.Close()
}
