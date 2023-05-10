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

var dp = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/dp1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/dp2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/dp3.jpg",
}

var ps = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ps1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ps2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ps3.jpg",
}

var kp = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kp1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kp2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kp3.jpg",
}

var tnbt = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/tnbt1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/tnbt2.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/tnbt3.jpg",
}

var klyr = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/klyr1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/kly2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/klyr3.jpg",
}

var sp = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/sp1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/sp2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/sp3.jpeg",
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
	dp, _ := json.Marshal(dp)
	ps, _ := json.Marshal(ps)
	kp, _ := json.Marshal(kp)
	tnbt, _ := json.Marshal(tnbt)
	klyr, _ := json.Marshal(klyr)
	sp, _ := json.Marshal(sp)

	ot, _ := json.Marshal(ot)

	var ecotourisms = []entity.Ecotourism{
		{
			ID:              1,
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
			ID:              2,
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
			ID:              3,
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
			ID:              4,
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
			ID:              5,
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
			ID:              6,
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
			ID:              7,
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
			ID:              8,
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
			ID:              9,
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
			ID:              10,
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
		{
			ID:              11,
			Thumbnail:       dp,
			Name:            "Dieng Plateau",
			Region:          "Jawa Tengah",
			TotalRatings:    9,
			Rating:          8.5,
			Description:     "Dieng Plateau adalah sebuah dataran tinggi yang terletak di Provinsi Jawa Tengah, Indonesia. Dieng Plateau dikenal karena keindahan alamnya yang menakjubkan, seperti kawah gunung berapi yang masih aktif, danau berwarna hijau kebiruan, dan pemandangan indah dari puncak gunung. Selain itu, Dieng Plateau juga memiliki banyak situs purbakala yang menarik untuk dikunjungi.",
			Price:           5000,
			Category:        enum.Gunung,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3958.2235632089855!2d109.89652366183016!3d-7.215319892760363!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e700ce34714804d%3A0xf71779c2d538ad1!2sDieng%20Plateau!5e0!3m2!1sen!2sid!4v1683699425484!5m2!1sen!2sid",
		},
		{
			ID:              12,
			Thumbnail:       ps,
			Name:            "Sempu",
			Region:          "Jawa Timur",
			TotalRatings:    8,
			Rating:          8.2,
			Description:     "Pulau Sempu adalah sebuah pulau yang terletak di Jawa Timur, Indonesia. Pulau ini terkenal karena keindahan alamnya yang masih sangat alami, seperti hutan tropis yang lebat, pantai yang indah, dan air laut yang jernih. Pulau Sempu juga merupakan tempat yang ideal untuk melakukan kegiatan seperti trekking, camping, dan snorkeling.",
			Price:           10000,
			Category:        enum.Pulau,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d15786.16525052208!2d112.68489276025714!3d-8.446621364930714!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd60120edbc901f%3A0x8efd89687a308993!2sSempu%20Island!5e0!3m2!1sen!2sid!4v1683699580865!5m2!1sen!2sid",
		},
		{
			ID:              13,
			Thumbnail:       kp,
			Name:            "Kawah Putih",
			Region:          "Jawa Barat",
			TotalRatings:    9,
			Rating:          9.2,
			Description:     "Kawah Putih adalah sebuah kawah gunung berapi yang terletak di Jawa Barat, Indonesia. Kawah Putih terkenal karena keindahan alamnya yang sangat menakjubkan, seperti danau berwarna putih kehijauan yang terbentuk dari gas vulkanik dan tanah vulkanik yang putih. Selain itu, Kawah Putih juga memiliki banyak aktivitas seperti hiking dan fotografi.",
			Price:           15000,
			Category:        enum.Gunung,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d1979.3260578735392!2d107.40084950668947!3d-7.166154048202182!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e688c1383dc510f%3A0xfab41bb8e4a3a83e!2sWhite%20Crater!5e0!3m2!1sen!2sid!4v1683699662292!5m2!1sen!2sid",
		},
		{
			ID:              14,
			Thumbnail:       tnbt,
			Name:            "Taman Nasional Bromo Tengger Semeru",
			Region:          "Jawa Timur",
			TotalRatings:    9,
			Rating:          9.2,
			Description:     "Taman Nasional Bromo Tengger Semeru adalah salah satu taman nasional yang terletak di Provinsi Jawa Timur, Indonesia. Taman nasional ini terkenal karena memiliki pemandangan gunung berapi yang indah, seperti Gunung Bromo, Gunung Tengger, dan Gunung Semeru. Selain itu, terdapat juga padang rumput luas yang hijau, dan banyak tempat wisata menarik yang dapat dikunjungi di sekitar taman nasional ini. Taman Nasional Bromo Tengger Semeru sangat cocok bagi para pecinta alam dan pendaki gunung yang ingin menikmati keindahan alam Jawa Timur.",
			Price:           50000,
			Category:        enum.Gunung,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3950.7773987379564!2d112.94985811183597!3d-8.02187469197122!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2dd63f365f69c945%3A0x36aec022815992c9!2sBromo%20Tengger%20Semeru%20National%20Park!5e0!3m2!1sen!2sid!4v1683700520813!5m2!1sen!2sid",
		},
		{
			ID:              15,
			Thumbnail:       klyr,
			Name:            "Klayar",
			Region:          "Jawa Tengah",
			TotalRatings:    8,
			Rating:          8.,
			Description:     "Pantai Klayar adalah sebuah pantai yang terletak di daerah Pacitan, Jawa Tengah. Pantai ini terkenal karena memiliki pemandangan yang indah dengan batu-batu karang yang unik dan air laut yang biru kehijauan. Di sekitar pantai ini terdapat banyak warung makan dan tempat penginapan yang murah meriah, sehingga pantai ini sangat cocok bagi para backpacker yang ingin menikmati indahnya pantai.",
			Price:           15000,
			Category:        enum.Pantai,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d15795.171332347978!2d110.9356009102288!3d-8.223577418397493!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e7bdd013dbec5f7%3A0x8c6f80dbac239271!2sPantai%20Klayar%20Pacitan!5e0!3m2!1sen!2sid!4v1683700555608!5m2!1sen!2sid",
		},
		{
			ID:              16,
			Thumbnail:       sp,
			Name:            "Situ Patenggang",
			Region:          "Jawa Barat",
			TotalRatings:    8,
			Rating:          8.3,
			Description:     "Situ Patenggang adalah sebuah danau yang terletak di kawasan wisata Ciwidey, Jawa Barat. Danau ini terkenal karena keindahan alamnya yang masih asri dan sejuk. Selain itu, di sekitar danau terdapat pepohonan pinus yang rindang dan padang rumput yang hijau yang sangat cocok untuk berfoto atau sekedar bersantai menikmati udara segar. Di Situ Patenggang, pengunjung juga dapat melakukan berbagai aktivitas seperti perahu dayung, memancing, berkuda, dan bersepeda.",
			Price:           20000,
			Category:        enum.Danau,
			OperationalTime: ot,
			Maps:            "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3958.642222454905!2d107.35509266182981!3d-7.167292842807568!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e688bc0fb5ad637%3A0x2585812940b0a366!2sPatenggang%20Lake!5e0!3m2!1sen!2sid!4v1683700586520!5m2!1sen!2sid",
		},
	}

	for _, ecotourism := range ecotourisms {
		db.Debug().Create(&ecotourism)
	}
}

