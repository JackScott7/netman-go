package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
)

func RouterRequestDo(method string, path string, payload *string, auth string) string {
	const addr = "http://192.168.1.1"
	url := addr + path
	var body io.Reader

	if payload != nil {
		body = strings.NewReader(*payload)
	}

	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		panic(err)
	}

	// Set Headers + Auth
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:141.0) Gecko/20100101 Firefox/141.0")
	req.Header.Add("Referer", addr)
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "Authorization="+auth)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer func(b io.ReadCloser) {
		err := b.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(respBody)
}

func getTplinkDslData(auth string) string {
	payload := strings.Join([]string{
		"[WAN_DSL_INTF_CFG#1,0,0,0,0,0#0,0,0,0,0,0]0,12",
		"status",
		"modulationType",
		"X_TP_AdslModulationCfg",
		"upstreamCurrRate",
		"downstreamCurrRate",
		"X_TP_AnnexType",
		"upstreamMaxRate",
		"downstreamMaxRate",
		"upstreamNoiseMargin",
		"downstreamNoiseMargin",
		"upstreamAttenuation",
		"downstreamAttenuation",
		"[WAN_DSL_INTF_STATS_TOTAL#1,0,0,0,0,0#0,0,0,0,0,0]1,8",
		"ATUCCRCErrors",
		"CRCErrors",
		"ATUCFECErrors",
		"FECErrors",
		"SeverelyErroredSecs",
		"X_TP_US_SeverelyErroredSecs",
		"erroredSecs",
		"X_TP_US_ErroredSecs",
	}, "\r\n") + "\r\n"

	body := RouterRequestDo(http.MethodPost, "/cgi?1&1", &payload, auth)
	return body
}

func getRouterLog(auth string) string {
	body := RouterRequestDo(http.MethodGet, "/cgi/log", nil, auth)
	return body
}

func parseRouterDslResponse(data string) DSLData {
	var dsl DSLData
	for line := range strings.Lines(data) {
		parts := strings.Split(strings.TrimSpace(line), "=")

		if len(parts) < 2 {
			continue
		}

		key, value := parts[0], parts[1]

		switch {
		case key == "status":
			dsl.Status = value
		case key == "modulationType":
			dsl.DSLModulationType = value
		case key == "X_TP_AnnexType":
			dsl.AnnexType = value
		case key == "upstreamCurrRate":
			conv, _ := strconv.ParseUint(value, 10, 16)
			dsl.CurrentRateUpStream = uint16(conv)
		case key == "downstreamCurrRate":
			conv, _ := strconv.ParseUint(value, 10, 16)
			dsl.CurrentRateDownStream = uint16(conv)
		case key == "upstreamMaxRate":
			conv, _ := strconv.ParseUint(value, 10, 32)
			dsl.MaxRateUpStream = uint32(conv)
		case key == "downstreamMaxRate":
			conv, _ := strconv.ParseUint(value, 10, 32)
			dsl.MaxRateDownStream = uint32(conv)
		case key == "upstreamNoiseMargin":
			conv, _ := strconv.ParseFloat(value, 32)
			dsl.SNRMarginUpStream = float32(conv / 10)
		case key == "downstreamNoiseMargin":
			conv, _ := strconv.ParseFloat(value, 32)
			dsl.SNRMarginDownStream = float32(conv / 10)
		case key == "upstreamAttenuation":
			conv, _ := strconv.ParseFloat(value, 32)
			dsl.LineAttenuationUpStream = float32(conv / 10)
		case key == "downstreamAttenuation":
			conv, _ := strconv.ParseFloat(value, 32)
			dsl.LineAttenuationDownStream = float32(conv / 10)
		case key == "ATUCCRCErrors":
			conv, _ := strconv.ParseInt(value, 10, 32)
			dsl.ATUCCRCErrors = int32(conv)
		case key == "CRCErrors":
			conv, _ := strconv.ParseInt(value, 10, 32)
			dsl.CRCErrors = int32(conv)
		}
	}

	return dsl
}

func prettyPrintJSON(data DSLData) error {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return quick.Highlight(
		os.Stdout,
		string(j),
		"json",
		"terminal256",
		"dracula",
	)
}
