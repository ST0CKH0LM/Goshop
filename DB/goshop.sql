-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 08, 2022 at 07:13 PM
-- Server version: 10.4.19-MariaDB
-- PHP Version: 7.3.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `goshop`
--

-- --------------------------------------------------------

--
-- Table structure for table `db_category`
--

CREATE TABLE `db_category` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `gender` varchar(255) NOT NULL,
  `color` varchar(255) NOT NULL,
  `pattern` varchar(255) NOT NULL,
  `size` varchar(255) NOT NULL,
  `price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `db_category`
--

INSERT INTO `db_category` (`id`, `name`, `gender`, `color`, `pattern`, `size`, `price`) VALUES
(1, 'test', 'male', 'red', 'pat1', 'xs', 200),
(2, 'test2', 'female', 'blue', 'pat2', 'm', 270),
(3, 'test3', 'male', 'red', 'pat1', 'xs', 270),
(4, 'ddd', 'male', 'red', 'pat1', 'xs', 100),
(5, 'test3', 'male', 'red', 'pat1', 'xs', 111),
(6, 'sddd', 'male', 'red', 'pat1', 'xs', 100),
(7, 'test3', 'male', 'red', 'pat1', 'xs', 270),
(8, 'test2', 'male', 'red', 'pat1', 'xs', 111),
(9, 'test3', 'male', 'red', 'pat2', 'xl', 111),
(10, 'sddd', 'male', 'red', 'pat1', 'xs', 200),
(11, 'sddd', 'male', 'red', 'pat1', 'xs', 111),
(12, 'name', 'male', 'red', 'pat1', 'xs', 111);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `db_category`
--
ALTER TABLE `db_category`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `db_category`
--
ALTER TABLE `db_category`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
