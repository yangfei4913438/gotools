package network

import (
	"fmt"
	"gotools/base"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//获取下载文件的存放路径
func GetDownloadFileAddr(download_path, download_url string) string {
	x := base.StrContains(download_path, "/")
	if x {
		url_slice := strings.Split(download_url, "/")
		return download_path + url_slice[len(url_slice)-1]
	} else {
		url_slice := strings.Split(download_url, "/")
		return download_path + "/" + url_slice[len(url_slice)-1]
	}
}

//总耗时计算,传入的值单位为秒
func GetElapsedTimeString(all_time int) string {
	var elapsed_time string
	if all_time < 60 {
		elapsed_time = base.IntToStr(all_time) + "秒"
	} else if all_time >= 60 && all_time < 60*60 {
		elapsed_second := base.IntToStr(all_time%60) + "秒"
		elapsed_time = base.IntToStr(all_time/60) + "分" + elapsed_second
	} else if all_time >= 3600 && all_time < 3600*24 {
		elapsed_h := base.IntToStr(all_time/3600) + "时"
		m := all_time % 3600
		elapsed_m := base.IntToStr(m/60) + "分"
		elapsed_s := base.IntToStr(m%60) + "秒"
		elapsed_time = elapsed_h + elapsed_m + elapsed_s
	} else {
		elapsed_d := base.IntToStr(all_time/(3600*24)) + "天"
		h := all_time % (3600 * 24)
		elapsed_h := base.IntToStr(h/3600) + "时"
		m := h % 3600
		elapsed_m := base.IntToStr(m/60) + "分"
		elapsed_s := base.IntToStr(m%60) + "秒"
		elapsed_time = elapsed_d + elapsed_h + elapsed_m + elapsed_s
	}

	return elapsed_time
}

//去程专用下载返回,传入下载地址,本地绝对路径
func GoDownload(wg *sync.WaitGroup, download_path, download_url string) {
	res, err := http.Get(download_url)
	if err != nil {
		fmt.Println("连接远程服务器时出错!" + err.Error())
		os.Exit(1)
	}

	//获取下载文件的存放路径
	target_addr := GetDownloadFileAddr(download_path, download_url)

	target_file, err := os.Create(target_addr)
	defer target_file.Close()
	if err != nil {
		fmt.Println("创建本地文件出错:" + err.Error())
		os.Exit(1)
	}
	io.Copy(target_file, res.Body)

	wg.Done()
}

//去程专用,检查下载信息。传入下载地址,本地保存路径,开始下载时间
func GoInfo(wg *sync.WaitGroup, download_path, download_url string, start_time time.Time) {

	res, err := http.Get(download_url)
	if err != nil {
		fmt.Println("连接远程服务器时出错!" + err.Error())
		os.Exit(1)
	}

	size, size_int := base.GetSize(res.ContentLength)
	fmt.Println("远程文件的大小为: " + size)

	//获取下载文件的存放路径
	target_addr := GetDownloadFileAddr(download_path, download_url)

	var local_size_last int64
	for {
		local_size, local_size_int := base.GetFileSize(target_addr)
		sd := local_size_int - local_size_last
		sd_str, _ := base.GetSize(sd)

		ps := int(float32(local_size_int) / float32(size_int) * 100)

		var jdt, showstr, avg_str, download_percent string
		if ps < 100 {
			jdt = strings.Repeat("=", ps/2) + ">" + strings.Repeat(" ", 50-ps/2)
			download_percent = strconv.Itoa(ps) + "%"
			showstr = "\r已下载: %s/%s [%s] 已完成:%s 当前下载速度:%s 已耗时:%s         "
			//因为还没有下载完成，所以下载速度就是上一秒和下一秒的差
			avg_str = sd_str
		} else {
			jdt = strings.Repeat("=", 50)
			download_percent = "100%"
			showstr = "\r已下载: %s/%s [%s] 已完成:%s 平均下载速度:%s 总耗时:%s          \n"
		}

		//计算时长
		all_time := int(math.Floor(time.Since(start_time).Seconds()))
		//耗时计算
		elapsed_time := GetElapsedTimeString(all_time)

		//平均速度计算
		if size_int < 1024 {
			avg_str = base.IntToStr(int(float32(size_int)/float32(all_time))) + "B/s"
		} else if 1024 <= size_int && size_int < 1024*1024 {
			avg_str = base.Float64ToStr(float64(size_int)/1024/float64(all_time)) + "K/s"
		} else {
			avg_str = base.Float64ToStr(float64(size_int)/1024/1024/float64(all_time)) + "M/s"
		}

		//打印下载进度条
		fmt.Printf(showstr, local_size, size, jdt, download_percent, avg_str, elapsed_time)

		if ps >= 100 {
			break
		} else {
			//每次循环的间隔位1秒,这个不能改
			time.Sleep(time.Second)
			//记录本次循环的已下载文件大小
			local_size_last = local_size_int
		}
	}
	wg.Done()
}

//功能整合,将下载代码抽象到一个函数中
func UrlDownload(download_path, download_url string) (string, error) {
	//清理下载目录
	_, err := base.ClearDir(download_path)
	if err != nil {
		return "清理下载目录出现异常:" + err.Error(), err
	}

	//记录下载开始时间
	start_download_time := time.Now()
	//开启等待组
	wg := sync.WaitGroup{}
	//设置为2个并发,一个下载,一个状态检查
	wg.Add(2)
	go GoDownload(&wg, download_path, download_url)
	go GoInfo(&wg, download_path, download_url, start_download_time)
	//程序等待所有去程都结束
	wg.Wait()

	//全部执行完成,返回0
	return "下载完成!", nil
}
