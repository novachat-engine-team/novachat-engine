/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2023/06/12 18:03
 * @File : codec_schema.service.pb.go
 * @Desc :
 *
 */

package mtproto

import (
	"novachat_engine/pkg/log"
)

// account.getWallPapers
func NewResponse_AccountGetWallPapers() *Response_AccountGetWallPapers {
	return &Response_AccountGetWallPapers{}
}

func (m *Response_AccountGetWallPapers) SetAccountGetWallPapersc04Cfac2(v *Vector_WallPaper) {
	m.AccountGetWallPapersc04Cfac2 = v
}

//func (m *Response_AccountGetWallPapers) GetAccountGetWallPapersc04Cfac2 () *Vector_WallPaper { return m.AccountGetWallPapersc04Cfac2 }

func (m *Response_AccountGetWallPapers) SetAccountGetWallPapersaabb1763(v *Account_WallPapers) {
	m.AccountGetWallPapersaabb1763 = v
}

//func (m *Response_AccountGetWallPapers) GetAccountGetWallPapersaabb1763 () *Account_WallPapers { return m.AccountGetWallPapersaabb1763 }

func (m *Response_AccountGetWallPapers) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_AccountGetWallPapers.ToUInt32(): //71
		return m.AccountGetWallPapersc04Cfac2.Encode(layer)
	case Cmd_AccountGetWallPapersaabb1763.ToUInt32(): //93
		return m.AccountGetWallPapersaabb1763.Encode(layer)

	default:
		log.Errorf("Response_AccountGetWallPapers Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_AccountGetWallPapers) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_AccountGetWallPapers.ToUInt32(): //71
		return m.AccountGetWallPapersc04Cfac2.Decode(dBuf, layer)
	case Cmd_AccountGetWallPapersaabb1763.ToUInt32(): //93
		return m.AccountGetWallPapersaabb1763.Decode(dBuf, layer)

	default:
		log.Errorf("Response_AccountGetWallPapers Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// contacts.deleteContacts
func NewResponse_ContactsDeleteContacts() *Response_ContactsDeleteContacts {
	return &Response_ContactsDeleteContacts{}
}

func (m *Response_ContactsDeleteContacts) SetContactsDeleteContacts59Ab389E(v *Bool) {
	m.ContactsDeleteContacts59Ab389E = v
}

//func (m *Response_ContactsDeleteContacts) GetContactsDeleteContacts59Ab389E () *Bool { return m.ContactsDeleteContacts59Ab389E }

func (m *Response_ContactsDeleteContacts) SetContactsDeleteContacts96A0E00(v *Updates) {
	m.ContactsDeleteContacts96A0E00 = v
}

//func (m *Response_ContactsDeleteContacts) GetContactsDeleteContacts96A0E00 () *Updates { return m.ContactsDeleteContacts96A0E00 }

func (m *Response_ContactsDeleteContacts) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_ContactsDeleteContacts.ToUInt32(): //71
		return m.ContactsDeleteContacts59Ab389E.Encode(layer)
	case Cmd_ContactsDeleteContacts96a0e00.ToUInt32(): //102
		return m.ContactsDeleteContacts96A0E00.Encode(layer)

	default:
		log.Errorf("Response_ContactsDeleteContacts Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_ContactsDeleteContacts) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_ContactsDeleteContacts.ToUInt32(): //71
		return m.ContactsDeleteContacts59Ab389E.Decode(dBuf, layer)
	case Cmd_ContactsDeleteContacts96a0e00.ToUInt32(): //102
		return m.ContactsDeleteContacts96A0E00.Decode(dBuf, layer)

	default:
		log.Errorf("Response_ContactsDeleteContacts Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// help.getAppChangelog
func NewResponse_HelpGetAppChangelog() *Response_HelpGetAppChangelog {
	return &Response_HelpGetAppChangelog{}
}

func (m *Response_HelpGetAppChangelog) SetHelpGetAppChangelog9010Ef6F(v *Updates) {
	m.HelpGetAppChangelog9010Ef6F = v
}

//func (m *Response_HelpGetAppChangelog) GetHelpGetAppChangelog9010Ef6F () *Updates { return m.HelpGetAppChangelog9010Ef6F }

func (m *Response_HelpGetAppChangelog) SetHelpGetAppChangelog5Bab7Fb2(v *Help_AppChangelog) {
	m.HelpGetAppChangelog5Bab7Fb2 = v
}

//func (m *Response_HelpGetAppChangelog) GetHelpGetAppChangelog5Bab7Fb2 () *Help_AppChangelog { return m.HelpGetAppChangelog5Bab7Fb2 }

func (m *Response_HelpGetAppChangelog) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_HelpGetAppChangelog.ToUInt32(): //71
		return m.HelpGetAppChangelog9010Ef6F.Encode(layer)
	case Cmd_HelpGetAppChangelog5bab7fb2.ToUInt32(): //51
		return m.HelpGetAppChangelog5Bab7Fb2.Encode(layer)

	default:
		log.Errorf("Response_HelpGetAppChangelog Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_HelpGetAppChangelog) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_HelpGetAppChangelog.ToUInt32(): //71
		return m.HelpGetAppChangelog9010Ef6F.Decode(dBuf, layer)
	case Cmd_HelpGetAppChangelog5bab7fb2.ToUInt32(): //51
		return m.HelpGetAppChangelog5Bab7Fb2.Decode(dBuf, layer)

	default:
		log.Errorf("Response_HelpGetAppChangelog Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// messages.getMessagesViews
func NewResponse_MessagesGetMessagesViews() *Response_MessagesGetMessagesViews {
	return &Response_MessagesGetMessagesViews{}
}

func (m *Response_MessagesGetMessagesViews) SetMessagesGetMessagesViewsc4C8A55D(v *VectorInt) {
	m.MessagesGetMessagesViewsc4C8A55D = v
}

//func (m *Response_MessagesGetMessagesViews) GetMessagesGetMessagesViewsc4C8A55D () *VectorInt { return m.MessagesGetMessagesViewsc4C8A55D }

func (m *Response_MessagesGetMessagesViews) SetMessagesGetMessagesViews5784D3E1(v *Messages_MessageViews) {
	m.MessagesGetMessagesViews5784D3E1 = v
}

//func (m *Response_MessagesGetMessagesViews) GetMessagesGetMessagesViews5784D3E1 () *Messages_MessageViews { return m.MessagesGetMessagesViews5784D3E1 }

func (m *Response_MessagesGetMessagesViews) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_MessagesGetMessagesViews.ToUInt32(): //71
		return m.MessagesGetMessagesViewsc4C8A55D.Encode(layer)
	case Cmd_MessagesGetMessagesViews5784d3e1.ToUInt32(): //119
		return m.MessagesGetMessagesViews5784D3E1.Encode(layer)

	default:
		log.Errorf("Response_MessagesGetMessagesViews Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_MessagesGetMessagesViews) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_MessagesGetMessagesViews.ToUInt32(): //71
		return m.MessagesGetMessagesViewsc4C8A55D.Decode(dBuf, layer)
	case Cmd_MessagesGetMessagesViews5784d3e1.ToUInt32(): //119
		return m.MessagesGetMessagesViews5784D3E1.Decode(dBuf, layer)

	default:
		log.Errorf("Response_MessagesGetMessagesViews Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// messages.installStickerSet
func NewResponse_MessagesInstallStickerSet() *Response_MessagesInstallStickerSet {
	return &Response_MessagesInstallStickerSet{}
}

func (m *Response_MessagesInstallStickerSet) SetMessagesInstallStickerSetc78Fe460(v *Messages_StickerSetInstallResult) {
	m.MessagesInstallStickerSetc78Fe460 = v
}

//func (m *Response_MessagesInstallStickerSet) GetMessagesInstallStickerSetc78Fe460 () *Messages_StickerSetInstallResult { return m.MessagesInstallStickerSetc78Fe460 }

func (m *Response_MessagesInstallStickerSet) SetMessagesInstallStickerSet7B30C3A6(v *Bool) {
	m.MessagesInstallStickerSet7B30C3A6 = v
}

//func (m *Response_MessagesInstallStickerSet) GetMessagesInstallStickerSet7B30C3A6 () *Bool { return m.MessagesInstallStickerSet7B30C3A6 }

func (m *Response_MessagesInstallStickerSet) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_MessagesInstallStickerSet.ToUInt32(): //71
		return m.MessagesInstallStickerSetc78Fe460.Encode(layer)
	case Cmd_MessagesInstallStickerSet7b30c3a6.ToUInt32(): //51
		return m.MessagesInstallStickerSet7B30C3A6.Encode(layer)

	default:
		log.Errorf("Response_MessagesInstallStickerSet Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_MessagesInstallStickerSet) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_MessagesInstallStickerSet.ToUInt32(): //71
		return m.MessagesInstallStickerSetc78Fe460.Decode(dBuf, layer)
	case Cmd_MessagesInstallStickerSet7b30c3a6.ToUInt32(): //51
		return m.MessagesInstallStickerSet7B30C3A6.Decode(dBuf, layer)

	default:
		log.Errorf("Response_MessagesInstallStickerSet Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// messages.readHistory
func NewResponse_MessagesReadHistory() *Response_MessagesReadHistory {
	return &Response_MessagesReadHistory{}
}

func (m *Response_MessagesReadHistory) SetMessagesReadHistorye306D3A(v *Messages_AffectedMessages) {
	m.MessagesReadHistorye306D3A = v
}

//func (m *Response_MessagesReadHistory) GetMessagesReadHistorye306D3A () *Messages_AffectedMessages { return m.MessagesReadHistorye306D3A }

func (m *Response_MessagesReadHistory) SetMessagesReadHistoryb04F2510(v *Messages_AffectedHistory) {
	m.MessagesReadHistoryb04F2510 = v
}

//func (m *Response_MessagesReadHistory) GetMessagesReadHistoryb04F2510 () *Messages_AffectedHistory { return m.MessagesReadHistoryb04F2510 }

func (m *Response_MessagesReadHistory) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_MessagesReadHistory.ToUInt32(): //71
		return m.MessagesReadHistorye306D3A.Encode(layer)
	case Cmd_MessagesReadHistoryb04f2510.ToUInt32(): //68
		return m.MessagesReadHistoryb04F2510.Encode(layer)

	default:
		log.Errorf("Response_MessagesReadHistory Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_MessagesReadHistory) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_MessagesReadHistory.ToUInt32(): //71
		return m.MessagesReadHistorye306D3A.Decode(dBuf, layer)
	case Cmd_MessagesReadHistoryb04f2510.ToUInt32(): //68
		return m.MessagesReadHistoryb04F2510.Decode(dBuf, layer)

	default:
		log.Errorf("Response_MessagesReadHistory Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// photos.updateProfilePhoto
func NewResponse_PhotosUpdateProfilePhoto() *Response_PhotosUpdateProfilePhoto {
	return &Response_PhotosUpdateProfilePhoto{}
}

func (m *Response_PhotosUpdateProfilePhoto) SetPhotosUpdateProfilePhotof0Bb5152(v *UserProfilePhoto) {
	m.PhotosUpdateProfilePhotof0Bb5152 = v
}

//func (m *Response_PhotosUpdateProfilePhoto) GetPhotosUpdateProfilePhotof0Bb5152 () *UserProfilePhoto { return m.PhotosUpdateProfilePhotof0Bb5152 }

func (m *Response_PhotosUpdateProfilePhoto) SetPhotosUpdateProfilePhoto72D4742C(v *Photos_Photo) {
	m.PhotosUpdateProfilePhoto72D4742C = v
}

//func (m *Response_PhotosUpdateProfilePhoto) GetPhotosUpdateProfilePhoto72D4742C () *Photos_Photo { return m.PhotosUpdateProfilePhoto72D4742C }

func (m *Response_PhotosUpdateProfilePhoto) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_PhotosUpdateProfilePhoto.ToUInt32(): //71
		return m.PhotosUpdateProfilePhotof0Bb5152.Encode(layer)
	case Cmd_PhotosUpdateProfilePhoto72d4742c.ToUInt32(): //117
		return m.PhotosUpdateProfilePhoto72D4742C.Encode(layer)

	default:
		log.Errorf("Response_PhotosUpdateProfilePhoto Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_PhotosUpdateProfilePhoto) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_PhotosUpdateProfilePhoto.ToUInt32(): //71
		return m.PhotosUpdateProfilePhotof0Bb5152.Decode(dBuf, layer)
	case Cmd_PhotosUpdateProfilePhoto72d4742c.ToUInt32(): //117
		return m.PhotosUpdateProfilePhoto72D4742C.Decode(dBuf, layer)

	default:
		log.Errorf("Response_PhotosUpdateProfilePhoto Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// upload.getCdnFileHashes
func NewResponse_UploadGetCdnFileHashes() *Response_UploadGetCdnFileHashes {
	return &Response_UploadGetCdnFileHashes{}
}

func (m *Response_UploadGetCdnFileHashes) SetUploadGetCdnFileHashesf715C87B(v *Vector_CdnFileHash) {
	m.UploadGetCdnFileHashesf715C87B = v
}

//func (m *Response_UploadGetCdnFileHashes) GetUploadGetCdnFileHashesf715C87B () *Vector_CdnFileHash { return m.UploadGetCdnFileHashesf715C87B }

func (m *Response_UploadGetCdnFileHashes) SetUploadGetCdnFileHashes4Da54231(v *Vector_FileHash) {
	m.UploadGetCdnFileHashes4Da54231 = v
}

//func (m *Response_UploadGetCdnFileHashes) GetUploadGetCdnFileHashes4Da54231 () *Vector_FileHash { return m.UploadGetCdnFileHashes4Da54231 }

func (m *Response_UploadGetCdnFileHashes) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_UploadGetCdnFileHashes.ToUInt32(): //71
		return m.UploadGetCdnFileHashesf715C87B.Encode(layer)
	case Cmd_UploadGetCdnFileHashes4da54231.ToUInt32(): //85
		return m.UploadGetCdnFileHashes4Da54231.Encode(layer)

	default:
		log.Errorf("Response_UploadGetCdnFileHashes Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_UploadGetCdnFileHashes) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_UploadGetCdnFileHashes.ToUInt32(): //71
		return m.UploadGetCdnFileHashesf715C87B.Decode(dBuf, layer)
	case Cmd_UploadGetCdnFileHashes4da54231.ToUInt32(): //85
		return m.UploadGetCdnFileHashes4Da54231.Decode(dBuf, layer)

	default:
		log.Errorf("Response_UploadGetCdnFileHashes Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

// upload.reuploadCdnFile
func NewResponse_UploadReuploadCdnFile() *Response_UploadReuploadCdnFile {
	return &Response_UploadReuploadCdnFile{}
}

func (m *Response_UploadReuploadCdnFile) SetUploadReuploadCdnFile1Af91C09(v *Vector_CdnFileHash) {
	m.UploadReuploadCdnFile1Af91C09 = v
}

//func (m *Response_UploadReuploadCdnFile) GetUploadReuploadCdnFile1Af91C09 () *Vector_CdnFileHash { return m.UploadReuploadCdnFile1Af91C09 }

func (m *Response_UploadReuploadCdnFile) SetUploadReuploadCdnFile9B2754A8(v *Vector_FileHash) {
	m.UploadReuploadCdnFile9B2754A8 = v
}

//func (m *Response_UploadReuploadCdnFile) GetUploadReuploadCdnFile9B2754A8 () *Vector_FileHash { return m.UploadReuploadCdnFile9B2754A8 }

func (m *Response_UploadReuploadCdnFile) Encode(layer int32) []byte {
	switch m.Cmd {
	case Cmd_UploadReuploadCdnFile.ToUInt32(): //71
		return m.UploadReuploadCdnFile1Af91C09.Encode(layer)
	case Cmd_UploadReuploadCdnFile9b2754a8.ToUInt32(): //85
		return m.UploadReuploadCdnFile9B2754A8.Encode(layer)

	default:
		log.Errorf("Response_UploadReuploadCdnFile Encode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *Response_UploadReuploadCdnFile) Decode(dBuf *DecodeBuf, layer int32) error {
	switch m.Cmd {
	case Cmd_UploadReuploadCdnFile.ToUInt32(): //71
		return m.UploadReuploadCdnFile1Af91C09.Decode(dBuf, layer)
	case Cmd_UploadReuploadCdnFile9b2754a8.ToUInt32(): //85
		return m.UploadReuploadCdnFile9B2754A8.Decode(dBuf, layer)

	default:
		log.Errorf("Response_UploadReuploadCdnFile Decode Invalid cmd:%x", uint32(m.Cmd))
		return nil
	}
}

func (m *VectorLong) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.VectorLong(m.Datas)
	return x.GetBuf()
}

func (m *VectorLong) Decode(dbuf *DecodeBuf, layer int32) error {
	m.Datas = dbuf.VectorLong()
	return dbuf.GetError()
}

func (m *VectorInt) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.VectorInt(m.Datas)
	return x.GetBuf()
}

func (m *VectorInt) Decode(dbuf *DecodeBuf, layer int32) error {
	// dbuf.Int()
	m.Datas = dbuf.VectorInt()
	return dbuf.GetError()
}

func (m *Vector_CdnFileHash) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_CdnFileHash) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*CdnFileHash, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &CdnFileHash{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_ContactStatus) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_ContactStatus) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*ContactStatus, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &ContactStatus{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_DialogFilter) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_DialogFilter) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*DialogFilter, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &DialogFilter{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_DialogFilterSuggested) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_DialogFilterSuggested) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*DialogFilterSuggested, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &DialogFilterSuggested{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_DialogPeer) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_DialogPeer) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*DialogPeer, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &DialogPeer{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_EmojiLanguage) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_EmojiLanguage) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*EmojiLanguage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &EmojiLanguage{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_FileHash) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_FileHash) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*FileHash, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &FileHash{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_LangPackString) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_LangPackString) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*LangPackString, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &LangPackString{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_MessageRange) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_MessageRange) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*MessageRange, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &MessageRange{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *VectorMessages_SearchCounter) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *VectorMessages_SearchCounter) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*Messages_SearchCounter, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &Messages_SearchCounter{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_ReceivedNotifyMessage) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_ReceivedNotifyMessage) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*ReceivedNotifyMessage, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &ReceivedNotifyMessage{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_SavedContact) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_SavedContact) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*SavedContact, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &SavedContact{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_SecureValue) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_SecureValue) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*SecureValue, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &SecureValue{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_StickerSetCovered) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, (*v).Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_StickerSetCovered) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*StickerSetCovered, l1)
	for i := int32(0); i < l1; i++ {
		m.Datas[i] = &StickerSetCovered{}
		if err := (*m.Datas[i]).Decode(dbuf, layer); err != nil {
			return err
		}
	}
	return dbuf.GetError()
}

