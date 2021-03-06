/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
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
 *
 * For any questions, please contact developer@tenta.io
 *
 * ydl-compat.go: Re-implementation of various functions from
 *                https://github.com/rg3/youtube-dl/blob/master/youtube_dl/compat.py
 */

package runtime

import "net/url"

// ParseUnquote implements compat.py/compat_urllib_parse_unquote
func ParseUnquote(str string) string {
	ret, err := url.PathUnescape(str)
	if err != nil {
		panic(newExtractorError(err.Error()))
	}
	return ret
}
