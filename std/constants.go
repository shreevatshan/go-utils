package std

const (
	EmptyString             = ""
	Comma                   = ","
	Colon                   = ":"
	Semicolon               = ";"
	AtSign                  = "@"
	NewLineAsString         = "\n"
	NewLineAsCharacter      = '\n'
	EqualTo                 = "="
	Space                   = " "
	ForwardSlashAsString    = "/"
	ForwardSlashAsCharacter = '/'
	BackwardSlash           = '\\'
	Hyphen                  = "-"
	QuestionMark            = "?"
	Ampersand               = "&"
)

const (
	MinutesIn1Day      = 1440
	SecondsIn1Minute   = 60
	SecondsIn2Minutes  = 120
	SecondsIn5Minutes  = 300
	SecondsIn15Minutes = 900
	SecondsIn15Hour    = 3600
)

const (
	JSONFileExtension          = ".json"
	WindowsExecutableExtension = ".exe"
	WindowsMSIExtension        = ".msi"
	ZipFileExtension           = ".zip"
	ChecksumFileExtension      = ".sha256"
)

const (
	WindowsOS = "windows"
	LinuxOS   = "linux"
)

const (
	Arch386   = "386"
	ArchAMD64 = "amd64"
	ArchARM   = "arm"
	ArchARM64 = "arm64"
)

const (
	PreviousDirectory = ".."
)

var (
	JSONNull *string
)
