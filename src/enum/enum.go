package enum

type EcoCategory string
type Gender string
type PayCategory string
type Status string
type TrashCategory string

const (
	Pantai    EcoCategory = "Pantai"
	AirTerjun EcoCategory = "Air Terjun"
	Gunung    EcoCategory = "Gunung"
	Danau     EcoCategory = "Danau"
	Pulau     EcoCategory = "Pulau"
	Lainnya   EcoCategory = "Lainnya"
)

const (
	Pria   Gender = "Pria"
	Wanita Gender = "Wanita"
)

const (
	Online PayCategory = "Pembayaran Online"
	Coin   PayCategory = "Pembayaran Coin"
)

const (
	Menunggu Status = "Menunggu"
	Berhasil Status = "Berhasil"
	Gagal    Status = "Gagal"
)

const (
	Plastik    TrashCategory = "Plastik"
	Kertas     TrashCategory = "Kertas"
	Elektronik TrashCategory = "Elektronik"
	Kaca       TrashCategory = "Kaca"
	Metal      TrashCategory = "Metal"
	Kardus     TrashCategory = "Kardus"
	Organik    TrashCategory = "Organik"
)
