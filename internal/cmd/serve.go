package cmd

import (
	// local
	"os"
	"os/signal"
	"syscall"

	actv1 "github.com/sqsinformatique/rosseti-back/domains/act/v1"
	objectv1 "github.com/sqsinformatique/rosseti-back/domains/object/v1"
	orderv1 "github.com/sqsinformatique/rosseti-back/domains/order/v1"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	sessionv1 "github.com/sqsinformatique/rosseti-back/domains/session/v1"
	userv1 "github.com/sqsinformatique/rosseti-back/domains/user/v1"
	"github.com/sqsinformatique/rosseti-back/internal/cfg"
	"github.com/sqsinformatique/rosseti-back/internal/context"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/internal/httpsrv"
	"github.com/sqsinformatique/rosseti-back/internal/mongo"
	"github.com/sqsinformatique/rosseti-back/internal/orm"

	// other
	"github.com/spf13/cobra"
)

type empty struct{}

func serveHandler(cmd *cobra.Command, args []string) {
	// Create context
	ctx := context.NewContext()

	// Initilize config
	config := cfg.NewConfig()
	ctx.RegisterConfig(config)

	// Register logger
	ctx.RegisterLogger()
	log := ctx.GetPackageLogger(empty{})

	log.Info().Msg(AppInfo)
	log.Info().Msg("Starting Rosseti-Back...")

	// Initialize connection
	DB, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create DB")
	}

	_, err = mongo.NewMongoDB(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create DB")
	}

	// Initilize ORM
	ORM, err := orm.NewORM("production", ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ORM")
	}

	// Initialize web-server, public/private-endpoints
	HTTPSrv, err := httpsrv.NewHTTPSrv(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create HTTPSrv")
	}

	// Initialize health/metrics
	// _ = health.Initialize(ctx, httpsrv.ProviderName, "private")
	// _ = metrics.Initialize(ctx)

	// Initilize domains
	SessionV1, err := sessionv1.NewSessionV1(ctx, ORM)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create SessionV1")
	}

	UserV1, err := userv1.NewUserV1(ctx, ORM, SessionV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create UserV1")
	}

	ProfileV1, err := profilev1.NewProfileV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ProfileV1")
	}

	_, err = actv1.NewActV1(ctx, ProfileV1, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ActV1")
	}

	ObjectV1, err := objectv1.NewObjectV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ObjectV1")
	}

	_, err = orderv1.NewOrderV1(ctx, ORM, UserV1, ObjectV1, ProfileV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create OrderV1")
	}

	// Start connect
	if err := DB.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed connect to DB")
	}

	// Start web-server, public/private-endpoints
	HTTPSrv.Start()

	// Start metrics
	// _ = health.Start(ctx)

	var closeSignal chan os.Signal

	exit := make(chan struct{})
	closeSignal = make(chan os.Signal, 1)
	signal.Notify(closeSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-closeSignal
		_ = HTTPSrv.Shutdown()
		log.Info().Msg("Exit program")
		close(exit)
	}()

	// Exit app if chan is closed
	<-exit
}
