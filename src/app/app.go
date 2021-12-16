package app

import (
	"context"
	"net/http"
	"time"

	"gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/models"
	rlr "gitlab.nekotal.tech/lachain/crosschain/relayer-smart-contract/src/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// App ...
type App struct {
	logger  *logrus.Logger
	router  *mux.Router
	server  *http.Server
	relayer *rlr.RelayerSRV
}

// NewApp is initializes the app
func NewApp(logger *logrus.Logger, addr string, db *gorm.DB,
	laCfg, posCfg, bscCfg, ethCfg *models.WorkerConfig) *App {
	// create new app
	inst := &App{
		logger:  logger,
		router:  mux.NewRouter(),
		server:  &http.Server{Addr: addr},
		relayer: rlr.CreateNewRelayerSRV(logger, db, laCfg, posCfg, bscCfg, ethCfg),
	}
	// set router
	inst.router = mux.NewRouter()
	inst.setRouters()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	inst.server.Handler = handlers.CORS(headers, methods, origins)(inst.router)

	inst.relayer.Run()

	return inst
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.router.HandleFunc(path, f).Methods("GET")
}

func (a *App) setRouters() {
	a.Get("/", a.Endpoints)
	a.Get("/status", a.StatusHandler)
	a.Get("/status/{destination_chain}/{sender}/{receipt}/{amount}", a.SwapStatusHandler)
	// a.Get("/resend_tx/{id}", a.ResendTxHandler)
	// a.Get("/set_mode/{mode}", a.SetModeHandler)
}

// Run the app on it's router
func (a *App) Run(ctx context.Context) {
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal(err)
		}
	}()

	a.logger.Infof("Relayer service has started on %s\nPress ctrl + C to exit.", a.server.Addr)

	<-ctx.Done()

	a.logger.Infoln("Relayer service has stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctxShutDown); err != nil {
		a.logger.Fatalf("Shutdown: %v\n", err)
	}
}
