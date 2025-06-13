-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1:3307
-- Thời gian đã tạo: Th6 12, 2025 lúc 06:31 AM
-- Phiên bản máy phục vụ: 10.4.28-MariaDB
-- Phiên bản PHP: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `test-tt-1`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(50) NOT NULL,
  `fullname` varchar(100) NOT NULL,
  `password` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `users`
--

INSERT INTO `users` (`id`, `username`, `fullname`, `password`) VALUES
(1, 'sơn', 'ldh sơn', 'pass123'),
(2, 'abc', 'abcxyz', 'pass123'),
(3, 'aaaaa', 'vvvvvv', 'pass123'),
(5, 'hẹ hẹ', 'aaaaaaaa', 'pass123'),
(6, 'user_6', 'User 6', 'pass123'),
(7, 'user_7', 'User 7', 'pass123'),
(8, 'user_8', 'User 8', 'pass123'),
(9, 'user_9', 'User 9', 'pass123'),
(10, 'user_10', 'User 10', 'pass123'),
(11, 'user_11', 'User 11', 'pass123'),
(13, 'user_13', 'User 13', 'pass123'),
(14, 'sonluu', 'sonluuldh', 'pass123'),
(15, 'user_15', 'User 15', 'pass123'),
(16, 'user_16', 'User 16', 'pass123'),
(17, 'user_17', 'User 17', 'pass123'),
(18, 'user_18', 'User 18', 'pass123'),
(19, 'user_19', 'User 19', 'pass123'),
(20, 'user_20', 'User 20', 'pass123'),
(21, 'user_21', 'User 21', 'pass123'),
(22, 'user_22', 'User 22', 'pass123'),
(23, 'user_23', 'User 23', 'pass123'),
(24, 'user_24', 'User 24', 'pass123'),
(25, 'user_25', 'User 25', 'pass123'),
(26, 'user_26', 'User 26', 'pass123'),
(27, 'user_27', 'User 27', 'pass123'),
(28, 'user_28', 'User 28', 'pass123'),
(29, 'user_29', 'User 29', 'pass123'),
(30, 'user_30', 'User 30', 'pass123'),
(31, 'user_31', 'User 31', 'pass123'),
(32, 'user_32', 'User 32', 'pass123'),
(33, 'user_33', 'User 33', 'pass123'),
(34, 'user_34', 'User 34', 'pass123'),
(35, 'user_35', 'User 35', 'pass123'),
(36, 'user_36', 'User 36', 'pass123'),
(37, 'user_37', 'User 37', 'pass123'),
(38, 'user_38', 'User 38', 'pass123'),
(39, 'user_39', 'User 39', 'pass123'),
(40, 'user_40', 'User 40', 'pass123'),
(41, 'a', 'a', '');

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `user_action_logs`
--

CREATE TABLE `user_action_logs` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `action` varchar(100) NOT NULL,
  `target_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Đang đổ dữ liệu cho bảng `user_action_logs`
--

INSERT INTO `user_action_logs` (`id`, `user_id`, `username`, `action`, `target_id`, `created_at`) VALUES
(1, 0, 'sơn', 'update_user', 5, '2025-06-10 17:42:29'),
(2, 0, 'abc', 'delete_user', 12, '2025-06-10 17:43:22'),
(3, 1, 'sơn', 'update_user', 14, '2025-06-10 18:08:47');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Chỉ mục cho bảng `user_action_logs`
--
ALTER TABLE `user_action_logs`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=42;

--
-- AUTO_INCREMENT cho bảng `user_action_logs`
--
ALTER TABLE `user_action_logs`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
