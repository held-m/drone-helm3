package run

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelotech/drone-helm3/internal/env"
	"gopkg.in/yaml.v2"
)

// Upgrade is an execution step that calls `helm upgrade` when executed.
type Upgrade struct {
	*config
	chart   string
	release string

	chartVersion    string
	appVersion      string
	dryRun          bool
	wait            bool
	values          string
	stringValues    string
	valuesFiles     []string
	reuseValues     bool
	timeout         string
	force           bool
	atomic          bool
	cleanupOnFail   bool
	historyMax      int
	certs           *repoCerts
	createNamespace bool
	skipCrds        bool

	cmd cmd
}

// NewUpgrade creates an Upgrade using fields from the given Config. No validation is performed at this time.
func NewUpgrade(cfg env.Config) *Upgrade {
	return &Upgrade{
		config:          newConfig(cfg),
		chart:           cfg.Chart,
		release:         cfg.Release,
		chartVersion:    cfg.ChartVersion,
		appVersion:      cfg.AppVersion,
		dryRun:          cfg.DryRun,
		wait:            cfg.Wait,
		values:          cfg.Values,
		stringValues:    cfg.StringValues,
		valuesFiles:     cfg.ValuesFiles,
		reuseValues:     cfg.ReuseValues,
		timeout:         cfg.Timeout,
		force:           cfg.Force,
		atomic:          cfg.AtomicUpgrade,
		cleanupOnFail:   cfg.CleanupOnFail,
		historyMax:      cfg.HistoryMax,
		certs:           newRepoCerts(cfg),
		createNamespace: cfg.CreateNamespace,
		skipCrds:        cfg.SkipCrds,
	}
}

// Execute executes the `helm upgrade` command.
func (u *Upgrade) Execute() error {
	return u.cmd.Run()
}

// Prepare gets the Upgrade ready to execute.
func (u *Upgrade) Prepare() error {
	if u.chart == "" {
		return fmt.Errorf("chart is required")
	}
	if u.release == "" {
		return fmt.Errorf("release is required")
	}

	args := u.globalFlags()
	args = append(args, "upgrade", "--install")

	if u.chartVersion != "" {
		args = append(args, "--version", u.chartVersion)
		addChartVersion(u.chartVersion, u.chart)
	}

	if u.appVersion != "" {
		addAppVersion(u.appVersion, u.chart)
	}

	if u.dryRun {
		args = append(args, "--dry-run")
	}
	if u.wait {
		args = append(args, "--wait")
	}
	if u.reuseValues {
		args = append(args, "--reuse-values")
	}
	if u.timeout != "" {
		args = append(args, "--timeout", u.timeout)
	}
	if u.force {
		args = append(args, "--force")
	}
	if u.atomic {
		args = append(args, "--atomic")
	}
	if u.cleanupOnFail {
		args = append(args, "--cleanup-on-fail")
	}
	if u.values != "" {
		args = append(args, "--set", u.values)
	}
	if u.stringValues != "" {
		args = append(args, "--set-string", u.stringValues)
	}
	if u.createNamespace {
		args = append(args, "--create-namespace")
	}
	if u.skipCrds {
		args = append(args, "--skip-crds")
	}
	for _, vFile := range u.valuesFiles {
		args = append(args, "--values", vFile)
	}
	args = append(args, u.certs.flags()...)

	// always set --history-max since it defaults to non-zero value
	args = append(args, fmt.Sprintf("--history-max=%d", u.historyMax))

	args = append(args, u.release, u.chart)
	u.cmd = command(helmBin, args...)
	u.cmd.Stdout(u.stdout)
	u.cmd.Stderr(u.stderr)

	if u.debug {
		fmt.Fprintf(u.stderr, "Generated command: '%s'\n", u.cmd.String())
	}

	return nil
}

// ChartYaml is the struct representation of a Chart.yaml file.
type ChartYaml struct {
	APIVersion  string   `yaml:"apiVersion"`
	Name        string   `yaml:"name"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	AppVersion  string   `yaml:"appVersion,omitempty"`
	Keywords    []string `yaml:"keywords,omitempty"`
	Home        string   `yaml:"home,omitempty"`
	Icon        string   `yaml:"icon,omitempty"`
	Sources     []string `yaml:"sources,omitempty"`
	Type        string   `yaml:"type,omitempty"`
	Maintainers []struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
		URL   string `yaml:"url,omitempty"`
	} `yaml:"maintainers,omitempty"`
	Dependencies []struct {
		Name       string   `yaml:"name"`
		Version    string   `yaml:"version"`
		Repository string   `yaml:"repository"`
		Condition  string   `yaml:"condition,omitempty"`
		Tags       []string `yaml:"tags,omitempty"`
	} `yaml:"dependencies,omitempty"`
}

func addAppVersion(appVersion, chartPath string) {

	println("Updating Chart.yaml appVersion to " + appVersion)

	yamlData, err := ioutil.ReadFile(chartPath + "/Chart.yaml")
	if err != nil {
		println("Try to read a file " + err.Error())
	}
	var chartYaml ChartYaml
	if err := yaml.Unmarshal(yamlData, &chartYaml); err != nil {
		println("Try to unmarshall the file: " + err.Error())
	}

	chartYaml.AppVersion = appVersion

	updateYaml, err := yaml.Marshal(chartYaml)
	if err != nil {
		println("Try to marshall the file " + err.Error())
	}

	err = ioutil.WriteFile(chartPath+"/Chart.yaml", updateYaml, os.ModePerm)
	if err != nil {
		println("try to write the file" + err.Error())
	}

}

func addChartVersion(chartVersion, chartPath string) {

	println("Updating Chart.yaml version to " + chartVersion)

	yamlData, err := ioutil.ReadFile(chartPath + "/Chart.yaml")
	if err != nil {
		println("Try to read a file " + err.Error())
	}
	var chartYaml ChartYaml
	if err := yaml.Unmarshal(yamlData, &chartYaml); err != nil {
		println("Try to unmarshall the file: " + err.Error())
	}

	chartYaml.Version = chartVersion

	updateYaml, err := yaml.Marshal(chartYaml)
	if err != nil {
		println("Try to marshall the file " + err.Error())
	}

	err = ioutil.WriteFile(chartPath+"/Chart.yaml", updateYaml, os.ModePerm)
	if err != nil {
		println("try to write the file" + err.Error())
	}

}
