#!/usr/bin/env bats
#
# Copyright 2021 HAProxy Technologies
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
load "../../libs/get_json_path"
load '../../libs/version'
load '../../libs/haproxy_config_setup'
load '../../libs/resource_client'

load 'utils/_helpers'

@test "acls: Replace one ACL" {
    resource_put "$_ACL_BASE_PATH/2" "data/put.json" "parent_name=fe_acl&parent_type=frontend"
    assert_equal "$SC" 202
    #
    # verify that ACL is actually changed
    #
    resource_get "$_ACL_BASE_PATH/2" "parent_name=fe_acl&parent_type=frontend"
    assert_equal "$SC" 200

    assert_equal "$(get_json_path "$BODY" " .data" )" "$(cat "$BATS_TEST_DIRNAME/data/put.json")"
}

@test "acls: Return an error when trying to replace non existing ACL" {
    resource_put "$_ACL_BASE_PATH/200" "data/put.json" "parent_name=fe_acl&parent_type=frontend"
    assert_equal "$SC" 404
}
