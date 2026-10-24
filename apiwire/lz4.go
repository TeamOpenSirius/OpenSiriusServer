package apiwire

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pierrec/lz4/v4"
	"github.com/vmihailenco/msgpack/v5"
)

func DecodeMaybeLZ4BlockArray(body []byte) ([]byte, error) {
	if len(body) == 0 {
		return body, nil
	}
	r := bytes.NewReader(body)
	arrayLen, ok, err := readMsgpackArrayLen(r)
	if err != nil {
		return nil, err
	}
	if !ok || arrayLen < 2 {
		return body, nil
	}
	extType, extPayload, ok, err := readMsgpackExt(r)
	if err != nil {
		return nil, err
	}
	if !ok || extType != 98 {
		return body, nil
	}
	sizes, err := decodeLz4BlockSizes(extPayload)
	if err != nil {
		return nil, err
	}
	if len(sizes) == 0 {
		return nil, fmt.Errorf("empty lz4 block sizes")
	}
	if arrayLen != len(sizes)+1 {
		return nil, fmt.Errorf("lz4 block count mismatch: array len %d sizes %d", arrayLen, len(sizes))
	}
	var out bytes.Buffer
	for i, size := range sizes {
		block, err := readMsgpackBin(r)
		if err != nil {
			return nil, fmt.Errorf("read lz4 block %d: %w", i, err)
		}
		dst := make([]byte, size)
		n, err := lz4.UncompressBlock(block, dst)
		if err != nil {
			return nil, fmt.Errorf("decompress lz4 block %d: %w", i, err)
		}
		if n != size {
			return nil, fmt.Errorf("lz4 block %d size mismatch: got %d want %d", i, n, size)
		}
		out.Write(dst)
	}
	return out.Bytes(), nil
}

func EncodeLZ4BlockArray(raw []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		blockSize = 64 * 1024
	}
	if len(raw) == 0 {
		return nil, fmt.Errorf("cannot encode empty lz4 block array")
	}
	blocks := make([][]byte, 0, (len(raw)+blockSize-1)/blockSize)
	sizes := make([]int, 0, cap(blocks))
	for start := 0; start < len(raw); start += blockSize {
		end := start + blockSize
		if end > len(raw) {
			end = len(raw)
		}
		block := raw[start:end]
		dst := make([]byte, lz4.CompressBlockBound(len(block)))
		n, err := lz4.CompressBlock(block, dst, nil)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			n, err = lz4.CompressBlockHC(block, dst, 0, nil, nil)
			if err != nil {
				return nil, err
			}
		}
		if n == 0 {
			dst = encodeLZ4LiteralBlock(block)
			n = len(dst)
		}
		blocks = append(blocks, dst[:n])
		sizes = append(sizes, len(block))
	}
	var sizesPayload bytes.Buffer
	sizeEnc := msgpack.NewEncoder(&sizesPayload)
	for _, size := range sizes {
		if err := sizeEnc.EncodeInt(int64(size)); err != nil {
			return nil, err
		}
	}
	var out bytes.Buffer
	writeMsgpackArrayLen(&out, len(blocks)+1)
	writeMsgpackExt(&out, 98, sizesPayload.Bytes())
	for _, block := range blocks {
		writeMsgpackBin(&out, block)
	}
	return out.Bytes(), nil
}

func readMsgpackArrayLen(r *bytes.Reader) (int, bool, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, false, err
	}
	if b >= 0x90 && b <= 0x9f {
		return int(b & 0x0f), true, nil
	}
	switch b {
	case 0xdc:
		var buf [2]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return 0, false, err
		}
		return int(binary.BigEndian.Uint16(buf[:])), true, nil
	case 0xdd:
		var buf [4]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return 0, false, err
		}
		return int(binary.BigEndian.Uint32(buf[:])), true, nil
	default:
		_ = r.UnreadByte()
		return 0, false, nil
	}
}

