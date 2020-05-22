package model

import (
	"fmt"
	"log"
	"net/http"

	"github.com/selcukusta/simple-image-server/internal/processor"
	"github.com/selcukusta/simple-image-server/internal/util/constant"
)

//HandlerFinalizer is using to create model for finalizing request
type HandlerFinalizer struct {
	ResponseWriter http.ResponseWriter
	Headers        map[string]string
}

//Finalize is using to finalize the request
func (hf HandlerFinalizer) Finalize(params map[string]string, imageAsByte []byte, contentType string) {
	result, errMessage, err := processor.ImageProcess(params, imageAsByte, contentType)
	if err != nil {
		log.Println(fmt.Sprintf(constant.LogErrorFormat, errMessage, err.Error()))
		hf.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		_, err = hf.ResponseWriter.Write([]byte(constant.ErrorMessage))
		if err != nil {
			log.Println(fmt.Sprintf(constant.LogErrorFormat, constant.LogErrorMessage, err.Error()))
		}
		return
	}

	if constant.CacheControlMaxAge != -1 {
		maxAge := constant.CacheControlMaxAge * 24 * 60 * 60
		hf.ResponseWriter.Header().Add("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge))
	}

	if hf.Headers != nil && len(hf.Headers) > 0 {
		for key, value := range hf.Headers {
			hf.ResponseWriter.Header().Add(key, value)
		}
	}

	_, err = hf.ResponseWriter.Write(result)
	if err != nil {
		log.Println(fmt.Sprintf(constant.LogErrorFormat, constant.LogErrorMessage, err.Error()))
	}
}
