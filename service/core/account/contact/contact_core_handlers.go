package contact

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/log"
	"novachat_engine/service/data/contact"
	"novachat_engine/service/mgo"
)

func (c *Core) GetContacts(userId int64) ([]*data_contact.Contact, error) {
	col := c.getCollection()

	cursor, err := col.Find(context.Background(), mgo.DBE.MarshalCustomSpecMap(&data_contact.Contact{UserId: userId}, "UserId", "Deleted"))
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetContacts userId:%v error:%s", userId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		return nil, nil
	}

	var ret []*data_contact.Contact
	err = cursor.All(context.Background(), &ret)
	if err != nil {
		log.Errorf("GetContacts userId:%v decoder error:%s", userId, err.Error())
		return nil, nil
	}

	return ret, nil
}

func (c *Core) GetContactById(userId int64, peerId int64) (*data_contact.Contact, error) {
	col := c.getCollection()

	ret := &data_contact.Contact{UserId: userId, PeerId: peerId, Deleted: false}
	singleResult := col.FindOne(context.Background(), mgo.DBE.MarshalCustomSpecMap(ret, "UserId", "PeerId", "Deleted"))
	if singleResult.Err() != nil && singleResult.Err() != mongo.ErrNoDocuments {
		log.Errorf("GetContactById userId:%v error:%s", userId, singleResult.Err().Error())
		return nil, singleResult.Err()
	}
	if singleResult.Err() == mongo.ErrNoDocuments {
		return nil, nil
	}

	err := singleResult.Decode(ret)
	if err != nil {
		log.Errorf("GetContacts userId:%v decoder error:%s", userId, err.Error())
		return nil, nil
	}

	return ret, nil
}

func (c *Core) GetContactByIdList(userId int64, peerId []int64) ([]*data_contact.Contact, error) {
	col := c.getCollection()

	var ret []*data_contact.Contact
	filter := mgo.DBE.MarshalCustomSpecMap(&data_contact.Contact{UserId: userId, Deleted: false}, "UserId", "Deleted")
	filter["peer_id"] = bson.M{mgo.IN: peerId}

	cursor, err := col.Find(context.Background(), filter)
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetContactByIdList userId:%v error:%s", userId, err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err == mongo.ErrNilDocument {
		return nil, nil
	}

	err = cursor.All(context.Background(), &ret)
	if err != nil {
		log.Errorf("GetContactByIdList userId:%v decoder error:%s", userId, err.Error())
		return nil, nil
	}

	return ret, nil
}

func (c *Core) AddContact(userId int64, phone string, peerId int64, name string, name2 string, now int32) (data_contact.MutualType, error) {
	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	var (
		values    []*data_contact.Contact
		myValue   = &data_contact.Contact{}
		peerValue = &data_contact.Contact{}
	)

	contact := data_contact.MutualTypeMyContact
	var err error
	ot := options.Transaction()
	ot.SetReadConcern(readconcern.Majority())
	ot.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	err = mgo.GetMongoDB().UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction(ot)
		if err != nil {
			return err
		}

		col := c.getCollection()
		cursor, err1 := col.Find(sessionContext, bson.M{mgo.OR: []bson.M{
			mgo.DBE.MarshalCustomSpecMap(&data_contact.Contact{UserId: userId, PeerId: peerId}, "UserId", "PeerId"),
			mgo.DBE.MarshalCustomSpecMap(&data_contact.Contact{UserId: peerId, PeerId: userId}, "UserId", "PeerId"),
		}})
		if err1 != nil && err1 != mongo.ErrNilDocument {
			sessionContext.AbortTransaction(sessionContext)

			log.Errorf("AddContact userId:%v peerId:%d Find error:%s", userId, peerId, err1.Error())
			return err1
		}

		err = cursor.All(sessionContext, &values)
		cursor.Close(sessionContext)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)

			log.Errorf("AddContact userId:%v peerId:%d decode error:%s", userId, peerId, err.Error())
			return err1
		}

		for _, v := range values {
			if v.UserId == userId {
				myValue = v
			} else if v.PeerId == peerId {
				peerValue = v
			}
		}

		if peerValue.GetContact() >= data_contact.MutualTypeMutual && myValue.GetContact() >= data_contact.MutualTypeMutual ||
			peerValue.GetContact() == data_contact.MutualTypeDefault && myValue.GetContact() == data_contact.MutualTypeMyContact {

			sessionContext.AbortTransaction(sessionContext)
			return nil
		}

		if peerValue.GetContact() == data_contact.MutualTypeMyContact {
			contact = data_contact.MutualTypeMutual
		}
		myValue.UserId = userId
		myValue.PeerId = peerId
		myValue.FirstName = name
		myValue.LastName = name2
		myValue.Contact = int32(contact)
		myValue.Deleted = false
		myValue.Date = now
		myValue.Block = false
		myValue.Phone = phone
		myValue.Id = makeContactId(userId, peerId)
		peerValue.Contact = int32(contact)
		peerValue.Date = now
		peerValue.Deleted = false
		peerValue.Id = makeContactId(peerId, userId)

		_, err = col.UpdateOne(sessionContext, mgo.DBE.MarshalCustomSpecMap(myValue, "UserId", "PeerId"),
			bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(myValue, "Id", "Phone", "Contact", "FirstName", "LastName", "NickName", "Block", "Date", "FoldId", "ClientId", "Deleted")}, updateOptions)
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)

			log.Errorf("AddContact userId:%v peerId:%d, UpdateOne myValue error:%s", userId, peerId, err.Error())
			return err
		}

		if contact == data_contact.MutualTypeMutual {
			_, err = col.UpdateOne(sessionContext, mgo.DBE.MarshalCustomSpecMap(peerValue, "UserId", "PeerId"),
				bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(peerValue, "Id", "Contact", "FirstName", "LastName", "NickName", "Block", "Date", "FoldId", "ClientId", "Deleted")}, updateOptions)
			if err != nil {
				sessionContext.AbortTransaction(sessionContext)

				log.Errorf("AddContact userId:%v UpdateOne peerValue error:%s", userId, err.Error())
				return err
			}
		}

		sessionContext.CommitTransaction(sessionContext)
		return nil
	})
	if err != nil {
		log.Errorf("AddContact userId:%v error:%s", userId, err.Error())
		return 0, err
	}
	return contact, nil
}

