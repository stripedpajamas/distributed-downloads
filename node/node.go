package node

import "os"

// Peer defines IP Address and Port of local or remote nodes
type Peer struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

// Configuration defines the fields a user can supply to the node
type Configuration struct {
	Port      int `json:"port"`
	ChunkSize int `json:"chunkSize"` // in bytes
	Peers     struct {
		Local  []Peer `json:"local"`
		Remote []Peer `json:"remote"`
	} `json:"peers"`
}

// Session defines the parameters with which peers can
// cooperatively and efficiently assemble the contents of a file
type Session struct {
	SessionID int `json:"sessionId"`
	// Owners have the file wanted and this node will
	// request chunks of data from peers in the Owners list
	Owners []Peer `json:"owners"`
	// Controller is the Peer to which we will serve all gathered chunks
	Controller Peer   `json:"controller"`
	Sum        []byte `json:"sum"`       // the final SHA256 hash of the file
	FileID     int    `json:"fileId"`    // to keep track of things locally
	FileSize   int    `json:"fileSize"`  // in bytes
	ChunkSize  int    `json:"chunkSize"` // in bytes
}

// File represents the data that is requested/downloaded/saved/served
type File struct {
	FileID   string
	FileSize int
	File     *os.File // references the actual file on disk
}

// Chunk represents the byte start and byte end of a file
type Chunk struct {
	Sum   []byte // the SHA256 hash of the chunk
	Start int
	End   int
}

// Request is the request that comes from controller
// chunk id is the 0-indexed index of the chunk in the file
type Request struct {
	SessionID int `json:"sessionId"`
	ChunkID   int `json:"chunkId"`
}
