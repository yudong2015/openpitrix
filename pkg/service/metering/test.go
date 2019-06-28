package main

import (
	"context"
	"fmt"
	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/models"
	"time"
)

func conn() *db.Conn {
	mysqlCfg := config.MysqlConfig{
		Host: "127.0.0.1",
		Port: "3306",
		User: "root",
		Password: "123456",
		Database: "mbing",
	}

	dbase, err := db.OpenDatabase(mysqlCfg)
	if err != nil {
		fmt.Printf("Failed to open database: [%+v]", err)
	}

	ctx := context.Background()
	conn := dbase.New(ctx)
	return conn
}

func selectAttributes(conn *db.Conn, ids []string) {
	var atts []*models.Attribute

	_, err := conn.Select(constants.IndexedColumns[constants.TableAttribute]...).
		From(constants.TableAttribute).
		Where(db.Eq(constants.ColumnAttributeId, ids)).
		Load(&atts)

	if err != nil {
		fmt.Printf("Failed to load data:\n [%+v]", err)
	}

	fmt.Printf("Loaded data: \n")
	for _, att := range atts {
		fmt.Printf("[%+v]\n", att)
	}
}

func deleteAtts(conn *db.Conn, ids []string) {
	_, err := conn.Update(constants.TableAttribute).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnAttributeId, ids)).
		Exec()

	if err != nil {
		fmt.Printf("Failed to delete data:\n [%+v]", err)
	}
}

func main() {

	//ids := []string{
	//	"att_0001",
	//	"att_0002",
	//}
	//conn := conn()
	//selectAttributes(conn, ids)
	//deleteAtts(conn, ids)

	//var aa []*models.Sku
	//for _, i := range ids {
	//	aa = append(aa, models.NewSku(i, "aaa", []string{}, []string{}))
	//}
	//fmt.Printf("%+v", aa)

	type AAA struct {
		id string
		name string
		now *time.Time
	}

	aaa := AAA{id: "aaa"}
	fmt.Printf("%+v", aaa)

}
