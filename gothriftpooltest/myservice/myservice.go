// Autogenerated by Thrift Compiler (0.13.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package myservice

import(
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

type UUID string

func UUIDPtr(v UUID) *UUID { return &v }

// Attributes:
//  - Req
type MyRequest struct {
  Req string `thrift:"req,1,required" db:"req" json:"req"`
}

func NewMyRequest() *MyRequest {
  return &MyRequest{}
}


func (p *MyRequest) GetReq() string {
  return p.Req
}
func (p *MyRequest) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetReq bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
        issetReq = true
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetReq{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Req is not set"));
  }
  return nil
}

func (p *MyRequest)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Req = v
}
  return nil
}

func (p *MyRequest) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("MyRequest"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyRequest) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err) }
  if err := oprot.WriteString(string(p.Req)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.req (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err) }
  return err
}

func (p *MyRequest) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyRequest(%+v)", *p)
}

// Attributes:
//  - Res
type MyResponse struct {
  Res string `thrift:"res,1,required" db:"res" json:"res"`
}

func NewMyResponse() *MyResponse {
  return &MyResponse{}
}


func (p *MyResponse) GetRes() string {
  return p.Res
}
func (p *MyResponse) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetRes bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
        issetRes = true
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetRes{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Res is not set"));
  }
  return nil
}

func (p *MyResponse)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Res = v
}
  return nil
}

func (p *MyResponse) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("MyResponse"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyResponse) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("res", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:res: ", p), err) }
  if err := oprot.WriteString(string(p.Res)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.res (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:res: ", p), err) }
  return err
}

func (p *MyResponse) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyResponse(%+v)", *p)
}

type MyException struct {
}

func NewMyException() *MyException {
  return &MyException{}
}

func (p *MyException) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyException) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("MyException"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyException) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyException(%+v)", *p)
}

func (p *MyException) Error() string {
  return p.String()
}

type MyService interface {
  Ping(ctx context.Context) (err error)
  // Parameters:
  //  - ID
  //  - Req
  GetResponse(ctx context.Context, id UUID, req *MyRequest) (r *MyResponse, err error)
}

type MyServiceClient struct {
  c thrift.TClient
}

func NewMyServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MyServiceClient {
  return &MyServiceClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewMyServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MyServiceClient {
  return &MyServiceClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewMyServiceClient(c thrift.TClient) *MyServiceClient {
  return &MyServiceClient{
    c: c,
  }
}

func (p *MyServiceClient) Client_() thrift.TClient {
  return p.c
}
func (p *MyServiceClient) Ping(ctx context.Context) (err error) {
  var _args0 MyServicePingArgs
  var _result1 MyServicePingResult
  if err = p.Client_().Call(ctx, "ping", &_args0, &_result1); err != nil {
    return
  }
  return nil
}

// Parameters:
//  - ID
//  - Req
func (p *MyServiceClient) GetResponse(ctx context.Context, id UUID, req *MyRequest) (r *MyResponse, err error) {
  var _args2 MyServiceGetResponseArgs
  _args2.ID = id
  _args2.Req = req
  var _result3 MyServiceGetResponseResult
  if err = p.Client_().Call(ctx, "get_response", &_args2, &_result3); err != nil {
    return
  }
  switch {
  case _result3.Me!= nil:
    return r, _result3.Me
  }

  return _result3.GetSuccess(), nil
}

type MyServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler MyService
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

  self4 := &MyServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self4.processorMap["ping"] = &myServiceProcessorPing{handler:handler}
  self4.processorMap["get_response"] = &myServiceProcessorGetResponse{handler:handler}
return self4
}

func (p *MyServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x5.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x5

}

type myServiceProcessorPing struct {
  handler MyService
}

func (p *myServiceProcessorPing) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := MyServicePingArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := MyServicePingResult{}
  var err2 error
  if err2 = p.handler.Ping(ctx); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: " + err2.Error())
    oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  if err2 = oprot.WriteMessageBegin("ping", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}

type myServiceProcessorGetResponse struct {
  handler MyService
}

func (p *myServiceProcessorGetResponse) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := MyServiceGetResponseArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("get_response", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := MyServiceGetResponseResult{}
var retval *MyResponse
  var err2 error
  if retval, err2 = p.handler.GetResponse(ctx, args.ID, args.Req); err2 != nil {
  switch v := err2.(type) {
    case *MyException:
  result.Me = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing get_response: " + err2.Error())
    oprot.WriteMessageBegin("get_response", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  }
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("get_response", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

type MyServicePingArgs struct {
}

func NewMyServicePingArgs() *MyServicePingArgs {
  return &MyServicePingArgs{}
}

func (p *MyServicePingArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServicePingArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ping_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServicePingArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServicePingArgs(%+v)", *p)
}

type MyServicePingResult struct {
}

func NewMyServicePingResult() *MyServicePingResult {
  return &MyServicePingResult{}
}

func (p *MyServicePingResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServicePingResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ping_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServicePingResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServicePingResult(%+v)", *p)
}

// Attributes:
//  - ID
//  - Req
type MyServiceGetResponseArgs struct {
  ID UUID `thrift:"id,1" db:"id" json:"id"`
  Req *MyRequest `thrift:"req,2" db:"req" json:"req"`
}

func NewMyServiceGetResponseArgs() *MyServiceGetResponseArgs {
  return &MyServiceGetResponseArgs{}
}


func (p *MyServiceGetResponseArgs) GetID() UUID {
  return p.ID
}
var MyServiceGetResponseArgs_Req_DEFAULT *MyRequest
func (p *MyServiceGetResponseArgs) GetReq() *MyRequest {
  if !p.IsSetReq() {
    return MyServiceGetResponseArgs_Req_DEFAULT
  }
return p.Req
}
func (p *MyServiceGetResponseArgs) IsSetReq() bool {
  return p.Req != nil
}

func (p *MyServiceGetResponseArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServiceGetResponseArgs)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := UUID(v)
  p.ID = temp
}
  return nil
}

func (p *MyServiceGetResponseArgs)  ReadField2(iprot thrift.TProtocol) error {
  p.Req = &MyRequest{}
  if err := p.Req.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
  }
  return nil
}

func (p *MyServiceGetResponseArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_response_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServiceGetResponseArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err) }
  if err := oprot.WriteString(string(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err) }
  return err
}

func (p *MyServiceGetResponseArgs) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("req", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:req: ", p), err) }
  if err := p.Req.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:req: ", p), err) }
  return err
}

func (p *MyServiceGetResponseArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServiceGetResponseArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Me
type MyServiceGetResponseResult struct {
  Success *MyResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
  Me *MyException `thrift:"me,1" db:"me" json:"me,omitempty"`
}

func NewMyServiceGetResponseResult() *MyServiceGetResponseResult {
  return &MyServiceGetResponseResult{}
}

var MyServiceGetResponseResult_Success_DEFAULT *MyResponse
func (p *MyServiceGetResponseResult) GetSuccess() *MyResponse {
  if !p.IsSetSuccess() {
    return MyServiceGetResponseResult_Success_DEFAULT
  }
return p.Success
}
var MyServiceGetResponseResult_Me_DEFAULT *MyException
func (p *MyServiceGetResponseResult) GetMe() *MyException {
  if !p.IsSetMe() {
    return MyServiceGetResponseResult_Me_DEFAULT
  }
return p.Me
}
func (p *MyServiceGetResponseResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *MyServiceGetResponseResult) IsSetMe() bool {
  return p.Me != nil
}

func (p *MyServiceGetResponseResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
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
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServiceGetResponseResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &MyResponse{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *MyServiceGetResponseResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Me = &MyException{}
  if err := p.Me.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Me), err)
  }
  return nil
}

func (p *MyServiceGetResponseResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("get_response_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServiceGetResponseResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *MyServiceGetResponseResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetMe() {
    if err := oprot.WriteFieldBegin("me", thrift.STRUCT, 1); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:me: ", p), err) }
    if err := p.Me.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Me), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 1:me: ", p), err) }
  }
  return err
}

func (p *MyServiceGetResponseResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServiceGetResponseResult(%+v)", *p)
}


