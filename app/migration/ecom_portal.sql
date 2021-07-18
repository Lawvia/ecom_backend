-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 18 Jul 2021 pada 18.25
-- Versi server: 10.4.13-MariaDB
-- Versi PHP: 7.4.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ecom_portal`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_items`
--

CREATE TABLE `ms_items` (
  `id` int(11) NOT NULL,
  `name` varchar(99) NOT NULL,
  `description` text NOT NULL,
  `price` varchar(45) NOT NULL,
  `stock` int(11) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `ms_items`
--

INSERT INTO `ms_items` (`id`, `name`, `description`, `price`, `stock`, `created`, `updated`) VALUES
(1, 'Pasta Gigi', 'pemutih gigi berlubang', '10000', 10, '2021-07-17 11:06:48', '2021-07-18 16:20:34'),
(2, 'Margarin', 'kualitas oke', '7000', 12, '2021-07-17 11:07:41', '2021-07-17 11:07:41'),
(3, 'Iphone XR', 'apple', '1200000', 8, '2021-07-17 11:08:21', '2021-07-18 16:20:34'),
(4, 'Dualshock 4', 'controller nintendo', '500000', 23, '2021-07-17 11:09:45', '2021-07-17 16:05:17'),
(5, 'Steam Deck', 'handheld gaming system', '2500000', 9, '2021-07-17 11:11:53', '2021-07-17 11:11:53');

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_user`
--

CREATE TABLE `ms_user` (
  `id` int(11) NOT NULL,
  `username` varchar(45) NOT NULL,
  `password` text NOT NULL,
  `active` int(11) NOT NULL,
  `last_login` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `ms_user`
--

INSERT INTO `ms_user` (`id`, `username`, `password`, `active`, `last_login`, `created`, `updated`) VALUES
(1, 'arsene', '$2b$10$g2JUB.o4DOa35dotoKPpOebh.6C0qrD1KUtKW9tLsF6AhefUt.KrG', 1, '2021-07-17 11:06:01', '2021-07-17 11:06:01', '2021-07-17 11:06:01'),
(2, 'lawvia', '$2a$14$yDQ/FP6x9E1f2jQnNRR6NuN7fe21uKxS7/bVA9cRrY.gBYvYKp40G', 1, '2021-07-18 03:03:04', '2021-07-18 03:03:04', '2021-07-18 03:03:04');

-- --------------------------------------------------------

--
-- Struktur dari tabel `ms_whitelist`
--

CREATE TABLE `ms_whitelist` (
  `id` int(11) NOT NULL,
  `name` varchar(45) NOT NULL,
  `ip` varchar(45) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Struktur dari tabel `tr_detail_purchase`
--

CREATE TABLE `tr_detail_purchase` (
  `id` int(11) NOT NULL,
  `id_items` int(11) NOT NULL,
  `qty` int(11) NOT NULL,
  `header_key` varchar(45) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `tr_detail_purchase`
--

INSERT INTO `tr_detail_purchase` (`id`, `id_items`, `qty`, `header_key`, `created`, `updated`) VALUES
(6, 1, 7, '58449c2aebe9dd83', '2021-07-18 16:20:34', '2021-07-18 16:20:34'),
(7, 3, 1, '58449c2aebe9dd83', '2021-07-18 16:20:34', '2021-07-18 16:20:34');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tr_header_purchase`
--

CREATE TABLE `tr_header_purchase` (
  `id` int(11) NOT NULL,
  `purchase_key` varchar(45) NOT NULL,
  `id_user` int(11) NOT NULL,
  `address` text NOT NULL,
  `total` varchar(45) NOT NULL,
  `status` int(11) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `tr_header_purchase`
--

INSERT INTO `tr_header_purchase` (`id`, `purchase_key`, `id_user`, `address`, `total`, `status`, `created`, `updated`) VALUES
(4, '58449c2aebe9dd83', 1, 'Jl. KS Tubun IIa', '1270000', 3, '2021-07-18 16:20:34', '2021-07-18 16:20:34');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tr_user_cart`
--

CREATE TABLE `tr_user_cart` (
  `id` int(11) NOT NULL,
  `id_items` int(11) NOT NULL,
  `id_user` int(11) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `tr_user_cart`
--

INSERT INTO `tr_user_cart` (`id`, `id_items`, `id_user`, `created`, `updated`) VALUES
(12, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(13, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(14, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(15, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(16, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(17, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(18, 1, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21'),
(19, 3, 1, '2021-07-18 16:21:21', '2021-07-18 16:21:21');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `ms_items`
--
ALTER TABLE `ms_items`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_user`
--
ALTER TABLE `ms_user`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `ms_whitelist`
--
ALTER TABLE `ms_whitelist`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tr_detail_purchase`
--
ALTER TABLE `tr_detail_purchase`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tr_header_purchase`
--
ALTER TABLE `tr_header_purchase`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tr_user_cart`
--
ALTER TABLE `tr_user_cart`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `ms_items`
--
ALTER TABLE `ms_items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `ms_user`
--
ALTER TABLE `ms_user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `ms_whitelist`
--
ALTER TABLE `ms_whitelist`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `tr_detail_purchase`
--
ALTER TABLE `tr_detail_purchase`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `tr_header_purchase`
--
ALTER TABLE `tr_header_purchase`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `tr_user_cart`
--
ALTER TABLE `tr_user_cart`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
