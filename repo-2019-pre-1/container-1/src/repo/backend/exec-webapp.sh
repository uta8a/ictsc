#!/bin/sh
timeout 30 sh -c "until nc db 3306 ;do sleep 1; done" && /usr/local/bin/server
