#!/bin/bash
systemctl stop zenbot

sudo go build -O /bin/zenbot/zenbot

systemctl start zenbot
