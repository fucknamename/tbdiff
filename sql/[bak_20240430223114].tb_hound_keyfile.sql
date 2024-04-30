-- MySQL dump 10.13  Distrib 5.5.27, for Win64 (x86)
--
-- Host: 127.0.0.1    Database: hound
-- ------------------------------------------------------
-- Server version	5.5.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tb_hound_keyfile`
--

DROP TABLE IF EXISTS `tb_hound_keyfile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_hound_keyfile` (
  `id` bigint(20) NOT NULL COMMENT '记录id',
  `gid` bigint(20) DEFAULT NULL COMMENT '关键词分组',
  `path` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件路径',
  `count` int(8) DEFAULT NULL COMMENT '关键词数',
  `catetory` tinyint(2) DEFAULT NULL COMMENT '1核心词，2前缀词，3尾缀词',
  `remark` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注信息',
  `state` tinyint(2) DEFAULT '1' COMMENT '启用状态：1启用，2禁用',
  `created_at` bigint(13) NOT NULL COMMENT '创建时间',
  `updated_at` bigint(13) DEFAULT NULL COMMENT '更新时间',
  `deleted_at` bigint(13) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='关键词库';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_hound_keyfile`
--

LOCK TABLES `tb_hound_keyfile` WRITE;
/*!40000 ALTER TABLE `tb_hound_keyfile` DISABLE KEYS */;
INSERT INTO `tb_hound_keyfile` VALUES (355919432907520,355903616236288,'./data/key/1709371190187_ty.txt',0,1,'ty110',1,1709371192799,1712578898,0),(1777311025306013696,355903616236288,'./data/key/1712578937516_采集词4.txt',0,1,'ty1102',1,1712578943932,1712578943,0);
/*!40000 ALTER TABLE `tb_hound_keyfile` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-04-30 22:31:14
