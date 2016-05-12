package infrastructure

import "io"

//IaasPrep deploy infrastructure requirement to iaas
type IaasPrep func(config *io.Reader) error

//Deploy deploy infrastructure
func (iaas IaasPrep) Deploy(config *io.Reader) error {
	return iaas(config)
}
