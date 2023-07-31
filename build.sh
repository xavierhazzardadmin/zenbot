#!/bin/bash
systemctl stop zenbot

sudo go build -o /bin/zenbot/zenbot -buildvcs=false

systemctl start zenbot
