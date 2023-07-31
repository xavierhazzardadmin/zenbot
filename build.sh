#!/bin/bash
systemctl stop zenbot

sudo go build -o /bin/zenbot/zenbot

systemctl start zenbot
