#!/usr/bin/env bash
thrift -r -out ../tf --gen go:package_prefix=pay-srv-client/tf/ PaymentApi.thrift