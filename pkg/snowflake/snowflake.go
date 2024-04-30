package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"novachat_engine/pkg/log"
	"sync"
	"time"
)

var baseEpoch int64
var once sync.Once

func init() {
	baseEpoch = snowflake.Epoch
}

func UpdateEpoch() {
	once.Do(func() {
		now := time.Now().UnixMilli()
		if now >= baseEpoch {
			snowflake.Epoch = now - baseEpoch
		}
		log.Debugf("epoch:%d %d %d\n", baseEpoch, now, snowflake.Epoch)
	})
}

type Node struct {
	n *snowflake.Node
}

func NewNode(node int64) (*Node, error) {
	n, err := snowflake.NewNode(node)
	if err != nil {
		return nil, err
	}
	return &Node{n: n}, nil
}

func (n *Node) Generate() snowflake.ID {
	return n.n.Generate()
}
