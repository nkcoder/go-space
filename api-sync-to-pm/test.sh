#!/bin/bash

echo "Running unit tests..."
go test -v ./...

echo ""
echo "Running tests with coverage..."
go test -cover ./...

echo ""
echo "Running benchmarks..."
go test -bench=. ./...

echo ""
echo "To run integration tests, set environment variables and run:"
echo "RUN_INTEGRATION_TESTS=1 TEST_DOC_API_KEY=xxx TEST_PM_API_KEY=yyy TEST_PM_WORKSPACE_ID=zzz go test -v -run TestMainIntegration"
