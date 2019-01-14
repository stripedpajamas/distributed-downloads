package node

// Peer defines IP Address and Port of local or remote nodes
type Peer struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// Configuration defines the fields a user can supply to the node
type Configuration struct {
	Port  int `json:"port"`
	Peers struct {
		Local  []Peer `json:"local"`
		Remote []Peer `json:"remote"`
	} `json:"peers"`
}
