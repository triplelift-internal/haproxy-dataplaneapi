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

@test "cache: Replace one cache" {
  resource_put "$_CACHE_BASE_PATH/test_cache" "data/cache_same_name.json" "force_reload=true"
  assert_equal "$SC" 200

  resource_get "$_CACHE_BASE_PATH/test_cache"
  assert_equal "$SC" 200
  assert_equal "test_cache" "$(get_json_path "$BODY" ".data.name")"
  assert_equal 1 "$(get_json_path "$BODY" ".data.max_object_size")"
  assert_equal 1000 "$(get_json_path "$BODY" ".data.total_max_size")"
}


@test "cache: Fail replacing cache that doesn't exist" {
  resource_put "$_CACHE_BASE_PATH/i_am_not_here" "data/cache_same_name.json" "force_reload=true"
  assert_equal "$SC" 404
}

@test "cache: Fail creating cache that isn't valid" {
  resource_put "$_CACHE_BASE_PATH/test_cache" "data/cache_unvalid.json" "force_reload=true"
  assert_equal "$SC" 400
}
