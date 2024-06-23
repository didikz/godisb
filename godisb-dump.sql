
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `godisb` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `godisb`;
DROP TABLE IF EXISTS `bank_accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bank_accounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `bank` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `account_number` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `account_name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `bank_accounts_account_number_IDX` (`account_number`,`bank`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `bank_accounts` WRITE;
/*!40000 ALTER TABLE `bank_accounts` DISABLE KEYS */;
INSERT INTO `bank_accounts` VALUES (1,'bca','12345678','John Doess',1),(2,'mandiri','87654321','Jane Doe',0),(3,'mandiri','87654322','Mike Myers',1);
/*!40000 ALTER TABLE `bank_accounts` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `disbursements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `disbursements` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `account_number` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `bank` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `amount` bigint unsigned NOT NULL,
  `remark` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` int unsigned NOT NULL,
  `beneficiary_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `failed_notes` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` bigint NOT NULL,
  `failed_at` bigint DEFAULT NULL,
  `completed_at` bigint DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `idempotency_key` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `disbursements_user_id_IDX` (`user_id`) USING BTREE,
  KEY `disbursements_bank_IDX` (`bank`,`account_number`) USING BTREE,
  KEY `disbursements_status_IDX` (`status`) USING BTREE,
  KEY `disbursements_created_at_IDX` (`created_at`) USING BTREE,
  KEY `disbursements_idempotency_key_IDX` (`idempotency_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `disbursements` WRITE;
/*!40000 ALTER TABLE `disbursements` DISABLE KEYS */;
INSERT INTO `disbursements` VALUES (1,'12345678','bca',100000,'payroll',1,'John Doe',NULL,1719110106,NULL,1719110106,1,'ik-123456'),(2,'12345678','bca',10000,'lorem ipsum',0,'John Doess',NULL,1719113594,NULL,NULL,1,'ik-123457'),(3,'12345678','bca',10000,'lorem ipsum',0,'John Doess',NULL,1719119489,NULL,NULL,1,'ik-123458'),(4,'12345678','bca',10000,'lorem ipsum',0,'John Doess',NULL,1719119662,NULL,NULL,1,'ik-123459'),(5,'12345678','bca',10000,'lorem ipsum',0,'John Doess',NULL,1719119774,NULL,NULL,1,'ik-123460'),(6,'12345678','bca',10000,'lorem ipsum',1,'John Doess',NULL,1719120010,NULL,1719120012,1,'ik-123461'),(7,'87654322','mandiri',10000,'lorem ipsum',1,'Mike Myers',NULL,1719120192,NULL,1719120194,1,'ik-123462'),(8,'87654322','mandiri',10000,'lorem ipsum',0,'Mike Myers',NULL,1719120389,NULL,NULL,1,'ik-123463'),(9,'87654322','mandiri',10000,'lorem ipsum',2,'Mike Myers','0',1719121148,1719121150,NULL,1,'ik-123464'),(10,'87654322','mandiri',10000,'lorem ipsum',2,'Mike Myers','0',1719122092,1719122094,NULL,1,'ik-123465'),(11,'87654322','mandiri',10000,'lorem ipsum',2,'Mike Myers','0',1719122452,1719122453,NULL,1,'ik-123466'),(12,'87654322','mandiri',10000,'lorem ipsum',2,'Mike Myers','0',1719122551,1719122552,NULL,1,'ik-123467'),(13,'12345678','bca',10000,'lorem ipsum',1,'John Doess',NULL,1719122586,NULL,1719122587,1,'ik-123468');
/*!40000 ALTER TABLE `disbursements` ENABLE KEYS */;
UNLOCK TABLES;
DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `balance` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'rezd14@gmail.com','secret',10010000);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

