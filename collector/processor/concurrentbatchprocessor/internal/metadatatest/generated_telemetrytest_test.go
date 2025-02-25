// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	// Had to add these two lines manually for some reason - mdatagen didn't import them properly
	"go.opentelemetry.io/otel/metric"
	"github.com/open-telemetry/otel-arrow/collector/processor/concurrentbatchprocessor/internal/metadata"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := SetupTelemetry()
	tb, err := metadata.NewTelemetryBuilder(testTel.NewTelemetrySettings())
	require.NoError(t, err)
	defer tb.Shutdown()
	require.NoError(t, tb.RegisterProcessorBatchMetadataCardinalityCallback(func(_ context.Context, observer metric.Int64Observer) error {
		observer.Observe(1)
		return nil
	}))
	tb.ProcessorBatchBatchSendSize.Record(context.Background(), 1)
	tb.ProcessorBatchBatchSendSizeBytes.Record(context.Background(), 1)
	tb.ProcessorBatchBatchSizeTriggerSend.Add(context.Background(), 1)
	tb.ProcessorBatchTimeoutTriggerSend.Add(context.Background(), 1)

	testTel.AssertMetrics(t, []metricdata.Metrics{
		{
			Name:        "otelcol_processor_batch_batch_send_size",
			Description: "Number of units in the batch",
			Unit:        "{units}",
			Data: metricdata.Histogram[int64]{
				Temporality: metricdata.CumulativeTemporality,
				DataPoints: []metricdata.HistogramDataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_processor_batch_batch_send_size_bytes",
			Description: "Number of bytes in batch that was sent",
			Unit:        "By",
			Data: metricdata.Histogram[int64]{
				Temporality: metricdata.CumulativeTemporality,
				DataPoints: []metricdata.HistogramDataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_processor_batch_batch_size_trigger_send",
			Description: "Number of times the batch was sent due to a size trigger",
			Unit:        "{times}",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_processor_batch_metadata_cardinality",
			Description: "Number of distinct metadata value combinations being processed",
			Unit:        "{combinations}",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: false,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
		{
			Name:        "otelcol_processor_batch_timeout_trigger_send",
			Description: "Number of times the batch was sent due to a timeout trigger",
			Unit:        "{times}",
			Data: metricdata.Sum[int64]{
				Temporality: metricdata.CumulativeTemporality,
				IsMonotonic: true,
				DataPoints: []metricdata.DataPoint[int64]{
					{},
				},
			},
		},
	}, metricdatatest.IgnoreTimestamp(), metricdatatest.IgnoreValue())
	AssertEqualProcessorBatchBatchSendSize(t, testTel.Telemetry,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorBatchBatchSendSizeBytes(t, testTel.Telemetry,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorBatchBatchSizeTriggerSend(t, testTel.Telemetry,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorBatchMetadataCardinality(t, testTel.Telemetry,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorBatchTimeoutTriggerSend(t, testTel.Telemetry,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
