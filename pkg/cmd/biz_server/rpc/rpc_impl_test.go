package rpc

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"novachat_engine/mtproto"
	help "novachat_engine/pkg/cmd/biz_server/rpc/help"
	"novachat_engine/pkg/log"
	"novachat_engine/pkg/rpc/metadata"
	"novachat_engine/pkg/rpc/rpc"
	"reflect"
	"testing"
)

type TestImpl struct {
}

func (t *TestImpl) Test() string {
	helloMessage := "I'm TestImpl::Test"
	fmt.Println(helloMessage)
	return helloMessage
}

func Test_TestImpl(t *testing.T) {
	impl := &TestImpl{}

	value := reflect.ValueOf(impl)
	t.Log(value.Type().String())

	testFuncValue := value.MethodByName("Test")
	t.Log(testFuncValue.Call(nil))

	op := log.DefaultOptions()
	log.Configure(op)
	helpServiceImpl := &help.HelpServiceImpl{}

	rpc.RegisterRpcImpl(mtproto.RPCHelpService_serviceDesc, helpServiceImpl)

	object1 := &mtproto.TLHelpAcceptTermsOfService{}
	recvData, _ := types.MarshalAny(object1)
	messageName, _ := types.AnyMessageName(recvData)
	typ := proto.MessageType(messageName)
	for {
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		} else {
			break
		}
	}
	value111 := reflect.New(typ)
	object, _ := value111.Interface().(mtproto.TLObject)
	objectName := rpc.GetObjectTypeName(object)
	cc, ok := rpc.GetRpcImplValue(objectName)
	if !ok {
		t.Fatalf("not found TLHelpAcceptTermsOfService")
	}

	vv := reflect.ValueOf(helpServiceImpl)
	mbn := vv.MethodByName("HelpAcceptTermsOfService")

	_ = mbn

	outValue := cc.RpcMethodValue.
		Call([]reflect.Value{
			reflect.ValueOf(metadata.NewIncomingContext(context.Background(), &metadata.RpcMetaData{
				AuthKeyId: 1111,
				SessionId: 2222,
			})),
			reflect.ValueOf(object)})
	resp, errResp := outValue[0].Interface(), outValue[1].Interface()
	if errResp != nil {
		t.Fatalf(errResp.(error).Error())
	}
	resp11 := resp.(proto.Message)
	t.Logf("%+v", resp11)
}
