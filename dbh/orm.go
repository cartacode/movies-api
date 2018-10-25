package dbh

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// ORM --
type ORM struct {
	collectionName string
	connection     *bongo.Connection
	resultSet      *bongo.ResultSet
	paginationInfo *bongo.PaginationInfo
	perPage        int
	page           int
}

// SetCollection --
func (o *ORM) SetCollection(col string) error {
	connection, err := NewConnection("controllers")
	if err != nil {
		return err
	}
	o.collectionName = col
	o.connection = connection

	o.perPage = 20
	o.page = 0

	return nil

}

// FindAll --
func (o *ORM) FindAll() {
	o.resultSet = o.connection.Collection(o.collectionName).Find(bson.M{})

}

// Pagination --
func (o *ORM) Pagination(perPage, page int) error {
	// Get pagination information
	o.perPage = perPage
	o.page = page
	paginationInfo, err := o.resultSet.Paginate(perPage, page)
	if err != nil {
		return err
	}
	o.paginationInfo = paginationInfo
	return nil
}

// GetPaginationResults =--
func (o *ORM) GetPaginationResults() {
	o.resultSet.Query.Skip(o.page * o.perPage)
	// Add the found results
	// for o.resultSet.Next(series) {
	// retval = append(retval, *series)

}
