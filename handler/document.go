package handler

import (
	"context"
	"encoding/json"
	"hkg-msa-collector/model"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/micro/go-micro/v2/logger"

	proto "github.com/xtech-cloud/hkg-msp-collector/proto/collector"
)

type Document struct{}

func (this *Document) Scrape(_ctx context.Context, _req *proto.DocumentScrapeRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Document.Scrape, req is %v", _req)

	_rsp.Status = &proto.Status{}

	if "" == _req.Name {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "name is required"
		return nil
	}

	if "" == _req.Address {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "address is required"
		return nil
	}

	u, err := url.Parse(_req.Address)
	if nil != err {
		return err
	}

	uuid := _req.Name + u.Hostname()
	for _, kw := range _req.Keyword {
		uuid += kw
	}
	document := &model.Document{
		ID:      model.ToUUID(uuid),
		Name:    _req.Name,
		Keyword: _req.Keyword,
		Address: _req.Address,
	}

	c := colly.NewCollector(func(c *colly.Collector) {
		extensions.RandomUserAgent(c)
	},
	)

	var daoErr error
	c.OnHTML(_req.Attribute, func(e *colly.HTMLElement) {
		text, err := e.DOM.Html()
		if nil != err {
			logger.Error(err)
			return
		}
		document.CrawledAt = time.Now()
		document.RawText = text
		dao := model.NewDocumentDAO(nil)
		daoErr = dao.UpsertOne(document)
	})

	c.OnError(func(r *colly.Response, e error) {
		logger.Error(e)
	})

	c.Visit(_req.Address)
	return daoErr
}

func (this *Document) Tidy(_ctx context.Context, _req *proto.DocumentTidyRequest, _rsp *proto.BlankResponse) error {
	logger.Infof("Received Document.Tidy, req is %v", _req)

	_rsp.Status = &proto.Status{}
	if "" == _req.Uuid {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "uuid is required"
		return nil
	}

	if 0 == len(_req.Rule) {
		_rsp.Status.Code = 1
		_rsp.Status.Message = "rule is required"
		return nil
	}
	dao := model.NewDocumentDAO(nil)
	document, err := dao.FindOne(_req.Uuid)
	if nil != err {
		return err
	}

	reader := strings.NewReader(document.RawText)
	doc, err := goquery.NewDocumentFromReader(reader)
	if nil != err {
		return err
	}

	jsonDoc := make(map[string]interface{})
	for k, v := range _req.Rule {
		regType := regexp.MustCompile(`\$t\=(.*?);`)
		regElement := regexp.MustCompile(`\$e\=(.*?);`)
		rType := regType.FindStringSubmatch(k)
		rElement := regElement.FindStringSubmatch(k)
		if len(rType) < 2 {
			continue
		}
		if len(rElement) < 2 {
			continue
		}
		if "text" == rType[1] {
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				text := strings.TrimSpace(s.Text())
				reg := regexp.MustCompile("\\s+")
				text = reg.ReplaceAllString(text, "")
				jsonDoc[v] = text
			})
		} else if "ary" == rType[1] {
			jsonValue := make([]string, 0)
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				text := strings.TrimSpace(s.Text())
				reg := regexp.MustCompile("\\s+")
				text = reg.ReplaceAllString(text, "")
				jsonValue = append(jsonValue, text)
				jsonDoc[v] = jsonValue
			})
		} else if "map" == rType[1] {
			jsonValue := make(map[string]string)
			regKeyClass := regexp.MustCompile(`\$pk\=\w*\[class\=\"(.*?)\"`)
			regValueClass := regexp.MustCompile(`\$pv\=\w*\[class\=\"(.*?)\"`)
			rKeyClass := regKeyClass.FindStringSubmatch(k)
			rValueClass := regValueClass.FindStringSubmatch(k)
			keyClass := ""
			valueClass := ""
			if len(rKeyClass) >= 2 {
				keyClass = rKeyClass[1]
			}
			if len(rValueClass) >= 2 {
				valueClass = rValueClass[1]
			}
			var siblingKey *goquery.Selection
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				if s.HasClass(keyClass) {
					siblingKey = s
				} else if s.HasClass(valueClass) {
					if nil != siblingKey {
						reg := regexp.MustCompile("\\s+")
						key := strings.ReplaceAll(siblingKey.Text(), "\u00A0", "\u0020")
						key = strings.TrimSpace(key)
						key = reg.ReplaceAllString(key, "")
						value := strings.ReplaceAll(s.Text(), "\u00A0", "\u0020")
						value = strings.TrimSpace(value)
						value = reg.ReplaceAllString(value, "")
						jsonValue[key] = value
					}
				}
			})
			jsonDoc[v] = jsonValue
		}
	}

	jsonStr, err := json.Marshal(jsonDoc)
	if nil != err {
		return err
	}
	document.TidyText = string(jsonStr)
	err = dao.UpdateOne(document)

	return err
}

func (this *Document) List(_ctx context.Context, _req *proto.ListRequest, _rsp *proto.DocumentListResponse) error {
	logger.Infof("Received Document.List, req is %v", _req)

	_rsp.Status = &proto.Status{}
	offset := int64(0)
	if _req.Offset > 0 {
		offset = _req.Offset
	}

	count := int64(50)
	if _req.Count > 0 {
		count = _req.Count
	}

	dao := model.NewDocumentDAO(nil)
	total, err := dao.Count()
	if nil != err {
		return err
	}
	_rsp.Total = total

	ary, err := dao.List(offset, count)
	if nil != err {
		return err
	}

	_rsp.Entity = make([]*proto.DocumentEntity, len(ary))
	for i, v := range ary {
		_rsp.Entity[i] = &proto.DocumentEntity{
			Uuid:      v.ID,
			Name:      v.Name,
			Address:   v.Address,
			RawText:   v.RawText,
			TidyText:  v.TidyText,
			Keyword:   v.Keyword,
			CrawledAt: v.CrawledAt.Unix(),
		}
	}
	return nil
}
