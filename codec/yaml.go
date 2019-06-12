package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type MetaData struct {
	Kind string `yaml:"kind"`
}

type Instance struct {
	InstanceId string `json:"instance-id"`
}

func main() {
	//instance := Instance{
	//	InstanceId: "jvessel-adhcxiaah",
	//}
	//data, err := json.Marshal(instance)
	//if err != nil {
	//	fmt.Println("json marshal err:", err)
	//	return
	//}
	//// replace " by \"
	//jsonData := string(data)
	//jsonData = strings.Replace(jsonData, "\"", "\\\"", -1)
	// construct yaml content
	jsonData := "[{\"registry\":\"my-registry\",\"registryAddrs\":[\"10.12.209.44:8500\",\"10.12.209.43:8500\"],\"services\":[\"jdsf-server\"]}]"
	yamlContent := fmt.Sprintf("kind: \"%s\"", jsonData)
	fmt.Println(yamlContent)

	// yamlContent := `kind: "{'instance-id':'jvessel-adhcxiaah'}"`	 	// ERROR
	// yamlContent := `kind: '{'instance-id':'jvessel-adhcxiaah'}'` 	// ERROR
	// yamlContent := `kind: '{\"instance-id\":\"jvessel-adhcxiaah\"}'` // ERROR
	// yamlContent := `kind: {\"instance-id\":\"jvessel-adhcxiaah\"}` // ERROR
	// yamlContent := `kind: "{\"instance-id\":\"jvessel-adhcxiaah\"}"` // OK

	// unmarshal yaml
	metaData := new(MetaData)
	yaml.Unmarshal([]byte(yamlContent), metaData)

	// unmarshal json content
	//var instance1 Instance
	//err := json.Unmarshal([]byte(metaData.Kind), &instance1)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(instance1)
	//}
}
