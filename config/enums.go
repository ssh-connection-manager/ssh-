package config

const (
	NameFileConfig string = "config"
	TypeFileConfig string = "yml"
	Separator      string = "/"
	NameApp        string = "ssh+"
	Space          string = " "
	Dot            string = "."
)

const (
	NameFileConnects string = "connect-ssh.json"
	NameFileCryptKey string = "key.txt"
)

const (
	DirectionApp       = Separator + Dot + NameApp + Separator
	FullNameFileConfig = NameFileConfig + Dot + TypeFileConfig
)
