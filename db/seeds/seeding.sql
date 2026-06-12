-- ================================================================
--  TICKITZ — DUMMY DATA v3.0 (JUNI 2026 — FILM ASLI)
--  Update:
--  ✓ Film diganti dengan film asli tayang Juni 2026 di Indonesia
--  ✓ Showtimes menggunakan CURRENT_DATE agar selalu relevan
--  ✓ Now Showing  : 10 film (rilis sebelum/tepat hari ini)
--  ✓ Upcoming     : 8 film (rilis setelah hari ini)
--  ✓ Booking      : HANYA untuk film Now Showing (logika bisnis benar)
--  ✓ Semua tabel lain (cinema, seats, users, dll) tetap sama
--  ✓ Poster via TMDB image path (ganti BASE_URL sesuai env kamu)
-- ================================================================


-- ============================================================
-- 1. LOCATIONS (sama seperti sebelumnya)
-- 1. LOCATIONS (12 kota Indonesia +realase_data Purwokerto sesuai desain)
-- ============================================================
INSERT INTO locations (id, city) VALUES
(1,  'Jakarta'),
(2,  'Surabaya'),
(3,  'Bandung'),
(4,  'Medan'),
(5,  'Semarang'),
(6,  'Makassar'),
(7,  'Palembang'),
(8,  'Denpasar'),
(9,  'Yogyakarta'),
(10, 'Malang'),
(11, 'Purwokerto'),
(12, 'Bekasi');


-- ============================================================
-- 2. PAYMENT METHODS
-- ============================================================
INSERT INTO payment_methods (id, logo, name) VALUES
(1, 'https://storage.tickitz.id/pm/gopay.png',     'GoPay'),
(2, 'https://storage.tickitz.id/pm/ovo.png',       'OVO'),
(3, 'https://storage.tickitz.id/pm/dana.png',      'DANA'),
(4, 'https://storage.tickitz.id/pm/bca.png',       'Transfer BCA'),
(5, 'https://storage.tickitz.id/pm/mandiri.png',   'Transfer Mandiri'),
(6, 'https://storage.tickitz.id/pm/bni.png',       'Transfer BNI'),
(7, 'https://storage.tickitz.id/pm/creditcard.png','Kartu Kredit'),
(8, 'https://storage.tickitz.id/pm/alfamart.png',  'Alfamart');


-- ============================================================
-- 3. GENRES
-- ============================================================
INSERT INTO genres (id, genre) VALUES
(1,  'Action'),
(2,  'Comedy'),
(3,  'Drama'),
(4,  'Horror'),
(5,  'Romance'),
(6,  'Sci-Fi'),
(7,  'Thriller'),
(8,  'Animation'),
(9,  'Adventure'),
(10, 'Fantasy'),
(11, 'Crime'),
(12, 'Documentary');


-- ============================================================
-- 4. DIRECTORS
-- ============================================================
INSERT INTO directors (id, name) VALUES
(1,  'Jose Poernomo'),
(2,  'Herwin Novianto'),
(3,  'Chiska Doppert'),
(4,  'Sondang Pratama'),
(5,  'Travis Knight'),
(6,  'Craig Gillespie'),
(7,  'Pete Docter'),
(8,  'Pierre Coffin'),
(9,  'Yeon Sang-ho'),
(10, 'Anggy Umbara'),
(11, 'Rizal Mantovani'),
(12, 'Ody C. Harahap'),
(13, 'Fajar Bustomi'),
(14, 'Steven Spielberg'),
(15, 'Wregas Bhanuteja'),
(16, 'Viva Westi'),
(17, 'Jake Kasdan'),
(18, 'Riri Riza'),
(19, 'Mouly Surya'),
(20, 'Adriyanto Dewo');


-- ============================================================
-- 5. CASTS
-- ============================================================
INSERT INTO casts (id, name) VALUES
(1,  'Iqbaal Ramadhan'),
(2,  'Rachel Amanda'),
(3,  'Caroline Zachrie'),
(4,  'Keiko Ananta'),
(5,  'Vino G. Bastian'),
(6,  'Tora Sudiro'),
(7,  'Deddy Mahendra Desta'),
(8,  'Indro Warkop'),
(9,  'Marsha Timothy'),
(10, 'Jefan Nathanio'),
(11, 'Hana Saraswati'),
(12, 'Dodit Mulyanto'),
(13, 'Nicholas Galitzine'),
(14, 'Camila Mendes'),
(15, 'Idris Elba'),
(16, 'Milly Alcock'),
(17, 'Jason Momoa'),
(18, 'Tom Hanks'),
(19, 'Tim Allen'),
(20, 'Steve Carell'),
(21, 'Joe Taslim'),
(22, 'Yayan Ruhian'),
(23, 'Emily Blunt'),
(24, 'Josh O''Connor'),
(25, 'Reza Rahadian'),
(26, 'Prilly Latuconsina'),
(27, 'Chelsea Islan'),
(28, 'Nicholas Saputra'),
(29, 'Adipati Dolken'),
(30, 'Refal Hady');


-- ============================================================
-- 6. MOVIES (18 film — 10 Now Showing + 8 Upcoming)
--
--  NOW SHOWING  (id 1-10)  : release_date <= CURRENT_DATE
--  UPCOMING     (id 11-18) : release_date >  CURRENT_DATE
--
--  Tanggal NOW SHOWING pakai offset negatif dari CURRENT_DATE
--  Tanggal UPCOMING pakai offset positif dari CURRENT_DATE
--
--  Poster: gunakan TMDB image URL (ganti dengan env kamu)
--  TMDB base: https://image.tmdb.org/t/p/w500
-- ============================================================
INSERT INTO movies (id, title, duration, poster, release_date, synopsis, category) VALUES
-- ── NOW SHOWING ──────────────────────────────────────────────

(1,  'Monster Pabrik Rambut',
     '01:48:00',
     'img/monster-pabrik-rambut.webp',
     (CURRENT_DATE - INTERVAL '7 days'),
     'Putri nekat mengungkap misteri di balik kematian ibunya yang bekerja di pabrik rambut misterius. Pihak pabrik mengklaim sang ibu bunuh diri, namun Putri menemukan kebenaran gelap yang melibatkan arwah Sulastri dan rahasia lama yang terkubur.',
     '17+'),

(2,  'Kucing Hitam',
     '01:42:00',
     'img/kucing-hitam.webp',
     (CURRENT_DATE - INTERVAL '7 days'),
     'Kehidupan psikiater Natalie berubah drastis setelah putri kecilnya membawa pulang seekor kucing hitam misterius. Teror demi teror datang menghantui keluarganya, dan Natalie perlahan menyadari kucing itu bukan sekadar hewan biasa.',
     '17+'),

(3,  'Warkop DKI: Viralin Doooong!!',
     '01:55:00',
     'img/warkop-dki.webp',
     (CURRENT_DATE - INTERVAL '7 days'),
     'Dono, Kasino, dan Indro yang terlilit masalah keuangan mencoba peruntungan sebagai konten kreator. Ide syuting konten horor di desa terpencil malah membawa mereka ke teror mistis yang nyata, memaksa mereka menebus dosa masa lalu.',
     'SU'),

(4,  'Colony',
     '02:02:00',
     'img/colony.webp',
     (CURRENT_DATE - INTERVAL '7 days'),
     'Sebuah virus mematikan menyebar di konferensi bioteknologi internasional dan mengubah para peserta menjadi makhluk mirip zombie. Sekelompok penyintas harus bertaruh nyawa menembus gedung yang telah berubah menjadi neraka untuk melarikan diri.',
     '17+'),

(5,  'Masters of the Universe',
     '02:10:00',
     'img/master-universe.webp',
     (CURRENT_DATE - INTERVAL '5 days'),
     'Setelah 15 tahun terpisah, pedang kekuatan membawa Adam kembali ke Eternia yang hancur di bawah kekuasaan Skeletor. Bersama Teela dan Man-At-Arms, Adam harus merangkul takdirnya sebagai He-Man untuk menyelamatkan dunianya.',
     '13+'),

(6,  'Nobody Loves Kay',
     '01:50:00',
     'img/nobody-loves-kay.webp',
     (CURRENT_DATE - INTERVAL '7 days'),
     'Kay adalah gamer pro wanita yang berjuang membuktikan dirinya di dunia esports yang didominasi pria. Ketika timnya terancam bubar dan sponsor menarik diri, Kay harus memilih antara gengsi dan hati nuraninya.',
     '13+'),

(7,  'Jangan Buang Ibu',
     '01:58:00',
     'img/jangan-buang-ibu.webp',
     (CURRENT_DATE - INTERVAL '5 days'),
     'Ristiana berjuang seorang diri membesarkan tiga anaknya setelah ditinggal suami yang meninggalkan utang. Ketika anak-anaknya sudah dewasa, mereka justru mengirimnya ke panti jompo. Drama keluarga yang menyentuh hati tentang pengorbanan seorang ibu.',
     'SU'),

(8,  'Garuda di Dadaku',
     '01:35:00',
     'img/garuda-di-dadaku.webp',
     (CURRENT_DATE - INTERVAL '3 days'),
     'Putra, remaja 13 tahun penderita asma, bermimpi menjadi pemain timnas Indonesia. Dengan tekad baja dan dukungan sahabatnya Gaga, Putra membuktikan bahwa keterbatasan fisik bukan halangan untuk meraih mimpi setinggi langit.',
     'SU'),

(9,  'Tanah Runtuh',
     '02:05:00',
     'img/tanah-runtuh.webp',
     (CURRENT_DATE - INTERVAL '3 days'),
     'Berlatar Kerusuhan Poso 2005, Kai dan adiknya Ringgo yang memiliki Down Syndrome terpisah dari sang ibu di tengah situasi mencekam. Dengan bantuan polisi Idham, dua bersaudara ini berjuang menyusuri pengungsian demi menemukan keluarganya.',
     '13+'),

(10, 'The Furious',
     '02:08:00',
     'img/the-forious.webp',
     (CURRENT_DATE - INTERVAL '2 days'),
     'Pendekar laga Indonesia Arman mencari balas dendam setelah putrinya diculik oleh jaringan kriminal internasional. Aksi brutal dan tak kenal ampun Joe Taslim dan Yayan Ruhian menjadi mesin penghancur setiap rintangan yang menghadang.',
     '17+'),


-- ── UPCOMING ─────────────────────────────────────────────────

(11, 'Toy Story 5',
     '01:45:00',
     'img/toy-story-5.webp',
     (CURRENT_DATE + INTERVAL '8 days'),
     'Woody, Buzz Lightyear, dan sahabat-sahabat mereka kembali dalam petualangan terbaru. Di era gadget yang semakin canggih, mereka harus membuktikan bahwa mainan fisik masih punya tempat di hati anak-anak masa kini.',
     'SU'),

(12, 'Supergirl: Woman of Tomorrow',
     '02:12:00',
     'img/supergirl.webp',
     (CURRENT_DATE + INTERVAL '8 days'),
     'Kara Zor-El alias Supergirl harus bersekutu dengan rekan tak terduga ketika musuh kejam menyerang orang-orang terdekatnya. Perjalanan antargalaksi penuh aksi ini menguji batas kemampuan dan keberanian sang pahlawan super.',
     '13+'),

(13, 'Disclosure Day',
     '02:18:00',
     'img/disclosure-day.webp',
     (CURRENT_DATE + INTERVAL '10 days'),
     'Pembawa berita cuaca Emily tiba-tiba berperilaku tak terkendali saat siaran langsung, seolah dikendalikan kekuatan asing. Seorang pria misterius mengklaim memiliki bukti bahwa pemerintah telah lama menyembunyikan kontak dengan peradaban alien.',
     '6'),

(14, 'Dukun Magang',
     '01:40:00',
     'img/dukun-magang.webp',
     (CURRENT_DATE + INTERVAL '7 days'),
     'Raka, mahasiswa yang tidak sengaja magang di pesugihan, harus menyelamatkan kampung halamannya dari teror Kuntilanak Hitam dengan kemampuan dukun setengah matang yang ia miliki. Horor komedi yang bikin ngakak sekaligus merinding.',
     '13+'),

(15, 'Cerita Lila',
     '02:00:00',
     'img/cerita-lila.webp',
     (CURRENT_DATE + INTERVAL '7 days'),
     'Diangkat dari kisah nyata, Lila dan Lili — anak kembar yang mengalami kekerasan oleh ibu kandung mereka sendiri. Kisah perjuangan mereka untuk bertahan hidup dan menemukan kebahagiaan di tengah trauma yang mendalam.',
     '17+'),

