#!/bin/sh
LC_ALL=POSIX
(find ./conf/ -type f -print0  | sort -z | xargs -0 sha1sum; find ./conf/ \( -type f -o -type d \) -print0 | sort -z | xargs -0 stat -c '%n %a') | sha1sum