func readMsgpackExt(r *bytes.Reader) (int8, []byte, bool, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, nil, false, err
	}
	var size uint32
	switch b {
	case 0xd4:
		size = 1
	case 0xd5:
		size = 2
	case 0xd6:
		size = 4
	case 0xd7:
		size = 8
	case 0xd8:
		size = 16
	case 0xc7:
		n, err := r.ReadByte()
		if err != nil {
			return 0, nil, false, err
		}
		size = uint32(n)
	case 0xc8:
		var buf [2]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return 0, nil, false, err
		}
		size = uint32(binary.BigEndian.Uint16(buf[:]))
	case 0xc9:
		var buf [4]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return 0, nil, false, err
		}
		size = binary.BigEndian.Uint32(buf[:])
	default:
		_ = r.UnreadByte()
		return 0, nil, false, nil
	}
	typeByte, err := r.ReadByte()
	if err != nil {
		return 0, nil, false, err
	}
	payload := make([]byte, size)
	if _, err := io.ReadFull(r, payload); err != nil {
		return 0, nil, false, err
	}
	return int8(typeByte), payload, true, nil
}

func decodeLz4BlockSizes(payload []byte) ([]int, error) {
	r := bytes.NewReader(payload)
	dec := msgpack.NewDecoder(r)
	sizes := make([]int, 0, 4)
	for r.Len() > 0 {
		v, err := dec.DecodeInt64()
		if err != nil {
			return nil, err
		}
		if v <= 0 {
			return nil, fmt.Errorf("invalid lz4 block size: %d", v)
		}
		sizes = append(sizes, int(v))
	}
	return sizes, nil
}

func readMsgpackBin(r *bytes.Reader) ([]byte, error) {
	b, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	var size uint32
	switch b {
	case 0xc4:
		n, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		size = uint32(n)
	case 0xc5:
		var buf [2]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return nil, err
		}
		size = uint32(binary.BigEndian.Uint16(buf[:]))
	case 0xc6:
		var buf [4]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return nil, err
		}
		size = binary.BigEndian.Uint32(buf[:])
	default:
		return nil, fmt.Errorf("expected msgpack bin, got 0x%02x", b)
	}
	data := make([]byte, size)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, err
	}
	return data, nil
}

func writeMsgpackArrayLen(w *bytes.Buffer, n int) {
	if n < 16 {
		w.WriteByte(0x90 | byte(n))
		return
	}
	if n <= 0xffff {
		w.WriteByte(0xdc)
		var b [2]byte
		binary.BigEndian.PutUint16(b[:], uint16(n))
		w.Write(b[:])
		return
	}
	w.WriteByte(0xdd)
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], uint32(n))
	w.Write(b[:])
}

func writeMsgpackExt(w *bytes.Buffer, typ int8, payload []byte) {
	size := len(payload)
	switch size {
	case 1:
		w.WriteByte(0xd4)
	case 2:
		w.WriteByte(0xd5)
	case 4:
		w.WriteByte(0xd6)
	case 8:
		w.WriteByte(0xd7)
	case 16:
		w.WriteByte(0xd8)
	default:
		if size <= 0xff {
			w.WriteByte(0xc7)
			w.WriteByte(byte(size))
		} else if size <= 0xffff {
			w.WriteByte(0xc8)
			var b [2]byte
			binary.BigEndian.PutUint16(b[:], uint16(size))
			w.Write(b[:])
		} else {
			w.WriteByte(0xc9)
			var b [4]byte
			binary.BigEndian.PutUint32(b[:], uint32(size))
			w.Write(b[:])
		}
	}
	w.WriteByte(byte(typ))
	w.Write(payload)
}

func writeMsgpackBin(w *bytes.Buffer, payload []byte) {
	size := len(payload)
	if size <= 0xff {
		w.WriteByte(0xc4)
		w.WriteByte(byte(size))
	} else if size <= 0xffff {
		w.WriteByte(0xc5)
		var b [2]byte
		binary.BigEndian.PutUint16(b[:], uint16(size))
		w.Write(b[:])
	} else {
		w.WriteByte(0xc6)
		var b [4]byte
		binary.BigEndian.PutUint32(b[:], uint32(size))
		w.Write(b[:])
	}
	w.Write(payload)
}

func encodeLZ4LiteralBlock(block []byte) []byte {
	var out bytes.Buffer
	literalLen := len(block)
	if literalLen < 15 {
		out.WriteByte(byte(literalLen << 4))
	} else {
		out.WriteByte(0xf0)
		remaining := literalLen - 15
		for remaining >= 255 {
			out.WriteByte(255)
			remaining -= 255
		}
		out.WriteByte(byte(remaining))
	}
	out.Write(block)
	return out.Bytes()
}
