package testhelper

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type RetryFunc func(res *dockertest.Resource) error

func IsIntegration() bool {
	return os.Getenv("TEST_INTEGRATION") == "true"
}

func StartDockerPool() *dockertest.Pool {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatalf("Could not construct pool")
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		logrus.WithError(err).Fatalf("Could not connect to Docker")
	}
	return pool
}

func StartDockerInstance(pool *dockertest.Pool, image string, tag string, retryFunc RetryFunc, env ...string) *dockertest.Resource {

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag:        tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	if err := resource.Expire(120); err != nil {
		logrus.WithError(err).Fatal("Could not set resource expiration")
	}
	if err := pool.Retry(func() error {
		return retryFunc(resource)
	}); err != nil {
		logrus.WithError(err).Fatal("Couldn't connect to the resource")
	}
	return resource
}
