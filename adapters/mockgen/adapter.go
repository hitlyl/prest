// Code generated by MockGen. DO NOT EDIT.
// Source: adapters/adapter.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	context "context"
	sql "database/sql"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	adapters "github.com/hitlyl/prest/adapters"
)

// MockAdapter is a mock of Adapter interface.
type MockAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterMockRecorder
}

// MockAdapterMockRecorder is the mock recorder for MockAdapter.
type MockAdapterMockRecorder struct {
	mock *MockAdapter
}

// NewMockAdapter creates a new mock instance.
func NewMockAdapter(ctrl *gomock.Controller) *MockAdapter {
	mock := &MockAdapter{ctrl: ctrl}
	mock.recorder = &MockAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdapter) EXPECT() *MockAdapterMockRecorder {
	return m.recorder
}

// BatchInsertCopy mocks base method.
func (m *MockAdapter) BatchInsertCopy(dbname, schema, table string, keys []string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{dbname, schema, table, keys}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchInsertCopy", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// BatchInsertCopy indicates an expected call of BatchInsertCopy.
func (mr *MockAdapterMockRecorder) BatchInsertCopy(dbname, schema, table, keys interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{dbname, schema, table, keys}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsertCopy", reflect.TypeOf((*MockAdapter)(nil).BatchInsertCopy), varargs...)
}

// BatchInsertCopyCtx mocks base method.
func (m *MockAdapter) BatchInsertCopyCtx(ctx context.Context, dbname, schema, table string, keys []string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dbname, schema, table, keys}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchInsertCopyCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// BatchInsertCopyCtx indicates an expected call of BatchInsertCopyCtx.
func (mr *MockAdapterMockRecorder) BatchInsertCopyCtx(ctx, dbname, schema, table, keys interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dbname, schema, table, keys}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsertCopyCtx", reflect.TypeOf((*MockAdapter)(nil).BatchInsertCopyCtx), varargs...)
}

// BatchInsertValues mocks base method.
func (m *MockAdapter) BatchInsertValues(SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchInsertValues", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// BatchInsertValues indicates an expected call of BatchInsertValues.
func (mr *MockAdapterMockRecorder) BatchInsertValues(SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsertValues", reflect.TypeOf((*MockAdapter)(nil).BatchInsertValues), varargs...)
}

// BatchInsertValuesCtx mocks base method.
func (m *MockAdapter) BatchInsertValuesCtx(ctx context.Context, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchInsertValuesCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// BatchInsertValuesCtx indicates an expected call of BatchInsertValuesCtx.
func (mr *MockAdapterMockRecorder) BatchInsertValuesCtx(ctx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInsertValuesCtx", reflect.TypeOf((*MockAdapter)(nil).BatchInsertValuesCtx), varargs...)
}

// CountByRequest mocks base method.
func (m *MockAdapter) CountByRequest(req *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountByRequest", req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountByRequest indicates an expected call of CountByRequest.
func (mr *MockAdapterMockRecorder) CountByRequest(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountByRequest", reflect.TypeOf((*MockAdapter)(nil).CountByRequest), req)
}

// DatabaseClause mocks base method.
func (m *MockAdapter) DatabaseClause(req *http.Request) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseClause", req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// DatabaseClause indicates an expected call of DatabaseClause.
func (mr *MockAdapterMockRecorder) DatabaseClause(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseClause", reflect.TypeOf((*MockAdapter)(nil).DatabaseClause), req)
}

// DatabaseOrderBy mocks base method.
func (m *MockAdapter) DatabaseOrderBy(order string, hasCount bool) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseOrderBy", order, hasCount)
	ret0, _ := ret[0].(string)
	return ret0
}

// DatabaseOrderBy indicates an expected call of DatabaseOrderBy.
func (mr *MockAdapterMockRecorder) DatabaseOrderBy(order, hasCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseOrderBy", reflect.TypeOf((*MockAdapter)(nil).DatabaseOrderBy), order, hasCount)
}

// DatabaseWhere mocks base method.
func (m *MockAdapter) DatabaseWhere(requestWhere string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseWhere", requestWhere)
	ret0, _ := ret[0].(string)
	return ret0
}

// DatabaseWhere indicates an expected call of DatabaseWhere.
func (mr *MockAdapterMockRecorder) DatabaseWhere(requestWhere interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseWhere", reflect.TypeOf((*MockAdapter)(nil).DatabaseWhere), requestWhere)
}

// Delete mocks base method.
func (m *MockAdapter) Delete(SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdapterMockRecorder) Delete(SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdapter)(nil).Delete), varargs...)
}

