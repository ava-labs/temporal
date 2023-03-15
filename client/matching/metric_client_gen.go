// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by cmd/tools/rpcwrappers. DO NOT EDIT.

package matching

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"google.golang.org/grpc"

	"go.temporal.io/server/common/metrics"
)

func (c *metricClient) CancelOutstandingPoll(
	ctx context.Context,
	request *matchingservice.CancelOutstandingPollRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.CancelOutstandingPollResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientCancelOutstandingPollScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.CancelOutstandingPoll(ctx, request, opts...)
}

func (c *metricClient) DescribeTaskQueue(
	ctx context.Context,
	request *matchingservice.DescribeTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.DescribeTaskQueueResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientDescribeTaskQueueScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.DescribeTaskQueue(ctx, request, opts...)
}

func (c *metricClient) GetTaskQueueMetadata(
	ctx context.Context,
	request *matchingservice.GetTaskQueueMetadataRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.GetTaskQueueMetadataResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientGetTaskQueueMetadataScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.GetTaskQueueMetadata(ctx, request, opts...)
}

func (c *metricClient) GetWorkerBuildIdCompatability(
	ctx context.Context,
	request *matchingservice.GetWorkerBuildIdCompatabilityRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.GetWorkerBuildIdCompatabilityResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientGetWorkerBuildIdCompatabilityScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.GetWorkerBuildIdCompatability(ctx, request, opts...)
}

func (c *metricClient) InvalidateTaskQueueMetadata(
	ctx context.Context,
	request *matchingservice.InvalidateTaskQueueMetadataRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.InvalidateTaskQueueMetadataResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientInvalidateTaskQueueMetadataScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.InvalidateTaskQueueMetadata(ctx, request, opts...)
}

func (c *metricClient) ListTaskQueuePartitions(
	ctx context.Context,
	request *matchingservice.ListTaskQueuePartitionsRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.ListTaskQueuePartitionsResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientListTaskQueuePartitionsScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.ListTaskQueuePartitions(ctx, request, opts...)
}

func (c *metricClient) RespondQueryTaskCompleted(
	ctx context.Context,
	request *matchingservice.RespondQueryTaskCompletedRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.RespondQueryTaskCompletedResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientRespondQueryTaskCompletedScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.RespondQueryTaskCompleted(ctx, request, opts...)
}

func (c *metricClient) UpdateWorkerBuildIdCompatability(
	ctx context.Context,
	request *matchingservice.UpdateWorkerBuildIdCompatabilityRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.UpdateWorkerBuildIdCompatabilityResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, metrics.MatchingClientUpdateWorkerBuildIdCompatabilityScope)
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.UpdateWorkerBuildIdCompatability(ctx, request, opts...)
}
