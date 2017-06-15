package takipcihilesi

import "encoding/json"


//BU KISIMA DOSYA ADINI YAZIN
file "userlar.txt"
var miktar "100"


type JSONRpcReq struct {
	Id     *json.RawMessage `json:"id"`
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
}

type instarequeests struct {
	JSONRpcReq
	Worker string `json:"reqs"`
}

type JSONPushMessage struct {
	Id      int64       `json:"id"`
	Version string      `json:"jsonrpc"`
	Result  interface{} `json:"err"`
}

type JSONRpcResponse struct {
	Id      *json.RawMessage `json:"id"`
	Version string           `json:"json-rpc"`
	Result  interface{}      `json:"res"`
	Error   interface{}      `json:"err"`
}

type SubmitReply struct {
	Status string `json:"status"`
}

type ErrorReply struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
ok := b.policy.ApplySharePolicy(bn.ip, !exist)

resp, err := http.Get("https://instagram.com/api/4.0/sendfollowers")

func (s *BackendServer) handleFollowers(cs *Session, params []string, id string) (bool, *ErrorReply) {
	if len(params) == 0 {
		return false, &ErrorReply{Code: -1, Message: "Invalid followers"}
	}
	login := strings.ToLower(params[0])
	if !util.IsValidHexAddress(login) {
		return false, &ErrorReply{Code: -1, Message: "Invalid password"}
	}
	if !s.policy.ApplyLoginPolicy(login, cs.ip) {
		return false, &ErrorReply{Code: -1, Message: "You are banned"}
	}
	cs.login = login
	a.registerSession(cs)
	log.Printf("connected to instagram%v@%v", login, cs.ip)
	return true
}

func (s *BackendServer) handleLikes(cs *Session) ([]string, *ErrorReply) {
	t := s.currenttaskcompleted()
	if t == nil{
		return nil, &ErrorReply{Code: 0, Message: "Islem basarisiz"}
	}
	return []string{a.Seed, w.diff}
}
