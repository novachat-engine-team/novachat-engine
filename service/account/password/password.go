/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/11/9 15:02
 * @File : password.go
 */

package password

import (
	"novachat_engine/mtproto"
	"novachat_engine/pkg/crypto"
)

// PASSWORD_HASH_INVALID
// NEW_PASSWORD_BAD
// NEW_SALT_INVALID
// EMAIL_INVALID
// EMAIL_UNCONFIRMED

// SRP_PASSWORD_CHANGED
// SRP_ID_INVALID

// EC F8 73 76 65 BC 77 5A

const (
	kStatePasswordNone             = 0
	kStateNoRecoveryPassword       = 1
	kStateEmailUnconfirmedPassword = 2
	kStateConfirmedPassword        = 3
)

const (
	kServerSaltLen  = 8
	kNewAlgoSaltLen = 40
	kHashLen        = 256
)

var (
	gNewAlgoSalt1 = []byte{0xEC, 0xF8, 0x73, 0x76, 0x65, 0xBC, 0x77, 0x5A}
	gNewAlgoSalt2 = []byte{0xBE, 0xDE, 0x48, 0x88, 0x8C, 0x0F, 0x42, 0xAC, 0x34, 0xFF, 0xD1, 0xD4, 0x93, 0x5D, 0x8B, 0x21}

	gNewAlgoP = []byte{
		0xc7, 0x1c, 0xae, 0xb9, 0xc6, 0xb1, 0xc9, 0x04, 0x8e, 0x6c, 0x52, 0x2f,
		0x70, 0xf1, 0x3f, 0x73, 0x98, 0x0d, 0x40, 0x23, 0x8e, 0x3e, 0x21, 0xc1,
		0x49, 0x34, 0xd0, 0x37, 0x56, 0x3d, 0x93, 0x0f, 0x48, 0x19, 0x8a, 0x0a,
		0xa7, 0xc1, 0x40, 0x58, 0x22, 0x94, 0x93, 0xd2, 0x25, 0x30, 0xf4, 0xdb,
		0xfa, 0x33, 0x6f, 0x6e, 0x0a, 0xc9, 0x25, 0x13, 0x95, 0x43, 0xae, 0xd4,
		0x4c, 0xce, 0x7c, 0x37, 0x20, 0xfd, 0x51, 0xf6, 0x94, 0x58, 0x70, 0x5a,
		0xc6, 0x8c, 0xd4, 0xfe, 0x6b, 0x6b, 0x13, 0xab, 0xdc, 0x97, 0x46, 0x51,
		0x29, 0x69, 0x32, 0x84, 0x54, 0xf1, 0x8f, 0xaf, 0x8c, 0x59, 0x5f, 0x64,
		0x24, 0x77, 0xfe, 0x96, 0xbb, 0x2a, 0x94, 0x1d, 0x5b, 0xcd, 0x1d, 0x4a,
		0xc8, 0xcc, 0x49, 0x88, 0x07, 0x08, 0xfa, 0x9b, 0x37, 0x8e, 0x3c, 0x4f,
		0x3a, 0x90, 0x60, 0xbe, 0xe6, 0x7c, 0xf9, 0xa4, 0xa4, 0xa6, 0x95, 0x81,
		0x10, 0x51, 0x90, 0x7e, 0x16, 0x27, 0x53, 0xb5, 0x6b, 0x0f, 0x6b, 0x41,
		0x0d, 0xba, 0x74, 0xd8, 0xa8, 0x4b, 0x2a, 0x14, 0xb3, 0x14, 0x4e, 0x0e,
		0xf1, 0x28, 0x47, 0x54, 0xfd, 0x17, 0xed, 0x95, 0x0d, 0x59, 0x65, 0xb4,
		0xb9, 0xdd, 0x46, 0x58, 0x2d, 0xb1, 0x17, 0x8d, 0x16, 0x9c, 0x6b, 0xc4,
		0x65, 0xb0, 0xd6, 0xff, 0x9c, 0xa3, 0x92, 0x8f, 0xef, 0x5b, 0x9a, 0xe4,
		0xe4, 0x18, 0xfc, 0x15, 0xe8, 0x3e, 0xbe, 0xa0, 0xf8, 0x7f, 0xa9, 0xff,
		0x5e, 0xed, 0x70, 0x05, 0x0d, 0xed, 0x28, 0x49, 0xf4, 0x7b, 0xf9, 0x59,
		0xd9, 0x56, 0x85, 0x0c, 0xe9, 0x29, 0x85, 0x1f, 0x0d, 0x81, 0x15, 0xf6,
		0x35, 0xb1, 0x05, 0xee, 0x2e, 0x4e, 0x15, 0xd0, 0x4b, 0x24, 0x54, 0xbf,
		0x6f, 0x4f, 0xad, 0xf0, 0x34, 0xb1, 0x04, 0x03, 0x11, 0x9c, 0xd8, 0xe3,
		0xb9, 0x2f, 0xcc, 0x5b,
	}

	gNewAlgoG = int32(3)

	gNewSecureAlgoSalt = []byte{0x7D, 0x04, 0xB3, 0x4B, 0x94, 0x82, 0x8C, 0x3D}
)

