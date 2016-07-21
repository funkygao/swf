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
echo "start zk, kafka, kateway first!"

# remove kafka topics
#====================
gk topics -z local -c me -del app1.d_$WORKFLOW.v1
gk topics -z local -c me -del app1.a_sms.v1
gk topics -z local -c me -del app1.a_marker.v1

# register a new workflow type
#=============================
swf workflow -register $WORKFLOW -domain app1 -version v1 -z local -c me

# register activity type
swf activity -register sms -domain app1 -z local -c me
swf activity -register marker -domain app1 -z local -c me

echo
echo "==============="
echo "alarmer decider"
echo "alarmer sms"
echo "alarmer marker"
echo "==============="
echo

# start a new workflow execution
swf kickoff -z local  -workflow-type $WORKFLOW,v1 -workflow-id order_312 -input '{"order": 312}'
echo "swf kickoff -z local  -workflow-type $WORKFLOW,v1 -workflow-id order_312 -input '{"order": 312}'"
