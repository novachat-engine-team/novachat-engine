package stickerset

import (
	"fmt"
	"novachat_engine/mtproto"
	sfsService "novachat_engine/pkg/cmd/sfs/rpc_client"
	data_fs "novachat_engine/service/data/fs"
)

type InputStickerSetType int32

func (d InputStickerSetType) ToInt32() int32 {
	return int32(d)
}

//  + TL_inputStickerSetEmpty
//  + TL_inputStickerSetID
//  + TL_inputStickerSetShortName
//  + TL_inputStickerSetAnimatedEmoji
//  + TL_inputStickerSetDice

const (
	InputStickerSetEmpty         InputStickerSetType = 0
	InputStickerSetID            InputStickerSetType = 1
	InputStickerSetShortName     InputStickerSetType = 2
	InputStickerSetAnimatedEmoji InputStickerSetType = 3
	InputStickerSetDice          InputStickerSetType = 4
)

func MtInputStickerSet2DataStickerSet(inputStickerSet *mtproto.InputStickerSet) *data_fs.StickerSet {
	switch inputStickerSet.ClassName {
	case mtproto.ClassInputStickerSetEmpty:
		return &data_fs.StickerSet{
			StickerSetType: InputStickerSetEmpty.ToInt32(),
		}
	case mtproto.ClassInputStickerSetID:
		return &data_fs.StickerSet{
			StickerSetType: InputStickerSetID.ToInt32(),
			StickerSetId:   inputStickerSet.GetId(),
			AccessHash:     inputStickerSet.GetAccessHash(),
		}
	case mtproto.ClassInputStickerSetShortName:
		return &data_fs.StickerSet{
			StickerSetType: InputStickerSetShortName.ToInt32(),
			ShortName:      inputStickerSet.GetShortName(),
		}
	case mtproto.ClassInputStickerSetAnimatedEmoji:
		return &data_fs.StickerSet{
			StickerSetType: InputStickerSetAnimatedEmoji.ToInt32(),
		}
	case mtproto.ClassInputStickerSetDice:
		return &data_fs.StickerSet{
			StickerSetType: InputStickerSetDice.ToInt32(),
			Emoticon:       inputStickerSet.GetEmoticon(),
		}
	default:
		panic(fmt.Errorf("InputStickerSet invalid type:%v", inputStickerSet.ClassName))
	}
}

func DataStickerSet2MtInputStickerSet(stickerSet *data_fs.StickerSet) *mtproto.InputStickerSet {
	switch stickerSet.StickerSetType {
	case InputStickerSetEmpty.ToInt32():
		return mtproto.NewTLInputStickerSetEmpty(nil).To_InputStickerSet()
	case InputStickerSetID.ToInt32():
		return mtproto.NewTLInputStickerSetID(&mtproto.InputStickerSet{
			Id:         stickerSet.StickerSetId,
			AccessHash: stickerSet.AccessHash,
		}).To_InputStickerSet()
	case InputStickerSetShortName.ToInt32():
		return mtproto.NewTLInputStickerSetShortName(&mtproto.InputStickerSet{
			ShortName: stickerSet.ShortName,
		}).To_InputStickerSet()
	case InputStickerSetAnimatedEmoji.ToInt32():
		return mtproto.NewTLInputStickerSetAnimatedEmoji(nil).To_InputStickerSet()

	case InputStickerSetDice.ToInt32():
		return mtproto.NewTLInputStickerSetDice(&mtproto.InputStickerSet{
			Emoticon: stickerSet.Emoticon,
		}).To_InputStickerSet()
	default:
		panic(fmt.Errorf("stickerSet invalid type:%v", stickerSet.GetStickerSetType()))
	}
}

func MtInputStickerSet2StickerSet(inputStickerSet *mtproto.InputStickerSet) *sfsService.StickerSet {
	switch inputStickerSet.ClassName {
	case mtproto.ClassInputStickerSetEmpty:
		return &sfsService.StickerSet{
			StickerSetType: InputStickerSetEmpty.ToInt32(),
		}
	case mtproto.ClassInputStickerSetID:
		return &sfsService.StickerSet{
			StickerSetType: InputStickerSetID.ToInt32(),
			StickerSetId:   inputStickerSet.GetId(),
			AccessHash:     inputStickerSet.GetAccessHash(),
		}
	case mtproto.ClassInputStickerSetShortName:
		return &sfsService.StickerSet{
			StickerSetType: InputStickerSetShortName.ToInt32(),
			ShortName:      inputStickerSet.GetShortName(),
		}
	case mtproto.ClassInputStickerSetAnimatedEmoji:
		return &sfsService.StickerSet{
			StickerSetType: InputStickerSetAnimatedEmoji.ToInt32(),
		}
	case mtproto.ClassInputStickerSetDice:
		inputStickerSet.To_InputStickerSetDice()
		return &sfsService.StickerSet{
			StickerSetType: InputStickerSetDice.ToInt32(),
			Emoticon:       inputStickerSet.GetEmoticon(),
		}
	default:
		panic(fmt.Errorf("InputStickerSet invalid type:%v", inputStickerSet.ClassName))
	}
}

func StickerSet2MtInputStickerSet(stickerSet *sfsService.StickerSet) *mtproto.InputStickerSet {
	switch stickerSet.StickerSetType {
	case InputStickerSetEmpty.ToInt32():
		return mtproto.NewTLInputStickerSetEmpty(nil).To_InputStickerSet()
	case InputStickerSetID.ToInt32():
		return mtproto.NewTLInputStickerSetID(&mtproto.InputStickerSet{
			Id:         stickerSet.StickerSetId,
			AccessHash: stickerSet.AccessHash,
		}).To_InputStickerSet()
	case InputStickerSetShortName.ToInt32():
		return mtproto.NewTLInputStickerSetShortName(&mtproto.InputStickerSet{
			ShortName: stickerSet.ShortName,
		}).To_InputStickerSet()
	case InputStickerSetAnimatedEmoji.ToInt32():
		return mtproto.NewTLInputStickerSetAnimatedEmoji(nil).To_InputStickerSet()

	case InputStickerSetDice.ToInt32():
		return mtproto.NewTLInputStickerSetDice(&mtproto.InputStickerSet{
			Emoticon: stickerSet.Emoticon,
		}).To_InputStickerSet()
	default:
		panic(fmt.Errorf("stickerSet invalid type:%v", stickerSet.GetStickerSetType()))
	}
}

func StickerSet2DataStickerSet(stickerSet *sfsService.StickerSet) *data_fs.StickerSet {
	return &data_fs.StickerSet{
		StickerSetType: stickerSet.StickerSetType,
		StickerSetId:   stickerSet.StickerSetId,
		AccessHash:     stickerSet.AccessHash,
		ShortName:      stickerSet.ShortName,
		Emoticon:       stickerSet.Emoticon,
	}
}

func DataStickerSet2StickerSet(stickerSet *data_fs.StickerSet) *sfsService.StickerSet {
	return &sfsService.StickerSet{
		StickerSetType: stickerSet.StickerSetType,
		StickerSetId:   stickerSet.StickerSetId,
		AccessHash:     stickerSet.AccessHash,
		ShortName:      stickerSet.ShortName,
		Emoticon:       stickerSet.Emoticon,
	}
}
