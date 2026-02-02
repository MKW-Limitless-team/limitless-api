package wwfc

import (
	"crypto/md5"
	"fmt"
	"net/url"
)

type User struct {
	ProfileID       uint64
	UserID          uint64
	GsbrCode        string
	Password        string
	NgDeviceID      []uint8
	Email           string
	UniqueNick      string
	FirstName       string
	LastName        string
	FriendInfo      string
	LastIPAddress   string
	LastInGameSn    string
	HasBan          bool
	BanIssued       string
	BanExpires      string
	BanReason       string
	BanReasonHidden string
	BanModerator    string
	BanTOS          bool
	OpenHost        bool
}

func GetMii(friendInfo string) string {
	return friendInfo[:102]
}

func ShowMii(mii string) string {
	return fmt.Sprintf("https://mii-unsecure.ariankordi.net/miis/image.png?data=%s&expression=normal&cameraYRotate=-30", url.QueryEscape(mii))
}

func PidToFC(pid uint64) uint64 {
	if pid == 0 {
		return 0
	}

	var buffer [8]byte

	buffer[0] = byte(pid)
	buffer[1] = byte(pid >> 8)
	buffer[2] = byte(pid >> 16)
	buffer[3] = byte(pid >> 24)

	buffer[4] = 'J'
	buffer[5] = 'C'
	buffer[6] = 'M'
	buffer[7] = 'R'

	sum := md5.Sum(buffer[:])

	return (uint64(sum[0]>>1) << 32) | pid
}

func FCToPid(fc uint64) uint64 {
	return fc & 0xFFFFFFFF
}
