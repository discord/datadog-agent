// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package metrics

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/metrics"
	"github.com/stretchr/testify/assert"
)

// TestOriginServiceMappingParity tests that the originServiceToMetricSource and
// metricSourceToOriginService functions maintain consistent behavior.
//
// The test verifies that for all MetricSource values from 0 to 1024:
//  1. Converting MetricSource -> OriginService -> MetricSource maintains idempotency
//     (the result maps back to the same OriginService)
//  2. The mapping is consistent in both directions
func TestOriginServiceMappingParity(t *testing.T) {
	for i := uint16(0); i <= 1024; i++ {
		ms := metrics.MetricSource(i)

		// Convert MetricSource to OriginService
		originService := metricSourceToOriginService(ms)

		// Convert back to MetricSource
		reversedMS := originServiceToMetricSource(originService)

		// Convert the reversed MetricSource back to OriginService
		reversedOriginService := metricSourceToOriginService(reversedMS)

		// The key invariant: the origin service values must match
		// This ensures idempotency: even if multiple MetricSources map to the same
		// OriginService, the reverse mapping should always produce a MetricSource
		// that maps back to the same OriginService.
		assert.Equal(t, originService, reversedOriginService,
			"Parity check failed for MetricSource(%d): "+
				"MetricSource(%d) -> OriginService(%d) -> MetricSource(%d) -> OriginService(%d). "+
				"Expected both OriginService values to match.",
			i, i, originService, reversedMS, reversedOriginService)
	}
}

// TestOriginServiceToMetricSourceKnownValues tests specific known mappings
// to ensure correctness of the reverse mapping.
func TestOriginServiceToMetricSourceKnownValues(t *testing.T) {
	tests := []struct {
		name          string
		originService int32
		expectedMS    metrics.MetricSource
	}{
		{
			name:          "Unknown maps to MetricSourceUnknown",
			originService: 0,
			expectedMS:    metrics.MetricSourceUnknown,
		},
		{
			name:          "JmxCustom",
			originService: 9,
			expectedMS:    metrics.MetricSourceJmxCustom,
		},
		{
			name:          "ActiveDirectory",
			originService: 10,
			expectedMS:    metrics.MetricSourceActiveDirectory,
		},
		{
			name:          "ActivemqXML",
			originService: 11,
			expectedMS:    metrics.MetricSourceActivemqXML,
		},
		{
			name:          "GPU",
			originService: 466,
			expectedMS:    metrics.MetricSourceGPU,
		},
		{
			name:          "CloudFoundry",
			originService: 440,
			expectedMS:    metrics.MetricSourceCloudFoundry,
		},
		{
			name:          "Postgres",
			originService: 128,
			expectedMS:    metrics.MetricSourcePostgres,
		},
		{
			name:          "Docker",
			originService: 183,
			expectedMS:    metrics.MetricSourceDocker,
		},
		{
			name:          "Kubernetes",
			originService: 99,
			expectedMS:    metrics.MetricSourceKubernetesState,
		},
		{
			name:          "OpenTelemetryCollectorPrometheusReceiver",
			originService: 238,
			expectedMS:    metrics.MetricSourceOpenTelemetryCollectorPrometheusReceiver,
		},
		{
			name:          "Invalid/unmapped value returns Unknown",
			originService: 9999,
			expectedMS:    metrics.MetricSourceUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := originServiceToMetricSource(tt.originService)
			assert.Equal(t, tt.expectedMS, result,
				"originServiceToMetricSource(%d) should return %v, got %v",
				tt.originService, tt.expectedMS, result)
		})
	}
}

// TestMetricSourceToOriginServiceKnownValues tests specific known forward mappings
// to ensure correctness.
func TestMetricSourceToOriginServiceKnownValues(t *testing.T) {
	tests := []struct {
		name                  string
		metricSource          metrics.MetricSource
		expectedOriginService int32
	}{
		{
			name:                  "Unknown",
			metricSource:          metrics.MetricSourceUnknown,
			expectedOriginService: 0,
		},
		{
			name:                  "Dogstatsd",
			metricSource:          metrics.MetricSourceDogstatsd,
			expectedOriginService: 0,
		},
		{
			name:                  "ActivemqXML",
			metricSource:          metrics.MetricSourceActivemqXML,
			expectedOriginService: 11,
		},
		{
			name:                  "GPU",
			metricSource:          metrics.MetricSourceGPU,
			expectedOriginService: 466,
		},
		{
			name:                  "CloudFoundry",
			metricSource:          metrics.MetricSourceCloudFoundry,
			expectedOriginService: 440,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := metricSourceToOriginService(tt.metricSource)
			assert.Equal(t, tt.expectedOriginService, result,
				"metricSourceToOriginService(%v) should return %d, got %d",
				tt.metricSource, tt.expectedOriginService, result)
		})
	}
}

