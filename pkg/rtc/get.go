// @File:     get.go
// @Created:  2020-03-29 00:53:50
// @Modified: 2020-04-03 12:05:42
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

package rtc

// GetInstallPathD returns the RunTimeCfg.InstallPath field if set, or the passed string otherwise.
func (rtc *RunTimeCfg) GetInstallPathD(str string) string {

	if rtc.InstallPath == "" {
		return str
	}
	return rtc.InstallPath
}

// GetOutputPathD returns the RunTimeCfg.OutputPath field if set, or the passed string otherwise.
func (rtc *RunTimeCfg) GetOutputPathD(str string) string {

	if rtc.OutputPath == "" {
		return str
	}
	return rtc.OutputPath
}

// GetDaemonArgsD returns the RunTimeCfg.DaemonArgs field if set, or the passed string otherwise.
func (rtc *RunTimeCfg) GetDaemonArgsD(str string) string {

	if rtc.DaemonArgs == "" {
		return str
	}
	return rtc.DaemonArgs
}
