package cmd

import (
	// local
	"os"
	"os/signal"
	"syscall"

	actv1 "github.com/sqsinformatique/rosseti-back/domains/act/v1"
	actsdetailv1 "github.com/sqsinformatique/rosseti-back/domains/acts_detail/v1"
	categoryv1 "github.com/sqsinformatique/rosseti-back/domains/category/v1"
	defectv1 "github.com/sqsinformatique/rosseti-back/domains/defect/v1"
	elementequipmentv1 "github.com/sqsinformatique/rosseti-back/domains/element_equpment/v1"
	journalv1 "github.com/sqsinformatique/rosseti-back/domains/journal/v1"
	objectv1 "github.com/sqsinformatique/rosseti-back/domains/object/v1"
	objectsdetailv1 "github.com/sqsinformatique/rosseti-back/domains/objects_detail/v1"
	orderv1 "github.com/sqsinformatique/rosseti-back/domains/order/v1"
	profilev1 "github.com/sqsinformatique/rosseti-back/domains/profile/v1"
	reviewv1 "github.com/sqsinformatique/rosseti-back/domains/review/v1"
	sessionv1 "github.com/sqsinformatique/rosseti-back/domains/session/v1"
	techtaskv1 "github.com/sqsinformatique/rosseti-back/domains/tech_task/v1"
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

	ObjectV1, err := objectv1.NewObjectV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ObjectV1")
	}

	TechTaskV1, err := techtaskv1.NewTechTaskV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create TechTaskV1")
	}

	_, err = orderv1.NewOrderV1(ctx, ORM, UserV1, ObjectV1, ProfileV1, TechTaskV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create OrderV1")
	}

	ReviewV1, err := reviewv1.NewReviewV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ReviewV1")
	}

	DefectsV1, err := defectv1.NewDefectV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create DefectsV1")
	}

	ObjectsDetailV1, err := objectsdetailv1.NewObjectsDetailV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ObjectsDetailV1")
	}

	_, err = categoryv1.NewCategoryV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ObjectsDetailV1")
	}

	ActsDetailV1, err := actsdetailv1.NewActsDetailV1(ctx, ORM, UserV1, DefectsV1, ObjectsDetailV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ActsDetailV1")
	}

	ActV1, err := actv1.NewActV1(ctx, ProfileV1, ORM, UserV1, ObjectV1, ReviewV1, ActsDetailV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ActV1")
	}

	CategoryV1, err := categoryv1.NewCategoryV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ActV1")
	}

	ElementEquipmentV1, err := elementequipmentv1.NewElementEquipmentV1(ctx, ORM, UserV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create ElementEquipmentV1")
	}

	_, err = journalv1.NewJournalV1(ctx, ORM, UserV1, ObjectV1, ProfileV1, ActV1, CategoryV1, ElementEquipmentV1)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed create JournalV1")
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