// TestAmbiguousOriginServiceMappings tests cases where multiple MetricSources
// map to the same OriginService value. This documents the expected canonical
// behavior of the reverse mapping.
func TestAmbiguousOriginServiceMappings(t *testing.T) {
	tests := []struct {
		name               string
		originService      int32
		canonicalMS        metrics.MetricSource
		alternativeSources []metrics.MetricSource
	}{
		{
			name:          "Serverless Custom (472)",
			originService: 472,
			canonicalMS:   metrics.MetricSourceAzureContainerAppCustom,
			alternativeSources: []metrics.MetricSource{
				metrics.MetricSourceAzureAppServiceCustom,
				metrics.MetricSourceGoogleCloudRunCustom,
			},
		},
		{
			name:          "Serverless Enhanced (473)",
			originService: 473,
			canonicalMS:   metrics.MetricSourceAzureContainerAppEnhanced,
			alternativeSources: []metrics.MetricSource{
				metrics.MetricSourceAzureAppServiceEnhanced,
				metrics.MetricSourceGoogleCloudRunEnhanced,
			},
		},
		{
			name:          "Serverless Runtime (474)",
			originService: 474,
			canonicalMS:   metrics.MetricSourceAzureContainerAppRuntime,
			alternativeSources: []metrics.MetricSource{
				metrics.MetricSourceAzureAppServiceRuntime,
				metrics.MetricSourceGoogleCloudRunRuntime,
			},
		},
		{
			name:          "Unknown/Dogstatsd (0)",
			originService: 0,
			canonicalMS:   metrics.MetricSourceUnknown,
			alternativeSources: []metrics.MetricSource{
				metrics.MetricSourceDogstatsd,
				metrics.MetricSourceOpenTelemetryCollectorUnknown,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify the canonical mapping
			result := originServiceToMetricSource(tt.originService)
			assert.Equal(t, tt.canonicalMS, result,
				"originServiceToMetricSource(%d) should return canonical value %v, got %v",
				tt.originService, tt.canonicalMS, result)

			// Verify all alternative sources map to the same origin service
			for _, altMS := range tt.alternativeSources {
				altOriginService := metricSourceToOriginService(altMS)
				assert.Equal(t, tt.originService, altOriginService,
					"Alternative MetricSource %v should map to OriginService %d, got %d",
					altMS, tt.originService, altOriginService)
			}

			// Verify idempotency: canonical MS should map back to the same origin service
			canonicalOriginService := metricSourceToOriginService(tt.canonicalMS)
			assert.Equal(t, tt.originService, canonicalOriginService,
				"Canonical MetricSource %v should map to OriginService %d, got %d",
				tt.canonicalMS, tt.originService, canonicalOriginService)
		})
	}
}

// TestOriginServiceRoundTripConsistency ensures that for every unique OriginService
// value in the forward mapping, the reverse mapping produces a MetricSource that
// maps back to the same OriginService.
func TestOriginServiceRoundTripConsistency(t *testing.T) {
	// Collect all unique origin service values by scanning all metric sources
	uniqueOriginServices := make(map[int32]metrics.MetricSource)

	for i := uint16(0); i <= 1024; i++ {
		ms := metrics.MetricSource(i)
		originService := metricSourceToOriginService(ms)

		// Store the first MetricSource we see for each OriginService
		if _, exists := uniqueOriginServices[originService]; !exists {
			uniqueOriginServices[originService] = ms
		}
	}

	// For each unique origin service, verify round-trip consistency
	for originService := range uniqueOriginServices {
		ms := originServiceToMetricSource(originService)
		roundTripOriginService := metricSourceToOriginService(ms)

		assert.Equal(t, originService, roundTripOriginService,
			"Round-trip failed: OriginService(%d) -> MetricSource(%v) -> OriginService(%d)",
			originService, ms, roundTripOriginService)
	}

	// Log the number of unique mappings for informational purposes
	t.Logf("Found %d unique OriginService values in the mapping", len(uniqueOriginServices))
}
