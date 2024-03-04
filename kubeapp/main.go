package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type info struct {
	cpu    string
	memory string
}

type container struct {
}

type Obj struct {
}

func main() {
	cmd := exec.Command("kubectl", "get", "--raw", "/apis/metrics.k8s.io/v1beta1/namespaces/traceable/pods/traceable-ebpf-tracer-ds-hnf5c")
	out, _ := cmd.Output()
	fmt.Println(string(out))
	json.Unmarshaler()
}
