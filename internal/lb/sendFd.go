package loadBalancer

import "golang.org/x/sys/unix"

func sendFD(socketFD int, fd int) error {
	return unix.Sendmsg(
		socketFD,
		[]byte{0},
		unix.UnixRights(fd),
		nil,
		0,
	)
}
