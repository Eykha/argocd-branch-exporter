package main

import (
	"fmt"
	"github.com/Eykha/argocd-branch-exporter/internal/metrics"
)

func main() {
	fmt.Println("Hello, world!")
	metrics.ServeMetrics()
}
