c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	brk := false
	for {
		go func(){
			var msg string
			_, err := fmt.Fscan(conn1, &msg)
			if err != nil{
				brk = true
			}
			c1 <- msg
		}()

		go func(){
			var msg string
			_, err := fmt.Fscan(conn2, &msg)
			if err != nil{
				brk = true
			}
			c2 <- msg
		}()

		if(brk){
			conn1.Close()
			conn2.Close()
			return 
		}
		var msg string
		
		select {
			case msg = <-c1:
				fmt.Fprintf(conn2, "%s\n", msg)
			case msg = <-c2:
				fmt.Fprintf(conn1, "%s\n", msg)
			case <-time.After(60 * time.Second):
				conn1.Close()
				conn2.Close()
				return
		}

		if msg == "end" {
			conn1.Close()
			conn2.Close()
			return
		}