(16, 'The Longest Wait',
     '01:28:00',
     'img/the-longest-wait.webp',
     (CURRENT_DATE + INTERVAL '7 days'),
     'Dokumenter yang mengabadikan momen bersejarah timnas Indonesia yang lolos ke putaran ketiga Kualifikasi Piala Dunia 2026 untuk pertama kalinya. Menyelami sisi humanis para pemain dari ruang ganti hingga cerita perjuangan yang selama ini tersembunyi.',
     'SU'),

(17, 'Minions & Monsters',
     '01:32:00',
     'img/minions-monster.webp',
     (CURRENT_DATE + INTERVAL '14 days'),
     'Para Minion kesayangan kini harus menghadapi serangan monster raksasa yang mengancam kediaman Gru. Petualangan kocak penuh warna yang menghibur seluruh anggota keluarga dari anak-anak hingga orang dewasa.',
     'SU'),

(18, 'Backrooms',
     '01:55:00',
     'img/backrooms.webp',
     (CURRENT_DATE + INTERVAL '14 days'),
     'Diadaptasi dari fenomena internet viral, sekelompok orang terjebak di dimensi paralel penuh lorong-lorong tanpa akhir yang dihuni entitas tak kasat mata. Thriller psikologis yang mengeksplorasi ketakutan terdalam manusia akan ketersesatan abadi.',
     '17+');


-- ============================================================
-- 7. MOVIE_GENRES
-- ============================================================
INSERT INTO movie_genres (movie_id, genre_id) VALUES
(1,  4), (1,  7),           -- Monster Pabrik Rambut: Horror, Thriller
(2,  4),                    -- Kucing Hitam: Horror
(3,  2), (3,  4),           -- Warkop DKI: Comedy, Horror
(4,  4), (4,  7), (4,  1),  -- Colony: Horror, Thriller, Action
(5,  1), (5,  9), (5, 10),  -- Masters of the Universe: Action, Adventure, Fantasy
(6,  1), (6,  3),           -- Nobody Loves Kay: Action, Drama
(7,  3), (7,  5),           -- Jangan Buang Ibu: Drama, Romance
(8,  3), (8,  9),           -- Garuda di Dadaku: Drama, Adventure
(9,  3), (9,  7),           -- Tanah Runtuh: Drama, Thriller
(10, 1), (10, 7), (10,11),  -- The Furious: Action, Thriller, Crime
(11, 8), (11, 9), (11, 2),  -- Toy Story 5: Animation, Adventure, Comedy
(12, 1), (12, 9), (12, 6),  -- Supergirl: Action, Adventure, Sci-Fi
(13, 6), (13, 7),           -- Disclosure Day: Sci-Fi, Thriller
(14, 4), (14, 2),           -- Dukun Magang: Horror, Comedy
(15, 3), (15, 7),           -- Cerita Lila: Drama, Thriller
(16,12),                    -- The Longest Wait: Documentary
(17, 8), (17, 2), (17, 9),  -- Minions & Monsters: Animation, Comedy, Adventure
(18, 4), (18, 7), (18, 6);  -- Backrooms: Horror, Thriller, Sci-Fi


-- ============================================================
-- 8. MOVIE_DIRECTORS
-- ============================================================
INSERT INTO movie_directors (movie_id, director_id) VALUES
(1,  1),   -- Monster Pabrik Rambut — Jose Poernomo
(2,  16),  -- Kucing Hitam — Viva Westi
(3,  2),   -- Warkop DKI — Herwin Novianto
(4,  9),   -- Colony — Yeon Sang-ho
(5,  5),   -- Masters of the Universe — Travis Knight
(6,  10),  -- Nobody Loves Kay — Anggy Umbara
(7,  18),  -- Jangan Buang Ibu — Riri Riza
(8,  12),  -- Garuda di Dadaku — Ody C. Harahap
(9,  20),  -- Tanah Runtuh — Adriyanto Dewo
(10, 13),  -- The Furious — Fajar Bustomi
(11, 7),   -- Toy Story 5 — Pete Docter
(12, 6),   -- Supergirl — Craig Gillespie
(13,14),   -- Disclosure Day — Steven Spielberg
(14, 3),   -- Dukun Magang — Chiska Doppert
(15, 4),   -- Cerita Lila — Sondang Pratama
(16,15),   -- The Longest Wait — Wregas Bhanuteja
(17, 8),   -- Minions & Monsters — Pierre Coffin
(18,11);   -- Backrooms — Rizal Mantovani


-- ============================================================
-- 9. MOVIE_CASTS
-- ============================================================
INSERT INTO movie_casts (movie_id, cast_id) VALUES
(1,  1), (1,  2),            -- Monster Pabrik Rambut: Iqbaal, Rachel Amanda
(2,  3), (2,  4),            -- Kucing Hitam: Caroline Zachrie, Keiko Ananta
(3,  5), (3,  6), (3,  7), (3, 8), (3, 9),  -- Warkop DKI
(4,  25),(4,  26),           -- Colony: (karakter Korea, pakai aktor lokal populer utk demo)
(5,  13),(5,  14),(5, 15),   -- Masters: Nicholas Galitzine, Camila Mendes, Idris Elba
(6,  27),(6,  28),           -- Nobody Loves Kay: Chelsea Islan, Nicholas Saputra
(7,  9), (7,  29),           -- Jangan Buang Ibu: Marsha Timothy, Adipati Dolken
(8,  26),(8,  30),           -- Garuda: Prilly Latuconsina, Refal Hady
(9,  5), (9,  25),           -- Tanah Runtuh: Vino G. Bastian, Reza Rahadian
(10,21), (10,22),            -- The Furious: Joe Taslim, Yayan Ruhian
(11,18), (11,19),            -- Toy Story 5: Tom Hanks (Woody), Tim Allen (Buzz)
(12,16), (12,17),            -- Supergirl: Milly Alcock, Jason Momoa
(13,23), (13,24),            -- Disclosure Day: Emily Blunt, Josh O''Connor
(14,10), (14,11),(14,12),    -- Dukun Magang: Jefan, Hana, Dodit
(15,27), (15,29),            -- Cerita Lila: Chelsea Islan, Adipati Dolken
(16,25),                     -- The Longest Wait
(17,20),                     -- Minions & Monsters: Steve Carell (Gru)
(18,26),(18,28);             -- Backrooms: Prilly, Nicholas Saputra


-- ============================================================
-- 10. CINEMAS (15 bioskop — sama seperti sebelumnya)
-- ============================================================
INSERT INTO cinemas (id, location_id, name, logo, capacity, isAvailable) VALUES
(1,  1, 'ebv.id Grand Indonesia',         'https://storage.tickitz.id/cinema/ebvid.png',      98, true),
(2,  1, 'hiflix FX Sudirman',             'https://storage.tickitz.id/cinema/hiflix.png',     98, true),
(3,  1, 'CineOne21 Taman Ismail Marzuki', 'https://storage.tickitz.id/cinema/cineone21.png',  98, true),
(4,  2, 'ebv.id Tunjungan Plaza',         'https://storage.tickitz.id/cinema/ebvid.png',      98, true),
(5,  2, 'hiflix Galaxy Mall',             'https://storage.tickitz.id/cinema/hiflix.png',     98, true),
(6,  2, 'CineOne21 Pakuwon Mall',         'https://storage.tickitz.id/cinema/cineone21.png',  98, true),
(7,  3, 'ebv.id Paris Van Java',          'https://storage.tickitz.id/cinema/ebvid.png',      98, true),
(8,  3, 'hiflix Bandung Indah Plaza',     'https://storage.tickitz.id/cinema/hiflix.png',     98, true),
(9,  4, 'CineOne21 Sun Plaza',            'https://storage.tickitz.id/cinema/cineone21.png',  98, true),
(10, 5, 'ebv.id Paragon Mall',            'https://storage.tickitz.id/cinema/ebvid.png',      98, true),
(11, 9, 'hiflix Ambarrukmo Plaza',        'https://storage.tickitz.id/cinema/hiflix.png',     98, true),
(12, 8, 'CineOne21 Bali Galeria',         'https://storage.tickitz.id/cinema/cineone21.png',  98, true),
(13, 6, 'ebv.id Panakukang Mall',         'https://storage.tickitz.id/cinema/ebvid.png',      98, true),
(14, 7, 'hiflix Palembang Square',        'https://storage.tickitz.id/cinema/hiflix.png',     98, true),
(15,11, 'CineOne21 Mayfair Purwokerto',   'https://storage.tickitz.id/cinema/cineone21.png',  98, false);


-- ============================================================
-- 11. SEATS (sama persis — 15 cinema × 98 kursi)
--     Layout: A,B,C,D,E,LN,G × 14 kolom
-- ============================================================

-- Macro: setiap cinema punya layout yang sama persis
-- Cinema 1
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(1,1,'A','regular'),(1,2,'A','regular'),(1,3,'A','regular'),(1,4,'A','regular'),(1,5,'A','regular'),(1,6,'A','regular'),(1,7,'A','regular'),(1,8,'A','regular'),(1,9,'A','regular'),(1,10,'A','regular'),(1,11,'A','regular'),(1,12,'A','regular'),(1,13,'A','regular'),(1,14,'A','regular'),
(1,1,'B','regular'),(1,2,'B','regular'),(1,3,'B','regular'),(1,4,'B','regular'),(1,5,'B','regular'),(1,6,'B','regular'),(1,7,'B','regular'),(1,8,'B','regular'),(1,9,'B','regular'),(1,10,'B','regular'),(1,11,'B','regular'),(1,12,'B','regular'),(1,13,'B','regular'),(1,14,'B','regular'),
(1,1,'C','regular'),(1,2,'C','regular'),(1,3,'C','regular'),(1,4,'C','regular'),(1,5,'C','regular'),(1,6,'C','regular'),(1,7,'C','regular'),(1,8,'C','regular'),(1,9,'C','regular'),(1,10,'C','regular'),(1,11,'C','regular'),(1,12,'C','regular'),(1,13,'C','regular'),(1,14,'C','regular'),
(1,1,'D','regular'),(1,2,'D','regular'),(1,3,'D','regular'),(1,4,'D','regular'),(1,5,'D','regular'),(1,6,'D','regular'),(1,7,'D','regular'),(1,8,'D','regular'),(1,9,'D','regular'),(1,10,'D','regular'),(1,11,'D','regular'),(1,12,'D','regular'),(1,13,'D','regular'),(1,14,'D','regular'),
(1,1,'E','regular'),(1,2,'E','regular'),(1,3,'E','regular'),(1,4,'E','regular'),(1,5,'E','regular'),(1,6,'E','regular'),(1,7,'E','regular'),(1,8,'E','regular'),(1,9,'E','regular'),(1,10,'E','regular'),(1,11,'E','regular'),(1,12,'E','regular'),(1,13,'E','regular'),(1,14,'E','regular'),
(1,1,'LN','love_nest'),(1,2,'LN','love_nest'),(1,3,'LN','love_nest'),(1,4,'LN','love_nest'),(1,5,'LN','love_nest'),(1,6,'LN','love_nest'),(1,7,'LN','love_nest'),(1,8,'LN','love_nest'),(1,9,'LN','love_nest'),(1,10,'LN','love_nest'),(1,11,'LN','love_nest'),(1,12,'LN','love_nest'),(1,13,'LN','love_nest'),(1,14,'LN','love_nest'),
(1,1,'G','regular'),(1,2,'G','regular'),(1,3,'G','regular'),(1,4,'G','regular'),(1,5,'G','regular'),(1,6,'G','regular'),(1,7,'G','regular'),(1,8,'G','regular'),(1,9,'G','regular'),(1,10,'G','regular'),(1,11,'G','regular'),(1,12,'G','regular'),(1,13,'G','regular'),(1,14,'G','regular');

