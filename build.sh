#!/bin/bash
systemctl stop zenbot

sudo go build -buildvcs=false -o /bin/zenbot/zenbot 

systemctl start zenbot
