package main

import (
	"fmt"
	"strings"

	"github.com/datianshi/pcfdeploy/infrastructure/aws"
)

func main() {
	//Prepare the infrastructure
	config := strings.NewReader("Hello")
	aws.CloudFormation.Deploy(config)
	fmt.Println("Hello World")
}
