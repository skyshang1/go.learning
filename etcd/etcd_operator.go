package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"time"
)

type EtcdOperator interface {
	// put

	// get

	// del

	// watch

	// txn
}

type etcdOperator struct {
	client *clientv3.Client
	logger *log.Logger
}

func NewEtcdOperator(endpoints []string, logger *log.Logger) *etcdOperator {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:        endpoints,
		DialTimeout:      3 * time.Second,
		AutoSyncInterval: 30 * time.Minute,
	})
	if err != nil {
		logger.Fatalf("construct etcd v3 client error: %s", err.Error())
	}

	operator := &etcdOperator{
		logger: logger,
		client: client,
	}

	return operator
}

func (o *etcdOperator) Put(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := o.client.Put(ctx, key, value)
	cancel()

	// handle error
	switch err {
	case context.Canceled:
		o.logger.Fatalf("ctx is canceld by another routine: %v", err)
	case context.DeadlineExceeded:
		o.logger.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
	case rpctypes.ErrEmptyKey:
		o.logger.Fatalf("client-side error: %v", err)
	default:
		o.logger.Fatalf("bad cluster endpoints, witch are not etcd servers: %v", err)
	}

	// print info
	o.logger.Printf("Header: [ClusterId: %v, MemberId: %v, RaftTerm: %v, Revision: %v]",
		resp.Header.ClusterId, resp.Header.MemberId, resp.Header.RaftTerm, resp.Header.Revision)
	o.logger.Printf("%v", resp.PrevKv)
}

// By default, Get will return the value for "key", if any.
// When passed WithRange(end), Get will return the keys in the range [key, end).
// When passed WithFromKey(), Get returns keys greater than or equal to key.
// When passed WithRev(rev) with rev > 0, Get retrieves keys at the given revision;
// if the required revision is compacted, the request will fail with ErrCompacted .
// When passed WithLimit(limit), the number of returned keys is bounded by limit.
// When passed WithSort(), the keys will be sorted.
func (o *etcdOperator) Get(key string) map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := o.client.Get(ctx, key)
	cancel()

	// handle error
	if err != nil {
		o.logger.Fatal(err)
	}

	value := make(map[string]string)
	// print info
	o.logger.Print(resp.Header.String())
	o.logger.Printf("[Count: %v, HasMore: %v]", resp.Count, resp.More)
	for _, kv := range resp.Kvs {
		o.logger.Print(kv.String())
		value[string(kv.Key)] = string(kv.Value)
	}

	return value
}

func (o *etcdOperator) Delete(key string, withPrefix bool) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	if withPrefix {
		_, err = o.client.Delete(ctx, key, clientv3.WithPrefix())
	} else {
		_, err = o.client.Delete(ctx, key)
	}

	cancel()

	// handle error
	if err != nil {
		o.logger.Fatal(err)
	}
}

func (o *etcdOperator) Txn() {
	_, err := o.client.Put(context.TODO(), "key", "xyz")
	if err != nil {
		o.logger.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = o.client.Txn(ctx).
		If(clientv3.Compare(clientv3.Value("key"), ">", "abc")).
		Then(clientv3.OpPut("key", "XYZ")).
		Else(clientv3.OpPut("key", "ABC")).
		Commit()
	cancel()
	if err != nil {
		o.logger.Fatal(err)
	}
}
