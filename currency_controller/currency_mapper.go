package currency_controller

import (
	"gopkg.in/mgo.v2"
	. "currency/currency-provider"
	. "currency/data_store"
)

type CurrencyMapper interface {
	Save(currency Currency)
}

type CurrencyMapperMongdb struct {
	session *mgo.Session
}

func (c *CurrencyMapperMongdb) Save(currency Currency) {
	c.session.DB("currency").C("currency").Insert(currency)
}

func NewCurrencyMapperMongdb(dataStore DataStore) CurrencyMapper {
	return &CurrencyMapperMongdb{session: dataStore.Session.Copy()}
}
