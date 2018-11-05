package distrodetector

import (
	"os/exec"
)

// Which returns the full path to the given executable, or the original string
func which(executable string) string {
	path, err := exec.LookPath(executable)
	if err != nil {
		return executable
	}
	return path
}

// Run a shell command and return the output, or an empty string
func run(shellCommand string) string {
	cmd := exec.Command("sh", "-c", shellCommand)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(stdoutStderr)
}

func (d *Distro) FindPackageByName(name string) []string {
	switch d.name {
	case "Arch Linux":
		return strings.Split(run("pacman -Ss " + name), "\n")
	case "Debian":
		// TODO
		return strings.Split(run("dpkg -Ss " + name), "\n")
	default:
		return []string{}
	}
}

// Return the path to the current package manager, if detected
func (d *Distro) PackageManager() string {
	switch d.name {
	case "Arch Linux":
		return "/usr/bin/pacman"
	}
	if which("dpkg") {
		return "/usr/bin/dpkg"
	}
}
