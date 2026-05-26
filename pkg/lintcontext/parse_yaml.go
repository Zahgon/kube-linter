package lintcontext

import (
	"fmt"
	"io"
	"os"

	ocsAppsV1 "github.com/openshift/api/apps/v1"
	ocpSecV1 "github.com/openshift/api/security/v1"
	k8sMonitoring "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	kedaV1Alpha1 "golang.stackrox.io/kube-linter/pkg/crds/keda/v1alpha1"
	"golang.stackrox.io/kube-linter/pkg/k8sutil"
	"helm.sh/helm/v3/pkg/chart"
	autoscalingV2Beta1 "k8s.io/api/autoscaling/v2beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	// The max file size, in bytes, that we will load.
	// TODO: make it configurable.
	maxFileSizeBytes = 10 * 1024 * 1024
)

var (
	decoder runtime.Decoder
)

func init() {
	clientScheme := scheme.Scheme

	// Add OpenShift and Autoscaling schema
	schemeBuilder := runtime.NewSchemeBuilder(ocsAppsV1.AddToScheme, autoscalingV2Beta1.AddToScheme, k8sMonitoring.AddToScheme, ocpSecV1.AddToScheme, kedaV1Alpha1.AddToScheme)
	if err := schemeBuilder.AddToScheme(clientScheme); err != nil {
		panic(fmt.Sprintf("Can not add OpenShift schema %v", err))
	}
	decoder = serializer.NewCodecFactory(clientScheme).UniversalDeserializer()
}

func parseObjects(data []byte, d runtime.Decoder) ([]k8sutil.Object, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this is for backward compatibility, should be replaced with kubeconform

// fallback to unstructured as schema validation will be performed by kubeconform check

// TODO: validate

type nopWriter struct{}

func (w nopWriter) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (l *lintContextImpl) renderHelmChart(dir string) (map[string]string, error) {
	_ = "STUB: not implemented"
	// Helm doesn't have great logging behaviour, and can spam stderr, so silence their logging.
	// TODO: capture these logs.
	return nil, nil
}

func (l *lintContextImpl) renderValues(chrt *chart.Chart, values map[string]interface{}) (map[string]string, error) {
	_ = "STUB: not implemented"
	// Process chart dependencies to handle import-values
	return nil, nil
}

func (l *lintContextImpl) loadObjectsFromHelmChart(dir string, ignorePaths []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Paths returned by helm include redundant directory in front, therefore we strip it out.

func (l *lintContextImpl) loadObjectsFromTgzHelmChart(tgzFile string, ignorePaths []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *lintContextImpl) renderTgzHelmChart(tgzFile string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *lintContextImpl) loadObjectFromYAMLReader(filePath string, r *yaml.YAMLReader) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *lintContextImpl) loadObjectsFromYAMLFile(filePath string, info os.FileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *lintContextImpl) loadObjectsFromReader(filePath string, reader io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *lintContextImpl) renderChart(fileName string, chart *chart.Chart) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *lintContextImpl) renderTgzHelmChartReader(fileName string, tgzReader io.Reader) (map[string]string, error) {
	_ = "STUB: not implemented"
	// Helm doesn't have great logging behaviour, and can spam stderr, so silence their logging.
	return nil, nil
}

func (l *lintContextImpl) readObjectsFromTgzHelmChart(fileName string, tgzReader io.Reader, ignoredPaths []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *lintContextImpl) loadHelmRenderedTemplates(chartPath string, renderedFiles map[string]string, ignorePaths []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip NOTES.txt file that may be present among templates but is not a kubernetes resource.

// normalizeDirectoryPaths removes the first element of the path that gets added by the Helm library.
// Helm adds chart name as the first path component, however this is not always correct, e.g. in case the helm chart
// directory was renamed, as shown in https://github.com/stackrox/kube-linter/issues/212
// The function converts mychart/templates/deployment.yaml to templates/deployment.yaml.
func normalizeDirectoryPaths(renderedFiles map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Go does not seem to have a library function that allows to split the first element of path, therefore
// splitting "by hand" on path separator char, which is ok if you check path.Split() implementation ;-)

func (l *lintContextImpl) loadObjectsFromKustomize(dir string) {
	_ = "STUB: not implemented"
	// Create a kustomize engine with source annotations enabled
	return
}

// Ignore warnings, similar to how we suppress Helm logs with nopWriter

// Render the kustomize manifests

// Convert each object to YAML and load it

// Marshal the underlying object content to YAML

// Extract the source file from annotations if available

// Prefix with the kustomization directory to show full path

// Fallback: Create a file path for the object for reporting purposes

// Load the object using the existing YAML reader
