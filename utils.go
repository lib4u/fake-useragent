package fakeUserAgent

import (
	_ "embed"
	"encoding/json"
	"io"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

//go:embed userAgents.json
var userAgentsFile []byte

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Filter[E any](s []E, f func(E) bool) []E {
	s2 := make([]E, 0, len(s))
	for _, e := range s {
		if f(e) {
			s2 = append(s2, e)
		}
	}
	return s2
}

func randFromLen(n int) int {
	rand.Seed(uint64(time.Now().UnixNano()))
	randomInt := rand.Intn(n)
	return randomInt
}

func ExtractMajorVersion(version string) int {
	parts := strings.Split(version, ".")
	if len(parts) > 0 {
		ver, _ := strconv.Atoi(parts[0])
		return ver
	}
	return 0
}

func getUserAgents(r io.Reader) ([]UserAgents, error) {
	var userAgents []UserAgents
	if err := json.NewDecoder(r).Decode(&userAgents); err != nil {
		return nil, err
	}
	return userAgents, nil
}