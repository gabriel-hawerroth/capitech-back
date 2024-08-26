/* Entity User */
create table users (
	id SERIAL UNIQUE NOT NULL,
	email VARCHAR(255) UNIQUE NOT NULL,
	password CHAR(60) NOT NULL,
	active BOOL NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create index user_email_idx on users(email);

/* Entity Category */
create table category (
	id SERIAL UNIQUE NOT NULL,
	description VARCHAR(40) NOT NULL
);

/* Entity Address */
create table address (
	id SERIAL UNIQUE NOT NULL,
	user_id INT4 NOT NULL,
	description VARCHAR(50),
	cep VARCHAR(9) NOT NULL,
	number NUMERIC(5) NOT NULL,
	complement VARCHAR(80),
	foreign key (user_id) references users(id)
);

/* Entity Product */
create table product (
	id SERIAL UNIQUE NOT NULL,
	name VARCHAR(80) NOT NULL,
	description text,
	value DECIMAL(8,2) NOT NULL,
	category_id INT4 NOT NULL,
	stock_quantity NUMERIC(5) NOT NULL,
	image BYTEA,
	foreign key (category_id) references category(id)
);

/* Entity ShoppingCart */
create table shopping_cart (
	id SERIAL UNIQUE NOT NULL,
	user_id INT4 NOT NULL,
	product_id INT4 NOT NULL,
	quantity NUMERIC(5) NOT NULL,
	foreign key (user_id) references users(id),
	foreign key (product_id) references product(id)
);

/* Entity Purchase */
create table purchase (
	id SERIAL UNIQUE NOT NULL,
	user_id INT4 NOT NULL,
	purchase_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	status VARCHAR(255) NOT NULL,
	payment_method VARCHAR(255) NOT NULL,
	address_id INT4 NOT NULL,
	shipping_value DECIMAL(7,2) NOT NULL,
	foreign key (user_id) references users(id),
	foreign key (address_id) references address(id)
);

/* Entity PurchaseItem */
create table purchase_item (
	id SERIAL UNIQUE NOT NULL,
	purchase_id INT4 NOT NULL,
	product_id INT4 NOT NULL,
	product_value DECIMAL(8,2) NOT NULL,
	quantity NUMERIC(5) NOT NULL,
	foreign key (purchase_id) references purchase(id),
	foreign key (product_id) references product(id)
);
