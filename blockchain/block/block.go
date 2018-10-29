package block

import (
	"crypto/sha256"
	"time"

	pb "github.com/mjkong/selfstudy/blockchain/protos"
)

func NewBlock(blockNum uint64, prevHash []byte) *pb.Block {

	block := &pb.Block{}
	block.Header = &pb.BlockHeader{}
	block.Data = &pb.BlockData{}

	block.Header.Timestamp = time.Now().Unix()
	block.Header.Version = blockNum
	block.Header.PrevBlockHash = prevHash

	return block

}

func (b *pb.Block) Hash() []byte {

	hash := sha256.Sum256(b.Bytes())

	return hash
}

func (b *pb.Block) Bytes() []byte {

	[]byte blockBytes;



}