package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type DataResponse struct {
	Err  int
	Data interface{}
	Msg  string
}

func Reponse(errCode int, data interface{}, msg string) DataResponse {

	if data == nil {
		data = ""
	} else {
		if reflect.ValueOf(data).IsNil() {
			data = ""
		}
		if reflect.TypeOf(data).String() == "[]interface {}" {
			if reflect.ValueOf(data).Len() == 0 {
				data = ""
			}
		}
		if reflect.TypeOf(data).String() == "[]orm.Params" {
			if reflect.ValueOf(data).Len() == 0 {
				data = ""
			}
		}
	}
	resp := DataResponse{
		Err:  errCode,
		Data: data,
		Msg:  msg,
	}
	return resp
}

func GetTodayDate() string {
	return time.Now().Format("2006-01-02")
}

func GetTimeAgo(timeStamp int64) (timeAgo string) {
	currentTime := time.Now().Unix()
	diffTime := currentTime - timeStamp
	log.Println(diffTime)
	if diffTime < 60 {
		return fmt.Sprintf("%d seconds ago.", diffTime)
	}
	if diffTime >= 60 && diffTime < 3600 {

		currentTime := time.Unix(diffTime, 0)
		return fmt.Sprintf("%d minutes ago.", currentTime.Minute())
	}
	if diffTime >= 3600 && diffTime < 3600*24 {
		currentTime := time.Unix(diffTime, 0)
		return fmt.Sprintf("%d hours ago.", currentTime.Hour())
	}
	if diffTime >= 3600*24 {
		currentTime := time.Unix(diffTime, 0)
		return fmt.Sprintf("%d days ago.", currentTime.Day())
	}

	return timeAgo
}

func GetMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetGuid(id int64) string {
	idStr := strconv.Itoa(int(id))
	return GetMd5(idStr + strconv.Itoa(time.Now().Nanosecond()))
}

func AddOne(table, fields string, where map[string]string) (int64, error) {
	o := orm.NewOrm()
	query := "UPDATE " + table + " SET " + fields + " = " + fields + "+1"

	i := 0
	for k, v := range where {
		if i == 0 {
			query = query + " WHERE " + k + " = \"" + v + "\""
		} else {
			query = query + " AND " + k + " = \"" + v + "\""
		}
		i++
	}

	res, err := o.Raw(query).Exec()
	num, err := res.RowsAffected()
	return num, err
}

func MinusOne(table, fields string, where map[string]string) (int64, error) {
	o := orm.NewOrm()
	query := "UPDATE " + table + " SET " + fields + " = " + fields + "-1"

	i := 0
	for k, v := range where {
		if i == 0 {
			query = query + " WHERE " + k + " = \"" + v + "\""
		} else {
			query = query + " AND " + k + " = \"" + v + "\""
		}
		i++
	}

	res, err := o.Raw(query).Exec()
	num, err := res.RowsAffected()
	return num, err
}
