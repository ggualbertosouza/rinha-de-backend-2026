package worker

import (
	"net"

	"golang.org/x/sys/unix"
)

func recvFD(conn *net.UnixConn) (int, error) {
	buf := [1]byte{}
	oob := make([]byte, 128)

	file, err := conn.File()
	if err != nil {
		return 0, err
	}
	defer file.Close()

	n, oobn, _, _, err := unix.Recvmsg(
		int(file.Fd()),
		buf[:],
		oob,
		0,
	)
	if err != nil {
		return 0, err
	}

	if n == 0 || oobn == 0 {
		return 0, nil
	}

	scms, err := unix.ParseSocketControlMessage(oob[:oobn])
	if err != nil {
		return 0, err
	}

	fds, err := unix.ParseUnixRights(&scms[0])
	if err != nil {
		return 0, err
	}

	return fds[0], nil
}
