package aws

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

//Config Parse the config
type Config struct {
	Template Template `yaml:"cloudformation"`
}

//Template CloudFormation info
type Template struct {
}

//Parse parse the yaml config file
func Parse(cfg io.Reader) (config *Config, err error) {
	var bytes []byte
	if bytes, err = ioutil.ReadAll(cfg); err != nil {
		err = yaml.Unmarshal(bytes, &config)
	}
	return
}

// func (config *Config) BuildConfigProvider() *client.ConfigProvider {
// 	return nil
// }
//
// func (config *Config) BuildAwsConfig() *aws.Config {
// 	return nil
// }
//
// func (config *Config) BuildStackInput() cloudformation.CreateStackInput {
// 	return nil
// }
