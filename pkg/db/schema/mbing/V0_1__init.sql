/**  Price  **/
CREATE TABLE IF NOT EXISTS attribute_name (
	attribute_name_id VARCHAR(50)  NOT NULL UNIQUE,
	name              VARCHAR(255) NOT NULL,

	status            VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	create_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	description       TEXT,
	PRIMARY KEY (attribute_name_id)
);


CREATE TABLE IF NOT EXISTS attribute_unit (
	attribute_unit_id VARCHAR(50) NOT NULL UNIQUE,
	name              VARCHAR(30) NOT NULL,
	status            VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	create_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (attribute_unit_id)
);


CREATE TABLE IF NOT EXISTS attribute (
	attribute_id      VARCHAR(50)  NOT NULL UNIQUE,
	attribute_name_id VARCHAR(50)  NOT NULL,
	attribute_unit_id VARCHAR(50),
	value             VARCHAR(255) NOT NULL
	COMMENT 'attribute value, the types: single int value, scope of value (min_value, max_value], string value',
	create_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	status            VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	PRIMARY KEY (attribute_id)
);


CREATE TABLE IF NOT EXISTS spu (
	spu_id                      VARCHAR(50) NOT NULL UNIQUE,
	resource_version_id         VARCHAR(50) NOT NULL,
	attribute_name_ids          JSON COMMENT 'sku attribute_name ids(Do not include metering_attribute_name_ids)',
	metering_attribute_name_ids JSON        NOT NULL
	COMMENT 'the attribute_name ids need to metering and billing',
	create_time                 TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time                 TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	status                      VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	PRIMARY KEY (spu_id)
);


CREATE TABLE IF NOT EXISTS sku (
	sku_id                 VARCHAR(50) NOT NULL UNIQUE,
	spu_id                 VARCHAR(50) NOT NULL,
	attribute_ids          JSON COMMENT 'sku attributes with value of spu.',
	metering_attribute_ids JSON        NOT NULL,
	create_time            TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time            TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	status                 VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	PRIMARY KEY (sku_id)
);


CREATE TABLE IF NOT EXISTS price (
	price_id     VARCHAR(50) NOT NULL UNIQUE,
	sku_id       VARCHAR(50) NOT NULL,
	attribute_id VARCHAR(50) NOT NULL,
	prices       JSON COMMENT '{upto: price1, ...}',
	currency     VARCHAR(10) NOT NULL  DEFAULT 'cny',
	start_time   TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
	end_time     TIMESTAMP   NULL,
	create_time  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
	status_time  TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	status       VARCHAR(16)           DEFAULT 'active'
	COMMENT 'active, deleted, disabled',
	INDEX price_sku_index (sku_id, price_id),
	PRIMARY KEY (price_id)
);


/**  Metering  **/
CREATE TABLE IF NOT EXISTS leasing (
	leasing_id           VARCHAR(50) NOT NULL UNIQUE,
	group_id             VARCHAR(50) NOT NULL,
	user_id              VARCHAR(50) NOT NULL,
	resource_id          VARCHAR(50) NOT NULL
	COMMENT 'the same as cluster_id',
	sku_id               VARCHAR(50) NOT NULL,
	sku_type             VARCHAR(16) NOT NULL DEFAULT 'sku',
	other_info           VARCHAR(50) COMMENT 'used for distinguish when resource_id and sku_id are same with others',
	metering_values      JSON COMMENT 'the values of metering_attributes, {att_id: value, ..}',
	lease_time           TIMESTAMP   NULL,
	update_duration_time TIMESTAMP   NULL,
	renewal_time         TIMESTAMP   NULL,
	create_time          TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
	status_time          TIMESTAMP            DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	close_time           JSON COMMENT '[{close_time: restart_time}, ..]',
	status               VARCHAR(16)          DEFAULT 'active'
	COMMENT 'active, handClosed, forceClosed',
	PRIMARY KEY (leasing_id)
);


CREATE TABLE IF NOT EXISTS leased (
	leased_id       VARCHAR(50) NOT NULL UNIQUE,
	group_id        VARCHAR(50) NOT NULL,
	user_id         VARCHAR(50) NOT NULL,
	resource_id     VARCHAR(50) NOT NULL
	COMMENT 'the same as cluster_id',
	sku_id          VARCHAR(50) NOT NULL,
	other_info      VARCHAR(50) COMMENT 'used for distinguish when resource_id and sku_id are same with others',
	metering_values JSON COMMENT 'the values of metering_attributes, {att_id: value, ..}',
	lease_time      TIMESTAMP   NULL,
	end_time        TIMESTAMP   NULL,
	create_time     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	close_time      JSON COMMENT '[{close_time: restart_time}, ..]',
	PRIMARY KEY (leased_id)
);


