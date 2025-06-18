/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES  */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for function public.uuid_generate_v1
DELIMITER //
CREATE FUNCTION "uuid_generate_v1"() RETURNS UUID AS $$ uuid_generate_v1 $$//
DELIMITER ;

-- Dumping structure for function public.uuid_generate_v1mc
DELIMITER //
CREATE FUNCTION "uuid_generate_v1mc"() RETURNS UUID AS $$ uuid_generate_v1mc $$//
DELIMITER ;

-- Dumping structure for function public.uuid_generate_v3
DELIMITER //
CREATE FUNCTION "uuid_generate_v3"(namespace UUID, name TEXT) RETURNS UUID AS $$ uuid_generate_v3 $$//
DELIMITER ;

-- Dumping structure for function public.uuid_generate_v4
DELIMITER //
CREATE FUNCTION "uuid_generate_v4"() RETURNS UUID AS $$ uuid_generate_v4 $$//
DELIMITER ;

-- Dumping structure for function public.uuid_generate_v5
DELIMITER //
CREATE FUNCTION "uuid_generate_v5"(namespace UUID, name TEXT) RETURNS UUID AS $$ uuid_generate_v5 $$//
DELIMITER ;

-- Dumping structure for function public.uuid_nil
DELIMITER //
CREATE FUNCTION "uuid_nil"() RETURNS UUID AS $$ uuid_nil $$//
DELIMITER ;

-- Dumping structure for function public.uuid_ns_dns
DELIMITER //
CREATE FUNCTION "uuid_ns_dns"() RETURNS UUID AS $$ uuid_ns_dns $$//
DELIMITER ;

-- Dumping structure for function public.uuid_ns_oid
DELIMITER //
CREATE FUNCTION "uuid_ns_oid"() RETURNS UUID AS $$ uuid_ns_oid $$//
DELIMITER ;

-- Dumping structure for function public.uuid_ns_url
DELIMITER //
CREATE FUNCTION "uuid_ns_url"() RETURNS UUID AS $$ uuid_ns_url $$//
DELIMITER ;

-- Dumping structure for function public.uuid_ns_x500
DELIMITER //
CREATE FUNCTION "uuid_ns_x500"() RETURNS UUID AS $$ uuid_ns_x500 $$//
DELIMITER ;

-- Dumping structure for table public.consumers
CREATE TABLE IF NOT EXISTS "consumers" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"nik" VARCHAR(20) NOT NULL,
	"full_name" VARCHAR(100) NOT NULL,
	"legal_name" VARCHAR(100) NOT NULL,
	"birth_place" VARCHAR(50) NULL DEFAULT NULL,
	"birth_date" DATE NULL DEFAULT NULL,
	"salary" BIGINT NULL DEFAULT NULL,
	"photo_ktp" TEXT NULL DEFAULT NULL,
	"photo_selfie" TEXT NULL DEFAULT NULL,
	"created_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	UNIQUE INDEX "consumers_nik_key" ("nik")
);

-- Dumping data for table public.consumers: 1 rows
/*!40000 ALTER TABLE "consumers" DISABLE KEYS */;
INSERT INTO "consumers" ("id", "nik", "full_name", "legal_name", "birth_place", "birth_date", "salary", "photo_ktp", "photo_selfie", "created_at", "updated_at") VALUES
	('2192643c-4dfb-431c-a6dd-1a2a842cbccd', '1234567890123456', 'Budi Santoso', 'Budi Santoso', 'Jakarta', '1990-01-15', 7500000, 'https://example.com/ktp/budi_santoso_ktp.jpg', 'https://example.com/selfie/budi_santoso_selfie.jpg', '2025-06-17 17:49:04.306463', '2025-06-17 17:49:04.306463'),
	('49993dbc-6ab6-4dff-894a-0645d26d1224', '2345678901234567', 'Rina Wulandari', 'Rina Wulandari', 'Yogyakarta', '1992-03-22', 6800000, 'https://example.com/ktp/rina_wulandari_ktp.jpg', 'https://example.com/selfie/rina_wulandari_selfie.jpg', '2025-06-17 22:18:01.729016', '2025-06-17 22:18:01.729016');
/*!40000 ALTER TABLE "consumers" ENABLE KEYS */;

-- Dumping structure for table public.consumer_limits
CREATE TABLE IF NOT EXISTS "consumer_limits" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"consumer_id" UUID NULL DEFAULT NULL,
	"tenor_month" INTEGER NULL DEFAULT NULL,
	"max_limit" BIGINT NOT NULL,
	"used_limit" BIGINT NULL DEFAULT 0,
	"created_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	CONSTRAINT "consumer_limits_consumer_id_fkey" FOREIGN KEY ("consumer_id") REFERENCES "consumers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT "consumer_limits_tenor_month_check" CHECK (((tenor_month = ANY (ARRAY[1, 2, 3, 6]))))
);

-- Dumping data for table public.consumer_limits: -1 rows
/*!40000 ALTER TABLE "consumer_limits" DISABLE KEYS */;
/*!40000 ALTER TABLE "consumer_limits" ENABLE KEYS */;

-- Dumping structure for table public.transactions
CREATE TABLE IF NOT EXISTS "transactions" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"contract_number" VARCHAR(50) NOT NULL,
	"consumer_id" UUID NULL DEFAULT NULL,
	"tenor_month" INTEGER NULL DEFAULT NULL,
	"otr" BIGINT NOT NULL,
	"admin_fee" BIGINT NOT NULL,
	"installment" BIGINT NOT NULL,
	"interest" BIGINT NOT NULL,
	"asset_name" VARCHAR(100) NOT NULL,
	"created_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	UNIQUE INDEX "transactions_contract_number_key" ("contract_number"),
	CONSTRAINT "transactions_consumer_id_fkey" FOREIGN KEY ("consumer_id") REFERENCES "consumers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
	CONSTRAINT "transactions_tenor_month_check" CHECK (((tenor_month = ANY (ARRAY[1, 2, 3, 6]))))
);

-- Dumping data for table public.transactions: -1 rows
/*!40000 ALTER TABLE "transactions" DISABLE KEYS */;
/*!40000 ALTER TABLE "transactions" ENABLE KEYS */;

-- Dumping structure for table public.users
CREATE TABLE IF NOT EXISTS "users" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"username" VARCHAR(255) NOT NULL,
	"email" VARCHAR(255) NOT NULL,
	"password" TEXT NOT NULL,
	"created_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	UNIQUE INDEX "users_email_key" ("email")
);

-- Dumping data for table public.users: 0 rows
/*!40000 ALTER TABLE "users" DISABLE KEYS */;
INSERT INTO "users" ("id", "username", "email", "password", "created_at", "updated_at") VALUES
	('6ab1a490-8fb0-42b8-897e-7712278dc206', 'alfahan', 'ali.farhan160@gmail.com', '$2a$10$vQpUeuKdIKG/.jybkmmVPeZ5ZG1v/t.ehK63E1wmfxR/3amxYl80G', '0001-01-01 00:00:00', '0001-01-01 00:00:00');
/*!40000 ALTER TABLE "users" ENABLE KEYS */;

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
