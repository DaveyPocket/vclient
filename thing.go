package main

import ("fmt"
			"net"
			"bufio"
			"strconv"
		)


func main() {
	conn, err := net.Dial("tcp", "games.recurse.com:7000")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	i := 5
	fmt.Println(strconv.Itoa(i) + " " + strconv.Itoa(1))

	m := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	s, _, _ := m.ReadLine()
	fmt.Println(string(s[:]))
	var temp string
	for {
		fmt.Scanln(&temp)
		fmt.Println(temp)
		m.WriteString(temp + " " + temp + "\n")
		m.Flush()
		s, _, _ := m.ReadLine()
		fmt.Println(string(s[:]))
	}
}
