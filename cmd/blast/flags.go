//  Copyright (c) 2018 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"math"

	"github.com/urfave/cli"
)

var (
	flBindAddr = cli.StringFlag{
		Name:   "bind-addr",
		Value:  "127.0.0.1:10000",
		Usage:  "Address to listen on for peer traffic",
		EnvVar: "BLAST_BIND_ADDR",
	}
	flGRPCAddr = cli.StringFlag{
		Name:   "grpc-addr",
		Value:  "127.0.0.1:10001",
		Usage:  "Address to listen on for client traffic via gRPC",
		EnvVar: "BLAST_GRPC_ADDR",
	}
	flHTTPAddr = cli.StringFlag{
		Name:   "http-addr",
		Value:  "127.0.0.1:10002",
		Usage:  "Address to listen on for client traffic via HTTP",
		EnvVar: "BLAST_HTTP_ADDR",
	}

	flNodeID = cli.StringFlag{
		Name:   "node-id",
		Value:  "node0",
		Usage:  "Node ID",
		EnvVar: "BLAST_NODE_ID",
	}
	flRaftDir = cli.StringFlag{
		Name:   "raft-dir",
		Usage:  "Raft data directory",
		EnvVar: "BLAST_RAFT_DIR",
	}
	flRaftSnapshotCount = cli.StringFlag{
		Name:   "raft-snapshot-count",
		Usage:  "Raft snapshot count",
		EnvVar: "BLAST_RAFT_SNAPSHOT_COUNT",
	}
	flRaftTimeout = cli.StringFlag{
		Name:   "raft-timeout",
		Usage:  "Raft timeout",
		EnvVar: "BLAST_RAFT_TIMEOUT",
	}

	flStoreDir = cli.StringFlag{
		Name:   "store-dir",
		Usage:  "Store data directory",
		EnvVar: "BLAST_STORE_DIR",
	}

	flIndexDir = cli.StringFlag{
		Name:   "index-dir",
		Usage:  "Index data directory",
		EnvVar: "BLAST_INDEX_DIR",
	}
	flIndexMapping = cli.StringFlag{
		Name:   "index-mapping",
		Usage:  "Index mapping path",
		EnvVar: "BLAST_INDEX_MAPPING",
	}
	flIndexType = cli.StringFlag{
		Name:   "index-type",
		Usage:  "Index type",
		EnvVar: "BLAST_INDEX_TYPE",
	}
	flIndexKvstore = cli.StringFlag{
		Name:   "index-kvstore",
		Usage:  "Index Key-Value store",
		EnvVar: "BLAST_INDEX_KVSTORE",
	}
	flIndexKvconfig = cli.StringFlag{
		Name:   "index-kvconfig",
		Usage:  "Index Key-Value config",
		EnvVar: "BLAST_INDEX_KVCONFIG",
	}

	flPeerGRPCAddr = cli.StringFlag{
		Name:   "peer-grpc-addr",
		Usage:  "Peer gRPC address to connect on for join the cluster",
		EnvVar: "BLAST_PEER_GRPC_ADDR",
	}

	flMaxSendMsgSize = cli.IntFlag{
		Name:  "max-send-msg-size",
		Value: math.MaxInt32,
		Usage: "Max size of send message via gRPC",
	}
	flMaxRecvMsgSize = cli.IntFlag{
		Name:  "max-recv-msg-size",
		Value: math.MaxInt32,
		Usage: "Max size of receive message via gRPC",
	}

	flLogLevel = cli.StringFlag{
		Name:   "log-level",
		Value:  "INFO",
		Usage:  "Log level",
		EnvVar: "BLAST_LOG_LEVEL",
	}
	flLogFilename = cli.StringFlag{
		Name:   "log-filename",
		Value:  "",
		Usage:  "Log filename",
		EnvVar: "BLAST_LOG_FILENAME",
	}
	flLogMaxSize = cli.IntFlag{
		Name:   "log-max-size",
		Value:  500,
		Usage:  "Max size of a log file (megabytes)",
		EnvVar: "BLAST_LOG_MAX_SIZE",
	}
	flLogMaxBackups = cli.IntFlag{
		Name:   "log-max-backups",
		Value:  3,
		Usage:  "Max backup count of log files",
		EnvVar: "BLAST_LOG_MAX_BACKUPS",
	}
	flLogMaxAge = cli.IntFlag{
		Name:   "log-max-age",
		Value:  30,
		Usage:  "Max age of a log file (days)",
		EnvVar: "BLAST_LOG_MAX_AGE",
	}
	flLogCompress = cli.BoolFlag{
		Name:   "log-compress",
		Usage:  "Compress a log file",
		EnvVar: "BLAST_LOG_COMPRESS",
	}

	flHTTPAccessLogFilename = cli.StringFlag{
		Name:   "http-access-log-filename",
		Value:  "",
		Usage:  "HTTP access log filename",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_FILENAME",
	}
	flHTTPAccessLogMaxSize = cli.IntFlag{
		Name:   "http-access-log-max-size",
		Value:  500,
		Usage:  "Max size of a HTTP access log file (megabytes)",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_MAX_SIZE",
	}
	flHTTPAccessLogMaxBackups = cli.IntFlag{
		Name:   "http-access-log-max-backups",
		Value:  3,
		Usage:  "Max backup count of HTTP access log files",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_MAX_BACKUPS",
	}
	flHTTPAccessLogMaxAge = cli.IntFlag{
		Name:   "http-access-log-max-age",
		Value:  30,
		Usage:  "Max age of a HTTP access log file (days)",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_MAX_AGE",
	}
	flHTTPAccessLogCompress = cli.BoolFlag{
		Name:   "http-access-log-compress",
		Usage:  "Compress a HTTP access log",
		EnvVar: "BLAST_HTTP_ACCESS_LOG_COMPRESS",
	}

	flBatchSize = cli.IntFlag{
		Name:   "batch-size",
		Value:  1000,
		Usage:  "Batch size for bulk update",
		EnvVar: "BLAST_BATCH_SIZE",
	}

	flPrettyPrint = cli.BoolFlag{
		Name:   "pretty-print",
		Usage:  "Pretty print JSON",
		EnvVar: "BLAST_PRETTY_PRINT",
	}
)