-- Cinema 2
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(2,1,'A','regular'),(2,2,'A','regular'),(2,3,'A','regular'),(2,4,'A','regular'),(2,5,'A','regular'),(2,6,'A','regular'),(2,7,'A','regular'),(2,8,'A','regular'),(2,9,'A','regular'),(2,10,'A','regular'),(2,11,'A','regular'),(2,12,'A','regular'),(2,13,'A','regular'),(2,14,'A','regular'),
(2,1,'B','regular'),(2,2,'B','regular'),(2,3,'B','regular'),(2,4,'B','regular'),(2,5,'B','regular'),(2,6,'B','regular'),(2,7,'B','regular'),(2,8,'B','regular'),(2,9,'B','regular'),(2,10,'B','regular'),(2,11,'B','regular'),(2,12,'B','regular'),(2,13,'B','regular'),(2,14,'B','regular'),
(2,1,'C','regular'),(2,2,'C','regular'),(2,3,'C','regular'),(2,4,'C','regular'),(2,5,'C','regular'),(2,6,'C','regular'),(2,7,'C','regular'),(2,8,'C','regular'),(2,9,'C','regular'),(2,10,'C','regular'),(2,11,'C','regular'),(2,12,'C','regular'),(2,13,'C','regular'),(2,14,'C','regular'),
(2,1,'D','regular'),(2,2,'D','regular'),(2,3,'D','regular'),(2,4,'D','regular'),(2,5,'D','regular'),(2,6,'D','regular'),(2,7,'D','regular'),(2,8,'D','regular'),(2,9,'D','regular'),(2,10,'D','regular'),(2,11,'D','regular'),(2,12,'D','regular'),(2,13,'D','regular'),(2,14,'D','regular'),
(2,1,'E','regular'),(2,2,'E','regular'),(2,3,'E','regular'),(2,4,'E','regular'),(2,5,'E','regular'),(2,6,'E','regular'),(2,7,'E','regular'),(2,8,'E','regular'),(2,9,'E','regular'),(2,10,'E','regular'),(2,11,'E','regular'),(2,12,'E','regular'),(2,13,'E','regular'),(2,14,'E','regular'),
(2,1,'LN','love_nest'),(2,2,'LN','love_nest'),(2,3,'LN','love_nest'),(2,4,'LN','love_nest'),(2,5,'LN','love_nest'),(2,6,'LN','love_nest'),(2,7,'LN','love_nest'),(2,8,'LN','love_nest'),(2,9,'LN','love_nest'),(2,10,'LN','love_nest'),(2,11,'LN','love_nest'),(2,12,'LN','love_nest'),(2,13,'LN','love_nest'),(2,14,'LN','love_nest'),
(2,1,'G','regular'),(2,2,'G','regular'),(2,3,'G','regular'),(2,4,'G','regular'),(2,5,'G','regular'),(2,6,'G','regular'),(2,7,'G','regular'),(2,8,'G','regular'),(2,9,'G','regular'),(2,10,'G','regular'),(2,11,'G','regular'),(2,12,'G','regular'),(2,13,'G','regular'),(2,14,'G','regular');

-- Cinema 3
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(3,1,'A','regular'),(3,2,'A','regular'),(3,3,'A','regular'),(3,4,'A','regular'),(3,5,'A','regular'),(3,6,'A','regular'),(3,7,'A','regular'),(3,8,'A','regular'),(3,9,'A','regular'),(3,10,'A','regular'),(3,11,'A','regular'),(3,12,'A','regular'),(3,13,'A','regular'),(3,14,'A','regular'),
(3,1,'B','regular'),(3,2,'B','regular'),(3,3,'B','regular'),(3,4,'B','regular'),(3,5,'B','regular'),(3,6,'B','regular'),(3,7,'B','regular'),(3,8,'B','regular'),(3,9,'B','regular'),(3,10,'B','regular'),(3,11,'B','regular'),(3,12,'B','regular'),(3,13,'B','regular'),(3,14,'B','regular'),
(3,1,'C','regular'),(3,2,'C','regular'),(3,3,'C','regular'),(3,4,'C','regular'),(3,5,'C','regular'),(3,6,'C','regular'),(3,7,'C','regular'),(3,8,'C','regular'),(3,9,'C','regular'),(3,10,'C','regular'),(3,11,'C','regular'),(3,12,'C','regular'),(3,13,'C','regular'),(3,14,'C','regular'),
(3,1,'D','regular'),(3,2,'D','regular'),(3,3,'D','regular'),(3,4,'D','regular'),(3,5,'D','regular'),(3,6,'D','regular'),(3,7,'D','regular'),(3,8,'D','regular'),(3,9,'D','regular'),(3,10,'D','regular'),(3,11,'D','regular'),(3,12,'D','regular'),(3,13,'D','regular'),(3,14,'D','regular'),
(3,1,'E','regular'),(3,2,'E','regular'),(3,3,'E','regular'),(3,4,'E','regular'),(3,5,'E','regular'),(3,6,'E','regular'),(3,7,'E','regular'),(3,8,'E','regular'),(3,9,'E','regular'),(3,10,'E','regular'),(3,11,'E','regular'),(3,12,'E','regular'),(3,13,'E','regular'),(3,14,'E','regular'),
(3,1,'LN','love_nest'),(3,2,'LN','love_nest'),(3,3,'LN','love_nest'),(3,4,'LN','love_nest'),(3,5,'LN','love_nest'),(3,6,'LN','love_nest'),(3,7,'LN','love_nest'),(3,8,'LN','love_nest'),(3,9,'LN','love_nest'),(3,10,'LN','love_nest'),(3,11,'LN','love_nest'),(3,12,'LN','love_nest'),(3,13,'LN','love_nest'),(3,14,'LN','love_nest'),
(3,1,'G','regular'),(3,2,'G','regular'),(3,3,'G','regular'),(3,4,'G','regular'),(3,5,'G','regular'),(3,6,'G','regular'),(3,7,'G','regular'),(3,8,'G','regular'),(3,9,'G','regular'),(3,10,'G','regular'),(3,11,'G','regular'),(3,12,'G','regular'),(3,13,'G','regular'),(3,14,'G','regular');

-- Cinema 4
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(4,1,'A','regular'),(4,2,'A','regular'),(4,3,'A','regular'),(4,4,'A','regular'),(4,5,'A','regular'),(4,6,'A','regular'),(4,7,'A','regular'),(4,8,'A','regular'),(4,9,'A','regular'),(4,10,'A','regular'),(4,11,'A','regular'),(4,12,'A','regular'),(4,13,'A','regular'),(4,14,'A','regular'),
(4,1,'B','regular'),(4,2,'B','regular'),(4,3,'B','regular'),(4,4,'B','regular'),(4,5,'B','regular'),(4,6,'B','regular'),(4,7,'B','regular'),(4,8,'B','regular'),(4,9,'B','regular'),(4,10,'B','regular'),(4,11,'B','regular'),(4,12,'B','regular'),(4,13,'B','regular'),(4,14,'B','regular'),
(4,1,'C','regular'),(4,2,'C','regular'),(4,3,'C','regular'),(4,4,'C','regular'),(4,5,'C','regular'),(4,6,'C','regular'),(4,7,'C','regular'),(4,8,'C','regular'),(4,9,'C','regular'),(4,10,'C','regular'),(4,11,'C','regular'),(4,12,'C','regular'),(4,13,'C','regular'),(4,14,'C','regular'),
(4,1,'D','regular'),(4,2,'D','regular'),(4,3,'D','regular'),(4,4,'D','regular'),(4,5,'D','regular'),(4,6,'D','regular'),(4,7,'D','regular'),(4,8,'D','regular'),(4,9,'D','regular'),(4,10,'D','regular'),(4,11,'D','regular'),(4,12,'D','regular'),(4,13,'D','regular'),(4,14,'D','regular'),
(4,1,'E','regular'),(4,2,'E','regular'),(4,3,'E','regular'),(4,4,'E','regular'),(4,5,'E','regular'),(4,6,'E','regular'),(4,7,'E','regular'),(4,8,'E','regular'),(4,9,'E','regular'),(4,10,'E','regular'),(4,11,'E','regular'),(4,12,'E','regular'),(4,13,'E','regular'),(4,14,'E','regular'),
(4,1,'LN','love_nest'),(4,2,'LN','love_nest'),(4,3,'LN','love_nest'),(4,4,'LN','love_nest'),(4,5,'LN','love_nest'),(4,6,'LN','love_nest'),(4,7,'LN','love_nest'),(4,8,'LN','love_nest'),(4,9,'LN','love_nest'),(4,10,'LN','love_nest'),(4,11,'LN','love_nest'),(4,12,'LN','love_nest'),(4,13,'LN','love_nest'),(4,14,'LN','love_nest'),
(4,1,'G','regular'),(4,2,'G','regular'),(4,3,'G','regular'),(4,4,'G','regular'),(4,5,'G','regular'),(4,6,'G','regular'),(4,7,'G','regular'),(4,8,'G','regular'),(4,9,'G','regular'),(4,10,'G','regular'),(4,11,'G','regular'),(4,12,'G','regular'),(4,13,'G','regular'),(4,14,'G','regular');

-- Cinema 5
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(5,1,'A','regular'),(5,2,'A','regular'),(5,3,'A','regular'),(5,4,'A','regular'),(5,5,'A','regular'),(5,6,'A','regular'),(5,7,'A','regular'),(5,8,'A','regular'),(5,9,'A','regular'),(5,10,'A','regular'),(5,11,'A','regular'),(5,12,'A','regular'),(5,13,'A','regular'),(5,14,'A','regular'),
(5,1,'B','regular'),(5,2,'B','regular'),(5,3,'B','regular'),(5,4,'B','regular'),(5,5,'B','regular'),(5,6,'B','regular'),(5,7,'B','regular'),(5,8,'B','regular'),(5,9,'B','regular'),(5,10,'B','regular'),(5,11,'B','regular'),(5,12,'B','regular'),(5,13,'B','regular'),(5,14,'B','regular'),
(5,1,'C','regular'),(5,2,'C','regular'),(5,3,'C','regular'),(5,4,'C','regular'),(5,5,'C','regular'),(5,6,'C','regular'),(5,7,'C','regular'),(5,8,'C','regular'),(5,9,'C','regular'),(5,10,'C','regular'),(5,11,'C','regular'),(5,12,'C','regular'),(5,13,'C','regular'),(5,14,'C','regular'),
(5,1,'D','regular'),(5,2,'D','regular'),(5,3,'D','regular'),(5,4,'D','regular'),(5,5,'D','regular'),(5,6,'D','regular'),(5,7,'D','regular'),(5,8,'D','regular'),(5,9,'D','regular'),(5,10,'D','regular'),(5,11,'D','regular'),(5,12,'D','regular'),(5,13,'D','regular'),(5,14,'D','regular'),
(5,1,'E','regular'),(5,2,'E','regular'),(5,3,'E','regular'),(5,4,'E','regular'),(5,5,'E','regular'),(5,6,'E','regular'),(5,7,'E','regular'),(5,8,'E','regular'),(5,9,'E','regular'),(5,10,'E','regular'),(5,11,'E','regular'),(5,12,'E','regular'),(5,13,'E','regular'),(5,14,'E','regular'),
(5,1,'LN','love_nest'),(5,2,'LN','love_nest'),(5,3,'LN','love_nest'),(5,4,'LN','love_nest'),(5,5,'LN','love_nest'),(5,6,'LN','love_nest'),(5,7,'LN','love_nest'),(5,8,'LN','love_nest'),(5,9,'LN','love_nest'),(5,10,'LN','love_nest'),(5,11,'LN','love_nest'),(5,12,'LN','love_nest'),(5,13,'LN','love_nest'),(5,14,'LN','love_nest'),
(5,1,'G','regular'),(5,2,'G','regular'),(5,3,'G','regular'),(5,4,'G','regular'),(5,5,'G','regular'),(5,6,'G','regular'),(5,7,'G','regular'),(5,8,'G','regular'),(5,9,'G','regular'),(5,10,'G','regular'),(5,11,'G','regular'),(5,12,'G','regular'),(5,13,'G','regular'),(5,14,'G','regular');

