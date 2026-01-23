package info

import (
	"encoding/json"
	"fmt"
	"hash"
	"strconv"

	"github.com/filevich/truco-mccfr-ai/abs"
	"github.com/filevich/truco-mccfr-ai/utils"
	"github.com/truquito/gotruco/enco"
	"github.com/truquito/gotruco/pdt"
)

type InfosetRondaBaseFullBoolean struct {
	InfosetRondaBase
	// Used by other "extended" structure
	Nuestros_pts int
	Opp_pts      int
}

func (info *InfosetRondaBaseFullBoolean) setPuntos(p *pdt.Partida, m *pdt.Manojo) {
	our_team := m.Jugador.Equipo
	opp_team := m.Jugador.GetEquipoContrario()
	info.Nuestros_pts = 0
	info.Opp_pts = 0
	const threshold = 5

	if p.Puntajes[our_team] >= int(p.Puntuacion)-threshold {
		info.Nuestros_pts = 1
	}

	if p.Puntajes[opp_team] >= int(p.Puntuacion)-threshold {
		info.Opp_pts = 1
	}
}

func (info *InfosetRondaBaseFullBoolean) HashBytes(h hash.Hash) []byte {
	h.Reset()
	hsep := []byte(sep)

	// 1
	h.Write([]byte(strconv.Itoa(info.Muestra)))
	h.Write(hsep)

	// 2
	// paso de un array de abstracciones a un array de primos
	nuestrasCartas := make([]int, len(info.NuestrasCartas))
	for mix, manojo := range info.NuestrasCartas {
		manojoPrimeID := 1
		for _, abstraccion := range manojo {
			manojoPrimeID *= utils.AllPrimes[abstraccion]
		}
		nuestrasCartas[mix] = manojoPrimeID
	}
	bs, _ := json.Marshal(nuestrasCartas)
	h.Write([]byte(bs))
	h.Write(hsep)

	// 3
	bs, _ = json.Marshal(info.ManojosEnJuego)
	h.Write(bs)
	h.Write(hsep)

	// 4
	h.Write([]byte(info.Envido))
	h.Write(hsep)

	// 5
	h.Write([]byte(info.Truco))
	h.Write(hsep)

	// 6
	bs, _ = json.Marshal(info.ResultadoManos)
	h.Write([]byte(bs))
	h.Write(hsep)

	// 7
	manoActual := fmt.Sprintf("%d.%d.%s",
		info.ManoActual.Max_us, info.ManoActual.Max_op, info.ManoActual.Vamos)
	h.Write([]byte(manoActual))
	h.Write(hsep)

	// 8
	bs, _ = json.Marshal(info.Chi)
	h.Write(bs)
	h.Write(hsep)

	// puntos
	h.Write([]byte(strconv.Itoa(info.Nuestros_pts)))
	h.Write(hsep)
	h.Write([]byte(strconv.Itoa(info.Opp_pts)))
	h.Write(hsep) // <- last separator is not really necessary

	return h.Sum(nil)
}

func infosetRondaBaseFullBooleanFactory(
	a abs.IAbstraction,
) InfosetBuilder {
	return func(
		p *pdt.Partida,
		m *pdt.Manojo,
		msgs []enco.IMessage,
	) Infoset {
		info := &InfosetRondaBaseFullBoolean{
			InfosetRondaBase{
				Vision: m.Jugador.ID,
			},
			0, 0,
		}
		chi_i := pdt.GetA(p, m)
		info.setMuestra(p)
		info.setNuestras_Cartas(p, m, a)
		info.setManojos_en_juego(p, m)
		info.setEnvido(p)
		info.setTruco(p)
		info.setChi(p, m, chi_i, a)
		info.setResultadoManos(p, m)
		info.setRonda(p, m, a)
		info.setPuntos(p, m)
		return info
	}
}
