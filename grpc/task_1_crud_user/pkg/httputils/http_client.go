/*
 * Copyright (c) 2024. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package httputils

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

func DoHttpRequest(ctx context.Context, method string, target string, headers map[string]any, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, target, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, fmt.Sprint(value))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