func (c *Core) ContactsBlock(userId int64, peerId int64, now int32) error {
	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	col := c.getCollection()
	value := &data_contact.Contact{
		UserId: userId,
		PeerId: peerId,
		Block:  true,
		Date:   now,
		Id:     makeContactId(userId, peerId),
	}
	_, err := col.UpdateOne(context.Background(), mgo.DBE.MarshalCustomSpecMap(value, "UserId", "PeerId"),
		bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(value, "Id", "FoldId", "Block", "Date")}, updateOptions)
	if err != nil {
		log.Errorf("ContactsBlock userId:%v error:%s", userId, err.Error())
		return err
	}

	return nil
}

func (c *Core) ContactsUnblock(userId int64, peerId int64, now int32) (bool, error) {
	updateOptions := options.Update()
	updateOptions.SetUpsert(true)

	col := c.getCollection()
	value := &data_contact.Contact{
		Id:     makeContactId(userId, peerId),
		UserId: userId,
		PeerId: peerId,
		Block:  false,
		Date:   now,
	}
	_, err := col.UpdateOne(context.Background(), mgo.DBE.MarshalCustomSpecMap(value, "UserId", "PeerId"),
		bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(value, "Id", "FoldId", "Block", "Date")}, updateOptions)
	if err != nil {
		log.Errorf("ContactsUnblock userId:%v error:%s", userId, err.Error())
		return false, err
	}

	return true, nil
}

func (c *Core) ContactsRemove(userId int64, list []int64) (bool, error) {

	col := c.getCollection()
	_, err := col.UpdateMany(context.Background(), bson.M{"user_id": userId, "peer_id": bson.M{mgo.IN: list}},
		bson.M{mgo.SET: mgo.DBE.MarshalCustomSpecMap(&data_contact.Contact{Contact: 0, Deleted: true}, "Contact", "Deleted", "Date")})
	if err != nil {
		log.Errorf("ContactsRemove userId:%v list:%v error:%s", userId, list, err.Error())
		return false, err
	}

	return true, nil
}

func (c *Core) GetBlocked(userId int64) ([]*data_contact.Contact, error) {
	var ret []*data_contact.Contact
	col := c.getCollection()
	of := options.Find()
	of.SetProjection(bson.M{"user_id": 1, "date": 1, "peer_id": 1})
	cursor, err := col.Find(context.Background(),
		mgo.DBE.MarshalCustomSpecMap(&data_contact.Contact{UserId: userId, Block: true, Deleted: false}, "UserId", "Block"))
	if err != nil && err != mongo.ErrNilDocument {
		log.Errorf("GetBlocked userId:%verror:%s", userId, err.Error())
		return nil, err
	}
	cursor.All(context.Background(), &ret)
	cursor.Close(context.Background())
	return ret, nil
}
