package convo

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/Sheerley/dsip/plog"
)

// Config struct holds all the informations necessary to configure
// dsip apps (client, PrimaryNode and SecondaryNode)
type Config struct {
	Type                     string
	Address                  net.IP
	Port                     int
	ExternalPort             int
	JobBinaryName            string
	GarbageCollectionTimeout int
	MaxThreads               int
	DatabaseName             string
	DatabaseUsername         string
	DatabasePassword         string
}

type configJSON struct {
	Type       string `json:"type"`
	Addr       string `json:"address"`
	Port       int    `json:"port"`
	EPort      int    `json:"external-port"`
	Cmd        string `json:"job-binary-name"`
	GcTimeout  int    `json:"garbage-collection-timeout"`
	MaxThreads int    `json:"max-threads"`
	DbName     string `json:"database-name"`
	DbUname    string `json:"database-username"`
	DbPasswd   string `json:"database-password"`
}

// SavedConfig holds the default config
var SavedConfig Config

func (cc *Config) jsonToConfig(cj configJSON) {
	stringGCPercentage := os.Getenv("GC_PERCENTAGE")
	gcPercentage, err := strconv.Atoi(stringGCPercentage)
	if err != nil {
		cc.GarbageCollectionTimeout = cj.GcTimeout
	} else {
		if cj.EPort <= 100 {
			cc.GarbageCollectionTimeout = cj.GcTimeout
		} else {
			cc.GarbageCollectionTimeout = gcPercentage
		}
	}

	stringThreads := os.Getenv("THREADS")
	threads, err := strconv.Atoi(stringThreads)
	if err != nil {
		cc.MaxThreads = cj.MaxThreads
	} else {
		cc.MaxThreads = threads
	}

	cc.Type = cj.Type

	extenralAddress := os.Getenv("EXTERNAL_IP")

	if len(extenralAddress) > 7 {
		cc.Address = net.ParseIP(extenralAddress)
	} else {
		cc.Address = net.ParseIP(cj.Addr)
	}

	cc.Port = cj.Port
	cc.JobBinaryName = cj.Cmd

	stringPort := os.Getenv("EXTERNAL_PORT")

	port, err := strconv.Atoi(stringPort)
	if err != nil {
		if cj.EPort == 0 {
			cc.ExternalPort = cj.Port
		} else {
			cc.ExternalPort = cj.EPort
		}
	} else {
		cc.ExternalPort = port
	}

	logLevel := os.Getenv("LOG_LEVEL")
	plog.SSetLogLevel(logLevel)

	cc.DatabaseName = cj.DbName
	cc.DatabaseUsername = cj.DbUname
	cc.DatabasePassword = cj.DbPasswd

	return
}

// Tell is used to create splash info about node
func (cc Config) Tell() (s string) {
	s = fmt.Sprintf("\tWelcome to dsip!\n"+
		"\tThis server runs as %v node.\n"+
		"\tYou can acces it at %v:%v.\n",
		cc.Type, cc.Address, cc.ExternalPort)

	if cc.Type == "secondary" {
		s = fmt.Sprintf("%v"+
			"\tNumber of available concurrent processing threads is set to %v.\n"+
			"\tYou will be processing jobs using %v.\n",
			s, cc.MaxThreads, cc.JobBinaryName)
	}

	s = fmt.Sprintf("%v"+
		"\tGarbage collection target percentage is set to %v.\n",
		s, cc.GarbageCollectionTimeout)

	return
}
