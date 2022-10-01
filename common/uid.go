package common

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

// UID is method to
type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(objectType int, localID, shardID uint32) UID {
	return UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

// Shard : 1 Object : 1 Local : 1 => 0001 0001 0001
// 1 << 8 => 0001 0000 0000
// 1 << 4 => 0000 0001 0000
// 1 << 0 =>              1
// => 0001 0001 0001
func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) getLocalID() uint32 {
	return uid.localID
}

func (uid UID) getObjectType() int {
	return uid.objectType
}

func (uid UID) getShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, err
	}
	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}
	// x = 1110 1110 0101 => x >> 4 = 1110 1110 & 0000 1111 = 1110
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid >> 0 & 0x3FFFF),
	}
	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", uid.String())), nil
}

func (uid UID) UnMarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}
	uid.localID = decodeUID.localID
	uid.shardID = decodeUID.shardID
	uid.objectType = decodeUID.objectType
	return nil
}
