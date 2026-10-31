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
}
