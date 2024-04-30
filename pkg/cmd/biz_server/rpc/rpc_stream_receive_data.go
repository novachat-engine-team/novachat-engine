package rpc

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	metadata2 "google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"io"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/rpc/rpc"
	"novachat_engine/pkg/runtime"
	"reflect"
)

func (m *Impl) ReceiveData(stream rpc.RpcStreamService_ReceiveDataServer) error {
	var address string
	var done = make(chan bool)
	var object mtproto.TLObject

	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		md, ok1 := metadata2.FromIncomingContext(stream.Context())
		if ok1 {
			address = md.Get(metadata.RpcMetadataAddressKey)[0]
		}
	} else {
		address = peer.Addr.String()
	}
	_ = address

	q, ok := m.loadOrStore(address)
	runtime.GoWithRecover(func() {
		for {
			select {
			case _, ok = <-done:
				if !ok {
					return
				}
			case msg, ok1 := <-q:
				if !ok1 {
					return
				}
				err := stream.Send(msg)
				if err != nil {
					return
				}
			}
		}
	}, nil)

	for {
		select {
		case <-stream.Context().Done():
			close(done)
			close(q)
			m.delete(address)
			return stream.Context().Err()

		default:
			recvData, err := stream.Recv()
			if err == io.EOF {
				close(done)
				close(q)
				m.delete(address)
				return nil
			}
			if err != nil {
				close(done)
				close(q)
				m.delete(address)
				return err
			}

			messageName, _ := types.AnyMessageName(recvData.GetData())
			typ := proto.MessageType(messageName)
			for {
				if typ.Kind() == reflect.Ptr {
					typ = typ.Elem()
				} else {
					break
				}
			}
			value := reflect.New(typ)
			object, _ = value.Interface().(mtproto.TLObject)
			if object == nil {
				log.Errorf("ReceiveDataStream messageName:%s recv data:%+v", messageName, recvData.Data)
				continue
			}
			err = types.UnmarshalAny(recvData.Data, object.(proto.Message))
			if err != nil {
				log.Errorf("ReceiveDataStream recv data:%+v", recvData.Data)
				continue
			}

			objectName := rpc.GetObjectTypeName(object)
			implValue, ok1 := rpc.GetRpcImplValue(objectName)
			if !ok1 {
				log.Errorf("ReceiveDataStream GetRpcImplValue objectName:%s", objectName)
				continue
			}

			runtime.GoSafeWithRecover(func(param []interface{}) {
				localImplValue, _ := param[0].(*rpc.ClientImpl)
				localRecvData, _ := param[1].(*rpc.RpcStreamData)
				localObject, _ := param[2].(mtproto.TLObject)
				//log.Debugf("call %+v", localImplValue)
				outValue := localImplValue.RpcMethodValue.
					Call([]reflect.Value{
						reflect.ValueOf(metadata.NewIncomingContext(context.Background(), localRecvData.Md)),
						reflect.ValueOf(localObject)})

				respRpcData := m.rpcDataPool.Get().(*rpc.RpcStreamData)
				defer m.rpcDataPool.Put(respRpcData)

				respRpcData.Md = localRecvData.Md
				resp, errResp := outValue[0].Interface(), outValue[1].Interface()
				if errResp != nil {
					_, ok = errResp.(*mtproto.Error)
					if !ok {
						log.Errorf("ReceiveDataStream GetRpcImplValue objectName:%s error:%s", objectName, errResp.(error).Error())
						return
					} else {
						respRpcData.Data, err = types.MarshalAny(errResp.(proto.Message))
						if err != nil {
							log.Errorf("ReceiveDataStream GetRpcImplValue objectName:%s errResp:`%+v` MarshalAny error:%s", objectName, errResp, err.Error())
							return
						}
					}
				} else {
					//log.Debugf("ReceiveDataStream reqMsgId:%d request:%+v resp:%+v", localRecvData.Md.ReqMsgId, localObject, resp)
					respRpcData.Data, err = types.MarshalAny(resp.(proto.Message))
					if err != nil {
						log.Errorf("ReceiveDataStream GetRpcImplValue objectName:%s resp:`%+v` MarshalAny error:%s", objectName, resp, err.Error())
						return
					}
				}

				q <- respRpcData
			}, []interface{}{implValue, recvData, object}, func(r interface{}) {
				log.Errorf("ReceiveDataStream GetRpcImplValue objectName:%s error:%+v", objectName, r)
			})
		}
	}
	return nil
}
