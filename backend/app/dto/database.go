package dto

import "time"

type MysqlDBInfo struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Name        string    `json:"name"`
	Format      string    `json:"format"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Permission  string    `json:"permission"`
	BackupCount int       `json:"backupCount"`
	Description string    `json:"description"`
}

type MysqlDBCreate struct {
	Name        string `json:"name" validate:"required"`
	MysqlName   string `json:"mysqlName" validate:"required"`
	Format      string `json:"format" validate:"required,oneof=utf8mb4 utf-8 gbk big5"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Permission  string `json:"permission" validate:"required"`
	Description string `json:"description"`
}

type MysqlStatus struct {
	AbortedClients               string `json:"Aborted_clients"`
	AbortedConnects              string `json:"Aborted_connects"`
	BytesReceived                string `json:"Bytes_received"`
	BytesSent                    string `json:"Bytes_sent"`
	ComCommit                    string `json:"Com_commit"`
	ComRollback                  string `json:"Com_rollback"`
	Connections                  string `json:"Connections"`
	CreatedTmpDiskTables         string `json:"Created_tmp_disk_tables"`
	CreatedTmpTables             string `json:"Created_tmp_tables"`
	InnodbBufferPoolPagesDirty   string `json:"Innodb_buffer_pool_pages_dirty"`
	InnodbBufferPoolReadRequests string `json:"Innodb_buffer_pool_read_requests"`
	InnodbBufferPoolReads        string `json:"Innodb_buffer_pool_reads"`
	KeyReadRequests              string `json:"Key_read_requests"`
	KeyReads                     string `json:"Key_reads"`
	KeyWriteEequests             string `json:"Key_write_requests"`
	KeyWrites                    string `json:"Key_writes"`
	MaxUsedConnections           string `json:"Max_used_connections"`
	OpenTables                   string `json:"Open_tables"`
	OpenedFiles                  string `json:"Opened_files"`
	OpenedTables                 string `json:"Opened_tables"`
	QcacheHits                   string `json:"Qcache_hits"`
	QcacheInserts                string `json:"Qcache_inserts"`
	Questions                    string `json:"Questions"`
	SelectFullJoin               string `json:"Select_full_join"`
	SelectRangeCheck             string `json:"Select_range_check"`
	SortMergePasses              string `json:"Sort_merge_passes"`
	TableLocksWaited             string `json:"Table_locks_waited"`
	ThreadsCached                string `json:"Threads_cached"`
	ThreadsConnected             string `json:"Threads_connected"`
	ThreadsCreated               string `json:"Threads_created"`
	ThreadsRunning               string `json:"Threads_running"`
	Uptime                       string `json:"Uptime"`
	Run                          string `json:"Run"`
	File                         string `json:"File"`
	Position                     string `json:"Position"`
}

type MysqlVariables struct {
	BinlogCachSize       string `json:"binlog_cache_size"`
	InnodbBufferPoolSize string `json:"innodb_buffer_pool_size"`
	InnodbLogBufferSize  string `json:"innodb_log_buffer_size"`
	JoinBufferSize       string `json:"join_buffer_size"`
	KeyBufferSize        string `json:"key_buffer_size"`
	MaxConnections       string `json:"max_connections"`
	MaxHeapTableSize     string `json:"max_heap_table_size"`
	QueryCacheSize       string `json:"query_cache_size"`
	QueryCache_type      string `json:"query_cache_type"`
	ReadBufferSize       string `json:"read_buffer_size"`
	ReadRndBufferSize    string `json:"read_rnd_buffer_size"`
	SortBufferSize       string `json:"sort_buffer_size"`
	TableOpenCache       string `json:"table_open_cache"`
	ThreadCacheSize      string `json:"thread_cache_size"`
	ThreadStack          string `json:"thread_stack"`
	Tmp_tableSize        string `json:"tmp_table_size"`
}

