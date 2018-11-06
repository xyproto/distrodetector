package distrodetector

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

const defaultName = "Unknown"

// Used when checking for Linux distros and BSDs (and NAME= is not defined in /etc)
var distroNames = []string{"Arch Linux", "Debian", "Ubuntu", "Void Linux", "FreeBSD", "NetBSD", "OpenBSD", "Manjaro", "Mint", "Elementary", "MX Linuyx", "Fedora", "openSUSE", "Solus", "Zorin", "CentOS", "KDE neon", "Lite", "Kali", "Antergos", "antiX", "Lubuntu", "PCLinuxOS", "Endless", "Peppermint", "SmartOS", "TrueOS", "Arco", "SparkyLinux", "deepin", "Puppy", "Slackware", "Bodhi", "Tails", "Xubuntu", "Archman", "Bluestar", "Mageia", "Deuvan", "Parrot", "Pop!", "ArchLabs", "Q4OS", "Kubuntu", "Nitrux", "Red Hat", "4MLinux", "Gentoo", "Pinguy", "LXLE", "KaOS", "Ultimate", "Alpine", "Feren", "KNOPPIX", "Robolinux", "Voyager", "Netrunner", "GhostBSD", "Budgie", "ClearOS", "Gecko", "SwagArch", "Emmabuntüs", "Scientific", "Omarine", "Neptune", "NixOS", "Slax", "Clonezilla", "DragonFly", "ExTiX", "OpenBSD", "Redcore", "Ubuntu Studio", "BunsenLabs", "BlackArch", "NuTyX", "ArchBang", "BackBox", "Sabayon", "AUSTRUMI", "Container", "ROSA", "SteamOS", "Tiny Core", "Kodachi", "Qubes", "siduction", "Parabola", "Trisquel", "Vector", "SolydXK", "Elive", "AV Linux", "Artix", "Raspbian", "Porteus"}

// TODO: Find a better way
//ihttps://en.wikipedia.org/wiki/List_of_Apple_operating_systems
//var codeNames = map[string]string{"10.0": "Cheetah", "10.1": "Puma", "10.2": "Jaguar", ...

// Distro represents the platform, contents of /etc/*release* and name of the
// detected Linux distribution or BSD.
type Distro struct {
	platform   string
	etcRelease string
	name       string
	codename   string
	version    string
}

// etcRelease returns the contents of /etc/*release*, or an empty string
func readEtcRelease() string {
	filenames, err := filepath.Glob("/etc/*release*")
	if err != nil {
		return ""
	}
	var bs strings.Builder
	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			continue
		}
		bs.Write(data)
	}
	return bs.String()
}

