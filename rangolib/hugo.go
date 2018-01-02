package rangolib

import "os/exec"

func RunHugo(path, cwd string) ([]byte, error) {
	if path == "" {
		path = "hugo"
	}
	hugo := exec.Command(path)
	hugo.Dir = cwd

	output, err := hugo.Output()
	if err != nil {
		return nil, err
	}

	return output, nil
}
