// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SyncSrcTableDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SyncSrcTableDao struct {
	gmvc.M                                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                              // DB is the raw underlying database management object.
	Table   string                              // Table is the table name of the DAO.
	Columns syncSrcTableColumns // Columns contains all the columns of Table that for convenient usage.
}

// SyncSrcTableColumns defines and stores column names for table sync_src_table.
type syncSrcTableColumns struct {
	Id          string //   
    CreateDate  string //   
    WriteDate   string //   
    TaskId      string //   
    DbType      string //   
    DbName      string //   
    TableName   string //   
    PkFiled     string //
}

func NewSyncSrcTableDao() *SyncSrcTableDao {
	return &SyncSrcTableDao{
		M:     g.DB("default").Model("sync_src_table").Safe(),
		DB:    g.DB("default"),
		Table: "sync_src_table",
		Columns: syncSrcTableColumns{
			Id:         "id",           
            CreateDate: "create_date",  
            WriteDate:  "write_date",   
            TaskId:     "task_id",      
            DbType:     "db_type",      
            DbName:     "db_name",      
            TableName:  "table_name",   
            PkFiled:    "pk_filed",
		},
	}
}