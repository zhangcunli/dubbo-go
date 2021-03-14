/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package match_judger

import (
	"github.com/apache/dubbo-go/config"
)

type DoubleMatchJudger struct {
	config.DoubleMatch
}

func (dmj *DoubleMatchJudger) Judge(input float64) bool {
	if dmj.Exact != 0 {
		return input == dmj.Exact
	}
	if dmj.Range != nil {
		return newDoubleRangeMatchJudger(dmj.Range).Judge(input)
	}
	if dmj.Mode != 0 {
		// todo  mod  match ??
	}
	return true
}

func newDoubleMatchJudger(matchConf *config.DoubleMatch) *DoubleMatchJudger {
	return &DoubleMatchJudger{
		DoubleMatch: *matchConf,
	}
}
