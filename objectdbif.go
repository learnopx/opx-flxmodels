package models

import (
	"database/sql"
	"fmt"
	//"strconv"
	"utils/dbutils"
	"strings"
)

func (obj IPV4Route) CreateDBTable(dbHdl *sql.DB) error {
	dbCmd := "CREATE TABLE IF NOT EXISTS IPV4Routes " +
		"( DestinationNw varchar(255) PRIMARY KEY ," +
		"NetworkMask varchar(255) ," +
		"Cost integer ," +
		"NextHopIp varchar(255) ," +
		"OutgoingInterface varchar(255) ," +
		"Protocol varchar(255) )"

	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj IPV4Route) StoreObjectInDb(dbHdl *sql.DB) (int64, error) {
	var objectId int64
	dbCmd := fmt.Sprintf(`INSERT INTO IPV4Routes (DestinationNw, NetworkMask, Cost, NextHopIp, OutgoingInterface, Protocol) VALUES ('%v', '%v', %v, '%v', '%v', '%s') ;`,
		obj.DestinationNw, obj.NetworkMask, obj.Cost, obj.NextHopIp, obj.OutgoingInterface, obj.Protocol)
	result, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	if err != nil {
		fmt.Println("### Failed to create IPV4Route ", err)
		return objectId, err
	}
	objectId, err = result.LastInsertId()
	if err != nil {
		fmt.Println("### Failed to get ObjectID after executing ", dbCmd, err)
	}
	return objectId, err
}

func (obj IPV4Route) DeleteObjectFromDb(objKey string, dbHdl *sql.DB) error {
	dbCmd := "delete from " + "IPV4Routes" + " where " + objKey
	_, err := dbutils.ExecuteSQLStmt(dbCmd, dbHdl)
	return err
}

func (obj IPV4Route) GetKey() (string, error) {
	v4RouteKey := obj.DestinationNw + ":" + obj.NetworkMask
	return v4RouteKey, nil
}

func (obj IPV4Route) GetSqlKeyStr(objKey string) (string, error) {
	str := strings.Split(objKey, ":")
	destNw, netMask := str[0], str[1]
	v4RouteSqlKey := "DestinationNw = " + "\"" + destNw + "\"" + " and NetworkMask = " + "\"" + netMask + "\""
	return v4RouteSqlKey, nil
}
