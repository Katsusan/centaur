package models

//PaginationParam：分页查询时的分页参数
type PaginationParam struct {
	PageIndex int //页索引，即第几页
	PageSize  int //页大小，即每页结果数
}

//PaginationResult: 分页查询结果
type PaginationResult struct {
	Total int //结果总数
}