var (
	// gPasswordKdfAlgoModPow *crypto2.PasswordKdfAlgoModPow
	gNewAlgo       *mtproto.PasswordKdfAlgo
	gNewSecureAlgo *mtproto.SecurePasswordKdfAlgo
	gSRPUtil       *crypto.SRPUtil
)

func init() {
	gNewAlgo = mtproto.NewTLPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512Iter100000SHA256ModPow(&mtproto.PasswordKdfAlgo{
		Salt1: gNewAlgoSalt1,
		Salt2: gNewAlgoSalt2,
		G:     gNewAlgoG,
		P:     gNewAlgoP,
	}).To_PasswordKdfAlgo()

	gNewSecureAlgo = mtproto.NewTLSecurePasswordKdfAlgoPBKDF2HMACSHA512Iter100000(&mtproto.SecurePasswordKdfAlgo{
		Salt: gNewSecureAlgoSalt,
	}).To_SecurePasswordKdfAlgo()

	gSRPUtil = crypto.MakeSRPUtil(&crypto.PasswordKdfAlgoModPow{
		Salt1: gNewAlgoSalt1,
		Salt2: gNewAlgoSalt2,
		G:     gNewAlgoG,
		P:     gNewAlgoP,
	})
}

func ToNoPassword() *mtproto.Account_Password {
	noPassword := mtproto.NewTLAccountPassword(&mtproto.Account_Password{
		NewAlgo:       gNewAlgo,
		NewSecureAlgo: gNewSecureAlgo,
		SecureRandom:  crypto.RandomBytes(256),
	})
	return noPassword.To_Account_Password()
}

