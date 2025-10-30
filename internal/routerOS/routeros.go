package routeros

import (
	"fmt"

	"github.com/go-routeros/routeros/v3"
)




type MikSvc struct {
	conn *routeros.Client
}


func  Initialize() (*MikSvc , error ) {
	conn , err := routeros.Dial("192.168.100.88:8728" , "ali" , "Ali@1384")
	if err != nil {
		return nil , fmt.Errorf("error , problem to connect mikrotik routerOS")
	}
	fmt.Println("connection to routerOS successful")

	return &MikSvc{
		conn: conn,
	}, nil
}

func (m *MikSvc) GetCpu () {
	reply , err := m.conn.Run("/system/resource/cpu/print")

	if err != nil {
		fmt.Println("invalid command")
		return
	}

	for _ , re := range reply.Re {
		fmt.Printf("%v\n" , re.Map["load"])
	}

}

func (m *MikSvc) GetFreeMem () {
	reply , err := m.conn.Run("/system/resource/print")

	if err != nil {
		fmt.Println("invalid command")
		return
	}

	for _ , re := range reply.Re {
		fmt.Printf("%v\n" , re.Map["free-memory"])
	}
}

func (m *MikSvc) GetFreeSpace () {
	reply , err := m.conn.Run("/system/resource/print")

	if err != nil {
		fmt.Println("invalid command")
		return
	}

	for _ , re := range reply.Re {
		fmt.Printf("%v\n" , re.Map["free-hdd-space"])
	
	}

}



