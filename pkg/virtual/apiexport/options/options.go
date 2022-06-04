/*
Copyright 2022 The KCP Authors.

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

package options

import (
	"path"

	"github.com/spf13/pflag"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	kcpclientset "github.com/kcp-dev/kcp/pkg/client/clientset/versioned"
	kcpinformer "github.com/kcp-dev/kcp/pkg/client/informers/externalversions"
	"github.com/kcp-dev/kcp/pkg/virtual/apiexport/builder"
	"github.com/kcp-dev/kcp/pkg/virtual/framework"
	"github.com/kcp-dev/kcp/pkg/virtual/framework/rootapiserver"
)

type APIExport struct{}

func NewAPIExport() *APIExport {
	return &APIExport{}
}

func (o *APIExport) AddFlags(flags *pflag.FlagSet, prefix string) {
	if o == nil {
		return
	}
}

func (o *APIExport) Validate(flagPrefix string) []error {
	if o == nil {
		return nil
	}
	errs := []error{}

	return errs
}

func (o *APIExport) NewVirtualWorkspaces(
	rootPathPrefix string,
	kubeClusterClient kubernetes.ClusterInterface,
	dynamicClusterClient dynamic.ClusterInterface,
	kcpClusterClient kcpclientset.ClusterInterface,
	wildcardKcpInformers kcpinformer.SharedInformerFactory,
) (extraInformers []rootapiserver.InformerStart, workspaces []framework.VirtualWorkspace, err error) {
	virtualWorkspaces := []framework.VirtualWorkspace{
		builder.BuildVirtualWorkspace(path.Join(rootPathPrefix, o.Name()), kubeClusterClient, dynamicClusterClient, kcpClusterClient, wildcardKcpInformers),
	}
	return nil, virtualWorkspaces, nil
}

func (o *APIExport) Name() string {
	return builder.VirtualWorkspaceName
}