type MysqlVariablesUpdate struct {
	MysqlName            string `json:"mysqlName" validate:"required"`
	KeyBufferSize        int64  `json:"key_buffer_size" validate:"required"`
	QueryCacheSize       int64  `json:"query_cache_size" validate:"required"`
	TmpTableSize         int64  `json:"tmp_table_size" validate:"required"`
	InnodbBufferPoolSize int64  `json:"innodb_buffer_pool_size" validate:"required"`
	InnodbLogBufferSize  int64  `json:"innodb_log_buffer_size" validate:"required"`
	SortBufferSize       int64  `json:"sort_buffer_size" validate:"required"`
	ReadBufferSize       int64  `json:"read_buffer_size" validate:"required"`

	ReadRndBufferSize int64 `json:"read_rnd_buffer_size" validate:"required"`
	JoinBufferSize    int64 `json:"join_buffer_size" validate:"required"`
	ThreadStack       int64 `json:"thread_stack" validate:"required"`
	BinlogCachSize    int64 `json:"binlog_cache_size" validate:"required"`
	ThreadCacheSize   int64 `json:"thread_cache_size" validate:"required"`
	TableOpenCache    int64 `json:"table_open_cache" validate:"required"`
	MaxConnections    int64 `json:"max_connections" validate:"required"`
}
type MysqlConfUpdateByFile struct {
	MysqlName string `json:"mysqlName" validate:"required"`
	File      string `json:"file"`
}

type ChangeDBInfo struct {
	ID        uint   `json:"id"`
	MysqlName string `json:"mysqlName" validate:"required"`
	Operation string `json:"operation" validate:"required,oneof=password privilege"`
	Value     string `json:"value" validate:"required"`
}

type DBBaseInfo struct {
	Name       string `json:"name"`
	Port       int64  `json:"port"`
	Password   string `json:"password"`
	RemoteConn bool   `json:"remoteConn"`
	MysqlKey   string `json:"mysqlKey"`
}

type SearchDBWithPage struct {
	PageInfo
	MysqlName string `json:"mysqlName" validate:"required"`
}

type SearchBackupsWithPage struct {
	PageInfo
	MysqlName string `json:"mysqlName" validate:"required"`
	DBName    string `json:"dbName" validate:"required"`
}

type BackupDB struct {
	MysqlName string `json:"mysqlName" validate:"required"`
	DBName    string `json:"dbName" validate:"required"`
}

type RecoverDB struct {
	MysqlName  string `json:"mysqlName" validate:"required"`
	DBName     string `json:"dbName" validate:"required"`
	BackupName string `json:"backupName" validate:"required"`
}

// redis
type RedisConfUpdate struct {
	Timeout     string `json:"timeout"`
	Maxclients  string `json:"maxclients"`
	Requirepass string `json:"requirepass"`
	Maxmemory   string `json:"maxmemory"`
}
type RedisConfPersistenceUpdate struct {
	Type        string `json:"type" validate:"required,oneof=aof rbd"`
	Appendonly  string `json:"appendonly"`
	Appendfsync string `json:"appendfsync"`
	Save        string `json:"save"`
}
type RedisConfUpdateByFile struct {
	File       string `json:"file"`
	RestartNow bool   `json:"restartNow"`
}

type RedisConf struct {
	Name          string `json:"name"`
	ContainerName string `json:"containerName"`
	Timeout       string `json:"timeout"`
	Maxclients    string `json:"maxclients"`
	Requirepass   string `json:"requirepass"`
	Maxmemory     string `json:"maxmemory"`
}

type RedisPersistence struct {
	Appendonly  string `json:"appendonly"`
	Appendfsync string `json:"appendfsync"`
	Save        string `json:"save"`
}

type RedisStatus struct {
	TcpPort                  string `json:"tcp_port"`
	UptimeInDays             string `json:"uptime_in_days"`
	ConnectedClients         string `json:"connected_clients"`
	UsedMemory               string `json:"used_memory"`
	UsedMemory_rss           string `json:"used_memory_rss"`
	UsedMemory_peak          string `json:"used_memory_peak"`
	MemFragmentationRatio    string `json:"mem_fragmentation_ratio"`
	TotalConnectionsReceived string `json:"total_connections_received"`
	TotalCommandsProcessed   string `json:"total_commands_processed"`
	InstantaneousOpsPerSec   string `json:"instantaneous_ops_per_sec"`
	KeyspaceHits             string `json:"keyspace_hits"`
	KeyspaceMisses           string `json:"keyspace_misses"`
	LatestForkUsec           string `json:"latest_fork_usec"`
}

type RedisBackupRecords struct {
	FileName  string `json:"fileName"`
	FileDir   string `json:"fileDir"`
	CreatedAt string `json:"createdAt"`
	Size      int    `json:"size"`
}
type RedisBackupRecover struct {
	FileName string `json:"fileName"`
	FileDir  string `json:"fileDir"`
}
type RedisBackupDelete struct {
	FileDir string   `json:"fileDir"`
	Names   []string `json:"names"`
}
