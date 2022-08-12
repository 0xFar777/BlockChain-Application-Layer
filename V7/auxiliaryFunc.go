package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

//辅助函数 -- 将uint64转换为byte类型
//--  因为在求当前区块的哈希时，需要将区块的所有信息进行拼接在取哈希，转换成相同的类型便于拼接
func Uint64ToByte(num uint64) []byte {
	// func Write(w io.Writer, order ByteOrder, data interface{}) error
	//将data的binary编码格式写入w，data必须是定长值、定长值的切片、定长值的指针。
	//order指定写入数据的字节序，写入结构体时，名字中有'_'的字段会置为0。
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panicln(err)
	}
	return buffer.Bytes()
}

//辅助函数 -- 将时间戳转为标准时间
func TimeStampToStandardTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

//辅助函数 -- 将区块数据由byte字节流转化为string类型
func (input *TXInput) SetBytesDataToStringData() string {
	bytesdata := input.PubKey
	stringdata := string(bytesdata)
	return stringdata
}