// DeleteCtx mocks base method.
func (m *MockAdapter) DeleteCtx(ctx context.Context, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// DeleteCtx indicates an expected call of DeleteCtx.
func (mr *MockAdapterMockRecorder) DeleteCtx(ctx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCtx", reflect.TypeOf((*MockAdapter)(nil).DeleteCtx), varargs...)
}

// DeleteSQL mocks base method.
func (m *MockAdapter) DeleteSQL(database, schema, table string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSQL", database, schema, table)
	ret0, _ := ret[0].(string)
	return ret0
}

// DeleteSQL indicates an expected call of DeleteSQL.
func (mr *MockAdapterMockRecorder) DeleteSQL(database, schema, table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSQL", reflect.TypeOf((*MockAdapter)(nil).DeleteSQL), database, schema, table)
}

// DeleteWithTransaction mocks base method.
func (m *MockAdapter) DeleteWithTransaction(tx *sql.Tx, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteWithTransaction", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// DeleteWithTransaction indicates an expected call of DeleteWithTransaction.
func (mr *MockAdapterMockRecorder) DeleteWithTransaction(tx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithTransaction", reflect.TypeOf((*MockAdapter)(nil).DeleteWithTransaction), varargs...)
}

// DistinctClause mocks base method.
func (m *MockAdapter) DistinctClause(r *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DistinctClause", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DistinctClause indicates an expected call of DistinctClause.
func (mr *MockAdapterMockRecorder) DistinctClause(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistinctClause", reflect.TypeOf((*MockAdapter)(nil).DistinctClause), r)
}

// ExecuteScripts mocks base method.
func (m *MockAdapter) ExecuteScripts(method, sql string, values []interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteScripts", method, sql, values)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// ExecuteScripts indicates an expected call of ExecuteScripts.
func (mr *MockAdapterMockRecorder) ExecuteScripts(method, sql, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteScripts", reflect.TypeOf((*MockAdapter)(nil).ExecuteScripts), method, sql, values)
}

// ExecuteScriptsCtx mocks base method.
func (m *MockAdapter) ExecuteScriptsCtx(ctx context.Context, method, sql string, values []interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteScriptsCtx", ctx, method, sql, values)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// ExecuteScriptsCtx indicates an expected call of ExecuteScriptsCtx.
func (mr *MockAdapterMockRecorder) ExecuteScriptsCtx(ctx, method, sql, values interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteScriptsCtx", reflect.TypeOf((*MockAdapter)(nil).ExecuteScriptsCtx), ctx, method, sql, values)
}

// FieldsPermissions mocks base method.
func (m *MockAdapter) FieldsPermissions(r *http.Request, table, op string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FieldsPermissions", r, table, op)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FieldsPermissions indicates an expected call of FieldsPermissions.
func (mr *MockAdapterMockRecorder) FieldsPermissions(r, table, op interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FieldsPermissions", reflect.TypeOf((*MockAdapter)(nil).FieldsPermissions), r, table, op)
}

// GetDatabase mocks base method.
func (m *MockAdapter) GetDatabase() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabase")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabase indicates an expected call of GetDatabase.
func (mr *MockAdapterMockRecorder) GetDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockAdapter)(nil).GetDatabase))
}

// GetScript mocks base method.
func (m *MockAdapter) GetScript(verb, folder, scriptName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScript", verb, folder, scriptName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScript indicates an expected call of GetScript.
func (mr *MockAdapterMockRecorder) GetScript(verb, folder, scriptName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScript", reflect.TypeOf((*MockAdapter)(nil).GetScript), verb, folder, scriptName)
}

// GetTransaction mocks base method.
func (m *MockAdapter) GetTransaction() (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockAdapterMockRecorder) GetTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockAdapter)(nil).GetTransaction))
}

// GetTransactionCtx mocks base method.
func (m *MockAdapter) GetTransactionCtx(ctx context.Context) (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionCtx", ctx)
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionCtx indicates an expected call of GetTransactionCtx.
func (mr *MockAdapterMockRecorder) GetTransactionCtx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionCtx", reflect.TypeOf((*MockAdapter)(nil).GetTransactionCtx), ctx)
}

// GroupByClause mocks base method.
func (m *MockAdapter) GroupByClause(r *http.Request) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GroupByClause", r)
	ret0, _ := ret[0].(string)
	return ret0
}

// GroupByClause indicates an expected call of GroupByClause.
func (mr *MockAdapterMockRecorder) GroupByClause(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GroupByClause", reflect.TypeOf((*MockAdapter)(nil).GroupByClause), r)
}

// Insert mocks base method.
func (m *MockAdapter) Insert(SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAdapterMockRecorder) Insert(SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAdapter)(nil).Insert), varargs...)
}

