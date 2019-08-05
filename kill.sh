#!/bin/bash
ps -ef | grep ./goCloud | grep -v grep | awk '{print $2}' | xargs kill -9