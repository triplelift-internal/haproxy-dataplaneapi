#!/usr/bin/env bats
#
# Copyright 2022 HAProxy Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

load '../../libs/dataplaneapi'
load '../../libs/get_json_path'
load '../../libs/haproxy_config_setup'
load '../../libs/resource_client'
load '../../libs/version'

load 'utils/_helpers'

@test "cache: Return an array of caches" {
  resource_get "$_CACHE_BASE_PATH"
  assert_equal "$SC" 200
  assert_equal 2 "$(get_json_path "$BODY" ".data | length")"

  for (( i=0; i<=1; i++ ))
  do
    if [ "$(get_json_path "$BODY" ".data[$i].name")" == "test_cache" ]
    then
      assert_equal 60 "$(get_json_path "$BODY" ".data[$i].max_age")"
      assert_equal 8 "$(get_json_path "$BODY" ".data[$i].max_object_size")"
      assert_equal 1024 "$(get_json_path "$BODY" ".data[$i].total_max_size")"
    elif [ "$(get_json_path "$BODY" ".data[$i].name")" == "test_cache2" ]
	then
      assert_equal 1024 "$(get_json_path "$BODY" ".data[$i].total_max_size")"
    else
      assert_failure
    fi
  done
}
