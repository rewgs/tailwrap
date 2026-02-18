package platform

import "log"

type Darwin struct {
	Kind Variant
}

type Variant int

const (
	AppStore = iota
	CommandLine
	Standalone
)

func (v Variant) String() (variant string) {
	switch v {
	case 0:
		variant = "App Store"
	case 1:
		variant = "Command-Line"
	case 2:
		variant = "Standalone"
	default:
		log.Fatalf("invalid value: %d", v)
	}
	return
}

func parseDarwinVariant(v string) (variant int) {
	switch v {
	case "App Store":
		variant = 0
	case "Command-Line":
		variant = 1
	case "Standalone":
		variant = 2
	default:
		log.Fatalf("invalid value: %s", v)
	}
	return
}

func GetDarwinVariantPath(path string) {
}

// path logic pseudo-code
//
// if "/Applications/Tailscale.app/" {
// 	if "/usr/local/bin/tailscale" {
// 		return "standalone"
// 	} else {
// 		return "app store"
// 	}
//
// 	if "~/go/bin/tailscale" {
// 		return "cli"
// 	}
// }
