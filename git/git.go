package git

import "os/exec"

func GetUserName() (string, error) {
	n := exec.Command("git", "config", "user.name")
	s, error := n.Output()
	if error != nil {
		return "", error
	}
	return string(s), nil
}
