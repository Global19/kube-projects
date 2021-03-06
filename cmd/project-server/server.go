package main

import (
	"flag"
	"os"
	"runtime"

	"k8s.io/apiserver/pkg/util/logs"

	"github.com/openshift/kube-projects/pkg/cmd/server"

	// install all APIs
	_ "github.com/openshift/kube-projects/pkg/apis/project/install"
	_ "github.com/openshift/kube-projects/pkg/project/auth"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()
	// defer serviceability.BehaviorOnPanic(os.Getenv("OPENSHIFT_ON_PANIC"))()
	// defer serviceability.Profile(os.Getenv("OPENSHIFT_PROFILE")).Stop()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	cmd := server.NewCommandStartProjectServer(os.Stdout)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
