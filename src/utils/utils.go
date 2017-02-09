package utils

import (

	//"fmt"
	//	_ "github.com/go-sql-driver/mysql"
	//sql "database/sql"
	//	"io/ioutil"
	log "github.com/siddontang/go/log"
	//	"regexp"
	//	"encoding/json"
	//	"time"
	//	"github.com/garyburd/redigo/redis"
	//"github.com/Unknwon/goconfig"
	//	"strconv"
	//"bytes"
	//	"os"
	//	"bufio"
	//	"io"
	//	"strings"
	//m "model"
	//"utils"
)



func CheckErr(err error) {
	if err != nil {
		log.Error("Error...", err)
	}
}
