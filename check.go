package main
import(
	"fmt"
	"net"
	"time"

)
func Check(destination string, port string) string{
	address := destination + ":" +port
	timeout := time.Duration(5 * time.Second)
	conn,err := net.DialTimeout("tcp", address, timeout)
	var status string

	if err!=nil{
		status = fmt.Sprintf("%s is down! Error: %s", destination, err.Error())
	} else {
		status = fmt.Sprintf("%s is up!\n From: %v\n To:%v\n", destination, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}