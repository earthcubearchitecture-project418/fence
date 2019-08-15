#!/bin/bash
DOCKERVER=$(<../VERSION)

docker save earthcube/p418webui:$DOCKERVER | bzip2 | pv |  ssh -i /home/fils/.ssh/id_rsa root@geodex.org 'bunzip2 | docker load'

