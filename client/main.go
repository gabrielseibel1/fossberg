/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/gabrielseibel1/fossberg/protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "Newbie"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type point struct{ x, y, z int64 }

var (
	player point
	guntip point
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGameServiceClient(conn)

	username := *name

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Join game
	enterRequest := pb.EnterGameRequest{Username: username}
	spawn, err := c.EnterGame(ctx, &enterRequest)
	if err != nil {
		log.Fatalf("could not enter game: %v", err)
	}
	player.x, player.y, player.z = spawn.GetX(), spawn.GetY(), spawn.GetZ()
	log.Printf("[EnterGame] %s -> %+v", enterRequest.String(), player)

	// Fire a shot
	guntip.x, guntip.y, guntip.z = player.x+1, player.y+2, player.z+3
	fireRequest := pb.FireRequest{Username: username, X1: player.x, Y1: player.y, Z1: player.z, X2: guntip.x, Y2: guntip.y + 2, Z2: guntip.z + 3}
	shot, err := c.Fire(ctx, &fireRequest)
	if err != nil {
		log.Fatalf("could not fire: %v", err)
	}
	log.Printf("[Fire] %+v-%+v -> %s", player, guntip, shot.String())
}