/** Billing **/
CREATE TABLE IF NOT EXISTS leasing_contract (
	id              VARCHAR(50)   NOT NULL UNIQUE,
	leasing_id      VARCHAR(50)   NOT NULL,
	sku_id          VARCHAR(50)   NOT NULL,
	user_id         VARCHAR(50)   NOT NULL,
	metering_values JSON          NOT NULL,
	start_time      TIMESTAMP     NULL
	COMMENT 'same as leasing_time',
	status_time     TIMESTAMP     NULL,
	create_time     TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	fee_info        TEXT,
	fee             DECIMAL(8, 2) NOT NULL
	COMMENT 'total fee from starting cluster to now',
	due_fee         DECIMAL(8, 2) NOT NULL default 0,
	done_fee        DECIMAL(8, 2) NOT NULL default 0,
	before_bill_fee DECIMAL(8, 2) NOT NULL DEFAULT 0
	COMMENT 'the total fee of the before bills ',
	currency        VARCHAR(10)   NOT NULL DEFAULT 'cny',
	remark          TEXT,
	PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS leased_contract (
	contract_id     VARCHAR(50)   NOT NULL UNIQUE,
	leasing_id      VARCHAR(50)   NOT NULL,
	sku_id          VARCHAR(50)   NOT NULL,
	user_id         VARCHAR(50)   NOT NULL,
	metering_values JSON          NOT NULL,
	start_time      TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	end_time        TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	create_time     TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	fee_info        TEXT,
	fee             DECIMAL(8, 2) NOT NULL
	COMMENT 'total fee from starting cluster to now',
	due_fee         DECIMAL(8, 2) NOT NULL default 0,
	done_fee        DECIMAL(8, 2) NOT NULL default 0,
	before_bill_fee DECIMAL(8, 2) NOT NULL DEFAULT 0
	COMMENT 'the total fee of the before bills ',
	currency        VARCHAR(10)   NOT NULL DEFAULT 'cny',
	remark          TEXT,
	PRIMARY KEY (contract_id)
);


/** Charge **/
CREATE TABLE IF NOT EXISTS charge (
	id          VARCHAR(50)   NOT NULL UNIQUE,
	user_id     VARCHAR(50)   NOT NULL,
	contract_id VARCHAR(50)   NOT NULL,
	fee         DECIMAL(8, 2) NOT NULL
	COMMENT 'total fee from starting cluster to now',
	currency    VARCHAR(10)   NOT NULL DEFAULT 'cny',
	info        JSON COMMENT '{couponReceivedID: fee}',
	status      VARCHAR(16)            DEFAULT 'successful'
	COMMENT 'successful, failed',
	create_time TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	remark      TEXT,
	PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS recharge (
	id          VARCHAR(50)   NOT NULL UNIQUE,
	user_id     VARCHAR(50)   NOT NULL,
	contract_id VARCHAR(50),
	fee         DECIMAL(8, 2) NOT NULL
	COMMENT 'total fee from starting cluster to now',
	currency    VARCHAR(10)   NOT NULL DEFAULT 'cny',
	info        JSON COMMENT '{couponReceivedID: fee}',
	status      VARCHAR(16)            DEFAULT 'successful'
	COMMENT 'successful, failed',
	operator    VARCHAR(50)   NOT NULL DEFAULT 'system',
	create_time TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	remark      TEXT,
	PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS income (
	id                  VARCHAR(50)   NOT NULL UNIQUE,
	user_id             VARCHAR(50)   NOT NULL,
	contract_id         VARCHAR(50)   NOT NULL,
	resource_version_id VARCHAR(50)   NOT NULL,
	balance             DECIMAL(9, 2) NOT NULL DEFAULT 0,
	currency            VARCHAR(10)   NOT NULL DEFAULT 'cny',
	create_time         TIMESTAMP              DEFAULT CURRENT_TIMESTAMP,
	remark              TEXT,
	PRIMARY KEY (id)
);


CREATE TABLE IF NOT EXISTS account (
	user_id            VARCHAR(50)   NOT NULL UNIQUE,
	user_type          TINYINT       NOT NULL DEFAULT 1
	COMMENT '0: deleted, 1: normal',
	balance            DECIMAL(9, 2) NOT NULL DEFAULT 0,
	income             JSON COMMENT '{cny: balance, ..}',
	currency           VARCHAR(10)   NOT NULL DEFAULT 'cny',
	credit_mode        TINYINT       NOT NULL DEFAULT 0
	COMMENT '0: close, 1: open',
	credit_amount      DECIMAL(8, 2),
	credit_duration    INT COMMENT 'unit: hour',
	first_in_debt_time TIMESTAMP     NULL,
	PRIMARY KEY (user_id)
);

#Init data about duration
INSERT INTO attribute_name (attribute_name_id, name, description)
VALUES ("att-name-000001", "duration", "default attribute: duration");

INSERT INTO attribute_unit (attribute_unit_id, name)
VALUES ("att-unit-000001", "hour"), ("att-unit-000002", "month"), ("att-unit-000003", "year");

INSERT INTO attribute (attribute_id, attribute_name_id, attribute_unit_id, value)
VALUES ("att-000001", "att-name-000001", "att-unit-000001", 1), ("att-000002", "att-name-000001", "att-unit-000002", 1),
	("att-000003", "att-name-000001", "att-unit-000003", 1);