-- Cinema 6
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(6,1,'A','regular'),(6,2,'A','regular'),(6,3,'A','regular'),(6,4,'A','regular'),(6,5,'A','regular'),(6,6,'A','regular'),(6,7,'A','regular'),(6,8,'A','regular'),(6,9,'A','regular'),(6,10,'A','regular'),(6,11,'A','regular'),(6,12,'A','regular'),(6,13,'A','regular'),(6,14,'A','regular'),
(6,1,'B','regular'),(6,2,'B','regular'),(6,3,'B','regular'),(6,4,'B','regular'),(6,5,'B','regular'),(6,6,'B','regular'),(6,7,'B','regular'),(6,8,'B','regular'),(6,9,'B','regular'),(6,10,'B','regular'),(6,11,'B','regular'),(6,12,'B','regular'),(6,13,'B','regular'),(6,14,'B','regular'),
(6,1,'C','regular'),(6,2,'C','regular'),(6,3,'C','regular'),(6,4,'C','regular'),(6,5,'C','regular'),(6,6,'C','regular'),(6,7,'C','regular'),(6,8,'C','regular'),(6,9,'C','regular'),(6,10,'C','regular'),(6,11,'C','regular'),(6,12,'C','regular'),(6,13,'C','regular'),(6,14,'C','regular'),
(6,1,'D','regular'),(6,2,'D','regular'),(6,3,'D','regular'),(6,4,'D','regular'),(6,5,'D','regular'),(6,6,'D','regular'),(6,7,'D','regular'),(6,8,'D','regular'),(6,9,'D','regular'),(6,10,'D','regular'),(6,11,'D','regular'),(6,12,'D','regular'),(6,13,'D','regular'),(6,14,'D','regular'),
(6,1,'E','regular'),(6,2,'E','regular'),(6,3,'E','regular'),(6,4,'E','regular'),(6,5,'E','regular'),(6,6,'E','regular'),(6,7,'E','regular'),(6,8,'E','regular'),(6,9,'E','regular'),(6,10,'E','regular'),(6,11,'E','regular'),(6,12,'E','regular'),(6,13,'E','regular'),(6,14,'E','regular'),
(6,1,'LN','love_nest'),(6,2,'LN','love_nest'),(6,3,'LN','love_nest'),(6,4,'LN','love_nest'),(6,5,'LN','love_nest'),(6,6,'LN','love_nest'),(6,7,'LN','love_nest'),(6,8,'LN','love_nest'),(6,9,'LN','love_nest'),(6,10,'LN','love_nest'),(6,11,'LN','love_nest'),(6,12,'LN','love_nest'),(6,13,'LN','love_nest'),(6,14,'LN','love_nest'),
(6,1,'G','regular'),(6,2,'G','regular'),(6,3,'G','regular'),(6,4,'G','regular'),(6,5,'G','regular'),(6,6,'G','regular'),(6,7,'G','regular'),(6,8,'G','regular'),(6,9,'G','regular'),(6,10,'G','regular'),(6,11,'G','regular'),(6,12,'G','regular'),(6,13,'G','regular'),(6,14,'G','regular');

-- Cinema 7
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(7,1,'A','regular'),(7,2,'A','regular'),(7,3,'A','regular'),(7,4,'A','regular'),(7,5,'A','regular'),(7,6,'A','regular'),(7,7,'A','regular'),(7,8,'A','regular'),(7,9,'A','regular'),(7,10,'A','regular'),(7,11,'A','regular'),(7,12,'A','regular'),(7,13,'A','regular'),(7,14,'A','regular'),
(7,1,'B','regular'),(7,2,'B','regular'),(7,3,'B','regular'),(7,4,'B','regular'),(7,5,'B','regular'),(7,6,'B','regular'),(7,7,'B','regular'),(7,8,'B','regular'),(7,9,'B','regular'),(7,10,'B','regular'),(7,11,'B','regular'),(7,12,'B','regular'),(7,13,'B','regular'),(7,14,'B','regular'),
(7,1,'C','regular'),(7,2,'C','regular'),(7,3,'C','regular'),(7,4,'C','regular'),(7,5,'C','regular'),(7,6,'C','regular'),(7,7,'C','regular'),(7,8,'C','regular'),(7,9,'C','regular'),(7,10,'C','regular'),(7,11,'C','regular'),(7,12,'C','regular'),(7,13,'C','regular'),(7,14,'C','regular'),
(7,1,'D','regular'),(7,2,'D','regular'),(7,3,'D','regular'),(7,4,'D','regular'),(7,5,'D','regular'),(7,6,'D','regular'),(7,7,'D','regular'),(7,8,'D','regular'),(7,9,'D','regular'),(7,10,'D','regular'),(7,11,'D','regular'),(7,12,'D','regular'),(7,13,'D','regular'),(7,14,'D','regular'),
(7,1,'E','regular'),(7,2,'E','regular'),(7,3,'E','regular'),(7,4,'E','regular'),(7,5,'E','regular'),(7,6,'E','regular'),(7,7,'E','regular'),(7,8,'E','regular'),(7,9,'E','regular'),(7,10,'E','regular'),(7,11,'E','regular'),(7,12,'E','regular'),(7,13,'E','regular'),(7,14,'E','regular'),
(7,1,'LN','love_nest'),(7,2,'LN','love_nest'),(7,3,'LN','love_nest'),(7,4,'LN','love_nest'),(7,5,'LN','love_nest'),(7,6,'LN','love_nest'),(7,7,'LN','love_nest'),(7,8,'LN','love_nest'),(7,9,'LN','love_nest'),(7,10,'LN','love_nest'),(7,11,'LN','love_nest'),(7,12,'LN','love_nest'),(7,13,'LN','love_nest'),(7,14,'LN','love_nest'),
(7,1,'G','regular'),(7,2,'G','regular'),(7,3,'G','regular'),(7,4,'G','regular'),(7,5,'G','regular'),(7,6,'G','regular'),(7,7,'G','regular'),(7,8,'G','regular'),(7,9,'G','regular'),(7,10,'G','regular'),(7,11,'G','regular'),(7,12,'G','regular'),(7,13,'G','regular'),(7,14,'G','regular');

-- Cinema 8
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(8,1,'A','regular'),(8,2,'A','regular'),(8,3,'A','regular'),(8,4,'A','regular'),(8,5,'A','regular'),(8,6,'A','regular'),(8,7,'A','regular'),(8,8,'A','regular'),(8,9,'A','regular'),(8,10,'A','regular'),(8,11,'A','regular'),(8,12,'A','regular'),(8,13,'A','regular'),(8,14,'A','regular'),
(8,1,'B','regular'),(8,2,'B','regular'),(8,3,'B','regular'),(8,4,'B','regular'),(8,5,'B','regular'),(8,6,'B','regular'),(8,7,'B','regular'),(8,8,'B','regular'),(8,9,'B','regular'),(8,10,'B','regular'),(8,11,'B','regular'),(8,12,'B','regular'),(8,13,'B','regular'),(8,14,'B','regular'),
(8,1,'C','regular'),(8,2,'C','regular'),(8,3,'C','regular'),(8,4,'C','regular'),(8,5,'C','regular'),(8,6,'C','regular'),(8,7,'C','regular'),(8,8,'C','regular'),(8,9,'C','regular'),(8,10,'C','regular'),(8,11,'C','regular'),(8,12,'C','regular'),(8,13,'C','regular'),(8,14,'C','regular'),
(8,1,'D','regular'),(8,2,'D','regular'),(8,3,'D','regular'),(8,4,'D','regular'),(8,5,'D','regular'),(8,6,'D','regular'),(8,7,'D','regular'),(8,8,'D','regular'),(8,9,'D','regular'),(8,10,'D','regular'),(8,11,'D','regular'),(8,12,'D','regular'),(8,13,'D','regular'),(8,14,'D','regular'),
(8,1,'E','regular'),(8,2,'E','regular'),(8,3,'E','regular'),(8,4,'E','regular'),(8,5,'E','regular'),(8,6,'E','regular'),(8,7,'E','regular'),(8,8,'E','regular'),(8,9,'E','regular'),(8,10,'E','regular'),(8,11,'E','regular'),(8,12,'E','regular'),(8,13,'E','regular'),(8,14,'E','regular'),
(8,1,'LN','love_nest'),(8,2,'LN','love_nest'),(8,3,'LN','love_nest'),(8,4,'LN','love_nest'),(8,5,'LN','love_nest'),(8,6,'LN','love_nest'),(8,7,'LN','love_nest'),(8,8,'LN','love_nest'),(8,9,'LN','love_nest'),(8,10,'LN','love_nest'),(8,11,'LN','love_nest'),(8,12,'LN','love_nest'),(8,13,'LN','love_nest'),(8,14,'LN','love_nest'),
(8,1,'G','regular'),(8,2,'G','regular'),(8,3,'G','regular'),(8,4,'G','regular'),(8,5,'G','regular'),(8,6,'G','regular'),(8,7,'G','regular'),(8,8,'G','regular'),(8,9,'G','regular'),(8,10,'G','regular'),(8,11,'G','regular'),(8,12,'G','regular'),(8,13,'G','regular'),(8,14,'G','regular');

-- Cinema 9
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(9,1,'A','regular'),(9,2,'A','regular'),(9,3,'A','regular'),(9,4,'A','regular'),(9,5,'A','regular'),(9,6,'A','regular'),(9,7,'A','regular'),(9,8,'A','regular'),(9,9,'A','regular'),(9,10,'A','regular'),(9,11,'A','regular'),(9,12,'A','regular'),(9,13,'A','regular'),(9,14,'A','regular'),
(9,1,'B','regular'),(9,2,'B','regular'),(9,3,'B','regular'),(9,4,'B','regular'),(9,5,'B','regular'),(9,6,'B','regular'),(9,7,'B','regular'),(9,8,'B','regular'),(9,9,'B','regular'),(9,10,'B','regular'),(9,11,'B','regular'),(9,12,'B','regular'),(9,13,'B','regular'),(9,14,'B','regular'),
(9,1,'C','regular'),(9,2,'C','regular'),(9,3,'C','regular'),(9,4,'C','regular'),(9,5,'C','regular'),(9,6,'C','regular'),(9,7,'C','regular'),(9,8,'C','regular'),(9,9,'C','regular'),(9,10,'C','regular'),(9,11,'C','regular'),(9,12,'C','regular'),(9,13,'C','regular'),(9,14,'C','regular'),
(9,1,'D','regular'),(9,2,'D','regular'),(9,3,'D','regular'),(9,4,'D','regular'),(9,5,'D','regular'),(9,6,'D','regular'),(9,7,'D','regular'),(9,8,'D','regular'),(9,9,'D','regular'),(9,10,'D','regular'),(9,11,'D','regular'),(9,12,'D','regular'),(9,13,'D','regular'),(9,14,'D','regular'),
(9,1,'E','regular'),(9,2,'E','regular'),(9,3,'E','regular'),(9,4,'E','regular'),(9,5,'E','regular'),(9,6,'E','regular'),(9,7,'E','regular'),(9,8,'E','regular'),(9,9,'E','regular'),(9,10,'E','regular'),(9,11,'E','regular'),(9,12,'E','regular'),(9,13,'E','regular'),(9,14,'E','regular'),
(9,1,'LN','love_nest'),(9,2,'LN','love_nest'),(9,3,'LN','love_nest'),(9,4,'LN','love_nest'),(9,5,'LN','love_nest'),(9,6,'LN','love_nest'),(9,7,'LN','love_nest'),(9,8,'LN','love_nest'),(9,9,'LN','love_nest'),(9,10,'LN','love_nest'),(9,11,'LN','love_nest'),(9,12,'LN','love_nest'),(9,13,'LN','love_nest'),(9,14,'LN','love_nest'),
(9,1,'G','regular'),(9,2,'G','regular'),(9,3,'G','regular'),(9,4,'G','regular'),(9,5,'G','regular'),(9,6,'G','regular'),(9,7,'G','regular'),(9,8,'G','regular'),(9,9,'G','regular'),(9,10,'G','regular'),(9,11,'G','regular'),(9,12,'G','regular'),(9,13,'G','regular'),(9,14,'G','regular');

