package handler

import (
	"context"
	"hkg-msa-collector/model"
	"net/url"
	"time"

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
		daoErr = dao.InsertOne(document)
	})

	c.OnError(func(r *colly.Response, e error) {
		logger.Error(e)
	})

	c.Visit(_req.Address)
	return daoErr
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
			Keyword:   v.Keyword,
			CrawledAt: v.CrawledAt.Unix(),
		}
	}
	return nil
}
