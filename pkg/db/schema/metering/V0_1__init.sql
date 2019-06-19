/**  Sku  **/
CREATE TABLE IF NOT EXISTS attribute_term (
	attribute_term_id VARCHAR(50)  NOT NULL UNIQUE,
	name              VARCHAR(255) NOT NULL,
	type              VARCHAR(16)  NOT NULL DEFAULT 'normal'
	COMMENT 'normal, metering',
	status            VARCHAR(16)           DEFAULT 'active'
	COMMENT 'active, deleted',
	create_time       TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
	status_time       TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	description       TEXT,
	PRIMARY KEY (attribute_term_id)
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
	attribute_term_id VARCHAR(50)  NOT NULL,
	attribute_unit_id VARCHAR(50),
	value             VARCHAR(255) NOT NULL
	COMMENT 'attribute value, the types: single int value, scope of value (min_value, max_value], string value',
	range             JSON COMMENT 'used for can not enumerated attribute',
	owner             VARCHAR(50)  NOT NULL,
	status            VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	create_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time       TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (attribute_id)
);


CREATE TABLE IF NOT EXISTS spu (
	spu_id      VARCHAR(50) NOT NULL UNIQUE,
	product_id  VARCHAR(50) NOT NULL UNIQUE
	COMMENT 'product_id: app_id/app_version_id/other_ids..',
	owner       VARCHAR(50) NOT NULL,
	status      VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	create_time TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (spu_id)
);


CREATE TABLE IF NOT EXISTS sku (
	sku_id                 VARCHAR(50) NOT NULL UNIQUE,
	spu_id                 VARCHAR(50) NOT NULL,
	attribute_ids          JSON COMMENT 'attributes of sku',
	metering_attribute_ids JSON COMMENT 'metering attributes of sku',
	fee_policy             VARCHAR(50) NOT NULL,
	status                 VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, deleted',
	create_time            TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time            TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (sku_id)
);


/**  Metering  **/
CREATE TABLE IF NOT EXISTS leasing (
	leasing_id           VARCHAR(50) NOT NULL UNIQUE,
	user_id              VARCHAR(50) NOT NULL,
	resource_id          VARCHAR(50) NOT NULL,
	sku_id               VARCHAR(50) NOT NULL,
	metering_values      JSON COMMENT 'the values of metering_attributes, {att_id: value, ..}',
	lease_time           TIMESTAMP   NOT NULL,
	update_duration_time TIMESTAMP   NULL,
	renewal_time         TIMESTAMP   NULL,
	create_time          TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
	status_time          TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
	ON UPDATE CURRENT_TIMESTAMP,
	stop_times           JSON COMMENT '[{close_time: restart_time}, ..]',
	status               VARCHAR(16) DEFAULT 'active'
	COMMENT 'active, handStop, forceStop, terminate',
	PRIMARY KEY (leasing_id)
);


CREATE TABLE IF NOT EXISTS leased (
	leased_id       VARCHAR(50) NOT NULL UNIQUE,
	user_id         VARCHAR(50) NOT NULL,
	resource_id     VARCHAR(50) NOT NULL
	COMMENT 'the same as cluster_id',
	sku_id          VARCHAR(50) NOT NULL,
	metering_values JSON COMMENT 'the values of metering_attributes, {att_id: value, ..}',
	lease_time      TIMESTAMP   NULL,
	end_time        TIMESTAMP   NULL,
	create_time     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	close_time      JSON COMMENT '[{close_time: restart_time}, ..]',
	PRIMARY KEY (leased_id)
);

#Init data about duration
INSERT INTO attribute_term (attribute_term_id, name, description)
VALUES ("att-term-000001", "duration", "default attribute: duration");

INSERT INTO attribute_unit (attribute_unit_id, name)
VALUES ("att-unit-000001", "hour"), ("att-unit-000002", "month"), ("att-unit-000003", "year");

INSERT INTO attribute (attribute_id, attribute_term_id, attribute_unit_id, value)
VALUES ("att-000001", "att-name-000001", "att-unit-000001", 1), ("att-000002", "att-name-000001", "att-unit-000002", 1),
	("att-000003", "att-name-000001", "att-unit-000003", 1);