-- Cinema 10
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(10,1,'A','regular'),(10,2,'A','regular'),(10,3,'A','regular'),(10,4,'A','regular'),(10,5,'A','regular'),(10,6,'A','regular'),(10,7,'A','regular'),(10,8,'A','regular'),(10,9,'A','regular'),(10,10,'A','regular'),(10,11,'A','regular'),(10,12,'A','regular'),(10,13,'A','regular'),(10,14,'A','regular'),
(10,1,'B','regular'),(10,2,'B','regular'),(10,3,'B','regular'),(10,4,'B','regular'),(10,5,'B','regular'),(10,6,'B','regular'),(10,7,'B','regular'),(10,8,'B','regular'),(10,9,'B','regular'),(10,10,'B','regular'),(10,11,'B','regular'),(10,12,'B','regular'),(10,13,'B','regular'),(10,14,'B','regular'),
(10,1,'C','regular'),(10,2,'C','regular'),(10,3,'C','regular'),(10,4,'C','regular'),(10,5,'C','regular'),(10,6,'C','regular'),(10,7,'C','regular'),(10,8,'C','regular'),(10,9,'C','regular'),(10,10,'C','regular'),(10,11,'C','regular'),(10,12,'C','regular'),(10,13,'C','regular'),(10,14,'C','regular'),
(10,1,'D','regular'),(10,2,'D','regular'),(10,3,'D','regular'),(10,4,'D','regular'),(10,5,'D','regular'),(10,6,'D','regular'),(10,7,'D','regular'),(10,8,'D','regular'),(10,9,'D','regular'),(10,10,'D','regular'),(10,11,'D','regular'),(10,12,'D','regular'),(10,13,'D','regular'),(10,14,'D','regular'),
(10,1,'E','regular'),(10,2,'E','regular'),(10,3,'E','regular'),(10,4,'E','regular'),(10,5,'E','regular'),(10,6,'E','regular'),(10,7,'E','regular'),(10,8,'E','regular'),(10,9,'E','regular'),(10,10,'E','regular'),(10,11,'E','regular'),(10,12,'E','regular'),(10,13,'E','regular'),(10,14,'E','regular'),
(10,1,'LN','love_nest'),(10,2,'LN','love_nest'),(10,3,'LN','love_nest'),(10,4,'LN','love_nest'),(10,5,'LN','love_nest'),(10,6,'LN','love_nest'),(10,7,'LN','love_nest'),(10,8,'LN','love_nest'),(10,9,'LN','love_nest'),(10,10,'LN','love_nest'),(10,11,'LN','love_nest'),(10,12,'LN','love_nest'),(10,13,'LN','love_nest'),(10,14,'LN','love_nest'),
(10,1,'G','regular'),(10,2,'G','regular'),(10,3,'G','regular'),(10,4,'G','regular'),(10,5,'G','regular'),(10,6,'G','regular'),(10,7,'G','regular'),(10,8,'G','regular'),(10,9,'G','regular'),(10,10,'G','regular'),(10,11,'G','regular'),(10,12,'G','regular'),(10,13,'G','regular'),(10,14,'G','regular');

-- Cinema 11
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(11,1,'A','regular'),(11,2,'A','regular'),(11,3,'A','regular'),(11,4,'A','regular'),(11,5,'A','regular'),(11,6,'A','regular'),(11,7,'A','regular'),(11,8,'A','regular'),(11,9,'A','regular'),(11,10,'A','regular'),(11,11,'A','regular'),(11,12,'A','regular'),(11,13,'A','regular'),(11,14,'A','regular'),
(11,1,'B','regular'),(11,2,'B','regular'),(11,3,'B','regular'),(11,4,'B','regular'),(11,5,'B','regular'),(11,6,'B','regular'),(11,7,'B','regular'),(11,8,'B','regular'),(11,9,'B','regular'),(11,10,'B','regular'),(11,11,'B','regular'),(11,12,'B','regular'),(11,13,'B','regular'),(11,14,'B','regular'),
(11,1,'C','regular'),(11,2,'C','regular'),(11,3,'C','regular'),(11,4,'C','regular'),(11,5,'C','regular'),(11,6,'C','regular'),(11,7,'C','regular'),(11,8,'C','regular'),(11,9,'C','regular'),(11,10,'C','regular'),(11,11,'C','regular'),(11,12,'C','regular'),(11,13,'C','regular'),(11,14,'C','regular'),
(11,1,'D','regular'),(11,2,'D','regular'),(11,3,'D','regular'),(11,4,'D','regular'),(11,5,'D','regular'),(11,6,'D','regular'),(11,7,'D','regular'),(11,8,'D','regular'),(11,9,'D','regular'),(11,10,'D','regular'),(11,11,'D','regular'),(11,12,'D','regular'),(11,13,'D','regular'),(11,14,'D','regular'),
(11,1,'E','regular'),(11,2,'E','regular'),(11,3,'E','regular'),(11,4,'E','regular'),(11,5,'E','regular'),(11,6,'E','regular'),(11,7,'E','regular'),(11,8,'E','regular'),(11,9,'E','regular'),(11,10,'E','regular'),(11,11,'E','regular'),(11,12,'E','regular'),(11,13,'E','regular'),(11,14,'E','regular'),
(11,1,'LN','love_nest'),(11,2,'LN','love_nest'),(11,3,'LN','love_nest'),(11,4,'LN','love_nest'),(11,5,'LN','love_nest'),(11,6,'LN','love_nest'),(11,7,'LN','love_nest'),(11,8,'LN','love_nest'),(11,9,'LN','love_nest'),(11,10,'LN','love_nest'),(11,11,'LN','love_nest'),(11,12,'LN','love_nest'),(11,13,'LN','love_nest'),(11,14,'LN','love_nest'),
(11,1,'G','regular'),(11,2,'G','regular'),(11,3,'G','regular'),(11,4,'G','regular'),(11,5,'G','regular'),(11,6,'G','regular'),(11,7,'G','regular'),(11,8,'G','regular'),(11,9,'G','regular'),(11,10,'G','regular'),(11,11,'G','regular'),(11,12,'G','regular'),(11,13,'G','regular'),(11,14,'G','regular');

-- Cinema 12
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(12,1,'A','regular'),(12,2,'A','regular'),(12,3,'A','regular'),(12,4,'A','regular'),(12,5,'A','regular'),(12,6,'A','regular'),(12,7,'A','regular'),(12,8,'A','regular'),(12,9,'A','regular'),(12,10,'A','regular'),(12,11,'A','regular'),(12,12,'A','regular'),(12,13,'A','regular'),(12,14,'A','regular'),
(12,1,'B','regular'),(12,2,'B','regular'),(12,3,'B','regular'),(12,4,'B','regular'),(12,5,'B','regular'),(12,6,'B','regular'),(12,7,'B','regular'),(12,8,'B','regular'),(12,9,'B','regular'),(12,10,'B','regular'),(12,11,'B','regular'),(12,12,'B','regular'),(12,13,'B','regular'),(12,14,'B','regular'),
(12,1,'C','regular'),(12,2,'C','regular'),(12,3,'C','regular'),(12,4,'C','regular'),(12,5,'C','regular'),(12,6,'C','regular'),(12,7,'C','regular'),(12,8,'C','regular'),(12,9,'C','regular'),(12,10,'C','regular'),(12,11,'C','regular'),(12,12,'C','regular'),(12,13,'C','regular'),(12,14,'C','regular'),
(12,1,'D','regular'),(12,2,'D','regular'),(12,3,'D','regular'),(12,4,'D','regular'),(12,5,'D','regular'),(12,6,'D','regular'),(12,7,'D','regular'),(12,8,'D','regular'),(12,9,'D','regular'),(12,10,'D','regular'),(12,11,'D','regular'),(12,12,'D','regular'),(12,13,'D','regular'),(12,14,'D','regular'),
(12,1,'E','regular'),(12,2,'E','regular'),(12,3,'E','regular'),(12,4,'E','regular'),(12,5,'E','regular'),(12,6,'E','regular'),(12,7,'E','regular'),(12,8,'E','regular'),(12,9,'E','regular'),(12,10,'E','regular'),(12,11,'E','regular'),(12,12,'E','regular'),(12,13,'E','regular'),(12,14,'E','regular'),
(12,1,'LN','love_nest'),(12,2,'LN','love_nest'),(12,3,'LN','love_nest'),(12,4,'LN','love_nest'),(12,5,'LN','love_nest'),(12,6,'LN','love_nest'),(12,7,'LN','love_nest'),(12,8,'LN','love_nest'),(12,9,'LN','love_nest'),(12,10,'LN','love_nest'),(12,11,'LN','love_nest'),(12,12,'LN','love_nest'),(12,13,'LN','love_nest'),(12,14,'LN','love_nest'),
(12,1,'G','regular'),(12,2,'G','regular'),(12,3,'G','regular'),(12,4,'G','regular'),(12,5,'G','regular'),(12,6,'G','regular'),(12,7,'G','regular'),(12,8,'G','regular'),(12,9,'G','regular'),(12,10,'G','regular'),(12,11,'G','regular'),(12,12,'G','regular'),(12,13,'G','regular'),(12,14,'G','regular');

-- Cinema 13
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(13,1,'A','regular'),(13,2,'A','regular'),(13,3,'A','regular'),(13,4,'A','regular'),(13,5,'A','regular'),(13,6,'A','regular'),(13,7,'A','regular'),(13,8,'A','regular'),(13,9,'A','regular'),(13,10,'A','regular'),(13,11,'A','regular'),(13,12,'A','regular'),(13,13,'A','regular'),(13,14,'A','regular'),
(13,1,'B','regular'),(13,2,'B','regular'),(13,3,'B','regular'),(13,4,'B','regular'),(13,5,'B','regular'),(13,6,'B','regular'),(13,7,'B','regular'),(13,8,'B','regular'),(13,9,'B','regular'),(13,10,'B','regular'),(13,11,'B','regular'),(13,12,'B','regular'),(13,13,'B','regular'),(13,14,'B','regular'),
(13,1,'C','regular'),(13,2,'C','regular'),(13,3,'C','regular'),(13,4,'C','regular'),(13,5,'C','regular'),(13,6,'C','regular'),(13,7,'C','regular'),(13,8,'C','regular'),(13,9,'C','regular'),(13,10,'C','regular'),(13,11,'C','regular'),(13,12,'C','regular'),(13,13,'C','regular'),(13,14,'C','regular'),
(13,1,'D','regular'),(13,2,'D','regular'),(13,3,'D','regular'),(13,4,'D','regular'),(13,5,'D','regular'),(13,6,'D','regular'),(13,7,'D','regular'),(13,8,'D','regular'),(13,9,'D','regular'),(13,10,'D','regular'),(13,11,'D','regular'),(13,12,'D','regular'),(13,13,'D','regular'),(13,14,'D','regular'),
(13,1,'E','regular'),(13,2,'E','regular'),(13,3,'E','regular'),(13,4,'E','regular'),(13,5,'E','regular'),(13,6,'E','regular'),(13,7,'E','regular'),(13,8,'E','regular'),(13,9,'E','regular'),(13,10,'E','regular'),(13,11,'E','regular'),(13,12,'E','regular'),(13,13,'E','regular'),(13,14,'E','regular'),
(13,1,'LN','love_nest'),(13,2,'LN','love_nest'),(13,3,'LN','love_nest'),(13,4,'LN','love_nest'),(13,5,'LN','love_nest'),(13,6,'LN','love_nest'),(13,7,'LN','love_nest'),(13,8,'LN','love_nest'),(13,9,'LN','love_nest'),(13,10,'LN','love_nest'),(13,11,'LN','love_nest'),(13,12,'LN','love_nest'),(13,13,'LN','love_nest'),(13,14,'LN','love_nest'),
(13,1,'G','regular'),(13,2,'G','regular'),(13,3,'G','regular'),(13,4,'G','regular'),(13,5,'G','regular'),(13,6,'G','regular'),(13,7,'G','regular'),(13,8,'G','regular'),(13,9,'G','regular'),(13,10,'G','regular'),(13,11,'G','regular'),(13,12,'G','regular'),(13,13,'G','regular'),(13,14,'G','regular');

