package fakeUserAgent

import (
	_ "embed"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"math/rand"
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

func Filter[E any](s *[]E, f func(E) bool) []E {
	s2 := make([]E, 0, len(*s))
	for _, e := range *s {
		if f(e) {
			s2 = append(s2, e)
		}
	}
	return s2
}

func randFromLen(n int) int {
	source := rand.NewSource(int64(time.Now().UnixNano()))
	rand.New(source)
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

func getUserAgents(data []byte) (*[]UserAgents, error) {
	var userAgents []UserAgents
	if err := json.Unmarshal(data, &userAgents); err != nil {
		return nil, err
	}
	return &userAgents, nil
}
