package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/filevich/truco-cfr/abs"
	"github.com/filevich/truco-cfr/cfr"
	"github.com/filevich/truco-cfr/eval"
	"github.com/filevich/truco-cfr/eval/dataset"
)

const (
	save_dir = "/tmp"
)

func main() {

	threads := 1
	num_players := 2
	tiny_eval := 1_000

	trainer := cfr.NewTrainer(cfr.ESVMCCFR_T, num_players, &abs.A1{})

	// trainer := cfr.Load(
	// 	cfr.CFR_T,
	// 	"/media/jp/DATA/models/2p/models-24h+48p-multi-core/extension-2d/a2/final_CFR_d24h2m_D24h0m_t8435_p0_a2_2205262321.json")

	// trainer := cfr.Load_model(
	// 	"/media/jp/DATA/models/2p/models-24h+48p-multi-core/extension-2d/a2/final_cfr_d48h4m_D48h0m_t26190_p0_a2_2210141233.model",
	// 	true,
	// 	1_000_000)

	// tiny eval
	log.Println("loading t1k22")
	var ds dataset.Dataset = dataset.LoadDataset("eval/t1k22.json")
	log.Println("done loading t1k22")

	post_save := func() {
		agent := &cfr.BotCFR{
			N:     trainer.String(),
			Model: trainer,
		}
		log.Println("tiny evaluating")
		res := eval.TinyEval(agent, num_players, ds[:tiny_eval])
		log.Printf("%s\n\n", res.String())
		runtime.GC()
	}

	post_save()

	// trainer.Train(
	// 	&cfr.ProfileTime{
	// 		TotalRunningTime:  24 * time.Hour,
	// 		Prunning_treshold: cfr.NEVER,
	// 		// multi
	// 		Threads: threads,
	// 		Mu:      &sync.Mutex{},
	// 		// io
	// 		Save_every:  25 * time.Hour,
	// 		Silent:      true,
	// 		Save_dir:    save_dir,
	// 		Save_prefix: "pre_",
	// 		// tiny eval
	// 		PostSave: post_save,
	// 		// GC
	// 		GC_every: 1 * time.Hour,
	// 	},
	// )

	// log.Printf("Resetting strategy sums")
	// trainer.Reset()

	trainer.Train(
		&cfr.ProfileTime{
			TotalRunningTime: 25 * time.Minute,
			PrunningTreshold: cfr.NEVER,
			// multi
			Threads: threads,
			Mu:      &sync.Mutex{},
			// io
			SaveEvery:  2 * time.Minute,
			Silent:     true,
			SaveDir:    save_dir,
			SavePrefix: "final_",
			// tiny eval
			PostSave:  post_save,
			EvalEvery: 1 * time.Minute,
			// GC
			GCEvery: 100 * time.Hour,
		},
	)

	// trainer.Train(
	// 	&cfr.ProfileTime{
	// 		TotalRunningTime:  4 * 24 * time.Hour,
	// 		Prunning_treshold: time.Nanosecond,
	// 		// multi
	// 		Threads: threads,
	// 		Mu:      &sync.Mutex{},
	// 		// io
	// 		Save_every:  24 * time.Hour,
	// 		Silent:      true,
	// 		Save_dir:    save_dir,
	// 		Save_prefix: "final_",
	// 		// tiny eval
	// 		PostSave: post_save,
	// 		// GC
	// 		GC_every: 1 * time.Hour,
	// 	},
	// )

}