-- Cinema 14
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(14,1,'A','regular'),(14,2,'A','regular'),(14,3,'A','regular'),(14,4,'A','regular'),(14,5,'A','regular'),(14,6,'A','regular'),(14,7,'A','regular'),(14,8,'A','regular'),(14,9,'A','regular'),(14,10,'A','regular'),(14,11,'A','regular'),(14,12,'A','regular'),(14,13,'A','regular'),(14,14,'A','regular'),
(14,1,'B','regular'),(14,2,'B','regular'),(14,3,'B','regular'),(14,4,'B','regular'),(14,5,'B','regular'),(14,6,'B','regular'),(14,7,'B','regular'),(14,8,'B','regular'),(14,9,'B','regular'),(14,10,'B','regular'),(14,11,'B','regular'),(14,12,'B','regular'),(14,13,'B','regular'),(14,14,'B','regular'),
(14,1,'C','regular'),(14,2,'C','regular'),(14,3,'C','regular'),(14,4,'C','regular'),(14,5,'C','regular'),(14,6,'C','regular'),(14,7,'C','regular'),(14,8,'C','regular'),(14,9,'C','regular'),(14,10,'C','regular'),(14,11,'C','regular'),(14,12,'C','regular'),(14,13,'C','regular'),(14,14,'C','regular'),
(14,1,'D','regular'),(14,2,'D','regular'),(14,3,'D','regular'),(14,4,'D','regular'),(14,5,'D','regular'),(14,6,'D','regular'),(14,7,'D','regular'),(14,8,'D','regular'),(14,9,'D','regular'),(14,10,'D','regular'),(14,11,'D','regular'),(14,12,'D','regular'),(14,13,'D','regular'),(14,14,'D','regular'),
(14,1,'E','regular'),(14,2,'E','regular'),(14,3,'E','regular'),(14,4,'E','regular'),(14,5,'E','regular'),(14,6,'E','regular'),(14,7,'E','regular'),(14,8,'E','regular'),(14,9,'E','regular'),(14,10,'E','regular'),(14,11,'E','regular'),(14,12,'E','regular'),(14,13,'E','regular'),(14,14,'E','regular'),
(14,1,'LN','love_nest'),(14,2,'LN','love_nest'),(14,3,'LN','love_nest'),(14,4,'LN','love_nest'),(14,5,'LN','love_nest'),(14,6,'LN','love_nest'),(14,7,'LN','love_nest'),(14,8,'LN','love_nest'),(14,9,'LN','love_nest'),(14,10,'LN','love_nest'),(14,11,'LN','love_nest'),(14,12,'LN','love_nest'),(14,13,'LN','love_nest'),(14,14,'LN','love_nest'),
(14,1,'G','regular'),(14,2,'G','regular'),(14,3,'G','regular'),(14,4,'G','regular'),(14,5,'G','regular'),(14,6,'G','regular'),(14,7,'G','regular'),(14,8,'G','regular'),(14,9,'G','regular'),(14,10,'G','regular'),(14,11,'G','regular'),(14,12,'G','regular'),(14,13,'G','regular'),(14,14,'G','regular');

-- Cinema 15
INSERT INTO seats (cinema_id, seat_number, row, seat_type) VALUES
(15,1,'A','regular'),(15,2,'A','regular'),(15,3,'A','regular'),(15,4,'A','regular'),(15,5,'A','regular'),(15,6,'A','regular'),(15,7,'A','regular'),(15,8,'A','regular'),(15,9,'A','regular'),(15,10,'A','regular'),(15,11,'A','regular'),(15,12,'A','regular'),(15,13,'A','regular'),(15,14,'A','regular'),
(15,1,'B','regular'),(15,2,'B','regular'),(15,3,'B','regular'),(15,4,'B','regular'),(15,5,'B','regular'),(15,6,'B','regular'),(15,7,'B','regular'),(15,8,'B','regular'),(15,9,'B','regular'),(15,10,'B','regular'),(15,11,'B','regular'),(15,12,'B','regular'),(15,13,'B','regular'),(15,14,'B','regular'),
(15,1,'C','regular'),(15,2,'C','regular'),(15,3,'C','regular'),(15,4,'C','regular'),(15,5,'C','regular'),(15,6,'C','regular'),(15,7,'C','regular'),(15,8,'C','regular'),(15,9,'C','regular'),(15,10,'C','regular'),(15,11,'C','regular'),(15,12,'C','regular'),(15,13,'C','regular'),(15,14,'C','regular'),
(15,1,'D','regular'),(15,2,'D','regular'),(15,3,'D','regular'),(15,4,'D','regular'),(15,5,'D','regular'),(15,6,'D','regular'),(15,7,'D','regular'),(15,8,'D','regular'),(15,9,'D','regular'),(15,10,'D','regular'),(15,11,'D','regular'),(15,12,'D','regular'),(15,13,'D','regular'),(15,14,'D','regular'),
(15,1,'E','regular'),(15,2,'E','regular'),(15,3,'E','regular'),(15,4,'E','regular'),(15,5,'E','regular'),(15,6,'E','regular'),(15,7,'E','regular'),(15,8,'E','regular'),(15,9,'E','regular'),(15,10,'E','regular'),(15,11,'E','regular'),(15,12,'E','regular'),(15,13,'E','regular'),(15,14,'E','regular'),
(15,1,'LN','love_nest'),(15,2,'LN','love_nest'),(15,3,'LN','love_nest'),(15,4,'LN','love_nest'),(15,5,'LN','love_nest'),(15,6,'LN','love_nest'),(15,7,'LN','love_nest'),(15,8,'LN','love_nest'),(15,9,'LN','love_nest'),(15,10,'LN','love_nest'),(15,11,'LN','love_nest'),(15,12,'LN','love_nest'),(15,13,'LN','love_nest'),(15,14,'LN','love_nest'),
(15,1,'G','regular'),(15,2,'G','regular'),(15,3,'G','regular'),(15,4,'G','regular'),(15,5,'G','regular'),(15,6,'G','regular'),(15,7,'G','regular'),(15,8,'G','regular'),(15,9,'G','regular'),(15,10,'G','regular'),(15,11,'G','regular'),(15,12,'G','regular'),(15,13,'G','regular'),(15,14,'G','regular');


-- ============================================================
-- 12. USERS (25 user — sama seperti sebelumnya)
-- ============================================================
INSERT INTO users (id, email, password, first_name, last_name, phone, photo, location_id, isActive, role) VALUES
(1,  'admin@tickitz.id',           '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y', 'Super',    'Admin',       '08100000001', 'https://storage.tickitz.id/users/admin1.jpg',   1,  true,  'admin'),
(2,  'budi.admin@tickitz.id',      '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y', 'Budi',     'Santoso',     '08100000002', 'https://storage.tickitz.id/users/admin2.jpg',   2,  true,  'admin'),
(3,  'sari.admin@tickitz.id',      '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y', 'Sari',     'Wulandari',   '08100000003', 'https://storage.tickitz.id/users/admin3.jpg',   3,  true,  'admin'),
(4,  'andi.pratama@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Andi',     'Pratama',     '08111111101', 'https://storage.tickitz.id/users/user4.jpg',    1,  true,  'user'),
(5,  'rina.lestari@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Rina',     'Lestari',     '08111111102', 'https://storage.tickitz.id/users/user5.jpg',    1,  true,  'user'),
(6,  'doni.kurniawan@yahoo.com',   '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Doni',     'Kurniawan',   '08111111103', 'https://storage.tickitz.id/users/user6.jpg',    2,  true,  'user'),
(7,  'bagas.wijaya@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Bagas',    'Wijaya',      '08111111104', 'https://storage.tickitz.id/users/user7.jpg',    3,  true,  'user'),
(8,  'nadia.putri@gmail.com',      '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Nadia',    'Putri',       '08111111105', 'https://storage.tickitz.id/users/user8.jpg',    3,  true,  'user'),
(9,  'reza.maulana@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Reza',     'Maulana',     '08111111106', 'https://storage.tickitz.id/users/user9.jpg',    4,  true,  'user'),
(10, 'fitri.handayani@gmail.com',  '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Fitri',    'Handayani',   '08111111107', 'https://storage.tickitz.id/users/user10.jpg',   1,  true,  'user'),
(11, 'hendra.gunawan@gmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Hendra',   'Gunawan',     '08111111108', 'https://storage.tickitz.id/users/user11.jpg',   5,  true,  'user'),
(12, 'dewi.anggraeni@gmail.com',   '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Dewi',     'Anggraeni',   '08111111109', 'https://storage.tickitz.id/users/user12.jpg',   1,  true,  'user'),
(13, 'agus.setiawan@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Agus',     'Setiawan',    '08111111110', 'https://storage.tickitz.id/users/user13.jpg',   8,  true,  'user'),
(14, 'sinta.permata@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Sinta',    'Permata',     '08111111111', 'https://storage.tickitz.id/users/user14.jpg',   9,  true,  'user'),
(15, 'lisa.amelia@gmail.com',      '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Lisa',     'Amelia',      '08111111112', 'https://storage.tickitz.id/users/user15.jpg',   2,  true,  'user'),
(16, 'tommy.wirawan@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Tommy',    'Wirawan',     '08111111113', 'https://storage.tickitz.id/users/user16.jpg',   3,  true,  'user'),
(17, 'yuni.astuti@gmail.com',      '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Yuni',     'Astuti',      '08111111114', 'https://storage.tickitz.id/users/user17.jpg',   1,  true,  'user'),
(18, 'kevin.pratama@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Kevin',    'Pratama',     '08111111115', 'https://storage.tickitz.id/users/user18.jpg',   6,  true,  'user'),
(19, 'amanda.putri@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Amanda',   'Putri',       '08111111116', 'https://storage.tickitz.id/users/user19.jpg',  11,  true,  'user'),
(20, 'rizky.hamdani@gmail.com',    '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Rizky',    'Hamdani',     '08111111117', 'https://storage.tickitz.id/users/user20.jpg',   1,  true,  'user'),
(21, 'clara.monica@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Clara',    'Monica',      '08111111118', 'https://storage.tickitz.id/users/user21.jpg',   2,  true,  'user'),
(22, 'dimas.arif@gmail.com',       '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Dimas',    'Arif',        '08111111119', 'https://storage.tickitz.id/users/user22.jpg',   7,  true,  'user'),
(23, 'maya.sari@gmail.com',        '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Maya',     'Sari',        '08111111120', NULL,                                            2,  false, 'user'),
(24, 'farhan.akbar@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Farhan',   'Akbar',       '08111111121', NULL,                                            1,  false, 'user'),
(25, 'putri.rahayu@gmail.com',     '$argon2id$v=19$m=65536,t=2,p=1$xIJYsW0jYK6HncD8JVUsZA$wGP3m49Qm7OrrFOAu/baOcutE3s/sUb0LaoGLy7B13Y',  'Putri',    'Rahayu',      '08111111122', NULL,                                           11,  false, 'user');


-- ============================================================
-- 13. SHOWTIMES — menggunakan CURRENT_DATE agar selalu relevan
--
--  ✅ NOW SHOWING (movie_id 1–10):
--     Jadwal di CURRENT_DATE, CURRENT_DATE-1, CURRENT_DATE-2, CURRENT_DATE+1, +2
--     → Query WHERE date >= CURRENT_DATE akan menemukan jadwal hari ini & besok
--
--  ❌ UPCOMING (movie_id 11–18):
--     Tidak ada showtime — film belum dijadwalkan tayang
--     → Tetap muncul di halaman "Upcoming" berdasarkan release_date > CURRENT_DATE
--
--  Pemetaan seat_id per cinema (sama seperti v2):
--    Cinema 1: 1–98    Cinema 2: 99–196   Cinema 3: 197–294
--    Cinema 4: 295–392 Cinema 5: 393–490  Cinema 6: 491–588
--    Cinema 7: 589–686 Cinema 8: 687–784  Cinema 9: 785–882
--    Cinema 10: 883–980  Cinema 11: 981–1078  Cinema 12: 1079–1176
-- ============================================================
INSERT INTO showtimes (id, movie_id, cinema_id, date, time, price) VALUES

-- ── FILM 1: Monster Pabrik Rambut ────────────────────────────
(1,  1, 1, CURRENT_DATE - 2, '08:30', 65000),
(2,  1, 1, CURRENT_DATE - 2, '16:00', 65000),
(3,  1, 1, CURRENT_DATE - 1, '13:30', 65000),
(4,  1, 1, CURRENT_DATE,     '20:00', 70000),
(5,  1, 1, CURRENT_DATE,     '21:30', 70000),
(6,  1, 2, CURRENT_DATE,     '18:30', 65000),
(7,  1, 4, CURRENT_DATE + 1, '20:00', 60000),
(8,  1, 7, CURRENT_DATE + 1, '21:30', 60000),

-- ── FILM 2: Kucing Hitam ─────────────────────────────────────
(9,  2, 1, CURRENT_DATE - 2, '21:30', 65000),
(10, 2, 2, CURRENT_DATE - 1, '21:30', 60000),
(11, 2, 1, CURRENT_DATE,     '21:30', 70000),
(12, 2, 3, CURRENT_DATE + 1, '20:00', 60000),
(13, 2, 8, CURRENT_DATE + 1, '21:30', 55000),

-- ── FILM 3: Warkop DKI: Viralin Doooong!! ───────────────────
(14, 3, 1, CURRENT_DATE - 2, '11:00', 65000),
(15, 3, 1, CURRENT_DATE - 1, '13:30', 65000),
(16, 3, 1, CURRENT_DATE,     '11:00', 65000),
(17, 3, 1, CURRENT_DATE,     '16:00', 65000),
(18, 3, 2, CURRENT_DATE,     '13:30', 60000),
(19, 3, 4, CURRENT_DATE + 1, '11:00', 60000),
(20, 3, 6, CURRENT_DATE + 1, '16:00', 55000),
(21, 3, 11,CURRENT_DATE + 2, '13:30', 55000),

