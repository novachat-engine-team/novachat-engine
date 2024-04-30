package contact

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"novachat_engine/pkg/config"
	"novachat_engine/pkg/log"
	"novachat_engine/service/mgo"
	"strings"
	"sync"
)

const (
	DbContacts    string = "db_contacts"
	TableContacts string = "tb_contacts"
)

type Core struct {
	pool sync.Pool
}

func makeContactId(userId int64, peerId int64) string {
	return fmt.Sprintf("%d%d", userId, peerId)
}

func NewContactCore(conf *config.MongodbConfig) *Core {
	mgo.InstallMongodb(conf)
	c := &Core{}

	mgo.DBE.Enable64ToString = false
	return c
}

func (c *Core) Init() {

}

func (c *Core) Migrator() {
	allOpts := options.ListDatabases().
		SetNameOnly(true).
		SetAuthorizedDatabases(true)
	nameList, err := mgo.GetMongoDB().ListDatabaseNames(context.Background(), bson.M{}, allOpts)
	if err != nil {
		log.Warnf("NewContactCore error:%s", err.Error())
		panic(err)
	}

	found := false
	for _, name := range nameList {
		if strings.Compare(name, DbContacts) == 0 {
			found = true
			break
		}
	}
	if found { // collection
		found = false
		op := options.ListCollections()
		op.SetNameOnly(true)
		nameList, err = mgo.GetMongoDB().Database(DbContacts).ListCollectionNames(context.Background(), bson.M{}, op)
		if err != nil {
			log.Warnf("NewContactCore table error:%s", err.Error())
			panic(err)
		}
		for _, name := range nameList {
			if strings.Compare(name, TableContacts) == 0 {
				found = true
				break
			}
		}
	}

	if !found {
		err = mgo.GetMongoDB().Database(DbContacts).CreateCollection(context.Background(), TableContacts)
		if err != nil {
			log.Fatalf("%s Create tableName:%s error:%s", DbContacts, TableContacts, err.Error())
			panic(err)
		}
	}

	err = mgo.GetMongoDB().CreateIndexes(DbContacts, TableContacts,
		[]mgo.IndexKey{
			{Key: "user_id", Order: 1, Unique: false},
			{Key: "peer_id", Order: 1, Unique: false},
		}, true)
	if err != nil {
		log.Warnf("Create Index DBName:%s TableName:%s", DbContacts, TableContacts, err.Error())
	}
}

func (c *Core) getCollection() *mongo.Collection {
	op := options.Collection()
	op.SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	op.SetReadConcern(readconcern.Majority())
	//op.SetRegistry(mgo.Registry())

	col := mgo.GetDatabase(DbContacts).
		Collection(TableContacts, op)
	return col
}