// InsertCtx mocks base method.
func (m *MockAdapter) InsertCtx(ctx context.Context, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// InsertCtx indicates an expected call of InsertCtx.
func (mr *MockAdapterMockRecorder) InsertCtx(ctx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCtx", reflect.TypeOf((*MockAdapter)(nil).InsertCtx), varargs...)
}

// InsertSQL mocks base method.
func (m *MockAdapter) InsertSQL(database, schema, table, names, placeholders string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertSQL", database, schema, table, names, placeholders)
	ret0, _ := ret[0].(string)
	return ret0
}

// InsertSQL indicates an expected call of InsertSQL.
func (mr *MockAdapterMockRecorder) InsertSQL(database, schema, table, names, placeholders interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSQL", reflect.TypeOf((*MockAdapter)(nil).InsertSQL), database, schema, table, names, placeholders)
}

// InsertWithTransaction mocks base method.
func (m *MockAdapter) InsertWithTransaction(tx *sql.Tx, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertWithTransaction", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// InsertWithTransaction indicates an expected call of InsertWithTransaction.
func (mr *MockAdapterMockRecorder) InsertWithTransaction(tx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithTransaction", reflect.TypeOf((*MockAdapter)(nil).InsertWithTransaction), varargs...)
}

// JoinByRequest mocks base method.
func (m *MockAdapter) JoinByRequest(r *http.Request) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JoinByRequest", r)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// JoinByRequest indicates an expected call of JoinByRequest.
func (mr *MockAdapterMockRecorder) JoinByRequest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JoinByRequest", reflect.TypeOf((*MockAdapter)(nil).JoinByRequest), r)
}

// OrderByRequest mocks base method.
func (m *MockAdapter) OrderByRequest(r *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderByRequest", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderByRequest indicates an expected call of OrderByRequest.
func (mr *MockAdapterMockRecorder) OrderByRequest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderByRequest", reflect.TypeOf((*MockAdapter)(nil).OrderByRequest), r)
}

// PaginateIfPossible mocks base method.
func (m *MockAdapter) PaginateIfPossible(r *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PaginateIfPossible", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PaginateIfPossible indicates an expected call of PaginateIfPossible.
func (mr *MockAdapterMockRecorder) PaginateIfPossible(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PaginateIfPossible", reflect.TypeOf((*MockAdapter)(nil).PaginateIfPossible), r)
}

// ParseBatchInsertRequest mocks base method.
func (m *MockAdapter) ParseBatchInsertRequest(r *http.Request) (string, string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseBatchInsertRequest", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].([]interface{})
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ParseBatchInsertRequest indicates an expected call of ParseBatchInsertRequest.
func (mr *MockAdapterMockRecorder) ParseBatchInsertRequest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseBatchInsertRequest", reflect.TypeOf((*MockAdapter)(nil).ParseBatchInsertRequest), r)
}

// ParseInsertRequest mocks base method.
func (m *MockAdapter) ParseInsertRequest(r *http.Request) (string, string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseInsertRequest", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].([]interface{})
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ParseInsertRequest indicates an expected call of ParseInsertRequest.
func (mr *MockAdapterMockRecorder) ParseInsertRequest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseInsertRequest", reflect.TypeOf((*MockAdapter)(nil).ParseInsertRequest), r)
}

// ParseScript mocks base method.
func (m *MockAdapter) ParseScript(scriptPath string, templateData map[string]interface{}) (string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseScript", scriptPath, templateData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseScript indicates an expected call of ParseScript.
func (mr *MockAdapterMockRecorder) ParseScript(scriptPath, templateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseScript", reflect.TypeOf((*MockAdapter)(nil).ParseScript), scriptPath, templateData)
}

// Query mocks base method.
func (m *MockAdapter) Query(SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockAdapterMockRecorder) Query(SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockAdapter)(nil).Query), varargs...)
}

