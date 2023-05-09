package usecase

import (
	"encoding/json"
	"gin/database/mysql"
	"gin/sdk/password"
	"gin/src/entity"
	"gin/src/enum"
)

var nd = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/nd1.png",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/nd2.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/nd3.jpeg",
}

var btr = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btr1.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btr2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btr3.jpg",
}

var btkr = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btkr1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btkr2.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btkr3.jpg",
}

var ag = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ag1.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ag2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ag3.jpg",
}

var kt = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kt1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kt2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kt3.jpeg",
}

var tgng = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/tg1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/tg2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/tg3.jpeg",
}

var skmp = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/sekumpul1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/sekumpul2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/sekumpul3.jpeg",
}

var pdw = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/pdw1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/pdw2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/pdw3.jpeg",
}

var brt = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/brt1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/brt2.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/brt3.jpg",
}

var klkng = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/klk1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/klk2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/klk3.jpg",
}

var ot = []string{
	"09.30 - 16.00",
	"09.30 - 16.00",
	"09.30 - 16.00",
	"09.30 - 16.00",
	"09.30 - 16.00",
	"08.30 - 18.00",
	"08.30 - 18.00",
}

func (uc *Usecase) GenerateEcotourismDummy(db *mysql.DB) {
	nd, _ := json.Marshal(nd)
	btr, _ := json.Marshal(btr)
	btkr, _ := json.Marshal(btkr)
	ag, _ := json.Marshal(ag)
	kt, _ := json.Marshal(kt)
	tgng, _ := json.Marshal(tgng)
	skmp, _ := json.Marshal(skmp)
	pdw, _ := json.Marshal(pdw)
	brt, _ := json.Marshal(brt)
	klkng, _ := json.Marshal(klkng)

	ot, _ := json.Marshal(ot)

	var ecotourisms = []entity.Ecotourism{
		{
			ID:              100,
			Thumbnail:       nd,
			Name:            "Nusa Dua",
			Region:          "Bali",
			TotalRatings:    8,
			Rating:          8.6,
			Description:     "Nusa Dua adalah sebuah pantai yang terletak di Bali, Indonesia. Pantai ini terkenal sebagai salah satu pantai yang sangat eksklusif dan mewah di Bali. Nusa Dua memiliki pasir putih yang halus dan air laut yang jernih dan tenang, sehingga pantai ini sangat cocok untuk berenang atau hanya sekedar bersantai di tepi pantai. Selain itu, pantai ini juga dikelilingi oleh hotel-hotel bintang lima dan tempat-tempat wisata yang menarik.",
			Price:           10000,
			Category:        enum.Pantai,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3942.8959373609796!2d115.23121217607502!3d-8.795846639920136!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd243213c2ea8df%3A0xfe9b8fbae7a12e8f!2sNusa%20Dua%20Beach!5e0!3m2!1sen!2sid!4v1683649529093!5m2!1sen!2sid",
		},
		{
			ID:              101,
			Thumbnail:       btr,
			Name:            "Batur",
			Region:          "Bali",
			TotalRatings:    2,
			Rating:          7.1,
			Description:     "Gunung Batur adalah salah satu gunung berapi yang terletak di Bali, Indonesia, dengan ketinggian sekitar 1.717 meter di atas permukaan laut. Menyajikan pemandangan yang sangat menakjubkan, pendakian ke puncak Gunung Batur adalah salah satu aktivitas yang populer di Bali, yang menarik wisatawan dari seluruh dunia. Dengan jalur pendakian yang cukup menantang namun tetap aman, pendaki akan disuguhkan dengan panorama indah dari keindahan alam Bali di sekitarnya, seperti Danau Batur yang sangat terkenal dan panorama pegunungan serta hijaunya hutan sekitarnya. Saat tiba di puncak, pengunjung dapat menikmati keindahan matahari terbit yang luar biasa sambil menikmati sarapan pagi. Pendakian Gunung Batur adalah pengalaman wisata yang sangat menarik bagi para pecinta alam dan petualangan, dan pastinya akan menjadi salah satu pengalaman tak terlupakan selama berada di Bali.",
			Price:           150000,
			Category:        enum.Gunung,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d31589.025837599722!2d115.35689482461532!3d-8.240086752017737!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd1f403c8e8ee3f%3A0xd38045afa18670b4!2sMt%20Batur!5e0!3m2!1sen!2sid!4v1683649573901!5m2!1sen!2sid",
		},
		{
			ID:              103,
			Thumbnail:       btkr,
			Name:            "Batukaru",
			Region:          "Bali",
			TotalRatings:    5,
			Rating:          7.9,
			Description:     "Gunung Batukaru adalah salah satu gunung yang terletak di Bali dan merupakan bagian dari Pegunungan Bali Barat. Dengan ketinggian mencapai 2.276 meter di atas permukaan laut, Gunung Batukaru menawarkan pemandangan alam yang indah dan memukau. Di sekitar gunung ini, terdapat hutan tropis yang lebat dan beragam, seperti hutan hujan tropis, hutan montane, dan hutan bambu yang hijau. Selain itu, terdapat pula berbagai jenis flora dan fauna langka yang hidup di sekitar gunung ini, seperti rusa, burung jalak Bali, dan bunga anggrek endemik Bali. Untuk para pendaki, Gunung Batukaru menawarkan jalur pendakian yang menantang namun sepadan dengan pemandangan yang spektakuler dari puncaknya. Dengan segala keindahan alam yang ditawarkannya, Gunung Batukaru patut menjadi destinasi wisata alam yang wajib dikunjungi bagi siapa saja yang berkunjung ke Bali.",
			Price:           20000,
			Category:        enum.Gunung,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d15790.713439521134!2d115.08777779999998!3d-8.33472205!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd187f8b41bd26f%3A0x688908b6268505d3!2sBukit%20Batukau!5e0!3m2!1sen!2sid!4v1683649605905!5m2!1sen!2sid",
		},
		{
			ID:              104,
			Thumbnail:       ag,
			Name:            "Agung",
			Region:          "Bali",
			TotalRatings:    9,
			Rating:          8.9,
			Description:     "Gunung Agung adalah gunung berapi yang terletak di pulau Bali, Indonesia dan merupakan salah satu gunung tertinggi di Indonesia dengan ketinggian 3.142 meter di atas permukaan laut. Selain menjadi ikon pariwisata Bali, Gunung Agung juga memiliki nilai religius yang penting bagi masyarakat Hindu Bali karena dianggap sebagai tempat suci. Dari puncaknya, para pendaki dapat menikmati pemandangan yang spektakuler dari atas awan serta menyaksikan matahari terbit yang menakjubkan di pagi hari. Namun, pendakian ke Gunung Agung memerlukan persiapan yang matang dan kondisi fisik yang baik karena jalur pendakian yang terjal dan medan yang sulit.",
			Price:           150000,
			Category:        enum.Gunung,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d31580.736468524676!2d115.4864579246991!3d-8.343267406772016!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd202e428b2eac7%3A0xa7d7d26cb3a3a7ad!2sMount%20Agung!5e0!3m2!1sen!2sid!4v1683649749595!5m2!1sen!2sid",
		},
		{
			ID:              105,
			Thumbnail:       kt,
			Name:            "Kuta",
			Region:          "Bali",
			TotalRatings:    7,
			Rating:          8.3,
			Description:     "Pantai Kuta adalah salah satu destinasi wisata alam yang terkenal di Bali, Indonesia. Dengan pasir putih yang lembut, ombak yang tinggi dan panjang garis pantai yang indah, Pantai Kuta menjadi daya tarik bagi para pengunjung yang mencari keindahan alam dan suasana pantai yang menyenangkan. Terletak di daerah Badung, Pantai Kuta juga dikelilingi oleh banyak restoran, kafe, dan toko-toko, sehingga membuatnya menjadi tempat yang populer bagi para wisatawan yang ingin bersantai sambil menikmati suasana Bali yang tenang dan damai.",
			Price:           0,
			Category:        enum.Pantai,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d7887.441994273812!2d115.1629545456012!3d-8.718021910420775!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd246bc2ab70d43%3A0x82feaae12f4ab48e!2sKuta%20Beach!5e0!3m2!1sen!2sid!4v1683649848668!5m2!1sen!2sid",
		},
		{
			ID:              106,
			Thumbnail:       tgng,
			Name:            "Tegenungan",
			Region:          "Bali",
			TotalRatings:    3,
			Rating:          5.9,
			Description:     "Air Terjun Tegenungan terletak di daerah Kemenuh, Gianyar, Bali, dan merupakan salah satu destinasi wisata alam yang populer di pulau Bali. Air terjun setinggi sekitar 15 meter ini dikelilingi oleh hutan tropis hijau dan sungai yang mengalir dengan air yang jernih dan segar. Para wisatawan dapat menikmati keindahan alam yang menakjubkan, berenang di kolam alami, atau bersantai di area sekitar air terjun yang asri dan tenang.",
			Price:           15000,
			Category:        enum.AirTerjun,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3945.214567433657!2d115.28638497607214!3d-8.575356787015146!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd2161beebd0c61%3A0xc1ae79ddb8410c5e!2sTegenungan%20Waterfall!5e0!3m2!1sen!2sid!4v1683650013129!5m2!1sen!2sid",
		},
		{
			ID:              107,
			Thumbnail:       skmp,
			Name:            "Sekumpul",
			Region:          "Bali",
			TotalRatings:    10,
			Rating:          9.8,
			Description:     "Air Terjun Sekumpul terletak di Desa Sekumpul, Kecamatan Sawan, Kabupaten Buleleng, Bali, dan merupakan salah satu air terjun terbesar dan terindah di Bali. Terdapat tujuh air terjun yang membentuk Sekumpul, masing-masing dengan ketinggian yang berbeda, dan dikelilingi oleh tebing-tebing curam serta hutan yang lebat. Pengunjung dapat menikmati pemandangan yang spektakuler dari atas tebing atau mengeksplorasi keindahan alam yang menakjubkan dengan cara trekking melalui jalan setapak yang cukup menantang.",
			Price:           20000,
			Category:        enum.AirTerjun,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3949.2502529437893!2d115.18000757606696!3d-8.177525481973687!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd18f2016762a37%3A0x18e8b8bb069da935!2sSekumpul%20Waterfall!5e0!3m2!1sen!2sid!4v1683650125102!5m2!1sen!2sid",
		},
		{
			ID:              108,
			Thumbnail:       pdw,
			Name:            "Pandawa",
			Region:          "Bali",
			TotalRatings:    7,
			Rating:          8.8,
			Description:     "Pantai Pandawa merupakan sebuah pantai yang terletak di Kabupaten Badung, Bali. Pantai ini terkenal dengan keindahan alamnya yang memukau, terdiri dari pasir putih yang lembut dan air laut berwarna biru jernih. Selain itu, di sepanjang pantai terdapat tebing-tebing batu karang yang menjulang tinggi dan terdapat pula pepohonan rindang yang memberikan nuansa yang asri dan menenangkan. Tidak hanya sebagai tempat wisata, Pantai Pandawa juga menjadi lokasi kegiatan olahraga air seperti selancar dan menyelam yang sangat populer di kalangan pengunjung.",
			Price:           15000,
			Category:        enum.Pantai,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d15769.493489522265!2d115.17622576502252!3d-8.84478675673298!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd25b7cd8ba1f31%3A0x41b8785dd055b2a4!2sPandawa%20Beach!5e0!3m2!1sen!2sid!4v1683650548263!5m2!1sen!2sid",
		},
		{
			ID:              109,
			Thumbnail:       brt,
			Name:            "Beratan",
			Region:          "Bali",
			TotalRatings:    6,
			Rating:          7.9,
			Description:     "Danau Beratan adalah sebuah danau yang terletak di kawasan pegunungan Bedugul, Bali. Danau ini memiliki pemandangan alam yang indah dengan air jernih yang tenang, di sekitarnya terdapat juga sebuah pura suci yang menjadi destinasi populer bagi para wisatawan. Dengan udara yang sejuk dan udara pegunungan yang segar, Danau Beratan menjadi tempat yang ideal untuk bersantai dan menikmati keindahan alam Bali.",
			Price:           20000,
			Category:        enum.Danau,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d15793.18333742596!2d115.16487746490454!3d-8.273325483705193!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd18968c2459d53%3A0xbe028c1761c867ac!2sDanau%20Beratan!5e0!3m2!1sen!2sid!4v1683650676729!5m2!1sen!2sid",
		},
		{
			ID:              110,
			Thumbnail:       klkng,
			Name:            "Kelingking",
			Region:          "Bali",
			TotalRatings:    4,
			Rating:          5.6,
			Description:     "Pantai Kelingking adalah salah satu destinasi wisata alam yang paling populer di Pulau Nusa Penida, Bali. Terkenal dengan formasi tebingnya yang mirip dengan rahang hiu yang membentang di sepanjang pantai, pantai Kelingking menawarkan pemandangan spektakuler dari laut biru turkis dan pasir putih yang bersih. Selain itu, pengunjung juga dapat menikmati kegiatan menyelam atau snorkeling di sekitar pantai yang kaya akan keanekaragaman hayati laut.",
			Price:           5000,
			Category:        enum.Pantai,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d63094.060275702046!2d115.40052402167966!3d-8.750386499999992!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd26f1616cd2cc7%3A0xee84b7df0afbff19!2sKelingking%20Beach%20Nusa%20Penida!5e0!3m2!1sen!2sid!4v1683650770874!5m2!1sen!2sid",
		},
	}

	for _, ecotourism := range ecotourisms {
		db.Model(&entity.Ecotourism{}).Create(&ecotourism)
	}
}

