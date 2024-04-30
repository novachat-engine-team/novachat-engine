package sms_service

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"go.uber.org/atomic"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/cache"
	"novachat_engine/pkg/log"
	"novachat_engine/service/common/errors"
	"novachat_engine/service/core/account/phone/phone_code"
	"novachat_engine/service/sms"
	"novachat_engine/service/sms/config"
	"strings"
)

type Service struct {
	phoneCode *phone_code.PhoneCode
	testMode  *atomic.Bool
	service   sms.Service
}

func NewSmsService(s sms.Service) *Service {
	return &Service{
		phoneCode: phone_code.NewPhoneCode(cache.GetRedisClient()),
		service:   s,
		testMode:  atomic.NewBool(true),
	}
}

func (s *Service) PlatForm() config.SmsPlatform {
	if s.service == nil {
		return config.SmsPlatformNone
	}
	return s.service.PlatForm()
}

func (s *Service) SetTestMode(testMode bool) {
	if s.testMode.Load() != testMode {
		s.testMode.Store(testMode)
	}
}

func (s *Service) SendSmsVerifyCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error {
	log.Debugf("SendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s", authKeyId, phoneNumber, phoneCodeHash)

	v, err := s.phoneCode.GetPhoneCode(phoneNumber, authKeyId)
	if err != nil {
		log.Errorf("SendSmsVerifyCode GetPhoneCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
		return err
	}
	if len(v) == 0 {
		buf, _ := jsoniter.MarshalToString(&sms.Code{
			PhoneCode:     phoneCode,
			PhoneCodeHash: phoneCodeHash,
		})

		if err = s.phoneCode.SetPhoneCode(phoneNumber, authKeyId, buf, sms.TTL); err != nil {
			log.Errorf("SendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
			return err
		}
		if s.testMode.Load() == true {
			err = s.sendSubMail(phoneNumber, phoneCode)
		}
	} else {
		c := &sms.Code{
			PhoneCode:     phoneCode,
			PhoneCodeHash: phoneCodeHash,
		}
		json.Unmarshal([]byte(v), c)
		log.Infof("------------ SendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s", authKeyId, c.PhoneCode, c.PhoneCodeHash)
	}

	if err != nil {
		log.Errorf("SendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
	} else {
		log.Debugf("SendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s ok!!!", authKeyId, phoneNumber, phoneCodeHash)
	}
	return err
}

func (s *Service) ReSendSmsVerifyCode(authKeyId int64, phoneNumber, phoneCodeHash string, newPhoneCode string) (int32, error) {
	var (
		v   string
		err error
		ttl int32
	)
	if v, err = s.phoneCode.GetPhoneCode(phoneNumber, authKeyId); len(v) != 0 {
		c := &sms.Code{}
		if err = json.Unmarshal([]byte(v), c); err != nil {
			return -1, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EXPIRED)
		}

		ttl, err = s.phoneCode.TTL(phoneNumber, authKeyId)
		if err != nil {
			return -1, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EXPIRED)
		}
		if ttl <= 0 {
			if len(c.PhoneCode) == 0 {
				err = errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EXPIRED)
				log.Errorf("ReSendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s c.PhoneCode error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
				return -1, err
			}

			err = s.SendSmsVerifyCode(authKeyId, phoneNumber, c.PhoneCode, phoneCodeHash)
			if err != nil {
				log.Errorf("ReSendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
				return 0, err
			}
			return sms.TTL, nil
		} else {
			c.PhoneCodeHash = phoneCodeHash
			buf, _ := jsoniter.MarshalToString(c)
			if err = s.phoneCode.SetPhoneCode(phoneNumber, authKeyId, buf, sms.TTL); err != nil {
				log.Errorf("ReSendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
				return 0, err
			}

			//if s.testMode.Load() == true {
			//err = s.sendSubMail(phoneNumber, c.PhoneCode)
			//}
		}
		return ttl, err
	} else {
		err = s.SendSmsVerifyCode(authKeyId, phoneNumber, newPhoneCode, phoneCodeHash)
		if err != nil {
			log.Errorf("ReSendSmsVerifyCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
			return 0, err
		}
		return sms.TTL, nil
	}

	return -1, errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EXPIRED)
}

// 验证登录验证码
func (s *Service) VerifySmsCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error {
	log.Debugf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s", authKeyId, phoneNumber, phoneCodeHash)

	v, err := s.phoneCode.GetPhoneCode(phoneNumber, authKeyId)
	if err != nil {
		log.Errorf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, err.Error())
		return errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EXPIRED)
	}

	c := &sms.Code{}
	err = json.Unmarshal([]byte(v), c)
	if err != nil {
		log.Errorf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s v:%s error:%s", authKeyId, phoneNumber, phoneCodeHash, v, err.Error())
		return errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EXPIRED)
	}

	//if len(c.PhoneCode) == 0 {
	//	log.Errorf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s c:%v", authKeyId, phoneNumber, phoneCodeHash, c)
	//	return errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_EMPTY)
	//}

	//if len(c.PhoneCodeHash) == 0 {
	//	log.Errorf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s c:%v", authKeyId, phoneNumber, phoneCodeHash, c)
	//	return errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_HASH_EMPTY)
	//}

	if len(phoneCode) != 0 && strings.Compare(phoneCode, c.PhoneCode) != 0 {
		log.Errorf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s c:%v", authKeyId, phoneNumber, phoneCode, c)
		return errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_INVALID)
	}

	if len(c.PhoneCodeHash) > 0 && strings.Compare(phoneCodeHash, c.PhoneCodeHash) != 0 {
		log.Errorf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s c:%v", authKeyId, phoneNumber, phoneCodeHash, c)
		return errors.NewRpcErrorWithRpcErrorCode(mtproto.RpcErrorCode_BAD_REQUEST_PHONE_CODE_INVALID)
	}

	log.Debugf("VerifySmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s ok!!!", authKeyId, phoneNumber, phoneCode)
	return nil
}

func (s *Service) ClearSmsCode(authKeyId int64, phoneNumber, phoneCode, phoneCodeHash string) error {
	_ = phoneCode

	log.Debugf("ClearSmsCode authKeyId:%d phoneNumber:%s phoneCodeHash:%s", authKeyId, phoneNumber, phoneCodeHash)
	return s.phoneCode.Clear(phoneNumber, authKeyId)
}

func (s *Service) sendSubMail(phoneNumber, Code string) error {
	err := s.service.Send(phoneNumber, Code)
	if err != nil {
		log.Errorf("sendSubMail error:%s", err.Error())
	}
	return err
}
