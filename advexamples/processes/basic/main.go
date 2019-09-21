package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {
	fmt.Printf("current process id : %d \n", os.Getpid())
	fmt.Printf("parent process id : %d \n", os.Getppid())
	fmt.Println("User ID:", os.Getuid())
	fmt.Println("Group ID:", os.Getgid())
	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Group IDs:", groups)
	fmt.Println()
	demo2()
}

func demo2() {
	uid := os.Getuid()
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %s (uid %d)\n", u.Username, uid)
	gid := os.Getgid()
	group, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Group: %s (uid %d)\n", group.Name, uid)

}
