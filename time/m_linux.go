package time

// 执行系统命令
func ExecBashShell(cmdLine string) (string, error) {
	// 所有变量都要先定义，避免使用直接赋值的变量
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成 cmd
	cmd = exec.Command("/bin/bash", "-c", cmdLine)

	// 执行命令，捕获了子进程的输出(pipe)
	if output, err = cmd.CombinedOutput(); err != nil {
		// 返回错误信息
		return "", err
	} else {
		// 返回输出内容
		return string(output), nil
	}
}

// 设置系统时间
func SetLinuxSystemTime(tt int64) error {
	// 时间戳转成本地时间
	localTime := TimestampToLocal(tt)

	// 组合命令语句
	cmdLine := fmt.Sprintf("date -s '%v'", localTime)

	// 执行命令
	if _, err := ExecBashShell(cmdLine); err != nil {
		return err
	} else {
		// 写入CMOS
		return syncToCOMS()
	}
}

// 设置系统时区
func SetLinuxTimeZone(tz int) error {
	// 先获取时区对应的城市
	city := GetTimeZoneCity(tz)

	// 判断值是否正确
	if len(city) == 0 {
		beego.Error("时区错误，取值范围应该是-12到12之间的整数。但是，接收到的参数是:", tz)
		// 1003 表示参数错误
		return errors.New("1003")
	}

	// 组合命令语句
	cmdLine := fmt.Sprintf("timedatectl set-timezone '%v'", city)

	// 执行命令
	if _, err := ExecBashShell(cmdLine); err != nil {
		return err
	} else {
		// 写入CMOS
		return syncToCOMS()
	}
}

// 将设置写入硬件时钟
func syncToCOMS() error {
	// 执行命令
	if _, err := ExecBashShell("hwclock -w"); err != nil {
		return err
	} else {
		return nil
	}
}
