#!/bin/sh
LC_ALL=POSIX
((find ./temp/ -type f -print0  | sort -z | xargs -0 sha1sum; find ./temp/ \( -type f -o -type d \) -print0 | sort -z | xargs -0 stat -c '%n %a') | sha1sum; sha1sum docker-compose* *.dockerfile) | sha1sum
