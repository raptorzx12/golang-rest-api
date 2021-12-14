-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 14, 2021 at 07:50 AM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.0.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pirate`
--

-- --------------------------------------------------------

--
-- Table structure for table `pirates`
--

CREATE TABLE `pirates` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `devil_fruit` varchar(255) DEFAULT NULL,
  `crew` varchar(255) DEFAULT NULL,
  `job` varchar(255) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `pirates`
--

INSERT INTO `pirates` (`id`, `name`, `devil_fruit`, `crew`, `job`, `user_id`) VALUES
(1, 'Eustass Captain Kid', 'Jiki Jiki no Mi', 'Kid Pirates', 'Captain', 1),
(2, 'Monkey D. Luffy', 'Gomu Gomu no Mi', 'Straw Hat Pirates', 'Captain', 1),
(3, 'Trafalgar D. Water Law', 'Ope Ope no Mi', 'Heart Pirates', 'Captain', 1),
(4, 'Roronoa Zoro', '', 'Straw Hat Pirates', 'Swordsman', 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(1, 'a', 'a@gmail.com', '$2a$04$/Hrmoirb.Xy9uUSTBvP6Tukypf2vbGiM/WZmAjM5f04RKUJsSpqza'),
(2, 'b', 'b@gmail.com', '$2a$04$NdDgZciTCH7mqnsNpL9VPOgQTGz40i.Rhy0fqaqIueKkHTSMJBGuK');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `pirates`
--
ALTER TABLE `pirates`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_users_pirates` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_users_email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pirates`
--
ALTER TABLE `pirates`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `pirates`
--
ALTER TABLE `pirates`
  ADD CONSTRAINT `fk_users_pirates` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
