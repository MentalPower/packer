// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type CreateElasticIpSpec struct {

    /* 购买弹性ip数量；取值范围：[1,100]  */
    MaxCount int `json:"maxCount"`

    /* 指定弹性ip地址进行创建，当申请创建多个弹性ip时，必须为空 (Optional) */
    ElasticIpAddress string `json:"elasticIpAddress"`

    /* 弹性ip规格  */
    ElasticIpSpec ElasticIpSpec `json:"elasticIpSpec"`
}
