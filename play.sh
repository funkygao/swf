#!/bin/sh

# config the workflow name
#=========================
if [[ $1 ]]; then
    WORKFLOW=$1
else
    WORKFLOW=w1
fi

# create a fresh database
#========================
mysql -uroot <<EOF
drop database swf;
create database swf;
EOF

mysql -uroot swf < services/history/mysql/db.sql
mysql -uroot swf < services/manager/mysql/db.sql

echo "database ready!"

# register a new workflow type
#=============================
swf workflow -register $WORKFLOW -domain app1 -version v1 -z local -c me

# register activity type
swf activity -register sms -domain app1 -z local -c me
swf activity -register marker -domain app1 -z local -c me

# start a new workflow execution
swf kickoff -z local  -workflow-type $WORKFLOW,v1 -workflow-id order_312 -input '{"order": 312}'
