package config

const (
	NameFileConfig = "config"
	TypeFileConfig = "yml"
	Separator      = "/"
	NameApp        = "ssh+"
	Space          = " "
	Dot            = "."
)

const (
	DirectionApp = Separator + Dot + Separator + NameApp + Separator
)

const (
	FullNameFileConfig = NameFileConfig + Dot + TypeFileConfig
	NameFileConnects   = "connect-ssh.json"
	NameFileCryptKey   = "key.txt"
)
