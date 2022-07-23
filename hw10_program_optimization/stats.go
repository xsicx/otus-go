package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/mailru/easyjson"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

var (
	ErrInvalidFileData = errors.New("invalid file data")
	ErrInvalidDomain   = errors.New("invalid domain")
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	domainStat := make(DomainStat)

	regex, err := regexp.Compile(fmt.Sprintf("\\.%s$", domain))
	if err != nil {
		return domainStat, ErrInvalidDomain
	}

	var user User
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := easyjson.Unmarshal(line, &user); err != nil {
			return domainStat, ErrInvalidFileData
		}

		if !regex.MatchString(user.Email) {
			continue
		}

		domainStat[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
	}

	return domainStat, nil
}
