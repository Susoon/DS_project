package main

import (
    "context"
    "flag"
    "fmt"
    "net"
    "bufio"
    //"net/rpc"
    "os"
    "sync"
    "errors"

    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
    pb "DS_project/proto"
)


type Room struct {
    sync.Mutex

    name string
    passwd string
    manager *User
    user_list []*User
    chat_log []string
}

type User struct {
    sync.Mutex

    name string
    passwd string
    active bool
    room_list []*Room
    cur_room *Room
    message_stream chan string
}

func newRoom(user *User, name string, passwd string) (*Room, string) {
    room := new(Room)

    room.name = name
    room.passwd = passwd
    room.manager = user;

    return room, room.GetInfoString("Room %s created.\n", room.name)
}

func (room *Room) GetInfoString(format string, a ...interface{}) string {
    return fmt.Sprintf("<" + room.name + "> : " + format, a...)
}

func (room *Room) AddUser(user *User) {
    room.Lock()
    defer room.Unlock()

    if idx := UserContains(room.user_list, user.name); idx == -1 {
        room.user_list = append(room.user_list, user)
    }
}

func (room *Room) DelUser(user *User) {
    room.Lock()
    defer room.Unlock()

    user_idx := UserContains(room.user_list, user.name)

    room.user_list[user_idx] = room.user_list[len(room.user_list) - 1]
    room.user_list = room.user_list[:len(room.user_list) - 1]

}

func (room *Room) CheckPasswd(passwd string) bool {
    if room.passwd == passwd {
        return true
    } else {
        return false
    }
}

func (room *Room) AddChatLog(chat string) {
    room.Lock()
    defer room.Unlock()
    room.chat_log = append(room.chat_log, chat)
}

func (room *Room) GetChatLog() string {
    room.Lock()
    defer room.Unlock()

    ret_str := fmt.Sprintf("Room %s Chat Log.\n", room.name)
    for _, chat := range room.chat_log {
        ret_str += chat
    }

    return ret_str
}

func RoomContains(room_list []*Room, room_name string) int {
    for idx, r := range room_list{
        if r.name == room_name {
            return idx
        }
    }

    return -1
}

func PrintRoomList(room_list []*Room) string {
    ret_str := "\n===================================\n"
    for _, room := range room_list {
        ret_str += fmt.Sprintf("Room name : %s.\n", room.name)
    }
    ret_str += "===================================\n\n"

    return ret_str
}

func newUser(name string, passwd string) (*User, string) {
    user := new(User)

    user.name = name
    user.passwd = passwd

    user.active = true

    user.message_stream = make(chan string)

    return user, user.GetInfoString("User %s Created!\n", user.name)
}

func (user *User) GetInfoString(format string, a ...interface{}) string {
    return fmt.Sprintf("[" + user.name + "] : " + format, a...)
}

func (user *User) GetChatString(chat_content string) string {
    if user.cur_room != nil {
        return "<" + user.cur_room.name + "> [" + user.name + "] : " + chat_content
    } else {
        return ""
    }
}

func (user *User) ListMyRooms() string {
    ret_str := user.GetInfoString("%s has following rooms.\n", user.name)
    for _, room := range user.room_list {
        user_num := 0
        for _, member := range room.user_list{
            if member.cur_room.name == room.name {
                user_num += 1
            }
        }
        ret_str += fmt.Sprintf("Room name : %s. Number of users in the room : %d.\n", room.name, user_num)
    }

    return ret_str
}

func (user *User) CheckPasswd(passwd string) bool {
    if user.passwd == passwd {
        return true
    } else {
        return false
    }
}

func (user *User) LeaveChatRoom() (string, *Room) {
    if user.cur_room == nil{
        return user.GetInfoString("You are not in any chat rooms!\n"), nil
    } else{

        cur_room := user.cur_room

        user.Lock()
        defer user.Unlock()

        user.cur_room = nil

        return cur_room.GetInfoString("%s left from %s.\n", user.name, cur_room.name), cur_room
    }
}

func (user *User) AddRoom(room *Room) {
    user.Lock()
    defer user.Unlock()

    if idx := RoomContains(user.room_list, room.name); idx == -1 {
        user.room_list = append(user.room_list, room)
    }
}

func (user *User) DelRoom(room *Room) {
    user.Lock()
    defer user.Unlock()

    room_idx := RoomContains(user.room_list, room.name)

    user.room_list[room_idx] = user.room_list[len(user.room_list) - 1]
    user.room_list = user.room_list[:len(user.room_list) - 1]

}

func UserContains(user_list []*User, user_name string) int {
    for idx, u := range user_list{
        if u.name == user_name {
            return idx
        }
    }

    return -1
}

func EnterChatRoom(user *User, room_name string, room_list []*Room) string {
    idx := RoomContains(room_list, room_name)

    if idx == -1 {
        return user.GetInfoString("There is no room which has name %s!\n", room_name)
    } else {
        user.LeaveChatRoom()

        my_idx := RoomContains(user.room_list, room_name)

        user.Lock()
        defer user.Unlock()

        if my_idx == -1 {
            user.room_list = append(user.room_list, room_list[idx])
        }

        user.cur_room = room_list[idx]

        return user.cur_room.GetInfoString("%s entered to %s.\n", user.name, room_name)
    }
}

