package integrationtest

import (
	"battle-of-kings/internal/repository/redis"
	"battle-of-kings/pkg/testhelper"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

var redisPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegration() {
		return
	}
	pool := testhelper.StartDockerPool()

	// set up the redis container for tests
	redisRes := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
			return err
		})
	redisPort = redisRes.GetPort("6379/tcp")

	// now run tests
	exitCode := m.Run()
	if err := redisRes.Close(); err != nil {
		logrus.WithError(err).Fatal("failed to close redis instance")
	}
	os.Exit(exitCode)

}
