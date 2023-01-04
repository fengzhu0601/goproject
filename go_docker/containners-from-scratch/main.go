/**
https://github.com/lizrice/containers-from-scratch/blob/master/README.md

https://www.bilibili.com/video/BV1EU4y1J7sz/?spm_id_from=333.999.0.0&vd_source=09679dd913f0422efa045f2d93646a59
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

/*
docker run <cmd> <args>
*/
func run() {
	fmt.Printf("Running run %v , pid:%d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	/*
		syscall.CLONE_NEWUTS 隔离主机与容器的UTS Namespace
		syscall.CLONE_NEWPID 隔离进程和修改容器跟目录
		syscall.CLONE_NEWNS 隔离挂载命名空间 Mount Namespace
	*/
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running child %v, pid:%d \n", os.Args[2:], os.Getpid())

	cg()

	// 修改容器的主机名
	must(syscall.Sethostname([]byte("container")))
	// 文件系统 http://cdimage.ubuntu.com/ubuntu-base/releases/22.04/release/
	must(syscall.Chroot("./ubuntu-fs"))
	//must(syscall.Chdir("/"))

	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	////must(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())

	must(syscall.Unmount("/proc", 0))
	//must(syscall.Unmount("thing", 0))
}

// 限制容器使用的资源
func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")

	if _, err := os.Stat(filepath.Join(pids, "container")); os.IsNotExist(err) {
		must(os.Mkdir(filepath.Join(pids, "container"), 0755))
	}

	//must(ioutil.WriteFile(filepath.Join(pids, "container/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	//must(ioutil.WriteFile(filepath.Join(pids, "container/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "container/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
