//go:build windows

package logger

type SyslogConfig struct {
	// Facility refers to the syslog facility of a given log
	Facility SyslogFacility `yaml:"facility" json:"facility"`
}