/////////////////////
//
//type passwordData struct {
//	userId       int32
//	hint         string //  process hint
//	hasRecovery  bool
//	email        string
//	NewAlgoSalt1 []byte
//	V            []byte
//	srpId        int64
//	srpb         []byte
//	B            []byte
//	Code         string
//	CodeExpired  int32
//	state        int
//	*AccountCore
//}
//
//func (p *passwordData) String() string {
//	var s []string
//	s = append(s, fmt.Sprintf("{user_id: %d", p.userId))
//	s = append(s, fmt.Sprintf("hint: %s", p.hint))
//	s = append(s, fmt.Sprintf("has_recovery: %v", p.hasRecovery))
//	s = append(s, fmt.Sprintf("email: %s", p.email))
//	s = append(s, fmt.Sprintf("new_algo_salt1: %s", hex.EncodeToString(p.NewAlgoSalt1)))
//	s = append(s, fmt.Sprintf("v: %s", hex.EncodeToString(p.V)))
//	s = append(s, fmt.Sprintf("srp_id: %d", p.srpId))
//	s = append(s, fmt.Sprintf("srp_b: %s", hex.EncodeToString(p.srpb)))
//	s = append(s, fmt.Sprintf("B: %s", hex.EncodeToString(p.B)))
//	s = append(s, fmt.Sprintf("state: %d}", p.state))
//
//	return strings.Join(s, ", ")
//}
//
//func makeEMailPattern(email string) string {
//	// make pattern --> axxxa@domain.com
//	return email
//}
//
//// hasRecovery
//
//func (m *AccountCore) MakePasswordData(userId int32) (*passwordData, error) {
//	var (
//		err   error
//		pData *passwordData
//	)
//
//	do, err := m.UserPasswordsDAO.SelectByUserId(userId)
//	if err != nil {
//		log.Errorf("%v: not found user_password row, user_id: %d", err, userId)
//		return nil, err
//	}
//
//	if do == nil {
//		pData = &passwordData{
//			userId:      userId,
//			state:       0,
//			AccountCore: m,
//		}
//		return pData, nil
//	} else {
//		// check data.
//		pData := &passwordData{
//			userId:      userId,
//			hint:        do.Hint,
//			hasRecovery: util.Int8ToBool(do.HasRecovery),
//			email:       do.Email,
//			srpId:       do.SrpId,
//			Code:        do.Code,
//			CodeExpired: do.CodeExpired,
//			state:       int(do.State),
//			AccountCore: m,
//		}
//
//		if len(do.NewAlgoSalt1) > 0 {
//			pData.NewAlgoSalt1, err = hex.DecodeString(do.NewAlgoSalt1)
//			if err != nil || len(pData.NewAlgoSalt1) != kNewAlgoSaltLen {
//				err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NEW_SALT_INVALID)
//				log.Errorf("%v: db's newAlgoSalt1 error", err)
//				return nil, err
//			}
//		}
//		if len(do.V) > 0 {
//			pData.V, err = hex.DecodeString(do.V)
//			if err != nil || len(pData.V) != kHashLen {
//				err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//				log.Errorf("%v: db's V error", err)
//				return nil, err
//			}
//		}
//		if len(do.SrpB) > 0 {
//			pData.srpb, err = hex.DecodeString(do.SrpB)
//			if err != nil || len(pData.srpb) != kHashLen {
//				err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//				log.Errorf("%v: db's srpB error", err)
//				return nil, err
//			}
//		}
//		if len(do.B) > 0 {
//			pData.B, err = hex.DecodeString(do.B)
//			if err != nil || len(pData.B) != kHashLen {
//				err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//				log.Errorf("%v: db's B error", err)
//				return nil, err
//			}
//		}
//
//		return pData, nil
//	}
//}
//
//// SESSION_PASSWORD_NEEDED
//func (m *AccountCore) CheckSessionPasswordNeeded(ctx context.Context, userId int32) (bool, error) {
//	// 仅仅从数据库里取state字段
//	if do, err := m.UserPasswordsDAO.SelectByUserId(userId); err != nil {
//		return false, err
//	} else if do != nil {
//		return do.State == kStateNoRecoveryPassword || do.State == kStateConfirmedPassword, nil
//	}
//
//	return false, nil
//}
//
//func (p *passwordData) getHasRecovery() *mtproto.Bool {
//	return mtproto.ToBool(p.state == kStateConfirmedPassword)
//}
//
//func (p *passwordData) makeCurrentAlgo() *mtproto.PasswordKdfAlgo {
//	alog := mtproto.MakeTLPasswordKdfAlgoModPow(&mtproto.PasswordKdfAlgo{
//		Salt1: p.NewAlgoSalt1,
//		Salt2: gNewAlgo.Salt2,
//		G:     gNewAlgo.G,
//		P:     gNewAlgo.P,
//	})
//	return alog.To_PasswordKdfAlgo()
//}
//func (p *passwordData) GetPassword() *mtproto.Account_Password {
//	log.Debugf("p: %s", p.String())
//
//	switch p.state {
//	case kStatePasswordNone:
//		noPassword := mtproto.MakeTLAccountPassword(&mtproto.Account_Password{
//			NewAlgo:       gNewAlgo,
//			NewSecureAlgo: gNewSecureAlgo,
//			SecureRandom:  crypto.RandomBytes(256),
//		})
//		return noPassword.To_Account_Password()
//	case kStateNoRecoveryPassword:
//		password := mtproto.MakeTLAccountPassword(&mtproto.Account_Password{
//			HasPassword:   true,
//			CurrentAlgo:   p.makeCurrentAlgo(),
//			Srp_B:         p.B,
//			SrpId:         &types.Int64Value{Value: p.srpId},
//			NewAlgo:       gNewAlgo,
//			NewSecureAlgo: gNewSecureAlgo,
//			SecureRandom:  crypto.RandomBytes(256),
//		})
//		return password.To_Account_Password()
//	case kStateEmailUnconfirmedPassword:
//		noPassword := mtproto.MakeTLAccountPassword(&mtproto.Account_Password{
//			EmailUnconfirmedPattern_68873BA5: &types.StringValue{Value: makeEMailPattern(p.email)},
//			NewAlgo:                          gNewAlgo,
//			NewSecureAlgo:                    gNewSecureAlgo,
//			SecureRandom:                     crypto.RandomBytes(256),
//		})
//		return noPassword.To_Account_Password()
//	case kStateConfirmedPassword:
//		password := mtproto.MakeTLAccountPassword(&mtproto.Account_Password{
//			HasRecovery_CA39B447: true,
//			HasPassword:          true,
//			CurrentAlgo:          p.makeCurrentAlgo(),
//			Srp_B:                p.B,
//			SrpId:                &types.Int64Value{Value: p.srpId},
//			NewAlgo:              gNewAlgo,
//			NewSecureAlgo:        gNewSecureAlgo,
//			SecureRandom:         crypto.RandomBytes(256),
//		})
//		if p.hint != "" {
//			password.SetHint_68873BA5(&types.StringValue{Value: p.hint})
//		}
//		return password.To_Account_Password()
//	default:
//		// bug.
//		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
//		log.Error(err.Error())
//		panic(err)
//		return nil
//	}
//}
//
////check salt, currentPasswordHash, newPasswordHash
//// password:InputCheckPasswordSRP new_settings:account.PasswordInputSettings
//// inputCheckPasswordEmpty#9880f658 = InputCheckPasswordSRP;
//// inputCheckPasswordSRP#d27ff082 srp_id:long A:bytes M1:bytes = InputCheckPasswordSRP;
////
//// account.passwordInputSettings#c23727c9 flags:#
//// 	new_algo:flags.0?PasswordKdfAlgo
//// 	new_password_hash:flags.0?bytes
//// 	hint:flags.0?string
//// 	email:flags.1?string
//// 	new_secure_settings:flags.2?SecureSecretSettings = account.PasswordInputSettings;
////
//// currentPasswordHash, newSalt, newPasswordHash []byte, hint, email string)
//func (p *passwordData) UpdatePasswordSetting(password *mtproto.InputCheckPasswordSRP, newSettings *mtproto.Account_PasswordInputSettings) error {
//	var err error
//
//	switch p.state {
//	case kStatePasswordNone:
//		// set password.
//		//
//		if password.Constructor != mtproto.CRC32_inputCheckPasswordEmpty {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
//			log.Errorf("%v: current_password_hash need empty.", err)
//			return err
//		}
//
//		if len(newSettings.NewPasswordHash) == 0 {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//			glog.Error(err, ": current_password_hash need empty.")
//			return err
//		}
//
//		if newSettings.NewAlgo == nil {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//			glog.Error(err, ": current_password_hash need empty.")
//			return err
//		}
//
//		// check salt
//		if !bytes.Equal(gNewAlgoSalt1, newSettings.NewAlgo.Salt1[:8]) {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_NEW_SALT_INVALID)
//			log.Errorf("%v: current_password_hash need empty.", err)
//			return err
//		}
//
//		// 1. new_algo_salt1
//		p.NewAlgoSalt1 = newSettings.NewAlgo.Salt1
//
//		// 2. calc v
//		p.srpId = id_facade.GetUUID()
//		p.V = newSettings.NewPasswordHash
//		p.srpb, p.B = gSRPUtil.CalcSRPB(p.V)
//

