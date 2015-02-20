// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package myservice

import (
	"fmt"
	"github.com/koofr/thrift/lib/go/thrift"
	"math"
)

// (needed to ensure safety because of naive import list construction.)
var _ = math.MinInt32
var _ = thrift.ZERO
var _ = fmt.Printf

type MyService interface {
	Ping() (err error)
	// Parameters:
	//  - Id
	//  - Req
	GetResult(id UUID, req *MyRequest) (r *MyResult, me *MyException, err error)
}

type MyServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewMyServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MyServiceClient {
	return &MyServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewMyServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MyServiceClient {
	return &MyServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *MyServiceClient) Ping() (err error) {
	if err = p.sendPing(); err != nil {
		return
	}
	return p.recvPing()
}

func (p *MyServiceClient) sendPing() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("ping", thrift.CALL, p.SeqId)
	args0 := NewPingArgs()
	err = args0.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *MyServiceClient) recvPing() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result1 := NewPingResult()
	err = result1.Read(iprot)
	iprot.ReadMessageEnd()
	return
}

// Parameters:
//  - Id
//  - Req
func (p *MyServiceClient) GetResult(id UUID, req *MyRequest) (r *MyResult, me *MyException, err error) {
	if err = p.sendGetResult(id, req); err != nil {
		return
	}
	return p.recvGetResult()
}

func (p *MyServiceClient) sendGetResult(id UUID, req *MyRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	oprot.WriteMessageBegin("get_result", thrift.CALL, p.SeqId)
	args4 := NewGetResultArgs()
	args4.Id = id
	args4.Req = req
	err = args4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return
}

func (p *MyServiceClient) recvGetResult() (value *MyResult, me *MyException, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	result5 := NewGetResultResult()
	err = result5.Read(iprot)
	iprot.ReadMessageEnd()
	value = result5.Success
	if result5.Me != nil {
		me = result5.Me
	}
	return
}

type MyServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      MyService
}

func (p *MyServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *MyServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *MyServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewMyServiceProcessor(handler MyService) *MyServiceProcessor {

	self8 := &MyServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["ping"] = &myServiceProcessorPing{handler: handler}
	self8.processorMap["get_result"] = &myServiceProcessorGetResult{handler: handler}
	return self8
}

func (p *MyServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x9.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x9

}

type myServiceProcessorPing struct {
	handler MyService
}

func (p *myServiceProcessorPing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewPingArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewPingResult()
	if err = p.handler.Ping(); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: "+err.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("ping", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type myServiceProcessorGetResult struct {
	handler MyService
}

func (p *myServiceProcessorGetResult) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := NewGetResultArgs()
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("get_result", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	iprot.ReadMessageEnd()
	result := NewGetResultResult()
	if result.Success, result.Me, err = p.handler.GetResult(args.Id, args.Req); err != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_result: "+err.Error())
		oprot.WriteMessageBegin("get_result", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return
	}
	if err2 := oprot.WriteMessageBegin("get_result", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 := result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 := oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type PingArgs struct {
}

func NewPingArgs() *PingArgs {
	return &PingArgs{}
}

func (p *PingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingArgs(%+v)", *p)
}

type PingResult struct {
}

func NewPingResult() *PingResult {
	return &PingResult{}
}

func (p *PingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *PingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingResult(%+v)", *p)
}

type GetResultArgs struct {
	Id  UUID       `thrift:"id,1"`
	Req *MyRequest `thrift:"req,2"`
}

func NewGetResultArgs() *GetResultArgs {
	return &GetResultArgs{}
}

func (p *GetResultArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetResultArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s")
	} else {
		p.Id = UUID(v)
	}
	return nil
}

func (p *GetResultArgs) readField2(iprot thrift.TProtocol) error {
	p.Req = NewMyRequest()
	if err := p.Req.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Req)
	}
	return nil
}

func (p *GetResultArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("get_result_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GetResultArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:id: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Id)); err != nil {
		return fmt.Errorf("%T.id (1) field write error: %s", p)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:id: %s", p, err)
	}
	return err
}

func (p *GetResultArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if p.Req != nil {
		if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:req: %s", p, err)
		}
		if err := p.Req.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Req)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:req: %s", p, err)
		}
	}
	return err
}

func (p *GetResultArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetResultArgs(%+v)", *p)
}

type GetResultResult struct {
	Success *MyResult    `thrift:"success,0"`
	Me      *MyException `thrift:"me,1"`
}

func NewGetResultResult() *GetResultResult {
	return &GetResultResult{}
}

func (p *GetResultResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GetResultResult) readField0(iprot thrift.TProtocol) error {
	p.Success = NewMyResult()
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success)
	}
	return nil
}

func (p *GetResultResult) readField1(iprot thrift.TProtocol) error {
	p.Me = NewMyException()
	if err := p.Me.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Me)
	}
	return nil
}

func (p *GetResultResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("get_result_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	switch {
	case p.Me != nil:
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	default:
		if err := p.writeField0(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("%T write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("%T write struct stop error: %s", err)
	}
	return nil
}

func (p *GetResultResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.Success != nil {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *GetResultResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.Me != nil {
		if err := oprot.WriteFieldBegin("me", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:me: %s", p, err)
		}
		if err := p.Me.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Me)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:me: %s", p, err)
		}
	}
	return err
}

func (p *GetResultResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetResultResult(%+v)", *p)
}
