/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mke2e

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/kueue/test/util"
)

var (
	k8sManagerClient client.Client
	k8sWorker1Client client.Client
	k8sWorker2Client client.Client
	ctx              context.Context
)

func TestAPIs(t *testing.T) {
	suiteName := "End To End MultiKueue Suite"
	if ver, found := os.LookupEnv("E2E_KIND_VERSION"); found {
		suiteName = fmt.Sprintf("%s: %s", suiteName, ver)
	}
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t,
		suiteName,
	)
}

var _ = ginkgo.BeforeSuite(func() {
	managerClusterName := os.Getenv("MANAGER_KIND_CLUSTER_NAME")
	gomega.Expect(managerClusterName).NotTo(gomega.BeEmpty(), "MANAGER_KIND_CLUSTER_NAME should not be empty")

	worker1ClusterName := os.Getenv("WORKER1_KIND_CLUSTER_NAME")
	gomega.Expect(worker1ClusterName).NotTo(gomega.BeEmpty(), "WORKER1_KIND_CLUSTER_NAME should not be empty")

	worker2ClusterName := os.Getenv("WORKER2_KIND_CLUSTER_NAME")
	gomega.Expect(worker1ClusterName).NotTo(gomega.BeEmpty(), "WORKER2_KIND_CLUSTER_NAME should not be empty")

	k8sManagerClient = util.CreateClientUsingCluster("kind-" + managerClusterName)
	k8sWorker1Client = util.CreateClientUsingCluster("kind-" + worker1ClusterName)
	k8sWorker2Client = util.CreateClientUsingCluster("kind-" + worker2ClusterName)

	ctx = context.Background()

	//wait for the managers to start
	// failing a this point might indicate a manifestation of
	// https://kind.sigs.k8s.io/docs/user/known-issues/#pod-errors-due-to-too-many-open-files
	util.KueueReadyForTesting(ctx, k8sManagerClient)
	util.KueueReadyForTesting(ctx, k8sWorker1Client)
	util.KueueReadyForTesting(ctx, k8sWorker2Client)
})