-- ── FILM 4: Colony ───────────────────────────────────────────
(22, 4, 1, CURRENT_DATE - 2, '18:30', 75000),
(23, 4, 1, CURRENT_DATE - 1, '20:00', 75000),
(24, 4, 1, CURRENT_DATE,     '18:30', 80000),
(25, 4, 2, CURRENT_DATE,     '20:00', 75000),
(26, 4, 4, CURRENT_DATE + 1, '18:30', 70000),
(27, 4, 9, CURRENT_DATE + 2, '20:00', 65000),

-- ── FILM 5: Masters of the Universe ─────────────────────────
(28, 5, 1, CURRENT_DATE - 1, '08:30', 85000),
(29, 5, 1, CURRENT_DATE,     '08:30', 90000),
(30, 5, 1, CURRENT_DATE,     '13:30', 90000),
(31, 5, 2, CURRENT_DATE,     '16:00', 85000),
(32, 5, 4, CURRENT_DATE + 1, '13:30', 80000),
(33, 5, 7, CURRENT_DATE + 1, '16:00', 80000),

-- ── FILM 6: Nobody Loves Kay ─────────────────────────────────
(34, 6, 1, CURRENT_DATE - 2, '16:00', 65000),
(35, 6, 1, CURRENT_DATE - 1, '16:00', 65000),
(36, 6, 1, CURRENT_DATE,     '16:00', 70000),
(37, 6, 3, CURRENT_DATE,     '13:30', 60000),
(38, 6, 5, CURRENT_DATE + 1, '16:00', 60000),

-- ── FILM 7: Jangan Buang Ibu ─────────────────────────────────
(39, 7, 1, CURRENT_DATE - 1, '13:30', 65000),
(40, 7, 1, CURRENT_DATE,     '11:00', 65000),
(41, 7, 2, CURRENT_DATE,     '13:30', 60000),
(42, 7, 6, CURRENT_DATE + 1, '11:00', 55000),
(43, 7,10, CURRENT_DATE + 2, '13:30', 55000),

-- ── FILM 8: Garuda di Dadaku ─────────────────────────────────
(44, 8, 1, CURRENT_DATE,     '11:00', 55000),
(45, 8, 3, CURRENT_DATE,     '13:30', 50000),
(46, 8, 4, CURRENT_DATE + 1, '11:00', 50000),
(47, 8,12, CURRENT_DATE + 1, '13:30', 50000),

-- ── FILM 9: Tanah Runtuh ─────────────────────────────────────
(48, 9, 1, CURRENT_DATE,     '13:30', 70000),
(49, 9, 2, CURRENT_DATE,     '16:00', 65000),
(50, 9, 9, CURRENT_DATE + 1, '16:00', 60000),
(51, 9,13, CURRENT_DATE + 2, '16:00', 60000),

-- ── FILM 10: The Furious ─────────────────────────────────────
(52,10, 1, CURRENT_DATE,     '20:00', 80000),
(53,10, 2, CURRENT_DATE,     '18:30', 75000),
(54,10, 4, CURRENT_DATE + 1, '20:00', 70000),
(55,10, 5, CURRENT_DATE + 1, '18:30', 70000),
(56,10,14, CURRENT_DATE + 2, '20:00', 65000);

-- ✋ NOTE: Film UPCOMING (id 11–18) tidak memiliki showtime
--    Query di aplikasi untuk "Now Showing" gunakan:
--    WHERE s.date >= CURRENT_DATE AND m.release_date <= CURRENT_DATE
--    Query untuk "Upcoming" gunakan:
--    WHERE m.release_date > CURRENT_DATE


-- ============================================================
-- 14. BOOKINGS
--
--  ⚠️  PENTING — LOGIKA BISNIS:
--  Booking HANYA untuk film yang sudah tayang (showtime sudah lewat / hari ini)
--  Film Upcoming = BELUM bisa dibooking → TIDAK ADA booking untuk movie_id 11–18
--
--  Showtimes yang sudah lewat (cocok untuk booking lama / not_active):
--    id 1,2 (film 1 CURRENT_DATE-2), id 9 (film 2), id 14,15 (film 3)
--    id 22,23 (film 4), id 28 (film 5), id 34,35 (film 6), id 39 (film 7)
--
--  Showtimes hari ini / besok (untuk booking active+paid dan active+not_paid):
--    id 3–8, 10–13, 16–21, 24–27, 29–33, 36–38, 40–47, 48–56
-- ============================================================
INSERT INTO bookings (id, user_id, showtime_id, status_ticket, status_paid, quantity, created_at, updated_at) VALUES

-- ── SKENARIO A: active + paid (tiket valid, nonton nanti) ───
-- Booking untuk showtime hari ini & besok
(1,  4,  4,  'active', 'paid', 2,  NOW() - INTERVAL '3 hours',  NOW() - INTERVAL '3 hours' + INTERVAL '10 minutes'),
(2,  5,  16, 'active', 'paid', 3,  NOW() - INTERVAL '4 hours',  NOW() - INTERVAL '4 hours' + INTERVAL '8 minutes'),
(3,  6,  24, 'active', 'paid', 2,  NOW() - INTERVAL '5 hours',  NOW() - INTERVAL '5 hours' + INTERVAL '7 minutes'),
(4,  7,  29, 'active', 'paid', 4,  NOW() - INTERVAL '6 hours',  NOW() - INTERVAL '6 hours' + INTERVAL '9 minutes'),
(5,  8,  36, 'active', 'paid', 2,  NOW() - INTERVAL '2 hours',  NOW() - INTERVAL '2 hours' + INTERVAL '6 minutes'),
(6,  9,  40, 'active', 'paid', 2,  NOW() - INTERVAL '7 hours',  NOW() - INTERVAL '7 hours' + INTERVAL '8 minutes'),
(7,  10, 44, 'active', 'paid', 2,  NOW() - INTERVAL '8 hours',  NOW() - INTERVAL '8 hours' + INTERVAL '7 minutes'),
(8,  11, 48, 'active', 'paid', 3,  NOW() - INTERVAL '9 hours',  NOW() - INTERVAL '9 hours' + INTERVAL '9 minutes'),
(9,  12, 52, 'active', 'paid', 2,  NOW() - INTERVAL '1 hours',  NOW() - INTERVAL '1 hours' + INTERVAL '6 minutes'),
(10, 13, 5,  'active', 'paid', 2,  NOW() - INTERVAL '2 hours',  NOW() - INTERVAL '2 hours' + INTERVAL '8 minutes'),
(11, 14, 11, 'active', 'paid', 1,  NOW() - INTERVAL '3 hours',  NOW() - INTERVAL '3 hours' + INTERVAL '5 minutes'),
(12, 15, 17, 'active', 'paid', 2,  NOW() - INTERVAL '4 hours',  NOW() - INTERVAL '4 hours' + INTERVAL '7 minutes'),
(13, 16, 25, 'active', 'paid', 3,  NOW() - INTERVAL '5 hours',  NOW() - INTERVAL '5 hours' + INTERVAL '8 minutes'),
(14, 17, 30, 'active', 'paid', 2,  NOW() - INTERVAL '6 hours',  NOW() - INTERVAL '6 hours' + INTERVAL '6 minutes'),
(15, 18, 37, 'active', 'paid', 2,  NOW() - INTERVAL '7 hours',  NOW() - INTERVAL '7 hours' + INTERVAL '7 minutes'),
(16, 19, 41, 'active', 'paid', 2,  NOW() - INTERVAL '8 hours',  NOW() - INTERVAL '8 hours' + INTERVAL '8 minutes'),
(17, 20, 45, 'active', 'paid', 3,  NOW() - INTERVAL '9 hours',  NOW() - INTERVAL '9 hours' + INTERVAL '9 minutes'),
(18, 21, 49, 'active', 'paid', 2,  NOW() - INTERVAL '2 hours',  NOW() - INTERVAL '2 hours' + INTERVAL '6 minutes'),
(19, 22, 53, 'active', 'paid', 2,  NOW() - INTERVAL '3 hours',  NOW() - INTERVAL '3 hours' + INTERVAL '7 minutes'),
(20, 4,  32, 'active', 'paid', 2,  NOW() - INTERVAL '4 hours',  NOW() - INTERVAL '4 hours' + INTERVAL '8 minutes'),
(21, 5,  46, 'active', 'paid', 3,  NOW() - INTERVAL '5 hours',  NOW() - INTERVAL '5 hours' + INTERVAL '9 minutes'),
(22, 7,  54, 'active', 'paid', 2,  NOW() - INTERVAL '6 hours',  NOW() - INTERVAL '6 hours' + INTERVAL '7 minutes'),

