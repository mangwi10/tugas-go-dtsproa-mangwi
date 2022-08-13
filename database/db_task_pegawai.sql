-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 13, 2022 at 05:52 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_task_pegawai`
--

-- --------------------------------------------------------

--
-- Table structure for table `tb_task`
--

CREATE TABLE `tb_task` (
  `id_task` int(11) NOT NULL,
  `deskripsi` text DEFAULT NULL,
  `pegawai` varchar(100) DEFAULT NULL,
  `deadline` varchar(150) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tb_task`
--

INSERT INTO `tb_task` (`id_task`, `deskripsi`, `pegawai`, `deadline`) VALUES
(1, 'Memperbaiki Wifi', 'nur', '18/12/2022'),
(2, 'Memperbaiki laptop FO', 'adi', '19/02/2022'),
(4, 'Memperbaiki CCTV, TV, AC', 'Mes', '01/11/2022'),
(5, 'memperbaiki remote', 'Agus', '21/06/2022'),
(6, 'Memperbaik Ac', 'Made', '23/06/2022'),
(7, 'Memasang Wifi', 'ketut', '12/09/2022'),
(8, 'Memasang kabel lan', 'komag', '12/09/2022');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tb_task`
--
ALTER TABLE `tb_task`
  ADD PRIMARY KEY (`id_task`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `tb_task`
--
ALTER TABLE `tb_task`
  MODIFY `id_task` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
