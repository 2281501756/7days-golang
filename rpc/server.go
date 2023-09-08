package geerpc

import (
	"encoding/json"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumber = 0x1337babe

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{}

var DefaultServer = Server{}

func NewServer() *Server {
	return &DefaultServer
}

func (s *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() {
		_ = conn.Close()
	}()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: option parse error")
		return
	}
	if opt.MagicNumber != DefaultOption.MagicNumber {
		log.Println("rpc server: MagicNumber not match", opt.MagicNumber, DefaultOption.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Println("rpc server: codec not support", opt.CodecType)
		return
	}
	s.serverCodec(f(conn))
}

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
}

var invalidRequest = struct{}{}

func (s *Server) serverCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	group := new(sync.WaitGroup)
	for {
		req, err := s.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			s.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		group.Add(1)
		go s.HandleResponse(cc, req, sending, group)
	}
	group.Wait()
	_ = cc.Close()
}

func (s *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		return nil, err
	}
	return &h, nil
}

func (s *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := s.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: h}
	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		return nil, err
	}
	return req, nil
}

func (s *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response failed", err)
	}
	log.Println("server send", body)
}
func (s *Server) HandleResponse(cc codec.Codec, req *request, lock *sync.Mutex, group *sync.WaitGroup) {
	defer group.Done()
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq))
	s.sendResponse(cc, req.h, req.replyv.Interface(), lock)
}

func (s *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error", err)
			return
		}
		go s.ServeConn(conn)
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}
