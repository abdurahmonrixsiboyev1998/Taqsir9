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

func GetEmail() (string, error) {
	e := exec.Command("git", "config", "user.email")
	s, error := e.Output()
	if error != nil {
		return "", error
	}
	return string(s), nil
}