func PrintUserList(user_list []*User) string {
    ret_str := "\n===================================\n"

    for _, user := range user_list {
        ret_str += fmt.Sprintf("User name : %s.\n", user.name)
    }
    ret_str += "===================================\n\n"

    return ret_str

}

type Server struct {
    //pb.UnimplementedChatManagerServer
    user_list []*User
    room_list []*Room
}

var (
    addr = flag.String("addr", "localhost", "The server addr")
    port = flag.Int("port", 8080, "The server port")
)

func (s *Server) CreateAccount(in *pb.CreateAccountRequest, stream pb.ChatManager_CreateAccountServer) error {
    if UserContains(s.user_list, in.GetName()) != -1 {
        fmt.Printf("Duplicated Name!\n");
        var err = errors.New("Duplicated name")
        stream.Send(&pb.CreateAccountReply{Message: fmt.Sprintf("[Server] User name %s already exists!\n", in.GetName())})
        return err
        //return &pb.CreateAccountReply{Message: fmt.Sprintf("[Server] User name %s already exists!\n", in.GetName())}, err
    } else{
        fmt.Printf("Create Account! name : %s passwd : %s\n", in.GetName(), in.GetPasswd());
        user, create_account_string := newUser(in.GetName(), in.GetPasswd())

        s.user_list = append(s.user_list, user)

        stream.Send(&pb.CreateAccountReply{Message: create_account_string})

        for {
            select {
            case <-stream.Context().Done():
                return nil
            case chat_message := <-user.message_stream:
                if s, ok := status.FromError(stream.Send(&pb.CreateAccountReply{Message: chat_message})); ok {
                    switch s.Code() {
                    case codes.OK:
                        // noop
                    case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
                        fmt.Printf("Client %s terminated connection", user.name)
                        break
                    default:
                        fmt.Printf("Failed to send to client %s: %v", user.name, s.Err())
                        break
                    }
                }
            }
        }


        <-stream.Context().Done()
        return nil
        //return &pb.CreateAccountReply{Message: create_account_string}, nil
    }
}

func (s *Server) LogInAccount(in *pb.LogInAccountRequest, stream pb.ChatManager_LogInAccountServer) error {
    idx := UserContains(s.user_list, in.GetName())
    if idx != -1 {
        user := s.user_list[idx]
        if user.CheckPasswd(in.GetPasswd()) {
            fmt.Printf("Log In Success! name : %s passwd : %s\n", in.GetName(), in.GetPasswd());

            stream.Send(&pb.LogInAccountReply{Message: user.GetInfoString("Log In Success!\n")})

            for {
                select {
                case <-stream.Context().Done():
                    return nil
                case chat_message := <-user.message_stream:
                    if s, ok := status.FromError(stream.Send(&pb.LogInAccountReply{Message: chat_message})); ok {
                        switch s.Code() {
                        case codes.OK:
                            // noop
                        case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
                            fmt.Printf("Client %s terminated connection", user.name)
                            break
                        default:
                            fmt.Printf("Failed to send to client %s: %v", user.name, s.Err())
                            break
                        }
                    }
                }
            }

            <-stream.Context().Done()

            return nil
            //return &pb.LogInAccountReply{Message: user.GetInfoString("Log In Success!\n")}, nil
        }
        var err = errors.New("Wrong Password")
        fmt.Printf("Wrong Passwd! name : %s passwd : %s\n", in.GetName(), in.GetPasswd())

        stream.Send(&pb.LogInAccountReply{Message: fmt.Sprintf("[Server] Wrong Password!\n")})

        return err
        //return &pb.LogInAccountReply{Message: fmt.Sprintf("[Server] Wrong Password!\n")}, err
    } else{
        var err = errors.New("Not exists")
        fmt.Printf("Not exists! name : %s passwd : %s\n", in.GetName(), in.GetPasswd());

        stream.Send(&pb.LogInAccountReply{Message: fmt.Sprintf("[Server] User name %s does not exist!\n", in.GetName())})

        return err
        //return &pb.LogInAccountReply{Message: fmt.Sprintf("[Server] User name %s does not exist!\n", in.GetName())}, err
    }

}

