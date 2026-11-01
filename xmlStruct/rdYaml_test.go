package xmlStruct

import (
	"os"
	"testing"
)


func TestRdYaml(t*testing.T) {
	yamlFilNam := "../yaml/testConfig.yaml"

	yamlTest, err := os.ReadFile(yamlFilNam)
	if err != nil {t.Errorf("could not open test file: %v", err)}

	if len(yamlTest) ==0 {t.Errorf("no content in test file!")}

	cfg, err := RdYaml(yamlFilNam)
	if err != nil {t.Errorf("could not decode yaml file: %v", err)}

	config := cfg.Config
	if config.AttrNameSuffix != "" {t.Errorf("should empty, but is %s", config.AttrNameSuffix)}
	if config.Header != "head" {t.Errorf("should be 'head', but is %s", config.Header)}
	if config.PackageName != "test yaml" {t.Errorf("should be 'test yaml', but is %s", config.PackageName)}

	PrintConfig(cfg)
}
