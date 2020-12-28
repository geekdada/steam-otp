package otp

import (
	hmac2 "crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"github.com/funny/binary"
	"math"
	"regexp"
	"time"
)

/**
Returns the current local Unix time
*/
func getTime(timeOffset int64) int64 {
	now := time.Now()
	return now.Unix() + timeOffset
}

func bufferizeSecret(secret string) (*binary.Buffer, error) {
	matched, err := regexp.MatchString(`[0-9a-f]{40}`, secret)

	if err != nil {
		return nil, err
	}

	if matched {
		hexValue, err := hex.DecodeString(secret)

		if err != nil {
			return nil, err
		} else {
			return &binary.Buffer{Data: hexValue}, nil
		}
	}

	base64Value, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {
		return nil, err
	}

	return &binary.Buffer{Data: base64Value}, nil
}

/**
Generate a Steam-style TOTP authentication code
 */
func GenerateAuthCode(secret string, timeOffset int64) (string, error) {
	localTime := getTime(timeOffset)
	secretBuffer, err := bufferizeSecret(secret)

	if err != nil {
		return "", err
	}

	buf := binary.Buffer{Data: make([]byte, 8)}

	buf.WriteUint32BE(0)
	buf.WriteUint32BE(uint32(math.Floor(float64(localTime) / 30)))

	hmac := hmac2.New(sha1.New, secretBuffer.Data)
	hmac.Write(buf.Data)
	secretHash := hmac.Sum(nil)

	startPos := secretHash[19] & 0x0F
	slicedSecretHashBuf := binary.Buffer{Data: secretHash[startPos : startPos + 4]}

	fullCode := slicedSecretHashBuf.ReadUint32BE() & 0x7FFFFFFF
	chars := "23456789BCDFGHJKMNPQRTVWXY"
	code := ""

	for i := 0; i < 5; i++ {
		code += string([]rune(chars)[int(fullCode) % len(chars)])
		fullCode /= uint32(len(chars))
	}

	return code, nil
}


