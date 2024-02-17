package fake

import (
	"context"
	"github.com/constantincuy/knowledgestore/internal/core/domain/common"
	"github.com/constantincuy/knowledgestore/internal/core/domain/file"
	"os"
	"time"
)

type Storage struct {
	files    file.List
	contents map[string]string
}

func (s Storage) Provider() string {
	return "fake_provider_1"
}

func (s Storage) GetChangedDocuments(ctx context.Context, filesystem file.Filesystem) (file.ChangeList, error) {
	return filesystem.Sync(s.files)
}

func (s Storage) DownloadDocument(ctx context.Context, path file.Path, target *os.File) {
	if target != nil {
		target.Write([]byte(s.contents[string(path)]))
	}
}

func NewStorage() Storage {
	files := make(file.List)
	created, _ := file.NewCreated(time.Now().UTC())
	updated, _ := file.NewUpdated(time.Now().UTC())
	files["my/path/hund.txt"] = file.File{
		Id:       common.NewId(),
		Path:     "my/path/hund.txt",
		Provider: "fake_provider_1",
		Created:  created,
		Updated:  updated,
	}

	files["my/path/flugzeug.txt"] = file.File{
		Id:       common.NewId(),
		Path:     "my/path/flugzeug.txt",
		Provider: "fake_provider_1",
		Created:  created,
		Updated:  updated,
	}

	files["my/path/computer.txt"] = file.File{
		Id:       common.NewId(),
		Path:     "my/path/computer.txt",
		Provider: "fake_provider_1",
		Created:  created,
		Updated:  updated,
	}

	contents := make(map[string]string)
	contents["my/path/hund.txt"] = `
Der Haushund (Canis lupus familiaris) ist ein Haustier und wird als Heim- und Nutztier gehalten. Seine wilde Stammform ist der Wolf, dem er als Unterart zugeordnet wird. Wann die Domestizierung stattfand, ist umstritten; wissenschaftliche Schätzungen variieren zwischen etwa 15.000 v. Chr. und 100.000 Jahren.

Im engeren Sinn bezeichnet man als Haushund die Hunde, die überwiegend im Haus gehalten werden, und kennzeichnet damit also eine Haltungsform. Historisch wurde ein Hund, der zur Bewachung des Hauses gehalten wird, als Haushund bezeichnet. Eine weitere Verwendung des Begriffs ist die Einschränkung auf sozialisierte (Haus-)Hunde, also Hunde, die an das Zusammenleben mit Menschen in der menschlichen Gesellschaft gewöhnt und an dieses angepasst sind. Damit wird der Haushund abgegrenzt gegen wild lebende, verwilderte oder streunende Hunde, die zwar auch domestiziert, aber nicht sozialisiert sind.

Der Dingo ist ebenfalls ein Haushund, wird jedoch provisorisch als eigenständige Unterart des Wolfes geführt`

	contents["my/path/flugzeug.txt"] = `
Ein Flugzeug ist ein Luftfahrzeug, das schwerer als Luft ist und den zu seinem Fliegen nötigen dynamischen Auftrieb mit nicht-rotierenden Auftriebsflächen erzeugt. In der enger gefassten Definition der Internationalen Zivilluftfahrtorganisation (ICAO) ist es auch immer ein motorisiertes Luftfahrzeug. Der Betrieb von Flugzeugen, die am Luftverkehr teilnehmen, wird durch Luftverkehrsgesetze geregelt.

Umgangssprachlich werden Flugzeuge mitunter auch Flieger genannt, der Ausdruck Flieger hat als Hauptbedeutung jedoch den Piloten.

## Definition
Die Internationale Zivilluftfahrtorganisation (International Civil Aviation Organization, ICAO) definiert den Begriff Flugzeug wie folgt:

Aeroplane. A power-driven heavier-than-air aircraft, deriving its lift in flight chiefly from aerodynamic reactions on surfaces which remain fixed under given conditions of flight.

– International Civil Aviation Organization[2]
Im rechtlichen Sprachgebrauch ist ein Flugzeug ein motorgetriebenes Luftfahrzeug, schwerer als (die von ihm verdrängte) Luft, das seinen Auftrieb durch Tragflächen erhält, die bei gleichbleibenden Flugbedingungen unverändert bleiben, allgemeinsprachlich Motorflugzeug genannt. Wenn in einem Gesetzestext also von Flugzeugen die Rede ist, dann sind immer nur Motorflugzeuge gemeint, nicht aber Segelflugzeuge, Motorsegler und Ultraleichtflugzeuge. Letztere sind in Deutschland eine Unterklasse der Luftsportgeräte.

Manche Autoren verwenden eine weiter gefasste Definition, nach der auch die Drehflügler eine Untergruppe der Flugzeuge darstellen. Die eigentlichen Flugzeuge werden dann zur besseren Abgrenzung als Starrflügler, Starrflügelflugzeug oder Flächenflugzeug bezeichnet.[3][4] Diese Einordnung widerspricht aber sowohl der rechtlichen Definition als auch dem allgemeinen Sprachgebrauch und kann damit als veraltet betrachtet werden.[5]

Die in diesem Artikel verwendete Definition richtet sich nach der umgangssprachlichen Bedeutung des Begriffes Flugzeug, die sämtliche Luftfahrzeuge umfasst, die einen Rumpf mit festen Tragflächen besitzen.[6][7]

## Abgrenzung zu anderen Luftfahrzeugen

Bei Flugzeugen wird der Auftrieb – bei der Vorwärtsbewegung des Luftfahrzeugs – durch die Umlenkung der notwendigen Luftströmung an den Tragflächen (mit geeignetem Profil und Anstellwinkel) erzeugt. Durch die Umlenkung wird der Luft ein senkrecht nach unten gerichteter Impuls übertragen. Nach dem ersten Newtonschen Gesetz erfordert diese Richtungsänderung der Strömung nach unten eine stetig wirkende Kraft. Nach dem dritten Newtonschen Gesetz (Actio und reactio) wirkt dabei eine gleiche und entgegengesetzte Kraft, der Auftrieb, auf die Tragfläche.[8]

Neben der starren Verbindung von Flügel und Flugzeugrumpf gibt es mit Wandel- und Schwenkflügelflugzeugen auch einige Flugzeugtypen, bei denen die Flügel flexibel am Flugzeugrumpf fixiert sind. Damit können bei diesen Typen Einsatzanforderungen realisiert werden, die mit einer starren Tragfläche nicht möglich sind. Im weiteren Sinn benutzen das Starrflügelprinzip auch Luftfahrzeuge mit vollkommen flexiblen Tragflächen, wie Gleit- und Motorschirme, sowie mit zerlegbaren Tragflächen wie bei Hängegleitern.
`

	contents["my/path/computer.txt"] = `
Ein Computer (englisch; deutsche Aussprache [kɔmˈpjuːtɐ]) oder Rechner ist ein Gerät, das mittels programmierbarer Rechenvorschriften Daten verarbeitet. Dementsprechend werden vereinzelt auch die abstrahierenden beziehungsweise veralteten, synonym gebrauchten Begriffe Rechenanlage, Datenverarbeitungsanlage oder elektronische Datenverarbeitungsanlage sowie Elektronengehirn verwendet.

Charles Babbage und Ada Lovelace (geborene Byron) gelten durch die von Babbage 1837 entworfene Rechenmaschine Analytical Engine als Vordenker des modernen universell programmierbaren Computers. Konrad Zuse (Z3, 1941 und Z4, 1945) in Berlin, John Presper Eckert und John William Mauchly (ENIAC, 1946) bauten die ersten funktionstüchtigen Geräte dieser Art. Bei der Klassifizierung eines Geräts als universell programmierbarer Computer spielt die Turing-Vollständigkeit eine wesentliche Rolle. Sie ist benannt nach dem englischen Mathematiker Alan Turing, der 1936 das logische Modell der Turingmaschine eingeführt hatte.`

	return Storage{files, contents}
}
