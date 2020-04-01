package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"playgo/netexamples/myzinx/internal/ziface"
	"playgo/netexamples/myzinx/utils"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4字节) + ID uint32(4字节)
	return 8
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 写消息的各个部件 注意顺序不要乱 先写消息长度 再写消息id 最后写消息数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// 拆包方法 先将Head信息读出来 之后根据Head信息里的Data长度 再进行一次读
	dataBuff := bytes.NewReader(binaryData)
	// 只解压Head信息 得到dataLen和MsgID
	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 判断datalen 是否超过了我们允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		errMsg := fmt.Sprintf("too Large msg data recv! : %d ", msg.DataLen)
		return nil, errors.New(errMsg)
	}
	return msg, nil
}

var _ ziface.IDataPack = &DataPack{}