// QueryCount mocks base method.
func (m *MockAdapter) QueryCount(SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryCount", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// QueryCount indicates an expected call of QueryCount.
func (mr *MockAdapterMockRecorder) QueryCount(SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCount", reflect.TypeOf((*MockAdapter)(nil).QueryCount), varargs...)
}

// QueryCountCtx mocks base method.
func (m *MockAdapter) QueryCountCtx(ctx context.Context, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryCountCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// QueryCountCtx indicates an expected call of QueryCountCtx.
func (mr *MockAdapterMockRecorder) QueryCountCtx(ctx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCountCtx", reflect.TypeOf((*MockAdapter)(nil).QueryCountCtx), varargs...)
}

// QueryCtx mocks base method.
func (m *MockAdapter) QueryCtx(ctx context.Context, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// QueryCtx indicates an expected call of QueryCtx.
func (mr *MockAdapterMockRecorder) QueryCtx(ctx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryCtx", reflect.TypeOf((*MockAdapter)(nil).QueryCtx), varargs...)
}

// ReturningByRequest mocks base method.
func (m *MockAdapter) ReturningByRequest(r *http.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturningByRequest", r)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReturningByRequest indicates an expected call of ReturningByRequest.
func (mr *MockAdapterMockRecorder) ReturningByRequest(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturningByRequest", reflect.TypeOf((*MockAdapter)(nil).ReturningByRequest), r)
}

// SchemaClause mocks base method.
func (m *MockAdapter) SchemaClause(req *http.Request) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaClause", req)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// SchemaClause indicates an expected call of SchemaClause.
func (mr *MockAdapterMockRecorder) SchemaClause(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaClause", reflect.TypeOf((*MockAdapter)(nil).SchemaClause), req)
}

// SchemaOrderBy mocks base method.
func (m *MockAdapter) SchemaOrderBy(order string, hasCount bool) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaOrderBy", order, hasCount)
	ret0, _ := ret[0].(string)
	return ret0
}

// SchemaOrderBy indicates an expected call of SchemaOrderBy.
func (mr *MockAdapterMockRecorder) SchemaOrderBy(order, hasCount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaOrderBy", reflect.TypeOf((*MockAdapter)(nil).SchemaOrderBy), order, hasCount)
}

// SchemaTablesClause mocks base method.
func (m *MockAdapter) SchemaTablesClause() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaTablesClause")
	ret0, _ := ret[0].(string)
	return ret0
}

// SchemaTablesClause indicates an expected call of SchemaTablesClause.
func (mr *MockAdapterMockRecorder) SchemaTablesClause() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaTablesClause", reflect.TypeOf((*MockAdapter)(nil).SchemaTablesClause))
}

// SchemaTablesOrderBy mocks base method.
func (m *MockAdapter) SchemaTablesOrderBy(order string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaTablesOrderBy", order)
	ret0, _ := ret[0].(string)
	return ret0
}

// SchemaTablesOrderBy indicates an expected call of SchemaTablesOrderBy.
func (mr *MockAdapterMockRecorder) SchemaTablesOrderBy(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaTablesOrderBy", reflect.TypeOf((*MockAdapter)(nil).SchemaTablesOrderBy), order)
}

// SchemaTablesWhere mocks base method.
func (m *MockAdapter) SchemaTablesWhere(requestWhere string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchemaTablesWhere", requestWhere)
	ret0, _ := ret[0].(string)
	return ret0
}

// SchemaTablesWhere indicates an expected call of SchemaTablesWhere.
func (mr *MockAdapterMockRecorder) SchemaTablesWhere(requestWhere interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchemaTablesWhere", reflect.TypeOf((*MockAdapter)(nil).SchemaTablesWhere), requestWhere)
}

// SelectFields mocks base method.
func (m *MockAdapter) SelectFields(fields []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectFields", fields)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectFields indicates an expected call of SelectFields.
func (mr *MockAdapterMockRecorder) SelectFields(fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectFields", reflect.TypeOf((*MockAdapter)(nil).SelectFields), fields)
}

// SelectSQL mocks base method.
func (m *MockAdapter) SelectSQL(selectStr, database, schema, table string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectSQL", selectStr, database, schema, table)
	ret0, _ := ret[0].(string)
	return ret0
}

// SelectSQL indicates an expected call of SelectSQL.
func (mr *MockAdapterMockRecorder) SelectSQL(selectStr, database, schema, table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectSQL", reflect.TypeOf((*MockAdapter)(nil).SelectSQL), selectStr, database, schema, table)
}

// SetByRequest mocks base method.
func (m *MockAdapter) SetByRequest(r *http.Request, initialPlaceholderID int) (string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetByRequest", r, initialPlaceholderID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SetByRequest indicates an expected call of SetByRequest.
func (mr *MockAdapterMockRecorder) SetByRequest(r, initialPlaceholderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetByRequest", reflect.TypeOf((*MockAdapter)(nil).SetByRequest), r, initialPlaceholderID)
}

// SetDatabase mocks base method.
func (m *MockAdapter) SetDatabase(name string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDatabase", name)
}

// SetDatabase indicates an expected call of SetDatabase.
func (mr *MockAdapterMockRecorder) SetDatabase(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDatabase", reflect.TypeOf((*MockAdapter)(nil).SetDatabase), name)
}

// ShowTable mocks base method.
func (m *MockAdapter) ShowTable(schema, table string) adapters.Scanner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowTable", schema, table)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// ShowTable indicates an expected call of ShowTable.
func (mr *MockAdapterMockRecorder) ShowTable(schema, table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowTable", reflect.TypeOf((*MockAdapter)(nil).ShowTable), schema, table)
}

// ShowTableCtx mocks base method.
func (m *MockAdapter) ShowTableCtx(ctx context.Context, schema, table string) adapters.Scanner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowTableCtx", ctx, schema, table)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// ShowTableCtx indicates an expected call of ShowTableCtx.
func (mr *MockAdapterMockRecorder) ShowTableCtx(ctx, schema, table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowTableCtx", reflect.TypeOf((*MockAdapter)(nil).ShowTableCtx), ctx, schema, table)
}

// TableClause mocks base method.
func (m *MockAdapter) TableClause() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableClause")
	ret0, _ := ret[0].(string)
	return ret0
}

// TableClause indicates an expected call of TableClause.
func (mr *MockAdapterMockRecorder) TableClause() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableClause", reflect.TypeOf((*MockAdapter)(nil).TableClause))
}

// TableOrderBy mocks base method.
func (m *MockAdapter) TableOrderBy(order string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableOrderBy", order)
	ret0, _ := ret[0].(string)
	return ret0
}

// TableOrderBy indicates an expected call of TableOrderBy.
func (mr *MockAdapterMockRecorder) TableOrderBy(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableOrderBy", reflect.TypeOf((*MockAdapter)(nil).TableOrderBy), order)
}

// TablePermissions mocks base method.
func (m *MockAdapter) TablePermissions(table, op string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TablePermissions", table, op)
	ret0, _ := ret[0].(bool)
	return ret0
}

// TablePermissions indicates an expected call of TablePermissions.
func (mr *MockAdapterMockRecorder) TablePermissions(table, op interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TablePermissions", reflect.TypeOf((*MockAdapter)(nil).TablePermissions), table, op)
}

// TableWhere mocks base method.
func (m *MockAdapter) TableWhere(requestWhere string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableWhere", requestWhere)
	ret0, _ := ret[0].(string)
	return ret0
}

// TableWhere indicates an expected call of TableWhere.
func (mr *MockAdapterMockRecorder) TableWhere(requestWhere interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableWhere", reflect.TypeOf((*MockAdapter)(nil).TableWhere), requestWhere)
}

// Update mocks base method.
func (m *MockAdapter) Update(SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAdapterMockRecorder) Update(SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdapter)(nil).Update), varargs...)
}