func (uc *Usecase) GenerateCommentDummy(db *mysql.DB) {
	nd, _ := json.Marshal(nd)
	btr, _ := json.Marshal(btr)
	var comments = []entity.Comment{
		{
			ID:           1,
			Thumbnail:    nd,
			Date:         "Selasa, 15 Desember 2022",
			UserID:       1,
			Rating:       8.5,
			EcotourismID: 100,
			Body:         "Tempatnya bagus banget, pemandangannya indah, dan cocok buat liburan keluarga",
		},
		{
			ID:           2,
			Thumbnail:    btr,
			Date:         "Rabu, 10 Mei 2022",
			UserID:       2,
			Rating:       7,
			EcotourismID: 101,
			Body:         "Pendakiannya cukup menantang, tapi pemandangannya sangat indah",
		},
		{
			ID:           3,
			Thumbnail:    btr,
			Date:         "Sabtu, 20 Januari 2020",
			UserID:       1,
			Rating:       8,
			EcotourismID: 101,
			Body:         "Sungguh pemandangangan yang sangat indah, Aku tak akan pernah melupakannya",
		},
	}

	for _, comment := range comments {
		db.Model(&entity.Comment{}).Create(&comment)
	}
}

func (uc *Usecase) GenerateUserDummy(db *mysql.DB) {
	kzh, _ := password.GeneratePassword("kazuha444")
	kai, _ := password.GeneratePassword("kai444")

	var users = []entity.User{
		{
			ID:           1,
			Username:     "kai",
			Password:     kai,
			Email:        "kai@gmail.com",
			Wallet:       100000,
			Contact:      "087654345678",
			Region:       "Bali",
			Gender:       enum.Pria,
			Birthday:     "15 Desember 2003",
			PhotoProfile: "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/Cr1.jpg",
		},
		{
			ID:           2,
			Username:     "kazuha",
			Password:     kzh,
			Email:        "kazuha@gmail.com",
			Wallet:       50000,
			Contact:      "089678765456",
			Region:       "Malang",
			Gender:       enum.Wanita,
			Birthday:     "10 November 2002",
			PhotoProfile: "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/anime-spy-x-family-anya-forger_169.jpeg",
		},
	}

	var trashes = []entity.Trash{
		{
			ID:            1,
			Category:      enum.Kertas,
			Location:      "Pantai Kuta, Bali",
			Mass:          10,
			Code:          "F123HGBH",
			Status:        enum.Berhasil,
			ExchangeTotal: 10000,
			UserID:        1,
		},
		{
			ID:            2,
			Category:      enum.Kaca,
			Location:      "-",
			Mass:          5,
			Code:          "F123HGBH",
			Status:        enum.Menunggu,
			ExchangeTotal: 25000,
			UserID:        1,
		},
		{
			ID:            3,
			Category:      enum.Kaca,
			Location:      "-",
			Mass:          5,
			Code:          "F123HGBH",
			Status:        enum.Gagal,
			ExchangeTotal: 25000,
			UserID:        1,
		},
	}

	var purchases = []entity.Purchase{
		{
			ID:          1,
			Date:        "Selasa, 15 Desember 2022",
			Place:       "Pantai Nusa Dua, Bali",
			Quantity:    1,
			TotalPrice:  10000,
			Code:        "G15HIHJ7",
			Status:      enum.Berhasil,
			PayCategory: enum.Online,
			UserID:      1,
		},
		{
			ID:          2,
			Date:        "Senin, 17 Agustus 2022",
			Place:       "Gunung Agung, Bali",
			Quantity:    1,
			TotalPrice:  150000,
			Code:        "-",
			Status:      enum.Menunggu,
			PayCategory: enum.Coin,
			UserID:      1,
		},
	}

	for _, user := range users {
		db.Model(&entity.User{}).Create(&user)
	}

	for _, trash := range trashes {
		db.Model(&entity.Trash{}).Create(&trash)
	}

	for _, purchase := range purchases {
		db.Model(&entity.Purchase{}).Create(&purchase)
	}
}
