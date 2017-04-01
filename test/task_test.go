package test_test

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"gopkg.in/yaml.v2"
)

var _ = Describe("test", func() {
	var defaultEnvironmentEnv = map[string]string{
		"BOSH_CA_CERT":       "fake-ca_cert",
		"BOSH_CLIENT":        "fake-client",
		"BOSH_CLIENT_SECRET": "fake-client_secret",
		"BOSH_DEPLOYMENT":    "fake-deployment",
		"BOSH_ENVIRONMENT":   "fake-environment",
	}

	Describe("clean-up-all.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("clean-up-all.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"clean-up",
				},
			}))
		})

		It("executes", func() {
			result := executeConfig("clean-up-all.yml", defaultEnvironmentParams())

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env: map[string]string{
					"BOSH_CA_CERT":       "fake-ca_cert",
					"BOSH_CLIENT":        "fake-client",
					"BOSH_CLIENT_SECRET": "fake-client_secret",
					"BOSH_ENVIRONMENT":   "fake-environment",
				},
				Arg: []string{
					"clean-up",
					"--all",
				},
			}))
		})
	})

	Describe("clean-up.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("clean-up.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"clean-up",
				},
			}))
		})

		It("executes", func() {
			result := executeConfig("clean-up.yml", defaultEnvironmentParams())

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env: map[string]string{
					"BOSH_CA_CERT":       "fake-ca_cert",
					"BOSH_CLIENT":        "fake-client",
					"BOSH_CLIENT_SECRET": "fake-client_secret",
					"BOSH_ENVIRONMENT":   "fake-environment",
				},
				Arg: []string{
					"clean-up",
				},
			}))
		})
	})

	Describe("recreate.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("recreate.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"recreate",
				},
			}))
		})

		It("executes for deployment", func() {
			result := executeConfig("recreate.yml", defaultDeploymentParams())

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"recreate",
				},
			}))
		})

		It("executes for instance group", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"

			result := executeConfig("recreate.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"recreate",
					"fake-instance_group",
				},
			}))
		})

		It("executes for instance", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"
			params["instance_id"] = "fake-instance_id"

			result := executeConfig("recreate.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"recreate",
					"fake-instance_group/fake-instance_id",
				},
			}))
		})
	})

	Describe("run-errand.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("run-errand.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"run-errand",
				},
			}))
		})

		It("executes for deployment", func() {
			params := defaultDeploymentParams()
			params["errand"] = "fake-errand"

			result := executeConfig("run-errand.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"run-errand",
					"fake-errand",
				},
			}))
		})
	})

	Describe("ssh.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("ssh.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"ssh",
				},
			}))
		})

		It("executes for deployment", func() {
			params := defaultDeploymentParams()
			params["command"] = `echo "uptime's "; uptime`

			result := executeConfig("ssh.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"ssh",
					"-c",
					`echo "uptime's "; uptime`,
				},
			}))
		})

		It("executes for instance group", func() {
			params := defaultDeploymentParams()
			params["command"] = `echo "uptime's "; uptime`
			params["instance_group"] = "fake-instance_group"

			result := executeConfig("ssh.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"ssh",
					"-c",
					`echo "uptime's "; uptime`,
					"fake-instance_group",
				},
			}))
		})

		It("executes for instance", func() {
			params := defaultDeploymentParams()
			params["command"] = `echo "uptime's "; uptime`
			params["instance_group"] = "fake-instance_group"
			params["instance_id"] = "fake-instance_id"

			result := executeConfig("ssh.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"ssh",
					"-c",
					`echo "uptime's "; uptime`,
					"fake-instance_group/fake-instance_id",
				},
			}))
		})
	})

	Describe("start.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("start.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"start",
				},
			}))
		})

		It("executes for deployment", func() {
			result := executeConfig("start.yml", defaultDeploymentParams())

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"start",
				},
			}))
		})

		It("executes for instance group", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"

			result := executeConfig("start.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"start",
					"fake-instance_group",
				},
			}))
		})

		It("executes for instance", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"
			params["instance_id"] = "fake-instance_id"

			result := executeConfig("start.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"start",
					"fake-instance_group/fake-instance_id",
				},
			}))
		})
	})

	Describe("stop.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("stop.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"stop",
				},
			}))
		})

		It("executes for deployment", func() {
			result := executeConfig("stop.yml", defaultDeploymentParams())

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"stop",
				},
			}))
		})

		It("executes for instance group", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"

			result := executeConfig("stop.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"stop",
					"fake-instance_group",
				},
			}))
		})

		It("executes for instance", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"
			params["instance_id"] = "fake-instance_id"

			result := executeConfig("stop.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"stop",
					"fake-instance_group/fake-instance_id",
				},
			}))
		})
	})

	Describe("take-snapshot.yml", func() {
		It("propagates errors", func() {
			params := defaultDeploymentParams()
			params["_test_bosh_exit"] = "2"

			result := executeConfig("take-snapshot.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 2,
				Env:  map[string]string{},
				Arg: []string{
					"take-snapshot",
				},
			}))
		})

		It("executes for deployment", func() {
			result := executeConfig("take-snapshot.yml", defaultDeploymentParams())

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"take-snapshot",
				},
			}))
		})

		It("executes for instance group", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"

			result := executeConfig("take-snapshot.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"take-snapshot",
					"fake-instance_group",
				},
			}))
		})

		It("executes for instance", func() {
			params := defaultDeploymentParams()
			params["instance_group"] = "fake-instance_group"
			params["instance_id"] = "fake-instance_id"

			result := executeConfig("take-snapshot.yml", params)

			Expect(result).To(Equal(executeResult{
				Exit: 0,
				Env:  defaultEnvironmentEnv,
				Arg: []string{
					"take-snapshot",
					"fake-instance_group/fake-instance_id",
				},
			}))
		})
	})
})

