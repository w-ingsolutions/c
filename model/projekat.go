package model

type WingProjekat struct {
	Id             int    `json:"id"`
	Naziv          string `json:"naziv"`
	Opis           string `json:"opis"`
	IdProjekta     string `json:"id_projekta"`
	BrojDokumenta  string `json:"broj_dokumenta"`
	Sveska         string `json:"sveska"`
	VrstaDokumenta string `json:"vrsta_dokumenta"`
	DatumDokumenta string `json:"datum"`
	Objekti        []*WingObjekat
	Investitori    []*WingLice
	Projektant     []*WingLice
	Radovi         []*WingIzabraniElement
	RabatRadovi    int
	RabatMaterijal int
	CenaRadova     float64
	CenaMaterijala float64

	//OdgovorniProjektant *WingOdgovorniProjektant
	//Dokumntacija        WingDokumentacija
}

type WingObjekat struct {
	BrojObjekta          string `json:"broj_objekta"`
	KategorijaObjekta    string `json:"kategorija_objekta"`
	KlasifikacijaObjekta string `json:"klasifikacija_objekta"`
	Funkcija             string `json:"funkcija"`
	Gradjenje            string `json:"gradjenje"`
	Spratnost            string `json:"spratnost"`
	Lokacija             string `json:"lokacija"`
	Ulica                string `json:"Ulica"`
	Broj                 string `json:"broj"`
	KP                   string `json:"KP"`
	KO                   string `json:"KO"`
}

type WingLice struct {
	Id             int
	Projektant     bool
	Admin          bool
	KratakNaziv    string
	DugiNaziv      string
	Ime            string
	Prezime        string
	JMBG           string
	BrojLicence    string
	Adresa         string
	Grad           string
	Region         string
	PIB            string
	MB             string
	DatumOsnivanja string
	Delatnost      string
	Racuni         []WingBankaRacun
	Email          string
	BrojTelefona   string
}
type WingBankaRacun struct {
	Banka string
	Racun string
}

type WingDokumentacija struct {
	Tekstualna string
	Numericka  string
	Graficka   string
}
