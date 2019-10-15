package block

import (
	"encoding/asn1"
	"log"
	"time"

	pb "github.com/mjkong/selfstudy/blockchain/protos"
)

type asn1Header struct {
	Version      uint64
	PreviousHash []byte
	Hash         []byte
}

func NewBlock(blockNum uint64, prevHash []byte) *pb.Block {

	block := &pb.Block{}
	block.Header = &pb.BlockHeader{}
	block.Data = &pb.BlockData{}

	block.Header.Timestamp = time.Now().Unix()
	block.Header.Version = blockNum
	block.Header.PrevBlockHash = prevHash

	return block

}

func Hash(b *pb.Block) []byte {

	header := asn1Header{b.Header.Version, b.Header.PrevBlockHash, b.Header.Hash}

	blockBytes, err := asn1.Marshal(header)
	if err != nil {
		log.Fatal(err)
	}

	return blockBytes
}

// 	hash := sha256.Sum256(b.Bytes())

// 	return hash
// }

// func (b *pb.Block) Bytes() []byte {

// 	b.Mashal(
// }
