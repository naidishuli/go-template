package postgresdb

import (
	"net/url"
	"regexp"
	"strconv"
)

var dbUrlRegex = regexp.MustCompile(`//(.+)?:(.+)?@(.+)?:(.\d+)?(\/(.*))?`)

type Config struct {
	Url                   string
	Host                  string
	Port                  int
	Username              string
	Password              string
	Database              string
	SSLMode               string
	MaxIdleConnections    int
	MaxOpenConnections    int
	SaveSQLAfterExecution bool
	CallbackSqlKey        string
}

func (c *Config) parse() {
	if c.MaxOpenConnections == 0 {
		c.MaxOpenConnections = 50
	}

	if c.MaxIdleConnections == 0 {
		c.MaxIdleConnections = 50
	}

	// get and remove extra query params passed to the connection string
	if c.Url != "" {
		u, _ := url.Parse(c.Url)

		for key, values := range u.Query() {
			for _, value := range values {
				if key == "sslmode" && value != "" {
					c.SSLMode = value
				}
			}
		}

		u.RawQuery = ""

		parts := dbUrlRegex.FindStringSubmatch(u.String())
		c.Username = parts[1]
		c.Password = parts[2]
		c.Host = parts[3]
		c.Port, _ = strconv.Atoi(parts[4])
		c.Database = parts[6]
	}
}
