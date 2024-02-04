-- 非 freepbx 环境要额外建的表
-- 如果按照书 <Asterisk权威指南> 的方法，只有 26 张表, 为了两个环境都能适配, 还要加一点表
--

DROP TABLE IF EXISTS devices;
CREATE TABLE `devices` (
  `id` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `tech` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `dial` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `devicetype` varchar(5) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `user` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `emergency_cid` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `hint_override` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  KEY `id` (`id`),
  KEY `tech` (`tech`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
