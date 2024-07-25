package docker

import (
	"bytes"
	"codeZone/internal/repository"
	"codeZone/internal/repository/docker/models"
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	dockerTimeout = 1 * time.Second
)

var ErrUnsupported = errors.New("Unsupported language")

// repo local implementation of DockerRepository
type repo struct {
	cli    *client.Client
	images map[string]*models.Image
}

// NewRepository creates new repository
func NewRepository(cli *client.Client) repository.DockerRepository {
	return &repo{
		cli:    cli,
		images: getActualImages(),
	}
}

// DownloadImages downloads repo images
func (r *repo) DownloadImages(ctx context.Context) {
	log.Println("downloading images...")

	var wg sync.WaitGroup
	wg.Add(len(r.images))

	var failed = false

	for _, img := range r.images {
		go func(img *models.Image) {
			defer func() {
				wg.Done()
			}()

			out, err := r.cli.ImagePull(ctx, img.ImageWithHost(), image.PullOptions{})
			io.Copy(os.Stdout, out)

			if err != nil {
				log.Printf("error downloading %s image: %v\n", img.Title, err)
				failed = true
			}
		}(img)
	}
	wg.Wait()

	if failed {
		os.Exit(1)
	}
	log.Println("images were downloaded!\n")
}

// RunContainer creates and runs container
func (r *repo) RunContainer(ctx context.Context, info *models.Info) (string, error) {
	containerName := uuid.New().String()
	img, ok := r.images[info.Language]
	if !ok {
		return "", ErrUnsupported
	}

	dir, err := os.MkdirTemp("", "tmp-*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dir)

	file, err := os.CreateTemp(dir, fmt.Sprintf("main-*.%s", img.FileExt))
	if err != nil {
		return "", err
	}
	_, err = file.Write([]byte(info.Code))
	if err != nil {
		return "", err
	}

	m := mount.Mount{
		Type:   mount.TypeBind,
		Source: dir,
		Target: "/app",
	}

	filename := file.Name()[strings.LastIndex(file.Name(), "\\")+1:]
	resp, err := r.cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: img.ImageWithHost(),
			Cmd:   img.Cmd(filename),
		}, &container.HostConfig{
			Mounts: []mount.Mount{m},
		}, nil, nil, containerName,
	)

	if err != nil {
		return "", err
	}

	err = r.cli.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

// CheckContainerLogs checks for container logs, if too long, cancels
func (r *repo) CheckContainerLogs(ctx context.Context, id string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, dockerTimeout)
	defer cancel()

	statusCh, errCh := r.cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", err
		}
	case <-statusCh:
	}

	out, err := r.cli.ContainerLogs(
		ctx,
		id,
		container.LogsOptions{
			ShowStdout: true,
			ShowStderr: true,
		},
	)
	if err != nil {
		return "", err
	}
	defer out.Close()

	var dst = &bytes.Buffer{}

	_, err = stdcopy.StdCopy(dst, dst, out)
	if err != nil {
		return "", err
	}

	err = r.cli.ContainerRemove(ctx, id, container.RemoveOptions{})
	if err != nil {
		return "", err
	}

	return dst.String(), nil
}