-- ── SKENARIO B: not_active + paid (tiket sudah dipakai) ─────
-- Booking untuk showtime yang sudah lewat (CURRENT_DATE-1 atau -2)
(23, 4,  1,  'not_active', 'paid', 2, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days' + INTERVAL '10 minutes'),
(24, 5,  9,  'not_active', 'paid', 2, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days' + INTERVAL '8 minutes'),
(25, 8,  14, 'not_active', 'paid', 3, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days' + INTERVAL '9 minutes'),
(26, 9,  22, 'not_active', 'paid', 2, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days' + INTERVAL '7 minutes'),
(27,10,  28, 'not_active', 'paid', 2, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days' + INTERVAL '6 minutes'),
(28,11,  34, 'not_active', 'paid', 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days' + INTERVAL '5 minutes'),
(29,12,  39, 'not_active', 'paid', 2, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days' + INTERVAL '8 minutes'),

-- ── SKENARIO C: active + not_paid (menunggu pembayaran) ─────
(30, 6,  6,  'active', 'not_paid', 2, NOW() - INTERVAL '30 minutes', NULL),
(31,13,  18, 'active', 'not_paid', 2, NOW() - INTERVAL '45 minutes', NULL),
(32,14,  26, 'active', 'not_paid', 3, NOW() - INTERVAL '20 minutes', NULL),
(33,15,  31, 'active', 'not_paid', 2, NOW() - INTERVAL '60 minutes', NULL),
(34,16,  38, 'active', 'not_paid', 2, NOW() - INTERVAL '15 minutes', NULL),
(35,17,  42, 'active', 'not_paid', 2, NOW() - INTERVAL '50 minutes', NULL),
(36,18,  50, 'active', 'not_paid', 3, NOW() - INTERVAL '25 minutes', NULL),
(37,20,  55, 'active', 'not_paid', 2, NOW() - INTERVAL '35 minutes', NULL);


-- ============================================================
-- 15. TRANSACTIONS
-- ============================================================
INSERT INTO transactions (id, booking_id, payment_method_id, virtual_rek, total_price, status, qr_code) VALUES
-- COMPLETED (booking 1–22)
(1,  1,  1, 88110001, 140000,  'completed', 'https://storage.tickitz.id/qr/TXN001.png'),
(2,  2,  4, 88110002, 195000,  'completed', 'https://storage.tickitz.id/qr/TXN002.png'),
(3,  3,  2, 88110003, 150000,  'completed', 'https://storage.tickitz.id/qr/TXN003.png'),
(4,  4,  1, 88110004, 360000,  'completed', 'https://storage.tickitz.id/qr/TXN004.png'),
(5,  5,  3, 88110005, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN005.png'),
(6,  6,  2, 88110006, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN006.png'),
(7,  7,  4, 88110007, 110000,  'completed', 'https://storage.tickitz.id/qr/TXN007.png'),
(8,  8,  1, 88110008, 210000,  'completed', 'https://storage.tickitz.id/qr/TXN008.png'),
(9,  9,  5, 88110009, 160000,  'completed', 'https://storage.tickitz.id/qr/TXN009.png'),
(10,10,  7, 88110010, 140000,  'completed', 'https://storage.tickitz.id/qr/TXN010.png'),
(11,11,  2, 88110011,  70000,  'completed', 'https://storage.tickitz.id/qr/TXN011.png'),
(12,12,  4, 88110012, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN012.png'),
(13,13,  1, 88110013, 225000,  'completed', 'https://storage.tickitz.id/qr/TXN013.png'),
(14,14,  6, 88110014, 170000,  'completed', 'https://storage.tickitz.id/qr/TXN014.png'),
(15,15,  3, 88110015, 120000,  'completed', 'https://storage.tickitz.id/qr/TXN015.png'),
(16,16,  1, 88110016, 110000,  'completed', 'https://storage.tickitz.id/qr/TXN016.png'),
(17,17,  4, 88110017, 165000,  'completed', 'https://storage.tickitz.id/qr/TXN017.png'),
(18,18,  2, 88110018, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN018.png'),
(19,19,  8, 88110019, 150000,  'completed', 'https://storage.tickitz.id/qr/TXN019.png'),
(20,20,  5, 88110020, 180000,  'completed', 'https://storage.tickitz.id/qr/TXN020.png'),
(21,21,  1, 88110021, 165000,  'completed', 'https://storage.tickitz.id/qr/TXN021.png'),
(22,22,  4, 88110022, 150000,  'completed', 'https://storage.tickitz.id/qr/TXN022.png'),
-- COMPLETED — tiket lama (booking 23–29)
(23,23,  2, 88110023, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN023.png'),
(24,24,  1, 88110024, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN024.png'),
(25,25,  4, 88110025, 195000,  'completed', 'https://storage.tickitz.id/qr/TXN025.png'),
(26,26,  3, 88110026, 150000,  'completed', 'https://storage.tickitz.id/qr/TXN026.png'),
(27,27,  5, 88110027, 170000,  'completed', 'https://storage.tickitz.id/qr/TXN027.png'),
(28,28,  6, 88110028,  65000,  'completed', 'https://storage.tickitz.id/qr/TXN028.png'),
(29,29,  1, 88110029, 130000,  'completed', 'https://storage.tickitz.id/qr/TXN029.png'),
-- PENDING (booking 30–36)
(30,30,  4, 88120001, 130000,  'pending', NULL),
(31,31,  1, 88120002, 150000,  'pending', NULL),
(32,32,  5, 88120003, 240000,  'pending', NULL),
(33,33,  2, 88120004, 180000,  'pending', NULL),
(34,34,  4, 88120005, 120000,  'pending', NULL),
(35,35,  3, 88120006, 130000,  'pending', NULL),
(36,36,  1, 88120007, 195000,  'pending', NULL),
-- FAILED
(37,37,  5, 88130001, 140000,  'failed',  NULL);


-- ============================================================
-- 16. BOOKING_SEATS
--
--  Peta seat_id (sama seperti v2):
--  Cinema 1: 1–98     Cinema 2: 99–196   Cinema 3: 197–294
--  Cinema 4: 295–392  Cinema 5: 393–490  Cinema 6: 491–588
--  Cinema 7: 589–686  Cinema 8: 687–784  Cinema 9: 785–882
--  ┌──────────────────────────────────────────────┐
--  │ Row A: seat 1–14  | B: 15–28  | C: 29–42    │
--  │ Row D: 43–56      | E: 57–70  | LN: 71–84   │
--  │ Row G: 85–98                                  │
--  └──────────────────────────────────────────────┘
-- ============================================================

-- Booking 1 (user 4, showtime 4 = Film 1 CURRENT_DATE, cinema 1, qty 2) → A1,A2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (1, 1, 4), (1, 2, 4);
-- Booking 2 (user 5, showtime 16 = Film 3, cinema 1, qty 3) → B1,B2,B3
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (2, 15, 16), (2, 16, 16), (2, 17, 16);
-- Booking 3 (user 6, showtime 24 = Film 4, cinema 1, qty 2) → C1,C2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (3, 29, 24), (3, 30, 24);
-- Booking 4 (user 7, showtime 29 = Film 5, cinema 1, qty 4) → D1,D2,D3,D4
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (4, 43, 29), (4, 44, 29), (4, 45, 29), (4, 46, 29);
-- Booking 5 (user 8, showtime 36 = Film 6, cinema 1, qty 2) → E1,E2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (5, 57, 36), (5, 58, 36);
-- Booking 6 (user 9, showtime 40 = Film 7, cinema 1, qty 2) → A3,A4
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (6, 3, 40), (6, 4, 40);
-- Booking 7 (user 10, showtime 44 = Film 8, cinema 1, qty 2) → LN1,LN2 (Love Nest!)
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (7, 71, 44), (7, 72, 44);
-- Booking 8 (user 11, showtime 48 = Film 9, cinema 1, qty 3) → G1,G2,G3
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (8, 85, 48), (8, 86, 48), (8, 87, 48);
-- Booking 9 (user 12, showtime 52 = Film 10, cinema 1, qty 2) → A5,A6
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (9, 5, 52), (9, 6, 52);
-- Booking 10 (user 13, showtime 5 = Film 1 late show, cinema 1, qty 2) → B5,B6
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (10, 19, 5), (10, 20, 5);
-- Booking 11 (user 14, showtime 11 = Film 2, cinema 1, qty 1) → C5
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (11, 33, 11);
-- Booking 12 (user 15, showtime 17 = Film 3, cinema 1, qty 2) → D7,D8
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (12, 49, 17), (12, 50, 17);
-- Booking 13 (user 16, showtime 25 = Film 4 hiflix, cinema 2, qty 3) → A1,A2,A3
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (13, 99, 25), (13, 100, 25), (13, 101, 25);
-- Booking 14 (user 17, showtime 30 = Film 5 hiflix, cinema 2, qty 2) → B3,B4
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (14, 116, 30), (14, 117, 30);
-- Booking 15 (user 18, showtime 37 = Film 6 Bandung CineOne, cinema 3, qty 2) → A1,A2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (15, 197, 37), (15, 198, 37);
-- Booking 16 (user 19, showtime 41 = Film 7 hiflix, cinema 2, qty 2) → LN3,LN4 (Love Nest!)
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (16, 171, 41), (16, 172, 41);
-- Booking 17 (user 20, showtime 45 = Film 8 CineOne Bali, cinema 12, qty 3) → A1,A2,A3
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (17, 1079, 45), (17, 1080, 45), (17, 1081, 45);
-- Booking 18 (user 21, showtime 49 = Film 9 hiflix, cinema 2, qty 2) → C3,C4
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (18, 127, 49), (18, 128, 49);
-- Booking 19 (user 22, showtime 53 = Film 10 hiflix, cinema 2, qty 2) → D5,D6
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (19, 145, 53), (19, 146, 53);
-- Booking 20 (user 4, showtime 32 = Film 5 ebv Bandung, cinema 7, qty 2) → A3,A4
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (20, 591, 32), (20, 592, 32);
-- Booking 21 (user 5, showtime 46 = Film 8 ebv Surabaya, cinema 4, qty 3) → B1,B2,B3
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (21, 309, 46), (21, 310, 46), (21, 311, 46);
-- Booking 22 (user 7, showtime 54 = Film 10 ebv Surabaya, cinema 4, qty 2) → E1,E2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (22, 351, 54), (22, 352, 54);

-- Tiket lama / not_active
-- Booking 23 (user 4, showtime 1 = Film 1, CURRENT_DATE-2, qty 2) → A7,A8
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (23, 7, 1), (23, 8, 1);
-- Booking 24 (user 5, showtime 9 = Film 2, CURRENT_DATE-2, qty 2) → A9,A10
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (24, 9, 9), (24, 10, 9);
-- Booking 25 (user 8, showtime 14 = Film 3, CURRENT_DATE-2, qty 3) → C7,C8,C9
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (25, 35, 14), (25, 36, 14), (25, 37, 14);
-- Booking 26 (user 9, showtime 22 = Film 4, CURRENT_DATE-2, qty 2) → E5,E6
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (26, 61, 22), (26, 62, 22);
-- Booking 27 (user 10, showtime 28 = Film 5, CURRENT_DATE-1, qty 2) → G5,G6
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (27, 89, 28), (27, 90, 28);
-- Booking 28 (user 11, showtime 34 = Film 6, CURRENT_DATE-2, qty 1) → D3
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (28, 45, 34);
-- Booking 29 (user 12, showtime 39 = Film 7, CURRENT_DATE-1, qty 2) → LN5,LN6 (Love Nest!)
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (29, 75, 39), (29, 76, 39);

-- Booking pending / not_paid (kursi di-hold)
-- Booking 30 (user 6, showtime 6 = Film 1 hiflix, qty 2) → A1,A2 cinema2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (30, 99, 6), (30, 100, 6);
-- Booking 31 (user 13, showtime 18 = Film 3 hiflix, qty 2) → B5,B6 cinema2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (31, 103, 18), (31, 104, 18);
-- Booking 32 (user 14, showtime 26 = Film 4 ebv Surabaya, qty 3) → C1,C2,C3 cinema4
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (32, 323, 26), (32, 324, 26), (32, 325, 26);
-- Booking 33 (user 15, showtime 31 = Film 5 hiflix, qty 2) → D9,D10 cinema2
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (33, 151, 31), (33, 152, 31);
-- Booking 34 (user 16, showtime 38 = Film 6 hiflix Surabaya, qty 2) → A5,A6 cinema5
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (34, 397, 38), (34, 398, 38);
-- Booking 35 (user 17, showtime 42 = Film 7 CineOne21 Surabaya, qty 2) → B1,B2 cinema6
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (35, 505, 42), (35, 506, 42);
-- Booking 36 (user 18, showtime 50 = Film 9 CineOne21 Medan, qty 3) → A1,A2,A3 cinema9
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (36, 785, 50), (36, 786, 50), (36, 787, 50);
-- Booking 37 (user 20, showtime 55 = Film 10 hiflix Surabaya, qty 2) → B7,B8 cinema5
INSERT INTO booking_seats (booking_id, seat_id, showtime_id) VALUES (37, 406, 55), (37, 407, 55);


-- ============================================================
-- SEQUENCE RESET
-- ============================================================
SELECT setval('locations_id_seq',       (SELECT MAX(id) FROM locations));
SELECT setval('payment_methods_id_seq', (SELECT MAX(id) FROM payment_methods));
SELECT setval('genres_id_seq',          (SELECT MAX(id) FROM genres));
SELECT setval('directors_id_seq',       (SELECT MAX(id) FROM directors));
SELECT setval('casts_id_seq',           (SELECT MAX(id) FROM casts));
SELECT setval('movies_id_seq',          (SELECT MAX(id) FROM movies));
SELECT setval('cinemas_id_seq',         (SELECT MAX(id) FROM cinemas));
SELECT setval('users_id_seq',           (SELECT MAX(id) FROM users));
SELECT setval('showtimes_id_seq',       (SELECT MAX(id) FROM showtimes));
SELECT setval('bookings_id_seq',        (SELECT MAX(id) FROM bookings));
SELECT setval('transactions_id_seq',    (SELECT MAX(id) FROM transactions));
SELECT setval('booking_seats_id_seq',   (SELECT MAX(id) FROM booking_seats));


-- ================================================================
-- SUMMARY DATA v3.0
-- ================================================================
-- locations       : 12 kota
-- users           : 25 (22 user + 3 admin; 3 belum verifikasi)
-- payment_methods : 8
-- genres          : 12
-- directors       : 20
-- casts           : 30
-- movies          : 18 film ASLI Juni 2026
--                   └ Now Showing (id 1-10)  : Monster Pabrik Rambut, Kucing Hitam,
--                                              Warkop DKI Viralin, Colony,
--                                              Masters of the Universe, Nobody Loves Kay,
--                                              Jangan Buang Ibu, Garuda di Dadaku,
--                                              Tanah Runtuh, The Furious
--                   └ Upcoming    (id 11-18) : Toy Story 5, Supergirl, Disclosure Day,
--                                              Dukun Magang, Cerita Lila,
--                                              The Longest Wait, Minions & Monsters, Backrooms
-- cinemas         : 15 (ebv.id / hiflix / CineOne21)
-- seats           : 1,470 kursi (15 × 98)
-- showtimes       : 56 jadwal — HANYA untuk Now Showing
--                   Semua tanggal = CURRENT_DATE ± N days (selalu relevan!)
-- bookings        : 37
--                   └ active+paid     : 22 (tiket valid, jadwal hari ini/besok)
--                   └ not_active+paid : 7  (tiket lama/sudah nonton)
--                   └ active+not_paid : 8  (menunggu pembayaran)
-- transactions    : 37 (29 completed | 7 pending | 1 failed)
-- ⚠️  Film UPCOMING tidak memiliki showtime → tidak bisa dibooking
-- ================================================================