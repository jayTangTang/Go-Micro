#!/bin/bash

sudo redis-server /etc/redis/redis.conf
sudo fdfs_trackerd /etc/fdfs/tracker.conf
sudo fdfs_storaged /etc/fdfs/storage.conf
sudo /usr/local/nginx/sbin/nginx

consul agent -dev
