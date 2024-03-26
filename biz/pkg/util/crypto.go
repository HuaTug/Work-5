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

import (
	"golang.org/x/crypto/bcrypt"
)

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	cost := 5

	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func removeQuotes(s string) string {
	// 检查字符串长度是否大于等于2，以确保字符串包含了至少一个引号
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		// 使用切片操作获取字符串中除了第一个和最后一个字符之外的所有字符
		return s[1 : len(s)-1]
	}
	// 如果字符串不包含引号或者引号不匹配，则原样返回
	return s
}