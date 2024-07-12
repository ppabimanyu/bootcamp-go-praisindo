/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package text

import (
	"strconv"
	"strings"
)

func Abbreviate(sentence string) string {
	words := strings.Split(sentence, " ")
	var result string
	for _, word := range words {
		if len(word) != 0 {
			result += string(word[0])
		}
	}
	return result
}

func ParseInt64(idString string) (int64, error) {
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
