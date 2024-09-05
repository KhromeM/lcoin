package za

import "encoding/binary"



type Transaction struct {
	From string
	To string
	Amount float64
	Nonce []byte
}

func (t Transaction) toBytes() []byte {
	byteSize := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteSize,uint64(t.Amount))
	return bytes.join([][]byte{
		[]byte(t.From),
		[]byte(t.To), 
		byteSize, 
		t.Nonce,
	}, []byte{})
}