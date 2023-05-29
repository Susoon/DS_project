package main

import (
    "bufio"
    "context"
    "flag"
    "fmt"
    //"net"
    //"net/rpc"
    "os"
    "sync"
    "time"
    "io"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "DS_project/proto"
)

var (
    server_addr = flag.String("server_addr", "localhost", "The address for server")
    server_port = flag.Int("server_port", 8080, "The port for server")
    wait = &sync.WaitGroup{}
)

const (
    CREATE_ACCOUNT = "create account"
    LOG_IN_ACCOUNT = "log in"
    ENTER_CHAT_ROOM = "enter chat room"
    LEAVE_CHAT_ROOM = "leave chat room"
    GET_LIST_OF_ROOM = "get list"
)

func connect_with_create(c pb.ChatManagerClient, name string, passwd string) {
    ctx := context.Background()
    stream_create, err := c.CreateAccount(ctx, &pb.CreateAccountRequest{Name: name, Passwd: passwd})

    if err != nil {
        fmt.Printf("\n[System] %v.\n", err)
        return
    }

    wait.Add(1)
    go func(stream pb.ChatManager_CreateAccountClient) {
        defer wait.Done()
        recv, recv_err := stream.Recv()
        if recv_err != nil && recv_err != io.EOF {
            //fmt.Printf("Recv Failed!\n")
            return
        }
        fmt.Printf(recv.Message)

        for {
            msg, err := stream.Recv()
            if err != nil {
                if err != io.EOF {
                    //fmt.Printf("Failed to Read! %v\n", err)
                    return
                }
                continue
            }
            fmt.Printf(msg.Message)
        }
    }(stream_create)

    go func(ctext context.Context) {
        select {
        case <- ctext.Done():
            fmt.Printf("SayHello: context err %+v \n", ctext.Err())
            return
        }
    }(ctx)
}

func connect_with_login(stream_login pb.ChatManager_LogInAccountClient){

    go func(stream pb.ChatManager_LogInAccountClient) {
        for {
            msg, err := stream.Recv()
            if err != nil {
                if err != io.EOF {
                    //fmt.Printf("Failed to Read! %v\n", err)
                    return
                }
                continue
            }
            fmt.Printf(msg.Message + "\n")
        }
    }(stream_login)

}

func PrintCommand() {
    fmt.Printf("\n===================================\n")
    fmt.Printf("Enter command.\n")
    fmt.Printf("Command List.\n")
    fmt.Printf("%s : Create your account.\n", CREATE_ACCOUNT)
    fmt.Printf("%s : Log in to your account.\n", LOG_IN_ACCOUNT)
    fmt.Printf("%s : Enter the chat room.\n", ENTER_CHAT_ROOM)
    fmt.Printf("%s : Leave current chat room.\n", LEAVE_CHAT_ROOM)
    fmt.Printf("%s : Get the list of chat room.\n", GET_LIST_OF_ROOM)
    fmt.Printf("===================================\n\n")
}

func main() {
    //timestamp := time.Now()
    done := make(chan int)

    flag.Parse()

    conn, err := grpc.Dial(*server_addr + fmt.Sprintf(":%d", *server_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        fmt.Printf("Failed to connect! %v\n", err)
        return
    }
    defer conn.Close()

    c := pb.NewChatManagerClient(conn)

    scanner := bufio.NewScanner(os.Stdin)

    var user_name = "My name"
    var user_passwd = "1234"

    var room_name = "My Room name"
    var room_passwd = "1234"

    PrintCommand()

    conn.Connect()

    for{
        scanner.Scan()

        if scanner.Text() == CREATE_ACCOUNT {
            fmt.Printf("\n[System] Enter your name.\n")

            scanner.Scan()

            user_name = scanner.Text()

            fmt.Printf("\n[System] Enter your passwd.\n")

            scanner.Scan()

            user_passwd = scanner.Text()

            connect_with_create(c, user_name, user_passwd)

        } else if scanner.Text() == LOG_IN_ACCOUNT {
            fmt.Printf("\n[System] Enter your name.\n")

            scanner.Scan()

            user_name = scanner.Text()

            fmt.Printf("\n[System] Enter your passwd.\n")

            scanner.Scan()

            user_passwd = scanner.Text()

            stream, err := c.LogInAccount(context.Background(), &pb.LogInAccountRequest{Name: user_name, Passwd: user_passwd})

            for err != nil {
                recv, recv_err := stream.Recv()
                if recv_err != nil {
                    fmt.Printf("Recv Failed!\n")
                    break
                }
                fmt.Printf(recv.Message)
                fmt.Printf("\n[System] Re-enter your passwd.\n")
                scanner.Scan()

                stream, err = c.LogInAccount(context.Background(), &pb.LogInAccountRequest{Name: user_name, Passwd: scanner.Text()})
            }
            connect_with_login(stream)
            recv, recv_err := stream.Recv()
            if recv_err != nil {
                fmt.Printf("Recv Failed!\n")
                break
            }
            fmt.Printf(recv.Message)

        } else if scanner.Text() == ENTER_CHAT_ROOM {
            fmt.Printf("\n[System] Enter your chat room name.\n")

            scanner.Scan()

            room_name = scanner.Text()

            fmt.Printf("\n[System] Enter your chat room passwd.\n")

            scanner.Scan()

            room_passwd = scanner.Text()

            r, err := c.EnterChatRoom(context.Background(), &pb.EnterChatRoomRequest{UserName: user_name, RoomName: room_name, Passwd: room_passwd})

            for err != nil {
                fmt.Printf(r.GetMessage())
                fmt.Printf("\n[System] Re-enter your chat room passwd.\n")
                scanner.Scan()

                r, err = c.EnterChatRoom(context.Background(), &pb.EnterChatRoomRequest{UserName: user_name, RoomName: room_name, Passwd: scanner.Text()})
            }
            fmt.Printf(r.GetMessage())

        } else if scanner.Text() == LEAVE_CHAT_ROOM {
            r, _ := c.LeaveChatRoom(context.Background(), &pb.LeaveChatRoomRequest{UserName: user_name})
            fmt.Printf(r.GetMessage())
        } else if scanner.Text() == GET_LIST_OF_ROOM {
            r, _ := c.GetListOfRoom(context.Background(), &pb.GetListOfRoomRequest{UserName: user_name})
            fmt.Printf(r.GetMessage())
        } else {
            r, _ := c.SendMessage(context.Background(), &pb.SendMessageRequest{UserName: user_name, RoomName: room_name, Content: scanner.Text(), Timestamp: time.Now().String()})
            fmt.Printf(r.GetMessage())
        }
    }
    <-done
}
