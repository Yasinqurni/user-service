package initproject

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Yasinqurni/be-project/pkg/database"
	grpccontroller "github.com/Yasinqurni/be-project/src/app/user/grpc"
	"github.com/Yasinqurni/be-project/src/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	grpcserver "github.com/Yasinqurni/be-project/pkg/grpc"
)

type logwriter struct{}

func (writer logwriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func InitProject() {
	log.SetFlags(0)
	log.SetOutput(new(logwriter))

	configEnv := database.InitDB()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	routes.Router(v1, configEnv.DB)

	// gRPC Server
	grpcPort := fmt.Sprintf(":%s", configEnv.PortGRPC)
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	grpcSrv := grpc.NewServer()

	grpccontroller.NewGrpcServer(grpcSrv, configEnv.DB)
	grpcserver.New(
		grpcSrv,
		grpcPort,
		lis,
	)

	r.Run(fmt.Sprintf(":%s", configEnv.PortAPP))

}
