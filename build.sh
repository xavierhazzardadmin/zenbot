#!/bin/bash
systemctl stop zenbot
echo "Stopped zenbot"

sudo go build -buildvcs=false -o /bin/zenbot/zenbot 

echo "Built executable"

systemctl start zenbot

echo "Started zenbot"
