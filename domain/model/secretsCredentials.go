package model


type DBCredentials struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int64  `json:"port"`
	DBClusterIdentifier string `json:"dbClusterIdentifier"`
}
