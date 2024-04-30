package leader

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"novachat_engine/pkg/log"
	"time"
)

type Result int32

const (
	ResultNone   Result = 0
	ResultLeader Result = 1
	ResultFollow Result = 2
)

type Config struct {
	EndPoints   []string
	ClusterName string
	NodeName    string
}

type Leader struct {
	client *clientv3.Client
	conf   Config
}

func NewLeader(conf Config) (*Leader, error) {
	cfg := clientv3.Config{
		Endpoints:   conf.EndPoints,
		DialTimeout: time.Second * 3,
	}
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Errorf("NewLeader error:%s", err.Error())
		return nil, err
	}

	return &Leader{
		conf:   conf,
		client: client,
	}, nil
}

func (l *Leader) Close() {
	if l.client != nil {
		l.client.Close()
	}
}

func (l *Leader) LeasesLock(result chan<- Result) error {
	lease := clientv3.NewLease(l.client)
	grantResponse, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		log.Errorf("LeasesLock Grant error:%s", err.Error())
		return err
	}

	leaseId := grantResponse.ID
	log.Debugf("LeasesLock leaseID:%v", leaseId)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	defer lease.Revoke(context.TODO(), leaseId)

	leaseKeepAliveResponse, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		log.Errorf("LeasesLock Grant KeepAlive:%s", err.Error())
		return err
	}

	key := fmt.Sprintf("/%s/lock", l.conf.ClusterName)
	kv := clientv3.NewKV(l.client)
	txn := kv.Txn(context.Background())
	txn.If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, l.conf.NodeName, clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(key))

	leaseKeepClose := make(chan struct{})
	var txnResponse *clientv3.TxnResponse
	go func() {
		for {
			select {
			case leaseKeepAlive, ok := <-leaseKeepAliveResponse:
				if !ok {
					result <- ResultNone
					close(leaseKeepClose)
					log.Warnf("LeaseId:%v Txn lock:%s close", leaseId, key)
					return
				}

				log.Debugf("leaseID:%+v ttl:%v keepAlive", leaseKeepAlive.ID, leaseKeepAlive.TTL)
			}
		}
	}()

	for {
		txnResponse, err = txn.Commit()
		if err != nil {
			log.Warnf("LeaseId:%v Txn lock:%s error:%s", leaseId, key, err.Error())
			return err
		}

		if txnResponse.Succeeded {
			result <- ResultLeader
			select {
			case <-leaseKeepClose:
				return nil
			}
		} else {
			result <- ResultFollow
			time.Sleep(3 * time.Second)
		}
	}
}
