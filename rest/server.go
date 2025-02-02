// Copyright 2019 GitBitEx.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	r.GET("/api/configs", GetConfigs)
	r.POST("/api/users", SignUp)
	r.POST("/api/users/accessToken", SignIn)
	r.GET("/api/products", GetProducts)
	r.GET("/api/products/:productId/trades", GetProductTrades)
	r.GET("/api/products/:productId/book", GetProductOrderBook)
	r.GET("/api/products/:productId/candles", GetProductCandles)

	private := r.Group("/", checkToken())
	{
		private.GET("/api/orders", GetOrders)
		private.POST("/api/orders", PlaceOrder)
		private.DELETE("/api/orders/:orderId", CancelOrder)
		private.DELETE("/api/orders", CancelOrders)
		private.GET("/api/accounts", GetAccounts)
		private.GET("/api/users/self", GetUsersSelf)
		private.DELETE("/api/users/accessToken", SignOut)

	}

	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}