var ndr = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ndr1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ndr2.jpeg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/ndr3.jpg",
}

var btrr = []string{
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btrr1.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btrr2.jpg",
	"https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/btrr3.jpg",
}

func (uc *Usecase) GenerateCommentDummy(db *mysql.DB) {
	ndr, _ := json.Marshal(ndr)
	btrr, _ := json.Marshal(btrr)
	var comments = []entity.Comment{
		{
			ID:           1,
			Thumbnail:    ndr,
			Date:         "Selasa, 15 Desember 2022",
			UserID:       1,
			Rating:       8,
			EcotourismID: 1,
			Body:         "Pantai Nusa Dua memiliki keindahan alam yang menakjubkan, pasir putih yang lembut, dan air laut yang jernih. Namun, sayangnya pantai ini terlalu ramai dan terlalu komersial. Pengunjung akan merasa sangat terganggu oleh penjual dan tawaran mereka yang berkeliling di sepanjang pantai.",
		},
		{
			ID:           2,
			Date:         "Rabu, 10 Mei 2022",
			UserID:       2,
			Rating:       7,
			EcotourismID: 1,
			Body:         "Pantai Nusa Dua memiliki pemandangan yang sangat indah, air laut yang jernih, dan pasir putih yang lembut. Selain itu, pantai ini terkenal sebagai pantai yang sangat tenang dan aman, sehingga cocok bagi keluarga yang ingin berlibur bersama.",
		},
		{
			ID:           3,
			Date:         "Sabtu, 20 Januari 2020",
			UserID:       3,
			Rating:       7,
			EcotourismID: 1,
			Body:         "Pantai Nusa Dua adalah salah satu pantai yang paling populer di Bali, dan pantai ini memiliki segala yang diinginkan oleh para wisatawan. Pantai ini memiliki pasir putih yang lembut, air laut yang jernih, dan banyak aktivitas yang dapat dilakukan, seperti snorkeling, parasailing, dan surfing.",
		},
		{
			ID:           4,
			Thumbnail:    btrr,
			Date:         "Jumat, 24 Desember 2021",
			UserID:       1,
			Rating:       9,
			EcotourismID: 2,
			Body:         "Gunung Batur adalah gunung berapi yang terletak di Bali, Indonesia, dan memiliki rating 9/10. Saya sangat terkesan dengan pemandangan yang luar biasa dari puncaknya. Pemandangan panorama danau Batur yang indah, serta matahari terbit di pagi hari benar-benar luar biasa.",
		},
		{
			ID:           5,
			Date:         "Selasa, 1 Februari 2022",
			UserID:       2,
			Rating:       7,
			EcotourismID: 2,
			Body:         "Gunung Batur memiliki rating 7/10 dari saya. Saya suka bahwa pendakian ke puncaknya relatif mudah dan dapat dilakukan oleh pendaki dengan tingkat kemampuan yang berbeda. Namun, saya merasa sedikit kecewa dengan pemandangan dari puncaknya karena kabut yang sering terjadi. Selain itu, banyaknya pengunjung bisa membuat jalur pendakian menjadi padat dan terkadang tidak nyaman.",
		},
		{
			ID:           6,
			Date:         "Jumat, 11 Maret 2022",
			UserID:       3,
			Rating:       8,
			EcotourismID: 2,
			Body:         "Saya memberi Gunung Batur rating 8/10. Saya menikmati pendakian yang menantang tapi menyenangkan dan pemandangan yang sangat memukau dari puncaknya. Namun, jalur pendakian bisa sangat berdebu dan kadang-kadang licin, terutama jika turun saat cuaca hujan",
		},
		{
			ID:           7,
			Date:         "Senin, 18 April 2022",
			UserID:       1,
			Rating:       9,
			EcotourismID: 3,
			Body:         "Saya sangat merekomendasikan untuk mendaki Gunung Batukaru. Pemandangan di sepanjang jalur pendakian sangat indah dan segar. Selain itu, ada banyak flora dan fauna yang dapat ditemukan selama perjalanan. Namun, perjalanan menuju puncak cukup sulit dan membutuhkan kekuatan fisik yang cukup. Oleh karena itu, saya hanya merekomendasikan pendakian ini untuk orang yang memiliki pengalaman mendaki yang cukup.",
		},
		{
			ID:           8,
			Date:         "Rabu, 25 Mei 2022",
			UserID:       1,
			Rating:       10,
			EcotourismID: 4,
			Body:         "Mendaki Gunung Agung adalah pengalaman yang sangat menantang dan menyenangkan. Selama perjalanan, kamu akan disuguhkan pemandangan alam yang indah dan beragam, seperti hutan tropis, sungai, dan panorama pegunungan. Puncak gunung ini menawarkan pemandangan yang luar biasa dan terkadang kamu bisa melihat sunrise dan sunset di sana.",
		},
		{
			ID:           9,
			Date:         "Jumat, 1 Juli 2022",
			UserID:       1,
			Rating:       9,
			EcotourismID: 5,
			Body:         "Pantai Kuta adalah salah satu pantai paling terkenal di Bali. Dengan pasir putih yang luas dan air laut yang jernih, ini adalah tempat yang sempurna untuk bersantai dan menikmati pemandangan yang indah. Meskipun terkenal sebagai tempat wisata, pantai ini tetap mempertahankan keindahan alamnya yang memukau.",
		},
		{
			ID:           10,
			Date:         "Senin, 8 Agustus 2022",
			UserID:       2,
			Rating:       9,
			EcotourismID: 8,
			Body:         "Air terjun Tegenungan menawarkan pemandangan yang luar biasa dan suasana yang damai. Saya memberikan rating 8 karena meskipun tempat ini kadang-kadang cukup ramai, tetapi tempat ini sangat cocok bagi mereka yang ingin bersantai dan menikmati keindahan alam. Namun, pengunjung harus berhati-hati ketika mendekati air terjun karena jalannya cukup licin.",
		},
		{
			ID:           11,
			Date:         "Rabu, 19 Oktober 2022",
			UserID:       1,
			Rating:       9,
			EcotourismID: 6,
			Body:         "Saya merasa kecewa dengan Air terjun Sekumpul. Meskipun pemandangan di sekitar air terjun indah, tapi jalan menuju air terjun sangat sulit dan berbahaya. Selain itu, saya merasa bahwa akses menuju air terjun terlalu ramai, dan suasana di sekitar air terjun terasa sangat ramai dan tidak tenang.",
		},
		{
			ID:           12,
			Date:         "Selasa, 12 September 2022",
			UserID:       2,
			Rating:       5,
			EcotourismID: 7,
			Body:         "Saya merasa kecewa dengan Air terjun Sekumpul. Meskipun pemandangan di sekitar air terjun indah, tapi jalan menuju air terjun sangat sulit dan berbahaya. Selain itu, saya merasa bahwa akses menuju air terjun terlalu ramai, dan suasana di sekitar air terjun terasa sangat ramai dan tidak tenang.",
		},
		{
			ID:           13,
			Date:         "Senin, 15 November 2021",
			UserID:       3,
			Rating:       8,
			EcotourismID: 9,
			Body:         "Saya memberi Danau Beratan rating 8/10. Danau ini memiliki pemandangan yang sangat indah dan menenangkan, dengan airnya yang jernih danau ini menjadi tempat yang cocok untuk berlibur dan bersantai. Sayangnya, di beberapa waktu danau ini cukup ramai dan sulit untuk menikmati keindahannya secara tenang.",
		},
		{
			ID:           14,
			Date:         "Kamis, 7 Oktober 2021",
			UserID:       3,
			Rating:       4,
			EcotourismID: 10,
			Body:         "Pantai Kelingking adalah destinasi wisata yang menakjubkan di Nusa Penida. Sayangnya, saya memberikan rating 4 dari 10 karena kondisi pantai yang kurang bersih. Terdapat banyak sampah dan debris di sepanjang pantai yang mengurangi pengalaman wisata. Namun, pemandangan alam yang spektakuler dan keindahan pantai tetap membuat saya terkesan.",
		},
		{
			ID:           15,
			Date:         "Senin, 11 Januari 2021",
			UserID:       1,
			Rating:       9,
			EcotourismID: 11,
			Body:         "Saya sangat terkesan dengan keindahan alam yang luar biasa di Dieng Plateau. Udara segar dan pemandangan yang spektakuler membuat saya merasa sangat damai dan tenang di sana. Selain itu, Dieng Plateau juga memiliki banyak situs purbakala yang menarik dan memiliki nilai sejarah yang tinggi. Namun, sayangnya akses ke tempat wisata ini agak sulit, dan terkadang terdapat beberapa penjual yang agak mengganggu kenyamanan pengunjung. Oleh karena itu, saya memberi rating 9 untuk Dieng Plateau.",
		},
		{
			ID:           16,
			Date:         "Rabu, 23 Februari 2022",
			UserID:       3,
			Rating:       8,
			EcotourismID: 12,
			Body:         "Pulau Sempu adalah salah satu tempat wisata alam yang paling indah di Jawa Timur. Keindahan pantai dan hutan tropis yang masih sangat alami membuat saya merasa seperti berada di surga. Selain itu, kegiatan seperti trekking dan camping juga sangat menyenangkan di sini. Namun, sayangnya akses ke pulau ini agak sulit, dan terdapat beberapa pengunjung yang kurang memperhatikan kebersihan lingkungan. Oleh karena itu, saya memberi rating 8 untuk Pulau Sempu.",
		},
		{
			ID:           17,
			Date:         "Jumat, 15 April 2022",
			UserID:       2,
			Rating:       8,
			EcotourismID: 13,
			Body:         "Saya sangat terkesan dengan keindahan alam yang menakjubkan di Kawah Putih. Danau berwarna putih kehijauan yang terbentuk dari gas vulkanik dan tanah vulkanik yang putih membuat saya merasa seperti berada di surga. Selain itu, pemandangan di sekitar kawah juga sangat indah dan menenangkan. Namun, sayangnya akses ke tempat wisata ini agak sulit, dan terkadang terdapat beberapa penjual yang agak mengganggu kenyamanan pengunjung. Oleh karena itu, saya memberi rating 8 untuk Kawah Putih.",
		},
		{
			ID:           18,
			Date:         "Kamis, 19 Mei 2022",
			UserID:       3,
			Rating:       4,
			EcotourismID: 14,
			Body:         "Taman Nasional Bromo Tengger Semeru adalah tempat yang sangat indah dan menakjubkan. Pemandangan gunung berapi dan padang rumput yang luas sangat memukau. Namun, sayangnya terdapat banyak sampah di sekitar taman nasional ini, dan pengunjung terkadang kurang disiplin dalam membuang sampah. Oleh karena itu, saya memberi rating 8 untuk Taman Nasional Bromo Tengger Semeru.",
		},
		{
			ID:           19,
			Date:         "Senin, 11 Juli 2022",
			UserID:       1,
			Rating:       7,
			EcotourismID: 15,
			Body:         "Pantai Klayar memiliki keindahan alam yang memukau dengan pasir putihnya yang halus dan air lautnya yang jernih. Selain itu, pantai ini juga dikelilingi oleh tebing-tebing karang yang menjulang tinggi, memberikan pemandangan yang spektakuler. Namun, sayangnya pantai ini cukup sulit diakses karena terletak di ujung selatan Jawa Timur dan jalan yang menuju pantai kurang memadai. Selain itu, terdapat beberapa penjual yang terlalu agresif dalam menawarkan barang dagangan mereka. Oleh karena itu, saya memberi rating 7 untuk Pantai Klayar.",
		},
		{
			ID:           20,
			Date:         "Sabtu, 13 Agustus 2022",
			UserID:       2,
			Rating:       4,
			EcotourismID: 16,
			Body:         "Situ Patenggang adalah sebuah danau yang terletak di kawasan wisata Ciwidey, Bandung, Jawa Barat. Danau ini memiliki pemandangan yang indah dengan air yang jernih dan tenang serta dikelilingi oleh pegunungan dan pepohonan yang hijau. Selain itu, di sekitar danau terdapat warung-warung makan yang menyediakan makanan dan minuman lokal yang lezat. Namun, sayangnya di sekitar danau terdapat banyak sampah yang menumpuk dan tidak tertangani dengan baik, sehingga membuat lingkungan sekitar menjadi tidak nyaman. Oleh karena itu, saya memberi rating 6 untuk Situ Patenggang.",
		},
	}

	for _, comment := range comments {
		db.Debug().Create(&comment)
	}
}

func (uc *Usecase) GenerateUserDummy(db *mysql.DB) {
	kzh, _ := password.GeneratePassword("kazuha444")
	kai, _ := password.GeneratePassword("kai444")
	nky, _ := password.GeneratePassword("nakya444")

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
		{
			ID:           3,
			Username:     "nakya",
			Password:     nky,
			Email:        "nakya@gmail.com",
			Wallet:       75000,
			Contact:      "089765145652",
			Region:       "Bandung",
			Gender:       enum.Pria,
			Birthday:     "10 Agustus 2000",
			PhotoProfile: "https://arcudskzafkijqukfool.supabase.co/storage/v1/object/public/traveleen/images/shigeo-kageyama-rage-mob-psycho-uhdpaper.com-4K-1.jpg",
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
		db.Debug().Create(&user)
	}

	for _, trash := range trashes {
		db.Debug().Create(&trash)
	}

	for _, purchase := range purchases {
		db.Debug().Create(&purchase)
	}
}
