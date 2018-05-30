/**
*  @file
*  @copyright defined in dashboard-api/LICENSE
 */

package common

import (
	"github.com/satori/go.uuid"
)

// GetUUID get uuid with GetUUIDV2 first, if error use GetUUIDV1, if error again use GetUUIDV3
func GetUUID() (string, error) {
	uuid, err := GetUUIDV2()
	if err != nil {
		uuid, err = GetUUIDV1()
		if err != nil {
			return GetUUIDV3()
		}
	}
	return uuid, err
}

// GetUUIDV1 Version 1, based on timestamp and MAC address (RFC 4122)
func GetUUIDV1() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GetUUIDV2 Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)
func GetUUIDV2() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GetUUIDV3 Version 3, based on MD5 hashing (RFC 4122)
func GetUUIDV3() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GetUUIDV4 Version 4, based on random numbers (RFC 4122)
func GetUUIDV4() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// GetUUIDV5 Version 5, based on SHA-1 hashing (RFC 4122)
func GetUUIDV5() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
