package mcprotocol

import (
	"encoding/binary"

	"github.com/layou233/ZBProxy/common/buf"
)

func ReadInt8(buffer *buf.Buffer) (int8, error) {
	b, err := buffer.ReadByte()
	if err != nil {
		return 0, err
	}
	return int8(b), nil
}

func ReadInt16(buffer *buf.Buffer) (int16, error) {
	bytes, err := buffer.Peek(2)
	if err != nil {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(bytes)), nil
}

func ReadUint16(buffer *buf.Buffer) (uint16, error) {
	bytes, err := buffer.Peek(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bytes), nil
}

func ReadInt(buffer *buf.Buffer) (int, error) {
	bytes, err := buffer.Peek(4)
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint32(bytes)), nil
}

func ReadInt32(buffer *buf.Buffer) (int32, error) {
	bytes, err := buffer.Peek(4)
	if err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(bytes)), nil
}

func ReadUint32(buffer *buf.Buffer) (uint32, error) {
	bytes, err := buffer.Peek(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(bytes), nil
}

func ReadInt64(buffer *buf.Buffer) (int64, error) {
	bytes, err := buffer.Peek(8)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(bytes)), nil
}

func ReadUint64(buffer *buf.Buffer) (uint64, error) {
	bytes, err := buffer.Peek(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bytes), nil
}

func ReadString(buffer *buf.Buffer) (string, error) {
	n, _, err := ReadVarIntFrom(buffer)
	if err != nil {
		return "", err
	}
	bytes, err := buffer.Peek(int(n))
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func WriteToPacket(buffer *buf.Buffer, item ...any) (err error) {
	for _, raw := range item {
		switch i := raw.(type) {
		case []byte:
			_, err = VarInt(int32(len(i))).WriteTo(buffer)
			if err != nil {
				break
			}
			_, err = buffer.Write(i)
			if err != nil {
				break
			}
		case string:
			_, err = VarInt(int32(len(i))).WriteTo(buffer)
			if err != nil {
				break
			}
			_, err = buffer.WriteString(i)
			if err != nil {
				break
			}
		case int8:
			return buffer.WriteByte(byte(i))
		case uint8: // aka byte
			return buffer.WriteByte(i)
		case int16:
			binary.BigEndian.PutUint16(buffer.Extend(2), uint16(i))
		case uint16:
			binary.BigEndian.PutUint16(buffer.Extend(2), i)
		case int:
			binary.BigEndian.PutUint32(buffer.Extend(4), uint32(i))
		case int32:
			binary.BigEndian.PutUint32(buffer.Extend(4), uint32(i))
		case uint32:
			binary.BigEndian.PutUint32(buffer.Extend(4), i)
		case int64:
			binary.BigEndian.PutUint64(buffer.Extend(8), uint64(i))
		case uint64:
			binary.BigEndian.PutUint64(buffer.Extend(8), i)
		case VarInt:
			_, err = i.WriteTo(buffer)
			if err != nil {
				break
			}
		case Message:
			_, err = i.WriteTo(buffer)
			if err != nil {
				break
			}
		case *Message:
			_, err = i.WriteTo(buffer)
			if err != nil {
				break
			}
		}
	}
	return
}
