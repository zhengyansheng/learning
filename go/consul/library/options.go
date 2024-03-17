package library

import "time"

type ServiceOption struct {
	Name    string
	Address string
	Port    int
	Tags    []string
}

type CheckOption struct {
	TCP                                    string
	Interval                               time.Duration `json:"-"`
	Timeout                                time.Duration `json:"-"`
	DeregisterCriticalServiceAfterDuration time.Duration `json:"-"`
}
