/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/3/31 16:29
 * @File : meta_data.go
 */

package metadata

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"
	"novachat_engine/pkg/log"
)

const rpcMetadata = "keyRpcMetaData"
const RpcMetadataKey = rpcMetadata
const RpcMetadataAddressKey = rpcMetadata + "address"

func RpcMetaDataToMD(md *RpcMetaData) metadata.MD {
	b, _ := proto.Marshal(md)
	return metadata.Pairs(rpcMetadata, base64.StdEncoding.EncodeToString(b))
}

func MDToRpcMetaData(md metadata.MD) *RpcMetaData {
	var ret string
	if true {
		val := metautils.NiceMD(md).Get(rpcMetadata)
		ret = val
	} else {
		val := md.Get(rpcMetadata)
		if len(val) != 0 {
			ret = val[0]
		}
	}
	if len(ret) == 0 {
		return nil
	}

	buf, err := base64.StdEncoding.DecodeString(ret)
	v := &RpcMetaData{}

	err = proto.Unmarshal(buf, v)
	if err != nil {
		log.Errorf("MDToRpcMetaData unmarshall:%s", ret)
		return nil
	}

	return v
}

func NewContextWithMetaData(ctx context.Context, md *RpcMetaData) context.Context {
	return metadata.NewOutgoingContext(ctx, RpcMetaDataToMD(md))
	//return metadata.NewIncomingContext(ctx, RpcMetaDataToMD(md))
}

func NewIncomingContext(ctx context.Context, md *RpcMetaData) context.Context {
	return metadata.NewIncomingContext(ctx, RpcMetaDataToMD(md))
}

func RpcMetaDataFromContext(ctx context.Context) *RpcMetaData {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		log.Errorf("RpcMetaDataFromIncoming error:ctx =%v", ctx)
		return nil
	}

	return MDToRpcMetaData(md)
}

func RpcMetaDataDebug(md *RpcMetaData) string {
	if md == nil {
		return ""
	}
	debugMd, ok := proto.Clone(md).(*RpcMetaData)
	if ok == false {
		return ""
	}
	return fmt.Sprintf("%v", debugMd)
}