// New detects the platform and distro/BSD, then returns a pointer to
// a Distro struct.
func New() *Distro {
	var d Distro
	d.platform = capitalize(runtime.GOOS)
	d.etcRelease = readEtcRelease()
	// Distro name, if not detected
	d.name = defaultName
	d.codename = ""
	d.version = ""
	// Check for Linux distros and BSD distros
	for _, distroName := range distroNames {
		if d.Grep(distroName) {
			d.name = distroName
			break
		}
	}
	// Examine all lines of text in /etc/*release*
	for _, line := range strings.Split(d.etcRelease, "\n") {
		// Check if NAME= is defined in /etc/*release*
		if strings.HasPrefix(line, "NAME=") {
			fields := strings.SplitN(strings.TrimSpace(line), "=", 2)
			name := fields[1]
			if name != "" {
				if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
					d.name = name[1 : len(name)-1]
					continue
				}
				d.name = name
			}
			// Check if DISTRIB_CODENAME= is defined in /etc/*release*
		} else if strings.HasPrefix(line, "DISTRIB_CODENAME=") {
			fields := strings.SplitN(strings.TrimSpace(line), "=", 2)
			codename := fields[1]
			if codename != "" {
				if strings.HasPrefix(codename, "\"") && strings.HasSuffix(codename, "\"") {
					d.codename = capitalize(codename[1 : len(codename)-1])
					continue
				}
				d.codename = capitalize(codename)
			}
			// Check if DISTRIB_RELEASE= is defined in /etc/*release*
		} else if strings.HasPrefix(line, "DISTRIB_RELEASE=") {
			fields := strings.SplitN(strings.TrimSpace(line), "=", 2)
			version := fields[1]
			if version != "" {
				if strings.HasPrefix(version, "\"") && strings.HasSuffix(version, "\"") {
					if containsDigit(version) {
						d.version = version[1 : len(version)-1]
					}
					continue
				}
				if containsDigit(version) {
					d.version = version
				}
			}
		}

	}
	// The following checks are only performed if no distro is detected so far
	// TODO: Generate a list of all files in PATH before performing these checks
	if d.name == defaultName {
		// Executables related to package managers
		if Has("xbsp-query") {
			d.name = "Void Linux"
		} else if Has("pacman") {
			d.name = "Arch Linux"
		} else if Has("dnf") {
			d.name = "Fedora"
		} else if Has("yum") {
			d.name = "Fedora"
		} else if Has("zypper") {
			d.name = "openSUSE"
		} else if Has("emerge") {
			d.name = "Gentoo"
		} else if Has("apk") {
			d.name = "Alpine"
		} else if Has("slapt-get") || Has("slackpkg") {
			d.name = "Slackware"
		} else if d.platform == "Darwin" {
			productName := strings.TrimSpace(Run("sw_vers -productName"))
			// Set the platform to either "macOS" or "OS X", if it is in the product name
			if strings.HasPrefix(productName, "Mac OS X") {
				d.platform = "OS X"
			} else if strings.Contains(productName, "macOS") {
				d.platform = "macOS"
			} else {
				d.platform = productName
			}
			// Version number
			d.version = strings.TrimSpace(Run("sw_vers -productVersion"))
			// Codename, thanks @rubynorails! https://unix.stackexchange.com/a/234173/3920
			d.codename = strings.TrimSpace(Run("awk '/SOFTWARE LICENSE AGREEMENT FOR OS X/' '/System/Library/CoreServices/Setup Assistant.app/Contents/Resources/en.lproj/OSXSoftwareLicense.rtf' | awk -F 'OS X ' '{print $NF}' | awk '{print substr($0, 0, length($0)-1)}'"))
			// Mac doesn't really have a distro name
			d.name = ""
		} else if Has("/usr/sbin/pkg") {
			d.name = "FreeBSD"
			// rpm and dpkg-query should come last, since many distros may include them
		} else if Has("rpm") {
			d.name = "Red Hat"
		} else if Has("dpkg-query") {
			d.name = "Debian"
		}
	}
	return &d
}

// Grep /etc/*release* for the given string.
// If the search fails, a case-insensitive string search is attempted.
// The contents of /etc/*release* is cached.
func (d *Distro) Grep(name string) bool {
	return strings.Contains(d.etcRelease, name) || strings.Contains(strings.ToLower(d.etcRelease), strings.ToLower(name))
}

// Platform returns the name of the current platform.
// This is the same as `runtime.GOOS`, but capitalized.
func (d *Distro) Platform() string {
	return d.platform
}

// Name returns the detected name of the current distro/BSD, or "Unknown".
func (d *Distro) Name() string {
	return d.name
}

// Codename returns the detected codename of the current distro/BSD,
// or an empty string.
func (d *Distro) Codename() string {
	return d.codename
}

// Version returns the detected release version of the current distro/BSD,
// or an empty string.
func (d *Distro) Version() string {
	return d.version
}

// EtcRelease returns the contents of /etc/*release, or an empty string.
// The contents are cached.
func (d *Distro) EtcRelease() string {
	return d.etcRelease
}

// String returns a string with the current platform, distro
// codename and release version (if available).
// Example strings:
//   Linux (Ubuntu Bionic 18.04)
//   Darwin (10.13.3)
func (d *Distro) String() string {
	var sb strings.Builder
	sb.WriteString(d.platform)
	sb.WriteString(" ")
	if d.name != "" || d.codename != "" || d.version != "" {
		sb.WriteString("(")
		needSpace := false
		if d.name != defaultName && d.name != "" {
			sb.WriteString(d.name)
			needSpace = true
		}
		if d.codename != "" {
			if needSpace {
				sb.WriteString(" ")
			}
			sb.WriteString(d.codename)
			needSpace = true
		}
		if d.version != "" {
			if needSpace {
				sb.WriteString(" ")
			}
			sb.WriteString(d.version)
		}
		sb.WriteString(")")
	}
	return sb.String()
}
