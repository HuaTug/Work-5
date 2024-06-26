/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import "time"

// MillTimeStampToTime convert ms timestamp to time.Time
func MillTimeStampToTime(timestamp int64) time.Time {
	second := timestamp / 1000
	nano := timestamp % 1000 * 1000000
	return time.Unix(second, nano)
}

// SecondTimeStampToTime convert s timestamp to time.Time
func SecondTimeStampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
