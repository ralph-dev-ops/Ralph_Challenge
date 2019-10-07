#!/usr/bin/env bash

echo "Creating AWS elastic beanstalk stack"
cd infrastructure
eb create ralph-secnet --elb-type application --cname ralph-secnet
