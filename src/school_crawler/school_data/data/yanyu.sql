-- MySQL dump 10.13  Distrib 5.7.16, for Linux (x86_64)
--
-- Host: localhost    Database: yanyu
-- ------------------------------------------------------
-- Server version	5.7.16-0ubuntu0.16.04.1

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
-- Table structure for table `yanyu_answer`
--

DROP TABLE IF EXISTS `yanyu_answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_answer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `question_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `content` text,
  `answer_date` datetime DEFAULT NULL,
  `read_count` int(11) DEFAULT '0',
  `like_count` int(11) DEFAULT '0',
  `fee` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_answer`
--

LOCK TABLES `yanyu_answer` WRITE;
/*!40000 ALTER TABLE `yanyu_answer` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_answer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_like_answer`
--

DROP TABLE IF EXISTS `yanyu_like_answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_like_answer` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `answer_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `like_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_like_answer`
--

LOCK TABLES `yanyu_like_answer` WRITE;
/*!40000 ALTER TABLE `yanyu_like_answer` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_like_answer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_like_share`
--

DROP TABLE IF EXISTS `yanyu_like_share`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_like_share` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `share_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `like_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_like_share`
--

LOCK TABLES `yanyu_like_share` WRITE;
/*!40000 ALTER TABLE `yanyu_like_share` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_like_share` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_major`
--

DROP TABLE IF EXISTS `yanyu_major`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_major` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `major` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_major`
--

LOCK TABLES `yanyu_major` WRITE;
/*!40000 ALTER TABLE `yanyu_major` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_major` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_question`
--

DROP TABLE IF EXISTS `yanyu_question`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_question` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `tag_universities` varchar(45) DEFAULT NULL,
  `tag_majors` varchar(45) DEFAULT NULL,
  `content` text,
  `ask_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_question`
--

LOCK TABLES `yanyu_question` WRITE;
/*!40000 ALTER TABLE `yanyu_question` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_question` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_share`
--

DROP TABLE IF EXISTS `yanyu_share`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_share` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `tag_universities` varchar(45) DEFAULT NULL,
  `tag_majors` varchar(45) DEFAULT NULL,
  `content` text,
  `post_date` datetime DEFAULT NULL,
  `fee` float DEFAULT NULL,
  `read_count` int(11) DEFAULT NULL,
  `like_count` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_share`
--

LOCK TABLES `yanyu_share` WRITE;
/*!40000 ALTER TABLE `yanyu_share` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_share` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_university`
--

DROP TABLE IF EXISTS `yanyu_university`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_university` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `location` varchar(45) DEFAULT NULL,
  `subjection` varchar(45) DEFAULT NULL,
  `attribution` varchar(45) DEFAULT NULL,
  `graduate_school` varchar(1) DEFAULT '0',
  `self_decision` varchar(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_university`
--

LOCK TABLES `yanyu_university` WRITE;
/*!40000 ALTER TABLE `yanyu_university` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_university` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `yanyu_user`
--

DROP TABLE IF EXISTS `yanyu_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `yanyu_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nickname` varchar(45) DEFAULT NULL,
  `telephone` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `university` varchar(45) DEFAULT NULL,
  `major` varchar(45) DEFAULT NULL,
  `avater_url` varchar(45) DEFAULT NULL,
  `avater_type` varchar(45) DEFAULT NULL,
  `user_state` varchar(45) DEFAULT NULL,
  `is_lock` varchar(1) DEFAULT '0',
  `register_date` datetime DEFAULT NULL,
  `auth_url` varchar(45) DEFAULT NULL,
  `is_auth` varchar(1) DEFAULT '0',
  `real_name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `yanyu_user`
--

LOCK TABLES `yanyu_user` WRITE;
/*!40000 ALTER TABLE `yanyu_user` DISABLE KEYS */;
/*!40000 ALTER TABLE `yanyu_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-11-19 16:21:39
