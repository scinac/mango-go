//go:build windows

package logger

func (sl MangoLogger) handleSyslogOutput(log *StructuredLog, jsonOut []byte) error {
	return nil
}
