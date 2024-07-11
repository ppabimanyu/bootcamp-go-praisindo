/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package rand

import gonanoid "github.com/matoous/go-nanoid"

func Generate(rawAlphabet string, size int, prefix ...string) (string, error) {
	r, e := gonanoid.Generate(rawAlphabet, size)
	if e != nil {
		return "", e
	}
	if len(prefix) > 0 {
		return prefix[0] + r, nil
	}
	return r, nil
}

func GenerateNumeric(size int, prefix ...string) (string, error) {
	r, e := gonanoid.Generate("123456789", size)
	if e != nil {
		return "", e
	}
	if len(prefix) > 0 {
		return prefix[0] + r, nil
	}
	return r, nil
}
