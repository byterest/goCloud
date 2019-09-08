#!/bin/bash
ps -ef | grep ./prejudice | grep -v grep | awk '{print $2}' | xargs kill -9