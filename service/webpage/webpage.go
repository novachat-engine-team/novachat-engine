/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/6/8 20:17
 * @File : webpage.go
 */

package webpage

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"novachat_engine/mtproto"
	"novachat_engine/pkg/log"
	"regexp"
	"time"
)

func GetWebPagePreview(rawurl string) *mtproto.WebPage {
	u, err := url.Parse(rawurl)
	if err != nil {
		return mtproto.NewTLWebPageEmpty(nil).To_WebPage()
	}

	ogContents, err := GetWebpageOgList(u.String(), []string{"image", "site_name", "title", "description", "url"})
	if len(ogContents) == 0 {
		return mtproto.NewTLWebPageEmpty(nil).To_WebPage()
	} else {
		var webPage = mtproto.NewTLWebPage(&mtproto.WebPage{
			Id:         rand.Int63(),
			Url:        rawurl,
			DisplayUrl: u.String()[len(u.Scheme)+3:],
			Type:       "article",
		})

		if v, ok := ogContents["title"]; ok {
			webPage.SetTitle(v)
		}
		if v, ok := ogContents["site_name"]; ok {
			webPage.SetSiteName(v)
		}
		if v, ok := ogContents["description"]; ok {
			webPage.SetDescription(v)
		}

		//var imageBody []byte
		//rawImageUrl, _ := ogContents["image"]
		//if rawImageUrl != "" {
		//	// var  *http.Response
		//	resp, err := http.Get(rawImageUrl)
		//	if err != nil {
		//		log.Warnf("get image body error - %s", err.Error())
		//	} else {
		//		imageBody, err = ioutil.ReadAll(resp.Body)
		//		if err != nil {
		//			log.Warnf("read image body error - %s", err.Error())
		//		}
		//	}
		//} else {
		//	log.Warnf("image empty")
		//}

		//if len(imageBody) > 0 {
		//	// TODO(Coder): gen photo
		//	// 1. upload photo
		//	// 2. getPhoto
		//	// webPage.SetPhoto(mtproto.NewTLPhotoEmpty().To_Photo())
		//}

		// webPage.SetPhoto(mtproto.NewTLPhotoEmpty().To_Photo())

		// image ==> photoSize
		return webPage.To_WebPage()
	}
}

//TODO: (Coder)
func GetWebpageOgList(url string, ogParams []string) (params map[string]string, err error) {
	return
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 3,
	}
	resp, err := c.Get(url)
	if err != nil {
		log.Warnf("get url error - %s", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warnf("get url RealAll error - %s", err.Error())
		return
	}

	params = GetWebpageOgListFromContent(string(body), ogParams)
	return
}

func GetWebpageOgListFromContent(content string, ogParams []string) map[string]string {
	pattern := regexp.MustCompile(`<meta\s+property\s*=\s*"og:([0-9a-zA-Z-]+)"\s+content\s*=\s*"([^"]*?)"\s*/>`)
	allMatches := pattern.FindAllStringSubmatch(content, -1)
	allParams := make(map[string]string)
	for _, val := range allMatches {
		//	log.Infof("og val: %v = %v\n", val[1], val[2])
		k := val[1]
		v := val[2]
		allParams[k] = v
	}

	return allParams
}
