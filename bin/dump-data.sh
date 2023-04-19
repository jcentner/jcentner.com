#!/bin/bash

dumpname="db_data_.$(date +%Y%m%d).dump"
touch ${dumpname}
pg_dump ubuntu --data-only --file=${dumpname}
