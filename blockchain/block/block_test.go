package block

import (
	"fmt"
	"testing"
)

func Test_NewBlock(t *testing.T) {

	newblock := NewBlock(1, []byte{244, 244, 244, 244, 244, 244, 244, 244})

	fmt.Println("Here is start of test!!!")
	fmt.Println(newblock)

}

func Test_Hash(t *testing.T) {

	fmt.Println("Here is start of Hash()")

	newblock := NewBlock(1, []byte{244, 244, 244, 244, 244, 244, 244, 244})

	fmt.Println("newly create block is ", newblock)

	Hash(newblock)

	// fmt.Println(blockbytes)

}
