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

@test "acls: Delete one ACL by its index" {
    resource_delete "$_ACL_BASE_PATH/1" "parent_name=fe_acl&parent_type=frontend"
    assert_equal "$SC" 202
}

@test "acls: Delete one ACL by its index and force reload" {
  resource_delete "$_ACL_BASE_PATH/0" "parent_name=fe_acl&parent_type=frontend&force_reload=true"
    assert_equal "$SC" 204
}
