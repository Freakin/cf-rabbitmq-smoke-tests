package service_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudfoundry-incubator/cf-test-helpers/workflowhelpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testConfig struct {
	AdminUser         string   `json:"admin_user"`
	AdminPassword     string   `json:"admin_password"`
	Api               string   `json:"api"`
	AppsDomain        string   `json:"apps_domain"`
	OrgName           string   `json:"org_name"`
	SpaceName         string   `json:"space_name"`
	PlanNames         []string `json:"plan_names"`
	ServiceName       string   `json:"service_name"`
	SkipSSLValidation bool     `json:"skip_ssl_validation"`
	TimeoutScale      float64  `json:"timeout_scale"`
}

func (c testConfig) ScaledTimeout(timeout time.Duration) time.Duration {
	return time.Duration(float64(timeout) * c.TimeoutScale)
}

func (c testConfig) Username() string {
	return c.AdminUser
}

func (c testConfig) Password() string {
	return c.AdminPassword
}

func (c testConfig) OrganizationName() string {
	return c.OrgName
}

func (c testConfig) SpaceName() string {
	return c.SpaceName
}

func loadConfig(path string) (cfg testConfig) {
	configFile, err := os.Open(path)
	if err != nil {
		fatal(err)
	}

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&cfg); err != nil {
		fatal(err)
	}

	return
}

var (
	config = loadConfig(os.Getenv("CONFIG_PATH"))
	ctx    workflowhelpers.UserContext
)

func fatal(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
	os.Exit(1)
}

func TestService(t *testing.T) {
	// if err := services.ValidateConfig(&config.Config); err != nil {
	// 	fatal(err)
	// }

	//ctx = services.NewContext(config.Config, "rabbitmq-smoke-tests")
	ctx = workflowhelpers.NewUserContext(config.Api, config, config, config.SkipSSLValidation, 30*time.Second)

	RegisterFailHandler(Fail)

	RunSpecs(t, "RabbitMQ Smoke Tests")
}

var _ = BeforeEach(func() {
	//ctx.Setup()
	ctx.Login()
})

var _ = AfterEach(func() {
	//ctx.Teardown()
	ctx.Logout()
})
