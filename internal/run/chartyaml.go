package run

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

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

// EditChartYaml edits the Chart.yaml file in the given chart path.
// The modifier function is called with the ChartYaml struct, and any changes
// made to it will be written back to the file.
// Example:
//
//	editChartYaml(chartPath, func(chartYaml *ChartYaml) {
//		chartYaml.Version = "1.2.3"
//	})
func EditChartYaml(chartPath string, modifier func(*ChartYaml)) {

	yamlData, err := ioutil.ReadFile(chartPath + "/Chart.yaml")
	if err != nil {
		println("Try to read a file " + err.Error())
	}
	var chartYaml ChartYaml
	if err := yaml.Unmarshal(yamlData, &chartYaml); err != nil {
		println("Try to unmarshall the file: " + err.Error())
	}

	modifier(&chartYaml)

	updateYaml, err := yaml.Marshal(chartYaml)
	if err != nil {
		println("Try to marshall the file " + err.Error())
	}

	err = ioutil.WriteFile(chartPath+"/Chart.yaml", updateYaml, os.ModePerm)
	if err != nil {
		println("try to write the file" + err.Error())
	}

}
