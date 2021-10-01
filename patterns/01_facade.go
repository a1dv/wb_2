package pattern

import (
    "fmt"
)

type Commandline struct {
    input string
}

func (cmd *Commandline) typing_address() {
    fmt.Scan(&cmd.input)
}

type local_machine struct {
    history []string
    cache []string
}

func (cmd *Commandline) Search_on_click() {
    res, num := Search_for_server_locally(&cmd.input)
    if !res {
        var dns DNS_server
        res, num = dns.Request_DNS(&cmd.input)
        if res {
            var s Server
            show_info(s.request_info(num))
        } else {
            fmt.Println("This server doesn't exist")
        }
    }
}

func Search_for_server_locally(input *string) (bool, int) {
    var my_pc local_machine
    res, num := my_pc.search_history(input)
    if !res {
        res, num = my_pc.search_cache(input)
    }
    return res, num
}

func (my_pc local_machine) search_history(input *string) (bool, int) {
    for i,v := range my_pc.history {
        if v == *input {
            return true, i
        }
    }
    return false, -1
}

func (my_pc local_machine) search_cache(input *string) (bool, int){
    for i,v := range my_pc.history {
        if v == *input {
            return true, i
        }
    }
    return false, -1
}

type DNS_server struct {
    IP_addresses []string
}

func (dns DNS_server) Request_DNS(input *string) (bool, int){
    res, num := dns.ask_local_DNS(input)
    if !res {
        res, num = dns.ask_region_DNS(input)
        if !res {
            return false, -1
        }
    }
    return res, num
}

func (dns DNS_server) ask_local_DNS(input *string) (bool, int){
    for i, v := range dns.IP_addresses {
        if v == *input {
            return true, i
        }
    }
    return false, -1
}

func(dns DNS_server) ask_region_DNS(input *string) (bool, int){
    for i, v := range dns.IP_addresses {
        if v == *input {
            return true, i
        }
    }
    return false, -1
}

type Server struct {
    info []string
}

func (s Server) request_info(address int) []string{
    return s.info
}

func show_info(information []string) {
    fmt.Println(information)
}
