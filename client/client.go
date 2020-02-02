/*
   Copyright 2019 Vadim Inshakov

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"fmt"
	pb "github.com/vadiminshakov/fabex/proto"
	"github.com/vadiminshakov/fabex/client/fabexclient"
	"log"
)

var (
	client *fabexclient.FabexClient
	addr   = "0.0.0.0"
	port   = "6000"
)

func init() {
	var err error
	client, err = fabexclient.New(addr, port)
	if err != nil {
		panic(err)
	}
}


func main() {
	//Explore(1, 15)
	//txs, err := client.GetByTxId(&pb.RequestFilter{Txid:"3a3e933a3d9953b0b10e6573254b6d3cf2347d72058c0347a55054babdd8e1a1"})
	//txs, err := client.GetByBlocknum(&pb.RequestFilter{Blocknum: 2})
	txs, err := client.GetBlockInfoByPayload(&pb.RequestFilter{Payload: "1440-"})
	if err != nil {
		log.Fatal(err)
	}

	blocks, err := client.PackTxsToBlocks(txs)
	fmt.Printf("%#v", blocks)
}
