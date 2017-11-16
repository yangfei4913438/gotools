package network

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"github.com/yangfei4913438/gotools/base"
	"math"
	"net/http"
	"strconv"
	"strings"
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

// 计算平均速度
func GetAvgStr(file_size int64, all_time int) string {
	var avg_str string

	// 平均速度计算: 当前文件大小/总共花费的时间(秒)=每秒的下载速度
	if file_size < 1024 {
		avg_str = base.IntToStr(int(float32(file_size)/float32(all_time))) + "B/s"
	} else if 1024 <= file_size && file_size < 1024*1024 {
		avg_str = base.Float64ToStr(float64(file_size)/1024/float64(all_time)) + "K/s"
	} else {
		avg_str = base.Float64ToStr(float64(file_size)/1024/1024/float64(all_time)) + "M/s"
	}

	return avg_str
}

// 文件下载
func UrlDownload(download_path, download_url string) error {

	//清理下载目录
	_, err := base.ClearDir(download_path)
	if err != nil {
		fmt.Println("清理下载目录出现异常:" + err.Error())
		return errors.New("清理下载目录出现异常:" + err.Error())
	}

	//记录下载开始时间
	start_time := time.Now()

	// 进度条, 显示的文字，平均速度，下载百分比
	var jdt, showstr, avg_str, download_percent string

	// 获取下载对象的信息
	res, err := http.Get(download_url)
	if err != nil {
		fmt.Println("连接远程服务器时出错!" + err.Error())
		return errors.New("连接远程服务器时出错!" + err.Error())
	}

	remote_size := base.GetSize(res.ContentLength)
	fmt.Println("远程文件的大小为: " + remote_size)

	// 回调函数，下载的时候，实时处理
	progress := func(current, total int64) {

		// 计算时长
		all_time := int(math.Floor(time.Since(start_time).Seconds()))

		// 耗时计算
		elapsed_time := GetElapsedTimeString(all_time)

		// 计算百分比
		ps := int(float32(current) / float32(total) * 100)

		// 绘制进度条
		jdt = strings.Repeat("=", ps/2) + ">" + strings.Repeat(" ", 50-ps/2)

		// 进度百分比
		download_percent = strconv.Itoa(ps) + "%"

		// 文字模板
		showstr = "\r已下载: %s/%s [%s] 已完成:%s 当前下载速度:%s 已耗时:%s         "

		// 平均速度计算
		avg_str = GetAvgStr(current, all_time)

		// 打印下载进度条
		fmt.Printf(showstr, base.GetSize(current), base.GetSize(total), jdt, download_percent, avg_str, elapsed_time)
	}

	// 启动下载
	r, err := req.Get(download_url, req.DownloadProgress(progress))
	if err != nil {
		fmt.Println("启动下载出现错误: " + err.Error())
		return errors.New("启动下载出现错误: " + err.Error())
	}

	//获取下载文件的存放路径
	target_addr := GetDownloadFileAddr("download", download_url)

	//保存文件
	err = r.ToFile(target_addr)
	if err != nil {
		fmt.Println("保存下载文件到本地磁盘出错: " + err.Error())
		return errors.New("保存下载文件到本地磁盘出错: " + err.Error())

	} else {
		// 计算时长
		all_time := int(math.Floor(time.Since(start_time).Seconds()))

		// 耗时计算
		elapsed_time := GetElapsedTimeString(all_time)

		// 本地文件的大小
		file_size_s, file_size := base.GetFileSize(target_addr)

		// 比较本地和远程文件的大小
		if file_size == res.ContentLength {

			// 完整进度条
			jdt := strings.Repeat("=", 50)

			// 最终百分比
			download_percent := "100%"

			// 文字模板
			showstr := "\r已下载: %s/%s [%s] 已完成:%s 平均下载速度:%s 总耗时:%s          \n"

			// 平均速度计算
			avg_str = GetAvgStr(file_size, all_time)

			// 打印下载进度条
			fmt.Printf(showstr, file_size_s, remote_size, jdt, download_percent, avg_str, elapsed_time)

		} else {
			fmt.Println("下载错误: 本地文件和远程文件的大小不一致！")
			return errors.New("下载错误: 本地文件和远程文件的大小不一致！")
		}
	}

	//没有错误,返回nil
	return nil
}
