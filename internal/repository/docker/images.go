package docker

import (
	"codeZone/internal/repository/docker/models"
	"fmt"
)

const (
	host = "docker.io"
)

// getActualImages returns slice of actual repository images
func getActualImages() map[string]*models.Image {
	var images map[string]*models.Image
	images = make(map[string]*models.Image)

	images["python"] = &models.Image{
		Title:   "python",
		Tag:     "3.12-alpine",
		Host:    host,
		FileExt: "py",
		Cmd: func(filename string) []string {
			return []string{"python", fmt.Sprintf("/tmp/%s", filename)}
		},
	}
	images["go"] = &models.Image{
		Title:   "golang",
		Tag:     "1.22.5-alpine",
		Host:    host,
		FileExt: "go",
		Cmd: func(filename string) []string {
			return []string{"go", "run", fmt.Sprintf("/tmp/%s", filename)}
		},
	}
	images["c++"] = &models.Image{
		Title:   "frolvlad/alpine-gxx",
		Tag:     "latest",
		Host:    host,
		FileExt: "cpp",
		Cmd: func(filename string) []string {
			return []string{"sh", "-c", fmt.Sprintf("g++ /tmp/%s -o main && ./main", filename)}
		},
	}
	images["c#"] = &models.Image{
		Title:   "mono",
		Tag:     "6.12",
		Host:    host,
		FileExt: "cs",
		Cmd: func(filename string) []string {
			return []string{"sh", "-c", fmt.Sprintf("mcs -out:main.exe /tmp/%s && mono main.exe", filename)}
		},
	}
	images["javascript"] = &models.Image{
		Title:   "node",
		Tag:     "22-alpine",
		Host:    host,
		FileExt: "js",
		Cmd: func(filename string) []string {
			return []string{"node", fmt.Sprintf("/tmp/%s", filename)}
		},
	}
	images["rust"] = &models.Image{
		Title:   "rust",
		Tag:     "1.79.0-alpine",
		Host:    host,
		FileExt: "rs",
		Cmd: func(filename string) []string {
			return []string{"sh", "-c", fmt.Sprintf("rustc /tmp/%s -o main && ./main", filename)}
		},
	}
	// FIXME: will work only if main class is "Main"
	images["java"] = &models.Image{
		Title:   "openjdk",
		Tag:     "24-slim",
		Host:    host,
		FileExt: "java",
		Cmd: func(filename string) []string {
			return []string{"sh", "-c", fmt.Sprintf("javac /tmp/%s && java -cp /tmp Main", filename)}
		},
	}

	return images
}
