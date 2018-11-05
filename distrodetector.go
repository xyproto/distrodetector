package distrodetector

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Used when checking for Linux distros and BSDs (and NAME= is not defined in /etc)
var distroNames = []string{"Arch Linux", "Debian", "Ubuntu", "Void Linux", "FreeBSD", "NetBSD", "OpenBSD", "Manjaro", "Mint", "Elementary", "MX Linuyx", "Fedora", "openSUSE", "Solus", "Zorin", "CentOS", "KDE neon", "Lite", "Kali", "Antergos", "antiX", "Lubuntu", "PCLinuxOS", "Endless", "Peppermint", "SmartOS", "TrueOS", "Arco", "SparkyLinux", "deepin", "Puppy", "Slackware", "Bodhi", "Tails", "Xubuntu", "Archman", "Bluestar", "Mageia", "Deuvan", "Parrot", "Pop!", "ArchLabs", "Q4OS", "Kubuntu", "Nitrux", "Red Hat", "4MLinux", "Gentoo", "Pinguy", "LXLE", "KaOS", "Ultimate", "Alpine", "Feren", "KNOPPIX", "Robolinux", "Voyager", "Netrunner", "GhostBSD", "Budgie", "ClearOS", "Gecko", "SwagArch", "Emmabuntüs", "Scientific", "Omarine", "Neptune", "NixOS", "Slax", "Clonezilla", "DragonFly", "ExTiX", "OpenBSD", "Redcore", "Ubuntu Studio", "BunsenLabs", "BlackArch", "NuTyX", "ArchBang", "BackBox", "Sabayon", "AUSTRUMI", "Container", "ROSA", "SteamOS", "Tiny Core", "Kodachi", "Qubes", "siduction", "Parabola", "Trisquel", "Vector", "SolydXK", "Elive", "AV Linux", "Artix", "Raspbian", "Porteus"}

// Distro represents the platform, contents of /etc/*release* and name of the
// detected Linux distribution or BSD.
type Distro struct {
	platform   string
	etcRelease string
	name       string
}

// Has returns the full path to the given executable, or the original string
func Has(executable string) bool {
	_, err := exec.LookPath(executable)
	if err != nil {
		return false
	}
	return true
}

// etcRelease returns the contents of /etc/*release*, or an empty string
func etcRelease() string {
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
	d.platform = runtime.GOOS
	d.etcRelease = etcRelease()
	// Distro name, if not detected
	d.name = "Unknown"
	// Check for Linux distros and BSD distros
	for _, distroName := range distroNames {
		if d.Grep(distroName) {
			d.name = distroName
			break
		}
	}
	// Check if NAME= is defined in /etc/*release*
	for _, line := range strings.Split(d.etcRelease, "\n") {
		if strings.HasPrefix(line, "NAME=") {
			fields := strings.SplitN(strings.TrimSpace(line), "=", 2)
			name := fields[1]
			if name != "" {
				if strings.HasPrefix(name, "\"") && strings.HasSuffix(name, "\"") {
					d.name = name[1 : len(name)-1]
					break
				}
				d.name = name
				break
			}
		}
	}
	// The following checks are only performed if no distro is detected so far
	if d.name == "Unknown" {
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
		} else if d.platform == "darwin" && Has("homebrew") {
			d.name = "Homebrew"
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

// Search /etc/*release* for the given string.
// If the search fails, a case-insensitive string search is attempted.
// The contents of /etc/*release* is cached.
func (d *Distro) Grep(name string) bool {
	return strings.Contains(d.etcRelease, name) || strings.Contains(strings.ToLower(d.etcRelease), strings.ToLower(name))
}

// Platform returns the name of the current platform.
// Same as runtime.GOOS
func (d *Distro) Platform() string {
	return d.platform
}

// Name returns the detected name of the current distro/BSD, or "Unknown"
func (d *Distro) Name() string {
	return d.name
}

// String returns the current platform and distro as a string
func (d *Distro) String() string {
	return fmt.Sprintf("%s (%s)", d.platform, d.name)
}
