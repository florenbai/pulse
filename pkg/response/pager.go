package response

import (
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

type DataPager struct {
	Page     uint64 `json:"current"`
	PageSize uint64 `json:"pageSize"`
}

func BasePager(c *app.RequestContext) DataPager {
	pageStr := c.Query("current")
	pagesizeStr := c.Query("pageSize")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	pagesize, err := strconv.ParseUint(pagesizeStr, 10, 64)
	if err != nil {
		page = 1
	}
	return DataPager{
		Page:     page,
		PageSize: pagesize,
	}
}
