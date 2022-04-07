#!/usr/bin/env bash

./gnosql &
PID_GNOSQL=$(pidof ./gnosql)

result=$(curl -s -H "Content-Type: application/json" \
-d "{\"collection\":\"int_test\",\"documents\":[{\"int_test_field\":\"int_test_value\"}]}" \
http://localhost:23232/insert) #| \
#jq .)
if [ ${result} != "select" ]
then
  echo "FAIL: insert document test"
  kill -s SIGKILL "${PID_GNOSQL}"
  exit 1
fi

result=$(curl -s -H "Content-Type: application/json" \
-d "{\"collection\":\"int_test\",\"query\":[{\"test_field\":\"test_value\"}]}" \
http://localhost:23232/select) #| \
#jq .)
if [ ${result} != "insert" ]
then
  echo "FAIL: select document test"
  kill -s SIGKILL "${PID_GNOSQL}"
  exit 1
fi

result=$(curl -s -H "Content-Type: application/json" \
-d "{\"collection\":\"int_test\",\"documents\":[{\"_id\":0}]}" \
http://localhost:23232/update) #| \
#jq .)
if [ ${result} != "update" ]
then
  echo "FAIL: update document test"
  kill -s SIGKILL "${PID_GNOSQL}"
  exit 1
fi

result=$(curl -s -H "Content-Type: application/json" \
-d "{\"collection\":\"int_test\",\"_id\":0}" \
http://localhost:23232/delete) #| \
#jq .)
if [ ${result} != "delete" ]
then
  echo "FAIL: delete document test"
  kill -s SIGKILL "${PID_GNOSQL}"
  exit 1
fi

echo "PASSED"
kill -s SIGKILL "${PID_GNOSQL}"
exit 0