// UpdateCtx mocks base method.
func (m *MockAdapter) UpdateCtx(ctx context.Context, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCtx", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// UpdateCtx indicates an expected call of UpdateCtx.
func (mr *MockAdapterMockRecorder) UpdateCtx(ctx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCtx", reflect.TypeOf((*MockAdapter)(nil).UpdateCtx), varargs...)
}

// UpdateSQL mocks base method.
func (m *MockAdapter) UpdateSQL(database, schema, table, setSyntax string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSQL", database, schema, table, setSyntax)
	ret0, _ := ret[0].(string)
	return ret0
}

// UpdateSQL indicates an expected call of UpdateSQL.
func (mr *MockAdapterMockRecorder) UpdateSQL(database, schema, table, setSyntax interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSQL", reflect.TypeOf((*MockAdapter)(nil).UpdateSQL), database, schema, table, setSyntax)
}

// UpdateWithTransaction mocks base method.
func (m *MockAdapter) UpdateWithTransaction(tx *sql.Tx, SQL string, params ...interface{}) adapters.Scanner {
	m.ctrl.T.Helper()
	varargs := []interface{}{tx, SQL}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateWithTransaction", varargs...)
	ret0, _ := ret[0].(adapters.Scanner)
	return ret0
}

// UpdateWithTransaction indicates an expected call of UpdateWithTransaction.
func (mr *MockAdapterMockRecorder) UpdateWithTransaction(tx, SQL interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{tx, SQL}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWithTransaction", reflect.TypeOf((*MockAdapter)(nil).UpdateWithTransaction), varargs...)
}

// WhereByRequest mocks base method.
func (m *MockAdapter) WhereByRequest(r *http.Request, initialPlaceholderID int) (string, []interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhereByRequest", r, initialPlaceholderID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]interface{})
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// WhereByRequest indicates an expected call of WhereByRequest.
func (mr *MockAdapterMockRecorder) WhereByRequest(r, initialPlaceholderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhereByRequest", reflect.TypeOf((*MockAdapter)(nil).WhereByRequest), r, initialPlaceholderID)
}