func (s *Server) EnterChatRoom(ctx context.Context, in *pb.EnterChatRoomRequest) (*pb.EnterChatRoomReply, error) {
    user_idx := UserContains(s.user_list, in.GetUserName())
    room_idx := RoomContains(s.room_list, in.GetRoomName())

    user := s.user_list[user_idx]

    if room_idx != -1 {
        room := s.room_list[room_idx]
        if room.CheckPasswd(in.GetPasswd()) {
            enter_str := EnterChatRoom(user, room.name, s.room_list)
            fmt.Printf("Enter the Chat Room! member name : %s passwd : %s\n", in.GetUserName(), in.GetPasswd())
            user.cur_room = room
            room.AddUser(user)

            for _, member := range room.user_list {
                if member.cur_room != nil && member.cur_room.name == room.name && member.name != user.name{
                    member.message_stream <-enter_str
                }
            }

            return &pb.EnterChatRoomReply{Message: enter_str}, nil
        } else {
            err := errors.New("Wrong Password")
            enter_str := room.GetInfoString("Wrong Password!\n")
            fmt.Printf("Wrong Passwd! name : %s passwd : %s\n", in.GetUserName(), in.GetPasswd())

            return &pb.EnterChatRoomReply{Message: enter_str}, err
        }
    } else {
        room, create_room_string := newRoom(user, in.GetRoomName(), in.GetPasswd())
        s.room_list = append(s.room_list, room)
        enter_str := EnterChatRoom(user, room.name, s.room_list)
        fmt.Printf("Create the Chat Room! manager name : %s passwd : %s\n", in.GetUserName(), in.GetPasswd())
        user.cur_room = room
        room.AddUser(user)

        return &pb.EnterChatRoomReply{Message: create_room_string + enter_str}, nil
    }
}

func (s *Server) LeaveChatRoom(ctx context.Context, in *pb.LeaveChatRoomRequest) (*pb.LeaveChatRoomReply, error) {
    user_idx := UserContains(s.user_list, in.GetUserName())

    user := s.user_list[user_idx]

    ret_str, cur_room := user.LeaveChatRoom()
    fmt.Printf("Leave the Chat Room! user name : %s room name : %s\n", in.GetUserName(), cur_room.name)

    for _, member := range cur_room.user_list {
        if member.cur_room != nil && member.cur_room.name == cur_room.name && member.name != user.name{
            member.message_stream <-ret_str
        }
    }

    cur_room.DelUser(user)
    user.cur_room = nil
    user.DelRoom(cur_room)

    return &pb.LeaveChatRoomReply{Message: ret_str}, nil
}

func (s *Server) GetListOfRoom(ctx context.Context, in *pb.GetListOfRoomRequest) (*pb.GetListOfRoomReply, error) {
    user_idx := UserContains(s.user_list, in.GetUserName())
    if user_idx == -1 {
        fmt.Printf("Try to get list of chat room without logging in!\n")
        err := errors.New("Try to get list of chat room without logging in!")
        return &pb.GetListOfRoomReply{Message: "Try to get list of chat room without logging in!\n"}, err
    }

    user := s.user_list[user_idx]

    ret_str := user.ListMyRooms()
    fmt.Printf("Get List of Chat Rooms! name : %s\n", in.GetUserName())

    return &pb.GetListOfRoomReply{Message: ret_str}, nil
}

func (s *Server) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*pb.SendMessageReply, error) {
    user_idx := UserContains(s.user_list, in.GetUserName())
    if user_idx == -1 {
        fmt.Printf("Try to get list of chat room without logging in!\n")
        err := errors.New("Try to get list of chat room without logging in!")
        return &pb.SendMessageReply{Message: "Try to get list of chat room without logging in!\n"}, err
    }

    user := s.user_list[user_idx]

    room_idx := RoomContains(s.room_list, in.GetRoomName())
    if room_idx == -1 {
        fmt.Printf("Try to enter non-exist room!\n")
        err := errors.New("Try to enter non-exist room!\n")
        return &pb.SendMessageReply{Message: "Try to enter non-exist room!\n"}, err
    }

    room := s.room_list[room_idx]

    content := user.GetChatString(in.GetContent()) + "\n"

    room.AddChatLog(content)

    for _, member := range room.user_list {
        if member.cur_room != nil && member.cur_room.name == room.name && member.name != user.name{
            member.message_stream <-content
        }
    }

    fmt.Printf("Send Message! name : %s content : %s\n", in.GetUserName(), in.GetContent())

    return &pb.SendMessageReply{Message: content}, nil
}

var user_list []*User
var room_list []*Room

func main() {
    flag.Parse()

    listener, err := net.Listen("tcp", *addr + fmt.Sprintf(":%d", *port))

    if err != nil {
        fmt.Printf("Failed to listen : %v\n", err)
    }

    server := &Server{}

    grpc_server:= grpc.NewServer()

    go func() {
        fmt.Printf("\n===================================\n")
        fmt.Printf("Enter command.\n")
        fmt.Printf("Command List.\n")
        fmt.Printf("print chat log : Create your account.\n")
        fmt.Printf("===================================\n\n")

        scanner := bufio.NewScanner(os.Stdin)

        for {
            scanner.Scan()

            if scanner.Text() == "print chat log" {
                fmt.Printf("Enter room name.\n")
                fmt.Printf("Room List")
                PrintRoomList(server.room_list)

                scanner.Scan()

                room_idx := RoomContains(server.room_list, scanner.Text())

                if room_idx == -1 {
                    fmt.Printf("Wrong Room Name.\n")
                } else {
                    room := server.room_list[room_idx]

                    fmt.Printf(room.GetChatLog())
                }
            }
        }

    }()

    pb.RegisterChatManagerServer(grpc_server, server)
    grpc_server.Serve(listener)
}