func (m *Vector_LangPackLanguage) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, v.To_LangPackLanguage().Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_LangPackLanguage) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*LangPackLanguage, l1)
	for i := int32(0); i < l1; i++ {
		m1 := NewTLLangPackLanguage(nil)
		if err := m1.Decode(dbuf, layer); err != nil {
			return err
		}
		m.Datas[i] = m1.To_LangPackLanguage()
	}
	return dbuf.GetError()
}

func (m *Vector_User) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, v.To_User().Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_User) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*User, l1)
	for i := int32(0); i < l1; i++ {
		m1 := NewTLUser(nil)
		if err := m1.Decode(dbuf, layer); err != nil {
			return err
		}
		m.Datas[i] = m1.To_User()
	}
	return dbuf.GetError()
}

func (m *Vector_WallPaper) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(Cmd_vector))
	x.Int(int32(len(m.Datas)))
	for _, v := range m.Datas {
		x.buf = append(x.buf, v.To_WallPaper().Encode(layer)...)
	}
	return x.GetBuf()
}

func (m *Vector_WallPaper) Decode(dbuf *DecodeBuf, layer int32) error {
	//dbuf.Int()
	l1 := dbuf.Int()
	if dbuf.GetError() != nil {
		return dbuf.GetError()
	}
	m.Datas = make([]*WallPaper, l1)
	for i := int32(0); i < l1; i++ {
		m1 := NewTLWallPaper(nil)
		if err := m1.Decode(dbuf, layer); err != nil {
			return err
		}
		m.Datas[i] = m1.To_WallPaper()
	}
	return dbuf.GetError()
}
