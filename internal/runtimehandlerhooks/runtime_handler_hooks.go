package runtimehandlerhooks

import (
	"context"
	"strings"

	"github.com/cri-o/cri-o/internal/lib/sandbox"
	"github.com/cri-o/cri-o/internal/log"
	"github.com/cri-o/cri-o/internal/oci"
	crioann "github.com/cri-o/cri-o/pkg/annotations"
	libconfig "github.com/cri-o/cri-o/pkg/config"
)

type RuntimeHandlerHooks interface {
	PreStart(ctx context.Context, c *oci.Container, s *sandbox.Sandbox) error
	PreStop(ctx context.Context, c *oci.Container, s *sandbox.Sandbox) error
}

// GetRuntimeHandlerHooks returns RuntimeHandlerHooks implementation by the runtime handler name
func GetRuntimeHandlerHooks(ctx context.Context, config *libconfig.Config, handler string, annotations map[string]string) (RuntimeHandlerHooks, error) {
	ctx, span := log.StartSpan(ctx)
	defer span.End()
	if strings.Contains(handler, HighPerformance) {
		log.Warnf(ctx, "The usage of the handler %q without adding high-performance feature annotations under allowed_annotations will be deprecated under 1.21", HighPerformance)
		return &HighPerformanceHooks{config.IrqBalanceConfigFile}, nil
	}
	if highPerformanceAnnotationsSpecified(annotations) {
		log.Warnf(ctx, "The usage of the handler %q without adding high-performance feature annotations under allowed_annotations will be deprecated under 1.21", HighPerformance)
		return &HighPerformanceHooks{config.IrqBalanceConfigFile}, nil
	}

	return nil, nil
}

func highPerformanceAnnotationsSpecified(annotations map[string]string) bool {
	for k := range annotations {
		if strings.HasPrefix(k, crioann.CPULoadBalancingAnnotation) ||
			strings.HasPrefix(k, crioann.CPUQuotaAnnotation) ||
			strings.HasPrefix(k, crioann.IRQLoadBalancingAnnotation) ||
			strings.HasPrefix(k, crioann.CPUCStatesAnnotation) ||
			strings.HasPrefix(k, crioann.CPUFreqGovernorAnnotation) {
			return true
		}
	}
	return false
}