type executeResult struct {
	Exit int
	Env  map[string]string
	Arg  []string
}

var reEnvArgLine = regexp.MustCompile("^(env|arg): (.+)$")

func executeConfig(file string, params map[string]string) executeResult {
	_, path, _, _ := runtime.Caller(0)

	configBytes, err := ioutil.ReadFile(filepath.Join(filepath.Dir(filepath.Dir(path)), file))
	Expect(err).ToNot(HaveOccurred())

	configStruct := struct {
		Run struct {
			Path string
			Args []string
		}
		Params map[string]string
	}{}

	err = yaml.Unmarshal(configBytes, &configStruct)
	Expect(err).ToNot(HaveOccurred())

	args := []string{"run", "--rm"}

	for key, val := range configStruct.Params {
		if _, ok := params[key]; ok {
			continue
		}

		params[key] = val
	}

	for key, val := range params {
		args = append(args, "-e", fmt.Sprintf("%s=%s", key, val))
	}

	args = append(args, "dpb587/bosh-concourse-tasks:master-spec", configStruct.Run.Path)

	for _, arg := range configStruct.Run.Args {
		args = append(args, arg)
	}

	command := exec.Command("/usr/local/bin/docker", args...)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	session.Wait(15 * time.Second)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit())

	result := executeResult{
		Exit: session.ExitCode(),
		Env:  map[string]string{},
		Arg:  []string{},
	}

	for _, line := range strings.Split(strings.TrimSpace(string(session.Out.Contents())), "\n") {
		match := reEnvArgLine.FindStringSubmatch(line)

		if len(match) > 0 {
			if match[1] == "env" {
				kv := strings.SplitN(match[2], "=", 2)

				result.Env[kv[0]] = kv[1]
			} else if match[1] == "arg" {
				result.Arg = append(result.Arg, match[2])
			}
		}
	}

	return result
}

func defaultEnvironmentParams() map[string]string {
	return map[string]string{
		"ca_cert":       "fake-ca_cert",
		"client":        "fake-client",
		"client_secret": "fake-client_secret",
		"environment":   "fake-environment",
	}
}

func defaultDeploymentParams() map[string]string {
	params := defaultEnvironmentParams()
	params["deployment"] = "fake-deployment"

	return params
}
