package config

import "fmt"

func PublicPath(path string) string {
	return fmt.Sprintf("%s/%s", EnvConfig.PUBLIC_PATH, path)
}
