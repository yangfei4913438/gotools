package base

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

//用于读取文件,一次性读取所有的。这个函数用于读取存储提示信息等小段信息的操作。
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

//获取指定目录下的所有目录名称，不进入下一级目录搜索。
func Get_app_name(dirPth string) ([]string, error) {
	dirs := make([]string, 0, 20)
	dir_names, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return dirs, err
	}
	
	for _, dir := range dir_names {
		//只匹配目录
		if dir.IsDir() {
			//过滤文件名是.开头的文件夹
			if Splitstr(dir.Name(), 0, 0) != "." {
				dirs = append(dirs, dir.Name())
			}
		}
	}
	return dirs, nil
}

//按行读取文件。注意：1行文字，返回空字符串切片
func Readline(file_path string) ([]string, int) {
	//声明空字符串
	var res []string
	
	//按行读取文件
	f, err := os.Open(file_path)
	if err != nil {
		res = append(res, err.Error())
		return res, 1
	}
	defer f.Close()
	
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		
		if err != nil || io.EOF == err {
			break
		} else {
			res = append(res, line)
		}
	}
	//返回结果
	return res, 0
}

//获取文件大小的精确值
func GetSize(obj int64) (string, int64) {
	var res string
	if obj < 1024 {
		res = strconv.FormatInt(obj, 10) + "B"
	} else if 1024 <= obj && obj < 1024*1024 {
		res = strconv.FormatFloat(float64(obj)/1024, 'f', 2, 64) + "k"
	} else {
		res = strconv.FormatFloat(float64(obj)/1024/1024, 'f', 2, 64) + "M"
	}
	return res, obj
}

//获取本地文件的大小,返回2个值,一个是带单位的字符串,一个是不带单位的字节
func GetFileSize(file_addr string) (string, int64) {
	
	_, err := os.Open(file_addr)
	if err != nil && os.IsNotExist(err) {
		return "0B", 0
	} else {
		
		file, err := os.Open(file_addr)
		if err != nil {
			fmt.Println("打开文件出错:" + err.Error())
		}
		defer file.Close()
		f, err := file.Stat()
		if err != nil {
			fmt.Println("打开文件出错:" + err.Error())
		}
		r, i := GetSize(f.Size())
		return r, i
	}
}

//检查路径是否存在。文件和目录都可以检查。
func CheckIsExist(addr string) int {
	_, err := os.Open(addr)
	if err != nil && os.IsNotExist(err) {
		//路径不存在
		return 1
	} else {
		//路径存在
		return 0
	}
}

//清空目录, 先删除，再创建
func ClearDir(dir_path string) (string, string) {
	cls_err, cls_out := ShExec("", "rm", "-rf", dir_path)
	if cls_err == "error" {
		return cls_err, cls_out
	} else {
		err, out := ShExec("", "mkdir", "-p", dir_path)
		return err, out
	}
}