//		if newSettings.Hint != nil {
//			p.hint = newSettings.Hint.Value
//		}
//		if newSettings.Email != nil {
//			p.email = newSettings.Email.Value
//		}
//
//		if p.email != "" {
//			p.state = kStateEmailUnconfirmedPassword
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_EMAIL_UNCONFIRMED)
//		} else {
//			p.state = kStateNoRecoveryPassword
//		}
//
//		// save
//		do := &dataobject.UserPasswordsDO{
//			UserId:       p.userId,
//			NewAlgoSalt1: hex.EncodeToString(p.NewAlgoSalt1),
//			V:            hex.EncodeToString(p.V),
//			SrpId:        p.srpId,
//			SrpB:         hex.EncodeToString(p.srpb),
//			B:            hex.EncodeToString(p.B),
//			Hint:         p.hint,
//			Email:        p.email,
//			State:        int8(p.state),
//		}
//		_, _, err = p.AccountCore.UserPasswordsDAO.InsertOrUpdate(do)
//
//		return err
//	case kStateEmailUnconfirmedPassword:
//		// check currentPasswordHash, currentPasswordHash: req.current_password_hash = new byte[0]
//		//
//		// email未验证状态时，客户端可以取消但不能更改密码
//		// email验证后由服务端修改状态 --> kStateConfirmedPassword
//		/*
//		  [12:40:00.488 01-0001042] (dc:5) Send: { core_message
//		  msg_id: 6669214788062407200 [LONG],
//		  seq_no: 62 [INT],
//		  bytes: 80 [INT],
//		  body: { msg_container
//		    messages: [ vector<0xffffffffffffffff>
//		      { core_message
//		        msg_id: 6669214788062378428 [LONG],
//		        seq_no: 61 [INT],
//		        bytes: 20 [INT],
//		        body: { account_updatePasswordSettings
//		          password: { inputCheckPasswordEmpty },
//		          new_settings: { account_passwordInputSettings
//		            flags: 2 [INT],
//		            new_algo: [ SKIPPED BY BIT 0 IN FIELD flags ],
//		            new_password_hash: [ SKIPPED BY BIT 0 IN FIELD flags ],
//		            hint: [ SKIPPED BY BIT 0 IN FIELD flags ],
//		            email: "" [STRING],
//		            new_secure_settings: [ SKIPPED BY BIT 2 IN FIELD flags ],
//		          },
//		        },
//		      },
//		*/
//
//		if password.Constructor != mtproto.CRC32_inputCheckPasswordEmpty {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
//			log.Errorf("%v: current_password_hash need empty.", err)
//			return err
//		}
//
//		if newSettings.Email == nil || newSettings.Email.Value != "" {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_EMAIL_INVALID)
//			log.Errorf("%v: email need empty.", err)
//			return err
//		}
//
//		// 检查是修改密码还是取消
//		if len(newSettings.NewPasswordHash) > 0 {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//			log.Errorf("%v: current_password_hash need empty.", err)
//			return err
//		}
//
//		//p.NewAlgoSalt1 = []byte{}
//		//p.V = []byte{}
//		//p.srpId = 0
//		//p.srpb = []byte{}
//		//p.B = []byte{}
//		//p.hint = ""
//		//p.email = ""
//		//p.state = kStatePasswordNone
//		//
//		//log.Debugf("cancel password")
//		//
//		//_, err = p.UserPasswordsDAO.Update(map[string]interface{}{
//		//	"new_algo_salt1": "",
//		//	"V":              "",
//		//	"srp_id":         0,
//		//	"srp_b":          "",
//		//	"B":              "",
//		//	"hint":           "",
//		//	"email":          "",
//		//	"state":          kStatePasswordNone,
//		//}, p.userId)
//		//
//		//return err
//		return p.clearPassword()
//	case kStateConfirmedPassword, kStateNoRecoveryPassword:
//		if password.Constructor != mtproto.CRC32_inputCheckPasswordSRP {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
//			log.Errorf("%v: current_password_hash need empty.", err)
//			return err
//		}
//
//		// email已经有了
//		if newSettings.Email != nil && newSettings.Email.Value != "" {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_EMAIL_INVALID)
//			log.Errorf("%v: email need empty.", err)
//			return err
//		}
//
//		// 检查是修改密码还是取消
//		if len(newSettings.NewPasswordHash) == 0 {
//			//  hint == ""
//			if newSettings.Hint != nil && newSettings.Hint.Value != "" {
//				//
//			}
//		} else {
//		}
//
//		// check password
//		if p.srpId != password.SrpId {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_SRP_ID_INVALID)
//			log.Errorf("%v: srp_id_invalid", err)
//			return err
//		}
//
//		M := gSRPUtil.CalcM(p.NewAlgoSalt1, p.V, password.A, p.srpb, p.B)
//		//log.Info(hex.EncodeToString(p.NewAlgoSalt1))
//		//log.Info(hex.EncodeToString(p.V))
//		//log.Info(hex.EncodeToString(password.A))
//		//log.Info(hex.EncodeToString(p.srpb))
//		//log.Info(hex.EncodeToString(p.B))
//		//log.Info(hex.EncodeToString(M))
//		//log.Info(hex.EncodeToString(password.M1))
//		if len(M) == 0 || !bytes.Equal(M, password.M1) {
//			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//			log.Errorf("%v: current_password_hash need empty.", err)
//			return err
//		}
//
//		// 检查是修改密码还是取消
//		if len(newSettings.NewPasswordHash) > 0 {
//			// 1. new_algo_salt1
//			p.NewAlgoSalt1 = newSettings.NewAlgo.Salt1
//
//			// 2. calc v
//			p.srpId = id_facade.GetUUID()
//			p.V = newSettings.NewPasswordHash
//			// gSRPUtil.GetVBytes(p.NewAlgoSalt1, newSettings.NewPasswordHash)
//			p.srpb, p.B = gSRPUtil.CalcSRPB(p.V)
//			if newSettings.Hint != nil {
//				p.hint = newSettings.Hint.Value
//			}
//			p.email = ""
//			// p.state = kStatePasswordNone
//
//			log.Debugf("update password")
//			_, err = p.UserPasswordsDAO.Update(map[string]interface{}{
//				"new_algo_salt1": hex.EncodeToString(p.NewAlgoSalt1),
//				"V":              hex.EncodeToString(p.V),
//				"srp_id":         p.srpId,
//				"srp_b":          hex.EncodeToString(p.srpb),
//				"B":              hex.EncodeToString(p.B),
//				"hint":           p.hint,
//				"email":          "",
//			}, p.userId)
//
//			return err
//		} else {
//			//p.NewAlgoSalt1 = []byte{}
//			//p.V = []byte{}
//			//p.srpId = 0
//			//p.srpb = []byte{}
//			//p.B = []byte{}
//			//p.hint = ""
//			//p.email = ""
//			//p.state = kStatePasswordNone
//			//
//			//log.Debugf("cancel password")
//			//
//			//_, err = p.UserPasswordsDAO.Update(map[string]interface{}{
//			//	"new_algo_salt1": "",
//			//	"V":              "",
//			//	"srp_id":         0,
//			//	"srp_b":          "",
//			//	"B":              "",
//			//	"hint":           "",
//			//	"email":          "",
//			//	"state":          kStatePasswordNone,
//			//}, p.userId)
//			//
//			//return err
//
//			return p.clearPassword()
//		}
//	default:
//		// bug.
//		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
//		log.Error(err.Error())
//		panic(err)
//		return err
//	}
//
//	// p.saveToDB()
//	return nil
//}
//
//func (p *passwordData) CheckPassword(password *mtproto.InputCheckPasswordSRP) bool {
//	if password.Constructor != mtproto.CRC32_inputCheckPasswordSRP {
//		return false
//	}
//
//	if p.srpId != password.SrpId {
//		return false
//	}
//
//	M := gSRPUtil.CalcM(p.NewAlgoSalt1, p.V, password.A, p.srpb, p.B)
//	if len(M) == 0 || !bytes.Equal(M, password.M1) {
//		return false
//	}
//	return true
//}
//
//func (p *passwordData) GetPasswordSetting(password *mtproto.InputCheckPasswordSRP) (*mtproto.Account_PasswordSettings, error) {
//	checked := p.CheckPassword(password)
//	if !checked {
//		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_PASSWORD_HASH_INVALID)
//		log.Errorf("%v: current_password_hash - %s", err, password)
//		return nil, err
//	}
//
//	setting := mtproto.MakeTLAccountPasswordSettings(&mtproto.Account_PasswordSettings{
//		Email_9A5C33E5: &types.StringValue{Value: p.email},
//	})
//	return setting.To_Account_PasswordSettings(), nil
//}
//
//func (p *passwordData) RequestPasswordRecovery() (*mtproto.Auth_PasswordRecovery, error) {
//	//  FLOOD_WAIT
//	passwordRecovery := mtproto.MakeTLAuthPasswordRecovery(&mtproto.Auth_PasswordRecovery{
//		EmailPattern: p.email,
//	})
//	return passwordRecovery.To_Auth_PasswordRecovery(), nil
//}
//
//func (p *passwordData) RecoverPassword(code string) error {
//	if p.Code != code {
//		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
//		log.Errorf("%s: userId - %d, code - %s", err, p.userId, code)
//		return err
//	}
//
//	//  check code_expired
//	return p.clearPassword()
//}
//
//func (m *AccountCore) CheckRecoverCode(userId int32, code string) error {
//	if do, err := m.UserPasswordsDAO.SelectCode(userId); err != nil {
//		return err
//	} else if do == nil {
//		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
//		log.Errorf("%v: not found user_password row, user_id - %d", err, userId)
//		return err
//	} else {
//		//  FLOOD_WAIT, check attempts??
//
//		if do.Code != code {
//			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
//			log.Errorf("%s: userId - %d, code - %s", err, userId, code)
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (p *passwordData) clearPassword() error {
//	// TODO(@benqi)
//	_, err := p.UserPasswordsDAO.Update(map[string]interface{}{
//		"new_algo_salt1": "",
//		"V":              "",
//		"srp_id":         0,
//		"srp_b":          "",
//		"B":              "",
//		"hint":           "",
//		"email":          "",
//		"state":          kStatePasswordNone,
//	}, p.userId)
//
//	if err != nil {
//		return err
//	}
//
//	p.NewAlgoSalt1 = []byte{}
//	p.V = []byte{}
//	p.srpId = 0
//	p.srpb = []byte{}
//	p.B = []byte{}
//	p.hint = ""
//	p.email = ""
//	p.state = kStatePasswordNone
//
//	return nil
//}
