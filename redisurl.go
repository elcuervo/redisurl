package redisurl

import (
	"net/url"
	"strconv"
	"strings"
)

type Url struct {
	User     string
	Password string
	Host     string
	Port     int
	Database int
}

func Parse(redisurl string) *Url {
	var username, password, host string
	var db, port int

	u, err := url.Parse(redisurl)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(u.Host, ":")
	host = parts[0]

	if host == "" {
		host = "127.0.0.1"
	}

	if len(u.Path) > 1 {
                path := u.Path[1:]
		db, err = strconv.Atoi(path)
                if err != nil {
                        panic("Invalid db" + path)
                }
	} else {
		db = 0
	}

	port, _ = strconv.Atoi(parts[1])

	if u.User != nil {
		username = u.User.Username()
		password, _ = u.User.Password()
	}

	return &Url{
		Host:     host,
		Port:     port,
		Database: db,
		User:     username,
		Password: password,
	}
}
