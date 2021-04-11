package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	DocumentCollectionName = "hkg_collector_document"
)

type Document struct {
	ID        string    `bson:"_id"`
	Name      string    `bason:"name"`
	Keyword   []string    `bason:"keyword"`
	Address   string    `bason:"address"`
	RawText   string    `bason:"rawText"`
	CrawledAt time.Time `bason:"crawledAt"`
}

type DocumentDAO struct {
	conn *Conn
}

func NewDocumentDAO(_conn *Conn) *DocumentDAO {
	if nil == _conn {
		return &DocumentDAO{
			conn: defaultConn,
		}
	} else {
		return &DocumentDAO{
			conn: _conn,
		}
	}
}

func (this *DocumentDAO) InsertOne(_doc *Document) (_err error) {
	_err = nil

	ctx, cancel := NewContext()
	defer cancel()

	document, err := bson.Marshal(_doc)
	if nil != err {
		_err = err
		return
	}

	_, err = this.conn.DB.Collection(DocumentCollectionName).InsertOne(ctx, document)
	if nil != err {
		// 忽略键重复的错误
		if mongo.IsDuplicateKeyError(err) {
			err = nil
		}
	}
	_err = err
	return
}

func (this *DocumentDAO) Count() (int64, error) {
	ctx, cancel := NewContext()
	defer cancel()
	count, err := this.conn.DB.Collection(DocumentCollectionName).EstimatedDocumentCount(ctx)
	return count, err
}

func (this *DocumentDAO) List(_offset int64, _count int64) ([]*Document, error) {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.D{}
	// 1: 倒叙  -1：正序
	sort := bson.D{{"crawledAt", -1}}

	findOptions := options.Find()
	findOptions.SetSort(sort)
	findOptions.SetSkip(_offset)
	findOptions.SetLimit(_count)

	cur, err := this.conn.DB.Collection(DocumentCollectionName).Find(ctx, filter, findOptions)
	if nil != err {
		return make([]*Document, 0), err
	}
	defer cur.Close(ctx)

	var ary []*Document
	for cur.Next(ctx) {
		var document Document
		err = cur.Decode(&document)
		if nil != err {
			return make([]*Document, 0), err
		}
		ary = append(ary, &document)
	}
	return ary, nil
}

func (this *DocumentDAO) UpdateOne(_doc *Document) error {
	ctx, cancel := NewContext()
	defer cancel()

	filter := bson.D{{"name", _doc.Name}}
	update := bson.D{
		{"$set", bson.D{
			{"rawText", _doc.RawText},
			{"crawledAt", _doc.CrawledAt},
		}},
	}
	_, err := this.conn.DB.Collection(DocumentCollectionName).UpdateOne(ctx, filter, update)
	if nil != err {
		return err
	}
	return nil
}
