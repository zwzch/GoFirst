package mobula

import (
	"os"
	"fmt"
	"log"
	"time"
)

func TestT(){
	logFile, err := os.Create("/Users/edz/goworkspace/src/GoFirst/mobula/" + time.Now().Format("goerror") + ".txt");
	if err != nil {
		fmt.Println(err);
	}
	loger := log.New(logFile, "test_", log.Ldate|log.Ltime|log.Lshortfile);

	defer func() {
		if err := recover(); err != nil {
			//fmt.Println("Exception has been caught.")
			//fmt.Println(err)
			userFile :="myfile"
			fout,err :=os.Create(userFile)
			defer fout.Close()
			if err != nil{
				fmt.Println(userFile,err)
				return
			}

			fout.WriteString("HelloWorld")
			fout.Write([]byte("abcd!\n"))
		}
	}()
	loger.Println("lala")
	//fmt.Println(i)

}